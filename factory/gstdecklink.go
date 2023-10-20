// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Decklink Sink
type GstDecklinkAudioSink struct {
	*base.GstBaseSink
}

func NewDecklinkAudioSink() (*GstDecklinkAudioSink, error) {
	e, err := gst.NewElement("decklinkaudiosink")
	if err != nil {
		return nil, err
	}
	return &GstDecklinkAudioSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewDecklinkAudioSinkWithName(name string) (*GstDecklinkAudioSink, error) {
	e, err := gst.NewElementWithName("decklinkaudiosink", name)
	if err != nil {
		return nil, err
	}
	return &GstDecklinkAudioSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// The serial number (hardware ID) of the Decklink card
// Default: NULL
func (e *GstDecklinkAudioSink) GetHwSerialNumber() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("hw-serial-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Output device instance to use. Higher priority than "device-number".
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstDecklinkAudioSink) SetPersistentId(value int64) error {
	return e.Element.SetProperty("persistent-id", value)
}

func (e *GstDecklinkAudioSink) GetPersistentId() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("persistent-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Timestamp alignment threshold in nanoseconds
// Default: 40000000, Min: 0, Max: 18446744073709551614
func (e *GstDecklinkAudioSink) SetAlignmentThreshold(value uint64) error {
	return e.Element.SetProperty("alignment-threshold", value)
}

func (e *GstDecklinkAudioSink) GetAlignmentThreshold() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("alignment-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Size of audio buffer in microseconds, this is the minimum latency that the sink reports
// Default: 50000, Min: 0, Max: 18446744073709551615
func (e *GstDecklinkAudioSink) SetBufferTime(value uint64) error {
	return e.Element.SetProperty("buffer-time", value)
}

func (e *GstDecklinkAudioSink) GetBufferTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("buffer-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Output device instance to use
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDecklinkAudioSink) SetDeviceNumber(value int) error {
	return e.Element.SetProperty("device-number", value)
}

func (e *GstDecklinkAudioSink) GetDeviceNumber() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Window of time in nanoseconds to wait before creating a discontinuity
// Default: 1000000000, Min: 0, Max: 18446744073709551614
func (e *GstDecklinkAudioSink) SetDiscontWait(value uint64) error {
	return e.Element.SetProperty("discont-wait", value)
}

func (e *GstDecklinkAudioSink) GetDiscontWait() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("discont-wait")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Decklink Source
type GstDecklinkAudioSrc struct {
	*base.GstPushSrc
}

func NewDecklinkAudioSrc() (*GstDecklinkAudioSrc, error) {
	e, err := gst.NewElement("decklinkaudiosrc")
	if err != nil {
		return nil, err
	}
	return &GstDecklinkAudioSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewDecklinkAudioSrcWithName(name string) (*GstDecklinkAudioSrc, error) {
	e, err := gst.NewElementWithName("decklinkaudiosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstDecklinkAudioSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Audio channels
// Default: 2 (2)
func (e *GstDecklinkAudioSrc) SetChannels(value GstDecklinkAudioChannels) error {
	e.Element.SetArg("channels", string(value))
	return nil
}

func (e *GstDecklinkAudioSrc) GetChannels() (GstDecklinkAudioChannels, error) {
	var value GstDecklinkAudioChannels
	var ok bool
	v, err := e.Element.GetProperty("channels")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkAudioChannels)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkAudioChannels")
	}
	return value, nil
}

// Audio input connection to use
// Default: auto (0)
func (e *GstDecklinkAudioSrc) SetConnection(value GstDecklinkAudioConnection) error {
	e.Element.SetArg("connection", string(value))
	return nil
}

func (e *GstDecklinkAudioSrc) GetConnection() (GstDecklinkAudioConnection, error) {
	var value GstDecklinkAudioConnection
	var ok bool
	v, err := e.Element.GetProperty("connection")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkAudioConnection)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkAudioConnection")
	}
	return value, nil
}

// Output device instance to use
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDecklinkAudioSrc) SetDeviceNumber(value int) error {
	return e.Element.SetProperty("device-number", value)
}

func (e *GstDecklinkAudioSrc) GetDeviceNumber() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Window of time in nanoseconds to wait before creating a discontinuity
// Default: 1000000000, Min: 0, Max: 18446744073709551614
func (e *GstDecklinkAudioSrc) SetDiscontWait(value uint64) error {
	return e.Element.SetProperty("discont-wait", value)
}

func (e *GstDecklinkAudioSrc) GetDiscontWait() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("discont-wait")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The serial number (hardware ID) of the Decklink card
// Default: NULL
func (e *GstDecklinkAudioSrc) GetHwSerialNumber() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("hw-serial-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Output device instance to use. Higher priority than "device-number".
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstDecklinkAudioSrc) SetPersistentId(value int64) error {
	return e.Element.SetProperty("persistent-id", value)
}

func (e *GstDecklinkAudioSrc) GetPersistentId() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("persistent-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Timestamp alignment threshold in nanoseconds
// Default: 40000000, Min: 0, Max: 18446744073709551614
func (e *GstDecklinkAudioSrc) SetAlignmentThreshold(value uint64) error {
	return e.Element.SetProperty("alignment-threshold", value)
}

func (e *GstDecklinkAudioSrc) GetAlignmentThreshold() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("alignment-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Size of internal buffer in number of video frames
// Default: 5, Min: 1, Max: 2147483647
func (e *GstDecklinkAudioSrc) SetBufferSize(value uint) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstDecklinkAudioSrc) GetBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Decklink Sink
type GstDecklinkVideoSink struct {
	*base.GstBaseSink
}

func NewDecklinkVideoSink() (*GstDecklinkVideoSink, error) {
	e, err := gst.NewElement("decklinkvideosink")
	if err != nil {
		return nil, err
	}
	return &GstDecklinkVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewDecklinkVideoSinkWithName(name string) (*GstDecklinkVideoSink, error) {
	e, err := gst.NewElementWithName("decklinkvideosink", name)
	if err != nil {
		return nil, err
	}
	return &GstDecklinkVideoSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Line number to use for inserting closed captions (0 = disabled)
// Default: 0, Min: 0, Max: 22
func (e *GstDecklinkVideoSink) SetCcLine(value int) error {
	return e.Element.SetProperty("cc-line", value)
}

func (e *GstDecklinkVideoSink) GetCcLine() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cc-line")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Output device instance to use
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDecklinkVideoSink) SetDeviceNumber(value int) error {
	return e.Element.SetProperty("device-number", value)
}

func (e *GstDecklinkVideoSink) GetDeviceNumber() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Keyer level
// Default: 255, Min: 0, Max: 255
func (e *GstDecklinkVideoSink) SetKeyerLevel(value int) error {
	return e.Element.SetProperty("keyer-level", value)
}

func (e *GstDecklinkVideoSink) GetKeyerLevel() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("keyer-level")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Keyer mode to be enabled
// Default: off (0)
func (e *GstDecklinkVideoSink) SetKeyerMode(value GstDecklinkKeyerMode) error {
	e.Element.SetArg("keyer-mode", string(value))
	return nil
}

func (e *GstDecklinkVideoSink) GetKeyerMode() (GstDecklinkKeyerMode, error) {
	var value GstDecklinkKeyerMode
	var ok bool
	v, err := e.Element.GetProperty("keyer-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkKeyerMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkKeyerMode")
	}
	return value, nil
}

// Video Mode to use for playback
// Default: ntsc (1)
func (e *GstDecklinkVideoSink) SetMode(value GstDecklinkModes) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstDecklinkVideoSink) GetMode() (GstDecklinkModes, error) {
	var value GstDecklinkModes
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkModes)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkModes")
	}
	return value, nil
}

