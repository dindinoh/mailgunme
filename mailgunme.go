package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/mailgun/mailgun-go"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/gcfg.v1"
)

// Config is the struct for config file in ~/.mailgunme
type Config struct {
	Mailgun struct {
		Privatekey, Publickey, Domain, Fromaddressname, Fromname, Subject, Toaddress string
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
func send(cfg Config, fromaddressname, fromname, to, message, subject string) {
	var tfromaddressname, tfromname, tto, tsubject string

	if cfg.Mailgun.Fromaddressname != "" {
		tfromaddressname = cfg.Mailgun.Fromaddressname
	}
	if fromaddressname != "" {
		tfromaddressname = fromaddressname
	}
	if tfromaddressname == "" {
		log.Fatalf("No From address name set!")
	}

	if cfg.Mailgun.Fromname != "" {
		tfromname = cfg.Mailgun.Fromname
	}
	if fromname != "" {
		tfromname = fromname
	}
	if tfromname == "" {
		log.Fatalf("No From name set!")
	}

	if cfg.Mailgun.Toaddress != "" {
		tto = cfg.Mailgun.Toaddress
	}
	if to != "" {
		tto = to
	}
	if tto == "" {
		log.Fatalf("No to address set!")
	}

	if cfg.Mailgun.Subject != "" {
		tsubject = cfg.Mailgun.Subject
	}
	if subject != "" {
		tsubject = subject
	}
	if tsubject == "" {
		log.Fatalf("No subject set!")
	}

	gun := mailgun.NewMailgun(cfg.Mailgun.Domain, cfg.Mailgun.Privatekey, cfg.Mailgun.Publickey)
	m := mailgun.NewMessage(tfromname+
		" <"+tfromaddressname+"@"+cfg.Mailgun.Domain+">",
		tsubject,
		message,
		tto)
	response, id, err := gun.Send(m)
	if err != nil {
		log.Fatalf("Error sending email!\n", err)
	}
	fmt.Printf("Response ID: %s\n", id)
	fmt.Printf("Message from server: %s\n", response)
}

// main parses config and cli options and calls send function
func main() {
	fromaddressnamePtr := flag.String("n", "", "from(@)")
	fromnamePtr := flag.String("f", "", "From Name")
	messagePtr := flag.String("m", "", "message")
	subjectPtr := flag.String("s", "", "subject")
	toPtr := flag.String("t", "", "to-address")
	flag.Parse()

	var cfg Config
	cfg = parse_config()

	send(cfg, *fromaddressnamePtr, *fromnamePtr, *toPtr, *messagePtr, *subjectPtr)
}
