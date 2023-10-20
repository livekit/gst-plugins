// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Multiplex audio and video into a MJ2 file
type GstMJ2Mux struct {
	*GstBaseQTMux
}

func NewMJ2Mux() (*GstMJ2Mux, error) {
	e, err := gst.NewElement("mj2mux")
	if err != nil {
		return nil, err
	}
	return &GstMJ2Mux{GstBaseQTMux: &GstBaseQTMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func NewMJ2MuxWithName(name string) (*GstMJ2Mux, error) {
	e, err := gst.NewElementWithName("mj2mux", name)
	if err != nil {
		return nil, err
	}
	return &GstMJ2Mux{GstBaseQTMux: &GstBaseQTMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// If set to true, the output should be as if it is to be streamed and hence no indexes written or duration written. (DEPRECATED, only valid for fragmented MP4)
// Default: false
func (e *GstMJ2Mux) SetStreamable(value bool) error {
	return e.Element.SetProperty("streamable", value)
}

func (e *GstMJ2Mux) GetStreamable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Multiplex audio and video into a MP4 file
type GstMP4Mux struct {
	*GstBaseQTMux
}

func NewMP4Mux() (*GstMP4Mux, error) {
	e, err := gst.NewElement("mp4mux")
	if err != nil {
		return nil, err
	}
	return &GstMP4Mux{GstBaseQTMux: &GstBaseQTMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func NewMP4MuxWithName(name string) (*GstMP4Mux, error) {
	e, err := gst.NewElementWithName("mp4mux", name)
	if err != nil {
		return nil, err
	}
	return &GstMP4Mux{GstBaseQTMux: &GstBaseQTMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// If set to true, the output should be as if it is to be streamed and hence no indexes written or duration written. (DEPRECATED, only valid for fragmented MP4)
// Default: false
func (e *GstMP4Mux) SetStreamable(value bool) error {
	return e.Element.SetProperty("streamable", value)
}

func (e *GstMP4Mux) GetStreamable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Demultiplex a QuickTime file into audio and video streams
type GstQTDemux struct {
	*gst.Element
}

func NewQTDemux() (*GstQTDemux, error) {
	e, err := gst.NewElement("qtdemux")
	if err != nil {
		return nil, err
	}
	return &GstQTDemux{Element: e}, nil
}

func NewQTDemuxWithName(name string) (*GstQTDemux, error) {
	e, err := gst.NewElementWithName("qtdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstQTDemux{Element: e}, nil
}

// Recovers unfinished qtmux files
type GstQTMoovRecover struct {
	*gst.Pipeline
}

func NewQTMoovRecover() (*GstQTMoovRecover, error) {
	e, err := gst.NewElement("qtmoovrecover")
	if err != nil {
		return nil, err
	}
	return &GstQTMoovRecover{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

func NewQTMoovRecoverWithName(name string) (*GstQTMoovRecover, error) {
	e, err := gst.NewElementWithName("qtmoovrecover", name)
	if err != nil {
		return nil, err
	}
	return &GstQTMoovRecover{Pipeline: &gst.Pipeline{Bin: &gst.Bin{Element: e}}}, nil
}

// If the broken input is from faststart mode
// Default: false
func (e *GstQTMoovRecover) SetFaststartMode(value bool) error {
	return e.Element.SetProperty("faststart-mode", value)
}

func (e *GstQTMoovRecover) GetFaststartMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("faststart-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Path to write the fixed file to (used as output)
// Default: NULL
func (e *GstQTMoovRecover) SetFixedOutput(value string) error {
	return e.Element.SetProperty("fixed-output", value)
}

func (e *GstQTMoovRecover) GetFixedOutput() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("fixed-output")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Path to recovery file (used as input)
// Default: NULL
func (e *GstQTMoovRecover) SetRecoveryInput(value string) error {
	return e.Element.SetProperty("recovery-input", value)
}

func (e *GstQTMoovRecover) GetRecoveryInput() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("recovery-input")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Path to broken input file. (If qtmux was on faststart mode, this file is the faststart file)
// Default: NULL
func (e *GstQTMoovRecover) SetBrokenInput(value string) error {
	return e.Element.SetProperty("broken-input", value)
}

func (e *GstQTMoovRecover) GetBrokenInput() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("broken-input")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Multiplex audio and video into a QuickTime file
type GstQTMux struct {
	*GstBaseQTMux
}

func NewQTMux() (*GstQTMux, error) {
	e, err := gst.NewElement("qtmux")
	if err != nil {
		return nil, err
	}
	return &GstQTMux{GstBaseQTMux: &GstBaseQTMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func NewQTMuxWithName(name string) (*GstQTMux, error) {
	e, err := gst.NewElementWithName("qtmux", name)
	if err != nil {
		return nil, err
	}
	return &GstQTMux{GstBaseQTMux: &GstBaseQTMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// If set to true, the output should be as if it is to be streamed and hence no indexes written or duration written. (DEPRECATED, only valid for fragmented MP4)
// Default: false
func (e *GstQTMux) SetStreamable(value bool) error {
	return e.Element.SetProperty("streamable", value)
}

func (e *GstQTMux) GetStreamable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Extracts Quicktime audio/video from RTP packets
type GstRtpXQTDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpXQTDepay() (*GstRtpXQTDepay, error) {
	e, err := gst.NewElement("rtpxqtdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpXQTDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpXQTDepayWithName(name string) (*GstRtpXQTDepay, error) {
	e, err := gst.NewElementWithName("rtpxqtdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpXQTDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Multiplex audio and video into a 3GPP file
type Gst3GPPMux struct {
	*GstBaseQTMux
}

func New3GPPMux() (*Gst3GPPMux, error) {
	e, err := gst.NewElement("3gppmux")
	if err != nil {
		return nil, err
	}
	return &Gst3GPPMux{GstBaseQTMux: &GstBaseQTMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func New3GPPMuxWithName(name string) (*Gst3GPPMux, error) {
	e, err := gst.NewElementWithName("3gppmux", name)
	if err != nil {
		return nil, err
	}
	return &Gst3GPPMux{GstBaseQTMux: &GstBaseQTMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// If set to true, the output should be as if it is to be streamed and hence no indexes written or duration written. (DEPRECATED, only valid for fragmented MP4)
// Default: false
func (e *Gst3GPPMux) SetStreamable(value bool) error {
	return e.Element.SetProperty("streamable", value)
}

func (e *Gst3GPPMux) GetStreamable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Multiplex audio and video into a ISML file
type GstISMLMux struct {
	*GstBaseQTMux
}

func NewISMLMux() (*GstISMLMux, error) {
	e, err := gst.NewElement("ismlmux")
	if err != nil {
		return nil, err
	}
	return &GstISMLMux{GstBaseQTMux: &GstBaseQTMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

func NewISMLMuxWithName(name string) (*GstISMLMux, error) {
	e, err := gst.NewElementWithName("ismlmux", name)
	if err != nil {
		return nil, err
	}
	return &GstISMLMux{GstBaseQTMux: &GstBaseQTMux{GstAggregator: &GstAggregator{Element: e}}}, nil
}

// If set to true, the output should be as if it is to be streamed and hence no indexes written or duration written.
// Default: true
func (e *GstISMLMux) SetStreamable(value bool) error {
	return e.Element.SetProperty("streamable", value)
}

func (e *GstISMLMux) GetStreamable() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("streamable")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstBaseQTMux struct {
	*GstAggregator
}

// Reports the approximate amount of remaining moov header space reserved using reserved-max-duration
// Default: 18446744073709551615, Min: 0, Max: 18446744073709551615
func (e *GstBaseQTMux) GetReservedDurationRemaining() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("reserved-duration-remaining")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// When set to a value > 0, reserves space for index tables at the beginning of the file.
// Default: 18446744073709551615, Min: 0, Max: 18446744073709551615
func (e *GstBaseQTMux) SetReservedMaxDuration(value uint64) error {
	return e.Element.SetProperty("reserved-max-duration", value)
}

func (e *GstBaseQTMux) GetReservedMaxDuration() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("reserved-max-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// File that will be used temporarily to store data from the stream when creating a faststart file. If null a filepath will be created automatically
// Default: NULL
func (e *GstBaseQTMux) SetFaststartFile(value string) error {
	return e.Element.SetProperty("faststart-file", value)
}

func (e *GstBaseQTMux) GetFaststartFile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("faststart-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Fragment durations in ms (produce a fragmented file if > 0)
// Default: 0, Min: 0, Max: -1
func (e *GstBaseQTMux) SetFragmentDuration(value uint) error {
	return e.Element.SetProperty("fragment-duration", value)
}

func (e *GstBaseQTMux) GetFragmentDuration() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("fragment-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// How to to write fragments to the file.  Only used when 'fragment-duration' is greater than 0
// Default: dash-or-mss (0)
func (e *GstBaseQTMux) SetFragmentMode(value GstQTMuxFragmentMode) error {
	e.Element.SetArg("fragment-mode", string(value))
	return nil
}

func (e *GstBaseQTMux) GetFragmentMode() (GstQTMuxFragmentMode, error) {
	var value GstQTMuxFragmentMode
	var ok bool
	v, err := e.Element.GetProperty("fragment-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQTMuxFragmentMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQTMuxFragmentMode")
	}
	return value, nil
}

// Interleave between streams in bytes
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstBaseQTMux) SetInterleaveBytes(value uint64) error {
	return e.Element.SetProperty("interleave-bytes", value)
}

func (e *GstBaseQTMux) GetInterleaveBytes() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("interleave-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// File to be used to store data for moov atom making movie file recovery possible in case of a crash during muxing. Null for disabled. (Experimental)
// Default: NULL
func (e *GstBaseQTMux) SetMoovRecoveryFile(value string) error {
	return e.Element.SetProperty("moov-recovery-file", value)
}

func (e *GstBaseQTMux) GetMoovRecoveryFile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("moov-recovery-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Threshold for creating an edit list for gaps at the start in nanoseconds
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstBaseQTMux) SetStartGapThreshold(value uint64) error {
	return e.Element.SetProperty("start-gap-threshold", value)
}

func (e *GstBaseQTMux) GetStartGapThreshold() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("start-gap-threshold")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Method to determine DTS time (DEPRECATED)
// Default: reorder (1)
func (e *GstBaseQTMux) SetDtsMethod(value GstQTMuxDtsMethods) error {
	e.Element.SetArg("dts-method", string(value))
	return nil
}

func (e *GstBaseQTMux) GetDtsMethod() (GstQTMuxDtsMethods, error) {
	var value GstQTMuxDtsMethods
	var ok bool
	v, err := e.Element.GetProperty("dts-method")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstQTMuxDtsMethods)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstQTMuxDtsMethods")
	}
	return value, nil
}

// If the file should be formatted for faststart (headers first)
// Default: false
func (e *GstBaseQTMux) SetFaststart(value bool) error {
	return e.Element.SetProperty("faststart", value)
}

func (e *GstBaseQTMux) GetFaststart() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("faststart")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Interleave between streams in nanoseconds
// Default: 250000000, Min: 0, Max: 18446744073709551615
func (e *GstBaseQTMux) SetInterleaveTime(value uint64) error {
	return e.Element.SetProperty("interleave-time", value)
}

func (e *GstBaseQTMux) GetInterleaveTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("interleave-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Maximum allowed drift of raw audio samples vs. timestamps in nanoseconds
// Default: 40000000, Min: 0, Max: 18446744073709551615
func (e *GstBaseQTMux) SetMaxRawAudioDrift(value uint64) error {
	return e.Element.SetProperty("max-raw-audio-drift", value)
}

func (e *GstBaseQTMux) GetMaxRawAudioDrift() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-raw-audio-drift")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// When used with reserved-max-duration, periodically updates the index tables with information muxed so far.
// Default: 18446744073709551615, Min: 0, Max: 18446744073709551615
func (e *GstBaseQTMux) SetReservedMoovUpdatePeriod(value uint64) error {
	return e.Element.SetProperty("reserved-moov-update-period", value)
}

func (e *GstBaseQTMux) GetReservedMoovUpdatePeriod() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("reserved-moov-update-period")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Force multiple chunks to be created even for single-stream files
// Default: false
func (e *GstBaseQTMux) SetForceChunks(value bool) error {
	return e.Element.SetProperty("force-chunks", value)
}

func (e *GstBaseQTMux) GetForceChunks() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-chunks")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Timescale to use in the movie (units per second, 0 == default)
// Default: 0, Min: 0, Max: -1
func (e *GstBaseQTMux) SetMovieTimescale(value uint) error {
	return e.Element.SetProperty("movie-timescale", value)
}

func (e *GstBaseQTMux) GetMovieTimescale() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("movie-timescale")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Prefill samples table of reserved duration
// Default: false
func (e *GstBaseQTMux) SetReservedPrefill(value bool) error {
	return e.Element.SetProperty("reserved-prefill", value)
}

func (e *GstBaseQTMux) GetReservedPrefill() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("reserved-prefill")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Timescale to use for the tracks (units per second, 0 is automatic)
// Default: 0, Min: 0, Max: -1
func (e *GstBaseQTMux) SetTrakTimescale(value uint) error {
	return e.Element.SetProperty("trak-timescale", value)
}

func (e *GstBaseQTMux) GetTrakTimescale() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("trak-timescale")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Create a timecode trak even in unsupported flavors
// Default: false
func (e *GstBaseQTMux) SetForceCreateTimecodeTrak(value bool) error {
	return e.Element.SetProperty("force-create-timecode-trak", value)
}

func (e *GstBaseQTMux) GetForceCreateTimecodeTrak() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-create-timecode-trak")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Calculate and include presentation/composition time (in addition to decoding time)
// Default: true
func (e *GstBaseQTMux) SetPresentationTime(value bool) error {
	return e.Element.SetProperty("presentation-time", value)
}

func (e *GstBaseQTMux) GetPresentationTime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("presentation-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Multiplier for converting reserved-max-duration into bytes of header to reserve, per second, per track
// Default: 550, Min: 0, Max: 10000
func (e *GstBaseQTMux) SetReservedBytesPerSec(value uint) error {
	return e.Element.SetProperty("reserved-bytes-per-sec", value)
}

func (e *GstBaseQTMux) GetReservedBytesPerSec() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("reserved-bytes-per-sec")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

type GstQTMuxDtsMethods string

const (
	GstQTMuxDtsMethodsDd      GstQTMuxDtsMethods = "dd"      // delta/duration
	GstQTMuxDtsMethodsReorder GstQTMuxDtsMethods = "reorder" // reorder
	GstQTMuxDtsMethodsAsc     GstQTMuxDtsMethods = "asc"     // ascending
)

type GstQTMuxFragmentMode string

const (
	GstQTMuxFragmentModeDashOrMss             GstQTMuxFragmentMode = "dash-or-mss"              // Dash or Smoothstreaming
	GstQTMuxFragmentModeFirstMoovThenFinalise GstQTMuxFragmentMode = "first-moov-then-finalise" // First MOOV Fragment Then Finalise
)
