package manager

import (
	"encoding/json"
	"errors"
	"github.com/tidwall/sjson"
	"io/ioutil"
)

type Nodejs struct {
}

func (instance Nodejs) GetSupportTypes() []string {
	return []string{"version"}
}

func (instance Nodejs) GetVersion(object string, filename string) (string, error) {

	if object != "version" {
		return "", errors.New("object[name='" + object + "'] isn't supported")
	}

	binary, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	m := make(map[string]interface{})

	err = json.Unmarshal(binary, &m)

	if err != nil {
		return "", err
	}

	v, isString := m[object].(string)

	if !isString {
		return "", errors.New("cannot read object[name='" + object + "'] from file[name='" + filename + "']")
	}

	if v == "" {
		return "1.0.0", nil
	}

	return v, nil
}

func (instance Nodejs) SetVersion(object string, filename string, value string) error {

	if object != "version" {
		return errors.New("object[name='" + object + "'] isn't supported")
	}

	binary, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	v, err := sjson.Set(string(binary), "version", value)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, []byte(v), 0644)

	if err != nil {
		return err
	}

	return nil
}