// Output device instance to use. Higher priority than "device-number".
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstDecklinkVideoSink) SetPersistentId(value int64) error {
	return e.Element.SetProperty("persistent-id", value)
}

func (e *GstDecklinkVideoSink) GetPersistentId() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("persistent-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Video format type to use for playback
// Default: 8bit-yuv (1)
func (e *GstDecklinkVideoSink) SetVideoFormat(value GstDecklinkVideoFormat) error {
	e.Element.SetArg("video-format", string(value))
	return nil
}

func (e *GstDecklinkVideoSink) GetVideoFormat() (GstDecklinkVideoFormat, error) {
	var value GstDecklinkVideoFormat
	var ok bool
	v, err := e.Element.GetProperty("video-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkVideoFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkVideoFormat")
	}
	return value, nil
}

// Line number to use for inserting AFD/Bar data (0 = disabled)
// Default: 0, Min: 0, Max: 10000
func (e *GstDecklinkVideoSink) SetAfdBarLine(value int) error {
	return e.Element.SetProperty("afd-bar-line", value)
}

func (e *GstDecklinkVideoSink) GetAfdBarLine() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("afd-bar-line")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The serial number (hardware ID) of the Decklink card
// Default: NULL
func (e *GstDecklinkVideoSink) GetHwSerialNumber() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("hw-serial-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// 3G-SDI Mapping Format (Level A/B)
// Default: default (0)
func (e *GstDecklinkVideoSink) SetMappingFormat(value GstDecklinkMappingFormat) error {
	e.Element.SetArg("mapping-format", string(value))
	return nil
}

func (e *GstDecklinkVideoSink) GetMappingFormat() (GstDecklinkMappingFormat, error) {
	var value GstDecklinkMappingFormat
	var ok bool
	v, err := e.Element.GetProperty("mapping-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkMappingFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkMappingFormat")
	}
	return value, nil
}

