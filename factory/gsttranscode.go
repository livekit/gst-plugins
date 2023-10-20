// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Autoplug and transcoder a stream
type GstTranscodeBin struct {
	*gst.Bin
}

func NewTranscodeBin() (*GstTranscodeBin, error) {
	e, err := gst.NewElement("transcodebin")
	if err != nil {
		return nil, err
	}
	return &GstTranscodeBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewTranscodeBinWithName(name string) (*GstTranscodeBin, error) {
	e, err := gst.NewElementWithName("transcodebin", name)
	if err != nil {
		return nil, err
	}
	return &GstTranscodeBin{Bin: &gst.Bin{Element: e}}, nil
}

// the audio filter(s) to apply, if possible

func (e *GstTranscodeBin) SetAudioFilter(value *gst.Element) error {
	return e.Element.SetProperty("audio-filter", value)
}

func (e *GstTranscodeBin) GetAudioFilter() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("audio-filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Whether to re-encode portions of compatible video streams that lay on segment boundaries
// Default: false
func (e *GstTranscodeBin) SetAvoidReencoding(value bool) error {
	return e.Element.SetProperty("avoid-reencoding", value)
}

func (e *GstTranscodeBin) GetAvoidReencoding() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("avoid-reencoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The GstEncodingProfile to use

func (e *GstTranscodeBin) SetProfile(value interface{}) error {
	return e.Element.SetProperty("profile", value)
}

func (e *GstTranscodeBin) GetProfile() (interface{}, error) {
	return e.Element.GetProperty("profile")
}

// the video filter(s) to apply, if possible

func (e *GstTranscodeBin) SetVideoFilter(value *gst.Element) error {
	return e.Element.SetProperty("video-filter", value)
}

func (e *GstTranscodeBin) GetVideoFilter() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("video-filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Autoplug and transcoder media from uris
type GstUriTranscodeBin struct {
	*gst.Pipeline
}

func NewUriTranscodeBin() (*GstUriTranscodeBin, error) {
	e, err := gst.NewElement("uritranscodebin")
	if err != nil {
		return nil, err
	}
	return &GstUriTranscodeBin{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

func NewUriTranscodeBinWithName(name string) (*GstUriTranscodeBin, error) {
	e, err := gst.NewElementWithName("uritranscodebin", name)
	if err != nil {
		return nil, err
	}
	return &GstUriTranscodeBin{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

// URI to decode
// Default: NULL
func (e *GstUriTranscodeBin) SetSourceUri(value string) error {
	return e.Element.SetProperty("source-uri", value)
}

func (e *GstUriTranscodeBin) GetSourceUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("source-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// the video filter(s) to apply, if possible

func (e *GstUriTranscodeBin) SetVideoFilter(value *gst.Element) error {
	return e.Element.SetProperty("video-filter", value)
}

func (e *GstUriTranscodeBin) GetVideoFilter() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("video-filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// the audio filter(s) to apply, if possible

func (e *GstUriTranscodeBin) SetAudioFilter(value *gst.Element) error {
	return e.Element.SetProperty("audio-filter", value)
}

func (e *GstUriTranscodeBin) GetAudioFilter() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("audio-filter")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Whether to re-encode portions of compatible video streams that lay on segment boundaries
// Default: false
func (e *GstUriTranscodeBin) SetAvoidReencoding(value bool) error {
	return e.Element.SetProperty("avoid-reencoding", value)
}

func (e *GstUriTranscodeBin) GetAvoidReencoding() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("avoid-reencoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The percentage of CPU to try to use with the processus running the pipeline driven by the clock
// Default: 100, Min: 0, Max: 100
func (e *GstUriTranscodeBin) SetCpuUsage(value uint) error {
	return e.Element.SetProperty("cpu-usage", value)
}

func (e *GstUriTranscodeBin) GetCpuUsage() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cpu-usage")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// URI to put output stream
// Default: NULL
func (e *GstUriTranscodeBin) SetDestUri(value string) error {
	return e.Element.SetProperty("dest-uri", value)
}

func (e *GstUriTranscodeBin) GetDestUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("dest-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The GstEncodingProfile to use

func (e *GstUriTranscodeBin) SetProfile(value interface{}) error {
	return e.Element.SetProperty("profile", value)
}

func (e *GstUriTranscodeBin) GetProfile() (interface{}, error) {
	return e.Element.GetProperty("profile")
}
