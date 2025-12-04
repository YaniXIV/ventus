package merkle

import (
	"fmt"
	"testing"
)

func TestCalcDepth(t *testing.T) {
	z := InitGSMTWithSHA256()

	var testData []byte = []byte("TestData")

	for range 50 {
		z.AddGSMT(testData)
	}

	fmt.Println(len(z.Levels[0]))
	got := z.CalcDepth()
	wanted := 6
	if wanted != got {
		t.Errorf("Error, got %v, wanted %v", got, wanted)
	}
}
