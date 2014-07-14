package main

import (
	"runtime"
	"sync"
	"os"
	"fmt"
	"encoding/csv"
)

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	conf := RunFlags()
	sims := buildSims(conf)
	// Run all simulators
	wg := &sync.WaitGroup{}
	numSims := len(sims)
	for i := 0; i < numSims; i += cpus {
		for j := 0; j < cpus && i+j < numSims; j++ {
			v := sims[j+i]
			wg.Add(1)
			go func() {
				v.Run()
				wg.Done()
			}()
		}
		wg.Wait()
	}
	output := AllStrings(sims)
	var file *os.File
	for i := 0; ; i++ {
		filename := fmt.Sprintf("statdata-%d.csv",i)
		if _,err := os.Stat(filename); os.IsNotExist(err) {
			file,err = os.Create(filename)
			if err != nil {
				panic(err)
			}
			break
		}
	}

	writeout := csv.NewWriter(file)
	err := writeout.WriteAll(output)
	if err != nil {
		panic(err)
	}
}
