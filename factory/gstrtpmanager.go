// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Client-to-Mixer Audio Level Indication (RFC6464) RTP Header Extension
type GstRTPHeaderExtensionClientAudioLevel struct {
	*GstRTPHeaderExtension
}

func NewRTPHeaderExtensionClientAudioLevel() (*GstRTPHeaderExtensionClientAudioLevel, error) {
	e, err := gst.NewElement("rtphdrextclientaudiolevel")
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionClientAudioLevel{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

func NewRTPHeaderExtensionClientAudioLevelWithName(name string) (*GstRTPHeaderExtensionClientAudioLevel, error) {
	e, err := gst.NewElementWithName("rtphdrextclientaudiolevel", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionClientAudioLevel{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

// If the vad extension attribute is enabled or not
// Default: true
func (e *GstRTPHeaderExtensionClientAudioLevel) GetVad() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("vad")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Retransmit RTP packets when needed, according to RFC4588
type GstRtpRtxSend struct {
	*gst.Element
}

func NewRtpRtxSend() (*GstRtpRtxSend, error) {
	e, err := gst.NewElement("rtprtxsend")
	if err != nil {
		return nil, err
	}
	return &GstRtpRtxSend{Element: e}, nil
}

func NewRtpRtxSendWithName(name string) (*GstRtpRtxSend, error) {
	e, err := gst.NewElementWithName("rtprtxsend", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpRtxSend{Element: e}, nil
}

//	Number of retransmission packets sent
//
// Default: 0, Min: 0, Max: -1
func (e *GstRtpRtxSend) GetNumRtxPackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-rtx-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of retransmission events received
// Default: 0, Min: 0, Max: -1
func (e *GstRtpRtxSend) GetNumRtxRequests() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-rtx-requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Map of original payload types to their retransmission payload types

func (e *GstRtpRtxSend) SetPayloadTypeMap(value *gst.Structure) error {
	return e.Element.SetProperty("payload-type-map", value)
}

func (e *GstRtpRtxSend) GetPayloadTypeMap() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("payload-type-map")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Map of SSRCs to their retransmission SSRCs for SSRC-multiplexed mode (default = random)

func (e *GstRtpRtxSend) SetSsrcMap(value *gst.Structure) error {
	return e.Element.SetProperty("ssrc-map", value)
}

// Map of payload types to their clock rates

func (e *GstRtpRtxSend) SetClockRateMap(value *gst.Structure) error {
	return e.Element.SetProperty("clock-rate-map", value)
}

func (e *GstRtpRtxSend) GetClockRateMap() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("clock-rate-map")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Amount of packets to queue (0 = unlimited)
// Default: 100, Min: 0, Max: 32767
func (e *GstRtpRtxSend) SetMaxSizePackets(value uint) error {
	return e.Element.SetProperty("max-size-packets", value)
}

func (e *GstRtpRtxSend) GetMaxSizePackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Amount of ms to queue (0 = unlimited)
// Default: 0, Min: 0, Max: -1
func (e *GstRtpRtxSend) SetMaxSizeTime(value uint) error {
	return e.Element.SetProperty("max-size-time", value)
}

func (e *GstRtpRtxSend) GetMaxSizeTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Real-Time Transport Protocol bin
type GstRtpBin struct {
	*gst.Bin
}

func NewRtpBin() (*GstRtpBin, error) {
	e, err := gst.NewElement("rtpbin")
	if err != nil {
		return nil, err
	}
	return &GstRtpBin{Bin: &gst.Bin{Element: e}}, nil
}

func NewRtpBinWithName(name string) (*GstRtpBin, error) {
	e, err := gst.NewElementWithName("rtpbin", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpBin{Bin: &gst.Bin{Element: e}}, nil
}

// The maximum number of nanoseconds per frame that time stamp offsets may be adjusted (0 = no limit).
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstRtpBin) SetMaxTsOffsetAdjustment(value uint64) error {
	return e.Element.SetProperty("max-ts-offset-adjustment", value)
}

func (e *GstRtpBin) GetMaxTsOffsetAdjustment() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-ts-offset-adjustment")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The minimum absolute value of the time offset in (nanoseconds). Used to set an lower limit for when a time offset is deemed large enough to be useful for sync corrections.Note, if the ntp-sync parameter is set the default value is changed to 0 (no limit)
// Default: 4000000, Min: 0, Max: 18446744073709551615
func (e *GstRtpBin) SetMinTsOffset(value uint64) error {
	return e.Element.SetProperty("min-ts-offset", value)
}

func (e *GstRtpBin) GetMinTsOffset() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("min-ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// RTCP SR / NTP-64 interval synchronization (ms) (0 = always)
// Default: 0, Min: 0, Max: -1
func (e *GstRtpBin) SetRtcpSyncInterval(value uint) error {
	return e.Element.SetProperty("rtcp-sync-interval", value)
}

func (e *GstRtpBin) GetRtcpSyncInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rtcp-sync-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Synchronize received streams to the RFC7273 clock (requires clock and offset to be provided)
// Default: false
func (e *GstRtpBin) SetRfc7273Sync(value bool) error {
	return e.Element.SetProperty("rfc7273-sync", value)
}

func (e *GstRtpBin) GetRfc7273Sync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rfc7273-sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The SDES items of this session

func (e *GstRtpBin) SetSdes(value *gst.Structure) error {
	return e.Element.SetProperty("sdes", value)
}

func (e *GstRtpBin) GetSdes() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("sdes")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Tells the jitterbuffer to never exceed the given latency in size
// Default: false
func (e *GstRtpBin) SetDropOnLatency(value bool) error {
	return e.Element.SetProperty("drop-on-latency", value)
}

func (e *GstRtpBin) GetDropOnLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-on-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Do not demultiplex based on PT values
// Default: false
func (e *GstRtpBin) SetIgnorePt(value bool) error {
	return e.Element.SetProperty("ignore-pt", value)
}

func (e *GstRtpBin) GetIgnorePt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-pt")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The maximum number of streams to create for one session
// Default: -1, Min: 0, Max: -1
func (e *GstRtpBin) SetMaxStreams(value uint) error {
	return e.Element.SetProperty("max-streams", value)
}

func (e *GstRtpBin) GetMaxStreams() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Automatically remove timed out sources
// Default: false
func (e *GstRtpBin) SetAutoremove(value bool) error {
	return e.Element.SetProperty("autoremove", value)
}

func (e *GstRtpBin) GetAutoremove() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("autoremove")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Send an event downstream when a packet is lost
// Default: false
func (e *GstRtpBin) SetDoLost(value bool) error {
	return e.Element.SetProperty("do-lost", value)
}

func (e *GstRtpBin) GetDoLost() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-lost")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The maximum absolute value of the time offset in (nanoseconds). Note, if the ntp-sync parameter is set the default value is changed to 0 (no limit)
// Default: 3000000000, Min: 0, Max: 9223372036854775807
func (e *GstRtpBin) SetMaxTsOffset(value int64) error {
	return e.Element.SetProperty("max-ts-offset", value)
}

func (e *GstRtpBin) GetMaxTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

// Default RTP profile of newly created sessions
// Default: avp (1)
func (e *GstRtpBin) SetRtpProfile(value interface{}) error {
	return e.Element.SetProperty("rtp-profile", value)
}

func (e *GstRtpBin) GetRtpProfile() (interface{}, error) {
	return e.Element.GetProperty("rtp-profile")
}

// Whether RTP NTP header extension should be updated with actual NTP time
// Default: true
func (e *GstRtpBin) SetUpdateNtp64HeaderExt(value bool) error {
	return e.Element.SetProperty("update-ntp64-header-ext", value)
}

func (e *GstRtpBin) GetUpdateNtp64HeaderExt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("update-ntp64-header-ext")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Enable retransmission on all streams
// Default: false
func (e *GstRtpBin) SetDoRetransmission(value bool) error {
	return e.Element.SetProperty("do-retransmission", value)
}

func (e *GstRtpBin) GetDoRetransmission() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-retransmission")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// GstStructure mapping from session index to FEC encoder factory, eg fec-encoders='fec,0="rtpst2022-1-fecenc\ rows\=5\ columns\=5";'
// Default: application/x-rtp-fec-encoders;
func (e *GstRtpBin) SetFecEncoders(value *gst.Structure) error {
	return e.Element.SetProperty("fec-encoders", value)
}

func (e *GstRtpBin) GetFecEncoders() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("fec-encoders")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Use of RTCP SR in synchronization
// Default: always (0)
func (e *GstRtpBin) SetRtcpSync(value GstRTCPSync) error {
	e.Element.SetArg("rtcp-sync", string(value))
	return nil
}

func (e *GstRtpBin) GetRtcpSync() (GstRTCPSync, error) {
	var value GstRTCPSync
	var ok bool
	v, err := e.Element.GetProperty("rtcp-sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRTCPSync)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRTCPSync")
	}
	return value, nil
}

// Use the pipeline running-time to set the NTP time in the RTCP SR messages (DEPRECATED: Use ntp-time-source property)
// Default: false
func (e *GstRtpBin) SetUsePipelineClock(value bool) error {
	return e.Element.SetProperty("use-pipeline-clock", value)
}

func (e *GstRtpBin) GetUsePipelineClock() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-pipeline-clock")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Maximum amount of time in ms that the RTP time in RTCP SRs is allowed to be ahead (-1 disabled)
// Default: 1000, Min: -1, Max: 2147483647
func (e *GstRtpBin) SetMaxRtcpRtpTimeDiff(value int) error {
	return e.Element.SetProperty("max-rtcp-rtp-time-diff", value)
}

func (e *GstRtpBin) GetMaxRtcpRtpTimeDiff() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-rtcp-rtp-time-diff")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Whether sources that don't receive RTP or RTCP packets for longer than 5x RTCP interval should be removed
// Default: true
func (e *GstRtpBin) SetTimeoutInactiveSources(value bool) error {
	return e.Element.SetProperty("timeout-inactive-sources", value)
}

func (e *GstRtpBin) GetTimeoutInactiveSources() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("timeout-inactive-sources")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Sets a smoothing factor for the timestamp offset in number of values for a calculated running moving average. (0 = no smoothing factor)
// Default: 0, Min: 0, Max: -1
func (e *GstRtpBin) SetTsOffsetSmoothingFactor(value uint) error {
	return e.Element.SetProperty("ts-offset-smoothing-factor", value)
}

func (e *GstRtpBin) GetTsOffsetSmoothingFactor() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ts-offset-smoothing-factor")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Add Reference Timestamp Meta to buffers with the original clock timestamp before any adjustments when syncing to an RFC7273 clock or after clock synchronization via RTCP or inband NTP-64 header extensions has happened.
// Default: false
func (e *GstRtpBin) SetAddReferenceTimestampMeta(value bool) error {
	return e.Element.SetProperty("add-reference-timestamp-meta", value)
}

func (e *GstRtpBin) GetAddReferenceTimestampMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("add-reference-timestamp-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Send event downstream when a stream is synchronized to the sender
// Default: false
func (e *GstRtpBin) SetDoSyncEvent(value bool) error {
	return e.Element.SetProperty("do-sync-event", value)
}

func (e *GstRtpBin) GetDoSyncEvent() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-sync-event")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Use send time or capture time for RTCP sync (TRUE = send time, FALSE = capture time)
// Default: true
func (e *GstRtpBin) SetRtcpSyncSendTime(value bool) error {
	return e.Element.SetProperty("rtcp-sync-send-time", value)
}

func (e *GstRtpBin) GetRtcpSyncSendTime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rtcp-sync-send-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Default amount of ms to buffer in the jitterbuffers
// Default: 200, Min: 0, Max: -1
func (e *GstRtpBin) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstRtpBin) GetLatency() (uint, error) {
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

// The maximum time (milliseconds) of missing packets tolerated.
// Default: 60000, Min: 0, Max: -1
func (e *GstRtpBin) SetMaxDropoutTime(value uint) error {
	return e.Element.SetProperty("max-dropout-time", value)
}

func (e *GstRtpBin) GetMaxDropoutTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-dropout-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Synchronize received streams to the NTP clock
// Default: false
func (e *GstRtpBin) SetNtpSync(value bool) error {
	return e.Element.SetProperty("ntp-sync", value)
}

func (e *GstRtpBin) GetNtpSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ntp-sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// NTP time source for RTCP packets
// Default: ntp (0)
func (e *GstRtpBin) SetNtpTimeSource(value GstRtpNtpTimeSource) error {
	e.Element.SetArg("ntp-time-source", string(value))
	return nil
}

func (e *GstRtpBin) GetNtpTimeSource() (GstRtpNtpTimeSource, error) {
	var value GstRtpNtpTimeSource
	var ok bool
	v, err := e.Element.GetProperty("ntp-time-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRtpNtpTimeSource)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRtpNtpTimeSource")
	}
	return value, nil
}

// Control the buffering algorithm in use
// Default: slave (1)
func (e *GstRtpBin) SetBufferMode(value RTPJitterBufferMode) error {
	e.Element.SetArg("buffer-mode", string(value))
	return nil
}

func (e *GstRtpBin) GetBufferMode() (RTPJitterBufferMode, error) {
	var value RTPJitterBufferMode
	var ok bool
	v, err := e.Element.GetProperty("buffer-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(RTPJitterBufferMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to RTPJitterBufferMode")
	}
	return value, nil
}

// GstStructure mapping from session index to FEC decoder factory, eg fec-decoders='fec,0="rtpst2022-1-fecdec\ size-time\=1000000000";'
// Default: application/x-rtp-fec-decoders;
func (e *GstRtpBin) SetFecDecoders(value *gst.Structure) error {
	return e.Element.SetProperty("fec-decoders", value)
}

func (e *GstRtpBin) GetFecDecoders() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("fec-decoders")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// The maximum time (milliseconds) of misordered packets tolerated.
// Default: 2000, Min: 0, Max: -1
func (e *GstRtpBin) SetMaxMisorderTime(value uint) error {
	return e.Element.SetProperty("max-misorder-time", value)
}

func (e *GstRtpBin) GetMaxMisorderTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-misorder-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Extends RTP packets to add or retrieve a RepairedStreamId (RID) value as specified in RFC8852
type GstRTPHeaderExtensionRepairedStreamId struct {
	*GstRTPHeaderExtension
}

func NewRTPHeaderExtensionRepairedStreamId() (*GstRTPHeaderExtensionRepairedStreamId, error) {
	e, err := gst.NewElement("rtphdrextrepairedstreamid")
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionRepairedStreamId{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

func NewRTPHeaderExtensionRepairedStreamIdWithName(name string) (*GstRTPHeaderExtensionRepairedStreamId, error) {
	e, err := gst.NewElementWithName("rtphdrextrepairedstreamid", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionRepairedStreamId{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

// The RepairedRtpStreamId (RID) value last read or to write from/to RTP buffers
// Default: NULL
func (e *GstRTPHeaderExtensionRepairedStreamId) SetRid(value string) error {
	return e.Element.SetProperty("rid", value)
}

func (e *GstRTPHeaderExtensionRepairedStreamId) GetRid() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("rid")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// A buffer that deals with network jitter and other transmission faults
type GstRtpJitterBuffer struct {
	*gst.Element
}

func NewRtpJitterBuffer() (*GstRtpJitterBuffer, error) {
	e, err := gst.NewElement("rtpjitterbuffer")
	if err != nil {
		return nil, err
	}
	return &GstRtpJitterBuffer{Element: e}, nil
}

func NewRtpJitterBufferWithName(name string) (*GstRtpJitterBuffer, error) {
	e, err := gst.NewElementWithName("rtpjitterbuffer", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpJitterBuffer{Element: e}, nil
}

// Add Reference Timestamp Meta to buffers with the original clock timestamp before any adjustments when syncing to an RFC7273 clock or after clock synchronization via RTCP or inband NTP-64 header extensions has happened.
// Default: false
func (e *GstRtpJitterBuffer) SetAddReferenceTimestampMeta(value bool) error {
	return e.Element.SetProperty("add-reference-timestamp-meta", value)
}

func (e *GstRtpJitterBuffer) GetAddReferenceTimestampMeta() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("add-reference-timestamp-meta")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The maximum number of retries to request a retransmission. (-1 not limited)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRtpJitterBuffer) SetRtxMaxRetries(value int) error {
	return e.Element.SetProperty("rtx-max-retries", value)
}

func (e *GstRtpJitterBuffer) GetRtxMaxRetries() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rtx-max-retries")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Adjust buffer timestamps with offset in nanoseconds
// Default: 0, Min: -9223372036854775808, Max: 9223372036854775807
func (e *GstRtpJitterBuffer) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

func (e *GstRtpJitterBuffer) GetTsOffset() (int64, error) {
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

// RTCP SR / NTP-64 interval synchronization (ms) (0 = always)
// Default: 0, Min: 0, Max: -1
func (e *GstRtpJitterBuffer) SetSyncInterval(value uint) error {
	return e.Element.SetProperty("sync-interval", value)
}

func (e *GstRtpJitterBuffer) GetSyncInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sync-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Tells the jitterbuffer to never exceed the given latency in size
// Default: false
func (e *GstRtpJitterBuffer) SetDropOnLatency(value bool) error {
	return e.Element.SetProperty("drop-on-latency", value)
}

func (e *GstRtpJitterBuffer) GetDropOnLatency() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("drop-on-latency")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The number of consecutive packets needed to start (set to 0 to disable faststart. The jitterbuffer will by default start after the latency has elapsed)
// Default: 0, Min: 0, Max: -1
func (e *GstRtpJitterBuffer) SetFaststartMinPackets(value uint) error {
	return e.Element.SetProperty("faststart-min-packets", value)
}

func (e *GstRtpJitterBuffer) GetFaststartMinPackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("faststart-min-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The maximum time (milliseconds) of misordered packets tolerated.
// Default: 2000, Min: 0, Max: -1
func (e *GstRtpJitterBuffer) SetMaxMisorderTime(value uint) error {
	return e.Element.SetProperty("max-misorder-time", value)
}

func (e *GstRtpJitterBuffer) GetMaxMisorderTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-misorder-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Maximum amount of time in ms that the RTP time in RTCP SRs is allowed to be ahead (-1 disabled)
// Default: 1000, Min: -1, Max: 2147483647
func (e *GstRtpJitterBuffer) SetMaxRtcpRtpTimeDiff(value int) error {
	return e.Element.SetProperty("max-rtcp-rtp-time-diff", value)
}

func (e *GstRtpJitterBuffer) GetMaxRtcpRtpTimeDiff() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-rtcp-rtp-time-diff")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Try to get a retransmission for this many ms (-1 automatic)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRtpJitterBuffer) SetRtxRetryPeriod(value int) error {
	return e.Element.SetProperty("rtx-retry-period", value)
}

func (e *GstRtpJitterBuffer) GetRtxRetryPeriod() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rtx-retry-period")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The time to wait for a retransmitted packet after it has been considered lost in order to collect statistics (ms)
// Default: 1000, Min: 0, Max: -1
func (e *GstRtpJitterBuffer) SetRtxStatsTimeout(value uint) error {
	return e.Element.SetProperty("rtx-stats-timeout", value)
}

func (e *GstRtpJitterBuffer) GetRtxStatsTimeout() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rtx-stats-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Various statistics
// Default: application/x-rtp-jitterbuffer-stats, num-pushed=(guint64)0, num-lost=(guint64)0, num-late=(guint64)0, num-duplicates=(guint64)0, avg-jitter=(guint64)0, rtx-count=(guint64)0, rtx-success-count=(guint64)0, rtx-per-packet=(double)0, rtx-rtt=(guint64)0;
func (e *GstRtpJitterBuffer) GetStats() (*gst.Structure, error) {
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

// The maximum number of nanoseconds per frame that time stamp offsets may be adjusted (0 = no limit).
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstRtpJitterBuffer) SetMaxTsOffsetAdjustment(value uint64) error {
	return e.Element.SetProperty("max-ts-offset-adjustment", value)
}

func (e *GstRtpJitterBuffer) GetMaxTsOffsetAdjustment() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-ts-offset-adjustment")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Minimum timeout between sending a transmission event in ms (-1 automatic)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRtpJitterBuffer) SetRtxMinRetryTimeout(value int) error {
	return e.Element.SetProperty("rtx-min-retry-timeout", value)
}

func (e *GstRtpJitterBuffer) GetRtxMinRetryTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rtx-min-retry-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Retry sending a transmission event after this timeout in ms (-1 automatic)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRtpJitterBuffer) SetRtxRetryTimeout(value int) error {
	return e.Element.SetProperty("rtx-retry-timeout", value)
}

func (e *GstRtpJitterBuffer) GetRtxRetryTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rtx-retry-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The buffer filled percent
// Default: 0, Min: 0, Max: 100
func (e *GstRtpJitterBuffer) GetPercent() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("percent")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Synchronize received streams to the RFC7273 clock (requires clock and offset to be provided)
// Default: false
func (e *GstRtpJitterBuffer) SetRfc7273Sync(value bool) error {
	return e.Element.SetProperty("rfc7273-sync", value)
}

func (e *GstRtpJitterBuffer) GetRfc7273Sync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rfc7273-sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Extra time in ms to wait before sending retransmission event (-1 automatic)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRtpJitterBuffer) SetRtxDelay(value int) error {
	return e.Element.SetProperty("rtx-delay", value)
}

func (e *GstRtpJitterBuffer) GetRtxDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rtx-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Send an event downstream when a packet is lost
// Default: false
func (e *GstRtpJitterBuffer) SetDoLost(value bool) error {
	return e.Element.SetProperty("do-lost", value)
}

func (e *GstRtpJitterBuffer) GetDoLost() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-lost")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Send retransmission events upstream when a packet is late
// Default: false
func (e *GstRtpJitterBuffer) SetDoRetransmission(value bool) error {
	return e.Element.SetProperty("do-retransmission", value)
}

func (e *GstRtpJitterBuffer) GetDoRetransmission() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("do-retransmission")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Minimal time between posting dropped packet messages
// Default: 200, Min: 0, Max: -1
func (e *GstRtpJitterBuffer) SetDropMessagesInterval(value uint) error {
	return e.Element.SetProperty("drop-messages-interval", value)
}

func (e *GstRtpJitterBuffer) GetDropMessagesInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("drop-messages-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The maximum time (milliseconds) of missing packets tolerated.
// Default: 60000, Min: 0, Max: 2147483647
func (e *GstRtpJitterBuffer) SetMaxDropoutTime(value uint) error {
	return e.Element.SetProperty("max-dropout-time", value)
}

func (e *GstRtpJitterBuffer) GetMaxDropoutTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-dropout-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The deadline for a valid RTX request in milliseconds. (-1 automatic)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRtpJitterBuffer) SetRtxDeadline(value int) error {
	return e.Element.SetProperty("rtx-deadline", value)
}

func (e *GstRtpJitterBuffer) GetRtxDeadline() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rtx-deadline")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Minimum time in ms to wait before sending retransmission event
// Default: 0, Min: 0, Max: -1
func (e *GstRtpJitterBuffer) SetRtxMinDelay(value uint) error {
	return e.Element.SetProperty("rtx-min-delay", value)
}

func (e *GstRtpJitterBuffer) GetRtxMinDelay() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rtx-min-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Control the buffering algorithm in use
// Default: slave (1)
func (e *GstRtpJitterBuffer) SetMode(value RTPJitterBufferMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstRtpJitterBuffer) GetMode() (RTPJitterBufferMode, error) {
	var value RTPJitterBufferMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(RTPJitterBufferMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to RTPJitterBufferMode")
	}
	return value, nil
}

// Post a custom message to the bus when a packet is dropped by the jitterbuffer
// Default: false
func (e *GstRtpJitterBuffer) SetPostDropMessages(value bool) error {
	return e.Element.SetProperty("post-drop-messages", value)
}

func (e *GstRtpJitterBuffer) GetPostDropMessages() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-drop-messages")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Estimate when the next packet should arrive and schedule a retransmission request for it.
// Default: true
func (e *GstRtpJitterBuffer) SetRtxNextSeqnum(value bool) error {
	return e.Element.SetProperty("rtx-next-seqnum", value)
}

func (e *GstRtpJitterBuffer) GetRtxNextSeqnum() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rtx-next-seqnum")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Amount of ms to buffer
// Default: 200, Min: 0, Max: -1
func (e *GstRtpJitterBuffer) SetLatency(value uint) error {
	return e.Element.SetProperty("latency", value)
}

func (e *GstRtpJitterBuffer) GetLatency() (uint, error) {
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

// Sending retransmission event when this much reordering (0 disable)
// Default: 3, Min: -1, Max: 2147483647
func (e *GstRtpJitterBuffer) SetRtxDelayReorder(value int) error {
	return e.Element.SetProperty("rtx-delay-reorder", value)
}

func (e *GstRtpJitterBuffer) GetRtxDelayReorder() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rtx-delay-reorder")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Parses codec streams transmitted in the same RTP session
type GstRtpPtDemux struct {
	*gst.Element
}

func NewRtpPtDemux() (*GstRtpPtDemux, error) {
	e, err := gst.NewElement("rtpptdemux")
	if err != nil {
		return nil, err
	}
	return &GstRtpPtDemux{Element: e}, nil
}

func NewRtpPtDemuxWithName(name string) (*GstRtpPtDemux, error) {
	e, err := gst.NewElementWithName("rtpptdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpPtDemux{Element: e}, nil
}

// Packets with these payload types will be dropped

func (e *GstRtpPtDemux) SetIgnoredPayloadTypes(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("ignored-payload-types", value)
}

func (e *GstRtpPtDemux) GetIgnoredPayloadTypes() (*gst.ValueArrayValue, error) {
	var value *gst.ValueArrayValue
	var ok bool
	v, err := e.Element.GetProperty("ignored-payload-types")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.ValueArrayValue)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.ValueArrayValue")
	}
	return value, nil
}

// Implement an RTP session
type GstRtpSession struct {
	*gst.Element
}

func NewRtpSession() (*GstRtpSession, error) {
	e, err := gst.NewElement("rtpsession")
	if err != nil {
		return nil, err
	}
	return &GstRtpSession{Element: e}, nil
}

func NewRtpSessionWithName(name string) (*GstRtpSession, error) {
	e, err := gst.NewElementWithName("rtpsession", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpSession{Element: e}, nil
}

// The maximum time (milliseconds) of misordered packets tolerated.
// Default: 2000, Min: 0, Max: -1
func (e *GstRtpSession) SetMaxMisorderTime(value uint) error {
	return e.Element.SetProperty("max-misorder-time", value)
}

func (e *GstRtpSession) GetMaxMisorderTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-misorder-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The RTCP bandwidth of the session in bytes per second (or as a real fraction of the RTP bandwidth if < 1.0)
// Default: 0.05, Min: 0, Max: 1.79769e+308
func (e *GstRtpSession) SetRtcpFraction(value float64) error {
	return e.Element.SetProperty("rtcp-fraction", value)
}

func (e *GstRtpSession) GetRtcpFraction() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("rtcp-fraction")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// Use send time or capture time for RTCP sync (TRUE = send time, FALSE = capture time)
// Default: true
func (e *GstRtpSession) SetRtcpSyncSendTime(value bool) error {
	return e.Element.SetProperty("rtcp-sync-send-time", value)
}

func (e *GstRtpSession) GetRtcpSyncSendTime() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rtcp-sync-send-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether sources that don't receive RTP or RTCP packets for longer than 5x RTCP interval should be removed
// Default: true
func (e *GstRtpSession) SetTimeoutInactiveSources(value bool) error {
	return e.Element.SetProperty("timeout-inactive-sources", value)
}

func (e *GstRtpSession) GetTimeoutInactiveSources() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("timeout-inactive-sources")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether RTP NTP header extension should be updated with actual NTP time
// Default: true
func (e *GstRtpSession) SetUpdateNtp64HeaderExt(value bool) error {
	return e.Element.SetProperty("update-ntp64-header-ext", value)
}

func (e *GstRtpSession) GetUpdateNtp64HeaderExt() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("update-ntp64-header-ext")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Use the pipeline running-time to set the NTP time in the RTCP SR messages (DEPRECATED: Use ntp-time-source property)
// Default: false
func (e *GstRtpSession) SetUsePipelineClock(value bool) error {
	return e.Element.SetProperty("use-pipeline-clock", value)
}

func (e *GstRtpSession) GetUsePipelineClock() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-pipeline-clock")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The internal RTPSession object

func (e *GstRtpSession) GetInternalSession() (RTPSession, error) {
	var value RTPSession
	var ok bool
	v, err := e.Element.GetProperty("internal-session")
	if err != nil {
		return value, err
	}
	value, ok = v.(RTPSession)
	if !ok {
		return value, fmt.Errorf("could not cast value to RTPSession")
	}
	return value, nil
}

// The number of active sources in the session
// Default: 0, Min: 0, Max: -1
func (e *GstRtpSession) GetNumActiveSources() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-active-sources")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Consecutive packet sequence numbers to accept the source
// Default: 2, Min: 0, Max: -1
func (e *GstRtpSession) SetProbation(value uint) error {
	return e.Element.SetProperty("probation", value)
}

func (e *GstRtpSession) GetProbation() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("probation")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// RTP profile to use
// Default: avp (1)
func (e *GstRtpSession) SetRtpProfile(value interface{}) error {
	return e.Element.SetProperty("rtp-profile", value)
}

func (e *GstRtpSession) GetRtpProfile() (interface{}, error) {
	return e.Element.GetProperty("rtp-profile")
}

// NTP time source for RTCP packets
// Default: ntp (0)
func (e *GstRtpSession) SetNtpTimeSource(value GstRtpNtpTimeSource) error {
	e.Element.SetArg("ntp-time-source", string(value))
	return nil
}

func (e *GstRtpSession) GetNtpTimeSource() (GstRtpNtpTimeSource, error) {
	var value GstRtpNtpTimeSource
	var ok bool
	v, err := e.Element.GetProperty("ntp-time-source")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRtpNtpTimeSource)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRtpNtpTimeSource")
	}
	return value, nil
}

// The number of sources in the session
// Default: 0, Min: 0, Max: -1
func (e *GstRtpSession) GetNumSources() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-sources")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The RTCP bandwidth used for receivers in bytes per second (-1 = default)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRtpSession) SetRtcpRrBandwidth(value int) error {
	return e.Element.SetProperty("rtcp-rr-bandwidth", value)
}

func (e *GstRtpSession) GetRtcpRrBandwidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rtcp-rr-bandwidth")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Various statistics
// Default: application/x-rtp-session-stats, rtx-drop-count=(uint)0, sent-nack-count=(uint)0, recv-nack-count=(uint)0, source-stats=(GValueArray)<  >, rtx-count=(uint)0, recv-rtx-req-count=(uint)0, sent-rtx-req-count=(uint)0;
func (e *GstRtpSession) GetStats() (*gst.Structure, error) {
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

// The maximum time (milliseconds) of missing packets tolerated.
// Default: 60000, Min: 0, Max: -1
func (e *GstRtpSession) SetMaxDropoutTime(value uint) error {
	return e.Element.SetProperty("max-dropout-time", value)
}

func (e *GstRtpSession) GetMaxDropoutTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-dropout-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Minimum interval between Regular RTCP packet (in ns)
// Default: 5000000000, Min: 0, Max: 18446744073709551615
func (e *GstRtpSession) SetRtcpMinInterval(value uint64) error {
	return e.Element.SetProperty("rtcp-min-interval", value)
}

func (e *GstRtpSession) GetRtcpMinInterval() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("rtcp-min-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// The RTCP bandwidth used for senders in bytes per second (-1 = default)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRtpSession) SetRtcpRsBandwidth(value int) error {
	return e.Element.SetProperty("rtcp-rs-bandwidth", value)
}

func (e *GstRtpSession) GetRtcpRsBandwidth() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("rtcp-rs-bandwidth")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The SDES items of this session

func (e *GstRtpSession) SetSdes(value *gst.Structure) error {
	return e.Element.SetProperty("sdes", value)
}

func (e *GstRtpSession) GetSdes() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("sdes")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Various statistics from TWCC

func (e *GstRtpSession) GetTwccStats() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("twcc-stats")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// The bandwidth of the session in bytes per second (0 for auto-discover)
// Default: 0, Min: 0, Max: 1.79769e+308
func (e *GstRtpSession) SetBandwidth(value float64) error {
	return e.Element.SetProperty("bandwidth", value)
}

func (e *GstRtpSession) GetBandwidth() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("bandwidth")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

// performs FEC as described by SMPTE 2022-1
type GstRTPST20221FecEnc struct {
	*gst.Element
}

func NewRTPST20221FecEnc() (*GstRTPST20221FecEnc, error) {
	e, err := gst.NewElement("rtpst2022-1-fecenc")
	if err != nil {
		return nil, err
	}
	return &GstRTPST20221FecEnc{Element: e}, nil
}

func NewRTPST20221FecEncWithName(name string) (*GstRTPST20221FecEnc, error) {
	e, err := gst.NewElementWithName("rtpst2022-1-fecenc", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPST20221FecEnc{Element: e}, nil
}

// Number of columns to apply row FEC on, 0=disabled
// Default: 0, Min: 0, Max: 255
func (e *GstRTPST20221FecEnc) SetColumns(value uint) error {
	return e.Element.SetProperty("columns", value)
}

func (e *GstRTPST20221FecEnc) GetColumns() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("columns")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Whether the encoder should compute and send column FEC
// Default: true
func (e *GstRTPST20221FecEnc) SetEnableColumnFec(value bool) error {
	return e.Element.SetProperty("enable-column-fec", value)
}

func (e *GstRTPST20221FecEnc) GetEnableColumnFec() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-column-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Whether the encoder should compute and send row FEC
// Default: true
func (e *GstRTPST20221FecEnc) SetEnableRowFec(value bool) error {
	return e.Element.SetProperty("enable-row-fec", value)
}

func (e *GstRTPST20221FecEnc) GetEnableRowFec() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-row-fec")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The payload type of FEC packets
// Default: 96, Min: 96, Max: 255
func (e *GstRTPST20221FecEnc) SetPt(value int) error {
	return e.Element.SetProperty("pt", value)
}

func (e *GstRTPST20221FecEnc) GetPt() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("pt")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Number of rows to apply column FEC on, 0=disabled
// Default: 0, Min: 0, Max: 255
func (e *GstRTPST20221FecEnc) SetRows(value uint) error {
	return e.Element.SetProperty("rows", value)
}

func (e *GstRTPST20221FecEnc) GetRows() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("rows")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Extends RTP packets to add or retrieve a 64-bit NTP timestamp as specified in RFC6501
type GstRTPHeaderExtensionNtp64 struct {
	*GstRTPHeaderExtension
}

func NewRTPHeaderExtensionNtp64() (*GstRTPHeaderExtensionNtp64, error) {
	e, err := gst.NewElement("rtphdrextntp64")
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionNtp64{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

func NewRTPHeaderExtensionNtp64WithName(name string) (*GstRTPHeaderExtensionNtp64, error) {
	e, err := gst.NewElementWithName("rtphdrextntp64", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionNtp64{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

// Add the header extension to every packet
// Default: false
func (e *GstRTPHeaderExtensionNtp64) SetEveryPacket(value bool) error {
	return e.Element.SetProperty("every-packet", value)
}

func (e *GstRTPHeaderExtensionNtp64) GetEveryPacket() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("every-packet")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Interval between consecutive packets that get the header extension added
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstRTPHeaderExtensionNtp64) SetInterval(value uint64) error {
	return e.Element.SetProperty("interval", value)
}

func (e *GstRTPHeaderExtensionNtp64) GetInterval() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// Extends RTP packets to add sequence number transport wide.
type GstRTPHeaderExtensionTWCC struct {
	*GstRTPHeaderExtension
}

func NewRTPHeaderExtensionTWCC() (*GstRTPHeaderExtensionTWCC, error) {
	e, err := gst.NewElement("rtphdrexttwcc")
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionTWCC{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

func NewRTPHeaderExtensionTWCCWithName(name string) (*GstRTPHeaderExtensionTWCC, error) {
	e, err := gst.NewElementWithName("rtphdrexttwcc", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionTWCC{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

// The number of separate RTP streams this header applies to
// Default: 1, Min: 1, Max: -1
func (e *GstRTPHeaderExtensionTWCC) SetNStreams(value uint) error {
	return e.Element.SetProperty("n-streams", value)
}

func (e *GstRTPHeaderExtensionTWCC) GetNStreams() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("n-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// multiplex N rtp streams into one
type GstRTPMux struct {
	*gst.Element
}

func NewRTPMux() (*GstRTPMux, error) {
	e, err := gst.NewElement("rtpmux")
	if err != nil {
		return nil, err
	}
	return &GstRTPMux{Element: e}, nil
}

func NewRTPMuxWithName(name string) (*GstRTPMux, error) {
	e, err := gst.NewElementWithName("rtpmux", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPMux{Element: e}, nil
}

// The RTP sequence number of the last processed packet
// Default: 0, Min: 0, Max: -1
func (e *GstRTPMux) GetSeqnum() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("seqnum")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Offset to add to all outgoing seqnum (-1 = random)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRTPMux) SetSeqnumOffset(value int) error {
	return e.Element.SetProperty("seqnum-offset", value)
}

func (e *GstRTPMux) GetSeqnumOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("seqnum-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The SSRC of the packets (default == random)
// Default: -1, Min: 0, Max: -1
func (e *GstRTPMux) SetSsrc(value uint) error {
	return e.Element.SetProperty("ssrc", value)
}

func (e *GstRTPMux) GetSsrc() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("ssrc")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Offset to add to all outgoing timestamps (-1 = random)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRTPMux) SetTimestampOffset(value int) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

func (e *GstRTPMux) GetTimestampOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timestamp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Receive retransmitted RTP packets according to RFC4588
type GstRtpRtxReceive struct {
	*gst.Element
}

func NewRtpRtxReceive() (*GstRtpRtxReceive, error) {
	e, err := gst.NewElement("rtprtxreceive")
	if err != nil {
		return nil, err
	}
	return &GstRtpRtxReceive{Element: e}, nil
}

func NewRtpRtxReceiveWithName(name string) (*GstRtpRtxReceive, error) {
	e, err := gst.NewElementWithName("rtprtxreceive", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpRtxReceive{Element: e}, nil
}

// Map of original payload types to their retransmission payload types

func (e *GstRtpRtxReceive) SetPayloadTypeMap(value *gst.Structure) error {
	return e.Element.SetProperty("payload-type-map", value)
}

func (e *GstRtpRtxReceive) GetPayloadTypeMap() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("payload-type-map")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// Map of SSRCs to their retransmission SSRCs for SSRC-multiplexed mode

func (e *GstRtpRtxReceive) SetSsrcMap(value *gst.Structure) error {
	return e.Element.SetProperty("ssrc-map", value)
}

// Number of retransmission packets correctly associated with retransmission requests
// Default: 0, Min: 0, Max: -1
func (e *GstRtpRtxReceive) GetNumRtxAssocPackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-rtx-assoc-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

//	Number of retransmission packets received
//
// Default: 0, Min: 0, Max: -1
func (e *GstRtpRtxReceive) GetNumRtxPackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-rtx-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Number of retransmission events received
// Default: 0, Min: 0, Max: -1
func (e *GstRtpRtxReceive) GetNumRtxRequests() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("num-rtx-requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Extends RTP packets to add or retrieve a Media Identification (MID) value as specified in RFC8843
type GstRTPHeaderExtensionMid struct {
	*GstRTPHeaderExtension
}

func NewRTPHeaderExtensionMid() (*GstRTPHeaderExtensionMid, error) {
	e, err := gst.NewElement("rtphdrextmid")
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionMid{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

func NewRTPHeaderExtensionMidWithName(name string) (*GstRTPHeaderExtensionMid, error) {
	e, err := gst.NewElementWithName("rtphdrextmid", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionMid{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

// The Media Identification (MID) value last read or to write from/to RTP buffers
// Default: NULL
func (e *GstRTPHeaderExtensionMid) SetMid(value string) error {
	return e.Element.SetProperty("mid", value)
}

func (e *GstRTPHeaderExtensionMid) GetMid() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mid")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Funnel RTP buffers together for multiplexing
type GstRtpFunnel struct {
	*gst.Element
}

func NewRtpFunnel() (*GstRtpFunnel, error) {
	e, err := gst.NewElement("rtpfunnel")
	if err != nil {
		return nil, err
	}
	return &GstRtpFunnel{Element: e}, nil
}

func NewRtpFunnelWithName(name string) (*GstRtpFunnel, error) {
	e, err := gst.NewElementWithName("rtpfunnel", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpFunnel{Element: e}, nil
}

// Use the same RTP timestamp offset for all sinkpads (-1 = disable)
// Default: -1, Min: -1, Max: 2147483647
func (e *GstRtpFunnel) SetCommonTsOffset(value int) error {
	return e.Element.SetProperty("common-ts-offset", value)
}

func (e *GstRtpFunnel) GetCommonTsOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("common-ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Extends RTP packets to add or retrieve a RtpStreamId (RID) value as specified in RFC8852
type GstRTPHeaderExtensionStreamId struct {
	*GstRTPHeaderExtension
}

func NewRTPHeaderExtensionStreamId() (*GstRTPHeaderExtensionStreamId, error) {
	e, err := gst.NewElement("rtphdrextstreamid")
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionStreamId{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

func NewRTPHeaderExtensionStreamIdWithName(name string) (*GstRTPHeaderExtensionStreamId, error) {
	e, err := gst.NewElementWithName("rtphdrextstreamid", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionStreamId{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

// The RtpStreamId (RID) value last read or to write from/to RTP buffers
// Default: NULL
func (e *GstRTPHeaderExtensionStreamId) SetRid(value string) error {
	return e.Element.SetProperty("rid", value)
}

func (e *GstRTPHeaderExtensionStreamId) GetRid() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("rid")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Keep RTP packets in a queue for retransmission
type GstRTPRtxQueue struct {
	*gst.Element
}

func NewRTPRtxQueue() (*GstRTPRtxQueue, error) {
	e, err := gst.NewElement("rtprtxqueue")
	if err != nil {
		return nil, err
	}
	return &GstRTPRtxQueue{Element: e}, nil
}

func NewRTPRtxQueueWithName(name string) (*GstRTPRtxQueue, error) {
	e, err := gst.NewElementWithName("rtprtxqueue", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPRtxQueue{Element: e}, nil
}

// Number of fulfilled retransmission requests
// Default: 0, Min: 0, Max: -1
func (e *GstRTPRtxQueue) GetFulfilledRequests() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("fulfilled-requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Amount of packets to queue (0 = unlimited)
// Default: 100, Min: 0, Max: -1
func (e *GstRTPRtxQueue) SetMaxSizePackets(value uint) error {
	return e.Element.SetProperty("max-size-packets", value)
}

func (e *GstRTPRtxQueue) GetMaxSizePackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Amount of ms to queue (0 = unlimited)
// Default: 0, Min: 0, Max: -1
func (e *GstRTPRtxQueue) SetMaxSizeTime(value uint) error {
	return e.Element.SetProperty("max-size-time", value)
}

func (e *GstRTPRtxQueue) GetMaxSizeTime() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Total number of retransmission requests
// Default: 0, Min: 0, Max: -1
func (e *GstRTPRtxQueue) GetRequests() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("requests")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Splits RTP streams based on the SSRC
type GstRtpSsrcDemux struct {
	*gst.Element
}

func NewRtpSsrcDemux() (*GstRtpSsrcDemux, error) {
	e, err := gst.NewElement("rtpssrcdemux")
	if err != nil {
		return nil, err
	}
	return &GstRtpSsrcDemux{Element: e}, nil
}

func NewRtpSsrcDemuxWithName(name string) (*GstRtpSsrcDemux, error) {
	e, err := gst.NewElementWithName("rtpssrcdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpSsrcDemux{Element: e}, nil
}

// The maximum number of streams allowed
// Default: -1, Min: 0, Max: -1
func (e *GstRtpSsrcDemux) SetMaxStreams(value uint) error {
	return e.Element.SetProperty("max-streams", value)
}

func (e *GstRtpSsrcDemux) GetMaxStreams() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("max-streams")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// performs FEC as described by SMPTE 2022-1
type GstRTPST20221FecDec struct {
	*gst.Element
}

func NewRTPST20221FecDec() (*GstRTPST20221FecDec, error) {
	e, err := gst.NewElement("rtpst2022-1-fecdec")
	if err != nil {
		return nil, err
	}
	return &GstRTPST20221FecDec{Element: e}, nil
}

func NewRTPST20221FecDecWithName(name string) (*GstRTPST20221FecDec, error) {
	e, err := gst.NewElementWithName("rtpst2022-1-fecdec", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPST20221FecDec{Element: e}, nil
}

// The amount of data to store (in ns, 0-disable)
// Default: 1000000000, Min: 0, Max: 18446744073709551615
func (e *GstRTPST20221FecDec) SetSizeTime(value uint64) error {
	return e.Element.SetProperty("size-time", value)
}

func (e *GstRTPST20221FecDec) GetSizeTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("size-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

// mixes RTP DTMF streams into other RTP streams
type GstRTPDTMFMux struct {
	*GstRTPMux
}

func NewRTPDTMFMux() (*GstRTPDTMFMux, error) {
	e, err := gst.NewElement("rtpdtmfmux")
	if err != nil {
		return nil, err
	}
	return &GstRTPDTMFMux{GstRTPMux: &GstRTPMux{Element: e}}, nil
}

func NewRTPDTMFMuxWithName(name string) (*GstRTPDTMFMux, error) {
	e, err := gst.NewElementWithName("rtpdtmfmux", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPDTMFMux{GstRTPMux: &GstRTPMux{Element: e}}, nil
}

type GstRTCPSync string

const (
	GstRTCPSyncAlways  GstRTCPSync = "always"   // always
	GstRTCPSyncInitial GstRTCPSync = "initial"  // initial
	GstRTCPSyncRtpInfo GstRTCPSync = "rtp-info" // rtp-info
)

type GstRtpNtpTimeSource string

const (
	GstRtpNtpTimeSourceNtp         GstRtpNtpTimeSource = "ntp"          // NTP time based on realtime clock
	GstRtpNtpTimeSourceUnix        GstRtpNtpTimeSource = "unix"         // UNIX time based on realtime clock
	GstRtpNtpTimeSourceRunningTime GstRtpNtpTimeSource = "running-time" // Running time based on pipeline clock
	GstRtpNtpTimeSourceClockTime   GstRtpNtpTimeSource = "clock-time"   // Pipeline clock time
)

type RTPJitterBufferMode string

const (
	RTPJitterBufferModeNone   RTPJitterBufferMode = "none"   // Only use RTP timestamps
	RTPJitterBufferModeSlave  RTPJitterBufferMode = "slave"  // Slave receiver to sender clock
	RTPJitterBufferModeBuffer RTPJitterBufferMode = "buffer" // Do low/high watermark buffering
	RTPJitterBufferModeSynced RTPJitterBufferMode = "synced" // Synchronized sender and receiver clocks
)
