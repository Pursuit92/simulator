package main

import (
	"os"
	"encoding/csv"
)

func main() {
	conf := RunFlags()
	sims := buildSims(conf)
	// Run all simulators
	for _,v := range sims {
		v.Run()
	}
	output := AllStrings(sims)
	stdout := csv.NewWriter(os.Stdout)
	err := stdout.WriteAll(output)
	if err != nil {
		panic(err)
	}
}
