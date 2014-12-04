package main

import "os"
import "github.com/binarypaean/go-poke/poke"

func main() {
	args := os.Args[1:]
	pokes := make([]*poke.Poke, len(args))

	for i, a := range args {
		pokes[i] = poke.NewPoke(a)
	}

	for _, p := range pokes {
		p.AddAction(poke.DNSLookup)
		p.AddAction(poke.GetRequest)
		p.Run()
	}
}
