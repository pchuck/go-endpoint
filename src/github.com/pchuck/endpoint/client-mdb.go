// simple mgo client
//   (adapted from http://labix.org/mgo examples)
package main

import (
        "fmt"
	"log"
	"os"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type Person struct {
        FirstName string
        LastName string
        Phone string
}

func mdbOps(server string, db string, collection string) {
//        session, err := mgo.Dial("server1", "server2")
        session, err := mgo.Dial(server)
        if err != nil {
                panic(err)
        }
        defer session.Close()

        // Optional. Switch the session to a monotonic behavior.
        session.SetMode(mgo.Monotonic, true)

        c := session.DB(db).C(collection)

	eric := Person{"Eric", "Brown", "+1-303-555-1212"}
        clara := Person{"Clara", "Young", "+1-303-555-1234"}
        err = c.Insert(&eric, &clara)

        if err != nil {
                log.Fatal(err)
        }

        result := Person{}
        err = c.Find(bson.M{"firstname": "Clara"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Phone:", result.Phone)
}

func main() {
	args := os.Args

	if len(args) < 4 {
		fmt.Println("   usage: ", args[0], " server db collection")
		fmt.Println("     e.g. ", args[0], " localhost:27017 test people")
		os.Exit(2)
	}

	server := os.Args[1]
	db := os.Args[2]
	collection := os.Args[3]

	mdbOps(server, db, collection)
}


