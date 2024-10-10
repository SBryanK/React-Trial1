package main
import ("fmt")

// switch function
// to select one of many code blocks to be executed.
func single() {
  day := 4 // number 4 will be executed

  switch day {
  case 1:
    fmt.Println("Monday")
  case 2:
    fmt.Println("Tuesday")
  case 3:
    fmt.Println("Wednesday")
  case 4:
    fmt.Println("Thursday")
  case 5:
    fmt.Println("Friday")
  case 6:
    fmt.Println("Saturday")
  case 7:
    fmt.Println("Sunday")
  }
}

// Return specified type
func myFunction(x int, y int) int {
	return x + y
  }

//Stuture
//A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

type Person struct {
	name string
	age int
	job string
	salary int
  }
  
  func specs() {
	var pers1 Person
	var pers2 Person
  
	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000
  
	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500
  
	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)
  
	// Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)
  }

//Map
//Maps are used to store data values in key:value pairs.
func maps(){
	a := make(map[string]string) // use := for simplicity, so no need to specify variable a type
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
  
	fmt.Printf(a["brand"])
  }