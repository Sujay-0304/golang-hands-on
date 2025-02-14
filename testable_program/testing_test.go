package testable_program

import "testing"

func TestDivide(t *testing.T) {
	got := Divide(10,2)
	if got != 5 {
		t.Errorf("Divide(10,2) = %d; want 5 testcase 1", got)
	}
	vot := Divide(0,10)
	if vot != 0 {
		t.Errorf("Divide(10,0) = %d; want 0 testcase 2", vot)
	}

	if Divide(10,10) != 1{
		t.Errorf("Divide(10,0) = %d; want 0 testcase 3", vot)
	}

	if Divide(10,5) != 2 {
		t.Errorf("Divide(10,2) = %d; want 5 testcase 4", vot)
	}

}