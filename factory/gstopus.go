// source: gst-plugins-base

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// decode opus streams to audio
type GstOpusDec struct {
	*GstAudioDecoder
}

func NewOpusDec() (*GstOpusDec, error) {
	e, err := gst.NewElement("opusdec")
	if err != nil {
		return nil, err
	}
	return &GstOpusDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

func NewOpusDecWithName(name string) (*GstOpusDec, error) {
	e, err := gst.NewElementWithName("opusdec", name)
	if err != nil {
		return nil, err
	}
	return &GstOpusDec{GstAudioDecoder: &GstAudioDecoder{Element: e}}, nil
}

// Apply gain if any is specified in the header
// Default: true
func (e *GstOpusDec) SetApplyGain(value bool) error {
	return e.Element.SetProperty("apply-gain", value)
}

func (e *GstOpusDec) GetApplyGain() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("apply-gain")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Set to true to enable phase inversion, this will slightly improve stereo quality, but will have side effects when downmixed to mono.
// Default: false
func (e *GstOpusDec) SetPhaseInversion(value bool) error {
	return e.Element.SetProperty("phase-inversion", value)
}

func (e *GstOpusDec) GetPhaseInversion() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("phase-inversion")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Various statistics
// Default: application/x-opusdec-stats, num-pushed=(guint64)0, num-gap=(guint64)0, plc-num-samples=(guint64)0, plc-duration=(guint64)0, bandwidth=(uint)0, sample-rate=(uint)0, gain=(uint)0, last-packet-duration=(uint)0, channels=(uint)0;
func (e *GstOpusDec) GetStats() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Use forward error correction if available (needs PLC enabled)
// Default: false
func (e *GstOpusDec) SetUseInbandFec(value bool) error {
	return e.Element.SetProperty("use-inband-fec", value)
}

func (e *GstOpusDec) GetUseInbandFec() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-inband-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Encodes audio in Opus format
type GstOpusEnc struct {
	*GstAudioEncoder
}

func NewOpusEnc() (*GstOpusEnc, error) {
	e, err := gst.NewElement("opusenc")
	if err != nil {
		return nil, err
	}
	return &GstOpusEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

func NewOpusEncWithName(name string) (*GstOpusEnc, error) {
	e, err := gst.NewElementWithName("opusenc", name)
	if err != nil {
		return nil, err
	}
	return &GstOpusEnc{GstAudioEncoder: &GstAudioEncoder{Element: e}}, nil
}

// Audio Band Width
// Default: fullband (1105)
func (e *GstOpusEnc) SetBandwidth(value GstOpusEncBandwidth) error {
	e.Element.SetArg("bandwidth", string(value))
	return nil
}

func (e *GstOpusEnc) GetBandwidth() (GstOpusEncBandwidth, error) {
	var value GstOpusEncBandwidth
	var ok bool
	v, err := e.Element.GetProperty("bandwidth")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpusEncBandwidth)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpusEncBandwidth")
	}
	return value, nil
}

// Specify an encoding bit-rate (in bps).
// Default: 64000, Min: 4000, Max: 650000
func (e *GstOpusEnc) SetBitrate(value int) error {
	return e.Element.SetProperty("bitrate", value)
}

func (e *GstOpusEnc) GetBitrate() (int, error) {
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

// The duration of an audio frame, in ms
// Default: 20 (20)
func (e *GstOpusEnc) SetFrameSize(value GstOpusEncFrameSize) error {
	e.Element.SetArg("frame-size", string(value))
	return nil
}

func (e *GstOpusEnc) GetFrameSize() (GstOpusEncFrameSize, error) {
	var value GstOpusEncFrameSize
	var ok bool
	v, err := e.Element.GetProperty("frame-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpusEncFrameSize)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpusEncFrameSize")
	}
	return value, nil
}

// Enable in-band forward error correction (use in combination with the packet-loss-percentage property)
// Default: false
func (e *GstOpusEnc) SetInbandFec(value bool) error {
	return e.Element.SetProperty("inband-fec", value)
}

