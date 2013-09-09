package message

import (
	"bytes"
	"fmt"
	"strings"
)

type Message struct {
	Cc []string
	To []string
	Bcc []string
	From string
	Subject string
	Body string
}

func New() *Message {
	return &Message{}
}

func (m Message) Encode() []byte {
	buffer := new(bytes.Buffer)

	fmt.Fprintf(buffer, "From: %s\n", m.From)
	fmt.Fprintf(buffer, "To: %s\n", strings.Join(m.To, ","))
	fmt.Fprintf(buffer, "Cc: %s\n", strings.Join(m.Cc, ","))
	fmt.Fprintf(buffer, "Subject: %s\n", m.Subject) // need to encode
	fmt.Fprint(buffer, "\n")
	fmt.Fprint(buffer, m.Body)

	return buffer.Bytes()
}
