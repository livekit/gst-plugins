// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Multiplexes media streams into an ATSC-compliant Transport Stream
type GstATSCMux struct {
	*GstBaseTsMux
}

func NewATSCMux() (*GstATSCMux, error) {
	e, err := gst.NewElement("atscmux")
	if err != nil {
		return nil, err
	}
	return &GstATSCMux{GstBaseTsMux: &GstBaseTsMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func NewATSCMuxWithName(name string) (*GstATSCMux, error) {
	e, err := gst.NewElementWithName("atscmux", name)
	if err != nil {
		return nil, err
	}
	return &GstATSCMux{GstBaseTsMux: &GstBaseTsMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// Multiplexes media streams into an MPEG Transport Stream
type GstMpegTsMux struct {
	*GstBaseTsMux
}

func NewMpegTsMux() (*GstMpegTsMux, error) {
	e, err := gst.NewElement("mpegtsmux")
	if err != nil {
		return nil, err
	}
	return &GstMpegTsMux{GstBaseTsMux: &GstBaseTsMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func NewMpegTsMuxWithName(name string) (*GstMpegTsMux, error) {
	e, err := gst.NewElementWithName("mpegtsmux", name)
	if err != nil {
		return nil, err
	}
	return &GstMpegTsMux{GstBaseTsMux: &GstBaseTsMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// Set to TRUE to output Blu-Ray disc format with 192 byte packets. FALSE for standard TS format with 188 byte packets.
// Default: false
func (e *GstMpegTsMux) SetM2tsMode(value bool) error {
	return e.Element.SetProperty("m2ts-mode", value)
}

func (e *GstMpegTsMux) GetM2tsMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("m2ts-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstBaseTsMux struct {
	*GstAggregator
}

// Set the interval (in ticks of the 90kHz clock) for writing out the PMT table
// Default: 9000, Min: 1, Max: -1
func (e *GstBaseTsMux) SetPmtInterval(value uint) error {
	return e.Element.SetProperty("pmt-interval", value)
}

func (e *GstBaseTsMux) GetPmtInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("pmt-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// A GstStructure specifies the mapping from elementary streams to programs

func (e *GstBaseTsMux) SetProgMap(value *gst.Structure) error {
	return e.Element.SetProperty("prog-map", value)
}

func (e *GstBaseTsMux) GetProgMap() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("prog-map")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Set the interval (in ticks of the 90kHz clock) for writing out the ServiceInformation tables
// Default: 9000, Min: 1, Max: -1
func (e *GstBaseTsMux) SetSiInterval(value uint) error {
	return e.Element.SetProperty("si-interval", value)
}

func (e *GstBaseTsMux) GetSiInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("si-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of packets per buffer (padded with dummy packets on EOS) (-1 = auto, 0 = all available packets, 7 for UDP streaming)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstBaseTsMux) SetAlignment(value int) error {
	return e.Element.SetProperty("alignment", value)
}

func (e *GstBaseTsMux) GetAlignment() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("alignment")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Set the target bitrate, will insert null packets as padding  to achieve multiplex-wide constant bitrate (0 means no padding)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstBaseTsMux) SetBitrate(value uint64) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstBaseTsMux) GetBitrate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Set the interval (in ticks of the 90kHz clock) for writing out the PAT table
// Default: 9000, Min: 1, Max: -1
func (e *GstBaseTsMux) SetPatInterval(value uint) error {
	return e.Element.SetProperty("pat-interval", value)
}

func (e *GstBaseTsMux) GetPatInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("pat-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Set the interval (in ticks of the 90kHz clock) for writing PCR
// Default: 3600, Min: 1, Max: -1
func (e *GstBaseTsMux) SetPcrInterval(value uint) error {
	return e.Element.SetProperty("pcr-interval", value)
}

func (e *GstBaseTsMux) GetPcrInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("pcr-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Set the interval (in ticks of the 90kHz clock) for writing SCTE-35 NULL (heartbeat) packets. (only valid if scte-35-pid is different from 0)
// Default: 27000000, Min: 1, Max: -1
func (e *GstBaseTsMux) SetScte35NullInterval(value uint) error {
	return e.Element.SetProperty("scte-35-null-interval", value)
}

func (e *GstBaseTsMux) GetScte35NullInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("scte-35-null-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// PID to use for inserting SCTE-35 packets (0: unused)
// Default: 0, Min: 0, Max: -1
func (e *GstBaseTsMux) SetScte35Pid(value uint) error {
	return e.Element.SetProperty("scte-35-pid", value)
}

func (e *GstBaseTsMux) GetScte35Pid() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("scte-35-pid")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}
