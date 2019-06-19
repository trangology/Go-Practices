# Parallel processing in Go works as Pool.map in Python

My task using GOMAXPROCS() [goroutine](https://gobyexample.com/goroutines), and [WaitGroup](https://stackoverflow.com/questions/19208725/example-for-sync-waitgroup-correct). 

Firstly, find memory comsumption of function Compute(). 
Because of memory usage limited (in this case is 15MB), 
then need to find maximum number of workers we can use at the same time.
