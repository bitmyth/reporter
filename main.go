package main

import (
	"flag"
	"github.com/segmentio/analytics-go/v3"
	"log"
	"net"
)

//import "github.com/segmentio/public-api-sdk-go"

func main() {
	var writeKey string
	var key string
	var value string
	flag.StringVar(&writeKey, "key", "EpwBUylzncVWV4v1A5BQHme1ukUc4l0T", "segment source write key")
	flag.StringVar(&key, "k", "test", "event type")
	flag.StringVar(&value, "v", "test", "event value")
	flag.Parse()
	log.Println("write key:", writeKey)

	client := analytics.New(writeKey)
	defer client.Close()

	err := client.Enqueue(analytics.Identify{
		UserId: "home mac pro: " + GetOutboundIP().String(),
		Traits: analytics.NewTraits().
			Set(key, value).
			Set("ip", GetOutboundIP()),
		//SetName("Michael Bolton").
		//SetEmail("mbolton@example.com").
		//Set("plan", "Enterprise").
		//Set("friends", 42),
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
