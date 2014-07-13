package main

const (
	// models
	SQSS string = "sqss"
	MQMS = "mqms"
	SQMS = "sqms"
	PSQ = "psq"

	// Factors
	Queues = "queues"
	Servers = "servers"
	Customers = "customers"
	IATime = "iatime"
	STime = "stime"

	// Results
	Wait = "wait"
	ProbWait = "probwait"
)

var (
	Models []string = []string{SQSS,MQMS,SQMS,PSQ}
	Factors = []string{Queues,Servers,Customers,IATime,STime}
	Results = []string{Wait,ProbWait}
)
