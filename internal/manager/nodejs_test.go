package manager

import (
	"io/ioutil"
	"path"
	"testing"
)

func Test_SetVersion_and_GetVersion_From_Nodejs(t *testing.T) {
	doTest(t, func(directory string) {
		filename := "testdata/nodejs/package.json"
		binary, err := ioutil.ReadFile(filename)

		if err != nil {
			t.Fatal(err)
		}

		dst := path.Join(directory, "package.json")
		err = ioutil.WriteFile(dst, binary, 0644)

		if err != nil {
			t.Fatal(err)
		}

		manager := Nodejs{}

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
