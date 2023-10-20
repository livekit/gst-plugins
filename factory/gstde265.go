// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decodes HEVC/H.265 video streams using libde265
type GstLibde265Dec struct {
	*GstVideoDecoder
}

func NewLibde265Dec() (*GstLibde265Dec, error) {
	e, err := gst.NewElement("libde265dec")
	if err != nil {
		return nil, err
	}
	return &GstLibde265Dec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

func NewLibde265DecWithName(name string) (*GstLibde265Dec, error) {
	e, err := gst.NewElementWithName("libde265dec", name)
	if err != nil {
		return nil, err
	}
	return &GstLibde265Dec{GstVideoDecoder: &GstVideoDecoder{Element: e}}, nil
}

// Maximum number of worker threads to spawn. (0 = auto)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstLibde265Dec) SetMaxThreads(value int) error {
	return e.Element.SetProperty("max-threads", value)
}

func (e *GstLibde265Dec) GetMaxThreads() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-threads")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