func (e *GstOpusEnc) GetInbandFec() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("inband-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum payload size in bytes
// Default: 4000, Min: 2, Max: 4000
func (e *GstOpusEnc) SetMaxPayloadSize(value uint) error {
	return e.Element.SetProperty("max-payload-size", value)
}

func (e *GstOpusEnc) GetMaxPayloadSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-payload-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Packet loss percentage
// Default: 0, Min: 0, Max: 100
func (e *GstOpusEnc) SetPacketLossPercentage(value int) error {
	return e.Element.SetProperty("packet-loss-percentage", value)
}

func (e *GstOpusEnc) GetPacketLossPercentage() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("packet-loss-percentage")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// What type of audio to optimize for
// Default: generic (2049)
func (e *GstOpusEnc) SetAudioType(value GstOpusEncAudioType) error {
	e.Element.SetArg("audio-type", string(value))
	return nil
}

func (e *GstOpusEnc) GetAudioType() (GstOpusEncAudioType, error) {
	var value GstOpusEncAudioType
	var ok bool
	v, err := e.Element.GetProperty("audio-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpusEncAudioType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpusEncAudioType")
	}
	return value, nil
}

// Bitrate type
// Default: constrained-vbr (2)
func (e *GstOpusEnc) SetBitrateType(value GstOpusEncBitrateType) error {
	e.Element.SetArg("bitrate-type", string(value))
	return nil
}

func (e *GstOpusEnc) GetBitrateType() (GstOpusEncBitrateType, error) {
	var value GstOpusEncBitrateType
	var ok bool
	v, err := e.Element.GetProperty("bitrate-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOpusEncBitrateType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOpusEncBitrateType")
	}
	return value, nil
}

// Complexity
// Default: 10, Min: 0, Max: 10
func (e *GstOpusEnc) SetComplexity(value int) error {
	return e.Element.SetProperty("complexity", value)
}

func (e *GstOpusEnc) GetComplexity() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("complexity")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// DTX
// Default: false
func (e *GstOpusEnc) SetDtx(value bool) error {
	return e.Element.SetProperty("dtx", value)
}

func (e *GstOpusEnc) GetDtx() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dtx")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstOpusEncAudioType string

const (
	GstOpusEncAudioTypeGeneric            GstOpusEncAudioType = "generic"             // Generic audio
	GstOpusEncAudioTypeVoice              GstOpusEncAudioType = "voice"               // Voice
	GstOpusEncAudioTypeRestrictedLowdelay GstOpusEncAudioType = "restricted-lowdelay" // Restricted low delay
)

type GstOpusEncBandwidth string

const (
	GstOpusEncBandwidthNarrowband    GstOpusEncBandwidth = "narrowband"    // Narrow band
	GstOpusEncBandwidthMediumband    GstOpusEncBandwidth = "mediumband"    // Medium band
	GstOpusEncBandwidthWideband      GstOpusEncBandwidth = "wideband"      // Wide band
	GstOpusEncBandwidthSuperwideband GstOpusEncBandwidth = "superwideband" // Super wide band
	GstOpusEncBandwidthFullband      GstOpusEncBandwidth = "fullband"      // Full band
	GstOpusEncBandwidthAuto          GstOpusEncBandwidth = "auto"          // Auto
)

type GstOpusEncBitrateType string

const (
	GstOpusEncBitrateTypeCbr            GstOpusEncBitrateType = "cbr"             // CBR
	GstOpusEncBitrateTypeVbr            GstOpusEncBitrateType = "vbr"             // VBR
	GstOpusEncBitrateTypeConstrainedVbr GstOpusEncBitrateType = "constrained-vbr" // Constrained VBR
)

type GstOpusEncFrameSize string

const (
	GstOpusEncFrameSize25 GstOpusEncFrameSize = "2.5" // 2.5
	GstOpusEncFrameSize5  GstOpusEncFrameSize = "5"   // 5
	GstOpusEncFrameSize10 GstOpusEncFrameSize = "10"  // 10
	GstOpusEncFrameSize20 GstOpusEncFrameSize = "20"  // 20
	GstOpusEncFrameSize40 GstOpusEncFrameSize = "40"  // 40
	GstOpusEncFrameSize60 GstOpusEncFrameSize = "60"  // 60
)
