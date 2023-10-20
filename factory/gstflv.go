// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Demux FLV feeds into digital streams
type GstFlvDemux struct {
	*gst.Element
}

func NewFlvDemux() (*GstFlvDemux, error) {
	e, err := gst.NewElement("flvdemux")
	if err != nil {
		return nil, err
	}
	return &GstFlvDemux{Element: e}, nil
}

func NewFlvDemuxWithName(name string) (*GstFlvDemux, error) {
	e, err := gst.NewElementWithName("flvdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstFlvDemux{Element: e}, nil
}

// Muxes video/audio streams into a FLV stream
type GstFlvMux struct {
	*GstAggregator
}

func NewFlvMux() (*GstFlvMux, error) {
	e, err := gst.NewElement("flvmux")
	if err != nil {
		return nil, err
	}
	return &GstFlvMux{GstAggregator: &GstAggregator{Element: e}}, nil
}

func NewFlvMuxWithName(name string) (*GstFlvMux, error) {
	e, err := gst.NewElementWithName("flvmux", name)
	if err != nil {
		return nil, err
	}
	return &GstFlvMux{GstAggregator: &GstAggregator{Element: e}}, nil
}

// The value of encoder in the meta packet.
// Default: GStreamer {VERSION} FLV muxer
func (e *GstFlvMux) SetEncoder(value string) error {
	return e.Element.SetProperty("encoder", value)
}

func (e *GstFlvMux) GetEncoder() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("encoder")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// If set to true, flvmux will modify buffers timestamps to ensure they are always strictly increasing, inside one stream and also between the audio and video streams
// Default: true
func (e *GstFlvMux) SetEnforceIncreasingTimestamps(value bool) error {
	return e.Element.SetProperty("enforce-increasing-timestamps", value)
}

func (e *GstFlvMux) GetEnforceIncreasingTimestamps() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enforce-increasing-timestamps")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The value of metadatacreator in the meta packet.
// Default: GStreamer {VERSION} FLV muxer
func (e *GstFlvMux) SetMetadatacreator(value string) error {
	return e.Element.SetProperty("metadatacreator", value)
}

func (e *GstFlvMux) GetMetadatacreator() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("metadatacreator")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// If set to true, streams that go backwards related to the other stream will have buffers dropped until they reach the correct timestamp
// Default: false
func (e *GstFlvMux) SetSkipBackwardsStreams(value bool) error {
	return e.Element.SetProperty("skip-backwards-streams", value)
}

func (e *GstFlvMux) GetSkipBackwardsStreams() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("skip-backwards-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// If set to true, the output should be as if it is to be streamed and hence no indexes written or duration written.
// Default: false
func (e *GstFlvMux) SetStreamable(value bool) error {
	return e.Element.SetProperty("streamable", value)
}

func (e *GstFlvMux) GetStreamable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
