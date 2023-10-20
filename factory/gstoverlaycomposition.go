// source: gst-plugins-base

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Overlay Composition
type GstOverlayComposition struct {
	*gst.Element
}

func NewOverlayComposition() (*GstOverlayComposition, error) {
	e, err := gst.NewElement("overlaycomposition")
	if err != nil {
		return nil, err
	}
	return &GstOverlayComposition{Element: e}, nil
}

func NewOverlayCompositionWithName(name string) (*GstOverlayComposition, error) {
	e, err := gst.NewElementWithName("overlaycomposition", name)
	if err != nil {
		return nil, err
	}
	return &GstOverlayComposition{Element: e}, nil
}
