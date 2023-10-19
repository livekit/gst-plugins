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

type AvencAc3 struct {
	Element *gst.Element
}

func NewAvencAc3() (*AvencAc3, error) {
	e, err := gst.NewElement("avenc_ac3")
	if err != nil {
		return nil, err
	}
	return &AvencAc3{Element: e}, nil
}

func NewAvencAc3WithName(name string) (*AvencAc3, error) {
	e, err := gst.NewElementWithName("avenc_ac3", name)
	if err != nil {
		return nil, err
	}
	return &AvencAc3{Element: e}, nil
}

// ----- Properties -----

// ac (int)
//
// set number of audio channels (Generic codec option, might have no effect)

func (e *AvencAc3) GetAc() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ac")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetAc(value int) error {
	return e.Element.SetProperty("ac", value)
}

// ad-conv-type (GstAc3EncoderAdConvType)
//
// A/D Converter Type (Private codec option)

func (e *AvencAc3) GetAdConvType() (interface{}, error) {
	return e.Element.GetProperty("ad-conv-type")
}

func (e *AvencAc3) SetAdConvType(value interface{}) error {
	return e.Element.SetProperty("ad-conv-type", value)
}

// ar (int)
//
// set audio sampling rate (in Hz) (Generic codec option, might have no effect)

func (e *AvencAc3) GetAr() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ar")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetAr(value int) error {
	return e.Element.SetProperty("ar", value)
}

// audio-service-type (GstAvcodeccontextAudioServiceType)
//
// audio service type (Generic codec option, might have no effect)

func (e *AvencAc3) GetAudioServiceType() (interface{}, error) {
	return e.Element.GetProperty("audio-service-type")
}

func (e *AvencAc3) SetAudioServiceType(value interface{}) error {
	return e.Element.SetProperty("audio-service-type", value)
}

// bitrate (int)
//
// set bitrate (in bits/s) (Generic codec option, might have no effect)

func (e *AvencAc3) GetBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// bufsize (int)
//
// set ratecontrol buffer size (in bits) (Generic codec option, might have no effect)

func (e *AvencAc3) GetBufsize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bufsize")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

// center-mixlev (float32)
//
// Center Mix Level (Private codec option)

func (e *AvencAc3) GetCenterMixlev() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("center-mixlev")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencAc3) SetCenterMixlev(value float32) error {
	return e.Element.SetProperty("center-mixlev", value)
}

// channel-coupling (GstAc3EncoderChannelCoupling)
//
// Channel Coupling (Private codec option)

func (e *AvencAc3) GetChannelCoupling() (interface{}, error) {
	return e.Element.GetProperty("channel-coupling")
}

func (e *AvencAc3) SetChannelCoupling(value interface{}) error {
	return e.Element.SetProperty("channel-coupling", value)
}

// channel-layout (uint64)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3) GetChannelLayout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("channel-layout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *AvencAc3) SetChannelLayout(value uint64) error {
	return e.Element.SetProperty("channel-layout", value)
}

// compression-level (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3) GetCompressionLevel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("compression-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetCompressionLevel(value int) error {
	return e.Element.SetProperty("compression-level", value)
}

// copyright (int)
//
// Copyright Bit (Private codec option)

func (e *AvencAc3) GetCopyright() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("copyright")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetCopyright(value int) error {
	return e.Element.SetProperty("copyright", value)
}

// cpl-start-band (GstAc3EncoderCplStartBand)
//
// Coupling Start Band (Private codec option)

func (e *AvencAc3) GetCplStartBand() (interface{}, error) {
	return e.Element.GetProperty("cpl-start-band")
}

func (e *AvencAc3) SetCplStartBand(value interface{}) error {
	return e.Element.SetProperty("cpl-start-band", value)
}

// cutoff (int)
//
// set cutoff bandwidth (Generic codec option, might have no effect)

func (e *AvencAc3) GetCutoff() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cutoff")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetCutoff(value int) error {
	return e.Element.SetProperty("cutoff", value)
}

// debug (GstAvcodeccontextDebug)
//
// print specific debug info (Generic codec option, might have no effect)

func (e *AvencAc3) GetDebug() (interface{}, error) {
	return e.Element.GetProperty("debug")
}

func (e *AvencAc3) SetDebug(value interface{}) error {
	return e.Element.SetProperty("debug", value)
}

// dheadphone-mode (GstAc3EncoderDheadphoneMode)
//
// Dolby Headphone Mode (Private codec option)

func (e *AvencAc3) GetDheadphoneMode() (interface{}, error) {
	return e.Element.GetProperty("dheadphone-mode")
}

