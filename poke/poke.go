package poke

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"time"
)

const Version = "0.0.2"

type Poke struct {
	Timestamp   time.Time
	Target      string
	Port        string
	Host        string
	Version     string
	*log.Logger `json:"-"`
	actions     []action  `json:"-"`
	Results     []*Result `json:"Latency"`
}

func NewPoke(target string) *Poke {
	host, err := os.Hostname()
	if err != nil {
		host = "Unknown"
	}

	var port string
	_, port, err = net.SplitHostPort(target)
	if err != nil {
		//Warn: We assume this means no port, but could be a malformed address
		//Add default port of 80 if user did not specify one
		port = "80"
	}

	return &Poke{
		Timestamp: time.Now(),
		Target:    target,
		Host:      host,
		Port:      port,
		Version:   Version,
		Logger:    log.New(os.Stderr, "", log.LstdFlags)}
}

func (p *Poke) AddAction(action action) {
	p.actions = append(p.actions, action)
}

func (p *Poke) Run() {
	p.Results = make(results, len(p.actions))
	for i, a := range p.actions {
		p.Results[i] = a(p)
	}
	p.marshalToLog()
}

func (p *Poke) marshalToLog() {
	rslt, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		p.Fatalf("Marshal failed: %v", err.Error())
	}
	p.Print(string(rslt))
}
