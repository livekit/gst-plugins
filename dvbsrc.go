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

	"github.com/livekit/gstplugins/base"
)

type Dvbsrc struct {
	*base.GstPushSrc
}

func NewDvbsrc() (*Dvbsrc, error) {
	e, err := gst.NewElement("dvbsrc")
	if err != nil {
		return nil, err
	}
	return &Dvbsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewDvbsrcWithName(name string) (*Dvbsrc, error) {
	e, err := gst.NewElementWithName("dvbsrc", name)
	if err != nil {
		return nil, err
	}
	return &Dvbsrc{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// adapter (int)
//
// The DVB adapter device number (eg. 0 for adapter0)

func (e *Dvbsrc) GetAdapter() (int, error) {
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

func (e *Dvbsrc) SetAdapter(value int) error {
	return e.Element.SetProperty("adapter", value)
}

// bandwidth (GstDvbSrcBandwidth)
//
// (DVB-T) Bandwidth. Deprecated

func (e *Dvbsrc) GetBandwidth() (GstDvbSrcBandwidth, error) {
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

func (e *Dvbsrc) SetBandwidth(value GstDvbSrcBandwidth) error {
	e.Element.SetArg("bandwidth", string(value))
	return nil
}

// bandwidth-hz (uint)
//
// Channel bandwidth in Hz

func (e *Dvbsrc) GetBandwidthHz() (uint, error) {
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

func (e *Dvbsrc) SetBandwidthHz(value uint) error {
	return e.Element.SetProperty("bandwidth-hz", value)
}

// code-rate-hp (GstDvbSrcCode_Rate)
//
// (DVB-T, DVB-S/S2 and DVB-C) High priority code rate

func (e *Dvbsrc) GetCodeRateHp() (interface{}, error) {
	return e.Element.GetProperty("code-rate-hp")
}

func (e *Dvbsrc) SetCodeRateHp(value interface{}) error {
	return e.Element.SetProperty("code-rate-hp", value)
}

// code-rate-lp (GstDvbSrcCode_Rate)
//
// (DVB-T) Low priority code rate

func (e *Dvbsrc) GetCodeRateLp() (interface{}, error) {
	return e.Element.GetProperty("code-rate-lp")
}

func (e *Dvbsrc) SetCodeRateLp(value interface{}) error {
	return e.Element.SetProperty("code-rate-lp", value)
}

// delsys (GstDvbSrcDelsys)
//
// Delivery System

func (e *Dvbsrc) GetDelsys() (GstDvbSrcDelsys, error) {
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

func (e *Dvbsrc) SetDelsys(value GstDvbSrcDelsys) error {
	e.Element.SetArg("delsys", string(value))
	return nil
}

// diseqc-source (int)
//
// (DVB-S/S2) Selected DiSEqC source. Only needed if you have a DiSEqC switch. Otherwise leave at -1 (disabled)

func (e *Dvbsrc) GetDiseqcSource() (int, error) {
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

func (e *Dvbsrc) SetDiseqcSource(value int) error {
	return e.Element.SetProperty("diseqc-source", value)
}

// dvb-buffer-size (uint)
//
// The kernel buffer size used by the DVB api

func (e *Dvbsrc) GetDvbBufferSize() (uint, error) {
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

func (e *Dvbsrc) SetDvbBufferSize(value uint) error {
	return e.Element.SetProperty("dvb-buffer-size", value)
}

// frequency (uint)
//
// Center frequency to tune into. Measured in kHz for the satellite distribution standards and Hz for all the rest

func (e *Dvbsrc) GetFrequency() (uint, error) {
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

func (e *Dvbsrc) SetFrequency(value uint) error {
	return e.Element.SetProperty("frequency", value)
}

// frontend (int)
//
// The frontend device number (eg. 0 for frontend0)

func (e *Dvbsrc) GetFrontend() (int, error) {
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

func (e *Dvbsrc) SetFrontend(value int) error {
	return e.Element.SetProperty("frontend", value)
}

// guard (GstDvbSrcGuard)
//
// (DVB-T) Guard Interval

func (e *Dvbsrc) GetGuard() (GstDvbSrcGuard, error) {
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

func (e *Dvbsrc) SetGuard(value GstDvbSrcGuard) error {
	e.Element.SetArg("guard", string(value))
	return nil
}

// hierarchy (GstDvbSrcHierarchy)
//
// (DVB-T) Hierarchy information

func (e *Dvbsrc) GetHierarchy() (GstDvbSrcHierarchy, error) {
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

func (e *Dvbsrc) SetHierarchy(value GstDvbSrcHierarchy) error {
	e.Element.SetArg("hierarchy", string(value))
	return nil
}

// interleaving (GstDvbSrcInterleaving)
//
// (DTMB) Interleaving type

func (e *Dvbsrc) GetInterleaving() (GstDvbSrcInterleaving, error) {
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

func (e *Dvbsrc) SetInterleaving(value GstDvbSrcInterleaving) error {
	e.Element.SetArg("interleaving", string(value))
	return nil
}

// inversion (GstDvbSrcInversion)
//
// (DVB-T and DVB-C) Inversion information

func (e *Dvbsrc) GetInversion() (GstDvbSrcInversion, error) {
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

func (e *Dvbsrc) SetInversion(value GstDvbSrcInversion) error {
	e.Element.SetArg("inversion", string(value))
	return nil
}

// isdbt-layer-enabled (uint)
//
// (ISDB-T) Layer Enabled (7 = All layers)

func (e *Dvbsrc) GetIsdbtLayerEnabled() (uint, error) {
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

func (e *Dvbsrc) SetIsdbtLayerEnabled(value uint) error {
	return e.Element.SetProperty("isdbt-layer-enabled", value)
}

// isdbt-layera-fec (GstDvbSrcCode_Rate)
//
// (ISDB-T) layer A Forward Error Correction

func (e *Dvbsrc) GetIsdbtLayeraFec() (interface{}, error) {
	return e.Element.GetProperty("isdbt-layera-fec")
}

func (e *Dvbsrc) SetIsdbtLayeraFec(value interface{}) error {
	return e.Element.SetProperty("isdbt-layera-fec", value)
}

// isdbt-layera-modulation (GstDvbSrcModulation)
//
// (ISDB-T) Layer A modulation type

func (e *Dvbsrc) GetIsdbtLayeraModulation() (GstDvbSrcModulation, error) {
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

func (e *Dvbsrc) SetIsdbtLayeraModulation(value GstDvbSrcModulation) error {
	e.Element.SetArg("isdbt-layera-modulation", string(value))
	return nil
}

// isdbt-layera-segment-count (int)
//
// (ISDB-T) Layer A segment count (-1 = AUTO)

func (e *Dvbsrc) GetIsdbtLayeraSegmentCount() (int, error) {
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

func (e *Dvbsrc) SetIsdbtLayeraSegmentCount(value int) error {
	return e.Element.SetProperty("isdbt-layera-segment-count", value)
}

// isdbt-layera-time-interleaving (int)
//
// (ISDB-T) Layer A time interleaving (-1 = AUTO)

func (e *Dvbsrc) GetIsdbtLayeraTimeInterleaving() (int, error) {
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

func (e *Dvbsrc) SetIsdbtLayeraTimeInterleaving(value int) error {
	return e.Element.SetProperty("isdbt-layera-time-interleaving", value)
}

// isdbt-layerb-fec (GstDvbSrcCode_Rate)
//
// (ISDB-T) layer B Forward Error Correction

func (e *Dvbsrc) GetIsdbtLayerbFec() (interface{}, error) {
	return e.Element.GetProperty("isdbt-layerb-fec")
}

func (e *Dvbsrc) SetIsdbtLayerbFec(value interface{}) error {
	return e.Element.SetProperty("isdbt-layerb-fec", value)
}

// isdbt-layerb-modulation (GstDvbSrcModulation)
//
// (ISDB-T) Layer B modulation type

func (e *Dvbsrc) GetIsdbtLayerbModulation() (GstDvbSrcModulation, error) {
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

func (e *Dvbsrc) SetIsdbtLayerbModulation(value GstDvbSrcModulation) error {
	e.Element.SetArg("isdbt-layerb-modulation", string(value))
	return nil
}

// isdbt-layerb-segment-count (int)
//
// (ISDB-T) Layer B segment count (-1 = AUTO)

func (e *Dvbsrc) GetIsdbtLayerbSegmentCount() (int, error) {
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

func (e *Dvbsrc) SetIsdbtLayerbSegmentCount(value int) error {
	return e.Element.SetProperty("isdbt-layerb-segment-count", value)
}

// isdbt-layerb-time-interleaving (int)
//
// (ISDB-T) Layer B time interleaving (-1 = AUTO)

func (e *Dvbsrc) GetIsdbtLayerbTimeInterleaving() (int, error) {
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

func (e *Dvbsrc) SetIsdbtLayerbTimeInterleaving(value int) error {
	return e.Element.SetProperty("isdbt-layerb-time-interleaving", value)
}

// isdbt-layerc-fec (GstDvbSrcCode_Rate)
//
// (ISDB-T) layer C Forward Error Correction

func (e *Dvbsrc) GetIsdbtLayercFec() (interface{}, error) {
	return e.Element.GetProperty("isdbt-layerc-fec")
}

func (e *Dvbsrc) SetIsdbtLayercFec(value interface{}) error {
	return e.Element.SetProperty("isdbt-layerc-fec", value)
}

// isdbt-layerc-modulation (GstDvbSrcModulation)
//
// (ISDB-T) Layer C modulation type

func (e *Dvbsrc) GetIsdbtLayercModulation() (GstDvbSrcModulation, error) {
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

func (e *Dvbsrc) SetIsdbtLayercModulation(value GstDvbSrcModulation) error {
	e.Element.SetArg("isdbt-layerc-modulation", string(value))
	return nil
}

// isdbt-layerc-segment-count (int)
//
// (ISDB-T) Layer C segment count (-1 = AUTO)

func (e *Dvbsrc) GetIsdbtLayercSegmentCount() (int, error) {
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

func (e *Dvbsrc) SetIsdbtLayercSegmentCount(value int) error {
	return e.Element.SetProperty("isdbt-layerc-segment-count", value)
}

// isdbt-layerc-time-interleaving (int)
//
// (ISDB-T) Layer C time interleaving (-1 = AUTO)

func (e *Dvbsrc) GetIsdbtLayercTimeInterleaving() (int, error) {
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

func (e *Dvbsrc) SetIsdbtLayercTimeInterleaving(value int) error {
	return e.Element.SetProperty("isdbt-layerc-time-interleaving", value)
}

// isdbt-partial-reception (int)
//
// (ISDB-T) Partial Reception (-1 = AUTO)

func (e *Dvbsrc) GetIsdbtPartialReception() (int, error) {
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

func (e *Dvbsrc) SetIsdbtPartialReception(value int) error {
	return e.Element.SetProperty("isdbt-partial-reception", value)
}

// isdbt-sb-segment-count (uint)
//
// (ISDB-T) SB segment count

func (e *Dvbsrc) GetIsdbtSbSegmentCount() (uint, error) {
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

func (e *Dvbsrc) SetIsdbtSbSegmentCount(value uint) error {
	return e.Element.SetProperty("isdbt-sb-segment-count", value)
}

// isdbt-sb-segment-idx (int)
//
// (ISDB-T) SB segment IDX

func (e *Dvbsrc) GetIsdbtSbSegmentIdx() (int, error) {
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

func (e *Dvbsrc) SetIsdbtSbSegmentIdx(value int) error {
	return e.Element.SetProperty("isdbt-sb-segment-idx", value)
}

// isdbt-sb-subchannel-id (int)
//
// (ISDB-T) SB Subchannel ID (-1 = AUTO)

func (e *Dvbsrc) GetIsdbtSbSubchannelId() (int, error) {
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

func (e *Dvbsrc) SetIsdbtSbSubchannelId(value int) error {
	return e.Element.SetProperty("isdbt-sb-subchannel-id", value)
}

// isdbt-sound-broadcasting (int)
//
// (ISDB-T) Sound Broadcasting

func (e *Dvbsrc) GetIsdbtSoundBroadcasting() (int, error) {
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

func (e *Dvbsrc) SetIsdbtSoundBroadcasting(value int) error {
	return e.Element.SetProperty("isdbt-sound-broadcasting", value)
}

// lnb-lof1 (uint)
//
// LNB's Local oscillator frequency used for low band reception (kHz)

func (e *Dvbsrc) GetLnbLof1() (uint, error) {
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

func (e *Dvbsrc) SetLnbLof1(value uint) error {
	return e.Element.SetProperty("lnb-lof1", value)
}

// lnb-lof2 (uint)
//
// LNB's Local oscillator frequency used for high band reception (kHz)

func (e *Dvbsrc) GetLnbLof2() (uint, error) {
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

func (e *Dvbsrc) SetLnbLof2(value uint) error {
	return e.Element.SetProperty("lnb-lof2", value)
}

// lnb-slof (uint)
//
// LNB's Upper bound for low band reception (kHz)

func (e *Dvbsrc) GetLnbSlof() (uint, error) {
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

func (e *Dvbsrc) SetLnbSlof(value uint) error {
	return e.Element.SetProperty("lnb-slof", value)
}

// modulation (GstDvbSrcModulation)
//
// (DVB-T/T2/C/S2, TURBO and ATSC) Modulation type

func (e *Dvbsrc) GetModulation() (GstDvbSrcModulation, error) {
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

func (e *Dvbsrc) SetModulation(value GstDvbSrcModulation) error {
	e.Element.SetArg("modulation", string(value))
	return nil
}

// pids (string)
//
// Colon-separated list of PIDs (eg. 110:120) to capture. ACT and CAT are automatically included but PMT should be added explicitly. Special value 8192 gets full MPEG-TS

func (e *Dvbsrc) GetPids() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pids")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Dvbsrc) SetPids(value string) error {
	return e.Element.SetProperty("pids", value)
}

// pilot (GstDvbSrcPilot)
//
// Pilot (DVB-S2)

func (e *Dvbsrc) GetPilot() (GstDvbSrcPilot, error) {
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

func (e *Dvbsrc) SetPilot(value GstDvbSrcPilot) error {
	e.Element.SetArg("pilot", string(value))
	return nil
}

// polarity (string)
//
// (DVB-S/S2) Polarity [vhHV] (eg. V for Vertical)

func (e *Dvbsrc) GetPolarity() (string, error) {
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

func (e *Dvbsrc) SetPolarity(value string) error {
	return e.Element.SetProperty("polarity", value)
}

// rolloff (GstDvbSrcRolloff)
//
// Rolloff (DVB-S2)

func (e *Dvbsrc) GetRolloff() (GstDvbSrcRolloff, error) {
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

func (e *Dvbsrc) SetRolloff(value GstDvbSrcRolloff) error {
	e.Element.SetArg("rolloff", string(value))
	return nil
}

// stats-reporting-interval (uint)
//
// The number of reads before reporting frontend stats

func (e *Dvbsrc) GetStatsReportingInterval() (uint, error) {
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

func (e *Dvbsrc) SetStatsReportingInterval(value uint) error {
	return e.Element.SetProperty("stats-reporting-interval", value)
}

// stream-id (int)
//
// (DVB-T2 and DVB-S2 max 255, ISDB max 65535) Stream ID (-1 = disabled)

func (e *Dvbsrc) GetStreamId() (int, error) {
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

func (e *Dvbsrc) SetStreamId(value int) error {
	return e.Element.SetProperty("stream-id", value)
}

// symbol-rate (uint)
//
// (DVB-S/S2, DVB-C) Symbol rate in kBd (kilo bauds)

func (e *Dvbsrc) GetSymbolRate() (uint, error) {
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

func (e *Dvbsrc) SetSymbolRate(value uint) error {
	return e.Element.SetProperty("symbol-rate", value)
}

// timeout (uint64)
//
// Post a message after timeout microseconds (0 = disabled)

func (e *Dvbsrc) GetTimeout() (uint64, error) {
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

func (e *Dvbsrc) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

// trans-mode (GstDvbSrcTransmission_Mode)
//
// (DVB-T) Transmission mode

func (e *Dvbsrc) GetTransMode() (interface{}, error) {
	return e.Element.GetProperty("trans-mode")
}

func (e *Dvbsrc) SetTransMode(value interface{}) error {
	return e.Element.SetProperty("trans-mode", value)
}

// tune (GstGpointer)
//
// Atomically tune to channel. (For Apps)

func (e *Dvbsrc) GetTune() (interface{}, error) {
	return e.Element.GetProperty("tune")
}

func (e *Dvbsrc) SetTune(value interface{}) error {
	return e.Element.SetProperty("tune", value)
}

// tuning-timeout (uint64)
//
// Microseconds to wait before giving up tuning/locking on a signal

func (e *Dvbsrc) GetTuningTimeout() (uint64, error) {
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

func (e *Dvbsrc) SetTuningTimeout(value uint64) error {
	return e.Element.SetProperty("tuning-timeout", value)
}

// ----- Constants -----

type GstDvbSrcPilot string

const (
	GstDvbSrcPilotOn GstDvbSrcPilot = "on" // on (0) – ON
	GstDvbSrcPilotOff GstDvbSrcPilot = "off" // off (1) – OFF
	GstDvbSrcPilotAuto GstDvbSrcPilot = "auto" // auto (2) – AUTO
)

type GstDvbSrcRolloff string

const (
	GstDvbSrcRolloff35 GstDvbSrcRolloff = "35" // 35 (0) – 35
	GstDvbSrcRolloff20 GstDvbSrcRolloff = "20" // 20 (1) – 20
	GstDvbSrcRolloff25 GstDvbSrcRolloff = "25" // 25 (2) – 25
	GstDvbSrcRolloffAuto GstDvbSrcRolloff = "auto" // auto (3) – auto
)

type GstDvbSrcCodeRate string

const (
	GstDvbSrcCodeRateNone GstDvbSrcCodeRate = "none" // none (0) – NONE
	GstDvbSrcCodeRate12 GstDvbSrcCodeRate = "1/2" // 1/2 (1) – 1/2
	GstDvbSrcCodeRate23 GstDvbSrcCodeRate = "2/3" // 2/3 (2) – 2/3
	GstDvbSrcCodeRate34 GstDvbSrcCodeRate = "3/4" // 3/4 (3) – 3/4
	GstDvbSrcCodeRate45 GstDvbSrcCodeRate = "4/5" // 4/5 (4) – 4/5
	GstDvbSrcCodeRate56 GstDvbSrcCodeRate = "5/6" // 5/6 (5) – 5/6
	GstDvbSrcCodeRate67 GstDvbSrcCodeRate = "6/7" // 6/7 (6) – 6/7
	GstDvbSrcCodeRate78 GstDvbSrcCodeRate = "7/8" // 7/8 (7) – 7/8
	GstDvbSrcCodeRate89 GstDvbSrcCodeRate = "8/9" // 8/9 (8) – 8/9
	GstDvbSrcCodeRateAuto GstDvbSrcCodeRate = "auto" // auto (9) – AUTO
	GstDvbSrcCodeRate35 GstDvbSrcCodeRate = "3/5" // 3/5 (10) – 3/5
	GstDvbSrcCodeRate910 GstDvbSrcCodeRate = "9/10" // 9/10 (11) – 9/10
	GstDvbSrcCodeRate25 GstDvbSrcCodeRate = "2/5" // 2/5 (12) – 2/5
)

type GstDvbSrcDelsys string

const (
	GstDvbSrcDelsysUndefined GstDvbSrcDelsys = "undefined" // undefined (0) – UNDEFINED
	GstDvbSrcDelsysDvbCA GstDvbSrcDelsys = "dvb-c-a" // dvb-c-a (1) – DVB-C-A
	GstDvbSrcDelsysDvbCB GstDvbSrcDelsys = "dvb-c-b" // dvb-c-b (2) – DVB-C-B
	GstDvbSrcDelsysDvbT GstDvbSrcDelsys = "dvb-t" // dvb-t (3) – DVB-T
	GstDvbSrcDelsysDss GstDvbSrcDelsys = "dss" // dss (4) – DSS
	GstDvbSrcDelsysDvbS GstDvbSrcDelsys = "dvb-s" // dvb-s (5) – DVB-S
	GstDvbSrcDelsysDvbS2 GstDvbSrcDelsys = "dvb-s2" // dvb-s2 (6) – DVB-S2
	GstDvbSrcDelsysDvbH GstDvbSrcDelsys = "dvb-h" // dvb-h (7) – DVB-H
	GstDvbSrcDelsysIsdbT GstDvbSrcDelsys = "isdb-t" // isdb-t (8) – ISDB-T
	GstDvbSrcDelsysIsdbS GstDvbSrcDelsys = "isdb-s" // isdb-s (9) – ISDB-S
	GstDvbSrcDelsysIsdbC GstDvbSrcDelsys = "isdb-c" // isdb-c (10) – ISDB-C
	GstDvbSrcDelsysAtsc GstDvbSrcDelsys = "atsc" // atsc (11) – ATSC
	GstDvbSrcDelsysAtscMh GstDvbSrcDelsys = "atsc-mh" // atsc-mh (12) – ATSC-MH
	GstDvbSrcDelsysDtmb GstDvbSrcDelsys = "dtmb" // dtmb (13) – DTMB
	GstDvbSrcDelsysCmmb GstDvbSrcDelsys = "cmmb" // cmmb (14) – CMMB
	GstDvbSrcDelsysDab GstDvbSrcDelsys = "dab" // dab (15) – DAB
	GstDvbSrcDelsysDvbT2 GstDvbSrcDelsys = "dvb-t2" // dvb-t2 (16) – DVB-T2
	GstDvbSrcDelsysTurbo GstDvbSrcDelsys = "turbo" // turbo (17) – TURBO
	GstDvbSrcDelsysDvbCC GstDvbSrcDelsys = "dvb-c-c" // dvb-c-c (18) – DVB-C-C
)

type GstDvbSrcHierarchy string

const (
	GstDvbSrcHierarchyNone GstDvbSrcHierarchy = "none" // none (0) – NONE
	GstDvbSrcHierarchy1 GstDvbSrcHierarchy = "1" // 1 (1) – 1
	GstDvbSrcHierarchy2 GstDvbSrcHierarchy = "2" // 2 (2) – 2
	GstDvbSrcHierarchy4 GstDvbSrcHierarchy = "4" // 4 (3) – 4
	GstDvbSrcHierarchyAuto GstDvbSrcHierarchy = "auto" // auto (4) – AUTO
)

type GstDvbSrcInterleaving string

const (
	GstDvbSrcInterleavingNone GstDvbSrcInterleaving = "none" // none (0) – NONE
	GstDvbSrcInterleavingAuto GstDvbSrcInterleaving = "auto" // auto (1) – AUTO
	GstDvbSrcInterleaving240 GstDvbSrcInterleaving = "240" // 240 (2) – 240
	GstDvbSrcInterleaving720 GstDvbSrcInterleaving = "720" // 720 (3) – 720
)

type GstDvbSrcInversion string

const (
	GstDvbSrcInversionOff GstDvbSrcInversion = "off" // off (0) – OFF
	GstDvbSrcInversionOn GstDvbSrcInversion = "on" // on (1) – ON
	GstDvbSrcInversionAuto GstDvbSrcInversion = "auto" // auto (2) – AUTO
)

type GstDvbSrcModulation string

const (
	GstDvbSrcModulationQpsk GstDvbSrcModulation = "qpsk" // qpsk (0) – QPSK
	GstDvbSrcModulationQam16 GstDvbSrcModulation = "qam-16" // qam-16 (1) – QAM 16
	GstDvbSrcModulationQam32 GstDvbSrcModulation = "qam-32" // qam-32 (2) – QAM 32
	GstDvbSrcModulationQam64 GstDvbSrcModulation = "qam-64" // qam-64 (3) – QAM 64
	GstDvbSrcModulationQam128 GstDvbSrcModulation = "qam-128" // qam-128 (4) – QAM 128
	GstDvbSrcModulationQam256 GstDvbSrcModulation = "qam-256" // qam-256 (5) – QAM 256
	GstDvbSrcModulationAuto GstDvbSrcModulation = "auto" // auto (6) – AUTO
	GstDvbSrcModulation8Vsb GstDvbSrcModulation = "8vsb" // 8vsb (7) – 8VSB
	GstDvbSrcModulation16Vsb GstDvbSrcModulation = "16vsb" // 16vsb (8) – 16VSB
	GstDvbSrcModulation8Psk GstDvbSrcModulation = "8psk" // 8psk (9) – 8PSK
	GstDvbSrcModulation16Apsk GstDvbSrcModulation = "16apsk" // 16apsk (10) – 16APSK
	GstDvbSrcModulation32Apsk GstDvbSrcModulation = "32apsk" // 32apsk (11) – 32APSK
	GstDvbSrcModulationDqpsk GstDvbSrcModulation = "dqpsk" // dqpsk (12) – DQPSK
	GstDvbSrcModulationQam4Nr GstDvbSrcModulation = "qam-4-nr" // qam-4-nr (13) – QAM 4 NR
)

type GstDvbSrcBandwidth string

const (
	GstDvbSrcBandwidth8 GstDvbSrcBandwidth = "8" // 8 (0) – 8
	GstDvbSrcBandwidth7 GstDvbSrcBandwidth = "7" // 7 (1) – 7
	GstDvbSrcBandwidth6 GstDvbSrcBandwidth = "6" // 6 (2) – 6
	GstDvbSrcBandwidthAuto GstDvbSrcBandwidth = "AUTO" // AUTO (3) – AUTO
	GstDvbSrcBandwidth5 GstDvbSrcBandwidth = "5" // 5 (4) – 5
	GstDvbSrcBandwidth10 GstDvbSrcBandwidth = "10" // 10 (5) – 10
	GstDvbSrcBandwidth1712 GstDvbSrcBandwidth = "1.712" // 1.712 (6) – 1.712
)

type GstDvbSrcGuard string

const (
	GstDvbSrcGuard32 GstDvbSrcGuard = "32" // 32 (0) – 32
	GstDvbSrcGuard16 GstDvbSrcGuard = "16" // 16 (1) – 16
	GstDvbSrcGuard8 GstDvbSrcGuard = "8" // 8 (2) – 8
	GstDvbSrcGuard4 GstDvbSrcGuard = "4" // 4 (3) – 4
	GstDvbSrcGuardAuto GstDvbSrcGuard = "auto" // auto (4) – AUTO
	GstDvbSrcGuard128 GstDvbSrcGuard = "128" // 128 (5) – 128
	GstDvbSrcGuard19128 GstDvbSrcGuard = "19/128" // 19/128 (6) – 19/128
	GstDvbSrcGuard19256 GstDvbSrcGuard = "19/256" // 19/256 (7) – 19/256
	GstDvbSrcGuardPn420 GstDvbSrcGuard = "pn420" // pn420 (8) – PN420
	GstDvbSrcGuardPn595 GstDvbSrcGuard = "pn595" // pn595 (9) – PN595
	GstDvbSrcGuardPn945 GstDvbSrcGuard = "pn945" // pn945 (10) – PN945
)

type GstDvbSrcTransmissionMode string

const (
	GstDvbSrcTransmissionMode2K GstDvbSrcTransmissionMode = "2k" // 2k (0) – 2K
	GstDvbSrcTransmissionMode8K GstDvbSrcTransmissionMode = "8k" // 8k (1) – 8K
	GstDvbSrcTransmissionModeAuto GstDvbSrcTransmissionMode = "auto" // auto (2) – AUTO
	GstDvbSrcTransmissionMode4K GstDvbSrcTransmissionMode = "4k" // 4k (3) – 4K
	GstDvbSrcTransmissionMode1K GstDvbSrcTransmissionMode = "1k" // 1k (4) – 1K
	GstDvbSrcTransmissionMode16K GstDvbSrcTransmissionMode = "16k" // 16k (5) – 16K
	GstDvbSrcTransmissionMode32K GstDvbSrcTransmissionMode = "32k" // 32k (6) – 32K
	GstDvbSrcTransmissionModeC1 GstDvbSrcTransmissionMode = "c1" // c1 (7) – C1
	GstDvbSrcTransmissionModeC3780 GstDvbSrcTransmissionMode = "c3780" // c3780 (8) – C3780
)

