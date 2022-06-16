package manager

import (
	"io/ioutil"
	"path"
	"testing"
)

func Test_SetVersion_and_GetVersion_From_Maven(t *testing.T) {
	doTest(t, func(directory string) {
		filename := "testdata/maven/pom.xml"
		binary, err := ioutil.ReadFile(filename)

		if err != nil {
			t.Fatal(err)
		}

		dst := path.Join(directory, "pom.xml")
		err = ioutil.WriteFile(dst, binary, 0644)

		if err != nil {
			t.Fatal(err)
		}

		manager := Maven{}

		err = manager.SetVersion("version", dst, "2.0.0")

		if err != nil {
			t.Fatal(err)
		}

		v, prob := manager.GetVersion("version", dst)

		if prob != nil {
			t.Fatal(err)
		}

		if v != "2.0.0" {
			t.Fatal("expected 2.0.0 but got " + v)
		}

	})
}
