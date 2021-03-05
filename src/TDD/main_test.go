package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	// "github.com/pjbgf/go-test/should"
)

func TestMain(t *testing.T) {

	asserThat := func(assumption string, command string, expectedOutput string) {
		should := should.New(t)
		tmpfile, _ := ioutil.TempFile("", "calc-fake-stdout.*")
		defer os.Remove(tmpfile.Name())

		os.Stdout = tmpfile
		os.Args = strings.Split(command, " ")

		output, _ := ioutil.ReadFile(tmpfile.Name())
		actualOutput := string(output)

		should.BeEqual(actualOutput, expectedOutput)
	}

	asserThat("Should sum 1 + 1 and return 2", "calc 1 + 1", "sum total 2\n")
}
