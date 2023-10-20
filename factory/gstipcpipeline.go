// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Allows splitting and continuing a pipeline in another process
type GstIpcPipelineSink struct {
	*gst.Element
}

func NewIpcPipelineSink() (*GstIpcPipelineSink, error) {
	e, err := gst.NewElement("ipcpipelinesink")
	if err != nil {
		return nil, err
	}
	return &GstIpcPipelineSink{Element: e}, nil
}

func NewIpcPipelineSinkWithName(name string) (*GstIpcPipelineSink, error) {
	e, err := gst.NewElementWithName("ipcpipelinesink", name)
	if err != nil {
		return nil, err
	}
	return &GstIpcPipelineSink{Element: e}, nil
}

// File descriptor to received data from
// Default: -1, Min: -1, Max: 65535
func (e *GstIpcPipelineSink) SetFdin(value int) error {
	return e.Element.SetProperty("fdin", value)
}

func (e *GstIpcPipelineSink) GetFdin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fdin")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// File descriptor to send data through
// Default: -1, Min: -1, Max: 65535
func (e *GstIpcPipelineSink) SetFdout(value int) error {
	return e.Element.SetProperty("fdout", value)
}

func (e *GstIpcPipelineSink) GetFdout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fdout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Read chunk size
// Default: 4096, Min: 1, Max: 16777216
func (e *GstIpcPipelineSink) SetReadChunkSize(value uint) error {
	return e.Element.SetProperty("read-chunk-size", value)
}

func (e *GstIpcPipelineSink) GetReadChunkSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("read-chunk-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum time to wait for a response to a message
// Default: 10000000, Min: 0, Max: 18446744073709551615
func (e *GstIpcPipelineSink) SetAckTime(value uint64) error {
	return e.Element.SetProperty("ack-time", value)
}

func (e *GstIpcPipelineSink) GetAckTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ack-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Continues a split pipeline from another process
type GstIpcPipelineSrc struct {
	*gst.Element
}

func NewIpcPipelineSrc() (*GstIpcPipelineSrc, error) {
	e, err := gst.NewElement("ipcpipelinesrc")
	if err != nil {
		return nil, err
	}
	return &GstIpcPipelineSrc{Element: e}, nil
}

func NewIpcPipelineSrcWithName(name string) (*GstIpcPipelineSrc, error) {
	e, err := gst.NewElementWithName("ipcpipelinesrc", name)
	if err != nil {
		return nil, err
	}
	return &GstIpcPipelineSrc{Element: e}, nil
}

// File descriptor to read data from
// Default: -1, Min: -1, Max: 65535
func (e *GstIpcPipelineSrc) SetFdin(value int) error {
	return e.Element.SetProperty("fdin", value)
}

func (e *GstIpcPipelineSrc) GetFdin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fdin")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// File descriptor to write data through
// Default: -1, Min: -1, Max: 65535
func (e *GstIpcPipelineSrc) SetFdout(value int) error {
	return e.Element.SetProperty("fdout", value)
}

func (e *GstIpcPipelineSrc) GetFdout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fdout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Read chunk size
// Default: 65536, Min: 1, Max: 16777216
func (e *GstIpcPipelineSrc) SetReadChunkSize(value uint) error {
	return e.Element.SetProperty("read-chunk-size", value)
}

func (e *GstIpcPipelineSrc) GetReadChunkSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("read-chunk-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum time to wait for a response to a message
// Default: 10000000, Min: 0, Max: 18446744073709551615
func (e *GstIpcPipelineSrc) SetAckTime(value uint64) error {
	return e.Element.SetProperty("ack-time", value)
}

func (e *GstIpcPipelineSrc) GetAckTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ack-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Contains the slave part of an inter-process pipeline
type GstIpcSlavePipeline struct {
	*gst.Pipeline
}

func NewIpcSlavePipeline() (*GstIpcSlavePipeline, error) {
	e, err := gst.NewElement("ipcslavepipeline")
	if err != nil {
		return nil, err
	}
	return &GstIpcSlavePipeline{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

func NewIpcSlavePipelineWithName(name string) (*GstIpcSlavePipeline, error) {
	e, err := gst.NewElementWithName("ipcslavepipeline", name)
	if err != nil {
		return nil, err
	}
	return &GstIpcSlavePipeline{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}
