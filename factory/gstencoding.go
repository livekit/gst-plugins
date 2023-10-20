// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Convenience encoding/muxing element
type GstEncodeBin2 struct {
	*GstEncodeBaseBin
}

func NewEncodeBin2() (*GstEncodeBin2, error) {
	e, err := gst.NewElement("encodebin2")
	if err != nil {
		return nil, err
	}
	return &GstEncodeBin2{GstEncodeBaseBin: &GstEncodeBaseBin{Bin: &gst.Bin{Element: e}}}, nil
}

func NewEncodeBin2WithName(name string) (*GstEncodeBin2, error) {
	e, err := gst.NewElementWithName("encodebin2", name)
	if err != nil {
		return nil, err
	}
	return &GstEncodeBin2{GstEncodeBaseBin: &GstEncodeBaseBin{Bin: &gst.Bin{Element: e}}}, nil
}

// Recombines streams split by the streamsplitter element
type GstStreamCombiner struct {
	*gst.Element
}

func NewStreamCombiner() (*GstStreamCombiner, error) {
	e, err := gst.NewElement("streamcombiner")
	if err != nil {
		return nil, err
	}
	return &GstStreamCombiner{Element: e}, nil
}

func NewStreamCombinerWithName(name string) (*GstStreamCombiner, error) {
	e, err := gst.NewElementWithName("streamcombiner", name)
	if err != nil {
		return nil, err
	}
	return &GstStreamCombiner{Element: e}, nil
}

// Splits streams based on their media type
type GstStreamSplitter struct {
	*gst.Element
}

func NewStreamSplitter() (*GstStreamSplitter, error) {
	e, err := gst.NewElement("streamsplitter")
	if err != nil {
		return nil, err
	}
	return &GstStreamSplitter{Element: e}, nil
}

func NewStreamSplitterWithName(name string) (*GstStreamSplitter, error) {
	e, err := gst.NewElementWithName("streamsplitter", name)
	if err != nil {
		return nil, err
	}
	return &GstStreamSplitter{Element: e}, nil
}

// Convenience encoding/muxing element
type GstEncodeBin struct {
	*GstEncodeBaseBin
}

func NewEncodeBin() (*GstEncodeBin, error) {
	e, err := gst.NewElement("encodebin")
	if err != nil {
		return nil, err
	}
	return &GstEncodeBin{GstEncodeBaseBin: &GstEncodeBaseBin{Bin: &gst.Bin{Element: e}}}, nil
}

func NewEncodeBinWithName(name string) (*GstEncodeBin, error) {
	e, err := gst.NewElementWithName("encodebin", name)
	if err != nil {
		return nil, err
	}
	return &GstEncodeBin{GstEncodeBaseBin: &GstEncodeBaseBin{Bin: &gst.Bin{Element: e}}}, nil
}

// Whether to re-encode portions of compatible video streams that lay on segment boundaries
// Default: false
func (e *GstEncodeBin) SetAvoidReencoding(value bool) error {
	return e.Element.SetProperty("avoid-reencoding", value)
}

func (e *GstEncodeBin) GetAvoidReencoding() (bool, error) {
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

// Flags to control behaviour
// Default: (none)
func (e *GstEncodeBin) SetFlags(value GstEncodeBinFlags) error {
	e.Element.SetArg("flags", fmt.Sprint(value))
	return nil
}

func (e *GstEncodeBin) GetFlags() (GstEncodeBinFlags, error) {
	var value GstEncodeBinFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstEncodeBinFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstEncodeBinFlags")
	}
	return value, nil
}

// The GstEncodingProfile to use

func (e *GstEncodeBin) SetProfile(value interface{}) error {
	return e.Element.SetProperty("profile", value)
}

func (e *GstEncodeBin) GetProfile() (interface{}, error) {
	return e.Element.GetProperty("profile")
}

// Max. number of buffers in the queue (0=disable)
// Default: 200, Min: 0, Max: -1
func (e *GstEncodeBin) SetQueueBuffersMax(value uint) error {
	return e.Element.SetProperty("queue-buffers-max", value)
}

