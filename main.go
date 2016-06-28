package main

import (
	"flag"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/url"
	"time"
)

// Version is initialized at compilation time
var Version = "0.0.0"

// BuildTime is initialized at compilation time
var BuildTime = "2016-06-29T00:39:38+0200"

func main() {
	var urlFlag = flag.String("url", "mongodb://localhost:27017/test", "MongoDB connection URI.")
	var cFlag = flag.String("c", "sessions", "MongoDB collection to cleanup.")
	var fFlag = flag.String("f", "updated_at", "MongoDB collection field with type 'time.Time'.")
	var rFlag = flag.Int("r", 168, "MongoDB retention delai in hour(s). Default is 7 days (168 hours).")
	var sFlag = flag.Bool("s", false, "Simulation mode, no deletion are send to the MongoDB database.")

	flag.Parse()

	hostPort, err := url.Parse(*urlFlag)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("[Mongo clean session - Version %s (build at %s)]", Version, BuildTime))
	fmt.Println("MongoDB connection URL:              ", hostPort.Host)
	fmt.Println("MongoDB collection to clean:         ", *cFlag)
	fmt.Println("MongoDB collection field:            ", *fFlag)
	fmt.Println("MongoDB retention periode in hour(s):", *rFlag)
	fmt.Println("Simulation mode:                     ", *sFlag)
	fmt.Println()

	session, err := mgo.Dial(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	Sessions := session.DB("").C(*cFlag)
	counter, err := Sessions.Count()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Number of item(s) in           %s: %d", *cFlag, counter))

	duration := time.Now().Add(time.Duration(-*rFlag) * time.Hour)
	query := bson.M{fmt.Sprintf("%s", *fFlag): bson.M{"$lte": duration}}

	counter_to_delete, err := Sessions.Find(query).Count()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Number of item(s) to delete in %s: %d", *cFlag, counter_to_delete))

	if *sFlag == false {
		info, err := Sessions.RemoveAll(query)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fmt.Sprintf("Number of item(s) deleted in   %s: %d", *cFlag, info.Removed))
	} else {
		fmt.Println("[SIMULATION_MODE] In simulation mode no elements are deleted!!!")
	}
}