// Certain DeckLink devices such as the DeckLink 8K Pro, the DeckLink Quad 2 and the DeckLink Duo 2 support multiple profiles to configure the capture and playback behavior of its sub-devices.For the DeckLink Duo 2 and DeckLink Quad 2, a profile is shared between any 2 sub-devices that utilize the same connectors. For the DeckLink 8K Pro, a profile is shared between all 4 sub-devices. Any sub-devices that share a profile are considered to be part of the same profile group.DeckLink Duo 2 support configuration of the duplex mode of individual sub-devices.
// Default: default (0)
func (e *GstDecklinkVideoSink) SetProfile(value GstDecklinkProfileId) error {
	e.Element.SetArg("profile", string(value))
	return nil
}

func (e *GstDecklinkVideoSink) GetProfile() (GstDecklinkProfileId, error) {
	var value GstDecklinkProfileId
	var ok bool
	v, err := e.Element.GetProperty("profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkProfileId)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkProfileId")
	}
	return value, nil
}

// Timecode format type to use for playback
// Default: rp188any (3)
func (e *GstDecklinkVideoSink) SetTimecodeFormat(value GstDecklinkTimecodeFormat) error {
	e.Element.SetArg("timecode-format", string(value))
	return nil
}

func (e *GstDecklinkVideoSink) GetTimecodeFormat() (GstDecklinkTimecodeFormat, error) {
	var value GstDecklinkTimecodeFormat
	var ok bool
	v, err := e.Element.GetProperty("timecode-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkTimecodeFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkTimecodeFormat")
	}
	return value, nil
}

// Decklink Source
type GstDecklinkVideoSrc struct {
	*base.GstPushSrc
}

func NewDecklinkVideoSrc() (*GstDecklinkVideoSrc, error) {
	e, err := gst.NewElement("decklinkvideosrc")
	if err != nil {
		return nil, err
	}
	return &GstDecklinkVideoSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewDecklinkVideoSrcWithName(name string) (*GstDecklinkVideoSrc, error) {
	e, err := gst.NewElementWithName("decklinkvideosrc", name)
	if err != nil {
		return nil, err
	}
	return &GstDecklinkVideoSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Video input connection to use
// Default: auto (0)
func (e *GstDecklinkVideoSrc) SetConnection(value GstDecklinkConnection) error {
	e.Element.SetArg("connection", string(value))
	return nil
}

func (e *GstDecklinkVideoSrc) GetConnection() (GstDecklinkConnection, error) {
	var value GstDecklinkConnection
	var ok bool
	v, err := e.Element.GetProperty("connection")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkConnection)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkConnection")
	}
	return value, nil
}

// Extract and output CC as GstMeta (if present)
// Default: false
func (e *GstDecklinkVideoSrc) SetOutputCc(value bool) error {
	return e.Element.SetProperty("output-cc", value)
}

func (e *GstDecklinkVideoSrc) GetOutputCc() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("output-cc")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Output stream time directly instead of translating to pipeline clock
// Default: false
func (e *GstDecklinkVideoSrc) SetOutputStreamTime(value bool) error {
	return e.Element.SetProperty("output-stream-time", value)
}

