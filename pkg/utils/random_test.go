package utils

import (
	"math/rand"
	"testing"
)

func TestRandomAlphaNumberString(t *testing.T) {
	var tests=[]struct{
		pLen int32
		expectedStr string
	}{
		{5,"60q86"},
		{6,"jh268o"},
		{7,"s4fz3g1"},
	}
	//seed with the same value results in the same random string each run
	rand.Seed(20330123)
	for _,test:=range tests{
		output:=RandomAlphaNumberString(test.pLen)
		if output!=test.expectedStr{
			t.Errorf("supposed output should be %s,real output is %s",test.expectedStr,output)
		}
	}
}
