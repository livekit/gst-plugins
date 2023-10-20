// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Computes an AccurateRip CRC
type GstAccurip struct {
	*GstAudioFilter
}

func NewAccurip() (*GstAccurip, error) {
	e, err := gst.NewElement("accurip")
	if err != nil {
		return nil, err
	}
	return &GstAccurip{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAccuripWithName(name string) (*GstAccurip, error) {
	e, err := gst.NewElementWithName("accurip", name)
	if err != nil {
		return nil, err
	}
	return &GstAccurip{GstAudioFilter: &GstAudioFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Indicate to the CRC calculation algorithm that this is the first track of a set
// Default: false
func (e *GstAccurip) SetFirstTrack(value bool) error {
	return e.Element.SetProperty("first-track", value)
}

func (e *GstAccurip) GetFirstTrack() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("first-track")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Indicate to the CRC calculation algorithm that this is the last track of a set
// Default: false
func (e *GstAccurip) SetLastTrack(value bool) error {
	return e.Element.SetProperty("last-track", value)
}

func (e *GstAccurip) GetLastTrack() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("last-track")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
