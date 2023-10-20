// source: gst-plugins-ugly

package factory

import (
	"github.com/go-gst/go-gst/gst"
)

// Extends RTSP so that it can handle WMS setup
type GstRTSPWMS struct {
	*gst.Element
}

func NewRTSPWMS() (*GstRTSPWMS, error) {
	e, err := gst.NewElement("rtspwms")
	if err != nil {
		return nil, err
	}
	return &GstRTSPWMS{Element: e}, nil
}

func NewRTSPWMSWithName(name string) (*GstRTSPWMS, error) {
	e, err := gst.NewElementWithName("rtspwms", name)
	if err != nil {
		return nil, err
	}
	return &GstRTSPWMS{Element: e}, nil
}

// Demultiplexes ASF Streams
type GstASFDemux struct {
	*gst.Element
}

func NewASFDemux() (*GstASFDemux, error) {
	e, err := gst.NewElement("asfdemux")
	if err != nil {
		return nil, err
	}
	return &GstASFDemux{Element: e}, nil
}

func NewASFDemuxWithName(name string) (*GstASFDemux, error) {
	e, err := gst.NewElementWithName("asfdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstASFDemux{Element: e}, nil
}

// Extracts ASF streams from RTP
type GstRtpAsfDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpAsfDepay() (*GstRtpAsfDepay, error) {
	e, err := gst.NewElement("rtpasfdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpAsfDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpAsfDepayWithName(name string) (*GstRtpAsfDepay, error) {
	e, err := gst.NewElementWithName("rtpasfdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpAsfDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}
