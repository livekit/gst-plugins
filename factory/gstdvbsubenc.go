// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Encodes AYUV video frames streams into DVB subtitles
type GstDvbSubEnc struct {
	*gst.Element
}

func NewDvbSubEnc() (*GstDvbSubEnc, error) {
	e, err := gst.NewElement("dvbsubenc")
	if err != nil {
		return nil, err
	}
	return &GstDvbSubEnc{Element: e}, nil
}

func NewDvbSubEncWithName(name string) (*GstDvbSubEnc, error) {
	e, err := gst.NewElementWithName("dvbsubenc", name)
	if err != nil {
		return nil, err
	}
	return &GstDvbSubEnc{Element: e}, nil
}

// Maximum Number of Colours to output
// Default: 16, Min: 1, Max: 256
func (e *GstDvbSubEnc) SetMaxColours(value int) error {
	return e.Element.SetProperty("max-colours", value)
}

func (e *GstDvbSubEnc) GetMaxColours() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-colours")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Apply an offset to incoming timestamps before output (in nanoseconds)
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstDvbSubEnc) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstDvbSubEnc) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}
