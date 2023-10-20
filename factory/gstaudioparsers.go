// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Advanced Audio Coding parser
type GstAacParse struct {
	*GstBaseParse
}

func NewAacParse() (*GstAacParse, error) {
	e, err := gst.NewElement("aacparse")
	if err != nil {
		return nil, err
	}
	return &GstAacParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewAacParseWithName(name string) (*GstAacParse, error) {
	e, err := gst.NewElementWithName("aacparse", name)
	if err != nil {
		return nil, err
	}
	return &GstAacParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// AC3 parser
type GstAc3Parse struct {
	*GstBaseParse
}

func NewAc3Parse() (*GstAc3Parse, error) {
	e, err := gst.NewElement("ac3parse")
	if err != nil {
		return nil, err
	}
	return &GstAc3Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewAc3ParseWithName(name string) (*GstAc3Parse, error) {
	e, err := gst.NewElementWithName("ac3parse", name)
	if err != nil {
		return nil, err
	}
	return &GstAc3Parse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Adaptive Multi-Rate audio parser
type GstAmrParse struct {
	*GstBaseParse
}

func NewAmrParse() (*GstAmrParse, error) {
	e, err := gst.NewElement("amrparse")
	if err != nil {
		return nil, err
	}
	return &GstAmrParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewAmrParseWithName(name string) (*GstAmrParse, error) {
	e, err := gst.NewElementWithName("amrparse", name)
	if err != nil {
		return nil, err
	}
	return &GstAmrParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// DCA parser
type GstDcaParse struct {
	*GstBaseParse
}

func NewDcaParse() (*GstDcaParse, error) {
	e, err := gst.NewElement("dcaparse")
	if err != nil {
		return nil, err
	}
	return &GstDcaParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewDcaParseWithName(name string) (*GstDcaParse, error) {
	e, err := gst.NewElementWithName("dcaparse", name)
	if err != nil {
		return nil, err
	}
	return &GstDcaParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Parses audio with the FLAC lossless audio codec
type GstFlacParse struct {
	*GstBaseParse
}

func NewFlacParse() (*GstFlacParse, error) {
	e, err := gst.NewElement("flacparse")
	if err != nil {
		return nil, err
	}
	return &GstFlacParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewFlacParseWithName(name string) (*GstFlacParse, error) {
	e, err := gst.NewElementWithName("flacparse", name)
	if err != nil {
		return nil, err
	}
	return &GstFlacParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Check the overall checksums of every frame
// Default: false
func (e *GstFlacParse) SetCheckFrameChecksums(value bool) error {
	return e.Element.SetProperty("check-frame-checksums", value)
}

func (e *GstFlacParse) GetCheckFrameChecksums() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("check-frame-checksums")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Parses and frames mpeg1 audio streams (levels 1-3), provides seek
type GstMpegAudioParse struct {
	*GstBaseParse
}

func NewMpegAudioParse() (*GstMpegAudioParse, error) {
	e, err := gst.NewElement("mpegaudioparse")
	if err != nil {
		return nil, err
	}
	return &GstMpegAudioParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewMpegAudioParseWithName(name string) (*GstMpegAudioParse, error) {
	e, err := gst.NewElementWithName("mpegaudioparse", name)
	if err != nil {
		return nil, err
	}
	return &GstMpegAudioParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Parses an SBC bluetooth audio stream
type GstSbcParse struct {
	*GstBaseParse
}

func NewSbcParse() (*GstSbcParse, error) {
	e, err := gst.NewElement("sbcparse")
	if err != nil {
		return nil, err
	}
	return &GstSbcParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewSbcParseWithName(name string) (*GstSbcParse, error) {
	e, err := gst.NewElementWithName("sbcparse", name)
	if err != nil {
		return nil, err
	}
	return &GstSbcParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Wavpack parser
type GstWavpackParse struct {
	*GstBaseParse
}

func NewWavpackParse() (*GstWavpackParse, error) {
	e, err := gst.NewElement("wavpackparse")
	if err != nil {
		return nil, err
	}
	return &GstWavpackParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewWavpackParseWithName(name string) (*GstWavpackParse, error) {
	e, err := gst.NewElementWithName("wavpackparse", name)
	if err != nil {
		return nil, err
	}
	return &GstWavpackParse{GstBaseParse: &GstBaseParse{Element: e}}, nil
}
