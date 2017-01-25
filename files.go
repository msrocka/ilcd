package ilcd

import (
	"encoding/xml"
	"io/ioutil"
)

// ReadProcessFile reads a process data set from the given file.
func ReadProcessFile(path string) (*Process, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	process := &Process{}
	err = xml.Unmarshal(data, process)
	return process, err
}

// ReadFlowFile reads a flow data set from the given file.
func ReadFlowFile(path string) (*Flow, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	flow := &Flow{}
	err = xml.Unmarshal(data, flow)
	return flow, err
}
