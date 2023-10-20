// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Access descramble and split DVB streams
type DvbBaseBin struct {
	*gst.Bin
}

func NewDvbBaseBin() (*DvbBaseBin, error) {
	e, err := gst.NewElement("dvbbasebin")
	if err != nil {
		return nil, err
	}
	return &DvbBaseBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewDvbBaseBinWithName(name string) (*DvbBaseBin, error) {
	e, err := gst.NewElementWithName("dvbbasebin", name)
	if err != nil {
		return nil, err
	}
	return &DvbBaseBin{Bin: &gst.Bin{Element: e}}, nil
}

// (ISDB-T) Layer C modulation type
// Default: auto (6)
func (e *DvbBaseBin) SetIsdbtLayercModulation(value GstDvbSrcModulation) error {
	e.Element.SetArg("isdbt-layerc-modulation", string(value))
	return nil
}

func (e *DvbBaseBin) GetIsdbtLayercModulation() (GstDvbSrcModulation, error) {
	var value GstDvbSrcModulation
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerc-modulation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcModulation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcModulation")
	}
	return value, nil
}

// (ISDB-T) Partial Reception (-1 = AUTO)
// Default: 1, Min: -1, Max: 1
func (e *DvbBaseBin) SetIsdbtPartialReception(value int) error {
	return e.Element.SetProperty("isdbt-partial-reception", value)
}

func (e *DvbBaseBin) GetIsdbtPartialReception() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-partial-reception")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (ISDB-T) SB segment count
// Default: 1, Min: 1, Max: 13
func (e *DvbBaseBin) SetIsdbtSbSegmentCount(value uint) error {
	return e.Element.SetProperty("isdbt-sb-segment-count", value)
}

func (e *DvbBaseBin) GetIsdbtSbSegmentCount() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("isdbt-sb-segment-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Delivery System
// Default: undefined (0)
func (e *DvbBaseBin) SetDelsys(value GstDvbSrcDelsys) error {
	e.Element.SetArg("delsys", string(value))
	return nil
}

func (e *DvbBaseBin) GetDelsys() (GstDvbSrcDelsys, error) {
	var value GstDvbSrcDelsys
	var ok bool
	v, err := e.Element.GetProperty("delsys")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcDelsys)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcDelsys")
	}
	return value, nil
}

// Center frequency to tune into. Measured in kHz for the satellite distribution standards and Hz for all the rest
// Default: 0, Min: 0, Max: -1
func (e *DvbBaseBin) SetFrequency(value uint) error {
	return e.Element.SetProperty("frequency", value)
}

func (e *DvbBaseBin) GetFrequency() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// (DVB-T) Guard Interval
// Default: 16 (1)
func (e *DvbBaseBin) SetGuard(value GstDvbSrcGuard) error {
	e.Element.SetArg("guard", string(value))
	return nil
}

func (e *DvbBaseBin) GetGuard() (GstDvbSrcGuard, error) {
	var value GstDvbSrcGuard
	var ok bool
	v, err := e.Element.GetProperty("guard")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcGuard)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcGuard")
	}
	return value, nil
}

// (ISDB-T) layer A Forward Error Correction
// Default: auto (9)
func (e *DvbBaseBin) SetIsdbtLayeraFec(value GstDvbSrcCodeRate) error {
	return e.Element.SetProperty("isdbt-layera-fec", value)
}

func (e *DvbBaseBin) GetIsdbtLayeraFec() (GstDvbSrcCodeRate, error) {
	var value GstDvbSrcCodeRate
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layera-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcCodeRate)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcCodeRate")
	}
	return value, nil
}

// (ISDB-T) Layer B modulation type
// Default: auto (6)
func (e *DvbBaseBin) SetIsdbtLayerbModulation(value GstDvbSrcModulation) error {
	e.Element.SetArg("isdbt-layerb-modulation", string(value))
	return nil
}

func (e *DvbBaseBin) GetIsdbtLayerbModulation() (GstDvbSrcModulation, error) {
	var value GstDvbSrcModulation
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerb-modulation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcModulation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcModulation")
	}
	return value, nil
}

// (DVB-T/T2/C/S2, TURBO and ATSC) Modulation type
// Default: qam-16 (1)
func (e *DvbBaseBin) SetModulation(value GstDvbSrcModulation) error {
	e.Element.SetArg("modulation", string(value))
	return nil
}

func (e *DvbBaseBin) GetModulation() (GstDvbSrcModulation, error) {
	var value GstDvbSrcModulation
	var ok bool
	v, err := e.Element.GetProperty("modulation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcModulation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcModulation")
	}
	return value, nil
}

// Rolloff (DVB-S2)
// Default: auto (3)
func (e *DvbBaseBin) SetRolloff(value GstDvbSrcRolloff) error {
	e.Element.SetArg("rolloff", string(value))
	return nil
}

func (e *DvbBaseBin) GetRolloff() (GstDvbSrcRolloff, error) {
	var value GstDvbSrcRolloff
	var ok bool
	v, err := e.Element.GetProperty("rolloff")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcRolloff)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcRolloff")
	}
	return value, nil
}

// (ISDB-T) Layer A segment count (-1 = AUTO)
// Default: -1, Min: -1, Max: 13
func (e *DvbBaseBin) SetIsdbtLayeraSegmentCount(value int) error {
	return e.Element.SetProperty("isdbt-layera-segment-count", value)
}

func (e *DvbBaseBin) GetIsdbtLayeraSegmentCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layera-segment-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (ISDB-T) SB segment IDX
// Default: 0, Min: 0, Max: 12
func (e *DvbBaseBin) SetIsdbtSbSegmentIdx(value int) error {
	return e.Element.SetProperty("isdbt-sb-segment-idx", value)
}

func (e *DvbBaseBin) GetIsdbtSbSegmentIdx() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-sb-segment-idx")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (ISDB-T) SB Subchannel ID (-1 = AUTO)
// Default: -1, Min: -1, Max: 41
func (e *DvbBaseBin) SetIsdbtSbSubchannelId(value int) error {
	return e.Element.SetProperty("isdbt-sb-subchannel-id", value)
}

