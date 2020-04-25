package main

import (
	"fmt"
	"log"
	"time"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

var (
	FetchAll = []imap.FetchItem{
		imap.FetchBody,
		imap.FetchBodyStructure,
		imap.FetchEnvelope,
		imap.FetchFlags,
		imap.FetchInternalDate,
		imap.FetchRFC822,
		imap.FetchRFC822Header,
		imap.FetchRFC822Size,
		imap.FetchRFC822Text,
		imap.FetchUid,
	}
)

type HandleMessage func(msg *imap.Message) error

func main() {
	log.Println("Connecting to server...")

	host := "host281.hostmonster.com:993"
	username := "hello@yoderstore.com"
	password := "TyvBK!.BRG1"
	for {
		log.Println("Processing messages")
		n, err := run(host, username, password, printMessage)
		if err != nil {
			log.Println(err)
		}
		// we didn't process any messages wait a bit before trying again
		if n == 0 {
			log.Println("Didn't process any messages. Pause a bit.")
			time.Sleep(30 * time.Second)
		}

		time.Sleep(3 * time.Second)
	}

}

func run(host string, username string, password string, process HandleMessage) (int, error) {
	// Connect to server
	c, err := client.DialTLS(host, nil)
	if err != nil {
		return 0, err
	}

	defer c.Logout()

	if err := c.Login(username, password); err != nil {
		return 0, err
	}

	mbox, err := c.Select("INBOX", false)
	if err != nil {
		return 0, err
	}

	seqset := new(imap.SeqSet)

	if mbox.Messages == 0 {
		return 0, nil
	}

	// Get the first message
	maxNumMsg := uint32(10)
	from := uint32(1)
	to := maxNumMsg
	if mbox.Messages < to {
		to = mbox.Messages
	}
	seqset.AddRange(from, to)

	messages := make(chan *imap.Message, maxNumMsg)
	if err := c.Fetch(seqset, []imap.FetchItem{imap.FetchRFC822}, messages); err != nil {
		return 0, err
	}

	if err := processMessages(messages, process); err != nil {
		return 0, err
	}

	item := imap.FormatFlagsOp(imap.AddFlags, true)
	flags := []interface{}{imap.DeletedFlag}
	err = c.Store(seqset, item, flags, nil)
	if err != nil {
		return 0, err
	}

	err = c.Expunge(nil)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

func processMessages(messages chan *imap.Message, process HandleMessage) error {

	for msg := range messages {
		process(msg)
	}

	return nil
}

func printMessage(msg *imap.Message) error {
	//	log.Println("* " + msg.Envelope.Subject)
	for name, value := range msg.Body {
		fmt.Printf("section: %#v\n\n", name)
		bytes := make([]byte, 5000)
		n, err := value.Read(bytes)
		fmt.Printf("\tvalue: %s\n", string(bytes[:n]))
		fmt.Print("-----------------------------------------------------------------------\n\n\n")
		if err != nil {
			return err
		}
	}
	return nil
}
