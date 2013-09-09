package main

import (
	"testing"
	"./message"
	"./smtp"
)

func TestSendMail(t *testing.T) {
	smtp := smtp.New()

	if err := smtp.Connect("127.0.0.1:25"); err != nil {
		t.Fatal(err)
	}

	msg := message.New()
	msg.From = "root@localhost"
	msg.To = []string{"root@localhost"}
	msg.Subject = "test mail from go"
	msg.Body = "test mail send successfully"

	smtp.SendMail(*msg)
}
