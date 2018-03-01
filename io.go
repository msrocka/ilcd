package ilcd

import (
	"encoding/xml"
	"io/ioutil"
)

// ReadModelFile reads a life cycle model from the given file.
func ReadModelFile(filePath string) (*Model, error) {
	m := &Model{}
	err := readFile(filePath, m)
	return m, err
}

// ReadModel reads a life cycle model from the given data.
func ReadModel(data []byte) (*Model, error) {
	m := &Model{}
	err := xml.Unmarshal(data, m)
	return m, err
}

// ReadProcessFile reads a process data set from the given file.
func ReadProcessFile(filePath string) (*Process, error) {
	p := &Process{}
	err := readFile(filePath, p)
	return p, err
}

// ReadProcess reads a process data set from the given data
func ReadProcess(data []byte) (*Process, error) {
	p := &Process{}
	err := xml.Unmarshal(data, p)
	return p, err
}

// ReadMethodFile reads a LCIA method data set from the given file.
func ReadMethodFile(filePath string) (*Method, error) {
	m := &Method{}
	err := readFile(filePath, m)
	return m, err
}

// ReadMethod reads a LCIA method data set from the given data.
func ReadMethod(data []byte) (*Method, error) {
	m := &Method{}
	err := xml.Unmarshal(data, m)
	return m, err
}

// ReadFlowFile reads a flow data set from the given file.
func ReadFlowFile(filePath string) (*Flow, error) {
	flow := &Flow{}
	err := readFile(filePath, flow)
	return flow, err
}

// ReadFlow reads a LCIA method data set from the given data.
func ReadFlow(data []byte) (*Flow, error) {
	f := &Flow{}
	err := xml.Unmarshal(data, f)
	return f, err
}

// ReadFlowPropertyFile reads a flow property data set from the given file.
func ReadFlowPropertyFile(filePath string) (*FlowProperty, error) {
	prop := &FlowProperty{}
	err := readFile(filePath, prop)
	return prop, err
}

// ReadFlowProperty reads a flow property data set from the given data.
func ReadFlowProperty(data []byte) (*FlowProperty, error) {
	fp := &FlowProperty{}
	err := xml.Unmarshal(data, fp)
	return fp, err
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

// ReadContact reads a contact data set from the given data
func ReadContact(data []byte) (*Contact, error) {
	c := &Contact{}
	err := xml.Unmarshal(data, c)
	return c, err
}

// ReadSourceFile reads a source data set from the given file
func ReadSourceFile(filePath string) (*Source, error) {
	s := &Source{}
	err := readFile(filePath, s)
	return s, err
}

// ReadSource reads a source data set from the given data
func ReadSource(data []byte) (*Source, error) {
	s := &Source{}
	err := xml.Unmarshal(data, s)
	return s, err
}

// ReadUnitGroupFile reads a unit group data set from the given file
func ReadUnitGroupFile(filePath string) (*UnitGroup, error) {
	ug := &UnitGroup{}
	err := readFile(filePath, ug)
	return ug, err
}

// ReadUnitGroup reads a unit group data set from the given data
func ReadUnitGroup(data []byte) (*UnitGroup, error) {
	ug := &UnitGroup{}
	err := xml.Unmarshal(data, ug)
	return ug, err
}

func readFile(filePath string, dataSet interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, dataSet)
}
