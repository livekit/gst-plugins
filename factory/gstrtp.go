// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Extracts G.726 audio from RTP packets
type GstRtpG726Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpG726Depay() (*GstRtpG726Depay, error) {
	e, err := gst.NewElement("rtpg726depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpG726Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpG726DepayWithName(name string) (*GstRtpG726Depay, error) {
	e, err := gst.NewElementWithName("rtpg726depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpG726Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Force AAL2 decoding for compatibility with bad payloaders
// Default: true
func (e *GstRtpG726Depay) SetForceAal2(value bool) error {
	return e.Element.SetProperty("force-aal2", value)
}

func (e *GstRtpG726Depay) GetForceAal2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aal2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Extracts MPEG4 video from RTP packets (RFC 3016)
type GstRtpMP4VDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpMP4VDepay() (*GstRtpMP4VDepay, error) {
	e, err := gst.NewElement("rtpmp4vdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpMP4VDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpMP4VDepayWithName(name string) (*GstRtpMP4VDepay, error) {
	e, err := gst.NewElementWithName("rtpmp4vdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpMP4VDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts MPEG audio from RTP packets (RFC 2038)
type GstRtpMPADepay struct {
	*GstRTPBaseDepayload
}

func NewRtpMPADepay() (*GstRtpMPADepay, error) {
	e, err := gst.NewElement("rtpmpadepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpMPADepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpMPADepayWithName(name string) (*GstRtpMPADepay, error) {
	e, err := gst.NewElementWithName("rtpmpadepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpMPADepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts QCELP (PureVoice) audio from RTP packets (RFC 2658)
type GstRtpQCELPDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpQCELPDepay() (*GstRtpQCELPDepay, error) {
	e, err := gst.NewElement("rtpqcelpdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpQCELPDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpQCELPDepayWithName(name string) (*GstRtpQCELPDepay, error) {
	e, err := gst.NewElementWithName("rtpqcelpdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpQCELPDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Encode Redundant Audio Data (RED)
type GstRtpRedEnc struct {
	*gst.Element
}

func NewRtpRedEnc() (*GstRtpRedEnc, error) {
	e, err := gst.NewElement("rtpredenc")
	if err != nil {
		return nil, err
	}
	return &GstRtpRedEnc{Element: e}, nil
}

func NewRtpRedEncWithName(name string) (*GstRtpRedEnc, error) {
	e, err := gst.NewElementWithName("rtpredenc", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpRedEnc{Element: e}, nil
}

// Payload type FEC packets (-1 disable)
// Default: 0, Min: 0, Max: 127
func (e *GstRtpRedEnc) SetPt(value int) error {
	return e.Element.SetProperty("pt", value)
}

func (e *GstRtpRedEnc) GetPt() (int, error) {
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

// Count of sent packets
// Default: 0, Min: 0, Max: -1
func (e *GstRtpRedEnc) GetSent() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("sent")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// true - can produce RED packets even without redundant blocks (distance==0) false - RED packets will be produced only if distance>0
// Default: true
func (e *GstRtpRedEnc) SetAllowNoRedBlocks(value bool) error {
	return e.Element.SetProperty("allow-no-red-blocks", value)
}

func (e *GstRtpRedEnc) GetAllowNoRedBlocks() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-no-red-blocks")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Tells which media packet to use as a redundant block (0 - no redundant blocks, 1 to use previous packet, 2 to use the packet before previous, etc.)
// Default: 0, Min: 0, Max: -1
func (e *GstRtpRedEnc) SetDistance(value uint) error {
	return e.Element.SetProperty("distance", value)
}

func (e *GstRtpRedEnc) GetDistance() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("distance")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Encodes RTP FEC (RFC5109)
type GstRtpUlpFecEnc struct {
	*gst.Element
}

func NewRtpUlpFecEnc() (*GstRtpUlpFecEnc, error) {
	e, err := gst.NewElement("rtpulpfecenc")
	if err != nil {
		return nil, err
	}
	return &GstRtpUlpFecEnc{Element: e}, nil
}

func NewRtpUlpFecEncWithName(name string) (*GstRtpUlpFecEnc, error) {
	e, err := gst.NewElementWithName("rtpulpfecenc", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpUlpFecEnc{Element: e}, nil
}

// FEC overhead percentage for important packets
// Default: 0, Min: 0, Max: 100
func (e *GstRtpUlpFecEnc) SetPercentageImportant(value uint) error {
	return e.Element.SetProperty("percentage-important", value)
}

func (e *GstRtpUlpFecEnc) GetPercentageImportant() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("percentage-important")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Count of protected packets
// Default: 0, Min: 0, Max: -1
func (e *GstRtpUlpFecEnc) GetProtected() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("protected")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The payload type of FEC packets
// Default: 255, Min: 0, Max: 255
func (e *GstRtpUlpFecEnc) SetPt(value uint) error {
	return e.Element.SetProperty("pt", value)
}

func (e *GstRtpUlpFecEnc) GetPt() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("pt")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Apply FEC on multiple packets
// Default: true
func (e *GstRtpUlpFecEnc) SetMultipacket(value bool) error {
	return e.Element.SetProperty("multipacket", value)
}

func (e *GstRtpUlpFecEnc) GetMultipacket() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("multipacket")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// FEC overhead percentage for the whole stream
// Default: 0, Min: 0, Max: 100
func (e *GstRtpUlpFecEnc) SetPercentage(value uint) error {
	return e.Element.SetProperty("percentage", value)
}

func (e *GstRtpUlpFecEnc) GetPercentage() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("percentage")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Extracts G722 audio from RTP packets
type GstRtpG722Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpG722Depay() (*GstRtpG722Depay, error) {
	e, err := gst.NewElement("rtpg722depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpG722Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpG722DepayWithName(name string) (*GstRtpG722Depay, error) {
	e, err := gst.NewElementWithName("rtpg722depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpG722Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload-encodes CELT audio into a RTP packet
type GstRtpCELTPay struct {
	*GstRTPBasePayload
}

func NewRtpCELTPay() (*GstRtpCELTPay, error) {
	e, err := gst.NewElement("rtpceltpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpCELTPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpCELTPayWithName(name string) (*GstRtpCELTPay, error) {
	e, err := gst.NewElementWithName("rtpceltpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpCELTPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Extracts G.723 audio from RTP packets (RFC 3551)
type GstRtpG723Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpG723Depay() (*GstRtpG723Depay, error) {
	e, err := gst.NewElement("rtpg723depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpG723Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpG723DepayWithName(name string) (*GstRtpG723Depay, error) {
	e, err := gst.NewElementWithName("rtpg723depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpG723Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts JPEG 2000 video from RTP packets (RFC 5371)
type GstRtpJ2KDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpJ2KDepay() (*GstRtpJ2KDepay, error) {
	e, err := gst.NewElement("rtpj2kdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpJ2KDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpJ2KDepayWithName(name string) (*GstRtpJ2KDepay, error) {
	e, err := gst.NewElementWithName("rtpj2kdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpJ2KDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts PCMU audio from RTP packets
type GstRtpPcmuDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpPcmuDepay() (*GstRtpPcmuDepay, error) {
	e, err := gst.NewElement("rtppcmudepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpPcmuDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpPcmuDepayWithName(name string) (*GstRtpPcmuDepay, error) {
	e, err := gst.NewElementWithName("rtppcmudepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpPcmuDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Decode Redundant Audio Data (RED)
type GstRtpRedDec struct {
	*gst.Element
}

func NewRtpRedDec() (*GstRtpRedDec, error) {
	e, err := gst.NewElement("rtpreddec")
	if err != nil {
		return nil, err
	}
	return &GstRtpRedDec{Element: e}, nil
}

func NewRtpRedDecWithName(name string) (*GstRtpRedDec, error) {
	e, err := gst.NewElementWithName("rtpreddec", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpRedDec{Element: e}, nil
}

// Payload type FEC packets
// Default: -1, Min: -1, Max: 127
func (e *GstRtpRedDec) SetPt(value int) error {
	return e.Element.SetProperty("pt", value)
}

func (e *GstRtpRedDec) GetPt() (int, error) {
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

// Count of received packets
// Default: 0, Min: 0, Max: -1
func (e *GstRtpRedDec) GetReceived() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("received")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// All the RED payloads this decoder may encounter

func (e *GstRtpRedDec) SetPayloads(value *gst.ValueArrayValue) error {
	return e.Element.SetProperty("payloads", value)
}

func (e *GstRtpRedDec) GetPayloads() (*gst.ValueArrayValue, error) {
	var value *gst.ValueArrayValue
	var ok bool
	v, err := e.Element.GetProperty("payloads")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.ValueArrayValue)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.ValueArrayValue")
	}
	return value, nil
}

// Extracts SVQ3 video from RTP packets (no RFC)
type GstRtpSV3VDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpSV3VDepay() (*GstRtpSV3VDepay, error) {
	e, err := gst.NewElement("rtpsv3vdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpSV3VDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpSV3VDepayWithName(name string) (*GstRtpSV3VDepay, error) {
	e, err := gst.NewElementWithName("rtpsv3vdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpSV3VDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload-encode Raw audio into RTP packets (RFC 3551)
type GstRtpL8Pay struct {
	*GstRTPBaseAudioPayload
}

func NewRtpL8Pay() (*GstRtpL8Pay, error) {
	e, err := gst.NewElement("rtpL8pay")
	if err != nil {
		return nil, err
	}
	return &GstRtpL8Pay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

func NewRtpL8PayWithName(name string) (*GstRtpL8Pay, error) {
	e, err := gst.NewElementWithName("rtpL8pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpL8Pay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

// Passes through RTP packets
type GstRtpPassthroughPay struct {
	*gst.Element
}

func NewRtpPassthroughPay() (*GstRtpPassthroughPay, error) {
	e, err := gst.NewElement("rtppassthroughpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpPassthroughPay{Element: e}, nil
}

func NewRtpPassthroughPayWithName(name string) (*GstRtpPassthroughPay, error) {
	e, err := gst.NewElementWithName("rtppassthroughpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpPassthroughPay{Element: e}, nil
}

// Maximum size of one packet
// Default: 0, Min: 28, Max: -1
func (e *GstRtpPassthroughPay) SetMtu(value uint) error {
	return e.Element.SetProperty("mtu", value)
}

func (e *GstRtpPassthroughPay) GetMtu() (uint, error) {
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

// The payload type of the packets
// Default: 128, Min: 0, Max: 128
func (e *GstRtpPassthroughPay) SetPt(value uint) error {
	return e.Element.SetProperty("pt", value)
}

func (e *GstRtpPassthroughPay) GetPt() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("pt")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The RTP sequence number of the last processed packet
// Default: 0, Min: 0, Max: 65535
func (e *GstRtpPassthroughPay) GetSeqnum() (uint, error) {
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
// Default: 0, Min: -1, Max: 65535
func (e *GstRtpPassthroughPay) SetSeqnumOffset(value int) error {
	return e.Element.SetProperty("seqnum-offset", value)
}

func (e *GstRtpPassthroughPay) GetSeqnumOffset() (int, error) {
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

// Various statistics
// Default: application/x-rtp-payload-stats, clock-rate=(uint)0, running-time=(guint64)18446744073709551615, seqnum=(uint)0, timestamp=(uint)0, ssrc=(uint)0, pt=(uint)128, seqnum-offset=(uint)0, timestamp-offset=(uint)0;
func (e *GstRtpPassthroughPay) GetStats() (*gst.Structure, error) {
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

// The RTP timestamp of the last processed packet
// Default: 0, Min: 0, Max: -1
func (e *GstRtpPassthroughPay) GetTimestamp() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("timestamp")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Offset to add to all outgoing timestamps (default = random)
// Default: 0, Min: 0, Max: -1
func (e *GstRtpPassthroughPay) SetTimestampOffset(value uint) error {
	return e.Element.SetProperty("timestamp-offset", value)
}

func (e *GstRtpPassthroughPay) GetTimestampOffset() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("timestamp-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Extracts PCMA audio from RTP packets
type GstRtpPcmaDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpPcmaDepay() (*GstRtpPcmaDepay, error) {
	e, err := gst.NewElement("rtppcmadepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpPcmaDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpPcmaDepayWithName(name string) (*GstRtpPcmaDepay, error) {
	e, err := gst.NewElementWithName("rtppcmadepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpPcmaDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload-encodes Speex audio into a RTP packet
type GstRtpSPEEXPay struct {
	*GstRTPBasePayload
}

func NewRtpSPEEXPay() (*GstRtpSPEEXPay, error) {
	e, err := gst.NewElement("rtpspeexpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpSPEEXPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpSPEEXPayWithName(name string) (*GstRtpSPEEXPay, error) {
	e, err := gst.NewElementWithName("rtpspeexpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpSPEEXPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Payload MPEG4 elementary streams as RTP packets (RFC 3640)
type GstRtpMP4GPay struct {
	*GstRTPBasePayload
}

func NewRtpMP4GPay() (*GstRtpMP4GPay, error) {
	e, err := gst.NewElement("rtpmp4gpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpMP4GPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpMP4GPayWithName(name string) (*GstRtpMP4GPay, error) {
	e, err := gst.NewElementWithName("rtpmp4gpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpMP4GPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Extracts AMR or AMR-WB audio from RTP packets (RFC 3267)
type GstRtpAMRDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpAMRDepay() (*GstRtpAMRDepay, error) {
	e, err := gst.NewElement("rtpamrdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpAMRDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpAMRDepayWithName(name string) (*GstRtpAMRDepay, error) {
	e, err := gst.NewElementWithName("rtpamrdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpAMRDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload-encode AMR or AMR-WB audio into RTP packets (RFC 3267)
type GstRtpAMRPay struct {
	*GstRTPBasePayload
}

func NewRtpAMRPay() (*GstRtpAMRPay, error) {
	e, err := gst.NewElement("rtpamrpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpAMRPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpAMRPayWithName(name string) (*GstRtpAMRPay, error) {
	e, err := gst.NewElementWithName("rtpamrpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpAMRPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Payload-encodes GSM audio into a RTP packet
type GstRTPGSMPay struct {
	*GstRTPBasePayload
}

func NewRTPGSMPay() (*GstRTPGSMPay, error) {
	e, err := gst.NewElement("rtpgsmpay")
	if err != nil {
		return nil, err
	}
	return &GstRTPGSMPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRTPGSMPayWithName(name string) (*GstRTPGSMPay, error) {
	e, err := gst.NewElementWithName("rtpgsmpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPGSMPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Payload-encodes H263 video in RTP packets (RFC 2190)
type GstRtpH263Pay struct {
	*GstRTPBasePayload
}

func NewRtpH263Pay() (*GstRtpH263Pay, error) {
	e, err := gst.NewElement("rtph263pay")
	if err != nil {
		return nil, err
	}
	return &GstRtpH263Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpH263PayWithName(name string) (*GstRtpH263Pay, error) {
	e, err := gst.NewElementWithName("rtph263pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpH263Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Disable packetization modes B and C
// Default: false
func (e *GstRtpH263Pay) SetModeaOnly(value bool) error {
	return e.Element.SetProperty("modea-only", value)
}

func (e *GstRtpH263Pay) GetModeaOnly() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("modea-only")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Extracts SBC audio from RTP packets
type GstRtpSbcDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpSbcDepay() (*GstRtpSbcDepay, error) {
	e, err := gst.NewElement("rtpsbcdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpSbcDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpSbcDepayWithName(name string) (*GstRtpSbcDepay, error) {
	e, err := gst.NewElementWithName("rtpsbcdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpSbcDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Various statistics
// Default: false
func (e *GstRtpSbcDepay) SetIgnoreTimestamps(value bool) error {
	return e.Element.SetProperty("ignore-timestamps", value)
}

func (e *GstRtpSbcDepay) GetIgnoreTimestamps() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ignore-timestamps")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Extracts Siren audio from RTP packets
type GstRTPSirenDepay struct {
	*GstRTPBaseDepayload
}

func NewRTPSirenDepay() (*GstRTPSirenDepay, error) {
	e, err := gst.NewElement("rtpsirendepay")
	if err != nil {
		return nil, err
	}
	return &GstRTPSirenDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRTPSirenDepayWithName(name string) (*GstRTPSirenDepay, error) {
	e, err := gst.NewElementWithName("rtpsirendepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPSirenDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts Speex audio from RTP packets
type GstRtpSPEEXDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpSPEEXDepay() (*GstRtpSPEEXDepay, error) {
	e, err := gst.NewElement("rtpspeexdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpSPEEXDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpSPEEXDepayWithName(name string) (*GstRtpSPEEXDepay, error) {
	e, err := gst.NewElementWithName("rtpspeexdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpSPEEXDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts raw audio from RTP packets
type GstRtpL16Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpL16Depay() (*GstRtpL16Depay, error) {
	e, err := gst.NewElement("rtpL16depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpL16Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpL16DepayWithName(name string) (*GstRtpL16Depay, error) {
	e, err := gst.NewElementWithName("rtpL16depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpL16Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts VP8 video from RTP packets)
type GstRtpVP8Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpVP8Depay() (*GstRtpVP8Depay, error) {
	e, err := gst.NewElement("rtpvp8depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpVP8Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpVP8DepayWithName(name string) (*GstRtpVP8Depay, error) {
	e, err := gst.NewElementWithName("rtpvp8depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpVP8Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Request new keyframe when packet loss is detected
// Default: false
func (e *GstRtpVP8Depay) SetRequestKeyframe(value bool) error {
	return e.Element.SetProperty("request-keyframe", value)
}

func (e *GstRtpVP8Depay) GetRequestKeyframe() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("request-keyframe")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Wait for the next keyframe after packet loss
// Default: false
func (e *GstRtpVP8Depay) SetWaitForKeyframe(value bool) error {
	return e.Element.SetProperty("wait-for-keyframe", value)
}

func (e *GstRtpVP8Depay) GetWaitForKeyframe() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("wait-for-keyframe")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Extracts GSM audio from RTP packets
type GstRTPGSMDepay struct {
	*GstRTPBaseDepayload
}

func NewRTPGSMDepay() (*GstRTPGSMDepay, error) {
	e, err := gst.NewElement("rtpgsmdepay")
	if err != nil {
		return nil, err
	}
	return &GstRTPGSMDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRTPGSMDepayWithName(name string) (*GstRTPGSMDepay, error) {
	e, err := gst.NewElementWithName("rtpgsmdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPGSMDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts MPEG1 System Streams from RTP packets (RFC 3555)
type GstRtpMP1SDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpMP1SDepay() (*GstRtpMP1SDepay, error) {
	e, err := gst.NewElement("rtpmp1sdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpMP1SDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpMP1SDepayWithName(name string) (*GstRtpMP1SDepay, error) {
	e, err := gst.NewElementWithName("rtpmp1sdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpMP1SDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload-encodes PCMA audio into a RTP packet
type GstRtpPcmaPay struct {
	*GstRTPBaseAudioPayload
}

func NewRtpPcmaPay() (*GstRtpPcmaPay, error) {
	e, err := gst.NewElement("rtppcmapay")
	if err != nil {
		return nil, err
	}
	return &GstRtpPcmaPay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

func NewRtpPcmaPayWithName(name string) (*GstRtpPcmaPay, error) {
	e, err := gst.NewElementWithName("rtppcmapay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpPcmaPay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

// Extracts Vorbis Audio from RTP packets (RFC 5215)
type GstRtpVorbisDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpVorbisDepay() (*GstRtpVorbisDepay, error) {
	e, err := gst.NewElement("rtpvorbisdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpVorbisDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpVorbisDepayWithName(name string) (*GstRtpVorbisDepay, error) {
	e, err := gst.NewElementWithName("rtpvorbisdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpVorbisDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts raw audio from RTP packets
type GstRtpL8Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpL8Depay() (*GstRtpL8Depay, error) {
	e, err := gst.NewElement("rtpL8depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpL8Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpL8DepayWithName(name string) (*GstRtpL8Depay, error) {
	e, err := gst.NewElementWithName("rtpL8depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpL8Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts H263/+/++ video from RTP packets (RFC 4629)
type GstRtpH263PDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpH263PDepay() (*GstRtpH263PDepay, error) {
	e, err := gst.NewElement("rtph263pdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpH263PDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpH263PDepayWithName(name string) (*GstRtpH263PDepay, error) {
	e, err := gst.NewElementWithName("rtph263pdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpH263PDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload-encodes H263/+/++ video in RTP packets (RFC 4629)
type GstRtpH263PPay struct {
	*GstRTPBasePayload
}

func NewRtpH263PPay() (*GstRtpH263PPay, error) {
	e, err := gst.NewElement("rtph263ppay")
	if err != nil {
		return nil, err
	}
	return &GstRtpH263PPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpH263PPayWithName(name string) (*GstRtpH263PPay, error) {
	e, err := gst.NewElementWithName("rtph263ppay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpH263PPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Packet Fragmentation Mode
// Default: normal (0)
func (e *GstRtpH263PPay) SetFragmentationMode(value GstFragmentationMode) error {
	e.Element.SetArg("fragmentation-mode", string(value))
	return nil
}

func (e *GstRtpH263PPay) GetFragmentationMode() (GstFragmentationMode, error) {
	var value GstFragmentationMode
	var ok bool
	v, err := e.Element.GetProperty("fragmentation-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFragmentationMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFragmentationMode")
	}
	return value, nil
}

// Extracts H265 video from RTP packets (RFC 7798)
type GstRtpH265Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpH265Depay() (*GstRtpH265Depay, error) {
	e, err := gst.NewElement("rtph265depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpH265Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpH265DepayWithName(name string) (*GstRtpH265Depay, error) {
	e, err := gst.NewElementWithName("rtph265depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpH265Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload-encodes PCMU audio into a RTP packet
type GstRtpPcmuPay struct {
	*GstRTPBaseAudioPayload
}

func NewRtpPcmuPay() (*GstRtpPcmuPay, error) {
	e, err := gst.NewElement("rtppcmupay")
	if err != nil {
		return nil, err
	}
	return &GstRtpPcmuPay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

func NewRtpPcmuPayWithName(name string) (*GstRtpPcmuPay, error) {
	e, err := gst.NewElementWithName("rtppcmupay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpPcmuPay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

// Payload SBC audio as RTP packets
type GstRtpSBCPay struct {
	*GstRTPBasePayload
}

func NewRtpSBCPay() (*GstRtpSBCPay, error) {
	e, err := gst.NewElement("rtpsbcpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpSBCPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpSBCPayWithName(name string) (*GstRtpSBCPay, error) {
	e, err := gst.NewElementWithName("rtpsbcpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpSBCPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Minimum quantity of frames to send in one packet (-1 for maximum allowed by the mtu)
// Default: 0, Min: -1, Max: 2147483647
func (e *GstRtpSBCPay) SetMinFrames(value int) error {
	return e.Element.SetProperty("min-frames", value)
}

func (e *GstRtpSBCPay) GetMinFrames() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-frames")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Extracts VP9 video from RTP packets)
type GstRtpVP9Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpVP9Depay() (*GstRtpVP9Depay, error) {
	e, err := gst.NewElement("rtpvp9depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpVP9Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpVP9DepayWithName(name string) (*GstRtpVP9Depay, error) {
	e, err := gst.NewElementWithName("rtpvp9depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpVP9Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Request new keyframe when packet loss is detected
// Default: false
func (e *GstRtpVP9Depay) SetRequestKeyframe(value bool) error {
	return e.Element.SetProperty("request-keyframe", value)
}

func (e *GstRtpVP9Depay) GetRequestKeyframe() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("request-keyframe")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Wait for the next keyframe after packet loss
// Default: false
func (e *GstRtpVP9Depay) SetWaitForKeyframe(value bool) error {
	return e.Element.SetProperty("wait-for-keyframe", value)
}

func (e *GstRtpVP9Depay) GetWaitForKeyframe() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("wait-for-keyframe")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Extracts H261 video from RTP packets (RFC 4587)
type GstRtpH261Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpH261Depay() (*GstRtpH261Depay, error) {
	e, err := gst.NewElement("rtph261depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpH261Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpH261DepayWithName(name string) (*GstRtpH261Depay, error) {
	e, err := gst.NewElementWithName("rtph261depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpH261Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload-encodes H261 video in RTP packets (RFC 4587)
type GstRtpH261Pay struct {
	*GstRTPBasePayload
}

func NewRtpH261Pay() (*GstRtpH261Pay, error) {
	e, err := gst.NewElement("rtph261pay")
	if err != nil {
		return nil, err
	}
	return &GstRtpH261Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpH261PayWithName(name string) (*GstRtpH261Pay, error) {
	e, err := gst.NewElementWithName("rtph261pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpH261Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Extracts MPEG4 elementary streams from RTP packets (RFC 3640)
type GstRtpMP4GDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpMP4GDepay() (*GstRtpMP4GDepay, error) {
	e, err := gst.NewElement("rtpmp4gdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpMP4GDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpMP4GDepayWithName(name string) (*GstRtpMP4GDepay, error) {
	e, err := gst.NewElementWithName("rtpmp4gdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpMP4GDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts MPEG video from RTP packets (RFC 2250)
type GstRtpMPVDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpMPVDepay() (*GstRtpMPVDepay, error) {
	e, err := gst.NewElement("rtpmpvdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpMPVDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpMPVDepayWithName(name string) (*GstRtpMPVDepay, error) {
	e, err := gst.NewElementWithName("rtpmpvdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpMPVDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Puts VP8 video in RTP packets
type GstRtpVP8Pay struct {
	*GstRTPBasePayload
}

func NewRtpVP8Pay() (*GstRtpVP8Pay, error) {
	e, err := gst.NewElement("rtpvp8pay")
	if err != nil {
		return nil, err
	}
	return &GstRtpVP8Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpVP8PayWithName(name string) (*GstRtpVP8Pay, error) {
	e, err := gst.NewElementWithName("rtpvp8pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpVP8Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Currently used picture-id for payloading
// Default: 0, Min: 0, Max: 32767
func (e *GstRtpVP8Pay) GetPictureId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("picture-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The picture ID mode for payloading
// Default: none (0)
func (e *GstRtpVP8Pay) SetPictureIdMode(value GstVP8RTPPayMode) error {
	e.Element.SetArg("picture-id-mode", string(value))
	return nil
}

func (e *GstRtpVP8Pay) GetPictureIdMode() (GstVP8RTPPayMode, error) {
	var value GstVP8RTPPayMode
	var ok bool
	v, err := e.Element.GetProperty("picture-id-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVP8RTPPayMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVP8RTPPayMode")
	}
	return value, nil
}

// Offset to add to the initial picture-id (-1 = random)
// Default: -1, Min: -1, Max: 32767
func (e *GstRtpVP8Pay) SetPictureIdOffset(value int) error {
	return e.Element.SetProperty("picture-id-offset", value)
}

func (e *GstRtpVP8Pay) GetPictureIdOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("picture-id-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Extracts raw video from RTP packets (RFC 4175)
type GstRtpVRawDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpVRawDepay() (*GstRtpVRawDepay, error) {
	e, err := gst.NewElement("rtpvrawdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpVRawDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpVRawDepayWithName(name string) (*GstRtpVRawDepay, error) {
	e, err := gst.NewElementWithName("rtpvrawdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpVRawDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload raw video as RTP packets (RFC 4175)
type GstRtpVRawPay struct {
	*GstRTPBasePayload
}

func NewRtpVRawPay() (*GstRtpVRawPay, error) {
	e, err := gst.NewElement("rtpvrawpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpVRawPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpVRawPayWithName(name string) (*GstRtpVRawPay, error) {
	e, err := gst.NewElementWithName("rtpvrawpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpVRawPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Split and send out each frame in multiple chunks to reduce overhead
// Default: 10, Min: 1, Max: 2147483647
func (e *GstRtpVRawPay) SetChunksPerFrame(value int) error {
	return e.Element.SetProperty("chunks-per-frame", value)
}

func (e *GstRtpVRawPay) GetChunksPerFrame() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("chunks-per-frame")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Extracts GStreamer buffers from RTP packets
type GstRtpGSTDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpGSTDepay() (*GstRtpGSTDepay, error) {
	e, err := gst.NewElement("rtpgstdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpGSTDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpGSTDepayWithName(name string) (*GstRtpGSTDepay, error) {
	e, err := gst.NewElementWithName("rtpgstdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpGSTDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Depayloads DV from RTP packets (RFC 3189)
type GstRTPDVDepay struct {
	*GstRTPBaseDepayload
}

func NewRTPDVDepay() (*GstRTPDVDepay, error) {
	e, err := gst.NewElement("rtpdvdepay")
	if err != nil {
		return nil, err
	}
	return &GstRTPDVDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRTPDVDepayWithName(name string) (*GstRTPDVDepay, error) {
	e, err := gst.NewElementWithName("rtpdvdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPDVDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts G.729 audio from RTP packets (RFC 3551)
type GstRtpG729Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpG729Depay() (*GstRtpG729Depay, error) {
	e, err := gst.NewElement("rtpg729depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpG729Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpG729DepayWithName(name string) (*GstRtpG729Depay, error) {
	e, err := gst.NewElementWithName("rtpg729depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpG729Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload GStreamer buffers as RTP packets
type GstRtpGSTPay struct {
	*GstRTPBasePayload
}

func NewRtpGSTPay() (*GstRtpGSTPay, error) {
	e, err := gst.NewElement("rtpgstpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpGSTPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpGSTPayWithName(name string) (*GstRtpGSTPay, error) {
	e, err := gst.NewElementWithName("rtpgstpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpGSTPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Interval for sending caps and TAG events in seconds (0 = disabled)
// Default: 0, Min: 0, Max: 3600
func (e *GstRtpGSTPay) SetConfigInterval(value uint) error {
	return e.Element.SetProperty("config-interval", value)
}

func (e *GstRtpGSTPay) GetConfigInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Extracts MPEG2 TS from RTP packets (RFC 2250)
type GstRtpMP2TDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpMP2TDepay() (*GstRtpMP2TDepay, error) {
	e, err := gst.NewElement("rtpmp2tdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpMP2TDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpMP2TDepayWithName(name string) (*GstRtpMP2TDepay, error) {
	e, err := gst.NewElementWithName("rtpmp2tdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpMP2TDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// The amount of bytes that need to be skipped at the beginning of the payload
// Default: 0, Min: 0, Max: -1
func (e *GstRtpMP2TDepay) SetSkipFirstBytes(value uint) error {
	return e.Element.SetProperty("skip-first-bytes", value)
}

func (e *GstRtpMP2TDepay) GetSkipFirstBytes() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("skip-first-bytes")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Payload-encodes MPEG2 TS into RTP packets (RFC 2250)
type GstRTPMP2TPay struct {
	*GstRTPBasePayload
}

func NewRTPMP2TPay() (*GstRTPMP2TPay, error) {
	e, err := gst.NewElement("rtpmp2tpay")
	if err != nil {
		return nil, err
	}
	return &GstRTPMP2TPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRTPMP2TPayWithName(name string) (*GstRTPMP2TPay, error) {
	e, err := gst.NewElementWithName("rtpmp2tpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPMP2TPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Extracts MPEG4 audio from RTP packets (RFC 3016)
type GstRtpMP4ADepay struct {
	*GstRTPBaseDepayload
}

func NewRtpMP4ADepay() (*GstRtpMP4ADepay, error) {
	e, err := gst.NewElement("rtpmp4adepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpMP4ADepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpMP4ADepayWithName(name string) (*GstRtpMP4ADepay, error) {
	e, err := gst.NewElementWithName("rtpmp4adepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpMP4ADepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload MPEG4 audio as RTP packets (RFC 3016)
type GstRtpMP4APay struct {
	*GstRTPBasePayload
}

func NewRtpMP4APay() (*GstRtpMP4APay, error) {
	e, err := gst.NewElement("rtpmp4apay")
	if err != nil {
		return nil, err
	}
	return &GstRtpMP4APay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpMP4APayWithName(name string) (*GstRtpMP4APay, error) {
	e, err := gst.NewElementWithName("rtpmp4apay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpMP4APay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Payload-encode Raw audio into RTP packets (RFC 3551)
type GstRtpL16Pay struct {
	*GstRTPBaseAudioPayload
}

func NewRtpL16Pay() (*GstRtpL16Pay, error) {
	e, err := gst.NewElement("rtpL16pay")
	if err != nil {
		return nil, err
	}
	return &GstRtpL16Pay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

func NewRtpL16PayWithName(name string) (*GstRtpL16Pay, error) {
	e, err := gst.NewElementWithName("rtpL16pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpL16Pay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

// Decodes RTP FEC (RFC5109)
type GstRtpUlpFecDec struct {
	*gst.Element
}

func NewRtpUlpFecDec() (*GstRtpUlpFecDec, error) {
	e, err := gst.NewElement("rtpulpfecdec")
	if err != nil {
		return nil, err
	}
	return &GstRtpUlpFecDec{Element: e}, nil
}

func NewRtpUlpFecDecWithName(name string) (*GstRtpUlpFecDec, error) {
	e, err := gst.NewElementWithName("rtpulpfecdec", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpUlpFecDec{Element: e}, nil
}

// The number of unrecovered packets
// Default: 0, Min: 0, Max: -1
func (e *GstRtpUlpFecDec) GetUnrecovered() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("unrecovered")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Whether to passthrough all data as-is without modification and never attempt to recover packets
// Default: false
func (e *GstRtpUlpFecDec) SetPassthrough(value bool) error {
	return e.Element.SetProperty("passthrough", value)
}

func (e *GstRtpUlpFecDec) GetPassthrough() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("passthrough")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// FEC packets payload type
// Default: 0, Min: 0, Max: 127
func (e *GstRtpUlpFecDec) SetPt(value uint) error {
	return e.Element.SetProperty("pt", value)
}

func (e *GstRtpUlpFecDec) GetPt() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("pt")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The number of recovered packets
// Default: 0, Min: 0, Max: -1
func (e *GstRtpUlpFecDec) GetRecovered() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("recovered")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// RTP storage

func (e *GstRtpUlpFecDec) SetStorage(value interface{}) error {
	return e.Element.SetProperty("storage", value)
}

func (e *GstRtpUlpFecDec) GetStorage() (interface{}, error) {
	return e.Element.GetProperty("storage")
}

// Payload-encodes JPEG 2000 pictures into RTP packets (RFC 5371)
type GstRtpJ2KPay struct {
	*GstRTPBasePayload
}

func NewRtpJ2KPay() (*GstRtpJ2KPay, error) {
	e, err := gst.NewElement("rtpj2kpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpJ2KPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpJ2KPayWithName(name string) (*GstRtpJ2KPay, error) {
	e, err := gst.NewElementWithName("rtpj2kpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpJ2KPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Extracts JPEG video from RTP packets (RFC 2435)
type GstRtpJPEGDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpJPEGDepay() (*GstRtpJPEGDepay, error) {
	e, err := gst.NewElement("rtpjpegdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpJPEGDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpJPEGDepayWithName(name string) (*GstRtpJPEGDepay, error) {
	e, err := gst.NewElementWithName("rtpjpegdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpJPEGDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload MPEG-4 video as RTP packets (RFC 3016)
type GstRtpMP4VPay struct {
	*GstRTPBasePayload
}

func NewRtpMP4VPay() (*GstRtpMP4VPay, error) {
	e, err := gst.NewElement("rtpmp4vpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpMP4VPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpMP4VPayWithName(name string) (*GstRtpMP4VPay, error) {
	e, err := gst.NewElementWithName("rtpmp4vpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpMP4VPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Send Config Insertion Interval in seconds (configuration headers will be multiplexed in the data stream when detected.) (0 = disabled, -1 = send with every IDR frame)
// Default: 0, Min: -1, Max: 3600
func (e *GstRtpMP4VPay) SetConfigInterval(value int) error {
	return e.Element.SetProperty("config-interval", value)
}

func (e *GstRtpMP4VPay) GetConfigInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Puts Opus audio in RTP packets
type GstRtpOPUSPay struct {
	*GstRTPBasePayload
}

func NewRtpOPUSPay() (*GstRtpOPUSPay, error) {
	e, err := gst.NewElement("rtpopuspay")
	if err != nil {
		return nil, err
	}
	return &GstRtpOPUSPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpOPUSPayWithName(name string) (*GstRtpOPUSPay, error) {
	e, err := gst.NewElementWithName("rtpopuspay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpOPUSPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// If enabled, the payloader will not transmit empty packets
// Default: false
func (e *GstRtpOPUSPay) SetDtx(value bool) error {
	return e.Element.SetProperty("dtx", value)
}

func (e *GstRtpOPUSPay) GetDtx() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("dtx")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Payload-encode Theora video into RTP packets (draft-01 RFC XXXX)
type GstRtpTheoraPay struct {
	*GstRTPBasePayload
}

func NewRtpTheoraPay() (*GstRtpTheoraPay, error) {
	e, err := gst.NewElement("rtptheorapay")
	if err != nil {
		return nil, err
	}
	return &GstRtpTheoraPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpTheoraPayWithName(name string) (*GstRtpTheoraPay, error) {
	e, err := gst.NewElementWithName("rtptheorapay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpTheoraPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Send Config Insertion Interval in seconds (configuration headers will be multiplexed in the data stream when detected.) (0 = disabled)
// Default: 0, Min: 0, Max: 3600
func (e *GstRtpTheoraPay) SetConfigInterval(value uint) error {
	return e.Element.SetProperty("config-interval", value)
}

func (e *GstRtpTheoraPay) GetConfigInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Extracts BroadcomVoice audio from RTP packets (RFC 4298)
type GstRTPBVDepay struct {
	*GstRTPBaseDepayload
}

func NewRTPBVDepay() (*GstRTPBVDepay, error) {
	e, err := gst.NewElement("rtpbvdepay")
	if err != nil {
		return nil, err
	}
	return &GstRTPBVDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRTPBVDepayWithName(name string) (*GstRTPBVDepay, error) {
	e, err := gst.NewElementWithName("rtpbvdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPBVDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Packetize G.723 audio into RTP packets
type GstRTPG723Pay struct {
	*GstRTPBasePayload
}

func NewRTPG723Pay() (*GstRTPG723Pay, error) {
	e, err := gst.NewElement("rtpg723pay")
	if err != nil {
		return nil, err
	}
	return &GstRTPG723Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRTPG723PayWithName(name string) (*GstRTPG723Pay, error) {
	e, err := gst.NewElementWithName("rtpg723pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPG723Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Payload-encode H265 video into RTP packets (RFC 7798)
type GstRtpH265Pay struct {
	*GstRTPBasePayload
}

func NewRtpH265Pay() (*GstRtpH265Pay, error) {
	e, err := gst.NewElement("rtph265pay")
	if err != nil {
		return nil, err
	}
	return &GstRtpH265Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpH265PayWithName(name string) (*GstRtpH265Pay, error) {
	e, err := gst.NewElementWithName("rtph265pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpH265Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Send VPS, SPS and PPS Insertion Interval in seconds (sprop parameter sets will be multiplexed in the data stream when detected.) (0 = disabled, -1 = send with every IDR frame)
// Default: 0, Min: -1, Max: 3600
func (e *GstRtpH265Pay) SetConfigInterval(value int) error {
	return e.Element.SetProperty("config-interval", value)
}

func (e *GstRtpH265Pay) GetConfigInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Bundle suitable SPS/PPS NAL units into aggregate packets.
// Default: none (0)
func (e *GstRtpH265Pay) SetAggregateMode(value GstRtpH265AggregateMode) error {
	e.Element.SetArg("aggregate-mode", string(value))
	return nil
}

func (e *GstRtpH265Pay) GetAggregateMode() (GstRtpH265AggregateMode, error) {
	var value GstRtpH265AggregateMode
	var ok bool
	v, err := e.Element.GetProperty("aggregate-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRtpH265AggregateMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRtpH265AggregateMode")
	}
	return value, nil
}

// Payload MPEG audio as RTP packets (RFC 2038)
type GstRtpMPAPay struct {
	*GstRTPBasePayload
}

func NewRtpMPAPay() (*GstRtpMPAPay, error) {
	e, err := gst.NewElement("rtpmpapay")
	if err != nil {
		return nil, err
	}
	return &GstRtpMPAPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpMPAPayWithName(name string) (*GstRtpMPAPay, error) {
	e, err := gst.NewElementWithName("rtpmpapay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpMPAPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Payload-encode Raw audio into RTP packets (RFC 3551)
type GstRtpG722Pay struct {
	*GstRTPBaseAudioPayload
}

func NewRtpG722Pay() (*GstRtpG722Pay, error) {
	e, err := gst.NewElement("rtpg722pay")
	if err != nil {
		return nil, err
	}
	return &GstRtpG722Pay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

func NewRtpG722PayWithName(name string) (*GstRtpG722Pay, error) {
	e, err := gst.NewElementWithName("rtpg722pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpG722Pay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

// Payloads KLV (SMPTE ST 336) metadata as RTP packets
type GstRtpKlvPay struct {
	*GstRTPBasePayload
}

func NewRtpKlvPay() (*GstRtpKlvPay, error) {
	e, err := gst.NewElement("rtpklvpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpKlvPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpKlvPayWithName(name string) (*GstRtpKlvPay, error) {
	e, err := gst.NewElementWithName("rtpklvpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpKlvPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Extracts CELT audio from RTP packets
type GstRtpCELTDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpCELTDepay() (*GstRtpCELTDepay, error) {
	e, err := gst.NewElement("rtpceltdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpCELTDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpCELTDepayWithName(name string) (*GstRtpCELTDepay, error) {
	e, err := gst.NewElementWithName("rtpceltdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpCELTDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payloads DV into RTP packets (RFC 3189)
type GstRTPDVPay struct {
	*GstRTPBasePayload
}

func NewRTPDVPay() (*GstRTPDVPay, error) {
	e, err := gst.NewElement("rtpdvpay")
	if err != nil {
		return nil, err
	}
	return &GstRTPDVPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRTPDVPayWithName(name string) (*GstRTPDVPay, error) {
	e, err := gst.NewElementWithName("rtpdvpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPDVPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// The payload mode of payloading
// Default: video (0)
func (e *GstRTPDVPay) SetMode(value GstDVPayMode) error {
	e.Element.SetArg("mode", string(value))
	return nil
}

func (e *GstRTPDVPay) GetMode() (GstDVPayMode, error) {
	var value GstDVPayMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDVPayMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDVPayMode")
	}
	return value, nil
}

// Extends RTP packets with color space and high dynamic range (HDR) information.
type GstRTPHeaderExtensionColorspace struct {
	*GstRTPHeaderExtension
}

func NewRTPHeaderExtensionColorspace() (*GstRTPHeaderExtensionColorspace, error) {
	e, err := gst.NewElement("rtphdrextcolorspace")
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionColorspace{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

func NewRTPHeaderExtensionColorspaceWithName(name string) (*GstRTPHeaderExtensionColorspace, error) {
	e, err := gst.NewElementWithName("rtphdrextcolorspace", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPHeaderExtensionColorspace{GstRTPHeaderExtension: &GstRTPHeaderExtension{Element: e}}, nil
}

// Extracts iSAC audio from RTP packets
type GstRtpIsacDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpIsacDepay() (*GstRtpIsacDepay, error) {
	e, err := gst.NewElement("rtpisacdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpIsacDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpIsacDepayWithName(name string) (*GstRtpIsacDepay, error) {
	e, err := gst.NewElementWithName("rtpisacdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpIsacDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts KLV (SMPTE ST 336) metadata from RTP packets
type GstRtpKlvDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpKlvDepay() (*GstRtpKlvDepay, error) {
	e, err := gst.NewElement("rtpklvdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpKlvDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpKlvDepayWithName(name string) (*GstRtpKlvDepay, error) {
	e, err := gst.NewElementWithName("rtpklvdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpKlvDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Packetize Siren audio streams into RTP packets
type GstRTPSirenPay struct {
	*GstRTPBaseAudioPayload
}

func NewRTPSirenPay() (*GstRTPSirenPay, error) {
	e, err := gst.NewElement("rtpsirenpay")
	if err != nil {
		return nil, err
	}
	return &GstRTPSirenPay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

func NewRTPSirenPayWithName(name string) (*GstRTPSirenPay, error) {
	e, err := gst.NewElementWithName("rtpsirenpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPSirenPay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

// Helper element for various purposes (ex. recovering from packet loss using RED/FEC). Saves given number of RTP packets. Should be instantiated before jitterbuffer
type GstRtpStorage struct {
	*gst.Element
}

func NewRtpStorage() (*GstRtpStorage, error) {
	e, err := gst.NewElement("rtpstorage")
	if err != nil {
		return nil, err
	}
	return &GstRtpStorage{Element: e}, nil
}

func NewRtpStorageWithName(name string) (*GstRtpStorage, error) {
	e, err := gst.NewElementWithName("rtpstorage", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpStorage{Element: e}, nil
}

// The amount of data to keep in the storage (in ns, 0-disable)
// Default: 0, Min: 0, Max: 18446744073709551615
func (e *GstRtpStorage) SetSizeTime(value uint64) error {
	return e.Element.SetProperty("size-time", value)
}

func (e *GstRtpStorage) GetSizeTime() (uint64, error) {
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

// Internal RtpStorage object

func (e *GstRtpStorage) GetInternalStorage() (interface{}, error) {
	return e.Element.GetProperty("internal-storage")
}

// Payload-encode Raw 24-bit audio into RTP packets (RFC 3190)
type GstRtpL24Pay struct {
	*GstRTPBaseAudioPayload
}

func NewRtpL24Pay() (*GstRtpL24Pay, error) {
	e, err := gst.NewElement("rtpL24pay")
	if err != nil {
		return nil, err
	}
	return &GstRtpL24Pay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

func NewRtpL24PayWithName(name string) (*GstRtpL24Pay, error) {
	e, err := gst.NewElementWithName("rtpL24pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpL24Pay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

// Payload-encode Vorbis audio into RTP packets (RFC 5215)
type GstRtpVorbisPay struct {
	*GstRTPBasePayload
}

func NewRtpVorbisPay() (*GstRtpVorbisPay, error) {
	e, err := gst.NewElement("rtpvorbispay")
	if err != nil {
		return nil, err
	}
	return &GstRtpVorbisPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpVorbisPayWithName(name string) (*GstRtpVorbisPay, error) {
	e, err := gst.NewElementWithName("rtpvorbispay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpVorbisPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Send Config Insertion Interval in seconds (configuration headers will be multiplexed in the data stream when detected.) (0 = disabled)
// Default: 0, Min: 0, Max: 3600
func (e *GstRtpVorbisPay) SetConfigInterval(value uint) error {
	return e.Element.SetProperty("config-interval", value)
}

func (e *GstRtpVorbisPay) GetConfigInterval() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Extracts Theora video from RTP packets (draft-01 of RFC XXXX)
type GstRtpTheoraDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpTheoraDepay() (*GstRtpTheoraDepay, error) {
	e, err := gst.NewElement("rtptheoradepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpTheoraDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpTheoraDepayWithName(name string) (*GstRtpTheoraDepay, error) {
	e, err := gst.NewElementWithName("rtptheoradepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpTheoraDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Puts VP9 video in RTP packets)
type GstRtpVP9Pay struct {
	*GstRTPBasePayload
}

func NewRtpVP9Pay() (*GstRtpVP9Pay, error) {
	e, err := gst.NewElement("rtpvp9pay")
	if err != nil {
		return nil, err
	}
	return &GstRtpVP9Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpVP9PayWithName(name string) (*GstRtpVP9Pay, error) {
	e, err := gst.NewElementWithName("rtpvp9pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpVP9Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Currently used picture-id for payloading
// Default: 0, Min: 0, Max: 32767
func (e *GstRtpVP9Pay) GetPictureId() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("picture-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The picture ID mode for payloading
// Default: none (0)
func (e *GstRtpVP9Pay) SetPictureIdMode(value GstVP9RTPPayMode) error {
	e.Element.SetArg("picture-id-mode", string(value))
	return nil
}

func (e *GstRtpVP9Pay) GetPictureIdMode() (GstVP9RTPPayMode, error) {
	var value GstVP9RTPPayMode
	var ok bool
	v, err := e.Element.GetProperty("picture-id-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstVP9RTPPayMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstVP9RTPPayMode")
	}
	return value, nil
}

// Offset to add to the initial picture-id (-1 = random)
// Default: -1, Min: -1, Max: 32767
func (e *GstRtpVP9Pay) SetPictureIdOffset(value int) error {
	return e.Element.SetProperty("picture-id-offset", value)
}

func (e *GstRtpVP9Pay) GetPictureIdOffset() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("picture-id-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Payload-encodes G.726 audio into a RTP packet
type GstRtpG726Pay struct {
	*GstRTPBaseAudioPayload
}

func NewRtpG726Pay() (*GstRtpG726Pay, error) {
	e, err := gst.NewElement("rtpg726pay")
	if err != nil {
		return nil, err
	}
	return &GstRtpG726Pay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

func NewRtpG726PayWithName(name string) (*GstRtpG726Pay, error) {
	e, err := gst.NewElementWithName("rtpg726pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpG726Pay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

// Force AAL2 encoding for compatibility with bad depayloaders
// Default: true
func (e *GstRtpG726Pay) SetForceAal2(value bool) error {
	return e.Element.SetProperty("force-aal2", value)
}

func (e *GstRtpG726Pay) GetForceAal2() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aal2")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Payload AC3 audio as RTP packets (RFC 4184)
type GstRtpAC3Pay struct {
	*GstRTPBasePayload
}

func NewRtpAC3Pay() (*GstRtpAC3Pay, error) {
	e, err := gst.NewElement("rtpac3pay")
	if err != nil {
		return nil, err
	}
	return &GstRtpAC3Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpAC3PayWithName(name string) (*GstRtpAC3Pay, error) {
	e, err := gst.NewElementWithName("rtpac3pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpAC3Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Packetize BroadcomVoice audio streams into RTP packets (RFC 4298)
type GstRTPBVPay struct {
	*GstRTPBaseAudioPayload
}

func NewRTPBVPay() (*GstRTPBVPay, error) {
	e, err := gst.NewElement("rtpbvpay")
	if err != nil {
		return nil, err
	}
	return &GstRTPBVPay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

func NewRTPBVPayWithName(name string) (*GstRTPBVPay, error) {
	e, err := gst.NewElementWithName("rtpbvpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPBVPay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

// Extracts H263 video from RTP packets (RFC 2190)
type GstRtpH263Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpH263Depay() (*GstRtpH263Depay, error) {
	e, err := gst.NewElement("rtph263depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpH263Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpH263DepayWithName(name string) (*GstRtpH263Depay, error) {
	e, err := gst.NewElementWithName("rtph263depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpH263Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts H264 video from RTP packets (RFC 3984)
type GstRtpH264Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpH264Depay() (*GstRtpH264Depay, error) {
	e, err := gst.NewElement("rtph264depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpH264Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpH264DepayWithName(name string) (*GstRtpH264Depay, error) {
	e, err := gst.NewElementWithName("rtph264depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpH264Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Request new keyframe when packet loss is detected
// Default: false
func (e *GstRtpH264Depay) SetRequestKeyframe(value bool) error {
	return e.Element.SetProperty("request-keyframe", value)
}

func (e *GstRtpH264Depay) GetRequestKeyframe() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("request-keyframe")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Wait for the next keyframe after packet loss, meaningful only when outputting access units
// Default: false
func (e *GstRtpH264Depay) SetWaitForKeyframe(value bool) error {
	return e.Element.SetProperty("wait-for-keyframe", value)
}

func (e *GstRtpH264Depay) GetWaitForKeyframe() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("wait-for-keyframe")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Packetize iLBC audio streams into RTP packets
type GstRTPILBCPay struct {
	*GstRTPBaseAudioPayload
}

func NewRTPILBCPay() (*GstRTPILBCPay, error) {
	e, err := gst.NewElement("rtpilbcpay")
	if err != nil {
		return nil, err
	}
	return &GstRTPILBCPay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

func NewRTPILBCPayWithName(name string) (*GstRTPILBCPay, error) {
	e, err := gst.NewElementWithName("rtpilbcpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPILBCPay{GstRTPBaseAudioPayload: &GstRTPBaseAudioPayload{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}}, nil
}

// Payload-encodes iSAC audio into a RTP packet
type GstRtpIsacPay struct {
	*GstRTPBasePayload
}

func NewRtpIsacPay() (*GstRtpIsacPay, error) {
	e, err := gst.NewElement("rtpisacpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpIsacPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpIsacPayWithName(name string) (*GstRtpIsacPay, error) {
	e, err := gst.NewElementWithName("rtpisacpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpIsacPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Payload LDAC audio as RTP packets
type GstRtpLdacPay struct {
	*GstRTPBasePayload
}

func NewRtpLdacPay() (*GstRtpLdacPay, error) {
	e, err := gst.NewElement("rtpldacpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpLdacPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpLdacPayWithName(name string) (*GstRtpLdacPay, error) {
	e, err := gst.NewElementWithName("rtpldacpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpLdacPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Extracts raw 24-bit audio from RTP packets
type GstRtpL24Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpL24Depay() (*GstRtpL24Depay, error) {
	e, err := gst.NewElement("rtpL24depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpL24Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpL24DepayWithName(name string) (*GstRtpL24Depay, error) {
	e, err := gst.NewElementWithName("rtpL24depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpL24Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts MPEG audio from RTP packets (RFC 5219)
type GstRtpMPARobustDepay struct {
	*GstRTPBaseDepayload
}

func NewRtpMPARobustDepay() (*GstRtpMPARobustDepay, error) {
	e, err := gst.NewElement("rtpmparobustdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpMPARobustDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpMPARobustDepayWithName(name string) (*GstRtpMPARobustDepay, error) {
	e, err := gst.NewElementWithName("rtpmparobustdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpMPARobustDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Extracts AC3 audio from RTP packets (RFC 4184)
type GstRtpAC3Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpAC3Depay() (*GstRtpAC3Depay, error) {
	e, err := gst.NewElement("rtpac3depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpAC3Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpAC3DepayWithName(name string) (*GstRtpAC3Depay, error) {
	e, err := gst.NewElementWithName("rtpac3depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpAC3Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payload-encode H264 video into RTP packets (RFC 3984)
type GstRtpH264Pay struct {
	*GstRTPBasePayload
}

func NewRtpH264Pay() (*GstRtpH264Pay, error) {
	e, err := gst.NewElement("rtph264pay")
	if err != nil {
		return nil, err
	}
	return &GstRtpH264Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpH264PayWithName(name string) (*GstRtpH264Pay, error) {
	e, err := gst.NewElementWithName("rtph264pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpH264Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Bundle suitable SPS/PPS NAL units into STAP-A aggregate packets
// Default: none (0)
func (e *GstRtpH264Pay) SetAggregateMode(value GstRtpH264AggregateMode) error {
	e.Element.SetArg("aggregate-mode", string(value))
	return nil
}

func (e *GstRtpH264Pay) GetAggregateMode() (GstRtpH264AggregateMode, error) {
	var value GstRtpH264AggregateMode
	var ok bool
	v, err := e.Element.GetProperty("aggregate-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstRtpH264AggregateMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstRtpH264AggregateMode")
	}
	return value, nil
}

// Send SPS and PPS Insertion Interval in seconds (sprop parameter sets will be multiplexed in the data stream when detected.) (0 = disabled, -1 = send with every IDR frame)
// Default: 0, Min: -1, Max: 3600
func (e *GstRtpH264Pay) SetConfigInterval(value int) error {
	return e.Element.SetProperty("config-interval", value)
}

func (e *GstRtpH264Pay) GetConfigInterval() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("config-interval")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// The base64 sprop-parameter-sets to set in out caps (set to NULL to extract from stream)
// Default: NULL
func (e *GstRtpH264Pay) SetSpropParameterSets(value string) error {
	return e.Element.SetProperty("sprop-parameter-sets", value)
}

func (e *GstRtpH264Pay) GetSpropParameterSets() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("sprop-parameter-sets")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Payload-encodes JPEG pictures into RTP packets (RFC 2435)
type GstRtpJPEGPay struct {
	*GstRTPBasePayload
}

func NewRtpJPEGPay() (*GstRtpJPEGPay, error) {
	e, err := gst.NewElement("rtpjpegpay")
	if err != nil {
		return nil, err
	}
	return &GstRtpJPEGPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRtpJPEGPayWithName(name string) (*GstRtpJPEGPay, error) {
	e, err := gst.NewElementWithName("rtpjpegpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpJPEGPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Quality factor on JPEG data (unused)
// Default: 255, Min: 0, Max: 255
func (e *GstRtpJPEGPay) SetQuality(value int) error {
	return e.Element.SetProperty("quality", value)
}

func (e *GstRtpJPEGPay) GetQuality() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("quality")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Default JPEG Type, overwritten by SOF when present
// Default: 1, Min: 0, Max: 255
func (e *GstRtpJPEGPay) SetType(value int) error {
	return e.Element.SetProperty("type", value)
}

func (e *GstRtpJPEGPay) GetType() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("type")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Payload-encodes MPEG2 ES into RTP packets (RFC 2250)
type GstRTPMPVPay struct {
	*GstRTPBasePayload
}

func NewRTPMPVPay() (*GstRTPMPVPay, error) {
	e, err := gst.NewElement("rtpmpvpay")
	if err != nil {
		return nil, err
	}
	return &GstRTPMPVPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRTPMPVPayWithName(name string) (*GstRTPMPVPay, error) {
	e, err := gst.NewElementWithName("rtpmpvpay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPMPVPay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

// Extracts QDM2 audio from RTP packets (no RFC)
type GstRtpQDM2Depay struct {
	*GstRTPBaseDepayload
}

func NewRtpQDM2Depay() (*GstRtpQDM2Depay, error) {
	e, err := gst.NewElement("rtpqdm2depay")
	if err != nil {
		return nil, err
	}
	return &GstRtpQDM2Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRtpQDM2DepayWithName(name string) (*GstRtpQDM2Depay, error) {
	e, err := gst.NewElementWithName("rtpqdm2depay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpQDM2Depay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Depayloads RTP/RTCP packets for streaming protocols according to RFC4571
type GstRtpStreamDepay struct {
	*GstBaseParse
}

func NewRtpStreamDepay() (*GstRtpStreamDepay, error) {
	e, err := gst.NewElement("rtpstreamdepay")
	if err != nil {
		return nil, err
	}
	return &GstRtpStreamDepay{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

func NewRtpStreamDepayWithName(name string) (*GstRtpStreamDepay, error) {
	e, err := gst.NewElementWithName("rtpstreamdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpStreamDepay{GstBaseParse: &GstBaseParse{Element: e}}, nil
}

// Extracts H263 video from RTP and encodes in Asterisk H263 format
type GstAsteriskh263 struct {
	*gst.Element
}

func NewAsteriskh263() (*GstAsteriskh263, error) {
	e, err := gst.NewElement("asteriskh263")
	if err != nil {
		return nil, err
	}
	return &GstAsteriskh263{Element: e}, nil
}

func NewAsteriskh263WithName(name string) (*GstAsteriskh263, error) {
	e, err := gst.NewElementWithName("asteriskh263", name)
	if err != nil {
		return nil, err
	}
	return &GstAsteriskh263{Element: e}, nil
}

// Extracts iLBC audio from RTP packets (RFC 3952)
type GstRTPiLBCDepay struct {
	*GstRTPBaseDepayload
}

func NewRTPiLBCDepay() (*GstRTPiLBCDepay, error) {
	e, err := gst.NewElement("rtpilbcdepay")
	if err != nil {
		return nil, err
	}
	return &GstRTPiLBCDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRTPiLBCDepayWithName(name string) (*GstRTPiLBCDepay, error) {
	e, err := gst.NewElementWithName("rtpilbcdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPiLBCDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// iLBC frame mode
// Default: 30ms (30)
func (e *GstRTPiLBCDepay) SetMode(value ILBCMode) error {
	return e.Element.SetProperty("mode", value)
}

func (e *GstRTPiLBCDepay) GetMode() (ILBCMode, error) {
	var value ILBCMode
	var ok bool
	v, err := e.Element.GetProperty("mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(ILBCMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to ILBCMode")
	}
	return value, nil
}

// Extracts Opus audio from RTP packets
type GstRTPOpusDepay struct {
	*GstRTPBaseDepayload
}

func NewRTPOpusDepay() (*GstRTPOpusDepay, error) {
	e, err := gst.NewElement("rtpopusdepay")
	if err != nil {
		return nil, err
	}
	return &GstRTPOpusDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

func NewRTPOpusDepayWithName(name string) (*GstRTPOpusDepay, error) {
	e, err := gst.NewElementWithName("rtpopusdepay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPOpusDepay{GstRTPBaseDepayload: &GstRTPBaseDepayload{Element: e}}, nil
}

// Payloads RTP/RTCP packets for streaming protocols according to RFC4571
type GstRtpStreamPay struct {
	*gst.Element
}

func NewRtpStreamPay() (*GstRtpStreamPay, error) {
	e, err := gst.NewElement("rtpstreampay")
	if err != nil {
		return nil, err
	}
	return &GstRtpStreamPay{Element: e}, nil
}

func NewRtpStreamPayWithName(name string) (*GstRtpStreamPay, error) {
	e, err := gst.NewElementWithName("rtpstreampay", name)
	if err != nil {
		return nil, err
	}
	return &GstRtpStreamPay{Element: e}, nil
}

// Packetize G.729 audio into RTP packets
type GstRTPG729Pay struct {
	*GstRTPBasePayload
}

func NewRTPG729Pay() (*GstRTPG729Pay, error) {
	e, err := gst.NewElement("rtpg729pay")
	if err != nil {
		return nil, err
	}
	return &GstRTPG729Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

func NewRTPG729PayWithName(name string) (*GstRTPG729Pay, error) {
	e, err := gst.NewElementWithName("rtpg729pay", name)
	if err != nil {
		return nil, err
	}
	return &GstRTPG729Pay{GstRTPBasePayload: &GstRTPBasePayload{Element: e}}, nil
}

type GstDVPayMode string

const (
	GstDVPayModeVideo   GstDVPayMode = "video"   // Video only
	GstDVPayModeBundled GstDVPayMode = "bundled" // Video and Audio bundled
	GstDVPayModeAudio   GstDVPayMode = "audio"   // Audio only
)

type GstFragmentationMode string

const (
	GstFragmentationModeNormal GstFragmentationMode = "normal" // Normal
	GstFragmentationModeSync   GstFragmentationMode = "sync"   // Fragment at sync points
)

type GstRtpH264AggregateMode string

const (
	GstRtpH264AggregateModeNone        GstRtpH264AggregateMode = "none"         // Do not aggregate NAL units
	GstRtpH264AggregateModeZeroLatency GstRtpH264AggregateMode = "zero-latency" // Aggregate NAL units until a VCL unit is included
	GstRtpH264AggregateModeMaxStap     GstRtpH264AggregateMode = "max-stap"     // Aggregate all NAL units with the same timestamp (adds one frame of latency)
)

type GstRtpH265AggregateMode string

const (
	GstRtpH265AggregateModeNone        GstRtpH265AggregateMode = "none"         // Do not aggregate NAL units
	GstRtpH265AggregateModeZeroLatency GstRtpH265AggregateMode = "zero-latency" // Aggregate NAL units until a VCL or suffix unit is included
	GstRtpH265AggregateModeMax         GstRtpH265AggregateMode = "max"          // Aggregate all NAL units with the same timestamp (adds one frame of latency)
)

type GstVP8RTPPayMode string

const (
	GstVP8RTPPayModeNone  GstVP8RTPPayMode = "none"   // No Picture ID
	GstVP8RTPPayMode7Bit  GstVP8RTPPayMode = "7-bit"  // 7-bit Picture ID
	GstVP8RTPPayMode15Bit GstVP8RTPPayMode = "15-bit" // 15-bit Picture ID
)

type GstVP9RTPPayMode string

const (
	GstVP9RTPPayModeNone  GstVP9RTPPayMode = "none"   // No Picture ID
	GstVP9RTPPayMode7Bit  GstVP9RTPPayMode = "7-bit"  // 7-bit Picture ID
	GstVP9RTPPayMode15Bit GstVP9RTPPayMode = "15-bit" // 15-bit Picture ID
)

type ILBCMode string

const (
	ILBCMode20ms ILBCMode = "20ms" // 20ms frames
	ILBCMode30ms ILBCMode = "30ms" // 30ms frames
)
