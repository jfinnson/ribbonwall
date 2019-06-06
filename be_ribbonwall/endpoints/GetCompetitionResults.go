package endpoints

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ribbonwall/be_ribbonwall/models"
	log "github.com/ribbonwall/common/logging"
	"net/http"
)

func (s *Endpoints) GetCompetitionResults() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//dbUser := os.Getenv("DB_USER")
		//dbPassword := os.Getenv("DB_PASSWORD")
		//dbName := os.Getenv("DB_NAME")
		//dbEndpoint := os.Getenv("DB_HOST")
		//
		//// Create the MySQL DNS string for the DB connection
		//// user:password@protocol(endpoint)/dbname?<params>
		//dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		//	dbUser, dbPassword, dbEndpoint, dbName,
		//)
		//
		//// Use db to perform SQL operations on database
		//db, err := sql.Open("mysql", dnsStr)
		//defer db.Close()
		//if err != nil {
		//	_, _ = fmt.Fprintf(w, "Error sql.Open %s", err.Error())
		//	return
		//}
		//
		//err = db.Ping()
		//if err != nil {
		//	_, _ = fmt.Fprintf(w, "Error ping %s. ", err.Error())
		//	return
		//}

		var competition_results []models.CompetitionResults
		s.Services.DB.Debug().Find(&competition_results)

		//cols, err := rows.Columns()
		//if err != nil {
		//	log.Errorf("Failed to get columns", err)
		//	return
		//}
		//
		//// Result is your slice string.
		//rawResult := make([][]byte, len(cols))
		//result := make([]string, len(cols))
		//
		//dest := make([]interface{}, len(cols)) // A temporary interface{} slice
		//for i := range rawResult {
		//	dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
		//}
		//
		//for rows.Next() {
		//	err = rows.Scan(dest...)
		//	if err != nil {
		//		log.Errorf("Failed to scan row", err)
		//		return
		//	}
		//
		//	for i, raw := range rawResult {
		//		if raw == nil {
		//			result[i] = "\\N"
		//		} else {
		//			result[i] = string(raw)
		//		}
		//	}
		//
		//	log.Infof("%#v\n", result)
		//}

		data, err := json.Marshal(competition_results)
		if err != nil {
			log.Errorf("error marshalling entity, %v", err)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "Database!\n Test results:\n%s", data)
	})
}
