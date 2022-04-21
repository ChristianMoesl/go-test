package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type File struct {
	Benchmarks []Benchmark `json:"benchmarks"`
}

// The same json tags will be used to encode data into JSON
type Benchmark struct {
	Timestamp  string               `json:"timestamp"`
	Parameters map[string]Parameter `json:"parameters"`
}

type ParameterType string

const (
	Number ParameterType = "Number"
	Text                 = "Text"
)

type Parameter struct {
	Type  ParameterType `json:"type"`
	Value any           `json:"value"`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("data.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
	}

	// we initialize our Users array
	var file File

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &file)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Loaded %d benchmarks\n", len(file.Benchmarks))

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for _, benchmark := range file.Benchmarks {
		fmt.Println("{")
		fmt.Printf("  timestamp: %s\n", benchmark.Timestamp)

		fmt.Println("  {")
		for name, parameter := range benchmark.Parameters {
			switch parameter.Type {
			case Number:
				fmt.Printf("    %s: %f\n", name, parameter.Value)
			case Text:
				fmt.Printf("    %s: %s\n", name, parameter.Value)
			}
		}
		fmt.Println("  }")
		fmt.Println("}")
	}
}