func (e *AvencAc3) SetDheadphoneMode(value interface{}) error {
	return e.Element.SetProperty("dheadphone-mode", value)
}

// dialnorm (int)
//
// Dialogue Level (dB) (Private codec option)

func (e *AvencAc3) GetDialnorm() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("dialnorm")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetDialnorm(value int) error {
	return e.Element.SetProperty("dialnorm", value)
}

// dmix-mode (GstAc3EncoderDmixMode)
//
// Preferred Stereo Downmix Mode (Private codec option)

func (e *AvencAc3) GetDmixMode() (interface{}, error) {
	return e.Element.GetProperty("dmix-mode")
}

func (e *AvencAc3) SetDmixMode(value interface{}) error {
	return e.Element.SetProperty("dmix-mode", value)
}

// dsur-mode (GstAc3EncoderDsurMode)
//
// Dolby Surround Mode (Private codec option)

func (e *AvencAc3) GetDsurMode() (interface{}, error) {
	return e.Element.GetProperty("dsur-mode")
}

func (e *AvencAc3) SetDsurMode(value interface{}) error {
	return e.Element.SetProperty("dsur-mode", value)
}

// dsurex-mode (GstAc3EncoderDsurexMode)
//
// Dolby Surround EX Mode (Private codec option)

func (e *AvencAc3) GetDsurexMode() (interface{}, error) {
	return e.Element.GetProperty("dsurex-mode")
}

func (e *AvencAc3) SetDsurexMode(value interface{}) error {
	return e.Element.SetProperty("dsurex-mode", value)
}

// dump-separator (string)
//
// set information dump field separator (Generic codec option, might have no effect)

func (e *AvencAc3) GetDumpSeparator() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("dump-separator")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *AvencAc3) SetDumpSeparator(value string) error {
	return e.Element.SetProperty("dump-separator", value)
}

// err-detect (GstAvcodeccontextErrDetect)
//
// set error detection flags (Generic codec option, might have no effect)

func (e *AvencAc3) GetErrDetect() (interface{}, error) {
	return e.Element.GetProperty("err-detect")
}

func (e *AvencAc3) SetErrDetect(value interface{}) error {
	return e.Element.SetProperty("err-detect", value)
}

// export-side-data (GstAvcodeccontextExportSideData)
//
// Export metadata as side data (Generic codec option, might have no effect)

func (e *AvencAc3) GetExportSideData() (interface{}, error) {
	return e.Element.GetProperty("export-side-data")
}

func (e *AvencAc3) SetExportSideData(value interface{}) error {
	return e.Element.SetProperty("export-side-data", value)
}

// flags (GstAvcodeccontextFlags)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

func (e *AvencAc3) SetFlags(value interface{}) error {
	return e.Element.SetProperty("flags", value)
}

// flags2 (GstAvcodeccontextFlags2)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3) GetFlags2() (interface{}, error) {
	return e.Element.GetProperty("flags2")
}

func (e *AvencAc3) SetFlags2(value interface{}) error {
	return e.Element.SetProperty("flags2", value)
}

// frame-size (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3) GetFrameSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("frame-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetFrameSize(value int) error {
	return e.Element.SetProperty("frame-size", value)
}

// global-quality (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3) GetGlobalQuality() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("global-quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetGlobalQuality(value int) error {
	return e.Element.SetProperty("global-quality", value)
}

// loro-cmixlev (float32)
//
// Lo/Ro Center Mix Level (Private codec option)

func (e *AvencAc3) GetLoroCmixlev() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("loro-cmixlev")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencAc3) SetLoroCmixlev(value float32) error {
	return e.Element.SetProperty("loro-cmixlev", value)
}

// loro-surmixlev (float32)
//
// Lo/Ro Surround Mix Level (Private codec option)

func (e *AvencAc3) GetLoroSurmixlev() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("loro-surmixlev")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencAc3) SetLoroSurmixlev(value float32) error {
	return e.Element.SetProperty("loro-surmixlev", value)
}

// ltrt-cmixlev (float32)
//
// Lt/Rt Center Mix Level (Private codec option)

func (e *AvencAc3) GetLtrtCmixlev() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ltrt-cmixlev")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencAc3) SetLtrtCmixlev(value float32) error {
	return e.Element.SetProperty("ltrt-cmixlev", value)
}

// ltrt-surmixlev (float32)
//
// Lt/Rt Surround Mix Level (Private codec option)

func (e *AvencAc3) GetLtrtSurmixlev() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("ltrt-surmixlev")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencAc3) SetLtrtSurmixlev(value float32) error {
	return e.Element.SetProperty("ltrt-surmixlev", value)
}

