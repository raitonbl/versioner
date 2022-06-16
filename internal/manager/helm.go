package manager

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Helm struct {
}

func (instance Helm) GetSupportTypes() []string {
	return []string{"appVersion", "version"}
}

func (instance Helm) GetVersion(object string, filename string) (string, error) {

	if object != "appVersion" && object != "version" {
		return "", errors.New("object[name='" + object + "'] isn't supported")
	}

	binary, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(binary, &m)

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

func (instance Helm) SetVersion(object string, filename string, value string) error {

	if object != "appVersion" && object != "version" {
		return errors.New("object[name='" + object + "'] isn't supported")
	}

	binary, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal(binary, &m)

	if err != nil {
		return err
	}

	m[object] = value

	binary, err = yaml.Marshal(m)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, binary, 0644)

	if err != nil {
		return err
	}

	return nil
}