func (e *DvbBaseBin) GetIsdbtSbSubchannelId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-sb-subchannel-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (DVB-S/S2) Polarity [vhHV] (eg. V for Vertical)
// Default: H
func (e *DvbBaseBin) SetPolarity(value string) error {
	return e.Element.SetProperty("polarity", value)
}

func (e *DvbBaseBin) GetPolarity() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("polarity")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Microseconds to wait before giving up tuning/locking on a signal
// Default: 10000000000, Min: 0, Max: 18446744073709551615
func (e *DvbBaseBin) SetTuningTimeout(value uint64) error {
	return e.Element.SetProperty("tuning-timeout", value)
}

func (e *DvbBaseBin) GetTuningTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("tuning-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// (ISDB-T) Sound Broadcasting
// Default: 0, Min: 0, Max: 1
func (e *DvbBaseBin) SetIsdbtSoundBroadcasting(value int) error {
	return e.Element.SetProperty("isdbt-sound-broadcasting", value)
}

func (e *DvbBaseBin) GetIsdbtSoundBroadcasting() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-sound-broadcasting")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// LNB's Upper bound for low band reception (kHz)
// Default: 11700000, Min: 0, Max: -1
func (e *DvbBaseBin) SetLnbSlof(value uint) error {
	return e.Element.SetProperty("lnb-slof", value)
}

func (e *DvbBaseBin) GetLnbSlof() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("lnb-slof")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The DVB adapter device number (eg. 0 for adapter0)
// Default: 0, Min: 0, Max: 16
func (e *DvbBaseBin) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

func (e *DvbBaseBin) GetAdapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (DVB-T, DVB-S/S2 and DVB-C) High priority code rate
// Default: auto (9)
func (e *DvbBaseBin) SetCodeRateHp(value GstDvbSrcCodeRate) error {
	return e.Element.SetProperty("code-rate-hp", value)
}

func (e *DvbBaseBin) GetCodeRateHp() (GstDvbSrcCodeRate, error) {
	var value GstDvbSrcCodeRate
	var ok bool
	v, err := e.Element.GetProperty("code-rate-hp")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcCodeRate)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcCodeRate")
	}
	return value, nil
}

// (DTMB) Interleaving type
// Default: auto (1)
func (e *DvbBaseBin) SetInterleaving(value GstDvbSrcInterleaving) error {
	e.Element.SetArg("interleaving", string(value))
	return nil
}

func (e *DvbBaseBin) GetInterleaving() (GstDvbSrcInterleaving, error) {
	var value GstDvbSrcInterleaving
	var ok bool
	v, err := e.Element.GetProperty("interleaving")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcInterleaving)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcInterleaving")
	}
	return value, nil
}

// (ISDB-T) Layer Enabled (7 = All layers)
// Default: 7, Min: 1, Max: 7
func (e *DvbBaseBin) SetIsdbtLayerEnabled(value uint) error {
	return e.Element.SetProperty("isdbt-layer-enabled", value)
}

func (e *DvbBaseBin) GetIsdbtLayerEnabled() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layer-enabled")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// (ISDB-T) Layer A modulation type
// Default: auto (6)
func (e *DvbBaseBin) SetIsdbtLayeraModulation(value GstDvbSrcModulation) error {
	e.Element.SetArg("isdbt-layera-modulation", string(value))
	return nil
}

func (e *DvbBaseBin) GetIsdbtLayeraModulation() (GstDvbSrcModulation, error) {
	var value GstDvbSrcModulation
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layera-modulation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcModulation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcModulation")
	}
	return value, nil
}

// (DVB-T) Bandwidth. Deprecated
// Default: AUTO (3)
func (e *DvbBaseBin) SetBandwidth(value GstDvbSrcBandwidth) error {
	e.Element.SetArg("bandwidth", string(value))
	return nil
}

func (e *DvbBaseBin) GetBandwidth() (GstDvbSrcBandwidth, error) {
	var value GstDvbSrcBandwidth
	var ok bool
	v, err := e.Element.GetProperty("bandwidth")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcBandwidth)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcBandwidth")
	}
	return value, nil
}

// (DVB-T) Hierarchy information
// Default: 1 (1)
func (e *DvbBaseBin) SetHierarchy(value GstDvbSrcHierarchy) error {
	e.Element.SetArg("hierarchy", string(value))
	return nil
}

func (e *DvbBaseBin) GetHierarchy() (GstDvbSrcHierarchy, error) {
	var value GstDvbSrcHierarchy
	var ok bool
	v, err := e.Element.GetProperty("hierarchy")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcHierarchy)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcHierarchy")
	}
	return value, nil
}

// LNB's Local oscillator frequency used for low band reception (kHz)
// Default: 9750000, Min: 0, Max: -1
func (e *DvbBaseBin) SetLnbLof1(value uint) error {
	return e.Element.SetProperty("lnb-lof1", value)
}

func (e *DvbBaseBin) GetLnbLof1() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("lnb-lof1")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Colon separated list of programs
// Default: NULL
func (e *DvbBaseBin) SetProgramNumbers(value string) error {
	return e.Element.SetProperty("program-numbers", value)
}

func (e *DvbBaseBin) GetProgramNumbers() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("program-numbers")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// (DVB-T) Transmission mode
// Default: 8k (1)
func (e *DvbBaseBin) SetTransMode(value GstDvbSrcTransmissionMode) error {
	return e.Element.SetProperty("trans-mode", value)
}

func (e *DvbBaseBin) GetTransMode() (GstDvbSrcTransmissionMode, error) {
	var value GstDvbSrcTransmissionMode
	var ok bool
	v, err := e.Element.GetProperty("trans-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcTransmissionMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcTransmissionMode")
	}
	return value, nil
}

// (DVB-S/S2) Selected DiSEqC source. Only needed if you have a DiSEqC switch. Otherwise leave at -1 (disabled)
// Default: -1, Min: -1, Max: 7
func (e *DvbBaseBin) SetDiseqcSource(value int) error {
	return e.Element.SetProperty("diseqc-source", value)
}

