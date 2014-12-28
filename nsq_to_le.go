package main

import (
	"flag"
	"github.com/bitly/go-nsq"
	"github.com/bsphere/le_go"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	var lookupd, topics, token string

	flag.StringVar(&lookupd, "lookupd", "http://127.0.0.1:4161",
		"lookupd HTTP address")

	flag.StringVar(&topics, "topics", "", "comma delimited NSQ topics")
	flag.StringVar(&token, "token", "", "Logentries token")
	flag.Parse()

	if lookupd == "" || topics == "" || token == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	le, err := le_go.Connect(token)
	if err != nil {
		log.Fatal(err)
	}

	defer le.Close()

	channel := "nsq_to_logentries" + strconv.FormatInt(time.Now().Unix(), 10) +
		"#ephemeral"

	for _, topic := range strings.Split(topics, ",") {
		le.Println("Logging messages from topic " + topic)

		c, err := nsq.NewConsumer(topic, channel, nsq.NewConfig())
		if err != nil {
			log.Fatal(err)
		}

		c.AddHandler(nsq.HandlerFunc(func(m *nsq.Message) error {
			defer m.Finish()

			le.Println(string(m.Body))

			return nil
		}))

		if err := c.ConnectToNSQLookupd(lookupd); err != nil {
			log.Fatal(err)
		}
	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
