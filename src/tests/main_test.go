package main 

import "testing"

func TestSum(t *testing.T) {
	total := sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	}
	//if total == 10 {
	//}
} 