// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decode a raw VBI stream containing teletext information to RGBA and text
type GstTeletextDec struct {
	*gst.Element
}

func NewTeletextDec() (*GstTeletextDec, error) {
	e, err := gst.NewElement("teletextdec")
	if err != nil {
		return nil, err
	}
	return &GstTeletextDec{Element: e}, nil
}

func NewTeletextDecWithName(name string) (*GstTeletextDec, error) {
	e, err := gst.NewElementWithName("teletextdec", name)
	if err != nil {
		return nil, err
	}
	return &GstTeletextDec{Element: e}, nil
}

// Number of sub-page that should displayed (-1 for all)
// Default: -1, Min: -1, Max: 153
func (e *GstTeletextDec) SetSubpage(value int) error {
	return e.Element.SetProperty("subpage", value)
}

func (e *GstTeletextDec) GetSubpage() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("subpage")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Enables subtitles mode for text output stripping the blank lines and the teletext state lines
// Default: false
func (e *GstTeletextDec) SetSubtitlesMode(value bool) error {
	return e.Element.SetProperty("subtitles-mode", value)
}

func (e *GstTeletextDec) GetSubtitlesMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("subtitles-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Output template used to print each one of the subtitles lines
// Default: %%s\n
func (e *GstTeletextDec) SetSubtitlesTemplate(value string) error {
	return e.Element.SetProperty("subtitles-template", value)
}

func (e *GstTeletextDec) GetSubtitlesTemplate() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitles-template")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Font description used for the pango output.
// Default: verdana 12
func (e *GstTeletextDec) SetFontDescription(value string) error {
	return e.Element.SetProperty("font-description", value)
}

func (e *GstTeletextDec) GetFontDescription() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("font-description")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Number of page that should displayed
// Default: 100, Min: 100, Max: 999
func (e *GstTeletextDec) SetPage(value int) error {
	return e.Element.SetProperty("page", value)
}

func (e *GstTeletextDec) GetPage() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("page")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
