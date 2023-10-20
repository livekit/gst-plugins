// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Set/merge caps on stream
type GstCapsSetter struct {
	*base.GstBaseTransform
}

func NewCapsSetter() (*GstCapsSetter, error) {
	e, err := gst.NewElement("capssetter")
	if err != nil {
		return nil, err
	}
	return &GstCapsSetter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCapsSetterWithName(name string) (*GstCapsSetter, error) {
	e, err := gst.NewElementWithName("capssetter", name)
	if err != nil {
		return nil, err
	}
	return &GstCapsSetter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Merge these caps (thereby overwriting) in the stream
// Default: ANY
func (e *GstCapsSetter) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstCapsSetter) GetCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// Match incoming caps' mime-type to mime-type of provided caps
// Default: true
func (e *GstCapsSetter) SetJoin(value bool) error {
	return e.Element.SetProperty("join", value)
}

func (e *GstCapsSetter) GetJoin() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("join")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Drop fields of incoming caps
// Default: false
func (e *GstCapsSetter) SetReplace(value bool) error {
	return e.Element.SetProperty("replace", value)
}

func (e *GstCapsSetter) GetReplace() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("replace")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Seek based on navigation keys left-right
type GstNavSeek struct {
	*base.GstBaseTransform
}

func NewNavSeek() (*GstNavSeek, error) {
	e, err := gst.NewElement("navseek")
	if err != nil {
		return nil, err
	}
	return &GstNavSeek{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewNavSeekWithName(name string) (*GstNavSeek, error) {
	e, err := gst.NewElementWithName("navseek", name)
	if err != nil {
		return nil, err
	}
	return &GstNavSeek{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Hold eos until the next 'Return' keystroke
// Default: false
func (e *GstNavSeek) SetHoldEos(value bool) error {
	return e.Element.SetProperty("hold-eos", value)
}

func (e *GstNavSeek) GetHoldEos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hold-eos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Time in seconds to seek by
// Default: 5, Min: 0, Max: 1.79769e+308
func (e *GstNavSeek) SetSeekOffset(value float64) error {
	return e.Element.SetProperty("seek-offset", value)
}

func (e *GstNavSeek) GetSeekOffset() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("seek-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Periodically query and report on processing progress
type GstProgressReport struct {
	*base.GstBaseTransform
}

func NewProgressReport() (*GstProgressReport, error) {
	e, err := gst.NewElement("progressreport")
	if err != nil {
		return nil, err
	}
	return &GstProgressReport{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewProgressReportWithName(name string) (*GstProgressReport, error) {
	e, err := gst.NewElementWithName("progressreport", name)
	if err != nil {
		return nil, err
	}
	return &GstProgressReport{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Use a query instead of buffer metadata to determine stream position
// Default: true
func (e *GstProgressReport) SetDoQuery(value bool) error {
	return e.Element.SetProperty("do-query", value)
}

func (e *GstProgressReport) GetDoQuery() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-query")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Format to use for the querying
// Default: auto
func (e *GstProgressReport) SetFormat(value string) error {
	return e.Element.SetProperty("format", value)
}

func (e *GstProgressReport) GetFormat() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("format")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Do not print output to stdout
// Default: false
func (e *GstProgressReport) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstProgressReport) GetSilent() (bool, error) {
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

// Number of seconds between reports when data is flowing
// Default: 5, Min: 1, Max: 2147483647
func (e *GstProgressReport) SetUpdateFreq(value int) error {
	return e.Element.SetProperty("update-freq", value)
}

func (e *GstProgressReport) GetUpdateFreq() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("update-freq")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// perform a number of tests
type GstTest struct {
	*base.GstBaseSink
}

func NewTest() (*GstTest, error) {
	e, err := gst.NewElement("testsink")
	if err != nil {
		return nil, err
	}
	return &GstTest{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewTestWithName(name string) (*GstTest, error) {
	e, err := gst.NewElementWithName("testsink", name)
	if err != nil {
		return nil, err
	}
	return &GstTest{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// md5 of processing the whole data
// Default: ---
func (e *GstTest) GetMd5() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("md5")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// average difference in usec between timestamp of next buffer and expected timestamp from analyzing last buffer
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstTest) GetTimestampDeviation() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("timestamp-deviation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// allowed average difference in usec between timestamp of next buffer and expected timestamp from analyzing last buffer
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstTest) SetAllowedTimestampDeviation(value int64) error {
	return e.Element.SetProperty("allowed-timestamp-deviation", value)
}

func (e *GstTest) GetAllowedTimestampDeviation() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("allowed-timestamp-deviation")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// number of buffers in stream
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstTest) GetBufferCount() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// expected number of buffers in stream
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstTest) SetExpectedBufferCount(value int64) error {
	return e.Element.SetProperty("expected-buffer-count", value)
}

func (e *GstTest) GetExpectedBufferCount() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("expected-buffer-count")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// expected length of stream
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstTest) SetExpectedLength(value int64) error {
	return e.Element.SetProperty("expected-length", value)
}

func (e *GstTest) GetExpectedLength() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("expected-length")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// expected md5 of processing the whole data
// Default: ---
func (e *GstTest) SetExpectedMd5(value string) error {
	return e.Element.SetProperty("expected-md5", value)
}

func (e *GstTest) GetExpectedMd5() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("expected-md5")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// length of stream
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstTest) GetLength() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("length")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// randomly change data in the stream
type GstBreakMyData struct {
	*base.GstBaseTransform
}

func NewBreakMyData() (*GstBreakMyData, error) {
	e, err := gst.NewElement("breakmydata")
	if err != nil {
		return nil, err
	}
	return &GstBreakMyData{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewBreakMyDataWithName(name string) (*GstBreakMyData, error) {
	e, err := gst.NewElementWithName("breakmydata", name)
	if err != nil {
		return nil, err
	}
	return &GstBreakMyData{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// seed for randomness (initialized when going from READY to PAUSED)
// Default: 0, Min: 0, Max: -1
func (e *GstBreakMyData) SetSeed(value uint) error {
	return e.Element.SetProperty("seed", value)
}

func (e *GstBreakMyData) GetSeed() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("seed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// set changed bytes to this value (-1 means random value
// Default: -1, Min: -1, Max: 255
func (e *GstBreakMyData) SetSetTo(value int) error {
	return e.Element.SetProperty("set-to", value)
}

func (e *GstBreakMyData) GetSetTo() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("set-to")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// amount of bytes skipped at the beginning of stream
// Default: 0, Min: 0, Max: -1
func (e *GstBreakMyData) SetSkip(value uint) error {
	return e.Element.SetProperty("skip", value)
}

func (e *GstBreakMyData) GetSkip() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("skip")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// probability for each byte in the buffer to be changed
// Default: 0, Min: 0, Max: 1
func (e *GstBreakMyData) SetProbability(value float64) error {
	return e.Element.SetProperty("probability", value)
}

func (e *GstBreakMyData) GetProbability() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("probability")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Post cpu usage information every buffer
type GstCpuReport struct {
	*base.GstBaseTransform
}

func NewCpuReport() (*GstCpuReport, error) {
	e, err := gst.NewElement("cpureport")
	if err != nil {
		return nil, err
	}
	return &GstCpuReport{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCpuReportWithName(name string) (*GstCpuReport, error) {
	e, err := gst.NewElementWithName("cpureport", name)
	if err != nil {
		return nil, err
	}
	return &GstCpuReport{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Implements pushfile:// URI-handler for push-based file access
type GstPushFileSrc struct {
	*gst.Bin
}

func NewPushFileSrc() (*GstPushFileSrc, error) {
	e, err := gst.NewElement("pushfilesrc")
	if err != nil {
		return nil, err
	}
	return &GstPushFileSrc{Bin: &gst.Bin{Element: e}}, nil
}

func NewPushFileSrcWithName(name string) (*GstPushFileSrc, error) {
	e, err := gst.NewElementWithName("pushfilesrc", name)
	if err != nil {
		return nil, err
	}
	return &GstPushFileSrc{Bin: &gst.Bin{Element: e}}, nil
}

// Initial Stream Time (if time-segment TRUE)
// Default: 0, Min: 0, Max: 9223372036854775807
func (e *GstPushFileSrc) SetStreamTime(value int64) error {
	return e.Element.SetProperty("stream-time", value)
}

func (e *GstPushFileSrc) GetStreamTime() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("stream-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Emit TIME SEGMENTS
// Default: false
func (e *GstPushFileSrc) SetTimeSegment(value bool) error {
	return e.Element.SetProperty("time-segment", value)
}

func (e *GstPushFileSrc) GetTimeSegment() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("time-segment")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Applied rate to use in TIME SEGMENT
// Default: 1, Min: 2.22507e-308, Max: 1.79769e+308
func (e *GstPushFileSrc) SetAppliedRate(value float64) error {
	return e.Element.SetProperty("applied-rate", value)
}

func (e *GstPushFileSrc) GetAppliedRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("applied-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Initial Buffer Timestamp (if time-segment TRUE)
// Default: 18446744073709551615, Min: 0, Max: 18446744073709551615
func (e *GstPushFileSrc) SetInitialTimestamp(value uint64) error {
	return e.Element.SetProperty("initial-timestamp", value)
}

func (e *GstPushFileSrc) GetInitialTimestamp() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("initial-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Location of the file to read
// Default: NULL
func (e *GstPushFileSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstPushFileSrc) GetLocation() (string, error) {
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

// Rate to use in TIME SEGMENT
// Default: 1, Min: 2.22507e-308, Max: 1.79769e+308
func (e *GstPushFileSrc) SetRate(value float64) error {
	return e.Element.SetProperty("rate", value)
}

func (e *GstPushFileSrc) GetRate() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Initial Start Time (if time-segment TRUE)
// Default: 0, Min: 0, Max: 9223372036854775807
func (e *GstPushFileSrc) SetStartTime(value int64) error {
	return e.Element.SetProperty("start-time", value)
}

func (e *GstPushFileSrc) GetStartTime() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("start-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// pull random sized buffers
type GstRndBufferSize struct {
	*gst.Element
}

func NewRndBufferSize() (*GstRndBufferSize, error) {
	e, err := gst.NewElement("rndbuffersize")
	if err != nil {
		return nil, err
	}
	return &GstRndBufferSize{Element: e}, nil
}

func NewRndBufferSizeWithName(name string) (*GstRndBufferSize, error) {
	e, err := gst.NewElementWithName("rndbuffersize", name)
	if err != nil {
		return nil, err
	}
	return &GstRndBufferSize{Element: e}, nil
}

// maximum buffer size
// Default: 8192, Min: 1, Max: 2147483647
func (e *GstRndBufferSize) SetMax(value int) error {
	return e.Element.SetProperty("max", value)
}

func (e *GstRndBufferSize) GetMax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// minimum buffer size
// Default: 1, Min: 0, Max: 2147483647
func (e *GstRndBufferSize) SetMin(value int) error {
	return e.Element.SetProperty("min", value)
}

func (e *GstRndBufferSize) GetMin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// seed for randomness (initialized when going from READY to PAUSED)
// Default: 0, Min: 0, Max: -1
func (e *GstRndBufferSize) SetSeed(value uint) error {
	return e.Element.SetProperty("seed", value)
}

func (e *GstRndBufferSize) GetSeed() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("seed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// inject metadata tags
type GstTagInject struct {
	*base.GstBaseTransform
}

func NewTagInject() (*GstTagInject, error) {
	e, err := gst.NewElement("taginject")
	if err != nil {
		return nil, err
	}
	return &GstTagInject{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewTagInjectWithName(name string) (*GstTagInject, error) {
	e, err := gst.NewElementWithName("taginject", name)
	if err != nil {
		return nil, err
	}
	return &GstTagInject{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Scope of tags to inject (stream | global)
// Default: stream (0)
func (e *GstTagInject) SetScope(value interface{}) error {
	return e.Element.SetProperty("scope", value)
}

// List of tags to inject into the target file
// Default: NULL
func (e *GstTagInject) SetTags(value string) error {
	return e.Element.SetProperty("tags", value)
}