func (e *GstDecklinkVideoSrc) GetOutputStreamTime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("output-stream-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Output device instance to use
// Default: 0, Min: 0, Max: 2147483647
func (e *GstDecklinkVideoSrc) SetDeviceNumber(value int) error {
	return e.Element.SetProperty("device-number", value)
}

func (e *GstDecklinkVideoSrc) GetDeviceNumber() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("device-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The serial number (hardware ID) of the Decklink card
// Default: NULL
func (e *GstDecklinkVideoSrc) GetHwSerialNumber() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("hw-serial-number")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Video Mode to use for playback
// Default: auto (0)
func (e *GstDecklinkVideoSrc) SetMode(value GstDecklinkModes) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstDecklinkVideoSrc) GetMode() (GstDecklinkModes, error) {
	var value GstDecklinkModes
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkModes)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkModes")
	}
	return value, nil
}

// Output device instance to use. Higher priority than "device-number".
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstDecklinkVideoSrc) SetPersistentId(value int64) error {
	return e.Element.SetProperty("persistent-id", value)
}

func (e *GstDecklinkVideoSrc) GetPersistentId() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("persistent-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Certain DeckLink devices such as the DeckLink 8K Pro, the DeckLink Quad 2 and the DeckLink Duo 2 support multiple profiles to configure the capture and playback behavior of its sub-devices.For the DeckLink Duo 2 and DeckLink Quad 2, a profile is shared between any 2 sub-devices that utilize the same connectors. For the DeckLink 8K Pro, a profile is shared between all 4 sub-devices. Any sub-devices that share a profile are considered to be part of the same profile group.DeckLink Duo 2 support configuration of the duplex mode of individual sub-devices.
// Default: default (0)
func (e *GstDecklinkVideoSrc) SetProfile(value GstDecklinkProfileId) error {
	e.Element.SetArg("profile", string(value))
	return nil
}

func (e *GstDecklinkVideoSrc) GetProfile() (GstDecklinkProfileId, error) {
	var value GstDecklinkProfileId
	var ok bool
	v, err := e.Element.GetProperty("profile")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkProfileId)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkProfileId")
	}
	return value, nil
}

// True if there is a valid input signal available
// Default: false
func (e *GstDecklinkVideoSrc) GetSignal() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("signal")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Skip that much time of initial frames after starting
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstDecklinkVideoSrc) SetSkipFirstTime(value uint64) error {
	return e.Element.SetProperty("skip-first-time", value)
}

func (e *GstDecklinkVideoSrc) GetSkipFirstTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("skip-first-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Timecode format type to use for input
// Default: rp188any (3)
func (e *GstDecklinkVideoSrc) SetTimecodeFormat(value GstDecklinkTimecodeFormat) error {
	e.Element.SetArg("timecode-format", string(value))
	return nil
}

func (e *GstDecklinkVideoSrc) GetTimecodeFormat() (GstDecklinkTimecodeFormat, error) {
	var value GstDecklinkTimecodeFormat
	var ok bool
	v, err := e.Element.GetProperty("timecode-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkTimecodeFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkTimecodeFormat")
	}
	return value, nil
}

// Video format type to use for input (Only use auto for mode=auto)
// Default: auto (0)
func (e *GstDecklinkVideoSrc) SetVideoFormat(value GstDecklinkVideoFormat) error {
	e.Element.SetArg("video-format", string(value))
	return nil
}

func (e *GstDecklinkVideoSrc) GetVideoFormat() (GstDecklinkVideoFormat, error) {
	var value GstDecklinkVideoFormat
	var ok bool
	v, err := e.Element.GetProperty("video-format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDecklinkVideoFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDecklinkVideoFormat")
	}
	return value, nil
}

// Size of internal buffer in number of video frames
// Default: 5, Min: 1, Max: 2147483647
func (e *GstDecklinkVideoSrc) SetBufferSize(value uint) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstDecklinkVideoSrc) GetBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Drop frames that are marked as having no input signal
// Default: false
func (e *GstDecklinkVideoSrc) SetDropNoSignalFrames(value bool) error {
	return e.Element.SetProperty("drop-no-signal-frames", value)
}

func (e *GstDecklinkVideoSrc) GetDropNoSignalFrames() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-no-signal-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Extract and output AFD/Bar as GstMeta (if present)
// Default: false
func (e *GstDecklinkVideoSrc) SetOutputAfdBar(value bool) error {
	return e.Element.SetProperty("output-afd-bar", value)
}

func (e *GstDecklinkVideoSrc) GetOutputAfdBar() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("output-afd-bar")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstDecklinkAudioChannels string

const (
	GstDecklinkAudioChannels2   GstDecklinkAudioChannels = "2"   // 2 Channels
	GstDecklinkAudioChannels8   GstDecklinkAudioChannels = "8"   // 8 Channels
	GstDecklinkAudioChannels16  GstDecklinkAudioChannels = "16"  // 16 Channels
	GstDecklinkAudioChannelsMax GstDecklinkAudioChannels = "max" // Maximum channels supported
)

