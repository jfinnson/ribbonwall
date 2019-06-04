package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"

	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/aws/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/rds/rdsutils"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

func mainHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello World! Ribbonwall! (Version info: %s, build date: %s)", os.Getenv("VERSION_INFO"), os.Getenv("BUILD_DATE"))
	})
}

func dbTest() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbUser := os.Getenv("db_user")
		dbName := os.Getenv("db_name")
		dbEndpoint := os.Getenv("db_endpoint")
		awsRegion := os.Getenv("aws_region")
		awsArn := os.Getenv("aws_arn")

		cfg, err := external.LoadDefaultAWSConfig()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "failed to load configuration, %v", err)
			os.Exit(1)
		}
		cfg.Region = awsRegion

		credProvider := stscreds.NewAssumeRoleProvider(sts.New(cfg), awsArn)

		authToken, err := rdsutils.BuildAuthToken(dbEndpoint, awsRegion, dbUser, credProvider)

		// Create the MySQL DNS string for the DB connection
		// user:password@protocol(endpoint)/dbname?<params>
		dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true",
			dbUser, authToken, dbEndpoint, dbName,
		)

		driver := mysql.MySQLDriver{}
		_ = driver
		// Use db to perform SQL operations on database
		db, err := sql.Open("mysql", dnsStr)
		if err != nil {
			_, _ = fmt.Fprintf(w, "Error %s", err.Error())
			panic(err.Error())
		}

		results, err := db.Query("SELECT * FROM ribbonwall_db.test_table")
		if err != nil {
			_, _ = fmt.Fprintf(w, "Error %s", err.Error())
			panic(err.Error())
		}

		_, _ = fmt.Fprintf(w, "Successfully opened connection to database! Test results: %s", results)
	})
}

func main() {
	http.HandleFunc("/", mainHandler())
	http.HandleFunc("/db", dbTest())
	_ = http.ListenAndServe(":8080", nil)
}
