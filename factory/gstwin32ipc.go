// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Receive video frames from the win32ipcvideosink
type GstWin32IpcVideoSrc struct {
	*base.GstBaseSrc
}

func NewWin32IpcVideoSrc() (*GstWin32IpcVideoSrc, error) {
	e, err := gst.NewElement("win32ipcvideosrc")
	if err != nil {
		return nil, err
	}
	return &GstWin32IpcVideoSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewWin32IpcVideoSrcWithName(name string) (*GstWin32IpcVideoSrc, error) {
	e, err := gst.NewElementWithName("win32ipcvideosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstWin32IpcVideoSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// The name of Win32 named pipe to communicate with server. Validation of the pipe name is caller's responsibility
// Default: \\.\pipe\gst.win32.ipc.video
func (e *GstWin32IpcVideoSrc) SetPipeName(value string) error {
	return e.Element.SetProperty("pipe-name", value)
}

func (e *GstWin32IpcVideoSrc) GetPipeName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pipe-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Maximum processing time for a buffer in nanoseconds
// Default: 20000000, Min: 0, Max: 18446744073709551615
func (e *GstWin32IpcVideoSrc) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

func (e *GstWin32IpcVideoSrc) GetProcessingDeadline() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("processing-deadline")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Send video frames to win32ipcvideosrc elements
type GstWin32IpcVideoSink struct {
	*base.GstBaseSink
}

func NewWin32IpcVideoSink() (*GstWin32IpcVideoSink, error) {
	e, err := gst.NewElement("win32ipcvideosink")
	if err != nil {
		return nil, err
	}
	return &GstWin32IpcVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewWin32IpcVideoSinkWithName(name string) (*GstWin32IpcVideoSink, error) {
	e, err := gst.NewElementWithName("win32ipcvideosink", name)
	if err != nil {
		return nil, err
	}
	return &GstWin32IpcVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// The name of Win32 named pipe to communicate with clients. Validation of the pipe name is caller's responsibility
// Default: \\.\pipe\gst.win32.ipc.video
func (e *GstWin32IpcVideoSink) SetPipeName(value string) error {
	return e.Element.SetProperty("pipe-name", value)
}

func (e *GstWin32IpcVideoSink) GetPipeName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pipe-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}
