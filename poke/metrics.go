package poke

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"net"
	"strings"
	"time"
)

type action func(*Poke) *Result
type metric func(*Poke, *Result) action

func timeAction(f action) action {
	return func(p *Poke) *Result {
		initiated := time.Now()
		r := f(p)
		duration := time.Now().Sub(initiated)
		r.Metrics["duration"] = duration.String()
		return r
	}
}

func contentHash(f action) action {
	return func(p *Poke) *Result {
		r := f(p)

		hash := sha256.Sum256([]byte(strings.Join(r.Response, "")))
		r.Metrics["hash"] = hash[:]
		return r
	}
}

func dnsLookup(p *Poke) *Result {
	r := NewResult()
	r.Name = "DNS Lookup"

	var resp []net.IP
	resp, r.Err = net.LookupIP(p.Target)
	if r.Err == nil {
		r.Response = make([]string, len(resp))
		for i, p := range resp {
			r.Response[i] = p.String()
		}
	}
	return r
}

func rootGetRequest(p *Poke) *Result {
	r := NewResult()
	r.Name = "HTTP GET"
	conn, err := net.Dial("tcp", net.JoinHostPort(p.Target, p.Port))
	if err != nil {
		r.Err = err
		return r
	}

	fmt.Fprintf(conn, "GET / HTTP/1.1\r\nHost: %v\r\nConnection: close\r\n\r\n", p.Target)
	response, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		r.Err = err
	}
	r.Response = []string{response}
	return r
}

func DNSLookup(p *Poke) *Result {
	return contentHash(timeAction(dnsLookup))(p)
}

func GetRequest(p *Poke) *Result {
	return contentHash(timeAction(rootGetRequest))(p)
}
