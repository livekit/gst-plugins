// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Payload-encode Raw audio into AAF AVTPDU (IEEE 1722)
type GstAvtpAafPay struct {
	*GstAvtpBasePayload
}

func NewAvtpAafPay() (*GstAvtpAafPay, error) {
	e, err := gst.NewElement("avtpaafpay")
	if err != nil {
		return nil, err
	}
	return &GstAvtpAafPay{GstAvtpBasePayload: &GstAvtpBasePayload{Element: e}}, nil
}

func NewAvtpAafPayWithName(name string) (*GstAvtpAafPay, error) {
	e, err := gst.NewElementWithName("avtpaafpay", name)
	if err != nil {
		return nil, err
	}
	return &GstAvtpAafPay{GstAvtpBasePayload: &GstAvtpBasePayload{Element: e}}, nil
}

// AAF timestamping mode
// Default: normal (0)
func (e *GstAvtpAafPay) SetTimestampMode(value GstAvtpAafTimestampMode) error {
	e.Element.SetArg("timestamp-mode", string(value))
	return nil
}

func (e *GstAvtpAafPay) GetTimestampMode() (GstAvtpAafTimestampMode, error) {
	var value GstAvtpAafTimestampMode
	var ok bool
	v, err := e.Element.GetProperty("timestamp-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAvtpAafTimestampMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAvtpAafTimestampMode")
	}
	return value, nil
}

// Payload-encode compressed video into CVF AVTPDU (IEEE 1722)
type GstAvtpCvfPay struct {
	*GstAvtpVfPayBase
}

func NewAvtpCvfPay() (*GstAvtpCvfPay, error) {
	e, err := gst.NewElement("avtpcvfpay")
	if err != nil {
		return nil, err
	}
	return &GstAvtpCvfPay{GstAvtpVfPayBase: &GstAvtpVfPayBase{GstAvtpBasePayload: &GstAvtpBasePayload{Element: e}}}, nil
}

func NewAvtpCvfPayWithName(name string) (*GstAvtpCvfPay, error) {
	e, err := gst.NewElementWithName("avtpcvfpay", name)
	if err != nil {
		return nil, err
	}
	return &GstAvtpCvfPay{GstAvtpVfPayBase: &GstAvtpVfPayBase{GstAvtpBasePayload: &GstAvtpBasePayload{Element: e}}}, nil
}

// Extracts raw video from RVF AVTPDUs
type GstAvtpRvfDepay struct {
	*GstAvtpVfDepayBase
}

func NewAvtpRvfDepay() (*GstAvtpRvfDepay, error) {
	e, err := gst.NewElement("avtprvfdepay")
	if err != nil {
		return nil, err
	}
	return &GstAvtpRvfDepay{GstAvtpVfDepayBase: &GstAvtpVfDepayBase{GstAvtpBaseDepayload: &GstAvtpBaseDepayload{Element: e}}}, nil
}

