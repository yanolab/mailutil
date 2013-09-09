package smtp

import (
	"errors"
	"net/smtp"
	"../message"
)

var (
	AlreadyConnected = errors.New("Already Connected")
	NotConnected = errors.New("Not Connected")
)

type SMTP struct {
	name string
	client *smtp.Client
	IsConnected bool
}

func New() *SMTP {
	return &SMTP{name: "gomail", IsConnected: false}
}

func (s *SMTP) SetName(name string) {
	s.name = name
}

func (s *SMTP) Close() error {
	var err error = nil

	if s.IsConnected {
		err = s.client.Quit()
		s.IsConnected = false
	}

	return err
}

func (s *SMTP) Connect(host string) error {
	if s.IsConnected {
		return AlreadyConnected
	}

	client, err := smtp.Dial(host)
	if err != nil {
		return err
	}

	if err = client.Hello(s.name); err != nil {
		client.Quit()
		return err
	}

	s.client = client
	s.IsConnected = true
	return nil
}

func (s *SMTP) sendRcpt(addrs []string) error {
	for _, addr := range addrs {
		if err := s.client.Rcpt(addr); err != nil {
			return err
		}
	}
	return nil
}

func (s *SMTP) SendMail(m message.Message) error {
	if !s.IsConnected {
		return NotConnected
	}

	if err := s.client.Mail(m.From); err != nil {
		return err
	}

	if err := s.sendRcpt(m.To); err != nil {
		return err
	}
	if err := s.sendRcpt(m.Cc); err != nil {
		return err
	}
	if err := s.sendRcpt(m.Bcc); err != nil {
		return err
	}

	data, err := s.client.Data()
	if err != nil {
		return err
	}
	defer data.Close()

	_, err = data.Write(m.Encode())

	return err
}
