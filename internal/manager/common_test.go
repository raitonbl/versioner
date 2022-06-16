package manager

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func doTest(t *testing.T, fn func(directory string)) {
	pid := fmt.Sprintf("%d", os.Getpid())

	dir, err := ioutil.TempDir("", fmt.Sprintf("versioner-%s", pid))

	if err != nil {
		t.Fatal(err)
	}

	fn(dir)

	defer func(path string) {
		_ = os.RemoveAll(path)
	}(dir)
}
