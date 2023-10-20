// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Parse and demultiplex a Smooth Streaming manifest into audio and video streams
type GstMssDemux struct {
	*GstAdaptiveDemux
}

func NewMssDemux() (*GstMssDemux, error) {
	e, err := gst.NewElement("mssdemux")
	if err != nil {
		return nil, err
	}
	return &GstMssDemux{GstAdaptiveDemux: &GstAdaptiveDemux{Bin: &gst.Bin{Element: e}}}, nil
}

func NewMssDemuxWithName(name string) (*GstMssDemux, error) {
	e, err := gst.NewElementWithName("mssdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstMssDemux{GstAdaptiveDemux: &GstAdaptiveDemux{Bin: &gst.Bin{Element: e}}}, nil
}

// Maximum buffers that can be stored in each internal stream queue (0 = infinite) (deprecated)
// Default: 0, Min: 0, Max: -1
func (e *GstMssDemux) SetMaxQueueSizeBuffers(value uint) error {
	return e.Element.SetProperty("max-queue-size-buffers", value)
}

func (e *GstMssDemux) GetMaxQueueSizeBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-queue-size-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}
