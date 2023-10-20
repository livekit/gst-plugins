// source: gst-plugins-ugly

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decodes ATSC A/52 encoded audio streams
type GstA52Dec struct {
	*GstAudioDecoder
}

func NewA52Dec() (*GstA52Dec, error) {
	e, err := gst.NewElement("a52dec")
	if err != nil {
		return nil, err
	}
	return &GstA52Dec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewA52DecWithName(name string) (*GstA52Dec, error) {
	e, err := gst.NewElementWithName("a52dec", name)
	if err != nil {
		return nil, err
	}
	return &GstA52Dec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Use Dynamic Range Compression
// Default: false
func (e *GstA52Dec) SetDrc(value bool) error {
	return e.Element.SetProperty("drc", value)
}

func (e *GstA52Dec) GetDrc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// LFE
// Default: false
func (e *GstA52Dec) SetLfe(value bool) error {
	return e.Element.SetProperty("lfe", value)
}

func (e *GstA52Dec) GetLfe() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("lfe")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Decoding Mode (default 3f2r)
// Default:  (0)
func (e *GstA52Dec) SetMode(value GstA52DecMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstA52Dec) GetMode() (GstA52DecMode, error) {
	var value GstA52DecMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstA52DecMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstA52DecMode")
	}
	return value, nil
}

type GstA52DecMode string

const (
	GstA52DecModeMono   GstA52DecMode = "mono"   // Mono
	GstA52DecModeStereo GstA52DecMode = "stereo" // Stereo
	GstA52DecMode3f     GstA52DecMode = "3f"     // 3 Front
	GstA52DecMode2f1r   GstA52DecMode = "2f1r"   // 2 Front, 1 Rear
	GstA52DecMode3f1r   GstA52DecMode = "3f1r"   // 3 Front, 1 Rear
	GstA52DecMode2f2r   GstA52DecMode = "2f2r"   // 2 Front, 2 Rear
	GstA52DecMode3f2r   GstA52DecMode = "3f2r"   // 3 Front, 2 Rear
	GstA52DecModeDolby  GstA52DecMode = "dolby"  // Dolby
)
