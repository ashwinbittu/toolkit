package toolkit

import (
	"fmt"
	"testing"
)

func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandomString(15)
	if len(s) != 15 {
		fmt.Println("Wrong length of random string returned")
	}
}
