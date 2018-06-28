package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"text/template"
)

type Secret struct {
	Data map[string]string
}

type Input struct {
	Service map[string]interface{}
	Key     map[string]string
	Env     map[string]string
	File    map[string]string
	Secret  map[string]map[string]string
}

func mockFromFilename(templateFileName, inputDataFileName string, wr io.Writer) error {
	templateText, err := ioutil.ReadFile(templateFileName)
	if err != nil {
		return fmt.Errorf("reading file %s: %s", templateFileName, err)
	}

	inputData, err := ioutil.ReadFile(inputDataFileName)
	if err != nil {
		return fmt.Errorf("reading file %s: %s", inputDataFileName, err)
	}

	return mock(templateText, inputData, wr)
}

func mock(templateText, inputData []byte, wr io.Writer) error {
	var input Input
	err := json.Unmarshal(inputData, &input)
	if err != nil {
		return fmt.Errorf("parsing json input : %s", err)
	}

	funcMap := template.FuncMap{
		"service": func(s string) (interface{}, error) {
			if i, ok := input.Service[s]; ok {
				return i, nil
			}
			return nil, fmt.Errorf("service '%s' doesn't exist: %s", s, err)
		},
		"secret": func(s string) (Secret, error) {
			if i, ok := input.Secret[s]; ok {
				return Secret{Data: i}, nil
			}
			return Secret{}, fmt.Errorf("secret path %s doesn't exist", s)
		},
		"file": func(fileName string) (string, error) {
			if i, ok := input.File[fileName]; ok {
				return i, nil
			}
			return "", fmt.Errorf("file '%s' doesn't exist: %s", fileName, err)
		},
		"key": func(key string) (string, error) {
			if i, ok := input.Key[key]; ok {
				return i, nil
			}
			return "", fmt.Errorf("key '%s' doesn't exist: %s", key, err)
		},
		"keyOrDefault": func(key, def string) string {
			if i, ok := input.Key[key]; ok {
				return i
			}
			return def
		},
		"parseJSON": func(j string) (interface{}, error) {
			var f interface{}
			err := json.Unmarshal([]byte(j), &f)
			if err != nil {
				return nil, fmt.Errorf("parsing JSON: %s '%s'", err, j)
			}
			return f, nil
		},
		"env": func(venv string) (string, error) {
			if i, ok := input.Env[venv]; ok {
				return i, nil
			}
			return "", fmt.Errorf("env variable '%s' doesn't exist: %s", venv, err)
		},
		// regexReplaceAll replaces all occurrences of a regular expression with
		// the given replacement value.
		"regexReplaceAll": func(re, pl, s string) (string, error) {
			compiled, err := regexp.Compile(re)
			if err != nil {
				return "", err
			}
			return compiled.ReplaceAllString(s, pl), nil
		},
		// regexMatch returns true or false if the string matches
		// the given regular expression
		"regexMatch": func(re, s string) (bool, error) {
			compiled, err := regexp.Compile(re)
			if err != nil {
				return false, err
			}
			return compiled.MatchString(s), nil
		},
	}

	tmpl, err := template.New("template").Funcs(funcMap).Option("missingkey=error").Parse(string(templateText))
	if err != nil {
		return fmt.Errorf("parsing template: %s", err)
	}

	err = tmpl.Execute(wr, "")
	if err != nil {
		return fmt.Errorf("execution: %s", err)
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("consul-template-mock - Render consul-template from JSON file datasource\n")
		fmt.Printf("Usage: consul-template-mock template-file.tmpl input-data-file.json")
		os.Exit(1)
	}

	if err := mockFromFilename(os.Args[1], os.Args[2], os.Stdout); err != nil {
		log.Fatal(err)
	}

}
