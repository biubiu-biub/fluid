package utils

import (
	"strings"
	"testing"
)

func TestNewReleaseName(t *testing.T) {
	var tests=[]struct{
		pStr string
		expectedStr string
	}{
		{"load","load-load-"},
		{"","-load-"},
		{"pkg","pkg-load-"},
	}
	for _,test:=range tests{
		output:=NewReleaseName(test.pStr)
		if !strings.HasPrefix(output,test.expectedStr){
			t.Errorf("supposed output should starts with %s,real output is %s",test.expectedStr,output)
		}
	}
}
func TestGetJobNameFromReleaseName(t *testing.T) {
	var tests=[]struct{
		pStr string
		expectedStr string
	}{
		{"load-load-12345","load-loader-job-12345"},
		{"-load-12345","-loader-job-12345"},
		{"pkg-load-12345","pkg-loader-job-12345"},
	}
	for _,test:=range tests{
		output:=GetJobNameFromReleaseName(test.pStr)
		if output!=test.expectedStr{
			t.Errorf("supposed output should be %s,real output is %s",test.expectedStr,output)
		}
	}
}