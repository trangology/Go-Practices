### Parallel processing in Go works similarly Pool.map in Python

My task using [GOMAXPROCS](https://github.com/dmuth/google-go-cpu-usage/blob/master/main.go), [goroutine](https://gobyexample.com/goroutines) and [WaitGroup](https://stackoverflow.com/questions/19208725/example-for-sync-waitgroup-correct). 

Firstly, find memory comsumption of function **Compute()**. <br>
Because of memory usage limited (in this task is 15MB), 
then need to find maximum number of workers we can use at the same time. <br>
Compare it with **min_workers**, **max_workers** and **number of cores available**. <br>
Finally, use the number found with **Gomaxprocs**. 
