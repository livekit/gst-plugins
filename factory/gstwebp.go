// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decode images from WebP format
type GstWebPDec struct {
	*GstVideoDecoder
}

func NewWebPDec() (*GstWebPDec, error) {
	e, err := gst.NewElement("webpdec")
	if err != nil {
		return nil, err
	}
	return &GstWebPDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewWebPDecWithName(name string) (*GstWebPDec, error) {
	e, err := gst.NewElementWithName("webpdec", name)
	if err != nil {
		return nil, err
	}
	return &GstWebPDec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

// When enabled, use multi-threaded decoding
// Default: false
func (e *GstWebPDec) SetUseThreads(value bool) error {
	return e.Element.SetProperty("use-threads", value)
}

func (e *GstWebPDec) GetUseThreads() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled, skip the in-loop filtering
// Default: false
func (e *GstWebPDec) SetBypassFiltering(value bool) error {
	return e.Element.SetProperty("bypass-filtering", value)
}

func (e *GstWebPDec) GetBypassFiltering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("bypass-filtering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// When enabled, use faster pointwise upsampler
// Default: false
func (e *GstWebPDec) SetNoFancyUpsampling(value bool) error {
	return e.Element.SetProperty("no-fancy-upsampling", value)
}

func (e *GstWebPDec) GetNoFancyUpsampling() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("no-fancy-upsampling")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Encode images in WEBP format
type GstWebpEnc struct {
	*GstVideoEncoder
}

func NewWebpEnc() (*GstWebpEnc, error) {
	e, err := gst.NewElement("webpenc")
	if err != nil {
		return nil, err
	}
	return &GstWebpEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

func NewWebpEncWithName(name string) (*GstWebpEnc, error) {
	e, err := gst.NewElementWithName("webpenc", name)
	if err != nil {
		return nil, err
	}
	return &GstWebpEnc{GstVideoEncoder: &GstVideoEncoder{Element: e}}, nil
}

// quality level, between 0 (smallest file) and 100 (biggest)
// Default: 90, Min: 0, Max: 100
func (e *GstWebpEnc) SetQuality(value float32) error {
	return e.Element.SetProperty("quality", value)
}

func (e *GstWebpEnc) GetQuality() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// quality/speed trade-off (0=fast, 6=slower-better)
// Default: 4, Min: 0, Max: 6
func (e *GstWebpEnc) SetSpeed(value uint) error {
	return e.Element.SetProperty("speed", value)
}

func (e *GstWebpEnc) GetSpeed() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Enable lossless encoding
// Default: false
func (e *GstWebpEnc) SetLossless(value bool) error {
	return e.Element.SetProperty("lossless", value)
}

func (e *GstWebpEnc) GetLossless() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lossless")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Preset name for visual tuning
// Default: photo (2)
func (e *GstWebpEnc) SetPreset(value GstWebpEncPreset) error {
	e.Element.SetArg("preset", string(value))
	return nil
}

func (e *GstWebpEnc) GetPreset() (GstWebpEncPreset, error) {
	var value GstWebpEncPreset
	var ok bool
	v, err := e.Element.GetProperty("preset")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstWebpEncPreset)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstWebpEncPreset")
	}
	return value, nil
}

type GstWebpEncPreset string

const (
	GstWebpEncPresetNone    GstWebpEncPreset = "none"    // Default
	GstWebpEncPresetPicture GstWebpEncPreset = "picture" // Digital picture,inner shot
	GstWebpEncPresetPhoto   GstWebpEncPreset = "photo"   // Outdoor photo, natural lighting
	GstWebpEncPresetDrawing GstWebpEncPreset = "drawing" // Hand or Line drawing
	GstWebpEncPresetIcon    GstWebpEncPreset = "icon"    // Small-sized colorful images
	GstWebpEncPresetText    GstWebpEncPreset = "text"    // text-like
)
