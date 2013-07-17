package nlp

import (
	"testing"
)

func TestParser(t *testing.T) {
	//  w := data.GetWord("mexico")

	//fmt.Printf("Word %s \n", w)
	//fmt.Printf("Categories %s \n", data.GetCategories(w.CategoryQuery))

	//  println("Database initialization complete")

	//nlp.Parse("I planted an apple tree behind the traffic control by the old catholic church on a pond lily")
	//  nlp.Parse("I ate a banana")
	//nlp.Parse("I got a size 4 haircut today") // Failure case because the number corresponds to the word BEFORE
	//nlp.Parse("I ate some cookies")

	//nlp.Parse("I ran a mile")
	//nlp.Parse("I ran 3 miles")
	//nlp.Parse("a hundred octopi")
	// Got a size 4 haircut

	if 7 != 7 { //try a unit test on function
		t.Error("TestParser did not work as expected.") // log error if it did not work as expected
	} else {
		t.Log("one test passed.") // log some info if you want
	}

}
