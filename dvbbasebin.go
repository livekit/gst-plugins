// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gstplugins

import (
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"
)

type Dvbbasebin struct {
	Element *gst.Element
}

func NewDvbbasebin() (*Dvbbasebin, error) {
	e, err := gst.NewElement("dvbbasebin")
	if err != nil {
		return nil, err
	}
	return &Dvbbasebin{Element: e}, nil
}

func NewDvbbasebinWithName(name string) (*Dvbbasebin, error) {
	e, err := gst.NewElementWithName("dvbbasebin", name)
	if err != nil {
		return nil, err
	}
	return &Dvbbasebin{Element: e}, nil
}

// ----- Properties -----

// adapter (int)
//
// The DVB adapter device number (eg. 0 for adapter0)

func (e *Dvbbasebin) GetAdapter() (int, error) {
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

func (e *Dvbbasebin) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

// bandwidth (GstDvbSrcBandwidth)
//
// (DVB-T) Bandwidth. Deprecated

func (e *Dvbbasebin) GetBandwidth() (interface{}, error) {
	return e.Element.GetProperty("bandwidth")
}

func (e *Dvbbasebin) SetBandwidth(value interface{}) error {
	return e.Element.SetProperty("bandwidth", value)
}

// bandwidth-hz (uint)
//
// Channel bandwidth in Hz

func (e *Dvbbasebin) GetBandwidthHz() (uint, error) {
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

func (e *Dvbbasebin) SetBandwidthHz(value uint) error {
	return e.Element.SetProperty("bandwidth-hz", value)
}

// code-rate-hp (GstDvbSrcCode_Rate)
//
// (DVB-T, DVB-S/S2 and DVB-C) High priority code rate

func (e *Dvbbasebin) GetCodeRateHp() (interface{}, error) {
	return e.Element.GetProperty("code-rate-hp")
}

func (e *Dvbbasebin) SetCodeRateHp(value interface{}) error {
	return e.Element.SetProperty("code-rate-hp", value)
}

// code-rate-lp (GstDvbSrcCode_Rate)
//
// (DVB-T) Low priority code rate

func (e *Dvbbasebin) GetCodeRateLp() (interface{}, error) {
	return e.Element.GetProperty("code-rate-lp")
}

func (e *Dvbbasebin) SetCodeRateLp(value interface{}) error {
	return e.Element.SetProperty("code-rate-lp", value)
}

// delsys (GstDvbSrcDelsys)
//
// Delivery System

func (e *Dvbbasebin) GetDelsys() (interface{}, error) {
	return e.Element.GetProperty("delsys")
}

func (e *Dvbbasebin) SetDelsys(value interface{}) error {
	return e.Element.SetProperty("delsys", value)
}

// diseqc-source (int)
//
// (DVB-S/S2) Selected DiSEqC source. Only needed if you have a DiSEqC switch. Otherwise leave at -1 (disabled)

func (e *Dvbbasebin) GetDiseqcSource() (int, error) {
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

func (e *Dvbbasebin) SetDiseqcSource(value int) error {
	return e.Element.SetProperty("diseqc-source", value)
}

// frequency (uint)
//
// Center frequency to tune into. Measured in kHz for the satellite distribution standards and Hz for all the rest

func (e *Dvbbasebin) GetFrequency() (uint, error) {
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

func (e *Dvbbasebin) SetFrequency(value uint) error {
	return e.Element.SetProperty("frequency", value)
}

// frontend (int)
//
// The frontend device number (eg. 0 for frontend0)

func (e *Dvbbasebin) GetFrontend() (int, error) {
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

func (e *Dvbbasebin) SetFrontend(value int) error {
	return e.Element.SetProperty("frontend", value)
}

// guard (GstDvbSrcGuard)
//
// (DVB-T) Guard Interval

func (e *Dvbbasebin) GetGuard() (interface{}, error) {
	return e.Element.GetProperty("guard")
}

func (e *Dvbbasebin) SetGuard(value interface{}) error {
	return e.Element.SetProperty("guard", value)
}

// hierarchy (GstDvbSrcHierarchy)
//
// (DVB-T) Hierarchy information

func (e *Dvbbasebin) GetHierarchy() (interface{}, error) {
	return e.Element.GetProperty("hierarchy")
}

func (e *Dvbbasebin) SetHierarchy(value interface{}) error {
	return e.Element.SetProperty("hierarchy", value)
}

// interleaving (GstDvbSrcInterleaving)
//
// (DTMB) Interleaving type

func (e *Dvbbasebin) GetInterleaving() (interface{}, error) {
	return e.Element.GetProperty("interleaving")
}

func (e *Dvbbasebin) SetInterleaving(value interface{}) error {
	return e.Element.SetProperty("interleaving", value)
}

// inversion (GstDvbSrcInversion)
//
// (DVB-T and DVB-C) Inversion information

func (e *Dvbbasebin) GetInversion() (interface{}, error) {
	return e.Element.GetProperty("inversion")
}

func (e *Dvbbasebin) SetInversion(value interface{}) error {
	return e.Element.SetProperty("inversion", value)
}

// isdbt-layer-enabled (uint)
//
// (ISDB-T) Layer Enabled (7 = All layers)

func (e *Dvbbasebin) GetIsdbtLayerEnabled() (uint, error) {
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

func (e *Dvbbasebin) SetIsdbtLayerEnabled(value uint) error {
	return e.Element.SetProperty("isdbt-layer-enabled", value)
}

// isdbt-layera-fec (GstDvbSrcCode_Rate)
//
// (ISDB-T) layer A Forward Error Correction

func (e *Dvbbasebin) GetIsdbtLayeraFec() (interface{}, error) {
	return e.Element.GetProperty("isdbt-layera-fec")
}

func (e *Dvbbasebin) SetIsdbtLayeraFec(value interface{}) error {
	return e.Element.SetProperty("isdbt-layera-fec", value)
}

// isdbt-layera-modulation (GstDvbSrcModulation)
//
// (ISDB-T) Layer A modulation type

func (e *Dvbbasebin) GetIsdbtLayeraModulation() (interface{}, error) {
	return e.Element.GetProperty("isdbt-layera-modulation")
}

func (e *Dvbbasebin) SetIsdbtLayeraModulation(value interface{}) error {
	return e.Element.SetProperty("isdbt-layera-modulation", value)
}

// isdbt-layera-segment-count (int)
//
// (ISDB-T) Layer A segment count (-1 = AUTO)

func (e *Dvbbasebin) GetIsdbtLayeraSegmentCount() (int, error) {
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

func (e *Dvbbasebin) SetIsdbtLayeraSegmentCount(value int) error {
	return e.Element.SetProperty("isdbt-layera-segment-count", value)
}

// isdbt-layera-time-interleaving (int)
//
// (ISDB-T) Layer A time interleaving (-1 = AUTO)

func (e *Dvbbasebin) GetIsdbtLayeraTimeInterleaving() (int, error) {
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

func (e *Dvbbasebin) SetIsdbtLayeraTimeInterleaving(value int) error {
	return e.Element.SetProperty("isdbt-layera-time-interleaving", value)
}

// isdbt-layerb-fec (GstDvbSrcCode_Rate)
//
// (ISDB-T) layer B Forward Error Correction

func (e *Dvbbasebin) GetIsdbtLayerbFec() (interface{}, error) {
	return e.Element.GetProperty("isdbt-layerb-fec")
}

func (e *Dvbbasebin) SetIsdbtLayerbFec(value interface{}) error {
	return e.Element.SetProperty("isdbt-layerb-fec", value)
}

// isdbt-layerb-modulation (GstDvbSrcModulation)
//
// (ISDB-T) Layer B modulation type

func (e *Dvbbasebin) GetIsdbtLayerbModulation() (interface{}, error) {
	return e.Element.GetProperty("isdbt-layerb-modulation")
}

func (e *Dvbbasebin) SetIsdbtLayerbModulation(value interface{}) error {
	return e.Element.SetProperty("isdbt-layerb-modulation", value)
}

// isdbt-layerb-segment-count (int)
//
// (ISDB-T) Layer B segment count (-1 = AUTO)

func (e *Dvbbasebin) GetIsdbtLayerbSegmentCount() (int, error) {
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

func (e *Dvbbasebin) SetIsdbtLayerbSegmentCount(value int) error {
	return e.Element.SetProperty("isdbt-layerb-segment-count", value)
}

// isdbt-layerb-time-interleaving (int)
//
// (ISDB-T) Layer B time interleaving (-1 = AUTO)

func (e *Dvbbasebin) GetIsdbtLayerbTimeInterleaving() (int, error) {
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

func (e *Dvbbasebin) SetIsdbtLayerbTimeInterleaving(value int) error {
	return e.Element.SetProperty("isdbt-layerb-time-interleaving", value)
}

// isdbt-layerc-fec (GstDvbSrcCode_Rate)
//
// (ISDB-T) layer C Forward Error Correction

func (e *Dvbbasebin) GetIsdbtLayercFec() (interface{}, error) {
	return e.Element.GetProperty("isdbt-layerc-fec")
}

func (e *Dvbbasebin) SetIsdbtLayercFec(value interface{}) error {
	return e.Element.SetProperty("isdbt-layerc-fec", value)
}

// isdbt-layerc-modulation (GstDvbSrcModulation)
//
// (ISDB-T) Layer C modulation type

func (e *Dvbbasebin) GetIsdbtLayercModulation() (interface{}, error) {
	return e.Element.GetProperty("isdbt-layerc-modulation")
}

func (e *Dvbbasebin) SetIsdbtLayercModulation(value interface{}) error {
	return e.Element.SetProperty("isdbt-layerc-modulation", value)
}

// isdbt-layerc-segment-count (int)
//
// (ISDB-T) Layer C segment count (-1 = AUTO)

func (e *Dvbbasebin) GetIsdbtLayercSegmentCount() (int, error) {
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

func (e *Dvbbasebin) SetIsdbtLayercSegmentCount(value int) error {
	return e.Element.SetProperty("isdbt-layerc-segment-count", value)
}

// isdbt-layerc-time-interleaving (int)
//
// (ISDB-T) Layer C time interleaving (-1 = AUTO)

func (e *Dvbbasebin) GetIsdbtLayercTimeInterleaving() (int, error) {
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

func (e *Dvbbasebin) SetIsdbtLayercTimeInterleaving(value int) error {
	return e.Element.SetProperty("isdbt-layerc-time-interleaving", value)
}

// isdbt-partial-reception (int)
//
// (ISDB-T) Partial Reception (-1 = AUTO)

func (e *Dvbbasebin) GetIsdbtPartialReception() (int, error) {
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

func (e *Dvbbasebin) SetIsdbtPartialReception(value int) error {
	return e.Element.SetProperty("isdbt-partial-reception", value)
}

// isdbt-sb-segment-count (uint)
//
// (ISDB-T) SB segment count

func (e *Dvbbasebin) GetIsdbtSbSegmentCount() (uint, error) {
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

func (e *Dvbbasebin) SetIsdbtSbSegmentCount(value uint) error {
	return e.Element.SetProperty("isdbt-sb-segment-count", value)
}

// isdbt-sb-segment-idx (int)
//
// (ISDB-T) SB segment IDX

func (e *Dvbbasebin) GetIsdbtSbSegmentIdx() (int, error) {
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

func (e *Dvbbasebin) SetIsdbtSbSegmentIdx(value int) error {
	return e.Element.SetProperty("isdbt-sb-segment-idx", value)
}

// isdbt-sb-subchannel-id (int)
//
// (ISDB-T) SB Subchannel ID (-1 = AUTO)

func (e *Dvbbasebin) GetIsdbtSbSubchannelId() (int, error) {
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

func (e *Dvbbasebin) SetIsdbtSbSubchannelId(value int) error {
	return e.Element.SetProperty("isdbt-sb-subchannel-id", value)
}

// isdbt-sound-broadcasting (int)
//
// (ISDB-T) Sound Broadcasting

func (e *Dvbbasebin) GetIsdbtSoundBroadcasting() (int, error) {
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

func (e *Dvbbasebin) SetIsdbtSoundBroadcasting(value int) error {
	return e.Element.SetProperty("isdbt-sound-broadcasting", value)
}

// lnb-lof1 (uint)
//
// LNB's Local oscillator frequency used for low band reception (kHz)

func (e *Dvbbasebin) GetLnbLof1() (uint, error) {
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

func (e *Dvbbasebin) SetLnbLof1(value uint) error {
	return e.Element.SetProperty("lnb-lof1", value)
}

// lnb-lof2 (uint)
//
// LNB's Local oscillator frequency used for high band reception (kHz)

func (e *Dvbbasebin) GetLnbLof2() (uint, error) {
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

func (e *Dvbbasebin) SetLnbLof2(value uint) error {
	return e.Element.SetProperty("lnb-lof2", value)
}

// lnb-slof (uint)
//
// LNB's Upper bound for low band reception (kHz)

func (e *Dvbbasebin) GetLnbSlof() (uint, error) {
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

func (e *Dvbbasebin) SetLnbSlof(value uint) error {
	return e.Element.SetProperty("lnb-slof", value)
}

// modulation (GstDvbSrcModulation)
//
// (DVB-T/T2/C/S2, TURBO and ATSC) Modulation type

func (e *Dvbbasebin) GetModulation() (interface{}, error) {
	return e.Element.GetProperty("modulation")
}

func (e *Dvbbasebin) SetModulation(value interface{}) error {
	return e.Element.SetProperty("modulation", value)
}

// pilot (GstDvbSrcPilot)
//
// Pilot (DVB-S2)

func (e *Dvbbasebin) GetPilot() (interface{}, error) {
	return e.Element.GetProperty("pilot")
}

func (e *Dvbbasebin) SetPilot(value interface{}) error {
	return e.Element.SetProperty("pilot", value)
}

// polarity (string)
//
// (DVB-S/S2) Polarity [vhHV] (eg. V for Vertical)

func (e *Dvbbasebin) GetPolarity() (string, error) {
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

func (e *Dvbbasebin) SetPolarity(value string) error {
	return e.Element.SetProperty("polarity", value)
}

// program-numbers (string)
//
// Colon separated list of programs

func (e *Dvbbasebin) GetProgramNumbers() (string, error) {
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

func (e *Dvbbasebin) SetProgramNumbers(value string) error {
	return e.Element.SetProperty("program-numbers", value)
}

// rolloff (GstDvbSrcRolloff)
//
// Rolloff (DVB-S2)

func (e *Dvbbasebin) GetRolloff() (interface{}, error) {
	return e.Element.GetProperty("rolloff")
}

func (e *Dvbbasebin) SetRolloff(value interface{}) error {
	return e.Element.SetProperty("rolloff", value)
}

// stats-reporting-interval (uint)
//
// The number of reads before reporting frontend stats

func (e *Dvbbasebin) GetStatsReportingInterval() (uint, error) {
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

func (e *Dvbbasebin) SetStatsReportingInterval(value uint) error {
	return e.Element.SetProperty("stats-reporting-interval", value)
}

// stream-id (int)
//
// (DVB-T2 and DVB-S2 max 255, ISDB max 65535) Stream ID (-1 = disabled)

func (e *Dvbbasebin) GetStreamId() (int, error) {
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

func (e *Dvbbasebin) SetStreamId(value int) error {
	return e.Element.SetProperty("stream-id", value)
}

// symbol-rate (uint)
//
// (DVB-S/S2, DVB-C) Symbol rate in kBd (kilo bauds)

func (e *Dvbbasebin) GetSymbolRate() (uint, error) {
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

func (e *Dvbbasebin) SetSymbolRate(value uint) error {
	return e.Element.SetProperty("symbol-rate", value)
}

// trans-mode (GstDvbSrcTransmission_Mode)
//
// (DVB-T) Transmission mode

func (e *Dvbbasebin) GetTransMode() (interface{}, error) {
	return e.Element.GetProperty("trans-mode")
}

func (e *Dvbbasebin) SetTransMode(value interface{}) error {
	return e.Element.SetProperty("trans-mode", value)
}

// tuning-timeout (uint64)
//
// Microseconds to wait before giving up tuning/locking on a signal

func (e *Dvbbasebin) GetTuningTimeout() (uint64, error) {
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

func (e *Dvbbasebin) SetTuningTimeout(value uint64) error {
	return e.Element.SetProperty("tuning-timeout", value)
}

