// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Plays audio to an ALSA device
type GstTinyalsaSink struct {
	*GstAudioSink
}

func NewTinyalsaSink() (*GstTinyalsaSink, error) {
	e, err := gst.NewElement("tinyalsasink")
	if err != nil {
		return nil, err
	}
	return &GstTinyalsaSink{GstAudioSink: &GstAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

func NewTinyalsaSinkWithName(name string) (*GstTinyalsaSink, error) {
	e, err := gst.NewElementWithName("tinyalsasink", name)
	if err != nil {
		return nil, err
	}
	return &GstTinyalsaSink{GstAudioSink: &GstAudioSink{GstAudioBaseSink: &GstAudioBaseSink{GstBaseSink: &base.GstBaseSink{Element: e}}}}, nil
}

// Enable the last-sample property
// Default: false
func (e *GstTinyalsaSink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

func (e *GstTinyalsaSink) GetEnableLastSample() (bool, error) {
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

// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)
// Default: 18446744073709551615, Min: -1, Max: 9223372036854775807
func (e *GstTinyalsaSink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

func (e *GstTinyalsaSink) GetMaxLateness() (int64, error) {
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

// Sink Statistics
// Default: application/x-gst-base-sink-stats, average-rate=(double)0, dropped=(guint64)0, rendered=(guint64)0;
func (e *GstTinyalsaSink) GetStats() (*gst.Structure, error) {
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

// Generate Quality-of-Service events upstream
// Default: false
func (e *GstTinyalsaSink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

func (e *GstTinyalsaSink) GetQos() (bool, error) {
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
func (e *GstTinyalsaSink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

func (e *GstTinyalsaSink) GetRenderDelay() (uint64, error) {
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

// Timestamp offset in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstTinyalsaSink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstTinyalsaSink) GetTsOffset() (int64, error) {
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

// Size of audio buffer in microseconds, this is the minimum latency that the sink reports
// Default: 200000, Min: 1, Max: 9223372036854775807
func (e *GstTinyalsaSink) SetBufferTime(value int64) error {
	return e.Element.SetProperty("buffer-time", value)
}

func (e *GstTinyalsaSink) GetBufferTime() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// The ALSA device to use
// Default: 0, Min: 0, Max: -1
func (e *GstTinyalsaSink) SetDevice(value uint) error {
	return e.Element.SetProperty("device", value)
}

func (e *GstTinyalsaSink) GetDevice() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("device")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Provide a clock to be used as the global pipeline clock
// Default: true
func (e *GstTinyalsaSink) SetProvideClock(value bool) error {
	return e.Element.SetProperty("provide-clock", value)
}

func (e *GstTinyalsaSink) GetProvideClock() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("provide-clock")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Window of time in nanoseconds to wait before creating a discontinuity
// Default: 1000000000, Min: 0, Max: 18446744073709551614
func (e *GstTinyalsaSink) SetDiscontWait(value uint64) error {
	return e.Element.SetProperty("discont-wait", value)
}

func (e *GstTinyalsaSink) GetDiscontWait() (uint64, error) {
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

// Tolerance for clock drift in microseconds
// Default: 40000, Min: 1, Max: 9223372036854775807
func (e *GstTinyalsaSink) SetDriftTolerance(value int64) error {
	return e.Element.SetProperty("drift-tolerance", value)
}

func (e *GstTinyalsaSink) GetDriftTolerance() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("drift-tolerance")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// The last sample received in the sink

func (e *GstTinyalsaSink) GetLastSample() (*gst.Sample, error) {
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

// The minimum amount of data to write in each iteration in microseconds
// Default: 10000, Min: 1, Max: 9223372036854775807
func (e *GstTinyalsaSink) SetLatencyTime(value int64) error {
	return e.Element.SetProperty("latency-time", value)
}

func (e *GstTinyalsaSink) GetLatencyTime() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("latency-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// The maximum bits per second to render (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstTinyalsaSink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

func (e *GstTinyalsaSink) GetMaxBitrate() (uint64, error) {
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

// Timestamp alignment threshold in nanoseconds
// Default: 40000000, Min: 1, Max: 18446744073709551614
func (e *GstTinyalsaSink) SetAlignmentThreshold(value uint64) error {
	return e.Element.SetProperty("alignment-threshold", value)
}

func (e *GstTinyalsaSink) GetAlignmentThreshold() (uint64, error) {
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

// Allow pull-based scheduling
// Default: false
func (e *GstTinyalsaSink) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

func (e *GstTinyalsaSink) GetCanActivatePull() (bool, error) {
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

// The ALSA card to use
// Default: 0, Min: 0, Max: -1
func (e *GstTinyalsaSink) SetCard(value uint) error {
	return e.Element.SetProperty("card", value)
}

func (e *GstTinyalsaSink) GetCard() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("card")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Algorithm used to match the rate of the masterclock
// Default: skew (1)
func (e *GstTinyalsaSink) SetSlaveMethod(value interface{}) error {
	return e.Element.SetProperty("slave-method", value)
}

func (e *GstTinyalsaSink) GetSlaveMethod() (interface{}, error) {
	return e.Element.GetProperty("slave-method")
}

// Sync on the clock
// Default: true
func (e *GstTinyalsaSink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

func (e *GstTinyalsaSink) GetSync() (bool, error) {
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

// The time to keep between rendered buffers (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstTinyalsaSink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

func (e *GstTinyalsaSink) GetThrottleTime() (uint64, error) {
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

// Go asynchronously to PAUSED
// Default: true
func (e *GstTinyalsaSink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

func (e *GstTinyalsaSink) GetAsync() (bool, error) {
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
func (e *GstTinyalsaSink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

func (e *GstTinyalsaSink) GetBlocksize() (uint, error) {
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

// Maximum processing time for a buffer in nanoseconds
// Default: 20000000, Min: 0, Max: 18446744073709551615
func (e *GstTinyalsaSink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

func (e *GstTinyalsaSink) GetProcessingDeadline() (uint64, error) {
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
