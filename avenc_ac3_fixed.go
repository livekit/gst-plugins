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

type AvencAc3Fixed struct {
	Element *gst.Element
}

func NewAvencAc3Fixed() (*AvencAc3Fixed, error) {
	e, err := gst.NewElement("avenc_ac3_fixed")
	if err != nil {
		return nil, err
	}
	return &AvencAc3Fixed{Element: e}, nil
}

func NewAvencAc3FixedWithName(name string) (*AvencAc3Fixed, error) {
	e, err := gst.NewElementWithName("avenc_ac3_fixed", name)
	if err != nil {
		return nil, err
	}
	return &AvencAc3Fixed{Element: e}, nil
}

// ----- Properties -----

// ac (int)
//
// set number of audio channels (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetAc() (int, error) {
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

func (e *AvencAc3Fixed) SetAc(value int) error {
	return e.Element.SetProperty("ac", value)
}

// ad-conv-type (GstFixedPointAc3EncoderAdConvType)
//
// A/D Converter Type (Private codec option)

func (e *AvencAc3Fixed) GetAdConvType() (interface{}, error) {
	return e.Element.GetProperty("ad-conv-type")
}

func (e *AvencAc3Fixed) SetAdConvType(value interface{}) error {
	return e.Element.SetProperty("ad-conv-type", value)
}

// ar (int)
//
// set audio sampling rate (in Hz) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetAr() (int, error) {
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

func (e *AvencAc3Fixed) SetAr(value int) error {
	return e.Element.SetProperty("ar", value)
}

// audio-service-type (GstAvcodeccontextAudioServiceType)
//
// audio service type (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetAudioServiceType() (interface{}, error) {
	return e.Element.GetProperty("audio-service-type")
}

func (e *AvencAc3Fixed) SetAudioServiceType(value interface{}) error {
	return e.Element.SetProperty("audio-service-type", value)
}

// bitrate (int)
//
// set bitrate (in bits/s) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetBitrate() (int, error) {
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

func (e *AvencAc3Fixed) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

// bufsize (int)
//
// set ratecontrol buffer size (in bits) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetBufsize() (int, error) {
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

func (e *AvencAc3Fixed) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

// center-mixlev (float32)
//
// Center Mix Level (Private codec option)