// max-pixels (int64)
//
// Maximum number of pixels (Generic codec option, might have no effect)

func (e *AvencAc3) GetMaxPixels() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-pixels")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *AvencAc3) SetMaxPixels(value int64) error {
	return e.Element.SetProperty("max-pixels", value)
}

// max-prediction-order (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3) GetMaxPredictionOrder() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-prediction-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetMaxPredictionOrder(value int) error {
	return e.Element.SetProperty("max-prediction-order", value)
}

// max-samples (int64)
//
// Maximum number of samples (Generic codec option, might have no effect)

func (e *AvencAc3) GetMaxSamples() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-samples")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *AvencAc3) SetMaxSamples(value int64) error {
	return e.Element.SetProperty("max-samples", value)
}

// maxrate (int64)
//
// maximum bitrate (in bits/s). Used for VBV together with bufsize. (Generic codec option, might have no effect)

func (e *AvencAc3) GetMaxrate() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("maxrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *AvencAc3) SetMaxrate(value int64) error {
	return e.Element.SetProperty("maxrate", value)
}

// min-prediction-order (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3) GetMinPredictionOrder() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-prediction-order")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetMinPredictionOrder(value int) error {
	return e.Element.SetProperty("min-prediction-order", value)
}

// minrate (int64)
//
// minimum bitrate (in bits/s). Most useful in setting up a CBR encode. It is of little use otherwise. (Generic codec option, might have no effect)

func (e *AvencAc3) GetMinrate() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("minrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *AvencAc3) SetMinrate(value int64) error {
	return e.Element.SetProperty("minrate", value)
}

// mixing-level (int)
//
// Mixing Level (Private codec option)

func (e *AvencAc3) GetMixingLevel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("mixing-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetMixingLevel(value int) error {
	return e.Element.SetProperty("mixing-level", value)
}

// original (int)
//
// Original Bit Stream (Private codec option)

func (e *AvencAc3) GetOriginal() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("original")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetOriginal(value int) error {
	return e.Element.SetProperty("original", value)
}

// per-frame-metadata (bool)
//
// Allow Changing Metadata Per-Frame (Private codec option)

func (e *AvencAc3) GetPerFrameMetadata() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("per-frame-metadata")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvencAc3) SetPerFrameMetadata(value bool) error {
	return e.Element.SetProperty("per-frame-metadata", value)
}

// room-type (GstAc3EncoderRoomType)
//
// Room Type (Private codec option)

func (e *AvencAc3) GetRoomType() (interface{}, error) {
	return e.Element.GetProperty("room-type")
}

func (e *AvencAc3) SetRoomType(value interface{}) error {
	return e.Element.SetProperty("room-type", value)
}

// side-data-only-packets (bool)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3) GetSideDataOnlyPackets() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("side-data-only-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvencAc3) SetSideDataOnlyPackets(value bool) error {
	return e.Element.SetProperty("side-data-only-packets", value)
}

// stereo-rematrixing (bool)
//
// Stereo Rematrixing (Private codec option)

func (e *AvencAc3) GetStereoRematrixing() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("stereo-rematrixing")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *AvencAc3) SetStereoRematrixing(value bool) error {
	return e.Element.SetProperty("stereo-rematrixing", value)
}

// strict (GstAvcodeccontextStrict)
//
// how strictly to follow the standards (Generic codec option, might have no effect)

func (e *AvencAc3) GetStrict() (interface{}, error) {
	return e.Element.GetProperty("strict")
}

func (e *AvencAc3) SetStrict(value interface{}) error {
	return e.Element.SetProperty("strict", value)
}

// surround-mixlev (float32)
//
// Surround Mix Level (Private codec option)

func (e *AvencAc3) GetSurroundMixlev() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("surround-mixlev")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *AvencAc3) SetSurroundMixlev(value float32) error {
	return e.Element.SetProperty("surround-mixlev", value)
}

// thread-type (GstAvcodeccontextThreadType)
//
// select multithreading type (Generic codec option, might have no effect)

func (e *AvencAc3) GetThreadType() (interface{}, error) {
	return e.Element.GetProperty("thread-type")
}

func (e *AvencAc3) SetThreadType(value interface{}) error {
	return e.Element.SetProperty("thread-type", value)
}

// threads (GstAvcodeccontextThreads)
//
// set the number of threads (Generic codec option, might have no effect)

func (e *AvencAc3) GetThreads() (interface{}, error) {
	return e.Element.GetProperty("threads")
}

func (e *AvencAc3) SetThreads(value interface{}) error {
	return e.Element.SetProperty("threads", value)
}

// ticks-per-frame (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3) GetTicksPerFrame() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("ticks-per-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetTicksPerFrame(value int) error {
	return e.Element.SetProperty("ticks-per-frame", value)
}

