// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// DVD playback element
type RsnDvdBin struct {
	*gst.Bin
}

func NewRsnDvdBin() (*RsnDvdBin, error) {
	e, err := gst.NewElement("rsndvdbin")
	if err != nil {
		return nil, err
	}
	return &RsnDvdBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewRsnDvdBinWithName(name string) (*RsnDvdBin, error) {
	e, err := gst.NewElementWithName("rsndvdbin", name)
	if err != nil {
		return nil, err
	}
	return &RsnDvdBin{Bin: &gst.Bin{Element: e}}, nil
}

// DVD device location
// Default: /dev/dvd
func (e *RsnDvdBin) SetDevice(value string) error {
	return e.Element.SetProperty("device", value)
}

func (e *RsnDvdBin) GetDevice() (string, error) {
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
