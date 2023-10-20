// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Demuxes MPEG2 transport streams
type GstTSDemux struct {
	*MpegTSBase
}

func NewTSDemux() (*GstTSDemux, error) {
	e, err := gst.NewElement("tsdemux")
	if err != nil {
		return nil, err
	}
	return &GstTSDemux{MpegTSBase: &MpegTSBase{Element: e}}, nil
}

func NewTSDemuxWithName(name string) (*GstTSDemux, error) {
	e, err := gst.NewElementWithName("tsdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstTSDemux{MpegTSBase: &MpegTSBase{Element: e}}, nil
}

// Emit messages for every pcr/opcr/pts/dts
// Default: false
func (e *GstTSDemux) SetEmitStats(value bool) error {
	return e.Element.SetProperty("emit-stats", value)
}

func (e *GstTSDemux) GetEmitStats() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("emit-stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Latency to add for smooth demuxing (in ms)
// Default: 700, Min: -1, Max: 2147483647
func (e *GstTSDemux) SetLatency(value int) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstTSDemux) GetLatency() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Program Number to demux for (-1 to ignore)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstTSDemux) SetProgramNumber(value int) error {
	return e.Element.SetProperty("program-number", value)
}

func (e *GstTSDemux) GetProgramNumber() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("program-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Whether SCTE 35 sections should be forwarded as events
// Default: false
func (e *GstTSDemux) SetSendScte35Events(value bool) error {
	return e.Element.SetProperty("send-scte35-events", value)
}

func (e *GstTSDemux) GetSendScte35Events() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("send-scte35-events")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Parses MPEG2 transport streams
type MpegTSParse2 struct {
	*MpegTSBase
}

func NewMpegTSParse2() (*MpegTSParse2, error) {
	e, err := gst.NewElement("tsparse")
	if err != nil {
		return nil, err
	}
	return &MpegTSParse2{MpegTSBase: &MpegTSBase{Element: e}}, nil
}

func NewMpegTSParse2WithName(name string) (*MpegTSParse2, error) {
	e, err := gst.NewElementWithName("tsparse", name)
	if err != nil {
		return nil, err
	}
	return &MpegTSParse2{MpegTSBase: &MpegTSBase{Element: e}}, nil
}

// If set, timestamps will be set on the output buffers using PCRs and smoothed over the smoothing-latency period
// Default: false
func (e *MpegTSParse2) SetSetTimestamps(value bool) error {
	return e.Element.SetProperty("set-timestamps", value)
}

func (e *MpegTSParse2) GetSetTimestamps() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("set-timestamps")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Additional latency in microseconds for smoothing jitter in input timestamps on live capture
// Default: 0, Min: 0, Max: -1
func (e *MpegTSParse2) SetSmoothingLatency(value uint) error {
	return e.Element.SetProperty("smoothing-latency", value)
}

func (e *MpegTSParse2) GetSmoothingLatency() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("smoothing-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// If set, buffers sized smaller than the alignment will be sent so that RAI packets are at the start of a new buffer
// Default: false
func (e *MpegTSParse2) SetSplitOnRai(value bool) error {
	return e.Element.SetProperty("split-on-rai", value)
}

func (e *MpegTSParse2) GetSplitOnRai() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("split-on-rai")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of packets per buffer (padded with dummy packets on EOS) (0 = auto)
// Default: 0, Min: 0, Max: -1
func (e *MpegTSParse2) SetAlignment(value uint) error {
	return e.Element.SetProperty("alignment", value)
}

func (e *MpegTSParse2) GetAlignment() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("alignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Set the PID to use for PCR values (-1 for auto)
// Default: -1, Min: -1, Max: 2147483647
func (e *MpegTSParse2) SetPcrPid(value int) error {
	return e.Element.SetProperty("pcr-pid", value)
}

func (e *MpegTSParse2) GetPcrPid() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pcr-pid")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type MpegTSBase struct {
	*gst.Element
}

// Ignore PCR stream for timing
// Default: false
func (e *MpegTSBase) SetIgnorePcr(value bool) error {
	return e.Element.SetProperty("ignore-pcr", value)
}

func (e *MpegTSBase) GetIgnorePcr() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-pcr")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Parse private sections
// Default: false
func (e *MpegTSBase) SetParsePrivateSections(value bool) error {
	return e.Element.SetProperty("parse-private-sections", value)
}

func (e *MpegTSBase) GetParsePrivateSections() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("parse-private-sections")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