func (e *AvencAc3Fixed) GetCenterMixlev() (float32, error) {
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

func (e *AvencAc3Fixed) SetCenterMixlev(value float32) error {
	return e.Element.SetProperty("center-mixlev", value)
}

// channel-coupling (GstFixedPointAc3EncoderChannelCoupling)
//
// Channel Coupling (Private codec option)

func (e *AvencAc3Fixed) GetChannelCoupling() (interface{}, error) {
	return e.Element.GetProperty("channel-coupling")
}

func (e *AvencAc3Fixed) SetChannelCoupling(value interface{}) error {
	return e.Element.SetProperty("channel-coupling", value)
}

// channel-layout (uint64)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetChannelLayout() (uint64, error) {
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

func (e *AvencAc3Fixed) SetChannelLayout(value uint64) error {
	return e.Element.SetProperty("channel-layout", value)
}

// compression-level (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetCompressionLevel() (int, error) {
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

func (e *AvencAc3Fixed) SetCompressionLevel(value int) error {
	return e.Element.SetProperty("compression-level", value)
}

// copyright (int)
//
// Copyright Bit (Private codec option)

func (e *AvencAc3Fixed) GetCopyright() (int, error) {
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

func (e *AvencAc3Fixed) SetCopyright(value int) error {
	return e.Element.SetProperty("copyright", value)
}

// cpl-start-band (GstFixedPointAc3EncoderCplStartBand)
//
// Coupling Start Band (Private codec option)

func (e *AvencAc3Fixed) GetCplStartBand() (interface{}, error) {
	return e.Element.GetProperty("cpl-start-band")
}

func (e *AvencAc3Fixed) SetCplStartBand(value interface{}) error {
	return e.Element.SetProperty("cpl-start-band", value)
}

// cutoff (int)
//
// set cutoff bandwidth (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetCutoff() (int, error) {
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

func (e *AvencAc3Fixed) SetCutoff(value int) error {
	return e.Element.SetProperty("cutoff", value)
}

// debug (GstAvcodeccontextDebug)
//
// print specific debug info (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetDebug() (interface{}, error) {
	return e.Element.GetProperty("debug")
}

func (e *AvencAc3Fixed) SetDebug(value interface{}) error {
	return e.Element.SetProperty("debug", value)
}

// dheadphone-mode (GstFixedPointAc3EncoderDheadphoneMode)
//
// Dolby Headphone Mode (Private codec option)

func (e *AvencAc3Fixed) GetDheadphoneMode() (interface{}, error) {
	return e.Element.GetProperty("dheadphone-mode")
}

func (e *AvencAc3Fixed) SetDheadphoneMode(value interface{}) error {
	return e.Element.SetProperty("dheadphone-mode", value)
}

// dialnorm (int)
//
// Dialogue Level (dB) (Private codec option)

func (e *AvencAc3Fixed) GetDialnorm() (int, error) {
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

func (e *AvencAc3Fixed) SetDialnorm(value int) error {
	return e.Element.SetProperty("dialnorm", value)
}

// dmix-mode (GstFixedPointAc3EncoderDmixMode)
//
// Preferred Stereo Downmix Mode (Private codec option)

func (e *AvencAc3Fixed) GetDmixMode() (interface{}, error) {
	return e.Element.GetProperty("dmix-mode")
}

func (e *AvencAc3Fixed) SetDmixMode(value interface{}) error {
	return e.Element.SetProperty("dmix-mode", value)
}

// dsur-mode (GstFixedPointAc3EncoderDsurMode)
//
// Dolby Surround Mode (Private codec option)

func (e *AvencAc3Fixed) GetDsurMode() (interface{}, error) {
	return e.Element.GetProperty("dsur-mode")
}

func (e *AvencAc3Fixed) SetDsurMode(value interface{}) error {
	return e.Element.SetProperty("dsur-mode", value)
}

// dsurex-mode (GstFixedPointAc3EncoderDsurexMode)
//
// Dolby Surround EX Mode (Private codec option)

func (e *AvencAc3Fixed) GetDsurexMode() (interface{}, error) {
	return e.Element.GetProperty("dsurex-mode")
}

func (e *AvencAc3Fixed) SetDsurexMode(value interface{}) error {
	return e.Element.SetProperty("dsurex-mode", value)
}

// dump-separator (string)
//
// set information dump field separator (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetDumpSeparator() (string, error) {
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

func (e *AvencAc3Fixed) SetDumpSeparator(value string) error {
	return e.Element.SetProperty("dump-separator", value)
}

// err-detect (GstAvcodeccontextErrDetect)
//
// set error detection flags (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetErrDetect() (interface{}, error) {
	return e.Element.GetProperty("err-detect")
}

func (e *AvencAc3Fixed) SetErrDetect(value interface{}) error {
	return e.Element.SetProperty("err-detect", value)
}

// export-side-data (GstAvcodeccontextExportSideData)
//
// Export metadata as side data (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetExportSideData() (interface{}, error) {
	return e.Element.GetProperty("export-side-data")
}

func (e *AvencAc3Fixed) SetExportSideData(value interface{}) error {
	return e.Element.SetProperty("export-side-data", value)
}

// flags (GstAvcodeccontextFlags)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

func (e *AvencAc3Fixed) SetFlags(value interface{}) error {
	return e.Element.SetProperty("flags", value)
}

// flags2 (GstAvcodeccontextFlags2)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetFlags2() (interface{}, error) {
	return e.Element.GetProperty("flags2")
}

func (e *AvencAc3Fixed) SetFlags2(value interface{}) error {
	return e.Element.SetProperty("flags2", value)
}