func (e *GstEncodeBin) GetQueueBuffersMax() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("queue-buffers-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max. amount of data in the queue (bytes, 0=disable)
// Default: 10485760, Min: 0, Max: -1
func (e *GstEncodeBin) SetQueueBytesMax(value uint) error {
	return e.Element.SetProperty("queue-bytes-max", value)
}

func (e *GstEncodeBin) GetQueueBytesMax() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("queue-bytes-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max. amount of data in the queue (in ns, 0=disable)
// Default: 1000000000, Min: 0, Max: 18446744073709551615
func (e *GstEncodeBin) SetQueueTimeMax(value uint64) error {
	return e.Element.SetProperty("queue-time-max", value)
}

func (e *GstEncodeBin) GetQueueTimeMax() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("queue-time-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Amount of timestamp jitter/imperfection to allow on audio streams before inserting/dropping samples (ns)
// Default: 20000000, Min: 0, Max: 18446744073709551615
func (e *GstEncodeBin) SetAudioJitterTolerance(value uint64) error {
	return e.Element.SetProperty("audio-jitter-tolerance", value)
}

func (e *GstEncodeBin) GetAudioJitterTolerance() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("audio-jitter-tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

type GstEncodeBaseBin struct {
	*gst.Bin
}

// Amount of timestamp jitter/imperfection to allow on audio streams before inserting/dropping samples (ns)
// Default: 20000000, Min: 0, Max: 18446744073709551615
func (e *GstEncodeBaseBin) SetAudioJitterTolerance(value uint64) error {
	return e.Element.SetProperty("audio-jitter-tolerance", value)
}

func (e *GstEncodeBaseBin) GetAudioJitterTolerance() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("audio-jitter-tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Whether to re-encode portions of compatible video streams that lay on segment boundaries
// Default: false
func (e *GstEncodeBaseBin) SetAvoidReencoding(value bool) error {
	return e.Element.SetProperty("avoid-reencoding", value)
}

func (e *GstEncodeBaseBin) GetAvoidReencoding() (bool, error) {
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

// Flags to control behaviour
// Default: (none)
func (e *GstEncodeBaseBin) SetFlags(value GstEncodeBinFlags) error {
	e.Element.SetArg("flags", fmt.Sprint(value))
	return nil
}

func (e *GstEncodeBaseBin) GetFlags() (GstEncodeBinFlags, error) {
	var value GstEncodeBinFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstEncodeBinFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstEncodeBinFlags")
	}
	return value, nil
}

// The GstEncodingProfile to use

func (e *GstEncodeBaseBin) SetProfile(value interface{}) error {
	return e.Element.SetProperty("profile", value)
}

func (e *GstEncodeBaseBin) GetProfile() (interface{}, error) {
	return e.Element.GetProperty("profile")
}

// Max. number of buffers in the queue (0=disable)
// Default: 200, Min: 0, Max: -1
func (e *GstEncodeBaseBin) SetQueueBuffersMax(value uint) error {
	return e.Element.SetProperty("queue-buffers-max", value)
}

func (e *GstEncodeBaseBin) GetQueueBuffersMax() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("queue-buffers-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max. amount of data in the queue (bytes, 0=disable)
// Default: 10485760, Min: 0, Max: -1
func (e *GstEncodeBaseBin) SetQueueBytesMax(value uint) error {
	return e.Element.SetProperty("queue-bytes-max", value)
}

func (e *GstEncodeBaseBin) GetQueueBytesMax() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("queue-bytes-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max. amount of data in the queue (in ns, 0=disable)
// Default: 1000000000, Min: 0, Max: 18446744073709551615
func (e *GstEncodeBaseBin) SetQueueTimeMax(value uint64) error {
	return e.Element.SetProperty("queue-time-max", value)
}

func (e *GstEncodeBaseBin) GetQueueTimeMax() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("queue-time-max")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

type GstEncodeBinFlags int

const (
	GstEncodeBinFlagsNoAudioConversion GstEncodeBinFlags = 0x00000001 // Do not use audio conversion elements
	GstEncodeBinFlagsNoVideoConversion GstEncodeBinFlags = 0x00000002 // Do not use video conversion elements
)
