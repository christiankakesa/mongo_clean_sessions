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

var mongoUrl = flag.String("url", "mongodb://localhost:27017/test", "MongoDB connection URI.")
var collection = flag.String("c", "sessions", "MongoDB collection to cleanup.")
var field = flag.String("f", "updated_at", "MongoDB collection field with type 'time.Time'.")
var retention = flag.Int("r", 168, "MongoDB retention delai in hour(s). Default is 7 days (168 hours).")

func main() {
	flag.Parse()

	hostPort, err := url.Parse(*mongoUrl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection URL:              ", hostPort.Host)
	fmt.Println("MongoDB collection to clean:         ", *collection)
	fmt.Println("MongoDB collection field:            ", *field)
	fmt.Println("MongoDB retention periode in hour(s):", *retention)
	fmt.Println()

	session, err := mgo.Dial(*mongoUrl)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	Sessions := session.DB("").C(*collection)
	counter, err := Sessions.Count()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Number of item(s) in           %s: %d", *collection, counter))

	duration := time.Now().Add(time.Duration(-*retention) * time.Hour)
	query := bson.M{fmt.Sprintf("%s", *field): bson.M{"$lte": duration}}

	counter_to_delete, err := Sessions.Find(query).Count()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Number of item(s) to delete in %s: %d", *collection, counter_to_delete))

	info, err := Sessions.RemoveAll(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Number of item(s) deleted in   %s: %d", *collection, info.Removed))
}
