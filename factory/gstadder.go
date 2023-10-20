// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Add N audio channels together
type GstAdder struct {
	*gst.Element
}

func NewAdder() (*GstAdder, error) {
	e, err := gst.NewElement("adder")
	if err != nil {
		return nil, err
	}
	return &GstAdder{Element: e}, nil
}

func NewAdderWithName(name string) (*GstAdder, error) {
	e, err := gst.NewElementWithName("adder", name)
	if err != nil {
		return nil, err
	}
	return &GstAdder{Element: e}, nil
}

// Set target format for mixing (NULL means ANY). Setting this property takes a reference to the supplied GstCaps object.

func (e *GstAdder) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstAdder) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}
