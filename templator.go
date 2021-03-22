package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"gopkg.in/yaml.v2"

	"templator/editor"
	"templator/template"
)

func main() {
	var templatePath, outputFile string
	flag.StringVar(&templatePath, "template", "", "Path to the template file")
	flag.StringVar(&outputFile, "outputFile", "", "Path of the output file")
	flag.Parse()

	content, err := ioutil.ReadFile(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	t := template.GetTemplate(string(content))
	rv := template.GetAllRequiredVariables(t)
	b, err := editor.CaptureInputFromEditor(rv)
	if err != nil {
		panic(err)
	}

	var data map[string]interface{}
	err = yaml.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}

	d := template.CreateDocument(t, data)

	if outputFile == "" {
		fmt.Println(d)
	} else {
		err := WriteToFile(outputFile, d)
		if err != nil {
			panic(fmt.Errorf("Failed to write output to file '%s' due to '%v'", outputFile, err))
		}
	}
}

// WriteToFile will print any string of text to a file safely by
// checking for errors and syncing at the end (used from https://golangcode.com/writing-to-file/).
func WriteToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}

	return file.Sync()
}
