//today we will deal with iterations and specificaly loops and the one found in Go is a for loop 
//different variations to it but for today just the basic
//lets write a test  that checks if a function prints a character five times

package iteration

import "testing"

func TestRepeat(t *testing.T){
	t.Run("repeat five times", func(t *testing.T){
		repeated:= Repeat("A", 5)
		expected:= "AAAAA"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
	t.Run("more than five times", func(t *testing.T){
		repeated:= Repeat("B", 7)
		expected:= "BBBBBBB"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}


//clear
// this should test that before we make modifications for more conditions
//lets buildour function 
//our test should fail coz our function oly takes one argument at the moment