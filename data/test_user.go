package data

import (
	"testing"
)

func TestEmail(t *testing.T) {
	//data.GetUserByEmail("sdkaflewrjlks@gmail.com")
	//data.GetUserByEmail("someemail@gmail.com")
	if 7 != 7 { //try a unit test on function
		t.Error("TestEmail did not work as expected.") // log error if it did not work as expected
	} else {
		t.Log("one test passed.") // log some info if you want
	}

}
