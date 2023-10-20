// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decode images from JPEG format
type GstJpegDec struct {
	*GstVideoDecoder
}

func NewJpegDec() (*GstJpegDec, error) {
	e, err := gst.NewElement("jpegdec")
	if err != nil {
		return nil, err
	}
	return &GstJpegDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewJpegDecWithName(name string) (*GstJpegDec, error) {
	e, err := gst.NewElementWithName("jpegdec", name)
	if err != nil {
		return nil, err
	}
	return &GstJpegDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

// The IDCT algorithm to use
// Default: ifast (1)
func (e *GstJpegDec) SetIdctMethod(value GstIDCTMethod) error {
	e.Element.SetArg("idct-method", string(value))
	return nil
}

func (e *GstJpegDec) GetIdctMethod() (GstIDCTMethod, error) {
	var value GstIDCTMethod
	var ok bool
	v, err := e.Element.GetProperty("idct-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstIDCTMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstIDCTMethod")
	}
	return value, nil
}

// (Deprecated) Error out after receiving N consecutive decoding errors (-1 = never fail, 0 = automatic, 1 = fail on first error)
// Default: 0, Min: -1, Max: 2147483647
func (e *GstJpegDec) SetMaxErrors(value int) error {
	return e.Element.SetProperty("max-errors", value)
}

func (e *GstJpegDec) GetMaxErrors() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-errors")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Encode images in JPEG format
type GstJpegEnc struct {
	*GstVideoEncoder
}

func NewJpegEnc() (*GstJpegEnc, error) {
	e, err := gst.NewElement("jpegenc")
	if err != nil {
		return nil, err
	}
	return &GstJpegEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewJpegEncWithName(name string) (*GstJpegEnc, error) {
	e, err := gst.NewElementWithName("jpegenc", name)
	if err != nil {
		return nil, err
	}
	return &GstJpegEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// The IDCT algorithm to use
// Default: ifast (1)
func (e *GstJpegEnc) SetIdctMethod(value GstIDCTMethod) error {
	e.Element.SetArg("idct-method", string(value))
	return nil
}

func (e *GstJpegEnc) GetIdctMethod() (GstIDCTMethod, error) {
	var value GstIDCTMethod
	var ok bool
	v, err := e.Element.GetProperty("idct-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstIDCTMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstIDCTMethod")
	}
	return value, nil
}

// Quality of encoding
// Default: 85, Min: 0, Max: 100
func (e *GstJpegEnc) SetQuality(value int) error {
	return e.Element.SetProperty("quality", value)
}

func (e *GstJpegEnc) GetQuality() (int, error) {
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

// Send EOS after encoding a frame, useful for snapshots
// Default: false
func (e *GstJpegEnc) SetSnapshot(value bool) error {
	return e.Element.SetProperty("snapshot", value)
}

func (e *GstJpegEnc) GetSnapshot() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("snapshot")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstIDCTMethod string

const (
	GstIDCTMethodIslow GstIDCTMethod = "islow" // Slow but accurate integer algorithm
	GstIDCTMethodIfast GstIDCTMethod = "ifast" // Faster, less accurate integer method
	GstIDCTMethodFloat GstIDCTMethod = "float" // Floating-point: accurate, fast on fast HW
)
