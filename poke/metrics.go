package poke

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"github.com/shenwei356/util/bytesize"
	"net"
	"net/http"
	"strings"
	"time"
)

type action func(*Poke) *Result
type metric func(action) action

func (a action) addMetric(m metric) action {
	return m(a)
}

func timeAction(f action) action {
	return func(p *Poke) *Result {
		initiated := time.Now()
		r := f(p)
		duration := time.Now().Sub(initiated)
		r.Metrics["duration"] = duration.String()
		return r
	}
}

func sizeInBytes(f action) action {
	return func(p *Poke) *Result {
		r := f(p)
		r.Metrics["size"] = fmt.Sprintf("%v", bytesize.ByteSize(
			len(
				[]byte(
					strings.Join(r.Response, "")))))
		return r
	}
}

func contentHash(f action) action {
	return func(p *Poke) *Result {
		r := f(p)

		hash := sha256.Sum256([]byte(strings.Join(r.Response, "")))
		r.Metrics["sha256"] = hash[:]
		return r
	}
}

func dnsLookup(p *Poke) *Result {
	r := NewResult("DNS Lookup", p.Target.Host)

	var resp []net.IP
	if resp, r.Err = net.LookupIP(r.Target); r.Err == nil {
		r.Response = make([]string, len(resp))
		for i, p := range resp {
			r.Response[i] = p.String()
		}
	}
	return r
}

func httpGet(p *Poke) *Result {
	r := NewResult("HTTP GET", p.Target.String())

	var err error
	var resp *http.Response
	if resp, err = http.Get(r.Target); err == nil {
		defer resp.Body.Close()
		r.Response = make([]string, 20) //Magic number alert! Should have default
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			r.Response = append(r.Response, scanner.Text())
		}
		//This was creating an empty error, rather than nil which caused inconsistent
		//output in the JSON. May have to handle unexpected EOF explicitly.
		//err = scanner.Err()
	}

	r.Err = err
	return r
}

func DNSLookup(p *Poke) *Result {
	a := action(dnsLookup).addMetric(timeAction).addMetric(contentHash).addMetric(sizeInBytes)
	return a(p)
}

func GetRequest(p *Poke) *Result {
	a := action(httpGet).addMetric(timeAction).addMetric(contentHash).addMetric(sizeInBytes)
	return a(p)
}
