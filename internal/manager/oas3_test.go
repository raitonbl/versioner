package manager

import (
	"io/ioutil"
	"path"
	"testing"
)

func Test_SetVersion_and_GetVersion_From_OAS3(t *testing.T) {
	doTest(t, func(directory string) {
		filename := "testdata/oas3/specification.yaml"
		binary, err := ioutil.ReadFile(filename)

		if err != nil {
			t.Fatal(err)
		}

		dst := path.Join(directory, "specification.yaml")
		err = ioutil.WriteFile(dst, binary, 0644)

		if err != nil {
			t.Fatal(err)
		}

		manager := OAS3{}

		err = manager.SetVersion("version", dst, "3.0.0")

		if err != nil {
			t.Fatal(err)
		}

		v, prob := manager.GetVersion("version", dst)

		if prob != nil {
			t.Fatal(err)
		}

		if v != "3.0.0" {
			t.Fatal("expected 3.0.0 but got " + v)
		}

	})
}
