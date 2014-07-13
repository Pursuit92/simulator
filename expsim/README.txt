Usage: ./expsim --models <comma separated models> --factors <see factor format> --result <result field>

Models: sqms (single queue, multiple server), mqms (multiple queue, multiple server), psq (per-server queues)

Factors: queues, servers, customers, iatime (average interarrival time), stime (average serving time)
Factor format: "factorname:val1,val2,val3:factorname:val1,val2..."

Results: wait, probwait
