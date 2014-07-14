Usage: ./expsim --factors <see factor format> --result <result fields> --reps <number of reps>
Outputs statdata-n.csv

Models: 
0: single queue, single server
1: single queue, multiple server
2: multiple queue, multiple server
3: per-server queues

Factors: models, queues, servers, customers, iatime (average interarrival time), stime (average serving time)
Factor format: "factorname:val1,val2,val3:factorname:val1,val2..."

Results: wait, probwait

Bounds:
stime and iatime both have upper/lower bounds at 0/100.
Don't put 0 for servers, queues, or customers. Idiot.
