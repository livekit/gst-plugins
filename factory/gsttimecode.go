// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Drops all audio/video until a specific timecode or running time has been reached
type GstAvWait struct {
	*gst.Element
}

func NewAvWait() (*GstAvWait, error) {
	e, err := gst.NewElement("avwait")
	if err != nil {
		return nil, err
	}
	return &GstAvWait{Element: e}, nil
}

func NewAvWaitWithName(name string) (*GstAvWait, error) {
	e, err := gst.NewElementWithName("avwait", name)
	if err != nil {
		return nil, err
	}
	return &GstAvWait{Element: e}, nil
}

// Whether the element is stopped or recording. If set to FALSE, all buffers will be dropped regardless of settings.
// Default: true
func (e *GstAvWait) SetRecording(value bool) error {
	return e.Element.SetProperty("recording", value)
}

func (e *GstAvWait) GetRecording() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("recording")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Running time to wait for in running-time mode
// Default: 18446744073709551615, Min: 0, Max: 18446744073709551615
func (e *GstAvWait) SetTargetRunningTime(value uint64) error {
	return e.Element.SetProperty("target-running-time", value)
}

func (e *GstAvWait) GetTargetRunningTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("target-running-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Timecode to wait for in timecode mode (object)

func (e *GstAvWait) SetTargetTimecode(value interface{}) error {
	return e.Element.SetProperty("target-timecode", value)
}

func (e *GstAvWait) GetTargetTimecode() (interface{}, error) {
	return e.Element.GetProperty("target-timecode")
}

// Timecode to wait for in timecode mode (string). Must take the form 00:00:00:00
// Default: 00:00:00:00
func (e *GstAvWait) SetTargetTimecodeString(value string) error {
	return e.Element.SetProperty("target-timecode-string", value)
}

func (e *GstAvWait) GetTargetTimecodeString() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("target-timecode-string")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Running time to end at in running-time mode
// Default: 18446744073709551615, Min: 0, Max: 18446744073709551615
func (e *GstAvWait) SetEndRunningTime(value uint64) error {
	return e.Element.SetProperty("end-running-time", value)
}

func (e *GstAvWait) GetEndRunningTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("end-running-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Timecode to end at in timecode mode (object)

func (e *GstAvWait) SetEndTimecode(value interface{}) error {
	return e.Element.SetProperty("end-timecode", value)
}

func (e *GstAvWait) GetEndTimecode() (interface{}, error) {
	return e.Element.GetProperty("end-timecode")
}

// Operation mode: What to wait for
// Default: timecode (0)
func (e *GstAvWait) SetMode(value GstAvWaitMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstAvWait) GetMode() (GstAvWaitMode, error) {
	var value GstAvWaitMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAvWaitMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAvWaitMode")
	}
	return value, nil
}

// Attaches a timecode meta into each video frame
type GstTimeCodeStamper struct {
	*base.GstBaseTransform
}

func NewTimeCodeStamper() (*GstTimeCodeStamper, error) {
	e, err := gst.NewElement("timecodestamper")
	if err != nil {
		return nil, err
	}
	return &GstTimeCodeStamper{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewTimeCodeStamperWithName(name string) (*GstTimeCodeStamper, error) {
	e, err := gst.NewElementWithName("timecodestamper", name)
	if err != nil {
		return nil, err
	}
	return &GstTimeCodeStamper{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// The daily jam of the LTC timecode

func (e *GstTimeCodeStamper) SetLtcDailyJam(value interface{}) error {
	return e.Element.SetProperty("ltc-daily-jam", value)
}

func (e *GstTimeCodeStamper) GetLtcDailyJam() (interface{}, error) {
	return e.Element.GetProperty("ltc-daily-jam")
}

// Extra latency to introduce for waiting for LTC timecodes
// Default: 150000000, Min: 0, Max: 18446744073709551615
func (e *GstTimeCodeStamper) SetLtcExtraLatency(value uint64) error {
	return e.Element.SetProperty("ltc-extra-latency", value)
}

func (e *GstTimeCodeStamper) GetLtcExtraLatency() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ltc-extra-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Use drop-frame timecodes for 29.97 and 59.94 FPS
// Default: false
func (e *GstTimeCodeStamper) SetDropFrame(value bool) error {
	return e.Element.SetProperty("drop-frame", value)
}

func (e *GstTimeCodeStamper) GetDropFrame() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// If true the RTC timecode will be automatically resynced if it drifts, otherwise it will only be counted up from the last known one
// Default: true
func (e *GstTimeCodeStamper) SetRtcAutoResync(value bool) error {
	return e.Element.SetProperty("rtc-auto-resync", value)
}

func (e *GstTimeCodeStamper) GetRtcAutoResync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rtc-auto-resync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum number of nanoseconds the RTC clock is allowed to drift from the video before it is resynced
// Default: 250000000, Min: 0, Max: 18446744073709551615
func (e *GstTimeCodeStamper) SetRtcMaxDrift(value uint64) error {
	return e.Element.SetProperty("rtc-max-drift", value)
}

func (e *GstTimeCodeStamper) GetRtcMaxDrift() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("rtc-max-drift")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// If set, take this timecode as the internal timecode for the first frame and increment from it. Only the values itself and daily jam are taken, flags and frame rate are always determined by timecodestamper itself. If unset, the internal timecode will start at 0 with the daily jam being the current real-time clock time

func (e *GstTimeCodeStamper) SetSetInternalTimecode(value interface{}) error {
	return e.Element.SetProperty("set-internal-timecode", value)
}

func (e *GstTimeCodeStamper) GetSetInternalTimecode() (interface{}, error) {
	return e.Element.GetProperty("set-internal-timecode")
}

// Time out upstream timecode if no new timecode was detected after this time
// Default: 18446744073709551615, Min: 0, Max: 18446744073709551615
func (e *GstTimeCodeStamper) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstTimeCodeStamper) GetTimeout() (uint64, error) {
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

// If true the LTC timecode will be automatically resynced if it drifts, otherwise it will only be counted up from the last known one
// Default: true
func (e *GstTimeCodeStamper) SetLtcAutoResync(value bool) error {
	return e.Element.SetProperty("ltc-auto-resync", value)
}

func (e *GstTimeCodeStamper) GetLtcAutoResync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ltc-auto-resync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Choose whether timecodes should be overridden or not
// Default: keep (1)
func (e *GstTimeCodeStamper) SetSet(value GstTimeCodeStamperSet) error {
	e.Element.SetArg("set", string(value))
	return nil
}

func (e *GstTimeCodeStamper) GetSet() (GstTimeCodeStamperSet, error) {
	var value GstTimeCodeStamperSet
	var ok bool
	v, err := e.Element.GetProperty("set")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTimeCodeStamperSet)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTimeCodeStamperSet")
	}
	return value, nil
}

// If true resync last known timecode from upstream, otherwise only count up from the last known one
// Default: true
func (e *GstTimeCodeStamper) SetAutoResync(value bool) error {
	return e.Element.SetProperty("auto-resync", value)
}

func (e *GstTimeCodeStamper) GetAutoResync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("auto-resync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Time out LTC timecode if no new timecode was detected after this time
// Default: 18446744073709551615, Min: 0, Max: 18446744073709551615
func (e *GstTimeCodeStamper) SetLtcTimeout(value uint64) error {
	return e.Element.SetProperty("ltc-timeout", value)
}

func (e *GstTimeCodeStamper) GetLtcTimeout() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ltc-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Post element message containing the current timecode
// Default: false
func (e *GstTimeCodeStamper) SetPostMessages(value bool) error {
	return e.Element.SetProperty("post-messages", value)
}

func (e *GstTimeCodeStamper) GetPostMessages() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-messages")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Choose from what source the timecode should be taken
// Default: internal (0)
func (e *GstTimeCodeStamper) SetSource(value GstTimeCodeStamperSource) error {
	e.Element.SetArg("source", string(value))
	return nil
}

func (e *GstTimeCodeStamper) GetSource() (GstTimeCodeStamperSource, error) {
	var value GstTimeCodeStamperSource
	var ok bool
	v, err := e.Element.GetProperty("source")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTimeCodeStamperSource)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTimeCodeStamperSource")
	}
	return value, nil
}

// Add this offset in frames to internal, LTC or RTC timecode, useful if there is an offset between the timecode source and video
// Default: 0, Min: -2147483648, Max: 2147483647
func (e *GstTimeCodeStamper) SetTimecodeOffset(value int) error {
	return e.Element.SetProperty("timecode-offset", value)
}

func (e *GstTimeCodeStamper) GetTimecodeOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timecode-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstTimeCodeStamperSet string

const (
	GstTimeCodeStamperSetNever  GstTimeCodeStamperSet = "never"  // Never set timecodes
	GstTimeCodeStamperSetKeep   GstTimeCodeStamperSet = "keep"   // Keep upstream timecodes and only set if no upstream timecode
	GstTimeCodeStamperSetAlways GstTimeCodeStamperSet = "always" // Always set timecode and remove upstream timecode
)

type GstTimeCodeStamperSource string

const (
	GstTimeCodeStamperSourceInternal        GstTimeCodeStamperSource = "internal"           // Use internal timecode counter, starting at zero or value set by property
	GstTimeCodeStamperSourceZero            GstTimeCodeStamperSource = "zero"               // Always use zero
	GstTimeCodeStamperSourceLastKnown       GstTimeCodeStamperSource = "last-known"         // Count up from the last known upstream timecode or internal if unknown
	GstTimeCodeStamperSourceLastKnownOrZero GstTimeCodeStamperSource = "last-known-or-zero" // Count up from the last known upstream timecode or zero if unknown
	GstTimeCodeStamperSourceLtc             GstTimeCodeStamperSource = "ltc"                // Linear timecode from an audio device
	GstTimeCodeStamperSourceRtc             GstTimeCodeStamperSource = "rtc"                // Timecode from real time clock
)

type GstAvWaitMode string

const (
	GstAvWaitModeTimecode    GstAvWaitMode = "timecode"     // time code (default)
	GstAvWaitModeRunningTime GstAvWaitMode = "running-time" // running time
	GstAvWaitModeVideoFirst  GstAvWaitMode = "video-first"  // video first
)
