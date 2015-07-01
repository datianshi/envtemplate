package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestEnvToMap(t *testing.T) {
	os.Setenv("ENV1", "VALUE1")
	os.Setenv("ENV2", "VALUE2")
	retv := envToMap()
	if retv["ENV1"] != "VALUE1" || retv["ENV2"] != "VALUE2" {
		t.Log("Failed to transform the envirables to the map")
		t.Log(retv)
		t.Fail()
	}
}

const expected = `host=hostname
username=username
`

func TestResolveEnv(t *testing.T) {
	os.Setenv("SMTP_USERNAME", "username")
	os.Setenv("SMTP_HOST", "hostname")
	var wr bytes.Buffer

	err := resolveEnv(&wr, "spec/sample")
	if err != nil {
		t.Error(err)
	}

	if wr.String() != expected {
		t.Log(fmt.Sprintf("The exepcted string: \n%s is not matching:\n %s", expected, wr.String()))
		t.Fail()
	}
}
