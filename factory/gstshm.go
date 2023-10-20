// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Send data over shared memory to the matching source
type GstShmSink struct {
	*base.GstBaseSink
}

func NewShmSink() (*GstShmSink, error) {
	e, err := gst.NewElement("shmsink")
	if err != nil {
		return nil, err
	}
	return &GstShmSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewShmSinkWithName(name string) (*GstShmSink, error) {
	e, err := gst.NewElementWithName("shmsink", name)
	if err != nil {
		return nil, err
	}
	return &GstShmSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Size of the shared memory area
// Default: 67108864, Min: 0, Max: -1
func (e *GstShmSink) SetShmSize(value uint) error {
	return e.Element.SetProperty("shm-size", value)
}

func (e *GstShmSink) GetShmSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("shm-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The path to the control socket used to control the shared memory transport. This may be modified during the NULL->READY transition
// Default: NULL
func (e *GstShmSink) SetSocketPath(value string) error {
	return e.Element.SetProperty("socket-path", value)
}

func (e *GstShmSink) GetSocketPath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("socket-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Block the stream until the shm pipe is connected
// Default: true
func (e *GstShmSink) SetWaitForConnection(value bool) error {
	return e.Element.SetProperty("wait-for-connection", value)
}

func (e *GstShmSink) GetWaitForConnection() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("wait-for-connection")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum Size of the shm buffer in nanoseconds (-1 to disable)
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstShmSink) SetBufferTime(value int64) error {
	return e.Element.SetProperty("buffer-time", value)
}

func (e *GstShmSink) GetBufferTime() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Permissions to set on the shm area
// Default: 416, Min: 0, Max: 4095
func (e *GstShmSink) SetPerms(value uint) error {
	return e.Element.SetProperty("perms", value)
}

func (e *GstShmSink) GetPerms() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("perms")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Receive data from the shared memory sink
type GstShmSrc struct {
	*base.GstPushSrc
}

func NewShmSrc() (*GstShmSrc, error) {
	e, err := gst.NewElement("shmsrc")
	if err != nil {
		return nil, err
	}
	return &GstShmSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewShmSrcWithName(name string) (*GstShmSrc, error) {
	e, err := gst.NewElementWithName("shmsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstShmSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// True if the element cannot produce data in PAUSED
// Default: false
func (e *GstShmSrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

func (e *GstShmSrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The name of the shared memory area used to get buffers
// Default: NULL
func (e *GstShmSrc) GetShmAreaName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("shm-area-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The path to the control socket used to control the shared memory
// Default: NULL
func (e *GstShmSrc) SetSocketPath(value string) error {
	return e.Element.SetProperty("socket-path", value)
}

func (e *GstShmSrc) GetSocketPath() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("socket-path")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}
