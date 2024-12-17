# go-scripts

Based on the book "100 mistakes of go and how to avoid them"

To run benchmarks

`
go test -bench=. -benchmem
`

## instruction_parallelism_optimization.go

In this experiment, function 2 was supposed to be faster than function 1 but according to the benchmarks - sometimes it isnt

![Flowchart](images/instruction_parallelism_optimization_flowchart.jpg)
