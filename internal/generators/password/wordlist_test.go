package password

import "testing"

func TestWordsByLengthPopulated(t *testing.T) {
	for length := 3; length<=9; length++ {
		if len(wordsByLength[length]) == 0 {
			t.Errorf("expected words of length %d, found none", length)
		}
	}
}