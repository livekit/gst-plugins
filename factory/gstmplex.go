// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// High-quality MPEG/DVD/SVCD/VCD video/audio multiplexer
type GstMplex struct {
	*gst.Element
}

func NewMplex() (*GstMplex, error) {
	e, err := gst.NewElement("mplex")
	if err != nil {
		return nil, err
	}
	return &GstMplex{Element: e}, nil
}

func NewMplexWithName(name string) (*GstMplex, error) {
	e, err := gst.NewElementWithName("mplex", name)
	if err != nil {
		return nil, err
	}
	return &GstMplex{Element: e}, nil
}

// Target decoders video buffer size (kB) [default determined by format if not explicitly set]
// Default: 0, Min: 20, Max: 4000
func (e *GstMplex) SetBufsize(value int) error {
	return e.Element.SetProperty("bufsize", value)
}

func (e *GstMplex) GetBufsize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("bufsize")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Encoding profile format
// Default: mpeg-1 (0)
func (e *GstMplex) SetFormat(value GstMplexFormat) error {
	e.Element.SetArg("format", string(value))
	return nil
}

func (e *GstMplex) GetFormat() (GstMplexFormat, error) {
	var value GstMplexFormat
	var ok bool
	v, err := e.Element.GetProperty("format")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstMplexFormat)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstMplexFormat")
	}
	return value, nil
}

// Bitrate of output stream in kbps (0 = autodetect)
// Default: 0, Min: 0, Max: 15360
func (e *GstMplex) SetMuxBitrate(value int) error {
	return e.Element.SetProperty("mux-bitrate", value)
}

func (e *GstMplex) GetMuxBitrate() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("mux-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of packets per pack for generic formats
// Default: 1, Min: 1, Max: 100
func (e *GstMplex) SetPacketsPerPack(value int) error {
	return e.Element.SetProperty("packets-per-pack", value)
}

func (e *GstMplex) GetPacketsPerPack() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("packets-per-pack")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Specify sector size in bytes for generic formats
// Default: 2048, Min: 256, Max: 16384
func (e *GstMplex) SetSectorSize(value int) error {
	return e.Element.SetProperty("sector-size", value)
}

func (e *GstMplex) GetSectorSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("sector-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Create system header in every pack for generic formats
// Default: false
func (e *GstMplex) SetSystemHeaders(value bool) error {
	return e.Element.SetProperty("system-headers", value)
}

func (e *GstMplex) GetSystemHeaders() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("system-headers")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether the input video stream is variable bitrate
// Default: false
func (e *GstMplex) SetVbr(value bool) error {
	return e.Element.SetProperty("vbr", value)
}

func (e *GstMplex) GetVbr() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vbr")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstMplexFormat string

const (
	GstMplexFormatMpeg1     GstMplexFormat = "mpeg-1"     // Generic MPEG-1
	GstMplexFormatVcd       GstMplexFormat = "vcd"        // Standard VCD
	GstMplexFormatVcdNsr    GstMplexFormat = "vcd-nsr"    // User VCD
	GstMplexFormatMpeg2     GstMplexFormat = "mpeg-2"     // Generic MPEG-2
	GstMplexFormatSvcd      GstMplexFormat = "svcd"       // Standard SVCD
	GstMplexFormatSvcdNsr   GstMplexFormat = "svcd-nsr"   // User SVCD
	GstMplexFormatVcdStill  GstMplexFormat = "vcd-still"  // VCD Stills sequences
	GstMplexFormatSvcdStill GstMplexFormat = "svcd-still" // SVCD Stills sequences
	GstMplexFormatDvdNav    GstMplexFormat = "dvd-nav"    // DVD MPEG-2 for dvdauthor
	GstMplexFormatDvd       GstMplexFormat = "dvd"        // DVD MPEG-2
	GstMplexFormatAtsc480i  GstMplexFormat = "atsc-480i"  // ATSC 480i
	GstMplexFormatAtsc480p  GstMplexFormat = "atsc-480p"  // ATSC 480p
	GstMplexFormatAtsc720p  GstMplexFormat = "atsc-720p"  // ATSC 720p
	GstMplexFormatAtsc1080i GstMplexFormat = "atsc-1080i" // ATSC 1080i
)
