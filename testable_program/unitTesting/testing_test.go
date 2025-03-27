package testable_program

import "testing"

func TestDivide(t *testing.T) {

	got := Divide(10,2)
	if got != 5{
		t.Fatalf("Divide(10,2) = %d; want 5 testcase 1", got)
	}
	vot := Divide(0,10)
	if vot != 0 {
		t.Fatalf("Divide(10,0) = %d; want 0 testcase 2", vot)
	}
	
	if Divide(10,10) != 1{
		t.Fatalf("Divide(10,0) = %d; want 0 testcase 3", vot)
	}
	
	if Divide(10,5) != 2 {
		t.Fatalf("Divide(10,2) = %d; want 5 testcase 4", vot)
	}

	if Divide(22,2) != 11 {
		t.Fatalf("Divide(10,2) = %d; want 5 testcase 4", vot)
	}

}

func TestDivide2(t *testing.T) {
    testCases := []struct {
        a, b     int
        expected int
        desc     string
    }{
        {10, 2, 5, "Divide(10,2)"},
        {0, 10, 0, "Divide(0,10)"},
        {10, 10, 1, "Divide(10,10)"},
        {10, 5, 2, "Divide(10,5)"},
        {22, 2, 11, "Divide(22,2)"},
    }

    for _, tc := range testCases {
		got := Divide(float64(tc.a), float64(tc.b))
        if got != tc.expected {
            t.Errorf("%s = %d; want %d", tc.desc, got, tc.expected)
        }
    }
}

func BenchmarkDivide(b *testing.B) {

    for i := 0; i < b.N; i++ {
        _ = Divide(10, 2) 
    }
}