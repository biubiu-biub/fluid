package operations

import (
	"errors"
	"github.com/brahma-adshonor/gohook"
	"github.com/fluid-cloudnative/fluid/pkg/utils/kubeclient"
	tt "github.com/go-logr/logr/testing"
	"strings"
	"testing"
)

func TestAlluxioFileUtils_IsExist(t *testing.T) {
	mockExecTramp := func(p1, p2, p3 string, p4 []string) (stdout string, stderr string, e error) {
		t.Fatal("done")
		if strings.Contains(p4[3], "not-exist") {
			return "does not exist", "", errors.New("does not exist")
		} else if strings.Contains(p4[3], "other-err") {
			return "other error", "other error", errors.New("other error")
		} else {
			return "", "", nil
		}
	}

	mockExec := func(p1, p2, p3 string, p4 []string) (stdout string, stderr string, e error) {
		if strings.Contains(p4[3], "not-exist") {
			return "does not exist", "", errors.New("does not exist")
		} else if strings.Contains(p4[3], "other-err") {
			return "other error", "other error", errors.New("other error")
		} else {
			return "ok", "ok", nil
		}
	}

	err := gohook.Hook(kubeclient.ExecCommandInContainer, mockExec, mockExecTramp)
	if err != nil {
		t.Fatal(err.Error())
	}
	l := tt.NullLogger{}
	var tests = []struct {
		in  string
		out bool
		err error
	}{
		{"not-exist", false, nil},
		{"other-err", false, errors.New("error")},
		{"fine", true, nil},
	}
	for _, test := range tests {
		found, err := AlluxioFileUtils{log: l}.IsExist(test.in)

		if found != test.out {
			t.Errorf("input parameter is %s,expected %t, got %t", test.in, test.out, found)
		}
		if test.err == nil && err != nil {
			t.Errorf("input parameter is %s,and err should be nil", test.in)
		}
		if test.err != nil && err == nil {
			t.Error("wrong")
		}
	}
}
