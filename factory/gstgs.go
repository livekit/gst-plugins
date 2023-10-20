// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Read from arbitrary point from a file in a Google Cloud Storage
type GstGsSrc struct {
	*base.GstBaseSrc
}

func NewGsSrc() (*GstGsSrc, error) {
	e, err := gst.NewElement("gssrc")
	if err != nil {
		return nil, err
	}
	return &GstGsSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewGsSrcWithName(name string) (*GstGsSrc, error) {
	e, err := gst.NewElementWithName("gssrc", name)
	if err != nil {
		return nil, err
	}
	return &GstGsSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Location of the file to read
// Default: NULL
func (e *GstGsSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstGsSrc) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Service Account Email to use for credentials
// Default: NULL
func (e *GstGsSrc) SetServiceAccountEmail(value string) error {
	return e.Element.SetProperty("service-account-email", value)
}

func (e *GstGsSrc) GetServiceAccountEmail() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("service-account-email")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Write buffers to a sequentially named set of files on Google Cloud Storage
type GstGsSink struct {
	*base.GstBaseSink
}

func NewGsSink() (*GstGsSink, error) {
	e, err := gst.NewElement("gssink")
	if err != nil {
		return nil, err
	}
	return &GstGsSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewGsSinkWithName(name string) (*GstGsSink, error) {
	e, err := gst.NewElementWithName("gssink", name)
	if err != nil {
		return nil, err
	}
	return &GstGsSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Google Cloud Storage Bucket Name
// Default: NULL
func (e *GstGsSink) SetBucketName(value string) error {
	return e.Element.SetProperty("bucket-name", value)
}

func (e *GstGsSink) GetBucketName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("bucket-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Index to use with location property to create file names.  The index is incremented by one for each buffer written.
// Default: 0, Min: 0, Max: 2147483647
func (e *GstGsSink) SetIndex(value int) error {
	return e.Element.SetProperty("index", value)
}

func (e *GstGsSink) GetIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// When to start a new file
// Default: buffer (0)
func (e *GstGsSink) SetNextFile(value GstGsSinkNext) error {
	e.Element.SetArg("next-file", string(value))
	return nil
}

func (e *GstGsSink) GetNextFile() (GstGsSinkNext, error) {
	var value GstGsSinkNext
	var ok bool
	v, err := e.Element.GetProperty("next-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGsSinkNext)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGsSinkNext")
	}
	return value, nil
}

// Full path name of the remote file
// Default: %%s_%%05d
func (e *GstGsSink) SetObjectName(value string) error {
	return e.Element.SetProperty("object-name", value)
}

func (e *GstGsSink) GetObjectName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("object-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Post a message for each file with information of the buffer
// Default: false
func (e *GstGsSink) SetPostMessages(value bool) error {
	return e.Element.SetProperty("post-messages", value)
}

func (e *GstGsSink) GetPostMessages() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-messages")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Service Account Email to use for credentials
// Default: NULL
func (e *GstGsSink) SetServiceAccountEmail(value string) error {
	return e.Element.SetProperty("service-account-email", value)
}

func (e *GstGsSink) GetServiceAccountEmail() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("service-account-email")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Start date in iso8601 format
// Default: NULL
func (e *GstGsSink) SetStartDate(value string) error {
	return e.Element.SetProperty("start-date", value)
}

func (e *GstGsSink) GetStartDate() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("start-date")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

type GstGsSinkNext string

const (
	GstGsSinkNextBuffer                GstGsSinkNext = "buffer"                      // New file for each buffer
	GstGsSinkNextOnlyOneFileNoNextFile GstGsSinkNext = "Only one file, no next file" // New file for each buffer
)
