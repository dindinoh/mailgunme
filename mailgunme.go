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

//defaultchecker
func defaultchecker(configvalue, arg string) string {
	var retval string
	if configvalue != "" {
		retval = configvalue
	}
	if arg != "" {
		retval = arg
	}
	if retval == "" {
		log.Fatalf("No " + arg + " set!")
	}
	return retval
}

// send will do the actual call
func send(cfg Config, fromaddressname, fromname, to, message, subject string) {
	var tfromaddressname, tfromname, tto, tsubject string

	tfromaddressname = defaultchecker(cfg.Mailgun.Fromaddressname, tfromaddressname)
	tfromname = defaultchecker(cfg.Mailgun.Fromname, fromname)
	tto = defaultchecker(cfg.Mailgun.Toaddress, to)
	tsubject = defaultchecker(cfg.Mailgun.Subject, subject)

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
