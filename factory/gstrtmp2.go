// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Sink element for RTMP streams
type GstRtmp2Sink struct {
	*base.GstBaseSink
}

func NewRtmp2Sink() (*GstRtmp2Sink, error) {
	e, err := gst.NewElement("rtmp2sink")
	if err != nil {
		return nil, err
	}
	return &GstRtmp2Sink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewRtmp2SinkWithName(name string) (*GstRtmp2Sink, error) {
	e, err := gst.NewElementWithName("rtmp2sink", name)
	if err != nil {
		return nil, err
	}
	return &GstRtmp2Sink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// Connect on READY, otherwise on first push
// Default: true
func (e *GstRtmp2Sink) SetAsyncConnect(value bool) error {
	return e.Element.SetProperty("async-connect", value)
}

func (e *GstRtmp2Sink) GetAsyncConnect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async-connect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// RTMP chunk size
// Default: 128, Min: 1, Max: 2147483647
func (e *GstRtmp2Sink) SetChunkSize(value uint) error {
	return e.Element.SetProperty("chunk-size", value)
}

func (e *GstRtmp2Sink) GetChunkSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("chunk-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Bitrate in kbit/sec to pace outgoing packets
// Default: 0, Min: 0, Max: 17179869
func (e *GstRtmp2Sink) SetPeakKbps(value uint) error {
	return e.Element.SetProperty("peak-kbps", value)
}

func (e *GstRtmp2Sink) GetPeakKbps() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("peak-kbps")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Retrieve a statistics structure
// Default: GstRtmpConnectionStats, in-chunk-size=(uint)0, out-chunk-size=(uint)0, in-window-ack-size=(uint)0, out-window-ack-size=(uint)0, in-bytes-total=(guint64)0, out-bytes-total=(guint64)0, in-bytes-acked=(guint64)0, out-bytes-acked=(guint64)0;
func (e *GstRtmp2Sink) GetStats() (*gst.Structure, error) {
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

// RTMP commands to send on EOS event before closing connection
// Default: deletestream+fcunpublish
func (e *GstRtmp2Sink) SetStopCommands(value GstRtmpStopCommands) error {
	e.Element.SetArg("stop-commands", fmt.Sprint(value))
	return nil
}

func (e *GstRtmp2Sink) GetStopCommands() (GstRtmpStopCommands, error) {
	var value GstRtmpStopCommands
	var ok bool
	v, err := e.Element.GetProperty("stop-commands")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRtmpStopCommands)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRtmpStopCommands")
	}
	return value, nil
}

// Source element for RTMP streams
type GstRtmp2Src struct {
	*base.GstPushSrc
}

func NewRtmp2Src() (*GstRtmp2Src, error) {
	e, err := gst.NewElement("rtmp2src")
	if err != nil {
		return nil, err
	}
	return &GstRtmp2Src{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewRtmp2SrcWithName(name string) (*GstRtmp2Src, error) {
	e, err := gst.NewElementWithName("rtmp2src", name)
	if err != nil {
		return nil, err
	}
	return &GstRtmp2Src{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// If set, an error is raised if the connection is closed without receiving an EOF RTMP message first. If not set, those are reported using EOS
// Default: false
func (e *GstRtmp2Src) SetNoEofIsError(value bool) error {
	return e.Element.SetProperty("no-eof-is-error", value)
}

func (e *GstRtmp2Src) GetNoEofIsError() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("no-eof-is-error")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Retrieve a statistics structure
// Default: GstRtmpConnectionStats, in-chunk-size=(uint)0, out-chunk-size=(uint)0, in-window-ack-size=(uint)0, out-window-ack-size=(uint)0, in-bytes-total=(guint64)0, out-bytes-total=(guint64)0, in-bytes-acked=(guint64)0, out-bytes-acked=(guint64)0;
func (e *GstRtmp2Src) GetStats() (*gst.Structure, error) {
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

// Connect on READY, otherwise on first push
// Default: true
func (e *GstRtmp2Src) SetAsyncConnect(value bool) error {
	return e.Element.SetProperty("async-connect", value)
}

func (e *GstRtmp2Src) GetAsyncConnect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async-connect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The maximum allowed time in seconds for valid packets not to arrive from the peer (0 = no timeout)
// Default: 0, Min: 0, Max: -1
func (e *GstRtmp2Src) SetIdleTimeout(value uint) error {
	return e.Element.SetProperty("idle-timeout", value)
}

func (e *GstRtmp2Src) GetIdleTimeout() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("idle-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstRtmpAuthmod string

const (
	GstRtmpAuthmodNone  GstRtmpAuthmod = "none"  // GST_RTMP_AUTHMOD_NONE
	GstRtmpAuthmodAuto  GstRtmpAuthmod = "auto"  // GST_RTMP_AUTHMOD_AUTO
	GstRtmpAuthmodAdobe GstRtmpAuthmod = "adobe" // GST_RTMP_AUTHMOD_ADOBE
)

type GstRtmpScheme string

const (
	GstRtmpSchemeRtmp  GstRtmpScheme = "rtmp"  // GST_RTMP_SCHEME_RTMP
	GstRtmpSchemeRtmps GstRtmpScheme = "rtmps" // GST_RTMP_SCHEME_RTMPS
)

type GstRtmpStopCommands int

const (
	GstRtmpStopCommandsNone         GstRtmpStopCommands = 0x00000000 // No command
	GstRtmpStopCommandsFcunpublish  GstRtmpStopCommands = 0x00000001 // FCUnpublish
	GstRtmpStopCommandsClosestream  GstRtmpStopCommands = 0x00000002 // closeStream
	GstRtmpStopCommandsDeletestream GstRtmpStopCommands = 0x00000004 // deleteStream
)
