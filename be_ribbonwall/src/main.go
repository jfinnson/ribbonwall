package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

func mainHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World! Ribbonwall! (Version info: %s, build date: %s)", os.Getenv("VERSION_INFO"), os.Getenv("BUILD_DATE"))
	})
}

func dbTest() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbUser := os.Getenv("db_user")
		dbPassword := os.Getenv("db_password")
		dbName := os.Getenv("db_name")
		dbEndpoint := os.Getenv("db_endpoint")
		//awsRegion := os.Getenv("aws_region")
		//awsArn := os.Getenv("aws_arn")

		//awsCreds := stscreds.NewCredentials(session.New(&aws.Config{Region: &awsRegion}), awsArn)
		//authToken, err := rdsutils.BuildAuthToken(dbEndpoint, awsRegion, dbUser, awsCreds)

		// Create the MySQL DNS string for the DB connection
		// user:password@protocol(endpoint)/dbname?<params>
		dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true",
			dbUser, dbPassword, dbEndpoint, dbName,
		)

		//
		//
		//rootCertPool := x509.NewCertPool() //NewCertPool returns a new, empty CertPool.
		//
		//pem, err := ioutil.ReadFile("rds-ca-bundle.pem") //reading the provided pem
		//if err != nil {
		//	_, _ = fmt.Fprintf(w, "Could not read certificates")
		//	log.Fatal("! Could not read certificates")
		//}
		//fmt.Println("Loading certificate seems to work")
		////AppendCertsFromPEM attempts to parse a series of PEM encoded certificates.
		////pushing in the pem
		//if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		//	_, _ = fmt.Fprintf(w, "Failed to append PEM")
		//	log.Fatal("Failed to append PEM.")
		//}
		//fmt.Println("Appending certificate seems to work too")
		//
		////setting up TLS
		////we dont need a client ca?
		//_ = mysql.RegisterTLSConfig("custom", &tls.Config{
		//	RootCAs:            rootCertPool,
		//	InsecureSkipVerify: true,
		//})

		// Use db to perform SQL operations on database
		db, err := sql.Open("mysql", dnsStr)
		defer db.Close()
		if err != nil {
			_, _ = fmt.Fprintf(w, "Error sql.Open %s", err.Error())
			return
		}

		err = db.Ping()
		if err != nil {
			_, _ = fmt.Fprintf(w, "Error ping %s. ", err.Error()) // authToken %s . dbEndpoint %s, awsRegion %s, dbUser %s, awsArn %s", err.Error(), authToken, dbEndpoint, awsRegion, dbUser, awsArn)
			return
		}

		rows, err := db.Query("SELECT * FROM ribbonwall_db.test_table")
		if err != nil {
			_, _ = fmt.Fprintf(w, "Error query %s", err.Error())
			return
		}

		cols, err := rows.Columns()
		if err != nil {
			fmt.Println("Failed to get columns", err)
			return
		}

		// Result is your slice string.
		rawResult := make([][]byte, len(cols))
		result := make([]string, len(cols))

		dest := make([]interface{}, len(cols)) // A temporary interface{} slice
		for i, _ := range rawResult {
			dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
		}

		for rows.Next() {
			err = rows.Scan(dest...)
			if err != nil {
				fmt.Println("Failed to scan row", err)
				return
			}

			for i, raw := range rawResult {
				if raw == nil {
					result[i] = "\\N"
				} else {
					result[i] = string(raw)
				}
			}

			fmt.Printf("%#v\n", result)
		}

		_, _ = fmt.Fprintf(w, "Successfully opened connection to database! Test results: %s", result)
	})
}

func main() {
	http.HandleFunc("/", mainHandler())
	http.HandleFunc("/db", dbTest())
	_ = http.ListenAndServe(":8080", nil)
}
