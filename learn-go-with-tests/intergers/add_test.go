// lets try to work with ints today
package tests

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

//it should fail we do not have an Add function
