package main

import "encoding/json"
import "github.com/binarypaean/go-poke/poke"

func main() {
	p := poke.NewPoke("www.google.com")
	p.Include(poke.DNSLookup)
	p.Run()

	var rslt []byte
	rslt, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		p.Fatalf("Marshal failed: %v", err.Error())
	}
	p.Print(string(rslt))
}
