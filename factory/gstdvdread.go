// source: gst-plugins-ugly

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Access a DVD title/chapter/angle using libdvdread
type GstDvdReadSrc struct {
	*base.GstPushSrc
}

func NewDvdReadSrc() (*GstDvdReadSrc, error) {
	e, err := gst.NewElement("dvdreadsrc")
	if err != nil {
		return nil, err
	}
	return &GstDvdReadSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewDvdReadSrcWithName(name string) (*GstDvdReadSrc, error) {
	e, err := gst.NewElementWithName("dvdreadsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstDvdReadSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// angle
// Default: 1, Min: 1, Max: 999
func (e *GstDvdReadSrc) SetAngle(value int) error {
	return e.Element.SetProperty("angle", value)
}

func (e *GstDvdReadSrc) GetAngle() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("angle")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// chapter
// Default: 1, Min: 1, Max: 999
func (e *GstDvdReadSrc) SetChapter(value int) error {
	return e.Element.SetProperty("chapter", value)
}

func (e *GstDvdReadSrc) GetChapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("chapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// DVD device location
// Default: /dev/dvd
func (e *GstDvdReadSrc) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstDvdReadSrc) GetDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// title
// Default: 1, Min: 1, Max: 999
func (e *GstDvdReadSrc) SetTitle(value int) error {
	return e.Element.SetProperty("title", value)
}

func (e *GstDvdReadSrc) GetTitle() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("title")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