type GstDecklinkAudioConnection string

const (
	GstDecklinkAudioConnectionAuto      GstDecklinkAudioConnection = "auto"       // Automatic
	GstDecklinkAudioConnectionEmbedded  GstDecklinkAudioConnection = "embedded"   // SDI/HDMI embedded audio
	GstDecklinkAudioConnectionAes       GstDecklinkAudioConnection = "aes"        // AES/EBU input
	GstDecklinkAudioConnectionAnalog    GstDecklinkAudioConnection = "analog"     // Analog input
	GstDecklinkAudioConnectionAnalogXlr GstDecklinkAudioConnection = "analog-xlr" // Analog input (XLR)
	GstDecklinkAudioConnectionAnalogRca GstDecklinkAudioConnection = "analog-rca" // Analog input (RCA)
)

type GstDecklinkConnection string

const (
	GstDecklinkConnectionAuto       GstDecklinkConnection = "auto"        // Auto
	GstDecklinkConnectionSdi        GstDecklinkConnection = "sdi"         // SDI
	GstDecklinkConnectionHdmi       GstDecklinkConnection = "hdmi"        // HDMI
	GstDecklinkConnectionOpticalSdi GstDecklinkConnection = "optical-sdi" // Optical SDI
	GstDecklinkConnectionComponent  GstDecklinkConnection = "component"   // Component
	GstDecklinkConnectionComposite  GstDecklinkConnection = "composite"   // Composite
	GstDecklinkConnectionSvideo     GstDecklinkConnection = "svideo"      // S-Video
)

type GstDecklinkModes string

