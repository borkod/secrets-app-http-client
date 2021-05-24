package cmdtests

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/borkod/secrets-http-client/cmd"
)

func TestCmd(t *testing.T) {

	type testCasesConfig struct {
		args           []string
		expectedResult string
	}

	testCases := []testCasesConfig{
		{
			args:           []string{"create", "-data=mysecret", "-url=http://localhost:8080"},
			expectedResult: "http://localhost:8080/06c219e5bc8378f3a8a3f83b4b7e4649\n",
		},
		{
			args:           []string{"create", "-data=mysecretvalue", "-url=http://localhost:8080"},
			expectedResult: "http://localhost:8080/11eed548a8f140ba4781f3eb11554680\n",
		},
		{
			args:           []string{"view", "-id=06c219e5bc8378f3a8a3f83b4b7e4649", "-url=http://localhost:8080"},
			expectedResult: "mysecret\n",
		},
		{
			args:           []string{"view", "-id=11eed548a8f140ba4781f3eb11554680", "-url=http://localhost:8080"},
			expectedResult: "mysecretvalue\n",
		},
		{
			args:           []string{"view", "-id=11eed548a8f140ba4781f3eb11554680", "-url=http://localhost:8080"},
			expectedResult: "404 Not Found\n",
		},
		{
			args:           []string{"view", "-id=06c219e5bc8378f3a8a3f83b4b7e4649", "-url=http://localhost:8080"},
			expectedResult: "404 Not Found\n",
		},
	}

	for _, tc := range testCases {

		response := captureOutput(cmd.Root, tc.args...)
		if response != tc.expectedResult {
			t.Errorf("Response is %s. Expected: %s", response, tc.expectedResult)
		}
	}
}

func captureOutput(f func([]string) error, args ...string) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stdout)
	}()
	f(args)
	return buf.String()
}
