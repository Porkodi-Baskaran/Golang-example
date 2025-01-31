package basic

import "fmt"

func BasicFn() {
	// 	// call goroutine
	// 	go goroutine("Process 1")

	// 	goroutine("Process 2")
	// }

	// // create a function
	// func goroutine(message string) {

	// 	fmt.Println(message)
	// }

	// func BasicFn() {
	// 	// call goroutine
	// 	go goroutine("Process 1")

	// 	goroutine("Process 2")
	// }
	// // create a function
	// func goroutine(message string) {

	// 	fmt.Println(message)
	// }

	// FOR LOOP
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("i: ", i)
	// }

	// i := 20

	// fmt.Print(i)

	// While loop
	number := 1

	// loop that runs infinitely
	for {

		// condition to terminate the loop
		if number > 5 {
			break
		}

		fmt.Printf("%d\n", number)
		number++

	}
}