func (e *DvbBaseBin) GetDiseqcSource() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("diseqc-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The frontend device number (eg. 0 for frontend0)
// Default: 0, Min: 0, Max: 16
func (e *DvbBaseBin) SetFrontend(value int) error {
	return e.Element.SetProperty("frontend", value)
}

func (e *DvbBaseBin) GetFrontend() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("frontend")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (ISDB-T) layer B Forward Error Correction
// Default: auto (9)
func (e *DvbBaseBin) SetIsdbtLayerbFec(value GstDvbSrcCodeRate) error {
	return e.Element.SetProperty("isdbt-layerb-fec", value)
}

func (e *DvbBaseBin) GetIsdbtLayerbFec() (GstDvbSrcCodeRate, error) {
	var value GstDvbSrcCodeRate
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerb-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcCodeRate)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcCodeRate")
	}
	return value, nil
}

// (ISDB-T) Layer B time interleaving (-1 = AUTO)
// Default: -1, Min: -1, Max: 8
func (e *DvbBaseBin) SetIsdbtLayerbTimeInterleaving(value int) error {
	return e.Element.SetProperty("isdbt-layerb-time-interleaving", value)
}

func (e *DvbBaseBin) GetIsdbtLayerbTimeInterleaving() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerb-time-interleaving")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (ISDB-T) Layer C time interleaving (-1 = AUTO)
// Default: -1, Min: -1, Max: 8
func (e *DvbBaseBin) SetIsdbtLayercTimeInterleaving(value int) error {
	return e.Element.SetProperty("isdbt-layerc-time-interleaving", value)
}

func (e *DvbBaseBin) GetIsdbtLayercTimeInterleaving() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerc-time-interleaving")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Channel bandwidth in Hz
// Default: 0, Min: 0, Max: -1
func (e *DvbBaseBin) SetBandwidthHz(value uint) error {
	return e.Element.SetProperty("bandwidth-hz", value)
}

func (e *DvbBaseBin) GetBandwidthHz() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("bandwidth-hz")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// (ISDB-T) Layer A time interleaving (-1 = AUTO)
// Default: -1, Min: -1, Max: 8
func (e *DvbBaseBin) SetIsdbtLayeraTimeInterleaving(value int) error {
	return e.Element.SetProperty("isdbt-layera-time-interleaving", value)
}

func (e *DvbBaseBin) GetIsdbtLayeraTimeInterleaving() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layera-time-interleaving")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (ISDB-T) layer C Forward Error Correction
// Default: auto (9)
func (e *DvbBaseBin) SetIsdbtLayercFec(value GstDvbSrcCodeRate) error {
	return e.Element.SetProperty("isdbt-layerc-fec", value)
}

func (e *DvbBaseBin) GetIsdbtLayercFec() (GstDvbSrcCodeRate, error) {
	var value GstDvbSrcCodeRate
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerc-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcCodeRate)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcCodeRate")
	}
	return value, nil
}

// (ISDB-T) Layer C segment count (-1 = AUTO)
// Default: -1, Min: -1, Max: 13
func (e *DvbBaseBin) SetIsdbtLayercSegmentCount(value int) error {
	return e.Element.SetProperty("isdbt-layerc-segment-count", value)
}

func (e *DvbBaseBin) GetIsdbtLayercSegmentCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerc-segment-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (DVB-T2 and DVB-S2 max 255, ISDB max 65535) Stream ID (-1 = disabled)
// Default: -1, Min: -1, Max: 65535
func (e *DvbBaseBin) SetStreamId(value int) error {
	return e.Element.SetProperty("stream-id", value)
}

func (e *DvbBaseBin) GetStreamId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("stream-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (DVB-T) Low priority code rate
// Default: 1/2 (1)
func (e *DvbBaseBin) SetCodeRateLp(value GstDvbSrcCodeRate) error {
	return e.Element.SetProperty("code-rate-lp", value)
}

func (e *DvbBaseBin) GetCodeRateLp() (GstDvbSrcCodeRate, error) {
	var value GstDvbSrcCodeRate
	var ok bool
	v, err := e.Element.GetProperty("code-rate-lp")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcCodeRate)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcCodeRate")
	}
	return value, nil
}

// (DVB-T and DVB-C) Inversion information
// Default: on (1)
func (e *DvbBaseBin) SetInversion(value GstDvbSrcInversion) error {
	e.Element.SetArg("inversion", string(value))
	return nil
}

func (e *DvbBaseBin) GetInversion() (GstDvbSrcInversion, error) {
	var value GstDvbSrcInversion
	var ok bool
	v, err := e.Element.GetProperty("inversion")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcInversion)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcInversion")
	}
	return value, nil
}

// (ISDB-T) Layer B segment count (-1 = AUTO)
// Default: -1, Min: -1, Max: 13
func (e *DvbBaseBin) SetIsdbtLayerbSegmentCount(value int) error {
	return e.Element.SetProperty("isdbt-layerb-segment-count", value)
}

func (e *DvbBaseBin) GetIsdbtLayerbSegmentCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerb-segment-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// LNB's Local oscillator frequency used for high band reception (kHz)
// Default: 10600000, Min: 0, Max: -1
func (e *DvbBaseBin) SetLnbLof2(value uint) error {
	return e.Element.SetProperty("lnb-lof2", value)
}

func (e *DvbBaseBin) GetLnbLof2() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("lnb-lof2")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Pilot (DVB-S2)
// Default: auto (2)
func (e *DvbBaseBin) SetPilot(value GstDvbSrcPilot) error {
	e.Element.SetArg("pilot", string(value))
	return nil
}

func (e *DvbBaseBin) GetPilot() (GstDvbSrcPilot, error) {
	var value GstDvbSrcPilot
	var ok bool
	v, err := e.Element.GetProperty("pilot")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcPilot)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcPilot")
	}
	return value, nil
}

