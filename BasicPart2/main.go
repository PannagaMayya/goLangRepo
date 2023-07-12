package main

import "fmt"

func main() {
	//MAP
	fmt.Println("MAP")
	mapEx := map[int]string{1: "1", 2: "2"}
	fmt.Println(mapEx)
	//Empty Map
	mapEx = make(map[int]string)
	fmt.Println(mapEx)
	//Insertion
	mapEx[3] = "3"
	mapEx[1] = "-1"
	mapEx[2] = "2"
	fmt.Println(mapEx)
	//Deletion
	delete(mapEx, 3)
	fmt.Println(mapEx)
	//Iteration
	for key, val := range mapEx {
		fmt.Println("key:", key, ", value:", val)
	}
}
