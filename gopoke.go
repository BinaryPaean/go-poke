package main

import "os"
import "github.com/binarypaean/go-poke/poke"

func main() {
	args := os.Args[1:] // Drop the application name
	pokes := make([]*poke.Poke, len(args))

	for i, a := range args {
		pokes[i] = poke.NewPoke(a)
	}

	//Here is where we should "actually" be catching comand line flags
	// or configuration options to decide which actions and metrics to run.
	//Current architectural plan is to use the same actions+metric set for all
	//created Pokes. Invoke the program with different arguments for heterogenious
	//action/metric sets.
	for _, p := range pokes {
		p.AddAction(poke.DNSLookup)
		p.AddAction(poke.GetRequest)
		p.Run()
	}
}
