package main

import (
	"flag"
	"fmt"
	"log"
  "gopkg.in/gcfg.v1"
	"github.com/mailgun/mailgun-go"
)

// Config is the struct for config file in ~/.mailgunme
type Config struct {
	Hardcoded struct {
		PublicKey, PrivateKey, Domain string
	}
}

// parse_config reads config file in ~/.mailgunme
func parse_config() Config {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "~/.mailgunme")

	if err != nil {
		log.Fatalf("Failed to parse mailgunme config file in ~/.mailgunme", err)
	}
	return
}

// send will do the actual call
func send(cfg Config, to, msg string) {
	fmt.Println("phone home")
}

// main parses config and cli options and calls send function
func main() {
	parse_config()

	toPtr := flag.String("to", "", "recipient address")
	msgPtr := flag.String("msg", "", "message")

	gun := mailgun.NewMailgun("valid-mailgun-domain", "private-mailgun-key", "public-mailgun-key")

	m := mailgun.NewMessage("Sender <sender@example.com>", "Subject", "Message Body", "Recipient <recipient@example.com>")
	response, id, _ := gun.Send(m)
	fmt.Printf("Response ID: %s\n", id)
	fmt.Printf("Message from server: %s\n", response)
}
