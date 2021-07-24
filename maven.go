package main

import (
	"encoding/xml"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/creekorful/mvnparser"
	"io/ioutil"
	"strings"
)

type Maven struct {
}

func (object Maven) IsSupported(value string) bool {
	return value == "maven"
}

func (object Maven) ReadVersion(value string) (string, error) {
	var project mvnparser.MavenProject

	container, err := ioutil.ReadFile(value)

	if err != nil {
		return "", err
	}

	if ex := xml.Unmarshal(container, &project); ex != nil {
		return "", ex
	}

	return project.Version, nil
}

func (object Maven) GetVersionFile() string {
	return "pom.xml"
}

func (object Maven) EditVersion(path string, newValue string) error {
	var project mvnparser.MavenProject

	container, ex := ioutil.ReadFile(path)

	if ex != nil {
		return ex
	}

	if err := xml.Unmarshal(container, &project); err != nil {
		return err
	}

	content, ex := ioutil.ReadFile(path)

	file, err := setVersion(content, newValue)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, []byte(file), 0644)

	if err != nil {
		return ex
	}

	return nil
}

func setVersion(data []byte, value string) (string, error) {

	doc, err := xmlquery.Parse(strings.NewReader(string(data)))

	if err != nil {
		panic(err)
	}

	vNode := xmlquery.FindOne(doc, "/project/version")

	if vNode != nil {
		xmlquery.RemoveFromTree(vNode)
		vNode = nil
	}

	src, err := xmlquery.Parse(strings.NewReader(fmt.Sprintf("<version>%s</version>", value)))

	if err != nil {
		return "", err
	}

	xmlquery.AddChild(xmlquery.FindOne(doc, "/project"), xmlquery.FindOne(src, "//version"))

	content := doc.OutputXML(false)

	if strings.HasPrefix(content, "<?xml?>") {
		content = content[7:len(content)]
	}

	return content, nil
}
