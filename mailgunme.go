package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"

	"github.com/mailgun/mailgun-go"
	"gopkg.in/gcfg.v1"
)

// Config is the struct for config file in ~/.mailgunme
type Config struct {
	Mailgun struct {
		Privatekey, Publickey, Domain string
	}
}

// parse_config reads config file in ~/.mailgunme
func parse_config() Config {
	var cfg Config
	home, err := homedir.Dir()
	if err != nil {
		log.Fatalf("Could not find homedir, really?")
	}

	err = gcfg.ReadFileInto(&cfg, home+"/.mailgunme")

	if err != nil {
		log.Fatalf("Failed to parse mailgunme config file in ~/.mailgunme\n", err)
	}
	return cfg
}

// send will do the actual call
func send(cfg Config, from, to, message, subject string) {
	gun := mailgun.NewMailgun(cfg.Mailgun.Domain, cfg.Mailgun.Privatekey, cfg.Mailgun.Publickey)
	m := mailgun.NewMessage(from+" <"+from+"@"+cfg.Mailgun.Domain+">", "Subject", message, to)
	response, id, _ := gun.Send(m)
	fmt.Printf("Response ID: %s\n", id)
	fmt.Printf("Message from server: %s\n", response)
}

// main parses config and cli options and calls send function
func main() {
	fromPtr := flag.String("f", "", "from-name")
	messagePtr := flag.String("m", "", "message")
	subjectPtr := flag.String("s", "", "subject")
	toPtr := flag.String("t", "", "to-address")
	flag.Parse()

	if *toPtr == "" {
		fmt.Println("No recipient provided!")
		os.Exit(1)
	}

	var cfg Config
	cfg = parse_config()

	send(cfg, *fromPtr, *toPtr, *messagePtr, *subjectPtr)
}