// The number of reads before reporting frontend stats
// Default: 100, Min: 0, Max: -1
func (e *DvbBaseBin) SetStatsReportingInterval(value uint) error {
	return e.Element.SetProperty("stats-reporting-interval", value)
}

func (e *DvbBaseBin) GetStatsReportingInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("stats-reporting-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// (DVB-S/S2, DVB-C) Symbol rate in kBd (kilo bauds)
// Default: 0, Min: 0, Max: -1
func (e *DvbBaseBin) SetSymbolRate(value uint) error {
	return e.Element.SetProperty("symbol-rate", value)
}

func (e *DvbBaseBin) GetSymbolRate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("symbol-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Digital Video Broadcast Source
type GstDvbSrc struct {
	*base.GstPushSrc
}

func NewDvbSrc() (*GstDvbSrc, error) {
	e, err := gst.NewElement("dvbsrc")
	if err != nil {
		return nil, err
	}
	return &GstDvbSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewDvbSrcWithName(name string) (*GstDvbSrc, error) {
	e, err := gst.NewElementWithName("dvbsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstDvbSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Microseconds to wait before giving up tuning/locking on a signal
// Default: 10000000000, Min: 0, Max: 18446744073709551615
func (e *GstDvbSrc) SetTuningTimeout(value uint64) error {
	return e.Element.SetProperty("tuning-timeout", value)
}

func (e *GstDvbSrc) GetTuningTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("tuning-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// (ISDB-T) Layer C modulation type
// Default: auto (6)
func (e *GstDvbSrc) SetIsdbtLayercModulation(value GstDvbSrcModulation) error {
	e.Element.SetArg("isdbt-layerc-modulation", string(value))
	return nil
}

func (e *GstDvbSrc) GetIsdbtLayercModulation() (GstDvbSrcModulation, error) {
	var value GstDvbSrcModulation
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerc-modulation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcModulation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcModulation")
	}
	return value, nil
}

// (ISDB-T) Partial Reception (-1 = AUTO)
// Default: 1, Min: -1, Max: 1
func (e *GstDvbSrc) SetIsdbtPartialReception(value int) error {
	return e.Element.SetProperty("isdbt-partial-reception", value)
}

func (e *GstDvbSrc) GetIsdbtPartialReception() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-partial-reception")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (DVB-T/T2/C/S2, TURBO and ATSC) Modulation type
// Default: qam-16 (1)
func (e *GstDvbSrc) SetModulation(value GstDvbSrcModulation) error {
	e.Element.SetArg("modulation", string(value))
	return nil
}

func (e *GstDvbSrc) GetModulation() (GstDvbSrcModulation, error) {
	var value GstDvbSrcModulation
	var ok bool
	v, err := e.Element.GetProperty("modulation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcModulation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcModulation")
	}
	return value, nil
}

// (ISDB-T) layer C Forward Error Correction
// Default: auto (9)
func (e *GstDvbSrc) SetIsdbtLayercFec(value GstDvbSrcCodeRate) error {
	return e.Element.SetProperty("isdbt-layerc-fec", value)
}

func (e *GstDvbSrc) GetIsdbtLayercFec() (GstDvbSrcCodeRate, error) {
	var value GstDvbSrcCodeRate
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerc-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcCodeRate)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcCodeRate")
	}
	return value, nil
}

// (ISDB-T) Sound Broadcasting
// Default: 0, Min: 0, Max: 1
func (e *GstDvbSrc) SetIsdbtSoundBroadcasting(value int) error {
	return e.Element.SetProperty("isdbt-sound-broadcasting", value)
}

func (e *GstDvbSrc) GetIsdbtSoundBroadcasting() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-sound-broadcasting")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (DVB-S/S2, DVB-C) Symbol rate in kBd (kilo bauds)
// Default: 0, Min: 0, Max: -1
func (e *GstDvbSrc) SetSymbolRate(value uint) error {
	return e.Element.SetProperty("symbol-rate", value)
}

func (e *GstDvbSrc) GetSymbolRate() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("symbol-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// (DVB-T) Transmission mode
// Default: 8k (1)
func (e *GstDvbSrc) SetTransMode(value GstDvbSrcTransmissionMode) error {
	return e.Element.SetProperty("trans-mode", value)
}

func (e *GstDvbSrc) GetTransMode() (GstDvbSrcTransmissionMode, error) {
	var value GstDvbSrcTransmissionMode
	var ok bool
	v, err := e.Element.GetProperty("trans-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcTransmissionMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcTransmissionMode")
	}
	return value, nil
}

// The frontend device number (eg. 0 for frontend0)
// Default: 0, Min: 0, Max: 16
func (e *GstDvbSrc) SetFrontend(value int) error {
	return e.Element.SetProperty("frontend", value)
}

func (e *GstDvbSrc) GetFrontend() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("frontend")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (DTMB) Interleaving type
// Default: auto (1)
func (e *GstDvbSrc) SetInterleaving(value GstDvbSrcInterleaving) error {
	e.Element.SetArg("interleaving", string(value))
	return nil
}

func (e *GstDvbSrc) GetInterleaving() (GstDvbSrcInterleaving, error) {
	var value GstDvbSrcInterleaving
	var ok bool
	v, err := e.Element.GetProperty("interleaving")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcInterleaving)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcInterleaving")
	}
	return value, nil
}

// (ISDB-T) Layer A time interleaving (-1 = AUTO)
// Default: -1, Min: -1, Max: 8
func (e *GstDvbSrc) SetIsdbtLayeraTimeInterleaving(value int) error {
	return e.Element.SetProperty("isdbt-layera-time-interleaving", value)
}

func (e *GstDvbSrc) GetIsdbtLayeraTimeInterleaving() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layera-time-interleaving")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (DVB-T) Guard Interval
// Default: 16 (1)
func (e *GstDvbSrc) SetGuard(value GstDvbSrcGuard) error {
	e.Element.SetArg("guard", string(value))
	return nil
}

