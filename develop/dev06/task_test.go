package main

import (
	"testing"
)

type testCut struct {
	f, d           string
	s              bool
	data           string
	expectedStr    string
	expectedErr    error
}

var testCuts = []testCut{
	{"0, 2,3", " ", false, "Lizzie_Blose_lblose0@apache.org_Female_29\nLizzie Blose ablose0@apache.org Female 29\nJerome Severs jsevers1@utexas.edu Male 59", "Lizzie_Blose_lblose0@apache.org_Female_29\nLizzie ablose0@apache.org Female\nJerome jsevers1@utexas.edu Male\n", nil},
	{"1,2", " ", false, "Lizzie_Blose_lblose0@apache.org_Female_29\nLizzie Blose ablose0@apache.org Female 29\nJerome Severs jsevers1@utexas.edu Male 59", "Lizzie_Blose_lblose0@apache.org_Female_29\nBlose ablose0@apache.org\nSevers jsevers1@utexas.edu\n", nil},
	{"1,2", "_", true, "Lizzie_Blose_lblose0@apache.org_Female_29\nLizzie Blose ablose0@apache.org Female 29\nJerome Severs jsevers1@utexas.edu Male 59", "Blose lblose0@apache.org\n", nil},
}

func TestCut(t *testing.T) {
	for _, test := range testCuts {
		output, err := Cut(test.data, test.f, test.d, test.s)
		if test.expectedStr != output || test.expectedErr != err {
			t.Errorf("Test failed: expected - (%s, %v), got - (%s, %v)", test.expectedStr, test.expectedErr, output, err)
		}
	}
}
