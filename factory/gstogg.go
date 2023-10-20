// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// mux ogg streams (info about ogg: http://xiph.org)
type GstOggMux struct {
	*gst.Element
}

func NewOggMux() (*GstOggMux, error) {
	e, err := gst.NewElement("oggmux")
	if err != nil {
		return nil, err
	}
	return &GstOggMux{Element: e}, nil
}

func NewOggMuxWithName(name string) (*GstOggMux, error) {
	e, err := gst.NewElementWithName("oggmux", name)
	if err != nil {
		return nil, err
	}
	return &GstOggMux{Element: e}, nil
}

// Maximum delay in multiplexing streams
// Default: 500000000, Min: 0, Max: 18446744073709551615
func (e *GstOggMux) SetMaxDelay(value uint64) error {
	return e.Element.SetProperty("max-delay", value)
}

func (e *GstOggMux) GetMaxDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Maximum delay for sending out a page
// Default: 500000000, Min: 0, Max: 18446744073709551615
func (e *GstOggMux) SetMaxPageDelay(value uint64) error {
	return e.Element.SetProperty("max-page-delay", value)
}

func (e *GstOggMux) GetMaxPageDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-page-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Maximum timestamp difference for maintaining perfect granules
// Default: 40000000, Min: 0, Max: 18446744073709551615
func (e *GstOggMux) SetMaxTolerance(value uint64) error {
	return e.Element.SetProperty("max-tolerance", value)
}

func (e *GstOggMux) GetMaxTolerance() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Whether to include a Skeleton track
// Default: false
func (e *GstOggMux) SetSkeleton(value bool) error {
	return e.Element.SetProperty("skeleton", value)
}

func (e *GstOggMux) GetSkeleton() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("skeleton")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// parse ogg streams into pages (info about ogg: http://xiph.org)
type GstOggParse struct {
	*gst.Element
}

func NewOggParse() (*GstOggParse, error) {
	e, err := gst.NewElement("oggparse")
	if err != nil {
		return nil, err
	}
	return &GstOggParse{Element: e}, nil
}

func NewOggParseWithName(name string) (*GstOggParse, error) {
	e, err := gst.NewElementWithName("oggparse", name)
	if err != nil {
		return nil, err
	}
	return &GstOggParse{Element: e}, nil
}

// parse an OGM audio header and stream
type GstOgmAudioParse struct {
	*GstOgmParse
}

func NewOgmAudioParse() (*GstOgmAudioParse, error) {
	e, err := gst.NewElement("ogmaudioparse")
	if err != nil {
		return nil, err
	}
	return &GstOgmAudioParse{GstOgmParse: &GstOgmParse{Element: e}}, nil
}

func NewOgmAudioParseWithName(name string) (*GstOgmAudioParse, error) {
	e, err := gst.NewElementWithName("ogmaudioparse", name)
	if err != nil {
		return nil, err
	}
	return &GstOgmAudioParse{GstOgmParse: &GstOgmParse{Element: e}}, nil
}

// parse an OGM text header and stream
type GstOgmTextParse struct {
	*GstOgmParse
}

func NewOgmTextParse() (*GstOgmTextParse, error) {
	e, err := gst.NewElement("ogmtextparse")
	if err != nil {
		return nil, err
	}
	return &GstOgmTextParse{GstOgmParse: &GstOgmParse{Element: e}}, nil
}

func NewOgmTextParseWithName(name string) (*GstOgmTextParse, error) {
	e, err := gst.NewElementWithName("ogmtextparse", name)
	if err != nil {
		return nil, err
	}
	return &GstOgmTextParse{GstOgmParse: &GstOgmParse{Element: e}}, nil
}

// parse an OGM video header and stream
type GstOgmVideoParse struct {
	*GstOgmParse
}

func NewOgmVideoParse() (*GstOgmVideoParse, error) {
	e, err := gst.NewElement("ogmvideoparse")
	if err != nil {
		return nil, err
	}
	return &GstOgmVideoParse{GstOgmParse: &GstOgmParse{Element: e}}, nil
}

func NewOgmVideoParseWithName(name string) (*GstOgmVideoParse, error) {
	e, err := gst.NewElementWithName("ogmvideoparse", name)
	if err != nil {
		return nil, err
	}
	return &GstOgmVideoParse{GstOgmParse: &GstOgmParse{Element: e}}, nil
}

// parse an ogg avi stream into pages (info about ogg: http://xiph.org)
type GstOggAviParse struct {
	*gst.Element
}

func NewOggAviParse() (*GstOggAviParse, error) {
	e, err := gst.NewElement("oggaviparse")
	if err != nil {
		return nil, err
	}
	return &GstOggAviParse{Element: e}, nil
}

func NewOggAviParseWithName(name string) (*GstOggAviParse, error) {
	e, err := gst.NewElementWithName("oggaviparse", name)
	if err != nil {
		return nil, err
	}
	return &GstOggAviParse{Element: e}, nil
}

// demux ogg streams (info about ogg: http://xiph.org)
type GstOggDemux struct {
	*gst.Element
}

func NewOggDemux() (*GstOggDemux, error) {
	e, err := gst.NewElement("oggdemux")
	if err != nil {
		return nil, err
	}
	return &GstOggDemux{Element: e}, nil
}

func NewOggDemuxWithName(name string) (*GstOggDemux, error) {
	e, err := gst.NewElementWithName("oggdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstOggDemux{Element: e}, nil
}

type GstOgmParse struct {
	*gst.Element
}
