package manager

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/antchfx/xmlquery"
	"github.com/creekorful/mvnparser"
	"io/ioutil"
	"strings"
)

type Maven struct {
}

func (instance Maven) GetSupportTypes() []string {
	return []string{"version"}
}

func (instance Maven) GetVersion(object string, filename string) (string, error) {

	if object != "version" {
		return "", errors.New("object[name='" + object + "'] isn't supported")
	}

	var project mvnparser.MavenProject

	container, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", err
	}

	if ex := xml.Unmarshal(container, &project); ex != nil {
		return "", ex
	}

	return project.Version, nil
}

func (instance Maven) SetVersion(object string, filename string, value string) error {

	if object != "version" {
		return errors.New("object[name='" + object + "'] isn't supported")
	}

	var project mvnparser.MavenProject

	container, ex := ioutil.ReadFile(filename)

	if ex != nil {
		return ex
	}

	if err := xml.Unmarshal(container, &project); err != nil {
		return err
	}

	content, ex := ioutil.ReadFile(filename)

	file, err := setVersion(content, value)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, []byte(file), 0644)

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
