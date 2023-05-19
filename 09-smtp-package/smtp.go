package main

import "log"

type Smtp struct {
	Address  string
	Port     string
	Username string
	Password string
}

func (s *Smtp) send() {
	log.Println("send email")
}

func newSmtp(address, port, username, password string) *Smtp {
	return &Smtp{
		Address:  address,
		Port:     port,
		Username: username,
		Password: password,
	}
}

type Option func(*Smtp)

func WithPort(port string) Option {
	return func(s *Smtp) {
		s.Port = port
	}
}
func main() {
	foo := newSmtp("127.0.0.1", "25")
	foo.send()
}
