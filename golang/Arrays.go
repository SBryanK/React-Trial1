package main

import (
	"fmt"
)

//Adding function for later use
func say(n string) {
  fmt.Printf("Hello %v \n", n)
}

//for loop in fucntion
func Names(n []string, f func(string)){
  for _, v := range n {
    f(v)
  }
} // _ nmeans do not consider (we only want the v), := means in (waslrus operator -> to assign a value to a variable within an expression)

func init() {
  var arr1 = [3]int{1,2,3} // decide the lengths [3] and the value
  arr2 := [6]int{4,5,6,7,8,9}

  arr1[2] = 10 //Chnaging value of third element of arr1

  var arr3 = []int{100,101,102}

  fmt.Println(arr1)
  fmt.Println(arr2)
  fmt.Println(arr1[1]) //Access the second elemnt of arr1
  fmt.Println(arr3)

  // For loop
  x := 0
  for x < 5 {
    fmt.Println("x is: ",x)
    x++
  }

  // Advanced loop
  for i:=0; i <= 100; i+=10 {
    if i ==10 {
      continue
    }
    fmt.Println(i)

    }
  // If else
  a := 20
  b := 19
  if a > b {
    fmt.Println("x is greater than y")
    if a == b{
      fmt.Println("x is equal to y")
    } 
  } else {
    fmt.Println("No Thanks Bro")
  }

  //Using Function 
  say("Koko")

  
}