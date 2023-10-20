// source: gst-plugins-ugly

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Accepts raw RTP and RTCP packets and sends them forward
type GstRDTManager struct {
	*gst.Element
}

func NewRDTManager() (*GstRDTManager, error) {
	e, err := gst.NewElement("rdtmanager")
	if err != nil {
		return nil, err
	}
	return &GstRDTManager{Element: e}, nil
}

func NewRDTManagerWithName(name string) (*GstRDTManager, error) {
	e, err := gst.NewElementWithName("rdtmanager", name)
	if err != nil {
		return nil, err
	}
	return &GstRDTManager{Element: e}, nil
}

// Amount of ms to buffer
// Default: 200, Min: 0, Max: -1
func (e *GstRDTManager) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstRDTManager) GetLatency() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Demultiplex a RealMedia file into audio and video streams
type GstRMDemux struct {
	*gst.Element
}

func NewRMDemux() (*GstRMDemux, error) {
	e, err := gst.NewElement("rmdemux")
	if err != nil {
		return nil, err
	}
	return &GstRMDemux{Element: e}, nil
}

func NewRMDemuxWithName(name string) (*GstRMDemux, error) {
	e, err := gst.NewElementWithName("rmdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstRMDemux{Element: e}, nil
}

// Extends RTSP so that it can handle RealMedia setup
type GstRTSPReal struct {
	*gst.Element
}

func NewRTSPReal() (*GstRTSPReal, error) {
	e, err := gst.NewElement("rtspreal")
	if err != nil {
		return nil, err
	}
	return &GstRTSPReal{Element: e}, nil
}

func NewRTSPRealWithName(name string) (*GstRTSPReal, error) {
	e, err := gst.NewElementWithName("rtspreal", name)
	if err != nil {
		return nil, err
	}
	return &GstRTSPReal{Element: e}, nil
}

// Receive data over the network via PNM
type GstPNMSrc struct {
	*base.GstPushSrc
}

func NewPNMSrc() (*GstPNMSrc, error) {
	e, err := gst.NewElement("pnmsrc")
	if err != nil {
		return nil, err
	}
	return &GstPNMSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewPNMSrcWithName(name string) (*GstPNMSrc, error) {
	e, err := gst.NewElementWithName("pnmsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstPNMSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Location of the PNM url to read
// Default: NULL
func (e *GstPNMSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstPNMSrc) GetLocation() (string, error) {
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

// Demultiplex a RealAudio file
type GstRealAudioDemux struct {
	*gst.Element
}

func NewRealAudioDemux() (*GstRealAudioDemux, error) {
	e, err := gst.NewElement("rademux")
	if err != nil {
		return nil, err
	}
	return &GstRealAudioDemux{Element: e}, nil
}

func NewRealAudioDemuxWithName(name string) (*GstRealAudioDemux, error) {
	e, err := gst.NewElementWithName("rademux", name)
	if err != nil {
		return nil, err
	}
	return &GstRealAudioDemux{Element: e}, nil
}

// Extracts RealMedia from RDT packets
type GstRDTDepay struct {
	*gst.Element
}

func NewRDTDepay() (*GstRDTDepay, error) {
	e, err := gst.NewElement("rdtdepay")
	if err != nil {
		return nil, err
	}
	return &GstRDTDepay{Element: e}, nil
}

func NewRDTDepayWithName(name string) (*GstRDTDepay, error) {
	e, err := gst.NewElementWithName("rdtdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRDTDepay{Element: e}, nil
}
