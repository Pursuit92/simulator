package main

import (
	"flag"
	"strings"
	"strconv"
)

type FlagConf struct {
	Models []string
	Factors map[string][]int
	Results []string
	Reps int
}

func in(i string, l []string) bool {
	for _,v := range l {
		if i == v {
			return true
		}
	}
	return false
}

func allIn(i []string, l []string) bool {
	for _,v := range i {
		if ! in(v,l) {
			return false
		}
	}
	return true
}

func allAll(i,j []string) bool {
	if allIn(i,j) && allIn(j,i) {
		return true
	}
	return false
}

func RunFlags() FlagConf {
	modelStr := flag.String("models","","CSV list of Models")
	factorStr := flag.String("factors","","Factor/value list")
	resultStr := flag.String("results","","CSV list of Results")
	reps := flag.Int("reps",1,"Number of repetitions")
	flag.Parse()

	if *modelStr == "" || *factorStr == "" || *resultStr == "" {
		panic("Missing Option!")
	}

	conf := FlagConf{}
	conf.Models = strings.Split(*modelStr,",")
	if ! allIn(conf.Models,Models) {
		panic("Invalid Model!")
	}
	conf.Results = strings.Split(*resultStr,",")
	if ! allIn(conf.Results,Results) {
		panic("Invalid Result!")
	}
	conf.Factors = make(map[string][]int)
	factorList := strings.Split(*factorStr,":")
	if len(factorList) % 2 != 0 {
		panic("Invalid Factor List!")
	}
	factorNames := make([]string,len(factorList) / 2)
	for i := 0; i < len(factorList); i += 2 {
		j := i+1
		levels := strings.Split(factorList[j],",")
		intLevels := make([]int, len(levels))
		for k,v := range levels {
			i,e := strconv.Atoi(v)
			if e != nil {
				panic("Invalid Factor Level: " + v)
			}
			intLevels[k] = i
		}
		conf.Factors[factorList[i]] = intLevels
		factorNames[i / 2] = factorList[i]
	}
	if ! allAll(factorNames,Factors) {
		panic("Invalid Factor!")
	}
	conf.Reps = *reps
	return conf
}
