package poke

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
)

const Version = "0.0.2"

type Poke struct {
	Timestamp time.Time
	Target    *url.URL `json:"-"`
	Host      string
	Version   string      `json:"GoPokeVersion"`
	err       *log.Logger `json:"-"`
	output    *os.File    `json:"-"`
	actions   []action    `json:"-"`
	Results   []*Result
}

func NewPoke(userTarget string) *Poke {

	p := &Poke{
		Timestamp: time.Now(),
		Version:   Version,
		err:       log.New(os.Stderr, "", log.LstdFlags),
		output:    os.Stdout}

	var err error
	p.Host, err = os.Hostname()
	if err != nil {
		p.Host = "Unknown"
	}

	p.Target, err = url.Parse(userTarget)
	if err != nil {
		p.err.Fatalf("Unable to parse URL from: %v", err)
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
		p.err.Fatalf("Marshal failed: %v", err.Error())
	}
	fmt.Fprintln(p.output, string(rslt))
}