func (e *GstDvbSrc) GetGuard() (GstDvbSrcGuard, error) {
	var value GstDvbSrcGuard
	var ok bool
	v, err := e.Element.GetProperty("guard")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcGuard)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcGuard")
	}
	return value, nil
}

// (ISDB-T) Layer B time interleaving (-1 = AUTO)
// Default: -1, Min: -1, Max: 8
func (e *GstDvbSrc) SetIsdbtLayerbTimeInterleaving(value int) error {
	return e.Element.SetProperty("isdbt-layerb-time-interleaving", value)
}

func (e *GstDvbSrc) GetIsdbtLayerbTimeInterleaving() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerb-time-interleaving")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (ISDB-T) Layer C segment count (-1 = AUTO)
// Default: -1, Min: -1, Max: 13
func (e *GstDvbSrc) SetIsdbtLayercSegmentCount(value int) error {
	return e.Element.SetProperty("isdbt-layerc-segment-count", value)
}

func (e *GstDvbSrc) GetIsdbtLayercSegmentCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerc-segment-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (ISDB-T) SB segment IDX
// Default: 0, Min: 0, Max: 12
func (e *GstDvbSrc) SetIsdbtSbSegmentIdx(value int) error {
	return e.Element.SetProperty("isdbt-sb-segment-idx", value)
}

func (e *GstDvbSrc) GetIsdbtSbSegmentIdx() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-sb-segment-idx")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Colon-separated list of PIDs (eg. 110:120) to capture. ACT and CAT are automatically included but PMT should be added explicitly. Special value 8192 gets full MPEG-TS
// Default: 8192
func (e *GstDvbSrc) SetPids(value string) error {
	return e.Element.SetProperty("pids", value)
}

// The DVB adapter device number (eg. 0 for adapter0)
// Default: 0, Min: 0, Max: 16
func (e *GstDvbSrc) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

func (e *GstDvbSrc) GetAdapter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("adapter")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (DVB-T) Bandwidth. Deprecated
// Default: AUTO (3)
func (e *GstDvbSrc) SetBandwidth(value GstDvbSrcBandwidth) error {
	e.Element.SetArg("bandwidth", string(value))
	return nil
}

func (e *GstDvbSrc) GetBandwidth() (GstDvbSrcBandwidth, error) {
	var value GstDvbSrcBandwidth
	var ok bool
	v, err := e.Element.GetProperty("bandwidth")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcBandwidth)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcBandwidth")
	}
	return value, nil
}

// The kernel buffer size used by the DVB api
// Default: 1925120, Min: 0, Max: -1
func (e *GstDvbSrc) SetDvbBufferSize(value uint) error {
	return e.Element.SetProperty("dvb-buffer-size", value)
}

func (e *GstDvbSrc) GetDvbBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("dvb-buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// (DVB-T and DVB-C) Inversion information
// Default: on (1)
func (e *GstDvbSrc) SetInversion(value GstDvbSrcInversion) error {
	e.Element.SetArg("inversion", string(value))
	return nil
}

func (e *GstDvbSrc) GetInversion() (GstDvbSrcInversion, error) {
	var value GstDvbSrcInversion
	var ok bool
	v, err := e.Element.GetProperty("inversion")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcInversion)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcInversion")
	}
	return value, nil
}

// (ISDB-T) Layer Enabled (7 = All layers)
// Default: 7, Min: 1, Max: 7
func (e *GstDvbSrc) SetIsdbtLayerEnabled(value uint) error {
	return e.Element.SetProperty("isdbt-layer-enabled", value)
}

func (e *GstDvbSrc) GetIsdbtLayerEnabled() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layer-enabled")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// (ISDB-T) Layer A modulation type
// Default: auto (6)
func (e *GstDvbSrc) SetIsdbtLayeraModulation(value GstDvbSrcModulation) error {
	e.Element.SetArg("isdbt-layera-modulation", string(value))
	return nil
}

func (e *GstDvbSrc) GetIsdbtLayeraModulation() (GstDvbSrcModulation, error) {
	var value GstDvbSrcModulation
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layera-modulation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcModulation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcModulation")
	}
	return value, nil
}

// Rolloff (DVB-S2)
// Default: auto (3)
func (e *GstDvbSrc) SetRolloff(value GstDvbSrcRolloff) error {
	e.Element.SetArg("rolloff", string(value))
	return nil
}

func (e *GstDvbSrc) GetRolloff() (GstDvbSrcRolloff, error) {
	var value GstDvbSrcRolloff
	var ok bool
	v, err := e.Element.GetProperty("rolloff")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcRolloff)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcRolloff")
	}
	return value, nil
}

// Delivery System
// Default: undefined (0)
func (e *GstDvbSrc) SetDelsys(value GstDvbSrcDelsys) error {
	e.Element.SetArg("delsys", string(value))
	return nil
}

func (e *GstDvbSrc) GetDelsys() (GstDvbSrcDelsys, error) {
	var value GstDvbSrcDelsys
	var ok bool
	v, err := e.Element.GetProperty("delsys")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcDelsys)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcDelsys")
	}
	return value, nil
}

// (DVB-S/S2) Selected DiSEqC source. Only needed if you have a DiSEqC switch. Otherwise leave at -1 (disabled)
// Default: -1, Min: -1, Max: 7
func (e *GstDvbSrc) SetDiseqcSource(value int) error {
	return e.Element.SetProperty("diseqc-source", value)
}

