package main

import "net"
import "log"
import "time"
import "os"
import "encoding/json"

const version = "0.0.1"

type Poke struct {
	Timestamp time.Time
	Target    string
	Host      string
	Version   string
	*log.Logger
}

func NewPoke(target string) *Poke {
	var host, err = os.Hostname()
	if err != nil {
		host = "Unknown"
	}
	var theLog = log.New(os.Stderr, "", log.LstdFlags)
	return &Poke{
		Timestamp: time.Now(),
		Target:    target,
		Host:      host,
		Version:   version,
		Logger:    theLog}
}

func (p *Poke) LookupIP() error {
	var tStart = time.Now()
	var _, err = net.LookupIP(p.Target)
	var duration = time.Since(tStart)
	p.Printf("Lookup of %v took %v", p.Target, duration)
	return err
}

func main() {
	var p = NewPoke("www.google.com")
	p.Printf("Gopoke called on: %v", p.Target)
	var err = p.LookupIP()
	if err != nil {
		p.Printf("Lookup failed: %v", err)
	}
	var rslt []byte
	rslt, err = json.MarshalIndent(p, "", "  ")
	if err != nil {
		log.Fatalf("Marshal failed: %v", err.Error())
	}
	p.Print(string(rslt))
}
