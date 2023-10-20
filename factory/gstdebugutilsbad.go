// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Sink to test video CODEC conformance
type GstVideoCodecTestSink struct {
	*base.GstBaseSink
}

func NewVideoCodecTestSink() (*GstVideoCodecTestSink, error) {
	e, err := gst.NewElement("videocodectestsink")
	if err != nil {
		return nil, err
	}
	return &GstVideoCodecTestSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewVideoCodecTestSinkWithName(name string) (*GstVideoCodecTestSink, error) {
	e, err := gst.NewElementWithName("videocodectestsink", name)
	if err != nil {
		return nil, err
	}
	return &GstVideoCodecTestSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// File path to store non-padded I420 stream (optional).
// Default: NULL
func (e *GstVideoCodecTestSink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstVideoCodecTestSink) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Calculates a checksum for buffers
type GstChecksumSink struct {
	*base.GstBaseSink
}

func NewChecksumSink() (*GstChecksumSink, error) {
	e, err := gst.NewElement("checksumsink")
	if err != nil {
		return nil, err
	}
	return &GstChecksumSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewChecksumSinkWithName(name string) (*GstChecksumSink, error) {
	e, err := gst.NewElementWithName("checksumsink", name)
	if err != nil {
		return nil, err
	}
	return &GstChecksumSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Checksum type
// Default: sha1 (1)
func (e *GstChecksumSink) SetHash(value GstChecksumSinkHash) error {
	e.Element.SetArg("hash", string(value))
	return nil
}

func (e *GstChecksumSink) GetHash() (GstChecksumSinkHash, error) {
	var value GstChecksumSinkHash
	var ok bool
	v, err := e.Element.GetProperty("hash")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstChecksumSinkHash)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstChecksumSinkHash")
	}
	return value, nil
}

// Split up a stream into randomly-sized buffers
type GstChopMyData struct {
	*gst.Element
}

func NewChopMyData() (*GstChopMyData, error) {
	e, err := gst.NewElement("chopmydata")
	if err != nil {
		return nil, err
	}
	return &GstChopMyData{Element: e}, nil
}

func NewChopMyDataWithName(name string) (*GstChopMyData, error) {
	e, err := gst.NewElementWithName("chopmydata", name)
	if err != nil {
		return nil, err
	}
	return &GstChopMyData{Element: e}, nil
}

// Maximum size of outgoing buffers
// Default: 4096, Min: 1, Max: 2147483647
func (e *GstChopMyData) SetMaxSize(value int) error {
	return e.Element.SetProperty("max-size", value)
}

func (e *GstChopMyData) GetMaxSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum size of outgoing buffers
// Default: 1, Min: 1, Max: 2147483647
func (e *GstChopMyData) SetMinSize(value int) error {
	return e.Element.SetProperty("min-size", value)
}

func (e *GstChopMyData) GetMinSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Step increment for random buffer sizes
// Default: 1, Min: 1, Max: 2147483647
func (e *GstChopMyData) SetStepSize(value int) error {
	return e.Element.SetProperty("step-size", value)
}

func (e *GstChopMyData) GetStepSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("step-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Pipeline that enables different clocks
type GstClockSelect struct {
	*gst.Pipeline
}

func NewClockSelect() (*GstClockSelect, error) {
	e, err := gst.NewElement("clockselect")
	if err != nil {
		return nil, err
	}
	return &GstClockSelect{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

func NewClockSelectWithName(name string) (*GstClockSelect, error) {
	e, err := gst.NewElementWithName("clockselect", name)
	if err != nil {
		return nil, err
	}
	return &GstClockSelect{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

// ID of pipeline clock
// Default: default (0)
func (e *GstClockSelect) SetClockId(value GstClockSelectClockId) error {
	e.Element.SetArg("clock-id", string(value))
	return nil
}

func (e *GstClockSelect) GetClockId() (GstClockSelectClockId, error) {
	var value GstClockSelectClockId
	var ok bool
	v, err := e.Element.GetProperty("clock-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstClockSelectClockId)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstClockSelectClockId")
	}
	return value, nil
}

// PTP clock domain (meaningful only when Clock ID is PTP)
// Default: 0, Min: 0, Max: 255
func (e *GstClockSelect) SetPtpDomain(value uint) error {
	return e.Element.SetProperty("ptp-domain", value)
}

func (e *GstClockSelect) GetPtpDomain() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ptp-domain")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Compares incoming buffers
type GstCompare struct {
	*gst.Element
}

func NewCompare() (*GstCompare, error) {
	e, err := gst.NewElement("compare")
	if err != nil {
		return nil, err
	}
	return &GstCompare{Element: e}, nil
}

func NewCompareWithName(name string) (*GstCompare, error) {
	e, err := gst.NewElementWithName("compare", name)
	if err != nil {
		return nil, err
	}
	return &GstCompare{Element: e}, nil
}

// Threshold beyond which to consider content different as determined by content-method
// Default: 0, Min: 0, Max: 1.79769e+308
func (e *GstCompare) SetThreshold(value float64) error {
	return e.Element.SetProperty("threshold", value)
}

func (e *GstCompare) GetThreshold() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Whether threshold value is upper bound or lower bound for difference measure
// Default: true
func (e *GstCompare) SetUpper(value bool) error {
	return e.Element.SetProperty("upper", value)
}

func (e *GstCompare) GetUpper() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("upper")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Indicates which metadata should be compared
// Default: memory+meta+timestamps+flags
func (e *GstCompare) SetMeta(value interface{}) error {
	return e.Element.SetProperty("meta", value)
}

func (e *GstCompare) GetMeta() (interface{}, error) {
	return e.Element.GetProperty("meta")
}

// Method to compare buffer content
// Default: mem (0)
func (e *GstCompare) SetMethod(value GstCompareMethod) error {
	e.Element.SetArg("method", string(value))
	return nil
}

func (e *GstCompare) GetMethod() (GstCompareMethod, error) {
	var value GstCompareMethod
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCompareMethod)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCompareMethod")
	}
	return value, nil
}

// Consider OFFSET and OFFSET_END part of timestamp metadata
// Default: false
func (e *GstCompare) SetOffsetTs(value bool) error {
	return e.Element.SetProperty("offset-ts", value)
}

func (e *GstCompare) GetOffsetTs() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("offset-ts")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Pass through all packets but ignore some GstFlowReturn types
type GstErrorIgnore struct {
	*gst.Element
}

func NewErrorIgnore() (*GstErrorIgnore, error) {
	e, err := gst.NewElement("errorignore")
	if err != nil {
		return nil, err
	}
	return &GstErrorIgnore{Element: e}, nil
}

func NewErrorIgnoreWithName(name string) (*GstErrorIgnore, error) {
	e, err := gst.NewElementWithName("errorignore", name)
	if err != nil {
		return nil, err
	}
	return &GstErrorIgnore{Element: e}, nil
}

// Whether to ignore GST_FLOW_ERROR
// Default: true
func (e *GstErrorIgnore) SetIgnoreError(value bool) error {
	return e.Element.SetProperty("ignore-error", value)
}

func (e *GstErrorIgnore) GetIgnoreError() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-error")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to ignore GST_FLOW_NOT_LINKED
// Default: false
func (e *GstErrorIgnore) SetIgnoreNotlinked(value bool) error {
	return e.Element.SetProperty("ignore-notlinked", value)
}

func (e *GstErrorIgnore) GetIgnoreNotlinked() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-notlinked")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to ignore GST_FLOW_NOT_NEGOTIATED
// Default: true
func (e *GstErrorIgnore) SetIgnoreNotnegotiated(value bool) error {
	return e.Element.SetProperty("ignore-notnegotiated", value)
}

func (e *GstErrorIgnore) GetIgnoreNotnegotiated() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-notnegotiated")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Which GstFlowReturn value we should convert to when ignoring
// Default: not-linked (-1)
func (e *GstErrorIgnore) SetConvertTo(value interface{}) error {
	return e.Element.SetProperty("convert-to", value)
}

func (e *GstErrorIgnore) GetConvertTo() (interface{}, error) {
	return e.Element.GetProperty("convert-to")
}

// Whether to ignore GST_FLOW_EOS
// Default: false
func (e *GstErrorIgnore) SetIgnoreEos(value bool) error {
	return e.Element.SetProperty("ignore-eos", value)
}

func (e *GstErrorIgnore) GetIgnoreEos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-eos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Fake audio renderer
type GstFakeAudioSink struct {
	*gst.Bin
}

func NewFakeAudioSink() (*GstFakeAudioSink, error) {
	e, err := gst.NewElement("fakeaudiosink")
	if err != nil {
		return nil, err
	}
	return &GstFakeAudioSink{Bin: &gst.Bin{Element: e}}, nil
}

func NewFakeAudioSinkWithName(name string) (*GstFakeAudioSink, error) {
	e, err := gst.NewElementWithName("fakeaudiosink", name)
	if err != nil {
		return nil, err
	}
	return &GstFakeAudioSink{Bin: &gst.Bin{Element: e}}, nil
}

// Mute the audio channel without changing the volume
// Default: false
func (e *GstFakeAudioSink) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

func (e *GstFakeAudioSink) GetMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mute")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of buffers to accept going EOS
// Default: -1, Min: -1, Max: 2147483647
func (e *GstFakeAudioSink) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

func (e *GstFakeAudioSink) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Go asynchronously to PAUSED
// Default: true
func (e *GstFakeAudioSink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

func (e *GstFakeAudioSink) GetAsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Size in bytes to pull per buffer (0 = default)
// Default: 4096, Min: 0, Max: -1
func (e *GstFakeAudioSink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstFakeAudioSink) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Can activate in pull mode
// Default: false
func (e *GstFakeAudioSink) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

func (e *GstFakeAudioSink) GetCanActivatePull() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-pull")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Can activate in push mode
// Default: true
func (e *GstFakeAudioSink) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

func (e *GstFakeAudioSink) GetCanActivatePush() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-push")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Drop and don't render / hand off out-of-segment buffers
// Default: true
func (e *GstFakeAudioSink) SetDropOutOfSegment(value bool) error {
	return e.Element.SetProperty("drop-out-of-segment", value)
}

func (e *GstFakeAudioSink) GetDropOutOfSegment() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-out-of-segment")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstFakeAudioSink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

func (e *GstFakeAudioSink) GetMaxLateness() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-lateness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Maximum processing time for a buffer in nanoseconds
// Default: 20000000, Min: 0, Max: 18446744073709551615
func (e *GstFakeAudioSink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

func (e *GstFakeAudioSink) GetProcessingDeadline() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("processing-deadline")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Generate a state change error
// Default: none (0)
func (e *GstFakeAudioSink) SetStateError(value GstFakeAudioSinkStateError) error {
	e.Element.SetArg("state-error", string(value))
	return nil
}

func (e *GstFakeAudioSink) GetStateError() (GstFakeAudioSinkStateError, error) {
	var value GstFakeAudioSinkStateError
	var ok bool
	v, err := e.Element.GetProperty("state-error")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeAudioSinkStateError)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeAudioSinkStateError")
	}
	return value, nil
}

// The audio volume, 1.0=100%%
// Default: 1, Min: 0, Max: 10
func (e *GstFakeAudioSink) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

func (e *GstFakeAudioSink) GetVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// The message describing current status
// Default: NULL
func (e *GstFakeAudioSink) GetLastMessage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("last-message")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The maximum bits per second to render (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstFakeAudioSink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstFakeAudioSink) GetMaxBitrate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Timestamp offset in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstFakeAudioSink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstFakeAudioSink) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Dump buffer contents to stdout
// Default: false
func (e *GstFakeAudioSink) SetDump(value bool) error {
	return e.Element.SetProperty("dump", value)
}

func (e *GstFakeAudioSink) GetDump() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dump")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The last sample received in the sink

func (e *GstFakeAudioSink) GetLastSample() (*gst.Sample, error) {
	var value *gst.Sample
	var ok bool
	v, err := e.Element.GetProperty("last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Sample)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Sample")
	}
	return value, nil
}

// Generate Quality-of-Service events upstream
// Default: true
func (e *GstFakeAudioSink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

func (e *GstFakeAudioSink) GetQos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("qos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Additional render delay of the sink in nanoseconds
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstFakeAudioSink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

func (e *GstFakeAudioSink) GetRenderDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("render-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Sink Statistics
// Default: application/x-gst-base-sink-stats, average-rate=(double)0, dropped=(guint64)0, rendered=(guint64)0;
func (e *GstFakeAudioSink) GetStats() (*gst.Structure, error) {
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

// Sync on the clock
// Default: true
func (e *GstFakeAudioSink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

func (e *GstFakeAudioSink) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable the last-sample property
// Default: true
func (e *GstFakeAudioSink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

func (e *GstFakeAudioSink) GetEnableLastSample() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Send a signal before unreffing the buffer
// Default: false
func (e *GstFakeAudioSink) SetSignalHandoffs(value bool) error {
	return e.Element.SetProperty("signal-handoffs", value)
}

func (e *GstFakeAudioSink) GetSignalHandoffs() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("signal-handoffs")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Don't produce last_message events
// Default: true
func (e *GstFakeAudioSink) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstFakeAudioSink) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The time to keep between rendered buffers (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstFakeAudioSink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

func (e *GstFakeAudioSink) GetThrottleTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("throttle-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Fake video display that allows zero-copy
type GstFakeVideoSink struct {
	*gst.Bin
}

func NewFakeVideoSink() (*GstFakeVideoSink, error) {
	e, err := gst.NewElement("fakevideosink")
	if err != nil {
		return nil, err
	}
	return &GstFakeVideoSink{Bin: &gst.Bin{Element: e}}, nil
}

func NewFakeVideoSinkWithName(name string) (*GstFakeVideoSink, error) {
	e, err := gst.NewElementWithName("fakevideosink", name)
	if err != nil {
		return nil, err
	}
	return &GstFakeVideoSink{Bin: &gst.Bin{Element: e}}, nil
}

// Can activate in pull mode
// Default: false
func (e *GstFakeVideoSink) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

func (e *GstFakeVideoSink) GetCanActivatePull() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-pull")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Drop and don't render / hand off out-of-segment buffers
// Default: true
func (e *GstFakeVideoSink) SetDropOutOfSegment(value bool) error {
	return e.Element.SetProperty("drop-out-of-segment", value)
}

func (e *GstFakeVideoSink) GetDropOutOfSegment() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-out-of-segment")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The message describing current status
// Default: NULL
func (e *GstFakeVideoSink) GetLastMessage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("last-message")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Number of buffers to accept going EOS
// Default: -1, Min: -1, Max: 2147483647
func (e *GstFakeVideoSink) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

func (e *GstFakeVideoSink) GetNumBuffers() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Maximum processing time for a buffer in nanoseconds
// Default: 15000000, Min: 0, Max: 18446744073709551615
func (e *GstFakeVideoSink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

func (e *GstFakeVideoSink) GetProcessingDeadline() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("processing-deadline")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Don't produce last_message events
// Default: true
func (e *GstFakeVideoSink) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstFakeVideoSink) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Flags to control behaviour
// Default: overlay-composition+crop
func (e *GstFakeVideoSink) SetAllocationMetaFlags(value GstFakeVideoSinkAllocationMetaFlags) error {
	e.Element.SetArg("allocation-meta-flags", fmt.Sprint(value))
	return nil
}

func (e *GstFakeVideoSink) GetAllocationMetaFlags() (GstFakeVideoSinkAllocationMetaFlags, error) {
	var value GstFakeVideoSinkAllocationMetaFlags
	var ok bool
	v, err := e.Element.GetProperty("allocation-meta-flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeVideoSinkAllocationMetaFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeVideoSinkAllocationMetaFlags")
	}
	return value, nil
}

// Go asynchronously to PAUSED
// Default: true
func (e *GstFakeVideoSink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

func (e *GstFakeVideoSink) GetAsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The time to keep between rendered buffers (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstFakeVideoSink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

func (e *GstFakeVideoSink) GetThrottleTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("throttle-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Sink Statistics
// Default: application/x-gst-base-sink-stats, average-rate=(double)0, dropped=(guint64)0, rendered=(guint64)0;
func (e *GstFakeVideoSink) GetStats() (*gst.Structure, error) {
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

// Sync on the clock
// Default: true
func (e *GstFakeVideoSink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

func (e *GstFakeVideoSink) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Generate Quality-of-Service events upstream
// Default: true
func (e *GstFakeVideoSink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

func (e *GstFakeVideoSink) GetQos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("qos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Additional render delay of the sink in nanoseconds
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstFakeVideoSink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

func (e *GstFakeVideoSink) GetRenderDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("render-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Dump buffer contents to stdout
// Default: false
func (e *GstFakeVideoSink) SetDump(value bool) error {
	return e.Element.SetProperty("dump", value)
}

func (e *GstFakeVideoSink) GetDump() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dump")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)
// Default: 5000000, Min: -1, Max: 9223372036854775807
func (e *GstFakeVideoSink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

func (e *GstFakeVideoSink) GetMaxLateness() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-lateness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Enable the last-sample property
// Default: true
func (e *GstFakeVideoSink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

func (e *GstFakeVideoSink) GetEnableLastSample() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Send a signal before unreffing the buffer
// Default: false
func (e *GstFakeVideoSink) SetSignalHandoffs(value bool) error {
	return e.Element.SetProperty("signal-handoffs", value)
}

func (e *GstFakeVideoSink) GetSignalHandoffs() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("signal-handoffs")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The last sample received in the sink

func (e *GstFakeVideoSink) GetLastSample() (*gst.Sample, error) {
	var value *gst.Sample
	var ok bool
	v, err := e.Element.GetProperty("last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Sample)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Sample")
	}
	return value, nil
}

// The maximum bits per second to render (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstFakeVideoSink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstFakeVideoSink) GetMaxBitrate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Generate a state change error
// Default: none (0)
func (e *GstFakeVideoSink) SetStateError(value GstFakeVideoSinkStateError) error {
	e.Element.SetArg("state-error", string(value))
	return nil
}

func (e *GstFakeVideoSink) GetStateError() (GstFakeVideoSinkStateError, error) {
	var value GstFakeVideoSinkStateError
	var ok bool
	v, err := e.Element.GetProperty("state-error")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeVideoSinkStateError)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeVideoSinkStateError")
	}
	return value, nil
}

// Timestamp offset in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstFakeVideoSink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstFakeVideoSink) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Size in bytes to pull per buffer (0 = default)
// Default: 4096, Min: 0, Max: -1
func (e *GstFakeVideoSink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstFakeVideoSink) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Can activate in push mode
// Default: true
func (e *GstFakeVideoSink) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

func (e *GstFakeVideoSink) GetCanActivatePush() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("can-activate-push")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// DebugSpy provides information on buffers with bus messages
type GstDebugSpy struct {
	*base.GstBaseTransform
}

func NewDebugSpy() (*GstDebugSpy, error) {
	e, err := gst.NewElement("debugspy")
	if err != nil {
		return nil, err
	}
	return &GstDebugSpy{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewDebugSpyWithName(name string) (*GstDebugSpy, error) {
	e, err := gst.NewElementWithName("debugspy", name)
	if err != nil {
		return nil, err
	}
	return &GstDebugSpy{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Checksum algorithm to use
// Default: sha1 (1)
func (e *GstDebugSpy) SetChecksumType(value interface{}) error {
	return e.Element.SetProperty("checksum-type", value)
}

func (e *GstDebugSpy) GetChecksumType() (interface{}, error) {
	return e.Element.GetProperty("checksum-type")
}

// Produce verbose output ?
// Default: false
func (e *GstDebugSpy) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstDebugSpy) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Shows the current frame-rate and drop-rate of the videosink as overlay or text on stdout
type GstFPSDisplaySink struct {
	*gst.Bin
}

func NewFPSDisplaySink() (*GstFPSDisplaySink, error) {
	e, err := gst.NewElement("fpsdisplaysink")
	if err != nil {
		return nil, err
	}
	return &GstFPSDisplaySink{Bin: &gst.Bin{Element: e}}, nil
}

func NewFPSDisplaySinkWithName(name string) (*GstFPSDisplaySink, error) {
	e, err := gst.NewElementWithName("fpsdisplaysink", name)
	if err != nil {
		return nil, err
	}
	return &GstFPSDisplaySink{Bin: &gst.Bin{Element: e}}, nil
}

// Sync on the clock (if the internally used sink doesn't have this property it will be ignored
// Default: true
func (e *GstFPSDisplaySink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

func (e *GstFPSDisplaySink) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Video sink to use (Must only be called on NULL state)

func (e *GstFPSDisplaySink) SetVideoSink(value *gst.Element) error {
	return e.Element.SetProperty("video-sink", value)
}

func (e *GstFPSDisplaySink) GetVideoSink() (*gst.Element, error) {
	var value *gst.Element
	var ok bool
	v, err := e.Element.GetProperty("video-sink")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Element)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Element")
	}
	return value, nil
}

// Time between consecutive frames per second measures and update  (in ms). Should be set on NULL state
// Default: 500, Min: 1, Max: 2147483647
func (e *GstFPSDisplaySink) SetFpsUpdateInterval(value int) error {
	return e.Element.SetProperty("fps-update-interval", value)
}

func (e *GstFPSDisplaySink) GetFpsUpdateInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fps-update-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of frames dropped by the sink
// Default: 0, Min: 0, Max: -1
func (e *GstFPSDisplaySink) GetFramesDropped() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("frames-dropped")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The message describing current status
// Default: NULL
func (e *GstFPSDisplaySink) GetLastMessage() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("last-message")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// If the fps-measurements signal should be emitted.
// Default: false
func (e *GstFPSDisplaySink) SetSignalFpsMeasurements(value bool) error {
	return e.Element.SetProperty("signal-fps-measurements", value)
}

func (e *GstFPSDisplaySink) GetSignalFpsMeasurements() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("signal-fps-measurements")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether to use text-overlay
// Default: true
func (e *GstFPSDisplaySink) SetTextOverlay(value bool) error {
	return e.Element.SetProperty("text-overlay", value)
}

func (e *GstFPSDisplaySink) GetTextOverlay() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("text-overlay")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Number of frames rendered
// Default: 0, Min: 0, Max: -1
func (e *GstFPSDisplaySink) GetFramesRendered() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("frames-rendered")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum fps rate measured. Reset when going from NULL to READY.-1 means no measurement has yet been done
// Default: -1, Min: -1, Max: 1.79769e+308
func (e *GstFPSDisplaySink) GetMaxFps() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("max-fps")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Minimum fps rate measured. Reset when going from NULL to READY.-1 means no measurement has yet been done
// Default: -1, Min: -1, Max: 1.79769e+308
func (e *GstFPSDisplaySink) GetMinFps() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("min-fps")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Don't produce last_message events
// Default: false
func (e *GstFPSDisplaySink) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstFPSDisplaySink) GetSilent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("silent")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Simple container object
type GstTestSrcBin struct {
	*gst.Bin
}

func NewTestSrcBin() (*GstTestSrcBin, error) {
	e, err := gst.NewElement("testsrcbin")
	if err != nil {
		return nil, err
	}
	return &GstTestSrcBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewTestSrcBinWithName(name string) (*GstTestSrcBin, error) {
	e, err := gst.NewElementWithName("testsrcbin", name)
	if err != nil {
		return nil, err
	}
	return &GstTestSrcBin{Bin: &gst.Bin{Element: e}}, nil
}

//	Whether to expose sources at random time to simulate a source that is reading a file and exposing the srcpads later.
//
// Default: false
func (e *GstTestSrcBin) SetExposeSourcesAsync(value bool) error {
	return e.Element.SetProperty("expose-sources-async", value)
}

func (e *GstTestSrcBin) GetExposeSourcesAsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("expose-sources-async")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// String describing the stream types to expose, eg. "video+audio".
// Default: NULL
func (e *GstTestSrcBin) SetStreamTypes(value string) error {
	return e.Element.SetProperty("stream-types", value)
}

func (e *GstTestSrcBin) GetStreamTypes() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("stream-types")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Watches for pauses in stream buffers
type GstWatchdog struct {
	*base.GstBaseTransform
}

func NewWatchdog() (*GstWatchdog, error) {
	e, err := gst.NewElement("watchdog")
	if err != nil {
		return nil, err
	}
	return &GstWatchdog{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewWatchdogWithName(name string) (*GstWatchdog, error) {
	e, err := gst.NewElementWithName("watchdog", name)
	if err != nil {
		return nil, err
	}
	return &GstWatchdog{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Timeout (in ms) after which an element error is sent to the bus if no buffers are received. 0 means disabled.
// Default: 1000, Min: 0, Max: 2147483647
func (e *GstWatchdog) SetTimeout(value int) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstWatchdog) GetTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

type GstFakeAudioSinkStateError string

const (
	GstFakeAudioSinkStateErrorNone            GstFakeAudioSinkStateError = "none"              // No state change errors
	GstFakeAudioSinkStateErrorNullToReady     GstFakeAudioSinkStateError = "null-to-ready"     // Fail state change from NULL to READY
	GstFakeAudioSinkStateErrorReadyToPaused   GstFakeAudioSinkStateError = "ready-to-paused"   // Fail state change from READY to PAUSED
	GstFakeAudioSinkStateErrorPausedToPlaying GstFakeAudioSinkStateError = "paused-to-playing" // Fail state change from PAUSED to PLAYING
	GstFakeAudioSinkStateErrorPlayingToPaused GstFakeAudioSinkStateError = "playing-to-paused" // Fail state change from PLAYING to PAUSED
	GstFakeAudioSinkStateErrorPausedToReady   GstFakeAudioSinkStateError = "paused-to-ready"   // Fail state change from PAUSED to READY
	GstFakeAudioSinkStateErrorReadyToNull     GstFakeAudioSinkStateError = "ready-to-null"     // Fail state change from READY to NULL
)

type GstFakeVideoSinkAllocationMetaFlags int

const (
	GstFakeVideoSinkAllocationMetaFlagsCrop               GstFakeVideoSinkAllocationMetaFlags = 0x00000001 // Expose the crop meta as supported
	GstFakeVideoSinkAllocationMetaFlagsOverlayComposition GstFakeVideoSinkAllocationMetaFlags = 0x00000002 // Expose the overlay composition meta as supported
)

type GstFakeVideoSinkStateError string

const (
	GstFakeVideoSinkStateErrorNone            GstFakeVideoSinkStateError = "none"              // No state change errors
	GstFakeVideoSinkStateErrorNullToReady     GstFakeVideoSinkStateError = "null-to-ready"     // Fail state change from NULL to READY
	GstFakeVideoSinkStateErrorReadyToPaused   GstFakeVideoSinkStateError = "ready-to-paused"   // Fail state change from READY to PAUSED
	GstFakeVideoSinkStateErrorPausedToPlaying GstFakeVideoSinkStateError = "paused-to-playing" // Fail state change from PAUSED to PLAYING
	GstFakeVideoSinkStateErrorPlayingToPaused GstFakeVideoSinkStateError = "playing-to-paused" // Fail state change from PLAYING to PAUSED
	GstFakeVideoSinkStateErrorPausedToReady   GstFakeVideoSinkStateError = "paused-to-ready"   // Fail state change from PAUSED to READY
	GstFakeVideoSinkStateErrorReadyToNull     GstFakeVideoSinkStateError = "ready-to-null"     // Fail state change from READY to NULL
)

type GstChecksumSinkHash string

const (
	GstChecksumSinkHashMd5    GstChecksumSinkHash = "md5"    // MD5
	GstChecksumSinkHashSha1   GstChecksumSinkHash = "sha1"   // SHA-1
	GstChecksumSinkHashSha256 GstChecksumSinkHash = "sha256" // SHA-256
	GstChecksumSinkHashSha512 GstChecksumSinkHash = "sha512" // SHA-512
)

type GstClockSelectClockId string

const (
	GstClockSelectClockIdDefault   GstClockSelectClockId = "default"   // Default (elected from elements) pipeline clock
	GstClockSelectClockIdMonotonic GstClockSelectClockId = "monotonic" // System monotonic clock
	GstClockSelectClockIdRealtime  GstClockSelectClockId = "realtime"  // System realtime clock
	GstClockSelectClockIdPtp       GstClockSelectClockId = "ptp"       // PTP clock
	GstClockSelectClockIdTai       GstClockSelectClockId = "tai"       // System TAI clock
)

type GstCompareMethod string

const (
	GstCompareMethodMem  GstCompareMethod = "mem"  // Memory
	GstCompareMethodMax  GstCompareMethod = "max"  // Maximum metric
	GstCompareMethodSsim GstCompareMethod = "ssim" // SSIM (raw video)
)