func (e *GstDvbSrc) GetDiseqcSource() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("diseqc-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Center frequency to tune into. Measured in kHz for the satellite distribution standards and Hz for all the rest
// Default: 0, Min: 0, Max: -1
func (e *GstDvbSrc) SetFrequency(value uint) error {
	return e.Element.SetProperty("frequency", value)
}

func (e *GstDvbSrc) GetFrequency() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("frequency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// LNB's Local oscillator frequency used for low band reception (kHz)
// Default: 9750000, Min: 0, Max: -1
func (e *GstDvbSrc) SetLnbLof1(value uint) error {
	return e.Element.SetProperty("lnb-lof1", value)
}

func (e *GstDvbSrc) GetLnbLof1() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("lnb-lof1")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// LNB's Local oscillator frequency used for high band reception (kHz)
// Default: 10600000, Min: 0, Max: -1
func (e *GstDvbSrc) SetLnbLof2(value uint) error {
	return e.Element.SetProperty("lnb-lof2", value)
}

func (e *GstDvbSrc) GetLnbLof2() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("lnb-lof2")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Pilot (DVB-S2)
// Default: auto (2)
func (e *GstDvbSrc) SetPilot(value GstDvbSrcPilot) error {
	e.Element.SetArg("pilot", string(value))
	return nil
}

func (e *GstDvbSrc) GetPilot() (GstDvbSrcPilot, error) {
	var value GstDvbSrcPilot
	var ok bool
	v, err := e.Element.GetProperty("pilot")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcPilot)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcPilot")
	}
	return value, nil
}

// (DVB-T, DVB-S/S2 and DVB-C) High priority code rate
// Default: auto (9)
func (e *GstDvbSrc) SetCodeRateHp(value GstDvbSrcCodeRate) error {
	return e.Element.SetProperty("code-rate-hp", value)
}

func (e *GstDvbSrc) GetCodeRateHp() (GstDvbSrcCodeRate, error) {
	var value GstDvbSrcCodeRate
	var ok bool
	v, err := e.Element.GetProperty("code-rate-hp")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcCodeRate)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcCodeRate")
	}
	return value, nil
}

// (ISDB-T) Layer A segment count (-1 = AUTO)
// Default: -1, Min: -1, Max: 13
func (e *GstDvbSrc) SetIsdbtLayeraSegmentCount(value int) error {
	return e.Element.SetProperty("isdbt-layera-segment-count", value)
}

func (e *GstDvbSrc) GetIsdbtLayeraSegmentCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layera-segment-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (ISDB-T) Layer B segment count (-1 = AUTO)
// Default: -1, Min: -1, Max: 13
func (e *GstDvbSrc) SetIsdbtLayerbSegmentCount(value int) error {
	return e.Element.SetProperty("isdbt-layerb-segment-count", value)
}

func (e *GstDvbSrc) GetIsdbtLayerbSegmentCount() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerb-segment-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Post a message after timeout microseconds (0 = disabled)
// Default: 1000000, Min: 0, Max: 18446744073709551615
func (e *GstDvbSrc) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstDvbSrc) GetTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// (DVB-T) Hierarchy information
// Default: 1 (1)
func (e *GstDvbSrc) SetHierarchy(value GstDvbSrcHierarchy) error {
	e.Element.SetArg("hierarchy", string(value))
	return nil
}

func (e *GstDvbSrc) GetHierarchy() (GstDvbSrcHierarchy, error) {
	var value GstDvbSrcHierarchy
	var ok bool
	v, err := e.Element.GetProperty("hierarchy")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcHierarchy)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcHierarchy")
	}
	return value, nil
}

// LNB's Upper bound for low band reception (kHz)
// Default: 11700000, Min: 0, Max: -1
func (e *GstDvbSrc) SetLnbSlof(value uint) error {
	return e.Element.SetProperty("lnb-slof", value)
}

func (e *GstDvbSrc) GetLnbSlof() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("lnb-slof")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The number of reads before reporting frontend stats
// Default: 100, Min: 0, Max: -1
func (e *GstDvbSrc) SetStatsReportingInterval(value uint) error {
	return e.Element.SetProperty("stats-reporting-interval", value)
}

func (e *GstDvbSrc) GetStatsReportingInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("stats-reporting-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// (ISDB-T) SB Subchannel ID (-1 = AUTO)
// Default: -1, Min: -1, Max: 41
func (e *GstDvbSrc) SetIsdbtSbSubchannelId(value int) error {
	return e.Element.SetProperty("isdbt-sb-subchannel-id", value)
}

func (e *GstDvbSrc) GetIsdbtSbSubchannelId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-sb-subchannel-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// (DVB-T2 and DVB-S2 max 255, ISDB max 65535) Stream ID (-1 = disabled)
// Default: -1, Min: -1, Max: 65535
func (e *GstDvbSrc) SetStreamId(value int) error {
	return e.Element.SetProperty("stream-id", value)
}

func (e *GstDvbSrc) GetStreamId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("stream-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Atomically tune to channel. (For Apps)

func (e *GstDvbSrc) SetTune(value interface{}) error {
	return e.Element.SetProperty("tune", value)
}

// Channel bandwidth in Hz
// Default: 0, Min: 0, Max: -1
func (e *GstDvbSrc) SetBandwidthHz(value uint) error {
	return e.Element.SetProperty("bandwidth-hz", value)
}

func (e *GstDvbSrc) GetBandwidthHz() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("bandwidth-hz")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// (DVB-T) Low priority code rate
// Default: 1/2 (1)
func (e *GstDvbSrc) SetCodeRateLp(value GstDvbSrcCodeRate) error {
	return e.Element.SetProperty("code-rate-lp", value)
}

func (e *GstDvbSrc) GetCodeRateLp() (GstDvbSrcCodeRate, error) {
	var value GstDvbSrcCodeRate
	var ok bool
	v, err := e.Element.GetProperty("code-rate-lp")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcCodeRate)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcCodeRate")
	}
	return value, nil
}

// (ISDB-T) layer A Forward Error Correction
// Default: auto (9)
func (e *GstDvbSrc) SetIsdbtLayeraFec(value GstDvbSrcCodeRate) error {
	return e.Element.SetProperty("isdbt-layera-fec", value)
}

func (e *GstDvbSrc) GetIsdbtLayeraFec() (GstDvbSrcCodeRate, error) {
	var value GstDvbSrcCodeRate
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layera-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcCodeRate)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcCodeRate")
	}
	return value, nil
}

// (ISDB-T) SB segment count
// Default: 1, Min: 1, Max: 13
func (e *GstDvbSrc) SetIsdbtSbSegmentCount(value uint) error {
	return e.Element.SetProperty("isdbt-sb-segment-count", value)
}

