// source: gstreamer

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Simple data queue
type GstQueue struct {
	*gst.Element
}

func NewQueue() (*GstQueue, error) {
	e, err := gst.NewElement("queue")
	if err != nil {
		return nil, err
	}
	return &GstQueue{Element: e}, nil
}

func NewQueueWithName(name string) (*GstQueue, error) {
	e, err := gst.NewElementWithName("queue", name)
	if err != nil {
		return nil, err
	}
	return &GstQueue{Element: e}, nil
}

// Min. amount of data in the queue to allow reading (in ns, 0=disable)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstQueue) SetMinThresholdTime(value uint64) error {
	return e.Element.SetProperty("min-threshold-time", value)
}

func (e *GstQueue) GetMinThresholdTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("min-threshold-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Current amount of data in the queue (bytes)
// Default: 0, Min: 0, Max: -1
func (e *GstQueue) GetCurrentLevelBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-level-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Current amount of data in the queue (in ns)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstQueue) GetCurrentLevelTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("current-level-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Discard all data in the queue when an EOS event is received
// Default: false
func (e *GstQueue) SetFlushOnEos(value bool) error {
	return e.Element.SetProperty("flush-on-eos", value)
}

func (e *GstQueue) GetFlushOnEos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("flush-on-eos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Max. amount of data in the queue (bytes, 0=disable)
// Default: 10485760, Min: 0, Max: -1
func (e *GstQueue) SetMaxSizeBytes(value uint) error {
	return e.Element.SetProperty("max-size-bytes", value)
}

func (e *GstQueue) GetMaxSizeBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max. amount of data in the queue (in ns, 0=disable)
// Default: 1000000000, Min: 0, Max: 18446744073709551615
func (e *GstQueue) SetMaxSizeTime(value uint64) error {
	return e.Element.SetProperty("max-size-time", value)
}

func (e *GstQueue) GetMaxSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Min. number of buffers in the queue to allow reading (0=disable)
// Default: 0, Min: 0, Max: -1
func (e *GstQueue) SetMinThresholdBuffers(value uint) error {
	return e.Element.SetProperty("min-threshold-buffers", value)
}

func (e *GstQueue) GetMinThresholdBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-threshold-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Min. amount of data in the queue to allow reading (bytes, 0=disable)
// Default: 0, Min: 0, Max: -1
func (e *GstQueue) SetMinThresholdBytes(value uint) error {
	return e.Element.SetProperty("min-threshold-bytes", value)
}

func (e *GstQueue) GetMinThresholdBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("min-threshold-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Don't emit queue signals
// Default: false
func (e *GstQueue) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstQueue) GetSilent() (bool, error) {
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

// Current number of buffers in the queue
// Default: 0, Min: 0, Max: -1
func (e *GstQueue) GetCurrentLevelBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-level-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Where the queue leaks, if at all
// Default: no (0)
func (e *GstQueue) SetLeaky(value GstQueueLeaky) error {
	e.Element.SetArg("leaky", string(value))
	return nil
}

func (e *GstQueue) GetLeaky() (GstQueueLeaky, error) {
	var value GstQueueLeaky
	var ok bool
	v, err := e.Element.GetProperty("leaky")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQueueLeaky)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQueueLeaky")
	}
	return value, nil
}

// Max. number of buffers in the queue (0=disable)
// Default: 200, Min: 0, Max: -1
func (e *GstQueue) SetMaxSizeBuffers(value uint) error {
	return e.Element.SetProperty("max-size-buffers", value)
}

func (e *GstQueue) GetMaxSizeBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// N-to-1 pipe fitting
type GstFunnel struct {
	*gst.Element
}

func NewFunnel() (*GstFunnel, error) {
	e, err := gst.NewElement("funnel")
	if err != nil {
		return nil, err
	}
	return &GstFunnel{Element: e}, nil
}

func NewFunnelWithName(name string) (*GstFunnel, error) {
	e, err := gst.NewElementWithName("funnel", name)
	if err != nil {
		return nil, err
	}
	return &GstFunnel{Element: e}, nil
}

// Forward sticky events on stream changes
// Default: true
func (e *GstFunnel) SetForwardStickyEvents(value bool) error {
	return e.Element.SetProperty("forward-sticky-events", value)
}

func (e *GstFunnel) GetForwardStickyEvents() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("forward-sticky-events")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Read from arbitrary point in a file
type GstFileSrc struct {
	*base.GstBaseSrc
}

func NewFileSrc() (*GstFileSrc, error) {
	e, err := gst.NewElement("filesrc")
	if err != nil {
		return nil, err
	}
	return &GstFileSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewFileSrcWithName(name string) (*GstFileSrc, error) {
	e, err := gst.NewElementWithName("filesrc", name)
	if err != nil {
		return nil, err
	}
	return &GstFileSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Location of the file to read
// Default: NULL
func (e *GstFileSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstFileSrc) GetLocation() (string, error) {
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

// Pass data without modification
type GstIdentity struct {
	*base.GstBaseTransform
}

func NewIdentity() (*GstIdentity, error) {
	e, err := gst.NewElement("identity")
	if err != nil {
		return nil, err
	}
	return &GstIdentity{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewIdentityWithName(name string) (*GstIdentity, error) {
	e, err := gst.NewElementWithName("identity", name)
	if err != nil {
		return nil, err
	}
	return &GstIdentity{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// The Probability a buffer is dropped
// Default: 0, Min: 0, Max: 1
func (e *GstIdentity) SetDropProbability(value float32) error {
	return e.Element.SetProperty("drop-probability", value)
}

func (e *GstIdentity) GetDropProbability() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("drop-probability")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// Dump buffer contents to stdout
// Default: false
func (e *GstIdentity) SetDump(value bool) error {
	return e.Element.SetProperty("dump", value)
}

func (e *GstIdentity) GetDump() (bool, error) {
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

// Statistics
// Default: application/x-identity-stats, num-bytes=(guint64)0, num-buffers=(guint64)0;
func (e *GstIdentity) GetStats() (*gst.Structure, error) {
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

// Send a signal before pushing the buffer
// Default: true
func (e *GstIdentity) SetSignalHandoffs(value bool) error {
	return e.Element.SetProperty("signal-handoffs", value)
}

func (e *GstIdentity) GetSignalHandoffs() (bool, error) {
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

// silent
// Default: true
func (e *GstIdentity) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstIdentity) GetSilent() (bool, error) {
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

// Timestamp buffers and eat segments so as to appear as one segment
// Default: false
func (e *GstIdentity) SetSingleSegment(value bool) error {
	return e.Element.SetProperty("single-segment", value)
}

func (e *GstIdentity) GetSingleSegment() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("single-segment")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Send element messages if timestamps and durations do not match up
// Default: false
func (e *GstIdentity) SetCheckImperfectTimestamp(value bool) error {
	return e.Element.SetProperty("check-imperfect-timestamp", value)
}

func (e *GstIdentity) GetCheckImperfectTimestamp() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("check-imperfect-timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// (Re)timestamps buffers with number of bytes per second (0 = inactive)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstIdentity) SetDatarate(value int) error {
	return e.Element.SetProperty("datarate", value)
}

func (e *GstIdentity) GetDatarate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("datarate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Don't forward allocation queries
// Default: false
func (e *GstIdentity) SetDropAllocation(value bool) error {
	return e.Element.SetProperty("drop-allocation", value)
}

func (e *GstIdentity) GetDropAllocation() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-allocation")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Error after N buffers
// Default: -1, Min: -1, Max: 2147483647
func (e *GstIdentity) SetErrorAfter(value int) error {
	return e.Element.SetProperty("error-after", value)
}

func (e *GstIdentity) GetErrorAfter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("error-after")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// last-message
// Default: NULL
func (e *GstIdentity) GetLastMessage() (string, error) {
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

// Microseconds to sleep between processing
// Default: 0, Min: 0, Max: -1
func (e *GstIdentity) SetSleepTime(value uint) error {
	return e.Element.SetProperty("sleep-time", value)
}

func (e *GstIdentity) GetSleepTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sleep-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Send element messages if offset and offset_end do not match up
// Default: false
func (e *GstIdentity) SetCheckImperfectOffset(value bool) error {
	return e.Element.SetProperty("check-imperfect-offset", value)
}

func (e *GstIdentity) GetCheckImperfectOffset() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("check-imperfect-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Drop buffers with the given flags
// Default: (none)
func (e *GstIdentity) SetDropBufferFlags(value *gst.BufferFlags) error {
	return e.Element.SetProperty("drop-buffer-flags", value)
}

func (e *GstIdentity) GetDropBufferFlags() (*gst.BufferFlags, error) {
	var value *gst.BufferFlags
	var ok bool
	v, err := e.Element.GetProperty("drop-buffer-flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.BufferFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.BufferFlags")
	}
	return value, nil
}

// EOS after N buffers
// Default: -1, Min: -1, Max: 2147483647
func (e *GstIdentity) SetEosAfter(value int) error {
	return e.Element.SetProperty("eos-after", value)
}

func (e *GstIdentity) GetEosAfter() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("eos-after")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Synchronize to pipeline clock
// Default: false
func (e *GstIdentity) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

func (e *GstIdentity) GetSync() (bool, error) {
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

// Timestamp offset in nanoseconds for synchronisation, negative for earlier sync
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstIdentity) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstIdentity) GetTsOffset() (int64, error) {
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

// Multiple data queue
type GstMultiQueue struct {
	*gst.Element
}

func NewMultiQueue() (*GstMultiQueue, error) {
	e, err := gst.NewElement("multiqueue")
	if err != nil {
		return nil, err
	}
	return &GstMultiQueue{Element: e}, nil
}

func NewMultiQueueWithName(name string) (*GstMultiQueue, error) {
	e, err := gst.NewElementWithName("multiqueue", name)
	if err != nil {
		return nil, err
	}
	return &GstMultiQueue{Element: e}, nil
}

// High threshold for buffering to finish. Only used if use-buffering is True
// Default: 0.99, Min: 0, Max: 1
func (e *GstMultiQueue) SetHighWatermark(value float64) error {
	return e.Element.SetProperty("high-watermark", value)
}

func (e *GstMultiQueue) GetHighWatermark() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("high-watermark")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Low threshold for buffering to start. Only used if use-buffering is True
// Default: 0.01, Min: 0, Max: 1
func (e *GstMultiQueue) SetLowWatermark(value float64) error {
	return e.Element.SetProperty("low-watermark", value)
}

func (e *GstMultiQueue) GetLowWatermark() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("low-watermark")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Max. amount of data in the queue (bytes, 0=disable)
// Default: 10485760, Min: 0, Max: -1
func (e *GstMultiQueue) SetMaxSizeBytes(value uint) error {
	return e.Element.SetProperty("max-size-bytes", value)
}

func (e *GstMultiQueue) GetMaxSizeBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Minimum extra buffering for deinterleaving (size of the queues) when use-interleave=true
// Default: 250000000, Min: 0, Max: 18446744073709551615
func (e *GstMultiQueue) SetMinInterleaveTime(value uint64) error {
	return e.Element.SetProperty("min-interleave-time", value)
}

func (e *GstMultiQueue) GetMinInterleaveTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("min-interleave-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// High threshold for buffering to finish. Only used if use-buffering is True (Deprecated: use high-watermark instead)
// Default: 99, Min: 0, Max: 100
func (e *GstMultiQueue) SetHighPercent(value int) error {
	return e.Element.SetProperty("high-percent", value)
}

func (e *GstMultiQueue) GetHighPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("high-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Max. number of buffers in the queue (0=disable)
// Default: 5, Min: 0, Max: -1
func (e *GstMultiQueue) SetMaxSizeBuffers(value uint) error {
	return e.Element.SetProperty("max-size-buffers", value)
}

func (e *GstMultiQueue) GetMaxSizeBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Synchronize deactivated or not-linked streams by running time
// Default: false
func (e *GstMultiQueue) SetSyncByRunningTime(value bool) error {
	return e.Element.SetProperty("sync-by-running-time", value)
}

func (e *GstMultiQueue) GetSyncByRunningTime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync-by-running-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Adjust time limits based on input interleave
// Default: false
func (e *GstMultiQueue) SetUseInterleave(value bool) error {
	return e.Element.SetProperty("use-interleave", value)
}

func (e *GstMultiQueue) GetUseInterleave() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-interleave")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Amount of buffers the queues can grow if one of them is empty (0=disable) (NOT IMPLEMENTED)
// Default: 5, Min: 0, Max: -1
func (e *GstMultiQueue) SetExtraSizeBuffers(value uint) error {
	return e.Element.SetProperty("extra-size-buffers", value)
}

func (e *GstMultiQueue) GetExtraSizeBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("extra-size-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Amount of time the queues can grow if one of them is empty (in ns, 0=disable) (NOT IMPLEMENTED)
// Default: 3000000000, Min: 0, Max: 18446744073709551615
func (e *GstMultiQueue) SetExtraSizeTime(value uint64) error {
	return e.Element.SetProperty("extra-size-time", value)
}

func (e *GstMultiQueue) GetExtraSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("extra-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Low threshold for buffering to start. Only used if use-buffering is True (Deprecated: use low-watermark instead)
// Default: 1, Min: 0, Max: 100
func (e *GstMultiQueue) SetLowPercent(value int) error {
	return e.Element.SetProperty("low-percent", value)
}

func (e *GstMultiQueue) GetLowPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("low-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Extra buffering in time for unlinked streams (if 'sync-by-running-time')
// Default: 250000000, Min: 0, Max: 18446744073709551615
func (e *GstMultiQueue) SetUnlinkedCacheTime(value uint64) error {
	return e.Element.SetProperty("unlinked-cache-time", value)
}

func (e *GstMultiQueue) GetUnlinkedCacheTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("unlinked-cache-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Amount of data the queues can grow if one of them is empty (bytes, 0=disable) (NOT IMPLEMENTED)
// Default: 10485760, Min: 0, Max: -1
func (e *GstMultiQueue) SetExtraSizeBytes(value uint) error {
	return e.Element.SetProperty("extra-size-bytes", value)
}

func (e *GstMultiQueue) GetExtraSizeBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("extra-size-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max. amount of data in the queue (in ns, 0=disable)
// Default: 2000000000, Min: 0, Max: 18446744073709551615
func (e *GstMultiQueue) SetMaxSizeTime(value uint64) error {
	return e.Element.SetProperty("max-size-time", value)
}

func (e *GstMultiQueue) GetMaxSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Multiqueue Statistics
// Default: application/x-gst-multi-queue-stats;
func (e *GstMultiQueue) GetStats() (*gst.Structure, error) {
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

// Emit GST_MESSAGE_BUFFERING based on low-/high-percent thresholds (0%% = low-watermark, 100%% = high-watermark)
// Default: false
func (e *GstMultiQueue) SetUseBuffering(value bool) error {
	return e.Element.SetProperty("use-buffering", value)
}

func (e *GstMultiQueue) GetUseBuffering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-buffering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// 1-to-N output stream selector
type GstOutputSelector struct {
	*gst.Element
}

func NewOutputSelector() (*GstOutputSelector, error) {
	e, err := gst.NewElement("output-selector")
	if err != nil {
		return nil, err
	}
	return &GstOutputSelector{Element: e}, nil
}

func NewOutputSelectorWithName(name string) (*GstOutputSelector, error) {
	e, err := gst.NewElementWithName("output-selector", name)
	if err != nil {
		return nil, err
	}
	return &GstOutputSelector{Element: e}, nil
}

// Resend latest buffer after a switch to a new pad
// Default: false
func (e *GstOutputSelector) SetResendLatest(value bool) error {
	return e.Element.SetProperty("resend-latest", value)
}

func (e *GstOutputSelector) GetResendLatest() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("resend-latest")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Currently active src pad

func (e *GstOutputSelector) SetActivePad(value *gst.Pad) error {
	return e.Element.SetProperty("active-pad", value)
}

func (e *GstOutputSelector) GetActivePad() (*gst.Pad, error) {
	var value *gst.Pad
	var ok bool
	v, err := e.Element.GetProperty("active-pad")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Pad)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Pad")
	}
	return value, nil
}

// The mode to be used for pad negotiation
// Default: all (1)
func (e *GstOutputSelector) SetPadNegotiationMode(value GstOutputSelectorPadNegotiationMode) error {
	e.Element.SetArg("pad-negotiation-mode", string(value))
	return nil
}

func (e *GstOutputSelector) GetPadNegotiationMode() (GstOutputSelectorPadNegotiationMode, error) {
	var value GstOutputSelectorPadNegotiationMode
	var ok bool
	v, err := e.Element.GetProperty("pad-negotiation-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstOutputSelectorPadNegotiationMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstOutputSelectorPadNegotiationMode")
	}
	return value, nil
}

// 1-to-N output stream by stream-id
type GstStreamidDemux struct {
	*gst.Element
}

func NewStreamidDemux() (*GstStreamidDemux, error) {
	e, err := gst.NewElement("streamiddemux")
	if err != nil {
		return nil, err
	}
	return &GstStreamidDemux{Element: e}, nil
}

func NewStreamidDemuxWithName(name string) (*GstStreamidDemux, error) {
	e, err := gst.NewElementWithName("streamiddemux", name)
	if err != nil {
		return nil, err
	}
	return &GstStreamidDemux{Element: e}, nil
}

// The currently active src pad

func (e *GstStreamidDemux) GetActivePad() (*gst.Pad, error) {
	var value *gst.Pad
	var ok bool
	v, err := e.Element.GetProperty("active-pad")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Pad)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Pad")
	}
	return value, nil
}

// Push empty (no data) buffers around
type GstFakeSrc struct {
	*base.GstBaseSrc
}

func NewFakeSrc() (*GstFakeSrc, error) {
	e, err := gst.NewElement("fakesrc")
	if err != nil {
		return nil, err
	}
	return &GstFakeSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewFakeSrcWithName(name string) (*GstFakeSrc, error) {
	e, err := gst.NewElementWithName("fakesrc", name)
	if err != nil {
		return nil, err
	}
	return &GstFakeSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// Don't produce last_message events
// Default: true
func (e *GstFakeSrc) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstFakeSrc) GetSilent() (bool, error) {
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

// How to determine buffer sizes
// Default: empty (1)
func (e *GstFakeSrc) SetSizetype(value GstFakeSrcSizeType) error {
	e.Element.SetArg("sizetype", string(value))
	return nil
}

func (e *GstFakeSrc) GetSizetype() (GstFakeSrcSizeType, error) {
	var value GstFakeSrcSizeType
	var ok bool
	v, err := e.Element.GetProperty("sizetype")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeSrcSizeType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeSrcSizeType")
	}
	return value, nil
}

// Can activate in push mode
// Default: true
func (e *GstFakeSrc) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

func (e *GstFakeSrc) GetCanActivatePush() (bool, error) {
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

// Data allocation method
// Default: allocate (1)
func (e *GstFakeSrc) SetData(value GstFakeSrcDataType) error {
	e.Element.SetArg("data", string(value))
	return nil
}

func (e *GstFakeSrc) GetData() (GstFakeSrcDataType, error) {
	var value GstFakeSrcDataType
	var ok bool
	v, err := e.Element.GetProperty("data")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeSrcDataType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeSrcDataType")
	}
	return value, nil
}

// Timestamps buffers with number of bytes per second (0 = none)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstFakeSrc) SetDatarate(value int) error {
	return e.Element.SetProperty("datarate", value)
}

func (e *GstFakeSrc) GetDatarate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("datarate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The format of the segment events
// Default: bytes (2)
func (e *GstFakeSrc) SetFormat(value *gst.Format) error {
	return e.Element.SetProperty("format", value)
}

func (e *GstFakeSrc) GetFormat() (*gst.Format, error) {
	var value *gst.Format
	var ok bool
	v, err := e.Element.GetProperty("format")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Format)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Format")
	}
	return value, nil
}

// Set the pattern (unused)
// Default: NULL
func (e *GstFakeSrc) SetPattern(value string) error {
	return e.Element.SetProperty("pattern", value)
}

func (e *GstFakeSrc) GetPattern() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pattern")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The last status message
// Default: NULL
func (e *GstFakeSrc) GetLastMessage() (string, error) {
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

// Size of parent buffer for sub-buffered allocation
// Default: 40960, Min: 0, Max: 2147483647
func (e *GstFakeSrc) SetParentsize(value int) error {
	return e.Element.SetProperty("parentsize", value)
}

func (e *GstFakeSrc) GetParentsize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("parentsize")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Send a signal before pushing the buffer
// Default: false
func (e *GstFakeSrc) SetSignalHandoffs(value bool) error {
	return e.Element.SetProperty("signal-handoffs", value)
}

func (e *GstFakeSrc) GetSignalHandoffs() (bool, error) {
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

// Maximum buffer size
// Default: 4096, Min: 0, Max: 2147483647
func (e *GstFakeSrc) SetSizemax(value int) error {
	return e.Element.SetProperty("sizemax", value)
}

func (e *GstFakeSrc) GetSizemax() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sizemax")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Can activate in pull mode
// Default: true
func (e *GstFakeSrc) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

func (e *GstFakeSrc) GetCanActivatePull() (bool, error) {
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

// Dump buffer contents to stdout
// Default: false
func (e *GstFakeSrc) SetDump(value bool) error {
	return e.Element.SetProperty("dump", value)
}

func (e *GstFakeSrc) GetDump() (bool, error) {
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

// How to fill the buffer, if at all
// Default: nothing (1)
func (e *GstFakeSrc) SetFilltype(value GstFakeSrcFillType) error {
	e.Element.SetArg("filltype", string(value))
	return nil
}

func (e *GstFakeSrc) GetFilltype() (GstFakeSrcFillType, error) {
	var value GstFakeSrcFillType
	var ok bool
	v, err := e.Element.GetProperty("filltype")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeSrcFillType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeSrcFillType")
	}
	return value, nil
}

// True if the element cannot produce data in PAUSED
// Default: false
func (e *GstFakeSrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

func (e *GstFakeSrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Minimum buffer size
// Default: 0, Min: 0, Max: 2147483647
func (e *GstFakeSrc) SetSizemin(value int) error {
	return e.Element.SetProperty("sizemin", value)
}

func (e *GstFakeSrc) GetSizemin() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sizemin")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Sync to the clock to the datarate
// Default: false
func (e *GstFakeSrc) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

func (e *GstFakeSrc) GetSync() (bool, error) {
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

// Black hole for data
type GstFakeSink struct {
	*base.GstBaseSink
}

func NewFakeSink() (*GstFakeSink, error) {
	e, err := gst.NewElement("fakesink")
	if err != nil {
		return nil, err
	}
	return &GstFakeSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewFakeSinkWithName(name string) (*GstFakeSink, error) {
	e, err := gst.NewElementWithName("fakesink", name)
	if err != nil {
		return nil, err
	}
	return &GstFakeSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Generate a state change error
// Default: none (0)
func (e *GstFakeSink) SetStateError(value GstFakeSinkStateError) error {
	e.Element.SetArg("state-error", string(value))
	return nil
}

func (e *GstFakeSink) GetStateError() (GstFakeSinkStateError, error) {
	var value GstFakeSinkStateError
	var ok bool
	v, err := e.Element.GetProperty("state-error")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFakeSinkStateError)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFakeSinkStateError")
	}
	return value, nil
}

// Can activate in push mode
// Default: true
func (e *GstFakeSink) SetCanActivatePush(value bool) error {
	return e.Element.SetProperty("can-activate-push", value)
}

func (e *GstFakeSink) GetCanActivatePush() (bool, error) {
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

// Dump buffer contents to stdout
// Default: false
func (e *GstFakeSink) SetDump(value bool) error {
	return e.Element.SetProperty("dump", value)
}

func (e *GstFakeSink) GetDump() (bool, error) {
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

// Number of buffers to accept going EOS
// Default: -1, Min: -1, Max: 2147483647
func (e *GstFakeSink) SetNumBuffers(value int) error {
	return e.Element.SetProperty("num-buffers", value)
}

func (e *GstFakeSink) GetNumBuffers() (int, error) {
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

// Send a signal before unreffing the buffer
// Default: false
func (e *GstFakeSink) SetSignalHandoffs(value bool) error {
	return e.Element.SetProperty("signal-handoffs", value)
}

func (e *GstFakeSink) GetSignalHandoffs() (bool, error) {
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
func (e *GstFakeSink) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstFakeSink) GetSilent() (bool, error) {
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

// Can activate in pull mode
// Default: false
func (e *GstFakeSink) SetCanActivatePull(value bool) error {
	return e.Element.SetProperty("can-activate-pull", value)
}

func (e *GstFakeSink) GetCanActivatePull() (bool, error) {
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
func (e *GstFakeSink) SetDropOutOfSegment(value bool) error {
	return e.Element.SetProperty("drop-out-of-segment", value)
}

func (e *GstFakeSink) GetDropOutOfSegment() (bool, error) {
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
func (e *GstFakeSink) GetLastMessage() (string, error) {
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

// Write data to a file descriptor
type GstFdSink struct {
	*base.GstBaseSink
}

func NewFdSink() (*GstFdSink, error) {
	e, err := gst.NewElement("fdsink")
	if err != nil {
		return nil, err
	}
	return &GstFdSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewFdSinkWithName(name string) (*GstFdSink, error) {
	e, err := gst.NewElementWithName("fdsink", name)
	if err != nil {
		return nil, err
	}
	return &GstFdSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// An open file descriptor to write to
// Default: 1, Min: 0, Max: 2147483647
func (e *GstFdSink) SetFd(value int) error {
	return e.Element.SetProperty("fd", value)
}

func (e *GstFdSink) GetFd() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fd")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// 1-to-N pipe fitting
type GstTee struct {
	*gst.Element
}

func NewTee() (*GstTee, error) {
	e, err := gst.NewElement("tee")
	if err != nil {
		return nil, err
	}
	return &GstTee{Element: e}, nil
}

func NewTeeWithName(name string) (*GstTee, error) {
	e, err := gst.NewElementWithName("tee", name)
	if err != nil {
		return nil, err
	}
	return &GstTee{Element: e}, nil
}

// The pad ALLOCATION queries will be proxied to (DEPRECATED, has no effect)

func (e *GstTee) SetAllocPad(value *gst.Pad) error {
	return e.Element.SetProperty("alloc-pad", value)
}

func (e *GstTee) GetAllocPad() (*gst.Pad, error) {
	var value *gst.Pad
	var ok bool
	v, err := e.Element.GetProperty("alloc-pad")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Pad)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Pad")
	}
	return value, nil
}

// Return GST_FLOW_OK even if there are no source pads or they are all unlinked
// Default: false
func (e *GstTee) SetAllowNotLinked(value bool) error {
	return e.Element.SetProperty("allow-not-linked", value)
}

func (e *GstTee) GetAllowNotLinked() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-not-linked")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// If the element can operate in push mode
// Default: true
func (e *GstTee) SetHasChain(value bool) error {
	return e.Element.SetProperty("has-chain", value)
}

func (e *GstTee) GetHasChain() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("has-chain")
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
func (e *GstTee) GetLastMessage() (string, error) {
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

// The number of source pads
// Default: 0, Min: 0, Max: 2147483647
func (e *GstTee) GetNumSrcPads() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("num-src-pads")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Behavior of tee in pull mode
// Default: never (0)
func (e *GstTee) SetPullMode(value GstTeePullMode) error {
	e.Element.SetArg("pull-mode", string(value))
	return nil
}

func (e *GstTee) GetPullMode() (GstTeePullMode, error) {
	var value GstTeePullMode
	var ok bool
	v, err := e.Element.GetProperty("pull-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstTeePullMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstTeePullMode")
	}
	return value, nil
}

// Don't produce last_message events
// Default: true
func (e *GstTee) SetSilent(value bool) error {
	return e.Element.SetProperty("silent", value)
}

func (e *GstTee) GetSilent() (bool, error) {
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

// Concatenate multiple streams
type GstConcat struct {
	*gst.Element
}

func NewConcat() (*GstConcat, error) {
	e, err := gst.NewElement("concat")
	if err != nil {
		return nil, err
	}
	return &GstConcat{Element: e}, nil
}

func NewConcatWithName(name string) (*GstConcat, error) {
	e, err := gst.NewElementWithName("concat", name)
	if err != nil {
		return nil, err
	}
	return &GstConcat{Element: e}, nil
}

// Currently active sink pad

func (e *GstConcat) GetActivePad() (*gst.Pad, error) {
	var value *gst.Pad
	var ok bool
	v, err := e.Element.GetProperty("active-pad")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Pad)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Pad")
	}
	return value, nil
}

// Adjust the base value of segments to ensure they are adjacent
// Default: true
func (e *GstConcat) SetAdjustBase(value bool) error {
	return e.Element.SetProperty("adjust-base", value)
}

func (e *GstConcat) GetAdjustBase() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("adjust-base")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Synchronise buffers to the clock
type GstClockSync struct {
	*gst.Element
}

func NewClockSync() (*GstClockSync, error) {
	e, err := gst.NewElement("clocksync")
	if err != nil {
		return nil, err
	}
	return &GstClockSync{Element: e}, nil
}

func NewClockSyncWithName(name string) (*GstClockSync, error) {
	e, err := gst.NewElementWithName("clocksync", name)
	if err != nil {
		return nil, err
	}
	return &GstClockSync{Element: e}, nil
}

// Generate Quality-of-Service events upstream
// Default: false
func (e *GstClockSync) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

func (e *GstClockSync) GetQos() (bool, error) {
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

// Synchronize to pipeline clock
// Default: true
func (e *GstClockSync) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

func (e *GstClockSync) GetSync() (bool, error) {
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

// Automatically set ts-offset based on running time of the first buffer and pipeline's running time (i.e., ts-offset = "pipeline running time" - "buffer running time"). When enabled, clocksync element will update ts-offset on the first buffer per flush event or READY to PAUSED state change. This property can be useful in case that buffer timestamp does not necessarily have to be synchronized with pipeline's running time, but duration of the buffer through clocksync element needs to be synchronized with the amount of clock time go. Note that mixed use of ts-offset and this property would be racy if clocksync element is running already.
// Default: false
func (e *GstClockSync) SetSyncToFirst(value bool) error {
	return e.Element.SetProperty("sync-to-first", value)
}

func (e *GstClockSync) GetSyncToFirst() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync-to-first")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Timestamp offset in nanoseconds for synchronisation, negative for earlier sync
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstClockSync) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstClockSync) GetTsOffset() (int64, error) {
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

// Handles data: uris
type GstDataURISrc struct {
	*base.GstBaseSrc
}

func NewDataURISrc() (*GstDataURISrc, error) {
	e, err := gst.NewElement("dataurisrc")
	if err != nil {
		return nil, err
	}
	return &GstDataURISrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

func NewDataURISrcWithName(name string) (*GstDataURISrc, error) {
	e, err := gst.NewElementWithName("dataurisrc", name)
	if err != nil {
		return nil, err
	}
	return &GstDataURISrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}, nil
}

// URI that should be used
// Default: NULL
func (e *GstDataURISrc) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

func (e *GstDataURISrc) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Download Buffer element
type GstDownloadBuffer struct {
	*gst.Element
}

func NewDownloadBuffer() (*GstDownloadBuffer, error) {
	e, err := gst.NewElement("downloadbuffer")
	if err != nil {
		return nil, err
	}
	return &GstDownloadBuffer{Element: e}, nil
}

func NewDownloadBufferWithName(name string) (*GstDownloadBuffer, error) {
	e, err := gst.NewElementWithName("downloadbuffer", name)
	if err != nil {
		return nil, err
	}
	return &GstDownloadBuffer{Element: e}, nil
}

// Location to store temporary files in (Only read this property, use temp-template to configure the name template)
// Default: NULL
func (e *GstDownloadBuffer) GetTempLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("temp-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Remove the temp-location after use
// Default: true
func (e *GstDownloadBuffer) SetTempRemove(value bool) error {
	return e.Element.SetProperty("temp-remove", value)
}

func (e *GstDownloadBuffer) GetTempRemove() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temp-remove")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// File template to store temporary files in, should contain directory and XXXXXX. (NULL == disabled)
// Default: NULL
func (e *GstDownloadBuffer) SetTempTemplate(value string) error {
	return e.Element.SetProperty("temp-template", value)
}

func (e *GstDownloadBuffer) GetTempTemplate() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("temp-template")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// High threshold for buffering to finish. Emits GST_MESSAGE_BUFFERING with a value of 100%%
// Default: 99, Min: 0, Max: 100
func (e *GstDownloadBuffer) SetHighPercent(value int) error {
	return e.Element.SetProperty("high-percent", value)
}

func (e *GstDownloadBuffer) GetHighPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("high-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Low threshold for buffering to start. Emits GST_MESSAGE_BUFFERING with a value of 0%%
// Default: 10, Min: 0, Max: 100
func (e *GstDownloadBuffer) SetLowPercent(value int) error {
	return e.Element.SetProperty("low-percent", value)
}

func (e *GstDownloadBuffer) GetLowPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("low-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Max. amount of data to buffer (bytes, 0=disable)
// Default: 2097152, Min: 0, Max: -1
func (e *GstDownloadBuffer) SetMaxSizeBytes(value uint) error {
	return e.Element.SetProperty("max-size-bytes", value)
}

func (e *GstDownloadBuffer) GetMaxSizeBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max. amount of data to buffer (in ns, 0=disable)
// Default: 2000000000, Min: 0, Max: 18446744073709551615
func (e *GstDownloadBuffer) SetMaxSizeTime(value uint64) error {
	return e.Element.SetProperty("max-size-time", value)
}

func (e *GstDownloadBuffer) GetMaxSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Read from a file descriptor
type GstFdSrc struct {
	*base.GstPushSrc
}

func NewFdSrc() (*GstFdSrc, error) {
	e, err := gst.NewElement("fdsrc")
	if err != nil {
		return nil, err
	}
	return &GstFdSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewFdSrcWithName(name string) (*GstFdSrc, error) {
	e, err := gst.NewElementWithName("fdsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstFdSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// An open file descriptor to read from
// Default: -1, Min: 0, Max: 2147483647
func (e *GstFdSrc) SetFd(value int) error {
	return e.Element.SetProperty("fd", value)
}

func (e *GstFdSrc) GetFd() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("fd")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Post a message after timeout microseconds (0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstFdSrc) SetTimeout(value uint64) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstFdSrc) GetTimeout() (uint64, error) {
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

// Write stream to a file
type GstFileSink struct {
	*base.GstBaseSink
}

func NewFileSink() (*GstFileSink, error) {
	e, err := gst.NewElement("filesink")
	if err != nil {
		return nil, err
	}
	return &GstFileSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewFileSinkWithName(name string) (*GstFileSink, error) {
	e, err := gst.NewElementWithName("filesink", name)
	if err != nil {
		return nil, err
	}
	return &GstFileSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Retry up to this many ms on transient errors (currently EACCES)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstFileSink) SetMaxTransientErrorTimeout(value int) error {
	return e.Element.SetProperty("max-transient-error-timeout", value)
}

func (e *GstFileSink) GetMaxTransientErrorTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-transient-error-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Open the file with O_SYNC for enabling synchronous IO
// Default: false
func (e *GstFileSink) SetOSync(value bool) error {
	return e.Element.SetProperty("o-sync", value)
}

func (e *GstFileSink) GetOSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("o-sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Append to an already existing file
// Default: false
func (e *GstFileSink) SetAppend(value bool) error {
	return e.Element.SetProperty("append", value)
}

func (e *GstFileSink) GetAppend() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("append")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The buffering mode to use
// Default: default (-1)
func (e *GstFileSink) SetBufferMode(value GstFileSinkBufferMode) error {
	e.Element.SetArg("buffer-mode", string(value))
	return nil
}

func (e *GstFileSink) GetBufferMode() (GstFileSinkBufferMode, error) {
	var value GstFileSinkBufferMode
	var ok bool
	v, err := e.Element.GetProperty("buffer-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFileSinkBufferMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFileSinkBufferMode")
	}
	return value, nil
}

// Size of buffer in number of bytes for line or full buffer-mode
// Default: 65536, Min: 0, Max: -1
func (e *GstFileSink) SetBufferSize(value uint) error {
	return e.Element.SetProperty("buffer-size", value)
}

func (e *GstFileSink) GetBufferSize() (uint, error) {
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

// Specify file mode used to open file
// Default: truncate (1)
func (e *GstFileSink) SetFileMode(value GstFileSinkFileMode) error {
	e.Element.SetArg("file-mode", string(value))
	return nil
}

func (e *GstFileSink) GetFileMode() (GstFileSinkFileMode, error) {
	var value GstFileSinkFileMode
	var ok bool
	v, err := e.Element.GetProperty("file-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFileSinkFileMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFileSinkFileMode")
	}
	return value, nil
}

// Location of the file to write
// Default: NULL
func (e *GstFileSink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstFileSink) GetLocation() (string, error) {
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

// N-to-1 input stream selector
type GstInputSelector struct {
	*gst.Element
}

func NewInputSelector() (*GstInputSelector, error) {
	e, err := gst.NewElement("input-selector")
	if err != nil {
		return nil, err
	}
	return &GstInputSelector{Element: e}, nil
}

func NewInputSelectorWithName(name string) (*GstInputSelector, error) {
	e, err := gst.NewElementWithName("input-selector", name)
	if err != nil {
		return nil, err
	}
	return &GstInputSelector{Element: e}, nil
}

// Behavior in sync-streams mode
// Default: active-segment (0)
func (e *GstInputSelector) SetSyncMode(value GstInputSelectorSyncMode) error {
	e.Element.SetArg("sync-mode", string(value))
	return nil
}

func (e *GstInputSelector) GetSyncMode() (GstInputSelectorSyncMode, error) {
	var value GstInputSelectorSyncMode
	var ok bool
	v, err := e.Element.GetProperty("sync-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstInputSelectorSyncMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstInputSelectorSyncMode")
	}
	return value, nil
}

// Synchronize inactive streams to the running time of the active stream or to the current clock
// Default: true
func (e *GstInputSelector) SetSyncStreams(value bool) error {
	return e.Element.SetProperty("sync-streams", value)
}

func (e *GstInputSelector) GetSyncStreams() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The currently active sink pad

func (e *GstInputSelector) SetActivePad(value *gst.Pad) error {
	return e.Element.SetProperty("active-pad", value)
}

func (e *GstInputSelector) GetActivePad() (*gst.Pad, error) {
	var value *gst.Pad
	var ok bool
	v, err := e.Element.GetProperty("active-pad")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Pad)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Pad")
	}
	return value, nil
}

// Cache buffers for active-pad
// Default: false
func (e *GstInputSelector) SetCacheBuffers(value bool) error {
	return e.Element.SetProperty("cache-buffers", value)
}

func (e *GstInputSelector) GetCacheBuffers() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("cache-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Drop backwards buffers on pad switch
// Default: false
func (e *GstInputSelector) SetDropBackwards(value bool) error {
	return e.Element.SetProperty("drop-backwards", value)
}

func (e *GstInputSelector) GetDropBackwards() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-backwards")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The number of sink pads
// Default: 0, Min: 0, Max: -1
func (e *GstInputSelector) GetNPads() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("n-pads")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Simple data queue
type GstQueue2 struct {
	*gst.Element
}

func NewQueue2() (*GstQueue2, error) {
	e, err := gst.NewElement("queue2")
	if err != nil {
		return nil, err
	}
	return &GstQueue2{Element: e}, nil
}

func NewQueue2WithName(name string) (*GstQueue2, error) {
	e, err := gst.NewElementWithName("queue2", name)
	if err != nil {
		return nil, err
	}
	return &GstQueue2{Element: e}, nil
}

// Average input data rate (bytes/s)
// Default: 0, Min: 0, Max: 9223372036854775807
func (e *GstQueue2) GetAvgInRate() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("avg-in-rate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Conversion value between data size and time
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstQueue2) GetBitrate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Current number of buffers in the queue
// Default: 0, Min: 0, Max: -1
func (e *GstQueue2) GetCurrentLevelBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-level-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max. number of buffers in the queue (0=disable)
// Default: 100, Min: 0, Max: -1
func (e *GstQueue2) SetMaxSizeBuffers(value uint) error {
	return e.Element.SetProperty("max-size-buffers", value)
}

func (e *GstQueue2) GetMaxSizeBuffers() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-buffers")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Use a bitrate from a downstream query to estimate buffer duration if not provided
// Default: true
func (e *GstQueue2) SetUseBitrateQuery(value bool) error {
	return e.Element.SetProperty("use-bitrate-query", value)
}

func (e *GstQueue2) GetUseBitrateQuery() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-bitrate-query")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Emit GST_MESSAGE_BUFFERING based on low-/high-percent thresholds (0%% = low-watermark, 100%% = high-watermark)
// Default: false
func (e *GstQueue2) SetUseBuffering(value bool) error {
	return e.Element.SetProperty("use-buffering", value)
}

func (e *GstQueue2) GetUseBuffering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-buffering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Use a bitrate from upstream tags to estimate buffer duration if not provided
// Default: false
func (e *GstQueue2) SetUseTagsBitrate(value bool) error {
	return e.Element.SetProperty("use-tags-bitrate", value)
}

func (e *GstQueue2) GetUseTagsBitrate() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-tags-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// High threshold for buffering to finish. Only used if use-buffering is True (Deprecated: use high-watermark instead)
// Default: 99, Min: 0, Max: 100
func (e *GstQueue2) SetHighPercent(value int) error {
	return e.Element.SetProperty("high-percent", value)
}

func (e *GstQueue2) GetHighPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("high-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Low threshold for buffering to start. Only used if use-buffering is True
// Default: 0.01, Min: 0, Max: 1
func (e *GstQueue2) SetLowWatermark(value float64) error {
	return e.Element.SetProperty("low-watermark", value)
}

func (e *GstQueue2) GetLowWatermark() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("low-watermark")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Max. amount of data in the queue (bytes, 0=disable)
// Default: 2097152, Min: 0, Max: -1
func (e *GstQueue2) SetMaxSizeBytes(value uint) error {
	return e.Element.SetProperty("max-size-bytes", value)
}

func (e *GstQueue2) GetMaxSizeBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Max. amount of data in the queue (in ns, 0=disable)
// Default: 2000000000, Min: 0, Max: 18446744073709551615
func (e *GstQueue2) SetMaxSizeTime(value uint64) error {
	return e.Element.SetProperty("max-size-time", value)
}

func (e *GstQueue2) GetMaxSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Max. amount of data in the ring buffer (bytes, 0 = disabled)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstQueue2) SetRingBufferMaxSize(value uint64) error {
	return e.Element.SetProperty("ring-buffer-max-size", value)
}

func (e *GstQueue2) GetRingBufferMaxSize() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ring-buffer-max-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// File template to store temporary files in, should contain directory and XXXXXX. (NULL == disabled)
// Default: NULL
func (e *GstQueue2) SetTempTemplate(value string) error {
	return e.Element.SetProperty("temp-template", value)
}

func (e *GstQueue2) GetTempTemplate() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("temp-template")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Current amount of data in the queue (bytes)
// Default: 0, Min: 0, Max: -1
func (e *GstQueue2) GetCurrentLevelBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("current-level-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Location to store temporary files in (Only read this property, use temp-template to configure the name template)
// Default: NULL
func (e *GstQueue2) GetTempLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("temp-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Current amount of data in the queue (in ns)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstQueue2) GetCurrentLevelTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("current-level-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// High threshold for buffering to finish. Only used if use-buffering is True
// Default: 0.99, Min: 0, Max: 1
func (e *GstQueue2) SetHighWatermark(value float64) error {
	return e.Element.SetProperty("high-watermark", value)
}

func (e *GstQueue2) GetHighWatermark() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("high-watermark")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Low threshold for buffering to start. Only used if use-buffering is True (Deprecated: use low-watermark instead)
// Default: 1, Min: 0, Max: 100
func (e *GstQueue2) SetLowPercent(value int) error {
	return e.Element.SetProperty("low-percent", value)
}

func (e *GstQueue2) GetLowPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("low-percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Remove the temp-location after use
// Default: true
func (e *GstQueue2) SetTempRemove(value bool) error {
	return e.Element.SetProperty("temp-remove", value)
}

func (e *GstQueue2) GetTempRemove() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("temp-remove")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Estimate the bitrate of the stream to calculate time level
// Default: true
func (e *GstQueue2) SetUseRateEstimate(value bool) error {
	return e.Element.SetProperty("use-rate-estimate", value)
}

func (e *GstQueue2) GetUseRateEstimate() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-rate-estimate")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Pass data without modification, limiting formats
type GstCapsFilter struct {
	*base.GstBaseTransform
}

func NewCapsFilter() (*GstCapsFilter, error) {
	e, err := gst.NewElement("capsfilter")
	if err != nil {
		return nil, err
	}
	return &GstCapsFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewCapsFilterWithName(name string) (*GstCapsFilter, error) {
	e, err := gst.NewElementWithName("capsfilter", name)
	if err != nil {
		return nil, err
	}
	return &GstCapsFilter{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// Restrict the possible allowed capabilities (NULL means ANY). Setting this property takes a reference to the supplied GstCaps object.
// Default: ANY
func (e *GstCapsFilter) SetCaps(value *gst.Caps) error {
	return e.Element.SetProperty("caps", value)
}

func (e *GstCapsFilter) GetCaps() (*gst.Caps, error) {
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

// Filter caps change behaviour
// Default: immediate (0)
func (e *GstCapsFilter) SetCapsChangeMode(value GstCapsFilterCapsChangeMode) error {
	e.Element.SetArg("caps-change-mode", string(value))
	return nil
}

func (e *GstCapsFilter) GetCapsChangeMode() (GstCapsFilterCapsChangeMode, error) {
	var value GstCapsFilterCapsChangeMode
	var ok bool
	v, err := e.Element.GetProperty("caps-change-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCapsFilterCapsChangeMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCapsFilterCapsChangeMode")
	}
	return value, nil
}

// Drops buffers and events or lets them through
type GstValve struct {
	*gst.Element
}

func NewValve() (*GstValve, error) {
	e, err := gst.NewElement("valve")
	if err != nil {
		return nil, err
	}
	return &GstValve{Element: e}, nil
}

func NewValveWithName(name string) (*GstValve, error) {
	e, err := gst.NewElementWithName("valve", name)
	if err != nil {
		return nil, err
	}
	return &GstValve{Element: e}, nil
}

// Whether to drop buffers and events or let them through
// Default: false
func (e *GstValve) SetDrop(value bool) error {
	return e.Element.SetProperty("drop", value)
}

func (e *GstValve) GetDrop() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The drop mode to use
// Default: drop-all (0)
func (e *GstValve) SetDropMode(value GstValveDropMode) error {
	e.Element.SetArg("drop-mode", string(value))
	return nil
}

func (e *GstValve) GetDropMode() (GstValveDropMode, error) {
	var value GstValveDropMode
	var ok bool
	v, err := e.Element.GetProperty("drop-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstValveDropMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstValveDropMode")
	}
	return value, nil
}

// Finds the media type of a stream
type GstTypeFindElement struct {
	*gst.Element
}

func NewTypeFindElement() (*GstTypeFindElement, error) {
	e, err := gst.NewElement("typefind")
	if err != nil {
		return nil, err
	}
	return &GstTypeFindElement{Element: e}, nil
}

func NewTypeFindElementWithName(name string) (*GstTypeFindElement, error) {
	e, err := gst.NewElementWithName("typefind", name)
	if err != nil {
		return nil, err
	}
	return &GstTypeFindElement{Element: e}, nil
}

// detected capabilities in stream

func (e *GstTypeFindElement) GetCaps() (*gst.Caps, error) {
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

// force caps without doing a typefind

func (e *GstTypeFindElement) SetForceCaps(value *gst.Caps) error {
	return e.Element.SetProperty("force-caps", value)
}

func (e *GstTypeFindElement) GetForceCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("force-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// minimum probability required to accept caps
// Default: 1, Min: 1, Max: 100
func (e *GstTypeFindElement) SetMinimum(value uint) error {
	return e.Element.SetProperty("minimum", value)
}

func (e *GstTypeFindElement) GetMinimum() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("minimum")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstCapsFilterCapsChangeMode string

const (
	GstCapsFilterCapsChangeModeImmediate GstCapsFilterCapsChangeMode = "immediate" // Only accept the current filter caps
	GstCapsFilterCapsChangeModeDelayed   GstCapsFilterCapsChangeMode = "delayed"   // Temporarily accept previous filter caps
)

type GstFileSinkBufferMode string

const (
	GstFileSinkBufferModeDefault    GstFileSinkBufferMode = "default"    // Default buffering
	GstFileSinkBufferModeFull       GstFileSinkBufferMode = "full"       // Fully buffered
	GstFileSinkBufferModeLine       GstFileSinkBufferMode = "line"       // Line buffered (deprecated, like full)
	GstFileSinkBufferModeUnbuffered GstFileSinkBufferMode = "unbuffered" // Unbuffered
)

type GstFileSinkFileMode string

const (
	GstFileSinkFileModeTruncate  GstFileSinkFileMode = "truncate"  // Truncate file (mode wb)
	GstFileSinkFileModeOutput    GstFileSinkFileMode = "output"    // Append file (mode ab)
	GstFileSinkFileModeOverwrite GstFileSinkFileMode = "overwrite" // Overwrite file without truncating (mode rb+)
)

type GstFakeSrcFillType string

const (
	GstFakeSrcFillTypeNothing     GstFakeSrcFillType = "nothing"      // Leave data as malloced
	GstFakeSrcFillTypeZero        GstFakeSrcFillType = "zero"         // Fill buffers with zeros
	GstFakeSrcFillTypeRandom      GstFakeSrcFillType = "random"       // Fill buffers with random data
	GstFakeSrcFillTypePattern     GstFakeSrcFillType = "pattern"      // Fill buffers with pattern 0x00 -> 0xff
	GstFakeSrcFillTypePatternSpan GstFakeSrcFillType = "pattern-span" // Fill buffers with pattern 0x00 -> 0xff that spans buffers
)

type GstOutputSelectorPadNegotiationMode string

const (
	GstOutputSelectorPadNegotiationModeNone   GstOutputSelectorPadNegotiationMode = "none"   // None
	GstOutputSelectorPadNegotiationModeAll    GstOutputSelectorPadNegotiationMode = "all"    // All
	GstOutputSelectorPadNegotiationModeActive GstOutputSelectorPadNegotiationMode = "active" // Active
)

type GstQueueLeaky string

const (
	GstQueueLeakyNo         GstQueueLeaky = "no"         // Not Leaky
	GstQueueLeakyUpstream   GstQueueLeaky = "upstream"   // Leaky on upstream (new buffers)
	GstQueueLeakyDownstream GstQueueLeaky = "downstream" // Leaky on downstream (old buffers)
)

type GstValveDropMode string

const (
	GstValveDropModeDropAll             GstValveDropMode = "drop-all"              // Drop all buffers and events
	GstValveDropModeForwardStickyEvents GstValveDropMode = "forward-sticky-events" // Drop all buffers but forward sticky events
	GstValveDropModeTransformToGap      GstValveDropMode = "transform-to-gap"      // Convert all dropped buffers into gap events and forward sticky events
)

type GstFakeSinkStateError string

const (
	GstFakeSinkStateErrorNone            GstFakeSinkStateError = "none"              // No state change errors
	GstFakeSinkStateErrorNullToReady     GstFakeSinkStateError = "null-to-ready"     // Fail state change from NULL to READY
	GstFakeSinkStateErrorReadyToPaused   GstFakeSinkStateError = "ready-to-paused"   // Fail state change from READY to PAUSED
	GstFakeSinkStateErrorPausedToPlaying GstFakeSinkStateError = "paused-to-playing" // Fail state change from PAUSED to PLAYING
	GstFakeSinkStateErrorPlayingToPaused GstFakeSinkStateError = "playing-to-paused" // Fail state change from PLAYING to PAUSED
	GstFakeSinkStateErrorPausedToReady   GstFakeSinkStateError = "paused-to-ready"   // Fail state change from PAUSED to READY
	GstFakeSinkStateErrorReadyToNull     GstFakeSinkStateError = "ready-to-null"     // Fail state change from READY to NULL
)

type GstFakeSrcDataType string

const (
	GstFakeSrcDataTypeAllocate  GstFakeSrcDataType = "allocate"  // Allocate data
	GstFakeSrcDataTypeSubbuffer GstFakeSrcDataType = "subbuffer" // Subbuffer data
)

type GstFakeSrcSizeType string

const (
	GstFakeSrcSizeTypeEmpty  GstFakeSrcSizeType = "empty"  // Send empty buffers
	GstFakeSrcSizeTypeFixed  GstFakeSrcSizeType = "fixed"  // Fixed size buffers (sizemax sized)
	GstFakeSrcSizeTypeRandom GstFakeSrcSizeType = "random" // Random sized buffers (sizemin <= size <= sizemax)
)

type GstInputSelectorSyncMode string

const (
	GstInputSelectorSyncModeActiveSegment GstInputSelectorSyncMode = "active-segment" // Sync using the current active segment
	GstInputSelectorSyncModeClock         GstInputSelectorSyncMode = "clock"          // Sync using the clock
)

type GstTeePullMode string

const (
	GstTeePullModeNever  GstTeePullMode = "never"  // Never activate in pull mode
	GstTeePullModeSingle GstTeePullMode = "single" // Only one src pad can be active in pull mode
)
