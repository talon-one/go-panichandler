package panichandler_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/require"
)

func TestOnPanic(t *testing.T) {
	var testFiles []string
	// collect all go files and run them
	err := filepath.Walk("./tests/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}
		testFiles = append(testFiles, path)
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	// run all test files
	// each test file will return an json encoded object that contains "expected" and "actual" fields
	// if they match each other the test passes
	for _, testFile := range testFiles {
		t.Run(testFile, func(t *testing.T) {
			cmd := exec.Command("go", "run", testFile)
			cmd.Env = os.Environ()
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("%v: output is %s", err, string(out))
			}
			var m map[string]string
			require.NoError(t, json.Unmarshal(out, &m))
			require.Equal(t, m["expected"], m["actual"])
		})
	}
}