// frame-size (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetFrameSize() (int, error) {
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

func (e *AvencAc3Fixed) SetFrameSize(value int) error {
	return e.Element.SetProperty("frame-size", value)
}

// global-quality (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetGlobalQuality() (int, error) {
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

func (e *AvencAc3Fixed) SetGlobalQuality(value int) error {
	return e.Element.SetProperty("global-quality", value)
}

// loro-cmixlev (float32)
//
// Lo/Ro Center Mix Level (Private codec option)

func (e *AvencAc3Fixed) GetLoroCmixlev() (float32, error) {
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

func (e *AvencAc3Fixed) SetLoroCmixlev(value float32) error {
	return e.Element.SetProperty("loro-cmixlev", value)
}

// loro-surmixlev (float32)
//
// Lo/Ro Surround Mix Level (Private codec option)

func (e *AvencAc3Fixed) GetLoroSurmixlev() (float32, error) {
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

func (e *AvencAc3Fixed) SetLoroSurmixlev(value float32) error {
	return e.Element.SetProperty("loro-surmixlev", value)
}

// ltrt-cmixlev (float32)
//
// Lt/Rt Center Mix Level (Private codec option)

func (e *AvencAc3Fixed) GetLtrtCmixlev() (float32, error) {
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

func (e *AvencAc3Fixed) SetLtrtCmixlev(value float32) error {
	return e.Element.SetProperty("ltrt-cmixlev", value)
}

// ltrt-surmixlev (float32)
//
// Lt/Rt Surround Mix Level (Private codec option)

func (e *AvencAc3Fixed) GetLtrtSurmixlev() (float32, error) {
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

func (e *AvencAc3Fixed) SetLtrtSurmixlev(value float32) error {
	return e.Element.SetProperty("ltrt-surmixlev", value)
}

// max-pixels (int64)
//
// Maximum number of pixels (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetMaxPixels() (int64, error) {
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

func (e *AvencAc3Fixed) SetMaxPixels(value int64) error {
	return e.Element.SetProperty("max-pixels", value)
}

// max-prediction-order (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetMaxPredictionOrder() (int, error) {
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

func (e *AvencAc3Fixed) SetMaxPredictionOrder(value int) error {
	return e.Element.SetProperty("max-prediction-order", value)
}

// max-samples (int64)
//
// Maximum number of samples (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetMaxSamples() (int64, error) {
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

func (e *AvencAc3Fixed) SetMaxSamples(value int64) error {
	return e.Element.SetProperty("max-samples", value)
}

// maxrate (int64)
//
// maximum bitrate (in bits/s). Used for VBV together with bufsize. (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetMaxrate() (int64, error) {
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

func (e *AvencAc3Fixed) SetMaxrate(value int64) error {
	return e.Element.SetProperty("maxrate", value)
}

// min-prediction-order (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetMinPredictionOrder() (int, error) {
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

func (e *AvencAc3Fixed) SetMinPredictionOrder(value int) error {
	return e.Element.SetProperty("min-prediction-order", value)
}

// minrate (int64)
//
// minimum bitrate (in bits/s). Most useful in setting up a CBR encode. It is of little use otherwise. (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetMinrate() (int64, error) {
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

func (e *AvencAc3Fixed) SetMinrate(value int64) error {
	return e.Element.SetProperty("minrate", value)
}

// mixing-level (int)
//
// Mixing Level (Private codec option)

func (e *AvencAc3Fixed) GetMixingLevel() (int, error) {
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

func (e *AvencAc3Fixed) SetMixingLevel(value int) error {
	return e.Element.SetProperty("mixing-level", value)
}

// original (int)
//
// Original Bit Stream (Private codec option)

func (e *AvencAc3Fixed) GetOriginal() (int, error) {
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

func (e *AvencAc3Fixed) SetOriginal(value int) error {
	return e.Element.SetProperty("original", value)
}

// per-frame-metadata (bool)
//
// Allow Changing Metadata Per-Frame (Private codec option)

func (e *AvencAc3Fixed) GetPerFrameMetadata() (bool, error) {
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

func (e *AvencAc3Fixed) SetPerFrameMetadata(value bool) error {
	return e.Element.SetProperty("per-frame-metadata", value)
}

// room-type (GstFixedPointAc3EncoderRoomType)
//
// Room Type (Private codec option)

func (e *AvencAc3Fixed) GetRoomType() (interface{}, error) {
	return e.Element.GetProperty("room-type")
}

func (e *AvencAc3Fixed) SetRoomType(value interface{}) error {
	return e.Element.SetProperty("room-type", value)
}

// side-data-only-packets (bool)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetSideDataOnlyPackets() (bool, error) {
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

func (e *AvencAc3Fixed) SetSideDataOnlyPackets(value bool) error {
	return e.Element.SetProperty("side-data-only-packets", value)
}

// stereo-rematrixing (bool)
//
// Stereo Rematrixing (Private codec option)

func (e *AvencAc3Fixed) GetStereoRematrixing() (bool, error) {
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

func (e *AvencAc3Fixed) SetStereoRematrixing(value bool) error {
	return e.Element.SetProperty("stereo-rematrixing", value)
}

// strict (GstAvcodeccontextStrict)
//
// how strictly to follow the standards (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetStrict() (interface{}, error) {
	return e.Element.GetProperty("strict")
}

func (e *AvencAc3Fixed) SetStrict(value interface{}) error {
	return e.Element.SetProperty("strict", value)
}

// surround-mixlev (float32)
//
// Surround Mix Level (Private codec option)

func (e *AvencAc3Fixed) GetSurroundMixlev() (float32, error) {
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

func (e *AvencAc3Fixed) SetSurroundMixlev(value float32) error {
	return e.Element.SetProperty("surround-mixlev", value)
}

// thread-type (GstAvcodeccontextThreadType)
//
// select multithreading type (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetThreadType() (interface{}, error) {
	return e.Element.GetProperty("thread-type")
}

func (e *AvencAc3Fixed) SetThreadType(value interface{}) error {
	return e.Element.SetProperty("thread-type", value)
}

// threads (GstAvcodeccontextThreads)
//
// set the number of threads (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetThreads() (interface{}, error) {
	return e.Element.GetProperty("threads")
}

func (e *AvencAc3Fixed) SetThreads(value interface{}) error {
	return e.Element.SetProperty("threads", value)
}

// ticks-per-frame (int)
//
// (null) (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetTicksPerFrame() (int, error) {
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

func (e *AvencAc3Fixed) SetTicksPerFrame(value int) error {
	return e.Element.SetProperty("ticks-per-frame", value)
}

// trellis (int)
//
// rate-distortion optimal quantization (Generic codec option, might have no effect)

func (e *AvencAc3Fixed) GetTrellis() (int, error) {
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

func (e *AvencAc3Fixed) SetTrellis(value int) error {
	return e.Element.SetProperty("trellis", value)
}

// ----- Constants -----

type FixedPointAc3EncoderChannelCoupling string

const (
	FixedPointAc3EncoderChannelCouplingAuto FixedPointAc3EncoderChannelCoupling = "auto" // auto (-1) – Selected by the Encoder
)

type FixedPointAc3EncoderCplStartBand string

const (
	FixedPointAc3EncoderCplStartBandAuto FixedPointAc3EncoderCplStartBand = "auto" // auto (-1) – Selected by the Encoder
)

type FixedPointAc3EncoderDheadphoneMode string

const (
	FixedPointAc3EncoderDheadphoneModeUnknown FixedPointAc3EncoderDheadphoneMode = "unknown" // unknown (-1) – Unspecified
	FixedPointAc3EncoderDheadphoneModeNotindicated FixedPointAc3EncoderDheadphoneMode = "notindicated" // notindicated (0) – Not Indicated (default)
	FixedPointAc3EncoderDheadphoneModeOff FixedPointAc3EncoderDheadphoneMode = "off" // off (1) – Not Dolby Headphone Encoded
	FixedPointAc3EncoderDheadphoneModeOn FixedPointAc3EncoderDheadphoneMode = "on" // on (2) – Dolby Headphone Encoded
)

type FixedPointAc3EncoderDmixMode string

const (
	FixedPointAc3EncoderDmixModeUnknown FixedPointAc3EncoderDmixMode = "unknown" // unknown (-1) – Unspecified
	FixedPointAc3EncoderDmixModeNotindicated FixedPointAc3EncoderDmixMode = "notindicated" // notindicated (0) – Not Indicated (default)
	FixedPointAc3EncoderDmixModeLtrt FixedPointAc3EncoderDmixMode = "ltrt" // ltrt (1) – Lt/Rt Downmix Preferred
	FixedPointAc3EncoderDmixModeLoro FixedPointAc3EncoderDmixMode = "loro" // loro (2) – Lo/Ro Downmix Preferred
	FixedPointAc3EncoderDmixModeDplii FixedPointAc3EncoderDmixMode = "dplii" // dplii (3) – Dolby Pro Logic II Downmix Preferred
)

type FixedPointAc3EncoderDsurMode string

const (
	FixedPointAc3EncoderDsurModeUnknown FixedPointAc3EncoderDsurMode = "unknown" // unknown (-1) – Unspecified
	FixedPointAc3EncoderDsurModeNotindicated FixedPointAc3EncoderDsurMode = "notindicated" // notindicated (0) – Not Indicated (default)
	FixedPointAc3EncoderDsurModeOff FixedPointAc3EncoderDsurMode = "off" // off (1) – Not Dolby Surround Encoded
	FixedPointAc3EncoderDsurModeOn FixedPointAc3EncoderDsurMode = "on" // on (2) – Dolby Surround Encoded
)

type FixedPointAc3EncoderDsurexMode string

const (
	FixedPointAc3EncoderDsurexModeUnknown FixedPointAc3EncoderDsurexMode = "unknown" // unknown (-1) – Unspecified
	FixedPointAc3EncoderDsurexModeNotindicated FixedPointAc3EncoderDsurexMode = "notindicated" // notindicated (0) – Not Indicated (default)
	FixedPointAc3EncoderDsurexModeOff FixedPointAc3EncoderDsurexMode = "off" // off (1) – Not Dolby Surround EX Encoded
	FixedPointAc3EncoderDsurexModeOn FixedPointAc3EncoderDsurexMode = "on" // on (2) – Dolby Surround EX Encoded
	FixedPointAc3EncoderDsurexModeDpliiz FixedPointAc3EncoderDsurexMode = "dpliiz" // dpliiz (3) – Dolby Pro Logic IIz-encoded
)

type FixedPointAc3EncoderRoomType string

const (
	FixedPointAc3EncoderRoomTypeUnknown FixedPointAc3EncoderRoomType = "unknown" // unknown (-1) – Unspecified
	FixedPointAc3EncoderRoomTypeNotindicated FixedPointAc3EncoderRoomType = "notindicated" // notindicated (0) – Not Indicated (default)
	FixedPointAc3EncoderRoomTypeLarge FixedPointAc3EncoderRoomType = "large" // large (1) – Large Room
	FixedPointAc3EncoderRoomTypeSmall FixedPointAc3EncoderRoomType = "small" // small (2) – Small Room
)

type FixedPointAc3EncoderAdConvType string

const (
	FixedPointAc3EncoderAdConvTypeUnknown FixedPointAc3EncoderAdConvType = "unknown" // unknown (-1) – Unspecified
	FixedPointAc3EncoderAdConvTypeStandard FixedPointAc3EncoderAdConvType = "standard" // standard (0) – Standard (default)
	FixedPointAc3EncoderAdConvTypeHdcd FixedPointAc3EncoderAdConvType = "hdcd" // hdcd (1) – HDCD
)

