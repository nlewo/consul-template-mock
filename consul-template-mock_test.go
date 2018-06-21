package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestMock(t *testing.T) {
	examplePrefixes := []string{"contrail-api", "skydive", "simple", "trivial"}

	for _, v := range examplePrefixes {
		t.Run(v, func(t *testing.T) {
			rw := new(bytes.Buffer)

			if err := mock("examples/"+v+".tmpl", "examples/"+v+".json", rw); err != nil {
				t.Error(err)
			}

			rendered, _ := ioutil.ReadAll(rw)
			expected, err := ioutil.ReadFile("examples/" + v + ".rendered")
			if err != nil {
				t.Error(err)
			}

			if string(rendered) != string(expected) {
				t.Fatalf("%s != %s", string(rendered), string(expected))
			}
		})
	}
}
