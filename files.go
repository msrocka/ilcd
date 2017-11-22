package ilcd

import (
	"encoding/xml"
	"io/ioutil"
)

// ReadProcessFile reads a process data set from the given file.
func ReadProcessFile(filePath string) (*Process, error) {
	process := &Process{}
	err := readFile(filePath, process)
	return process, err
}

// ReadMethodFile reads a LCIA method data set from the given file.
func ReadMethodFile(filePath string) (*Method, error) {
	method := &Method{}
	err := readFile(filePath, method)
	return method, err
}

// ReadFlowFile reads a flow data set from the given file.
func ReadFlowFile(filePath string) (*Flow, error) {
	flow := &Flow{}
	err := readFile(filePath, flow)
	return flow, err
}

// ReadFlowPropertyFile reads a flow property data set from the given file.
func ReadFlowPropertyFile(filePath string) (*FlowProperty, error) {
	prop := &FlowProperty{}
	err := readFile(filePath, prop)
	return prop, err
}

// ReadCategoryFile reads a category system from the given file.
func ReadCategoryFile(filePath string) (*CategorySystem, error) {
	system := &CategorySystem{}
	err := readFile(filePath, system)
	return system, err
}

// ReadContactFile reads a contact data set from the given file
func ReadContactFile(filePath string) (*Contact, error) {
	c := &Contact{}
	err := readFile(filePath, c)
	return c, err
}

// ReadSourceFile reads a source data set from the given file
func ReadSourceFile(filePath string) (*Source, error) {
	s := &Source{}
	err := readFile(filePath, s)
	return s, err
}

func readFile(filePath string, dataSet interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, dataSet)
}
