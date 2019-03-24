package main

import (
	"fmt"
	"git.darknebu.la/GalaxySimulator/db-actions"
	"git.darknebu.la/GalaxySimulator/structs"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func processStars() {
	log.Println("[processStars] Initializing the processing process")

	for {
		log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

		// print the amount of stars allready processed
		printStarsProcessed()

		// get some values such as a starID,
		starID := getStarID()
		star := getStar(starID)
		timestep := getTimestep(starID)

		//force := calcAllForces(star, timestep, theta)
		force := structs.Vec2{10, 10}
		newStar := calcNewPos(star, force, 1e4)

		insertStar(newStar, timestep+1)

		time.Sleep(5 * time.Second)
	}
}

func printStarsProcessed() {
	log.Printf("[   ] Stars processed: %d", starsProcessed)
}

func getStarID() int64 {
	log.Printf("[getStarID] Getting a new starID from the distributor (%s)...", DISTRIBUTORURL)

	log.Println("[getStarID] Preparing the http request")
	requestURL := fmt.Sprintf("http://%s/distributor", DISTRIBUTORURL)
	resp, err := http.Get(requestURL)
	log.Printf("[getStarID] Sent http request to %s\n", requestURL)
	defer resp.Body.Close()
	if err != nil {
		log.Printf("resp: %v", resp)
		panic(err)
	}

	log.Println("[getStarID] Reading the response body")
	body, err := ioutil.ReadAll(resp.Body)
	log.Println("[getStarID] Done Reading the response body")
	if err != nil {
		panic(err)
	}

	log.Printf("[getStarID] Got the response: %v", string(body))
	log.Println("[getStarID] Parsing the integer")
	starID, parseIntErr := strconv.ParseInt(string(body), 10, 64)
	if parseIntErr != nil {
		panic(parseIntErr)
	}
	log.Println("[getStarID] Done parsing the integer")

	return starID
}

func getStar(starID int64) structs.Star2D {
	var star = db_actions.GetStar(db, starID)
	return star
}

func getTimestep(starID int64) int64 {
	var timestep = db_actions.GetStarIDTimestep(db, starID)
	return timestep
}

func calcNewPos(star structs.Star2D, force structs.Vec2, timestep float64) structs.Star2D {
	star.CalcNewPos(force, timestep)
	returnStar := structs.Star2D{
		C: structs.Vec2{
			X: star.C.X,
			Y: star.C.Y,
		},
		V: structs.Vec2{
			X: force.X,
			Y: force.Y,
		},
		M: star.M,
	}
	return returnStar
}

func insertStar(star structs.Star2D, timestep int64) {
	log.Printf("Inserting the star %v into the timestep %d", star, timestep)
}
