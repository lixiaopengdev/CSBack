package db

import (
	"CSBackendTmp/ent"
	"CSBackendTmp/ent/migrate"
	"context"
	"database/sql"
	"log"
	"os"
	"strings"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	// "github.com/influxdata/influxdb-client-go/api"
	// "github.com/influxdata/influxdb-client-go/api"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
	// influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	// "github.com/influxdata/influxdb-client-go/v2/api"
)

var DBClient *ent.Client

// var MonAPI api.WriteAPI

func Init() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbNode, dbNodeExists := os.LookupEnv("MYSQL_NODE")
	dbPort, dbPortExists := os.LookupEnv("MYSQL_PORT")
	dbUserName, dbUserNameExists := os.LookupEnv("MYSQL_USERNAME")
	dbPW, dbPWExists := os.LookupEnv("MYSQL_PASSWORD")
	dbName, dbNameExists := os.LookupEnv("MYSQL_DBNAME")
	if !dbNodeExists || !dbPortExists || !dbUserNameExists || !dbPWExists || !dbNameExists {
		log.Fatal("FATAL ERROR: ENV not properly configured, check .env file or Database config")
	}
	path := strings.Join([]string{dbUserName, ":", dbPW, "@tcp(", dbNode, ":", dbPort, ")/", dbName, "?charset=utf8mb4&parseTime=true"}, "")
	db, err := sql.Open("mysql", path)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)
	drv := entsql.OpenDB("mysql", db)
	ctx := context.Background()
	DBClient = ent.NewClient(ent.Driver(drv))
	if err := DBClient.Schema.Create(ctx, migrate.WithForeignKeys(false)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// initTSDB()
}

// func initTSDB() {

// 	dbNode, dbNodeExists := os.LookupEnv("INFLUXDB_NODE")
// 	dbPort, dbPortExists := os.LookupEnv("INFLUXDB_PORT")
// 	dbToken, dbTokenExists := os.LookupEnv("INFLUXDB_TOKEN")
// 	if !dbNodeExists || !dbPortExists || !dbTokenExists {
// 		log.Fatal("FATAL ERROR: ENV not properly configured, check .env file or Database config")
// 	}
// 	// Create a new client using an InfluxDB server base URL and an authentication token
// 	client := influxdb2.NewClientWithOptions(strings.Join([]string{"http://", dbNode, ":", dbPort}, ""), dbToken, influxdb2.DefaultOptions().SetBatchSize(100))
// 	// Use blocking write client for writes to desired bucket
// 	MonAPI = client.WriteAPI("neoworld", "neoworld")
// }