const (
	GstDecklinkModesAuto               GstDecklinkModes = "auto"                // Automatic detection
	GstDecklinkModesNtsc               GstDecklinkModes = "ntsc"                // NTSC SD 60i
	GstDecklinkModesNtsc2398           GstDecklinkModes = "ntsc2398"            // NTSC SD 60i (24 fps)
	GstDecklinkModesPal                GstDecklinkModes = "pal"                 // PAL SD 50i
	GstDecklinkModesNtscP              GstDecklinkModes = "ntsc-p"              // NTSC SD 60p
	GstDecklinkModesPalP               GstDecklinkModes = "pal-p"               // PAL SD 50p
	GstDecklinkModes1080p2398          GstDecklinkModes = "1080p2398"           // HD1080 23.98p
	GstDecklinkModes1080p24            GstDecklinkModes = "1080p24"             // HD1080 24p
	GstDecklinkModes1080p25            GstDecklinkModes = "1080p25"             // HD1080 25p
	GstDecklinkModes1080p2997          GstDecklinkModes = "1080p2997"           // HD1080 29.97p
	GstDecklinkModes1080p30            GstDecklinkModes = "1080p30"             // HD1080 30p
	GstDecklinkModes1080i50            GstDecklinkModes = "1080i50"             // HD1080 50i
	GstDecklinkModes1080i5994          GstDecklinkModes = "1080i5994"           // HD1080 59.94i
	GstDecklinkModes1080i60            GstDecklinkModes = "1080i60"             // HD1080 60i
	GstDecklinkModes1080p50            GstDecklinkModes = "1080p50"             // HD1080 50p
	GstDecklinkModes1080p5994          GstDecklinkModes = "1080p5994"           // HD1080 59.94p
	GstDecklinkModes1080p60            GstDecklinkModes = "1080p60"             // HD1080 60p
	GstDecklinkModes720p50             GstDecklinkModes = "720p50"              // HD720 50p
	GstDecklinkModes720p5994           GstDecklinkModes = "720p5994"            // HD720 59.94p
	GstDecklinkModes720p60             GstDecklinkModes = "720p60"              // HD720 60p
	GstDecklinkModes1556p2398          GstDecklinkModes = "1556p2398"           // 2k 23.98p
	GstDecklinkModes1556p24            GstDecklinkModes = "1556p24"             // 2k 24p
	GstDecklinkModes1556p25            GstDecklinkModes = "1556p25"             // 2k 25p
	GstDecklinkModes2kdcip2398         GstDecklinkModes = "2kdcip2398"          // 2k dci 23.98p
	GstDecklinkModes2kdcip24           GstDecklinkModes = "2kdcip24"            // 2k dci 24p
	GstDecklinkModes2kdcip25           GstDecklinkModes = "2kdcip25"            // 2k dci 25p
	GstDecklinkModes2kdcip2997         GstDecklinkModes = "2kdcip2997"          // 2k dci 29.97p
	GstDecklinkModes2kdcip30           GstDecklinkModes = "2kdcip30"            // 2k dci 30p
	GstDecklinkModes2kdcip50           GstDecklinkModes = "2kdcip50"            // 2k dci 50p
	GstDecklinkModes2kdcip5994         GstDecklinkModes = "2kdcip5994"          // 2k dci 59.94p
	GstDecklinkModes2kdcip60           GstDecklinkModes = "2kdcip60"            // 2k dci 60p
	GstDecklinkModes2160p2398          GstDecklinkModes = "2160p2398"           // 4k 23.98p
	GstDecklinkModes2160p24            GstDecklinkModes = "2160p24"             // 4k 24p
	GstDecklinkModes2160p25            GstDecklinkModes = "2160p25"             // 4k 25p
	GstDecklinkModes2160p2997          GstDecklinkModes = "2160p2997"           // 4k 29.97p
	GstDecklinkModes2160p30            GstDecklinkModes = "2160p30"             // 4k 30p
	GstDecklinkModes2160p50            GstDecklinkModes = "2160p50"             // 4k 50p
	GstDecklinkModes2160p5994          GstDecklinkModes = "2160p5994"           // 4k 59.94p
	GstDecklinkModes2160p60            GstDecklinkModes = "2160p60"             // 4k 60p
	GstDecklinkModesNtscWidescreen     GstDecklinkModes = "ntsc-widescreen"     // NTSC SD 60i Widescreen
	GstDecklinkModesNtsc2398Widescreen GstDecklinkModes = "ntsc2398-widescreen" // NTSC SD 60i Widescreen (24 fps)
	GstDecklinkModesPalWidescreen      GstDecklinkModes = "pal-widescreen"      // PAL SD 50i Widescreen
	GstDecklinkModesNtscPWidescreen    GstDecklinkModes = "ntsc-p-widescreen"   // NTSC SD 60p Widescreen
	GstDecklinkModesPalPWidescreen     GstDecklinkModes = "pal-p-widescreen"    // PAL SD 50p Widescreen
	GstDecklinkModes4kdcip2398         GstDecklinkModes = "4kdcip2398"          // 4k dci 23.98p
	GstDecklinkModes4kdcip24           GstDecklinkModes = "4kdcip24"            // 4k dci 24p
	GstDecklinkModes4kdcip25           GstDecklinkModes = "4kdcip25"            // 4k dci 25p
	GstDecklinkModes4kdcip2997         GstDecklinkModes = "4kdcip2997"          // 4k dci 29.97p
	GstDecklinkModes4kdcip30           GstDecklinkModes = "4kdcip30"            // 4k dci 30p
	GstDecklinkModes4kdcip50           GstDecklinkModes = "4kdcip50"            // 4k dci 50p
	GstDecklinkModes4kdcip5994         GstDecklinkModes = "4kdcip5994"          // 4k dci 59.94p
	GstDecklinkModes4kdcip60           GstDecklinkModes = "4kdcip60"            // 4k dci 60p
	GstDecklinkModes8kp2398            GstDecklinkModes = "8kp2398"             // 8k 23.98p
	GstDecklinkModes8kp24              GstDecklinkModes = "8kp24"               // 8k 24p
	GstDecklinkModes8kp25              GstDecklinkModes = "8kp25"               // 8k 25p
	GstDecklinkModes8kp2997            GstDecklinkModes = "8kp2997"             // 8k 29.97p
	GstDecklinkModes8kp30              GstDecklinkModes = "8kp30"               // 8k 30p
	GstDecklinkModes8kp50              GstDecklinkModes = "8kp50"               // 8k 50p
	GstDecklinkModes8kp5994            GstDecklinkModes = "8kp5994"             // 8k 59.94p
	GstDecklinkModes8kp60              GstDecklinkModes = "8kp60"               // 8k 60p
	GstDecklinkModes8kdcip2398         GstDecklinkModes = "8kdcip2398"          // 8k dci 23.98p
	GstDecklinkModes8kdcip24           GstDecklinkModes = "8kdcip24"            // 8k dci 24p
	GstDecklinkModes8kdcip25           GstDecklinkModes = "8kdcip25"            // 8k dci 25p
	GstDecklinkModes8kdcip2997         GstDecklinkModes = "8kdcip2997"          // 8k dci 29.97p
	GstDecklinkModes8kdcip30           GstDecklinkModes = "8kdcip30"            // 8k dci 30p
	GstDecklinkModes8kdcip50           GstDecklinkModes = "8kdcip50"            // 8k dci 50p
	GstDecklinkModes8kdcip5994         GstDecklinkModes = "8kdcip5994"          // 8k dci 59.94p
	GstDecklinkModes8kdcip60           GstDecklinkModes = "8kdcip60"            // 8k dci 60p
)

