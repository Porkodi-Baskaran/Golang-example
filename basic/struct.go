package basic

import "fmt"

func Structfunc() {

	// Creating structure
	type Student struct {
		name   string
		branch string
		year   int
	}

	// Creating nested structure
	type Teacher struct {
		name    string
		subject string
		exp     int
		details Student
	}

	// Initializing the fields
	// of the structure
	result := Teacher{
		name:    "Suman",
		subject: "Java",
		exp:     5,
		details: Student{"Bongo", "CSE", 2},
	}

	// Display the values
	fmt.Println("Details of the Teacher")
	fmt.Println("Teacher's name: ", result.name)
	fmt.Println("Subject: ", result.subject)
	fmt.Println("Experience: ", result.exp)

	fmt.Println("\nDetails of Student")
	fmt.Println("Student's name: ", result.details.name)
	fmt.Println("Student's branch name: ", result.details.branch)
	fmt.Println("Year: ", result.details.year)
}

// 	// pointers in structure
// 	type Employee struct {
// 		firstName, lastName string
// 		age, salary         int
// 	}

// 	// passing the address of struct variable
// 	// emp8 is a pointer to the Employee struct
// 	emp8 := &Employee{"Sam", "Anderson", 55, 6000}

// 	// (*emp8).firstName is the syntax to access
// 	// the firstName field of the emp8 struct
// 	fmt.Println("First Name:", emp8.firstName)
// 	fmt.Println("Age:", (*emp8).age)
// }

// 	// initialize the function Rectangle
// 	type Rectangle func(int, int) int

// 	// create structure
// 	type rectanglePara struct {
// 		length  int
// 		breadth int
// 		color   string

// 		// function as a field of struct
// 		rect Rectangle
// 	}
// 	// assign values to struct variables
// 	result := rectanglePara{
// 		length:  10,
// 		breadth: 20,
// 		color:   "Red",
// 		rect: func(length int, breadth int) int {
// 			return length * breadth
// 		},
// 	}

// 	fmt.Println("Color of Rectangle: ", result.color)
// 	fmt.Println("Area of Rectangle: ", result.rect(result.length, result.breadth))
// }

// type Student struct {
// 	ID      int
// 	name    string
// 	class   int
// 	address string
// }

// var stud1 Student
// var stud2 Student

// stud1.ID = 1
// stud1.name = "Aishu"
// stud1.class = 12
// stud1.address = "Chennai"

// stud2.ID = 2
// stud2.name = "ABC"
// stud2.class = 12
// stud2.address = "Madurai"

// PrintStudent(stud1)
// PrintStudent(stud2)

// func (stud Student) {

// 	fmt.Println("ID: ", stud.ID)
// 	fmt.Println("Name: ", stud.name)
// 	fmt.Println("Class: ", stud.class)
// 	fmt.Println("Address: ", stud.address)
// }

// func(i int) {
// 	fmt.Println("Hello", i)
// }(10)

// }

// 	type Person struct {
// 		name   string
// 		age    int
// 		job    string
// 		salary int
// 	}

// 	var pers1 Person
// 	var pers2 Person

// 	// Pers1 specification
// 	pers1.name = "Hege"
// 	pers1.age = 45
// 	pers1.job = "Teacher"
// 	pers1.salary = 6000

// 	// Pers2 specification
// 	pers2.name = "Cecilie"
// 	pers2.age = 24
// 	pers2.job = "Marketing"
// 	pers2.salary = 4500

// 	// Print Pers1 info by calling a function
// 	printPerson(pers1)

// 	// Print Pers2 info by calling a function
// 	printPerson(pers2)
// }

// var pers3 Person

// func printPerson(pers Person) {
// 	fmt.Println("Name: ", pers.name)
// 	fmt.Println("Age: ", pers.age)
// 	fmt.Println("Job: ", pers.job)
// 	fmt.Println("Salary: ", pers.salary)
// }
