// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Uses libdv to decode DV video (smpte314) (libdv.sourceforge.net)
type GstDVDec struct {
	*gst.Element
}

func NewDVDec() (*GstDVDec, error) {
	e, err := gst.NewElement("dvdec")
	if err != nil {
		return nil, err
	}
	return &GstDVDec{Element: e}, nil
}

func NewDVDecWithName(name string) (*GstDVDec, error) {
	e, err := gst.NewElementWithName("dvdec", name)
	if err != nil {
		return nil, err
	}
	return &GstDVDec{Element: e}, nil
}

// Decoding quality
// Default: best (5)
func (e *GstDVDec) SetQuality(value GstDVDecQualityEnum) error {
	e.Element.SetArg("quality", string(value))
	return nil
}

func (e *GstDVDec) GetQuality() (GstDVDecQualityEnum, error) {
	var value GstDVDecQualityEnum
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDVDecQualityEnum)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDVDecQualityEnum")
	}
	return value, nil
}

// Clamp chroma
// Default: false
func (e *GstDVDec) SetClampChroma(value bool) error {
	return e.Element.SetProperty("clamp-chroma", value)
}

func (e *GstDVDec) GetClampChroma() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("clamp-chroma")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Clamp luma
// Default: false
func (e *GstDVDec) SetClampLuma(value bool) error {
	return e.Element.SetProperty("clamp-luma", value)
}

func (e *GstDVDec) GetClampLuma() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("clamp-luma")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Only decode Nth frame
// Default: 1, Min: 1, Max: 2147483647
func (e *GstDVDec) SetDropFactor(value int) error {
	return e.Element.SetProperty("drop-factor", value)
}

func (e *GstDVDec) GetDropFactor() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("drop-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Uses libdv to separate DV audio from DV video (libdv.sourceforge.net)
type GstDVDemux struct {
	*gst.Element
}

func NewDVDemux() (*GstDVDemux, error) {
	e, err := gst.NewElement("dvdemux")
	if err != nil {
		return nil, err
	}
	return &GstDVDemux{Element: e}, nil
}

func NewDVDemuxWithName(name string) (*GstDVDemux, error) {
	e, err := gst.NewElementWithName("dvdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstDVDemux{Element: e}, nil
}

type GstDVDecQualityEnum string

const (
	GstDVDecQualityEnumFastest        GstDVDecQualityEnum = "fastest"         // Monochrome, DC (Fastest)
	GstDVDecQualityEnumMonochromeAc   GstDVDecQualityEnum = "monochrome-ac"   // Monochrome, first AC coefficient
	GstDVDecQualityEnumMonochromeBest GstDVDecQualityEnum = "monochrome-best" // Monochrome, highest quality
	GstDVDecQualityEnumColourFastest  GstDVDecQualityEnum = "colour-fastest"  // Colour, DC, fastest
	GstDVDecQualityEnumColourAc       GstDVDecQualityEnum = "colour-ac"       // Colour, using only the first AC coefficient
	GstDVDecQualityEnumBest           GstDVDecQualityEnum = "best"            // Highest quality colour decoding
)