// trellis (int)
//
// rate-distortion optimal quantization (Generic codec option, might have no effect)

func (e *AvencAc3) GetTrellis() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("trellis")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *AvencAc3) SetTrellis(value int) error {
	return e.Element.SetProperty("trellis", value)
}

// ----- Constants -----

type Ac3EncoderRoomType string

const (
	Ac3EncoderRoomTypeUnknown Ac3EncoderRoomType = "unknown" // unknown (-1) – Unspecified
	Ac3EncoderRoomTypeNotindicated Ac3EncoderRoomType = "notindicated" // notindicated (0) – Not Indicated (default)
	Ac3EncoderRoomTypeLarge Ac3EncoderRoomType = "large" // large (1) – Large Room
	Ac3EncoderRoomTypeSmall Ac3EncoderRoomType = "small" // small (2) – Small Room
)

type Ac3EncoderAdConvType string

const (
	Ac3EncoderAdConvTypeUnknown Ac3EncoderAdConvType = "unknown" // unknown (-1) – Unspecified
	Ac3EncoderAdConvTypeStandard Ac3EncoderAdConvType = "standard" // standard (0) – Standard (default)
	Ac3EncoderAdConvTypeHdcd Ac3EncoderAdConvType = "hdcd" // hdcd (1) – HDCD
)

type Ac3EncoderChannelCoupling string

const (
	Ac3EncoderChannelCouplingAuto Ac3EncoderChannelCoupling = "auto" // auto (-1) – Selected by the Encoder
)

type Ac3EncoderCplStartBand string

const (
	Ac3EncoderCplStartBandAuto Ac3EncoderCplStartBand = "auto" // auto (-1) – Selected by the Encoder
)

type Ac3EncoderDheadphoneMode string

const (
	Ac3EncoderDheadphoneModeUnknown Ac3EncoderDheadphoneMode = "unknown" // unknown (-1) – Unspecified
	Ac3EncoderDheadphoneModeNotindicated Ac3EncoderDheadphoneMode = "notindicated" // notindicated (0) – Not Indicated (default)
	Ac3EncoderDheadphoneModeOff Ac3EncoderDheadphoneMode = "off" // off (1) – Not Dolby Headphone Encoded
	Ac3EncoderDheadphoneModeOn Ac3EncoderDheadphoneMode = "on" // on (2) – Dolby Headphone Encoded
)

type Ac3EncoderDmixMode string

const (
	Ac3EncoderDmixModeUnknown Ac3EncoderDmixMode = "unknown" // unknown (-1) – Unspecified
	Ac3EncoderDmixModeNotindicated Ac3EncoderDmixMode = "notindicated" // notindicated (0) – Not Indicated (default)
	Ac3EncoderDmixModeLtrt Ac3EncoderDmixMode = "ltrt" // ltrt (1) – Lt/Rt Downmix Preferred
	Ac3EncoderDmixModeLoro Ac3EncoderDmixMode = "loro" // loro (2) – Lo/Ro Downmix Preferred
	Ac3EncoderDmixModeDplii Ac3EncoderDmixMode = "dplii" // dplii (3) – Dolby Pro Logic II Downmix Preferred
)

type Ac3EncoderDsurMode string

const (
	Ac3EncoderDsurModeUnknown Ac3EncoderDsurMode = "unknown" // unknown (-1) – Unspecified
	Ac3EncoderDsurModeNotindicated Ac3EncoderDsurMode = "notindicated" // notindicated (0) – Not Indicated (default)
	Ac3EncoderDsurModeOff Ac3EncoderDsurMode = "off" // off (1) – Not Dolby Surround Encoded
	Ac3EncoderDsurModeOn Ac3EncoderDsurMode = "on" // on (2) – Dolby Surround Encoded
)

type Ac3EncoderDsurexMode string

const (
	Ac3EncoderDsurexModeUnknown Ac3EncoderDsurexMode = "unknown" // unknown (-1) – Unspecified
	Ac3EncoderDsurexModeNotindicated Ac3EncoderDsurexMode = "notindicated" // notindicated (0) – Not Indicated (default)
	Ac3EncoderDsurexModeOff Ac3EncoderDsurexMode = "off" // off (1) – Not Dolby Surround EX Encoded
	Ac3EncoderDsurexModeOn Ac3EncoderDsurexMode = "on" // on (2) – Dolby Surround EX Encoded
	Ac3EncoderDsurexModeDpliiz Ac3EncoderDsurexMode = "dpliiz" // dpliiz (3) – Dolby Pro Logic IIz-encoded
)

