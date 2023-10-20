// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Add absolute timestamps and flags of recorded data in a playback session
type GstRtpOnvifParse struct {
	*gst.Element
}

func NewRtpOnvifParse() (*GstRtpOnvifParse, error) {
	e, err := gst.NewElement("rtponvifparse")
	if err != nil {
		return nil, err
	}
	return &GstRtpOnvifParse{Element: e}, nil
}

func NewRtpOnvifParseWithName(name string) (*GstRtpOnvifParse, error) {
	e, err := gst.NewElementWithName("rtponvifparse", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpOnvifParse{Element: e}, nil
}

// Add absolute timestamps and flags of recorded data in a playback session
type GstRtpOnvifTimestamp struct {
	*gst.Element
}

func NewRtpOnvifTimestamp() (*GstRtpOnvifTimestamp, error) {
	e, err := gst.NewElement("rtponviftimestamp")
	if err != nil {
		return nil, err
	}
	return &GstRtpOnvifTimestamp{Element: e}, nil
}

func NewRtpOnvifTimestampWithName(name string) (*GstRtpOnvifTimestamp, error) {
	e, err := gst.NewElementWithName("rtponviftimestamp", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpOnvifTimestamp{Element: e}, nil
}

// The RTSP CSeq which initiated the playback
// Default: 0, Min: 0, Max: -1
func (e *GstRtpOnvifTimestamp) SetCseq(value uint) error {
	return e.Element.SetProperty("cseq", value)
}

func (e *GstRtpOnvifTimestamp) GetCseq() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("cseq")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Whether the element should drop buffers that fall outside the segment, not part of the specification but allows full reverse playback.
// Default: true
func (e *GstRtpOnvifTimestamp) SetDropOutOfSegment(value bool) error {
	return e.Element.SetProperty("drop-out-of-segment", value)
}

func (e *GstRtpOnvifTimestamp) GetDropOutOfSegment() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-out-of-segment")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Offset between the pipeline running time and the absolute UTC time, in nano-seconds since 1900 (-1 for automatic computation)
// Default: 18446744073709551615, Min: 0, Max: 18446744073709551615
func (e *GstRtpOnvifTimestamp) SetNtpOffset(value uint64) error {
	return e.Element.SetProperty("ntp-offset", value)
}

func (e *GstRtpOnvifTimestamp) GetNtpOffset() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ntp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// If the element should set the 'E' bit as defined in the ONVIF RTP extension. This increases latency by one packet
// Default: false
func (e *GstRtpOnvifTimestamp) SetSetEBit(value bool) error {
	return e.Element.SetProperty("set-e-bit", value)
}

func (e *GstRtpOnvifTimestamp) GetSetEBit() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("set-e-bit")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// If the element should set the 'T' bit as defined in the ONVIF RTP extension. This increases latency by one packet
// Default: false
func (e *GstRtpOnvifTimestamp) SetSetTBit(value bool) error {
	return e.Element.SetProperty("set-t-bit", value)
}

func (e *GstRtpOnvifTimestamp) GetSetTBit() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("set-t-bit")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether the element should use reference UTC timestamps from the buffers instead of using the ntp-offset mechanism.
// Default: false
func (e *GstRtpOnvifTimestamp) SetUseReferenceTimestamps(value bool) error {
	return e.Element.SetProperty("use-reference-timestamps", value)
}

func (e *GstRtpOnvifTimestamp) GetUseReferenceTimestamps() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-reference-timestamps")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
