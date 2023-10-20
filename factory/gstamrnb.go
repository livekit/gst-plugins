// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Adaptive Multi-Rate Narrow-Band audio decoder
type GstAmrnbDec struct {
	*GstAudioDecoder
}

func NewAmrnbDec() (*GstAmrnbDec, error) {
	e, err := gst.NewElement("amrnbdec")
	if err != nil {
		return nil, err
	}
	return &GstAmrnbDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewAmrnbDecWithName(name string) (*GstAmrnbDec, error) {
	e, err := gst.NewElementWithName("amrnbdec", name)
	if err != nil {
		return nil, err
	}
	return &GstAmrnbDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// The decoder variant
// Default: IF1 (0)
func (e *GstAmrnbDec) SetVariant(value GstAmrnbVariant) error {
	e.Element.SetArg("variant", string(value))
	return nil
}

func (e *GstAmrnbDec) GetVariant() (GstAmrnbVariant, error) {
	var value GstAmrnbVariant
	var ok bool
	v, err := e.Element.GetProperty("variant")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmrnbVariant)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmrnbVariant")
	}
	return value, nil
}

// Adaptive Multi-Rate Narrow-Band audio encoder
type GstAmrnbEnc struct {
	*GstAudioEncoder
}

func NewAmrnbEnc() (*GstAmrnbEnc, error) {
	e, err := gst.NewElement("amrnbenc")
	if err != nil {
		return nil, err
	}
	return &GstAmrnbEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewAmrnbEncWithName(name string) (*GstAmrnbEnc, error) {
	e, err := gst.NewElementWithName("amrnbenc", name)
	if err != nil {
		return nil, err
	}
	return &GstAmrnbEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Encoding Band Mode (Kbps)
// Default: MR122 (7)
func (e *GstAmrnbEnc) SetBandMode(value GstAmrnbEncBandMode) error {
	e.Element.SetArg("band-mode", string(value))
	return nil
}

func (e *GstAmrnbEnc) GetBandMode() (GstAmrnbEncBandMode, error) {
	var value GstAmrnbEncBandMode
	var ok bool
	v, err := e.Element.GetProperty("band-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAmrnbEncBandMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAmrnbEncBandMode")
	}
	return value, nil
}

type GstAmrnbEncBandMode string

const (
	GstAmrnbEncBandModeMR475 GstAmrnbEncBandMode = "MR475" // MR475
	GstAmrnbEncBandModeMR515 GstAmrnbEncBandMode = "MR515" // MR515
	GstAmrnbEncBandModeMR59  GstAmrnbEncBandMode = "MR59"  // MR59
	GstAmrnbEncBandModeMR67  GstAmrnbEncBandMode = "MR67"  // MR67
	GstAmrnbEncBandModeMR74  GstAmrnbEncBandMode = "MR74"  // MR74
	GstAmrnbEncBandModeMR795 GstAmrnbEncBandMode = "MR795" // MR795
	GstAmrnbEncBandModeMR102 GstAmrnbEncBandMode = "MR102" // MR102
	GstAmrnbEncBandModeMR122 GstAmrnbEncBandMode = "MR122" // MR122
	GstAmrnbEncBandModeMRDTX GstAmrnbEncBandMode = "MRDTX" // MRDTX
)

type GstAmrnbVariant string

const (
	GstAmrnbVariantIF1 GstAmrnbVariant = "IF1" // IF1
	GstAmrnbVariantIF2 GstAmrnbVariant = "IF2" // IF2
)
