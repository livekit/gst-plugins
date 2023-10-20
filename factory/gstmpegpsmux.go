// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Multiplexes media streams into an MPEG Program Stream
type MpegPsMux struct {
	*gst.Element
}

func NewMpegPsMux() (*MpegPsMux, error) {
	e, err := gst.NewElement("mpegpsmux")
	if err != nil {
		return nil, err
	}
	return &MpegPsMux{Element: e}, nil
}

func NewMpegPsMuxWithName(name string) (*MpegPsMux, error) {
	e, err := gst.NewElementWithName("mpegpsmux", name)
	if err != nil {
		return nil, err
	}
	return &MpegPsMux{Element: e}, nil
}

// Whether to aggregate GOPs and push them out as buffer lists
// Default: false
func (e *MpegPsMux) SetAggregateGops(value bool) error {
	return e.Element.SetProperty("aggregate-gops", value)
}

func (e *MpegPsMux) GetAggregateGops() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("aggregate-gops")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
