package main

import (
	"os"
	"time"

	"github.com/savaki/loggly"
)

func main() {
	token := os.Getenv("LOGGLY_TOKEN")
	client := loggly.New(token, loggly.Interval(5*time.Second))
	client.Write([]byte("{\"hello\":\"world\"}\n"))
	client.Flush()
}
