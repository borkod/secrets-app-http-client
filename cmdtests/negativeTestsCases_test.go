package cmdtests

import (
	"fmt"
	"testing"

	"github.com/borkod/secrets-http-client/cmd"
)

func TestRootBad(t *testing.T) {
	testCase := []string{}

	err := cmd.Root(testCase)
	if !(err != nil) {
		t.Errorf("Expected error. Got nil")
	}

	testCase = []string{"notvalidaction"}

	err = cmd.Root(testCase)
	if !(err != nil) {
		t.Errorf("Expected error. Got nil")
	}

	testCase = []string{"create"}

	err = cmd.Root(testCase)
	if !(err != nil) {
		t.Errorf("Expected error. Got nil")
	}

	testCase = []string{"view"}

	err = cmd.Root(testCase)
	if !(err != nil) {
		t.Errorf("Expected error. Got nil")
	}
}
func TestCreateSecretBad(t *testing.T) {

	type testCaseConfig struct {
		data        string
		url         string
		id          string
		description string
		result      string
		err         bool
	}

	testCasesBad := []testCaseConfig{
		{
			data:        "",
			url:         "",
			id:          "",
			description: "No input test",
			result:      "",
			err:         true,
		},
		{
			data:        "",
			url:         "test2",
			id:          "",
			description: "Bad Protocol",
			result:      "",
			err:         true,
		},
		{
			data:        "",
			url:         "http://localhost:8080",
			id:          "",
			description: "Missing Data",
			result:      "",
			err:         true,
		},
		{
			data:        "secretvalue",
			url:         "",
			id:          "",
			description: "Missing URL",
			result:      "",
			err:         true,
		},
		{
			data:        "secretvalue",
			url:         "http://localhost:8080",
			id:          "asdf",
			description: "ID Provided",
			result:      "",
			err:         true,
		},
	}

	c := cmd.NewCreateSecretCommand()

	for _, tc := range testCasesBad {
		args := []string{"-data=" + tc.data, "-url=" + tc.url}
		if len(tc.id) > 0 {
			args = append(args, "-id="+tc.id)
		}

		fmt.Println(args)
		err := c.Init(args)
		if !(err != nil && tc.err == true) {
			t.Errorf("Expected error. Got nil")
		}
	}
}

func TestGetSecretBad(t *testing.T) {

	type testCaseConfig struct {
		data        string
		url         string
		id          string
		description string
		result      string
		err         bool
	}

	testCasesBad := []testCaseConfig{
		{
			data:        "",
			url:         "",
			id:          "",
			description: "No input test",
			result:      "",
			err:         true,
		},
		{
			data:        "",
			url:         "test2",
			id:          "",
			description: "Bad Protocol",
			result:      "",
			err:         true,
		},
		{
			data:        "",
			url:         "http://localhost:8080",
			id:          "",
			description: "Missing ID",
			result:      "",
			err:         true,
		},
		{
			data:        "",
			url:         "",
			id:          "asdf",
			description: "Missing URL",
			result:      "",
			err:         true,
		},
		{
			data:        "secretvalue",
			url:         "http://localhost:8080",
			id:          "asdf",
			description: "Data Provided",
			result:      "",
			err:         true,
		},
	}

	c := cmd.NewGetSecretCommand()

	for _, tc := range testCasesBad {
		args := []string{"-id=" + tc.id, "-url=" + tc.url}
		if len(tc.data) > 0 {
			args = append(args, "-data="+tc.data)
		}

		fmt.Println(args)
		err := c.Init(args)
		if !(err != nil && tc.err == true) {
			t.Errorf("Expected error. Got nil")
		}
	}
}
