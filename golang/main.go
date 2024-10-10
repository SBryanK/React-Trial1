package main

import "fmt"

//making constant for later use --> Thisis unchangeable
const A int = 19

func main() {

	//Strings
	var name string = "Bryan"
	var name2 = "Kusno"
	name3 := "guo"
	fmt.Println(name, name2, name3)

	//integers
	age := 17
	var age2 int = 20

	fmt.Println(age, age2)

	//Bits and Memory
	var num int8 = 127
	num2 := 256 // the compiler decides the type of variabel to assign
	var score float32 = 12345.9

	fmt.Println(num, num2, score)

	//print
	fmt.Print("hello, ")
	fmt.Print("world") // will display hello, world

	//combine print
	fmt.Print("hello, ", name, "you are ", age2, " years old")

	var i, j string = "Learn", "Golang"

	fmt.Println(i, j)

	//printf (must use '\n')

	fmt.Printf("i has value: %v and type: %T\n", name, name)
	fmt.Printf("j has value: %v and type: %T\n", age, age)
	fmt.Printf("My age is %q and my name is %q\n", age, name)

	//Constant
	fmt.Println(A)

	var arr1 = [3]int{1, 2, 3} // decide the lengths [3] and the value
	arr2 := [6]int{4, 5, 6, 7, 8, 9}

	arr1[2] = 10 //Chnaging value of third element of arr1

	var arr3 = []int{100, 101, 102}
	arr3 = append(arr3, 103)

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr1[1]) //Access the second elemnt of arr1
	fmt.Println(arr3)
}