type GstDecklinkProfileId string

const (
	GstDecklinkProfileIdDefault            GstDecklinkProfileId = "default"               // Default, don't change profile
	GstDecklinkProfileIdOneSubDeviceFull   GstDecklinkProfileId = "one-sub-device-full"   // One sub-device, Full-Duplex
	GstDecklinkProfileIdOneSubDeviceHalf   GstDecklinkProfileId = "one-sub-device-half"   // One sub-device, Half-Duplex
	GstDecklinkProfileIdTwoSubDevicesFull  GstDecklinkProfileId = "two-sub-devices-full"  // Two sub-devices, Full-Duplex
	GstDecklinkProfileIdTwoSubDevicesHalf  GstDecklinkProfileId = "two-sub-devices-half"  // Two sub-devices, Half-Duplex
	GstDecklinkProfileIdFourSubDevicesHalf GstDecklinkProfileId = "four-sub-devices-half" // Four sub-devices, Half-Duplex
)

type GstDecklinkKeyerMode string

const (
	GstDecklinkKeyerModeOff      GstDecklinkKeyerMode = "off"      // Off
	GstDecklinkKeyerModeInternal GstDecklinkKeyerMode = "internal" // Internal
	GstDecklinkKeyerModeExternal GstDecklinkKeyerMode = "external" // External
)

type GstDecklinkMappingFormat string

const (
	GstDecklinkMappingFormatDefault GstDecklinkMappingFormat = "default" // Default, don't change mapping format
	GstDecklinkMappingFormatLevelA  GstDecklinkMappingFormat = "level-a" // Level A
	GstDecklinkMappingFormatLevelB  GstDecklinkMappingFormat = "level-b" // Level B
)

type GstDecklinkTimecodeFormat string

const (
	GstDecklinkTimecodeFormatRp188vitc1 GstDecklinkTimecodeFormat = "rp188vitc1" // bmdTimecodeRP188VITC1
	GstDecklinkTimecodeFormatRp188vitc2 GstDecklinkTimecodeFormat = "rp188vitc2" // bmdTimecodeRP188VITC2
	GstDecklinkTimecodeFormatRp188ltc   GstDecklinkTimecodeFormat = "rp188ltc"   // bmdTimecodeRP188LTC
	GstDecklinkTimecodeFormatRp188any   GstDecklinkTimecodeFormat = "rp188any"   // bmdTimecodeRP188Any
	GstDecklinkTimecodeFormatVitc       GstDecklinkTimecodeFormat = "vitc"       // bmdTimecodeVITC
	GstDecklinkTimecodeFormatVitcfield2 GstDecklinkTimecodeFormat = "vitcfield2" // bmdTimecodeVITCField2
	GstDecklinkTimecodeFormatSerial     GstDecklinkTimecodeFormat = "serial"     // bmdTimecodeSerial
)

type GstDecklinkVideoFormat string

const (
	GstDecklinkVideoFormatAuto     GstDecklinkVideoFormat = "auto"      // Auto
	GstDecklinkVideoFormat8bitYuv  GstDecklinkVideoFormat = "8bit-yuv"  // bmdFormat8BitYUV
	GstDecklinkVideoFormat10bitYuv GstDecklinkVideoFormat = "10bit-yuv" // bmdFormat10BitYUV
	GstDecklinkVideoFormat8bitArgb GstDecklinkVideoFormat = "8bit-argb" // bmdFormat8BitARGB
	GstDecklinkVideoFormat8bitBgra GstDecklinkVideoFormat = "8bit-bgra" // bmdFormat8BitBGRA
	GstDecklinkVideoFormat10bitRgb GstDecklinkVideoFormat = "10bit-rgb" // bmdFormat10BitRGB
)
