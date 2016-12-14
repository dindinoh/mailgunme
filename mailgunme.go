package main

import (
	"flag"
	"fmt"
	"github.com/mailgun/mailgun-go"
	"github.com/mitchellh/go-homedir"
	"gopkg.in/gcfg.v1"
	"log"
	"errors"
	"os"
	"bufio"
)

// Config is the struct for config file in ~/.mailgunme
type Config struct {
	Mailgun struct {
		Privatekey, Publickey, Domain, Fromaddressname, Fromname, Subject, Toaddress string
	}
}

// parse_config reads config file in ~/.mailgunme
func ParseConfig() (Config, error) {
	var cfg Config
	var err error
	var home string

	home, err = homedir.Dir()

	// how to test non existent file, erroneus file?
	if err == nil {
		err = gcfg.ReadFileInto(&cfg, home+"/.mailgunme")
	}

	return cfg, err
}

//defaultchecker will try to find a default if a value is not provided
func Defaultchecker(configvalue, arg, name string) (string, error) {
	var retval string
	var err error
	
	if configvalue != "" {
		retval = configvalue
	}

	if arg != "" {
		retval = arg
	}
	if retval == "" || name == ""{
		err = errors.New("Empty or missing value.")
	}
	return retval, err
}

// send will do the actual call
func send(cfg Config, fromaddressname, fromname, to, message, subject string) {
	var tfromaddressname, tfromname, tto, tsubject string

	tfromaddressname,_ = Defaultchecker(cfg.Mailgun.Fromaddressname, fromaddressname, "From address name")
	tfromname,_ = Defaultchecker(cfg.Mailgun.Fromname, fromname, "From name")
	tto,_ = Defaultchecker(cfg.Mailgun.Toaddress, to, "To")
	tsubject,_ = Defaultchecker(cfg.Mailgun.Subject, subject, "Subject")

	gun := mailgun.NewMailgun(cfg.Mailgun.Domain, cfg.Mailgun.Privatekey, cfg.Mailgun.Publickey)
	m := mailgun.NewMessage(tfromname+
		" <"+tfromaddressname+"@"+cfg.Mailgun.Domain+">",
		tsubject,
		message,
		tto)
	response, id, err := gun.Send(m)
	if err != nil {
		log.Fatal("Error sending email!\n", err)
	}
	fmt.Printf("Response ID: %s\n", id)
	fmt.Printf("Message from server: %s\n", response)
}

// main parses config and cli options and calls send function
func main() {
	var err error
	var message string
	fromaddressnamePtr := flag.String("n", "", "from(@)")
	fromnamePtr := flag.String("f", "", "From Name")
	messagePtr := flag.String("m", "", "message")
	subjectPtr := flag.String("s", "", "subject")
	toPtr := flag.String("t", "", "to-address")
	flag.Parse()

	// look for pipe for message
	messin, err := os.Stdin.Stat()
	if messin.Mode() & os.ModeNamedPipe != 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			message = scanner.Text()
		}
	} else {
		message = *messagePtr
	}
	
	// parse config
	var cfg Config
	cfg, err = ParseConfig()
	if err != nil {
		log.Fatalf("FATAL")
	}
	
	send(cfg, *fromaddressnamePtr, *fromnamePtr, *toPtr, message, *subjectPtr)
}
