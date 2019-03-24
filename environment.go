package main

import (
	"flag"
	"log"
	"os"
	"strconv"
)

func getEnvironment() {
	log.Println("[environment] Initializing the environment")
	getFlags()
	getEnvironmentVars()
}

func getEnvironmentVars() {
	log.Println("[ ] Getting the environment variables...")

	// get the data that should be used to connect to the database

	DISTRIBUTORURL = os.Getenv("DISTRIBUTORURL")
	if DISTRIBUTORURL == "" {
		DISTRIBUTORURL = "localhost:8081"
	}

	METRICBUNDLERURL = os.Getenv("METRICBUNDLERURL")
	if METRICBUNDLERURL == "" {
		METRICBUNDLERURL = "metrics-bundler.nbg1.emile.space"
	}

	DBURL = os.Getenv("DBURL")
	if DBURL == "" {
		DBURL = "postgresql.docker.localhost"
	}

	// get the data that should be used to connect to the database
	DBUSER = os.Getenv("DBUSER")
	if DBUSER == "" {
		DBUSER = "postgres"
	}

	DBPASSWD = os.Getenv("DBPASSWD")
	if DBPASSWD == "" {
		DBPASSWD = ""
	}

	DBPORT, _ := strconv.ParseInt(os.Getenv("DBPORT"), 10, 64)
	if DBPORT == 0 {
		DBPORT = 5432
	}

	DBNAME = os.Getenv("DBNAME")
	if DBNAME == "" {
		DBNAME = "postgres"
	}

	log.Println("--------------------")
	log.Printf("DBURL: %s", DBHOST)
	log.Printf("DBUSER: %s", DBUSER)
	log.Printf("DBPASSWD: %s", DBPASSWD)
	log.Printf("DBPORT: %d", DBPORT)
	log.Printf("DBPROJECTNAME: %s", DBNAME)
	log.Println("--------------------")
	log.Printf("DISTRIBUTORURL: %s", DISTRIBUTORURL)
	log.Printf("METRICBUNDLERURL: %s", METRICBUNDLERURL)
	log.Println("--------------------")
}

func getFlags() {
	// Parse the metricsPusherDelay from the command line.
	// If no argument is provided, the delay defaults at 5 seconds
	flag.Int64Var(&metricsPusherDelay, "metricsPusherDelay", 5, "The delay (in seconds) between the metric Pushes")
	flag.Parse()
	log.Println("[ ] Done loading the flags")
}
