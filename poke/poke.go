package poke

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"time"
)

const Version = "0.0.2"

type Poke struct {
	Timestamp   time.Time
	Target      *url.URL `json:"-"`
	Host        string
	Version     string
	*log.Logger `json:"-"`
	actions     []action `json:"-"`
	Results     []*Result
}

func NewPoke(userTarget string) *Poke {

	p := &Poke{
		Timestamp: time.Now(),
		Version:   Version,
		Logger:    log.New(os.Stderr, "", log.LstdFlags)}

	var err error
	p.Host, err = os.Hostname()
	if err != nil {
		p.Host = "Unknown"
	}

	p.Target, err = url.Parse(userTarget)
	if err != nil {
		p.Fatalf("Unable to parse URL from given path: %v", err)
		os.Exit(1)
	}

	return p
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