func (e *GstDvbSrc) GetIsdbtSbSegmentCount() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("isdbt-sb-segment-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// (DVB-S/S2) Polarity [vhHV] (eg. V for Vertical)
// Default: H
func (e *GstDvbSrc) SetPolarity(value string) error {
	return e.Element.SetProperty("polarity", value)
}

func (e *GstDvbSrc) GetPolarity() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("polarity")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// (ISDB-T) layer B Forward Error Correction
// Default: auto (9)
func (e *GstDvbSrc) SetIsdbtLayerbFec(value GstDvbSrcCodeRate) error {
	return e.Element.SetProperty("isdbt-layerb-fec", value)
}

func (e *GstDvbSrc) GetIsdbtLayerbFec() (GstDvbSrcCodeRate, error) {
	var value GstDvbSrcCodeRate
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerb-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcCodeRate)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcCodeRate")
	}
	return value, nil
}

// (ISDB-T) Layer B modulation type
// Default: auto (6)
func (e *GstDvbSrc) SetIsdbtLayerbModulation(value GstDvbSrcModulation) error {
	e.Element.SetArg("isdbt-layerb-modulation", string(value))
	return nil
}

func (e *GstDvbSrc) GetIsdbtLayerbModulation() (GstDvbSrcModulation, error) {
	var value GstDvbSrcModulation
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerb-modulation")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDvbSrcModulation)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDvbSrcModulation")
	}
	return value, nil
}

// (ISDB-T) Layer C time interleaving (-1 = AUTO)
// Default: -1, Min: -1, Max: 8
func (e *GstDvbSrc) SetIsdbtLayercTimeInterleaving(value int) error {
	return e.Element.SetProperty("isdbt-layerc-time-interleaving", value)
}

func (e *GstDvbSrc) GetIsdbtLayercTimeInterleaving() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("isdbt-layerc-time-interleaving")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstDvbSrcBandwidth string

const (
	GstDvbSrcBandwidth8    GstDvbSrcBandwidth = "8"     // 8
	GstDvbSrcBandwidth7    GstDvbSrcBandwidth = "7"     // 7
	GstDvbSrcBandwidth6    GstDvbSrcBandwidth = "6"     // 6
	GstDvbSrcBandwidthAUTO GstDvbSrcBandwidth = "AUTO"  // AUTO
	GstDvbSrcBandwidth5    GstDvbSrcBandwidth = "5"     // 5
	GstDvbSrcBandwidth10   GstDvbSrcBandwidth = "10"    // 10
	GstDvbSrcBandwidth1712 GstDvbSrcBandwidth = "1.712" // 1.712
)

type GstDvbSrcCodeRate string

const (
	GstDvbSrcCodeRateNone GstDvbSrcCodeRate = "none" // NONE
	GstDvbSrcCodeRate12   GstDvbSrcCodeRate = "1/2"  // 1/2
	GstDvbSrcCodeRate23   GstDvbSrcCodeRate = "2/3"  // 2/3
	GstDvbSrcCodeRate34   GstDvbSrcCodeRate = "3/4"  // 3/4
	GstDvbSrcCodeRate45   GstDvbSrcCodeRate = "4/5"  // 4/5
	GstDvbSrcCodeRate56   GstDvbSrcCodeRate = "5/6"  // 5/6
	GstDvbSrcCodeRate67   GstDvbSrcCodeRate = "6/7"  // 6/7
	GstDvbSrcCodeRate78   GstDvbSrcCodeRate = "7/8"  // 7/8
	GstDvbSrcCodeRate89   GstDvbSrcCodeRate = "8/9"  // 8/9
	GstDvbSrcCodeRateAuto GstDvbSrcCodeRate = "auto" // AUTO
	GstDvbSrcCodeRate35   GstDvbSrcCodeRate = "3/5"  // 3/5
	GstDvbSrcCodeRate910  GstDvbSrcCodeRate = "9/10" // 9/10
	GstDvbSrcCodeRate25   GstDvbSrcCodeRate = "2/5"  // 2/5
)

type GstDvbSrcHierarchy string

const (
	GstDvbSrcHierarchyNone GstDvbSrcHierarchy = "none" // NONE
	GstDvbSrcHierarchy1    GstDvbSrcHierarchy = "1"    // 1
	GstDvbSrcHierarchy2    GstDvbSrcHierarchy = "2"    // 2
	GstDvbSrcHierarchy4    GstDvbSrcHierarchy = "4"    // 4
	GstDvbSrcHierarchyAuto GstDvbSrcHierarchy = "auto" // AUTO
)

type GstDvbSrcInterleaving string

const (
	GstDvbSrcInterleavingNone GstDvbSrcInterleaving = "none" // NONE
	GstDvbSrcInterleavingAuto GstDvbSrcInterleaving = "auto" // AUTO
	GstDvbSrcInterleaving240  GstDvbSrcInterleaving = "240"  // 240
	GstDvbSrcInterleaving720  GstDvbSrcInterleaving = "720"  // 720
)

type GstDvbSrcTransmissionMode string

const (
	GstDvbSrcTransmissionMode2k    GstDvbSrcTransmissionMode = "2k"    // 2K
	GstDvbSrcTransmissionMode8k    GstDvbSrcTransmissionMode = "8k"    // 8K
	GstDvbSrcTransmissionModeAuto  GstDvbSrcTransmissionMode = "auto"  // AUTO
	GstDvbSrcTransmissionMode4k    GstDvbSrcTransmissionMode = "4k"    // 4K
	GstDvbSrcTransmissionMode1k    GstDvbSrcTransmissionMode = "1k"    // 1K
	GstDvbSrcTransmissionMode16k   GstDvbSrcTransmissionMode = "16k"   // 16K
	GstDvbSrcTransmissionMode32k   GstDvbSrcTransmissionMode = "32k"   // 32K
	GstDvbSrcTransmissionModeC1    GstDvbSrcTransmissionMode = "c1"    // C1
	GstDvbSrcTransmissionModeC3780 GstDvbSrcTransmissionMode = "c3780" // C3780
)

