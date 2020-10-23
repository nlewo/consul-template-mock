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

			if err := mockFromFilename("examples/"+v+".tmpl", "examples/"+v+".json", rw); err != nil {
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

func TestMissing(t *testing.T) {
	rw := new(bytes.Buffer)

	template := []byte("{{- with secret \"a\" -}}{{- .Data.d }}{{- end }}")
	mockData := []byte("{ \"secret\": {\"a\": {\"b\": \"c\"}}}")

	if err := mock(template, mockData, rw); err == nil {
        if rw.String() != "<no value>" {
            t.Errorf("Keys of map must be present")
        }
	}
}
