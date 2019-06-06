//package main
//
//import (
//	"database/sql"
//	"fmt"
//	"net/http"
//	"os"
//
//	_ "github.com/go-sql-driver/mysql"
//)
//
//func mainHandler() http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		_, _ = fmt.Fprintf(w, "Hello World! Ribbonwall! (Version info: %s, build date: %s)", os.Getenv("VERSION_INFO"), os.Getenv("BUILD_DATE"))
//	})
//}
//
//func dbTest() http.HandlerFunc {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		dbUser := os.Getenv("db_user")
//		dbPassword := os.Getenv("db_password")
//		dbName := os.Getenv("db_name")
//		dbEndpoint := os.Getenv("db_endpoint")
//
//		// Create the MySQL DNS string for the DB connection
//		// user:password@protocol(endpoint)/dbname?<params>
//		dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s",
//			dbUser, dbPassword, dbEndpoint, dbName,
//		)
//
//		// Use db to perform SQL operations on database
//		db, err := sql.Open("mysql", dnsStr)
//		defer db.Close()
//		if err != nil {
//			_, _ = fmt.Fprintf(w, "Error sql.Open %s", err.Error())
//			return
//		}
//
//		err = db.Ping()
//		if err != nil {
//			_, _ = fmt.Fprintf(w, "Error ping %s. ", err.Error())
//			return
//		}
//
//		rows, err := db.Query("SELECT * FROM competition_results")
//		if err != nil {
//			_, _ = fmt.Fprintf(w, "Error query %s", err.Error())
//			return
//		}
//
//		cols, err := rows.Columns()
//		if err != nil {
//			fmt.Println("Failed to get columns", err)
//			return
//		}
//
//		// Result is your slice string.
//		rawResult := make([][]byte, len(cols))
//		result := make([]string, len(cols))
//
//		dest := make([]interface{}, len(cols)) // A temporary interface{} slice
//		for i, _ := range rawResult {
//			dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
//		}
//
//		for rows.Next() {
//			err = rows.Scan(dest...)
//			if err != nil {
//				fmt.Println("Failed to scan row", err)
//				return
//			}
//
//			for i, raw := range rawResult {
//				if raw == nil {
//					result[i] = "\\N"
//				} else {
//					result[i] = string(raw)
//				}
//			}
//
//			fmt.Printf("%#v\n", result)
//		}
//
//		_, _ = fmt.Fprintf(w, "Successfully opened connection to database!\n Test results:\n%s", result)
//	})
//}
//
//func main() {
//	http.HandleFunc("/", mainHandler())
//	http.HandleFunc("/db", dbTest())
//	_ = http.ListenAndServe(":8080", nil)
//}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/ribbonwall/be_ribbonwall/endpoints"
	singletonConfig "github.com/ribbonwall/be_ribbonwall/singletons/config"
	singletonDatabase "github.com/ribbonwall/be_ribbonwall/singletons/db"
	logger "github.com/ribbonwall/common/logging"
)

type key int

const (
	requestIDKey key = 0
)

var (
	listenAddr string
	healthy    int32
)

func main() {
	config := singletonConfig.Get()
	flag.StringVar(&listenAddr, "listen-addr", ":8080", "server listen address")
	flag.Parse()

	logger.NewLogger()
	logger.Info("Server is starting...")
	logWriter := logger.Writer()
	outputLog := log.New(logWriter, "http: ", log.LstdFlags)

	// Init Cloud SQL MySQL
	dbClient, err := singletonDatabase.Get()
	if err != nil {
		log.Panicf("could not init MySQL client: %v", err)
	}

	// Init Endpoints Server injecting dependencies
	endpointService := endpoints.NewEndpointService(
		dbClient,
		config,
	)

	router := http.NewServeMux()
	router.Handle("/", endpointService.GetIndex())
	router.Handle("/healthz", healthz())
	router.Handle("/competition_results/", endpointService.GetCompetitionResults())

	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      tracing(nextRequestID)(logging(outputLog)(router)),
		ErrorLog:     outputLog,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Println("Server is ready to handle requests at", listenAddr)
	atomic.StoreInt32(&healthy, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	logger.Println("Server stopped")
}

func healthz() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&healthy) == 1 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
