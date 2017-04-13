package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mattn/go-xmpp"
)

var server = flag.String("server", "", "server:port")
var username = flag.String("username", "", "username")
var password = flag.String("password", "", "password")
var status = flag.String("status", "", "Status")
var statusMessage = flag.String("status-msg", "", "Status message")
var notls = flag.Bool("notls", true, "No TLS")
var starttls = flag.Bool("starttls", true, "Enable StartTLS")
var debug = flag.Bool("debug", false, "Enable debug output")
var session = flag.Bool("session", false, "Use server session")

var remote = flag.String("remote", "", "Jid of the receiver/remote party.")
var message = flag.String("message", "", "Message to send.")

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: example [options]\n")
		flag.PrintDefaults()
		os.Exit(2)
	}
	flag.Parse()

	if !*notls {
		xmpp.DefaultConfig = tls.Config{
			ServerName:         serverName(*server),
			InsecureSkipVerify: false,
		}
	}

	xmpp.DefaultConfig = tls.Config{
		ServerName:         serverName(*server),
		InsecureSkipVerify: true,
	}

	var talk *xmpp.Client
	var err error
	options := xmpp.Options{
		Host:                         *server,
		User:                         *username,
		Password:                     *password,
		NoTLS:                        *notls,
		Debug:                        *debug,
		Session:                      *session,
		Status:                       *status,
		StatusMessage:                *statusMessage,
		InsecureAllowUnencryptedAuth: false,
		StartTLS:                     *starttls,
	}

	talk, err = options.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	talk.Send(xmpp.Chat{Remote: *remote, Type: "chat", Text: *message})

}

//Seperate domain name and port.
func serverName(host string) string {
	return strings.Split(host, ":")[0]
}
