package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"
	_ "github.com/go-sql-driver/mysql"
)

func mainHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World! Ribbonwall! (Version info: %s, build date: %s)", os.Getenv("VERSION_INFO"), os.Getenv("BUILD_DATE"))
	})
}

func dbTest() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbUser := os.Getenv("db_user")
		//dbPassword := os.Getenv("db_password")
		dbName := os.Getenv("db_name")
		dbEndpoint := os.Getenv("db_endpoint")
		awsRegion := os.Getenv("aws_region")
		awsArn := os.Getenv("aws_arn")

		awsCreds := stscreds.NewCredentials(session.New(&aws.Config{Region: &awsRegion}), awsArn)
		authToken, err := rdsutils.BuildAuthToken(dbEndpoint, awsRegion, dbUser, awsCreds)

		// Create the MySQL DNS string for the DB connection
		// user:password@protocol(endpoint)/dbname?<params>
		dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true",
			dbUser, authToken, dbEndpoint, dbName,
		)

		// Use db to perform SQL operations on database
		db, err := sql.Open("mysql", dnsStr)
		defer db.Close()
		if err != nil {
			_, _ = fmt.Fprintf(w, "Error sql.Open %s", err.Error())
			return
		}

		err = db.Ping()
		if err != nil {
			_, _ = fmt.Fprintf(w, "Error ping %s. authToken %s ", err.Error(), authToken)
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
