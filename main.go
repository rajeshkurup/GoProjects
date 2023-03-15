package main

import (
	"fmt"
	"sync"

	"matrix"
)

func main() {
	fmt.Println("hello world")

	matrix2d := make([][]int, 8)
	for i := 0; i < len(matrix2d); i++ {
		matrix2d[i] = make([]int, 8)
	}

	matrix2d[1][1] = 1 // blocked
	matrix2d[2][2] = 1 // blocked
	matrix2d[1][2] = 1 // blocked
	matrix2d[2][1] = 1 // blocked
	matrix2d[4][4] = 1 // blocked
	matrix2d[3][3] = 1 // blocked
	matrix2d[5][4] = 1 // blocked
	matrix2d[4][5] = 1 // blocked
	matrix2d[5][7] = 1 // blocked

	start := matrix.Pos{X: 7, Y: 7}
	end := matrix.Pos{X: 5, Y: 5}

	matrixTraverser, _ := matrix.MakeMatrixTraverser(matrix2d)
	path, err := matrixTraverser.FindShortestPath(&start, &end)
	if err == nil {
		fmt.Println("")
		fmt.Println("Shortest Path")
		fmt.Println(path)
	} else {
		fmt.Println("")
		fmt.Println("No Path Found")
	}

	fmt.Println("")
	fmt.Println("Matrix")
	for _, matrix := range matrix2d {
		fmt.Println(matrix)
	}

	fmt.Println("")
	fmt.Println("Start")
	fmt.Println(start)

	fmt.Println("")
	fmt.Println("End")
	fmt.Println(end)
	fmt.Println("")

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
