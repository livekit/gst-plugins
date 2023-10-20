// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Encode ADPCM audio
type ADPCMEnc struct {
	*GstAudioEncoder
}

func NewADPCMEnc() (*ADPCMEnc, error) {
	e, err := gst.NewElement("adpcmenc")
	if err != nil {
		return nil, err
	}
	return &ADPCMEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewADPCMEncWithName(name string) (*ADPCMEnc, error) {
	e, err := gst.NewElementWithName("adpcmenc", name)
	if err != nil {
		return nil, err
	}
	return &ADPCMEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Block size for output stream
// Default: 1024, Min: 64, Max: 8192
func (e *ADPCMEnc) SetBlockalign(value int) error {
	return e.Element.SetProperty("blockalign", value)
}

func (e *ADPCMEnc) GetBlockalign() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("blockalign")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Layout for output stream
// Default: dvi (0)
func (e *ADPCMEnc) SetLayout(value GstADPCMEncLayout) error {
	e.Element.SetArg("layout", string(value))
	return nil
}

func (e *ADPCMEnc) GetLayout() (GstADPCMEncLayout, error) {
	var value GstADPCMEncLayout
	var ok bool
	v, err := e.Element.GetProperty("layout")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstADPCMEncLayout)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstADPCMEncLayout")
	}
	return value, nil
}

type GstADPCMEncLayout string

const (
	GstADPCMEncLayoutDvi GstADPCMEncLayout = "dvi" // DVI/IMA APDCM
)
