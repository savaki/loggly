package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/segmentio/go-loggly"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const (
	fieldPort  = "port"
	fieldToken = "token"
	fieldTag   = "tag"
)

func main() {
	app := cli.NewApp()
	app.Name = "loggly"
	app.Usage = "send logs from the specified port to loggly"
	app.Author = "Matt Ho"
	app.Version = "0.1"
	app.Flags = []cli.Flag{
		cli.IntFlag{fieldPort, 6030, "the default port loggly will listen to", "PORT"},
		cli.StringFlag{fieldToken, "", "your loggly token", "TOKEN"},
		cli.StringSliceFlag{fieldTag, &cli.StringSlice{}, "the tags to apply to the stream", "TAGS"},
	}
	app.Action = server
	app.Run(os.Args)
}

func server(c *cli.Context) {
	port := c.Int(fieldPort)
	tags := findTags(c)
	token := c.String(fieldToken)
	if token == "" {
		log.Fatalln("ERROR: loggly token not specified")
	}

	client := loggly.New(token, tags...)

	fmt.Printf("starting loggly listener on port %d\n", port)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleRequest(conn, client)
	}
}

func findTags(c *cli.Context) []string {
	tags := []string{}

	for _, parts := range c.StringSlice(fieldTag) {
		for _, word := range strings.Split(parts, ",") {
			tags = append(tags, strings.TrimSpace(word))
		}
	}

	return tags
}

func handleRequest(conn net.Conn, client *loggly.Client) {
	defer conn.Close()
	defer func() { fmt.Println("closing connection") }()
	io.Copy(client, conn)
}
