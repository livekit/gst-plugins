// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Render overlay on a video stream using Cairo
type GstCairoOverlay struct {
	*base.GstBaseTransform
}

func NewCairoOverlay() (*GstCairoOverlay, error) {
	e, err := gst.NewElement("cairooverlay")
	if err != nil {
		return nil, err
	}
	return &GstCairoOverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCairoOverlayWithName(name string) (*GstCairoOverlay, error) {
	e, err := gst.NewElementWithName("cairooverlay", name)
	if err != nil {
		return nil, err
	}
	return &GstCairoOverlay{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Let the draw signal work on a transparent surface and blend the results with the video at a later time
// Default: false
func (e *GstCairoOverlay) SetDrawOnTransparentSurface(value bool) error {
	return e.Element.SetProperty("draw-on-transparent-surface", value)
}

func (e *GstCairoOverlay) GetDrawOnTransparentSurface() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("draw-on-transparent-surface")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
