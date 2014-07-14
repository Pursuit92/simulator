package main

const (
	// models
	SQSS int = iota
	SQMS
	PSQ
	MQMS

	// Factors
	Model string = "models"
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
	Factors = []string{Model,Queues,Servers,Customers,IATime,STime}
	Results = []string{Wait,ProbWait}
)
