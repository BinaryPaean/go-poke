package poke

import (
	"log"
	"net"
	"os"
	"time"
)

const Version = "0.0.1"

type action func(*Poke) *Result
type actionChain []action

type Result struct {
	Name      string
	Initiated time.Time
	Completed time.Time
	Result    string
	Err       error
}

type resultChain []*Result

type Poke struct {
	Timestamp   time.Time
	Target      string
	Host        string
	Version     string
	*log.Logger `json:"-"`
	actions     actionChain `json:"-"`
	Results     resultChain `json:"Latency"`
}

func NewPoke(target string) *Poke {
	var host, err = os.Hostname()
	if err == nil {
		host = "Unknown"
	}

	return &Poke{
		Timestamp: time.Now(),
		Target:    target,
		Host:      host,
		Version:   Version,
		Logger:    log.New(os.Stderr, "", log.LstdFlags)}
}

func (p *Poke) Include(action action) {
	p.actions = append(p.actions, action)
}

func (p *Poke) Run() {
	p.Results = make(resultChain, len(p.actions))
	for i, a := range p.actions {
		p.Results[i] = a(p)
	}
}

func DNSLookup(p *Poke) *Result {
	r := &Result{}
	r.Name = "DNS Lookup"
	r.Initiated = time.Now()
	rslt, err := net.LookupIP(p.Target)
	r.Completed = time.Now()
	if err != nil {
		r.Err = err
	} else {
		for _, s := range rslt {
			bytes, merr := s.MarshalText()
			if err != nil {
				r.Err = merr
			} else {
				r.Result += string(bytes)
			}
		}
	}
	return r
}