type GstDvbSrcDelsys string

const (
	GstDvbSrcDelsysUndefined GstDvbSrcDelsys = "undefined" // UNDEFINED
	GstDvbSrcDelsysDvbCA     GstDvbSrcDelsys = "dvb-c-a"   // DVB-C-A
	GstDvbSrcDelsysDvbCB     GstDvbSrcDelsys = "dvb-c-b"   // DVB-C-B
	GstDvbSrcDelsysDvbT      GstDvbSrcDelsys = "dvb-t"     // DVB-T
	GstDvbSrcDelsysDss       GstDvbSrcDelsys = "dss"       // DSS
	GstDvbSrcDelsysDvbS      GstDvbSrcDelsys = "dvb-s"     // DVB-S
	GstDvbSrcDelsysDvbS2     GstDvbSrcDelsys = "dvb-s2"    // DVB-S2
	GstDvbSrcDelsysDvbH      GstDvbSrcDelsys = "dvb-h"     // DVB-H
	GstDvbSrcDelsysIsdbT     GstDvbSrcDelsys = "isdb-t"    // ISDB-T
	GstDvbSrcDelsysIsdbS     GstDvbSrcDelsys = "isdb-s"    // ISDB-S
	GstDvbSrcDelsysIsdbC     GstDvbSrcDelsys = "isdb-c"    // ISDB-C
	GstDvbSrcDelsysAtsc      GstDvbSrcDelsys = "atsc"      // ATSC
	GstDvbSrcDelsysAtscMh    GstDvbSrcDelsys = "atsc-mh"   // ATSC-MH
	GstDvbSrcDelsysDtmb      GstDvbSrcDelsys = "dtmb"      // DTMB
	GstDvbSrcDelsysCmmb      GstDvbSrcDelsys = "cmmb"      // CMMB
	GstDvbSrcDelsysDab       GstDvbSrcDelsys = "dab"       // DAB
	GstDvbSrcDelsysDvbT2     GstDvbSrcDelsys = "dvb-t2"    // DVB-T2
	GstDvbSrcDelsysTurbo     GstDvbSrcDelsys = "turbo"     // TURBO
	GstDvbSrcDelsysDvbCC     GstDvbSrcDelsys = "dvb-c-c"   // DVB-C-C
)

type GstDvbSrcGuard string

const (
	GstDvbSrcGuard32    GstDvbSrcGuard = "32"     // 32
	GstDvbSrcGuard16    GstDvbSrcGuard = "16"     // 16
	GstDvbSrcGuard8     GstDvbSrcGuard = "8"      // 8
	GstDvbSrcGuard4     GstDvbSrcGuard = "4"      // 4
	GstDvbSrcGuardAuto  GstDvbSrcGuard = "auto"   // AUTO
	GstDvbSrcGuard128   GstDvbSrcGuard = "128"    // 128
	GstDvbSrcGuard19128 GstDvbSrcGuard = "19/128" // 19/128
	GstDvbSrcGuard19256 GstDvbSrcGuard = "19/256" // 19/256
	GstDvbSrcGuardPn420 GstDvbSrcGuard = "pn420"  // PN420
	GstDvbSrcGuardPn595 GstDvbSrcGuard = "pn595"  // PN595
	GstDvbSrcGuardPn945 GstDvbSrcGuard = "pn945"  // PN945
)

type GstDvbSrcInversion string

const (
	GstDvbSrcInversionOff  GstDvbSrcInversion = "off"  // OFF
	GstDvbSrcInversionOn   GstDvbSrcInversion = "on"   // ON
	GstDvbSrcInversionAuto GstDvbSrcInversion = "auto" // AUTO
)

type GstDvbSrcModulation string

const (
	GstDvbSrcModulationQpsk   GstDvbSrcModulation = "qpsk"     // QPSK
	GstDvbSrcModulationQam16  GstDvbSrcModulation = "qam-16"   // QAM 16
	GstDvbSrcModulationQam32  GstDvbSrcModulation = "qam-32"   // QAM 32
	GstDvbSrcModulationQam64  GstDvbSrcModulation = "qam-64"   // QAM 64
	GstDvbSrcModulationQam128 GstDvbSrcModulation = "qam-128"  // QAM 128
	GstDvbSrcModulationQam256 GstDvbSrcModulation = "qam-256"  // QAM 256
	GstDvbSrcModulationAuto   GstDvbSrcModulation = "auto"     // AUTO
	GstDvbSrcModulation8vsb   GstDvbSrcModulation = "8vsb"     // 8VSB
	GstDvbSrcModulation16vsb  GstDvbSrcModulation = "16vsb"    // 16VSB
	GstDvbSrcModulation8psk   GstDvbSrcModulation = "8psk"     // 8PSK
	GstDvbSrcModulation16apsk GstDvbSrcModulation = "16apsk"   // 16APSK
	GstDvbSrcModulation32apsk GstDvbSrcModulation = "32apsk"   // 32APSK
	GstDvbSrcModulationDqpsk  GstDvbSrcModulation = "dqpsk"    // DQPSK
	GstDvbSrcModulationQam4Nr GstDvbSrcModulation = "qam-4-nr" // QAM 4 NR
)

type GstDvbSrcPilot string

const (
	GstDvbSrcPilotOn   GstDvbSrcPilot = "on"   // ON
	GstDvbSrcPilotOff  GstDvbSrcPilot = "off"  // OFF
	GstDvbSrcPilotAuto GstDvbSrcPilot = "auto" // AUTO
)

type GstDvbSrcRolloff string

const (
	GstDvbSrcRolloff35   GstDvbSrcRolloff = "35"   // 35
	GstDvbSrcRolloff20   GstDvbSrcRolloff = "20"   // 20
	GstDvbSrcRolloff25   GstDvbSrcRolloff = "25"   // 25
	GstDvbSrcRolloffAuto GstDvbSrcRolloff = "auto" // auto
)