func NewAvtpRvfDepayWithName(name string) (*GstAvtpRvfDepay, error) {
	e, err := gst.NewElementWithName("avtprvfdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstAvtpRvfDepay{GstAvtpVfDepayBase: &GstAvtpVfDepayBase{GstAvtpBaseDepayload: &GstAvtpBaseDepayload{Element: e}}}, nil
}

// Payload-encode raw video into RVF AVTPDU (IEEE 1722)
type GstAvtpRvfPay struct {
	*GstAvtpVfPayBase
}

func NewAvtpRvfPay() (*GstAvtpRvfPay, error) {
	e, err := gst.NewElement("avtprvfpay")
	if err != nil {
		return nil, err
	}
	return &GstAvtpRvfPay{GstAvtpVfPayBase: &GstAvtpVfPayBase{GstAvtpBasePayload: &GstAvtpBasePayload{Element: e}}}, nil
}

func NewAvtpRvfPayWithName(name string) (*GstAvtpRvfPay, error) {
	e, err := gst.NewElementWithName("avtprvfpay", name)
	if err != nil {
		return nil, err
	}
	return &GstAvtpRvfPay{GstAvtpVfPayBase: &GstAvtpVfPayBase{GstAvtpBasePayload: &GstAvtpBasePayload{Element: e}}}, nil
}

// Send AVTPDUs over the network
type GstAvtpSink struct {
	*base.GstBaseSink
}

func NewAvtpSink() (*GstAvtpSink, error) {
	e, err := gst.NewElement("avtpsink")
	if err != nil {
		return nil, err
	}
	return &GstAvtpSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewAvtpSinkWithName(name string) (*GstAvtpSink, error) {
	e, err := gst.NewElementWithName("avtpsink", name)
	if err != nil {
		return nil, err
	}
	return &GstAvtpSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Destination MAC address from Ethernet frames
// Default: 01:AA:AA:AA:AA:AA
func (e *GstAvtpSink) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

func (e *GstAvtpSink) GetAddress() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("address")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Network interface utilized to transmit AVTPDUs
// Default: eth0
func (e *GstAvtpSink) SetIfname(value string) error {
	return e.Element.SetProperty("ifname", value)
}

func (e *GstAvtpSink) GetIfname() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ifname")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Priority configured into socket (SO_PRIORITY)
// Default: 0, Min: 0, Max: 2147483647
func (e *GstAvtpSink) SetPriority(value int) error {
	return e.Element.SetProperty("priority", value)
}

func (e *GstAvtpSink) GetPriority() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("priority")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Receive AVTPDUs from the network
type GstAvtpSrc struct {
	*base.GstPushSrc
}

func NewAvtpSrc() (*GstAvtpSrc, error) {
	e, err := gst.NewElement("avtpsrc")
	if err != nil {
		return nil, err
	}
	return &GstAvtpSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewAvtpSrcWithName(name string) (*GstAvtpSrc, error) {
	e, err := gst.NewElementWithName("avtpsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstAvtpSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Destination MAC address to listen to
// Default: 01:AA:AA:AA:AA:AA
func (e *GstAvtpSrc) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

func (e *GstAvtpSrc) GetAddress() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("address")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Network interface utilized to receive AVTPDUs
// Default: eth0
func (e *GstAvtpSrc) SetIfname(value string) error {
	return e.Element.SetProperty("ifname", value)
}

func (e *GstAvtpSrc) GetIfname() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ifname")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Extracts raw audio from AAF AVTPDUs
type GstAvtpAafDepay struct {
	*GstAvtpBaseDepayload
}

func NewAvtpAafDepay() (*GstAvtpAafDepay, error) {
	e, err := gst.NewElement("avtpaafdepay")
	if err != nil {
		return nil, err
	}
	return &GstAvtpAafDepay{GstAvtpBaseDepayload: &GstAvtpBaseDepayload{Element: e}}, nil
}

func NewAvtpAafDepayWithName(name string) (*GstAvtpAafDepay, error) {
	e, err := gst.NewElementWithName("avtpaafdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstAvtpAafDepay{GstAvtpBaseDepayload: &GstAvtpBaseDepayload{Element: e}}, nil
}

// Check if the AVTP presentation time is synchronized with clock provided by a CRF stream
type GstAvtpCrfCheck struct {
	*GstAvtpCrfBase
}

func NewAvtpCrfCheck() (*GstAvtpCrfCheck, error) {
	e, err := gst.NewElement("avtpcrfcheck")
	if err != nil {
		return nil, err
	}
	return &GstAvtpCrfCheck{GstAvtpCrfBase: &GstAvtpCrfBase{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAvtpCrfCheckWithName(name string) (*GstAvtpCrfCheck, error) {
	e, err := gst.NewElementWithName("avtpcrfcheck", name)
	if err != nil {
		return nil, err
	}
	return &GstAvtpCrfCheck{GstAvtpCrfBase: &GstAvtpCrfBase{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Drop the packets which are not within 25%%%% of the sample period of the CRF timestamps
// Default: false
func (e *GstAvtpCrfCheck) SetDropInvalid(value bool) error {
	return e.Element.SetProperty("drop-invalid", value)
}

func (e *GstAvtpCrfCheck) GetDropInvalid() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-invalid")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Synchronize Presentation Time from AVTPDUs so they are phase-locked with clock provided by CRF stream
type GstAvtpCrfSync struct {
	*GstAvtpCrfBase
}

func NewAvtpCrfSync() (*GstAvtpCrfSync, error) {
	e, err := gst.NewElement("avtpcrfsync")
	if err != nil {
		return nil, err
	}
	return &GstAvtpCrfSync{GstAvtpCrfBase: &GstAvtpCrfBase{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

func NewAvtpCrfSyncWithName(name string) (*GstAvtpCrfSync, error) {
	e, err := gst.NewElementWithName("avtpcrfsync", name)
	if err != nil {
		return nil, err
	}
	return &GstAvtpCrfSync{GstAvtpCrfBase: &GstAvtpCrfBase{GstBaseTransform: &base.GstBaseTransform{Element: e}}}, nil
}

// Extracts compressed video from CVF AVTPDUs
type GstAvtpCvfDepay struct {
	*GstAvtpVfDepayBase
}

func NewAvtpCvfDepay() (*GstAvtpCvfDepay, error) {
	e, err := gst.NewElement("avtpcvfdepay")
	if err != nil {
		return nil, err
	}
	return &GstAvtpCvfDepay{GstAvtpVfDepayBase: &GstAvtpVfDepayBase{GstAvtpBaseDepayload: &GstAvtpBaseDepayload{Element: e}}}, nil
}

func NewAvtpCvfDepayWithName(name string) (*GstAvtpCvfDepay, error) {
	e, err := gst.NewElementWithName("avtpcvfdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstAvtpCvfDepay{GstAvtpVfDepayBase: &GstAvtpVfDepayBase{GstAvtpBaseDepayload: &GstAvtpBaseDepayload{Element: e}}}, nil
}

type GstAvtpVfDepayBase struct {
	*GstAvtpBaseDepayload
}

type GstAvtpVfPayBase struct {
	*GstAvtpBasePayload
}

// Maximum number of network frames to be sent on each Measurement Interval
// Default: 1, Min: 1, Max: -1
func (e *GstAvtpVfPayBase) SetMaxIntervalFrames(value uint) error {
	return e.Element.SetProperty("max-interval-frames", value)
}

func (e *GstAvtpVfPayBase) GetMaxIntervalFrames() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-interval-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Measurement interval of stream in nanoseconds
// Default: 250000, Min: 0, Max: 18446744073709551615
func (e *GstAvtpVfPayBase) SetMeasurementInterval(value uint64) error {
	return e.Element.SetProperty("measurement-interval", value)
}

func (e *GstAvtpVfPayBase) GetMeasurementInterval() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("measurement-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Maximum Transit Unit (MTU) of underlying network in bytes
// Default: 1500, Min: 0, Max: -1
func (e *GstAvtpVfPayBase) SetMtu(value uint) error {
	return e.Element.SetProperty("mtu", value)
}

func (e *GstAvtpVfPayBase) GetMtu() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("mtu")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstAvtpAafTimestampMode string

const (
	GstAvtpAafTimestampModeNormal GstAvtpAafTimestampMode = "normal" // Normal timestamping mode
	GstAvtpAafTimestampModeSparse GstAvtpAafTimestampMode = "sparse" // Sparse timestamping mode
)

type GstAvtpBaseDepayload struct {
	*gst.Element
}

// Stream ID associated with the AVTPDU
// Default: 12302652060662169600, Min: 0, Max: 18446744073709551615
func (e *GstAvtpBaseDepayload) SetStreamid(value uint64) error {
	return e.Element.SetProperty("streamid", value)
}

func (e *GstAvtpBaseDepayload) GetStreamid() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("streamid")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

type GstAvtpBasePayload struct {
	*gst.Element
}

// Stream ID associated with the AVTPDU
// Default: 12302652060662169600, Min: 0, Max: 18446744073709551615
func (e *GstAvtpBasePayload) SetStreamid(value uint64) error {
	return e.Element.SetProperty("streamid", value)
}

func (e *GstAvtpBasePayload) GetStreamid() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("streamid")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Timing Uncertainty (TU) in nanoseconds
// Default: 1000000, Min: 0, Max: -1
func (e *GstAvtpBasePayload) SetTu(value uint) error {
	return e.Element.SetProperty("tu", value)
}

func (e *GstAvtpBasePayload) GetTu() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("tu")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum Transit Time (MTT) in nanoseconds
// Default: 50000000, Min: 0, Max: -1
func (e *GstAvtpBasePayload) SetMtt(value uint) error {
	return e.Element.SetProperty("mtt", value)
}

func (e *GstAvtpBasePayload) GetMtt() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("mtt")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum amount of time (in ns) the pipeline can take for processing the buffer
// Default: 20000000, Min: 0, Max: 18446744073709551615
func (e *GstAvtpBasePayload) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

func (e *GstAvtpBasePayload) GetProcessingDeadline() (uint64, error) {
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

type GstAvtpCrfBase struct {
	*base.GstBaseTransform
}

// Destination MAC address expected on the Ethernet frames
// Default: 01:AA:AA:AA:AA:AA
func (e *GstAvtpCrfBase) SetAddress(value string) error {
	return e.Element.SetProperty("address", value)
}

func (e *GstAvtpCrfBase) GetAddress() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("address")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Network interface utilized to receive CRF AVTPDUs
// Default: eth0
func (e *GstAvtpCrfBase) SetIfname(value string) error {
	return e.Element.SetProperty("ifname", value)
}

func (e *GstAvtpCrfBase) GetIfname() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ifname")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Stream ID associated with the CRF AVTPDU
// Default: 12302652060662173696, Min: 0, Max: 18446744073709551615
func (e *GstAvtpCrfBase) SetStreamid(value uint64) error {
	return e.Element.SetProperty("streamid", value)
}

func (e *GstAvtpCrfBase) GetStreamid() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("streamid")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}
