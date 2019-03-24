package main

import "database/sql"

var (
	metricsPusherDelay int64 = 5

	DBURL            string = "postgresql.docker.localhost"
	DISTRIBUTORURL   string = "localhost:8081"
	METRICBUNDLERURL string = "metrics-bundler.nbg1.emile.space"

	DBHOST    = "postgres.docker.localhost"
	DBPORT    = 5432
	DBUSER    = "postgres"
	DBPASSWD  = ""
	DBNAME    = "postgres"
	DBSSLMODE = "disable"

	db *sql.DB

	metrics        map[string]float64
	starsProcessed int64
	theta          = 0.5
)

func main() {
	getEnvironment()
	connectToDatabase()
	initMetricPusher()
	processStars()
}
