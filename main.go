package main

import "fmt"
import "sync"


func main() {
	fmt.Println("hello world")

	var wg sync.WaitGroup
	wg.Add(1)

	go simple_function(&wg)

	fmt.Println("Waiting To Finish All Threads")
    wg.Wait()

    fmt.Println("All Done")
}

func simple_function(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("my simple function")
	
	int_val, float32_val, string_val := multiple_return_values()
	
	fmt.Println("int=", int_val, ", float32=", float32_val, ", string=", string_val)
}

func multiple_return_values() (int, float32, string) {
	return 101, 499.99, "my function with multiple retrun values"
}
