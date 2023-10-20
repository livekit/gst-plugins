// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Free MPEG-2/4 AAC encoder
type GstFaac struct {
	*GstAudioEncoder
}

func NewFaac() (*GstFaac, error) {
	e, err := gst.NewElement("faac")
	if err != nil {
		return nil, err
	}
	return &GstFaac{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewFaacWithName(name string) (*GstFaac, error) {
	e, err := gst.NewElementWithName("faac", name)
	if err != nil {
		return nil, err
	}
	return &GstFaac{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Use temporal noise shaping
// Default: false
func (e *GstFaac) SetTns(value bool) error {
	return e.Element.SetProperty("tns", value)
}

func (e *GstFaac) GetTns() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("tns")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Average Bitrate (ABR) in bits/sec
// Default: 128000, Min: 8000, Max: 320000
func (e *GstFaac) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstFaac) GetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Allow mid/side encoding
// Default: true
func (e *GstFaac) SetMidside(value bool) error {
	return e.Element.SetProperty("midside", value)
}

func (e *GstFaac) GetMidside() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("midside")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Variable bitrate (VBR) quantizer quality in %%
// Default: 100, Min: 1, Max: 1000
func (e *GstFaac) SetQuality(value int) error {
	return e.Element.SetProperty("quality", value)
}

func (e *GstFaac) GetQuality() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Encoding bitrate type (VBR/ABR)
// Default: VBR encoding (1)
func (e *GstFaac) SetRateControl(value GstFaacBrtype) error {
	e.Element.SetArg("rate-control", string(value))
	return nil
}

func (e *GstFaac) GetRateControl() (GstFaacBrtype, error) {
	var value GstFaacBrtype
	var ok bool
	v, err := e.Element.GetProperty("rate-control")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFaacBrtype)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFaacBrtype")
	}
	return value, nil
}

// Block type encorcing
// Default: Normal block type (0)
func (e *GstFaac) SetShortctl(value GstFaacShortCtl) error {
	e.Element.SetArg("shortctl", string(value))
	return nil
}

func (e *GstFaac) GetShortctl() (GstFaacShortCtl, error) {
	var value GstFaacShortCtl
	var ok bool
	v, err := e.Element.GetProperty("shortctl")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFaacShortCtl)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFaacShortCtl")
	}
	return value, nil
}

type GstFaacBrtype string

const (
	GstFaacBrtypeVBREncoding GstFaacBrtype = "VBR encoding" // VBR
	GstFaacBrtypeABREncoding GstFaacBrtype = "ABR encoding" // ABR
)

type GstFaacShortCtl string

const (
	GstFaacShortCtlNormalBlockType GstFaacShortCtl = "Normal block type" // SHORTCTL_NORMAL
	GstFaacShortCtlNoShortBlocks   GstFaacShortCtl = "No short blocks"   // SHORTCTL_NOSHORT
	GstFaacShortCtlNoLongBlocks    GstFaacShortCtl = "No long blocks"    // SHORTCTL_NOLONG
)
