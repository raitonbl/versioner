package manager

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type OAS3 struct {
}

func (instance OAS3) GetSupportTypes() []string {
	return []string{"version"}
}

func (instance OAS3) GetVersion(object string, filename string) (string, error) {

	if object != "version" {
		return "", errors.New("object[name='" + object + "'] isn't supported")
	}

	binary, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	m := yaml.MapSlice{}
	err = yaml.Unmarshal(binary, &m)

	if err != nil {
		return "", err
	}

	for _, each := range m {
		if each.Key == "info" {
			for _, s := range each.Value.(yaml.MapSlice) {
				if s.Key == "version" {
					return s.Value.(string), nil
				}
			}
		}
	}

	return "1.0.0", nil
}

func (instance OAS3) SetVersion(object string, filename string, value string) error {

	if object != "version" {
		return errors.New("object[name='" + object + "'] isn't supported")
	}

	binary, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	m := yaml.MapSlice{}
	err = yaml.Unmarshal(binary, &m)

	if err != nil {
		return err
	}

	infoIndex := -1
	var infoObject yaml.MapSlice

	for index, each := range m {
		if each.Key == "info" {
			infoIndex = index
			infoObject = each.Value.(yaml.MapSlice)
		}
	}

	if infoObject == nil {
		infoObject = make([]yaml.MapItem, 0)
	}

	isAssigned := false

	for index, item := range infoObject {
		if item.Key == "version" {
			infoObject[index] = yaml.MapItem{Key: "version", Value: value}
			isAssigned = true
			break
		}
	}

	if !isAssigned {
		infoObject = append(infoObject, yaml.MapItem{Key: "version", Value: value})
	}

	if infoIndex == -1 {
		m = append(m, yaml.MapItem{Key: "info", Value: infoObject})
	} else {
		m[infoIndex] = yaml.MapItem{Key: "info", Value: infoObject}
	}

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
