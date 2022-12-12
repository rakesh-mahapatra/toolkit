package toolkit

import "testing"

func TestTools_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandomString(20)
	if len(s) != 20 {
		t.Error("Wrong len of string returned.")
	}
}
