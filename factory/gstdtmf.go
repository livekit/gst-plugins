// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Generates DTMF tones
type GstDTMFSrc struct {
	*base.GstBaseSrc
}

func NewDTMFSrc() (*GstDTMFSrc, error) {
	e, err := gst.NewElement("dtmfsrc")
	if err != nil {
		return nil, err
	}
	return &GstDTMFSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewDTMFSrcWithName(name string) (*GstDTMFSrc, error) {
	e, err := gst.NewElementWithName("dtmfsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstDTMFSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Interval in ms between two tone packets
// Default: 50, Min: 10, Max: 50
func (e *GstDTMFSrc) SetInterval(value uint) error {
	return e.Element.SetProperty("interval", value)
}

func (e *GstDTMFSrc) GetInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Generates DTMF Sound from telephone-event RTP packets
type GstRtpDTMFDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpDTMFDepay() (*GstRtpDTMFDepay, error) {
	e, err := gst.NewElement("rtpdtmfdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpDTMFDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpDTMFDepayWithName(name string) (*GstRtpDTMFDepay, error) {
	e, err := gst.NewElementWithName("rtpdtmfdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpDTMFDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// The maxumimum duration (ms) of the outgoing soundpacket. (0 = no limit)
// Default: 0, Min: 0, Max: -1
func (e *GstRtpDTMFDepay) SetMaxDuration(value uint) error {
	return e.Element.SetProperty("max-duration", value)
}

func (e *GstRtpDTMFDepay) GetMaxDuration() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The smallest unit (ms) the duration must be a multiple of (0 disables it)
// Default: 0, Min: 0, Max: 1000
func (e *GstRtpDTMFDepay) SetUnitTime(value uint) error {
	return e.Element.SetProperty("unit-time", value)
}

func (e *GstRtpDTMFDepay) GetUnitTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("unit-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Generates RTP DTMF packets
type GstRTPDTMFSrc struct {
	*base.GstBaseSrc
}

func NewRTPDTMFSrc() (*GstRTPDTMFSrc, error) {
	e, err := gst.NewElement("rtpdtmfsrc")
	if err != nil {
		return nil, err
	}
	return &GstRTPDTMFSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewRTPDTMFSrcWithName(name string) (*GstRTPDTMFSrc, error) {
	e, err := gst.NewElementWithName("rtpdtmfsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPDTMFSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// The clock-rate at which to generate the dtmf packets
// Default: 8000, Min: 0, Max: -1
func (e *GstRTPDTMFSrc) SetClockRate(value uint) error {
	return e.Element.SetProperty("clock-rate", value)
}

func (e *GstRTPDTMFSrc) GetClockRate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("clock-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of packets to send to indicate start and stop dtmf events
// Default: 1, Min: 1, Max: 5
func (e *GstRTPDTMFSrc) SetPacketRedundancy(value uint) error {
	return e.Element.SetProperty("packet-redundancy", value)
}

func (e *GstRTPDTMFSrc) GetPacketRedundancy() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("packet-redundancy")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The payload type of the packets
// Default: 96, Min: 0, Max: 128
func (e *GstRTPDTMFSrc) SetPt(value uint) error {
	return e.Element.SetProperty("pt", value)
}

func (e *GstRTPDTMFSrc) GetPt() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("pt")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The RTP sequence number of the last processed packet
// Default: 0, Min: 0, Max: -1
func (e *GstRTPDTMFSrc) GetSeqnum() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("seqnum")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Offset to add to all outgoing seqnum (-1 = random)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRTPDTMFSrc) SetSeqnumOffset(value int) error {
	return e.Element.SetProperty("seqnum-offset", value)
}

func (e *GstRTPDTMFSrc) GetSeqnumOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("seqnum-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The SSRC of the packets (-1 == random)
// Default: -1, Min: 0, Max: -1
func (e *GstRTPDTMFSrc) SetSsrc(value uint) error {
	return e.Element.SetProperty("ssrc", value)
}

func (e *GstRTPDTMFSrc) GetSsrc() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ssrc")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The RTP timestamp of the last processed packet
// Default: 0, Min: 0, Max: -1
func (e *GstRTPDTMFSrc) GetTimestamp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Offset to add to all outgoing timestamps (-1 = random)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRTPDTMFSrc) SetTimestampOffset(value int) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

func (e *GstRTPDTMFSrc) GetTimestampOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timestamp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}
