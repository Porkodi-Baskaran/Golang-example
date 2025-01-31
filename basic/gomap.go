package basic

import "fmt"

func GoMapfun() {
	var food = make(map[string]string)
	food["morning"] = "breakfast"
	food["afternoon"] = "lunch"
	food["night"] = "Dinner"

	fmt.Println(food["morning"])

	// food["veg"] = "Mushroom"
	// food["night"] = "Night Dinner"
	// delete(food, "veg")

	val1, ok1 := food["afternoon"]
	val2, ok2 := food["brunch"]
	val3 := food["night"]

	fmt.Println(food)
	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3)

}
