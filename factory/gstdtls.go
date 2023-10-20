// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// Decodes DTLS packets
type GstDtlsDec struct {
	*gst.Element
}

func NewDtlsDec() (*GstDtlsDec, error) {
	e, err := gst.NewElement("dtlsdec")
	if err != nil {
		return nil, err
	}
	return &GstDtlsDec{Element: e}, nil
}

func NewDtlsDecWithName(name string) (*GstDtlsDec, error) {
	e, err := gst.NewElementWithName("dtlsdec", name)
	if err != nil {
		return nil, err
	}
	return &GstDtlsDec{Element: e}, nil
}

// SRTP key that should be used by the decoder

func (e *GstDtlsDec) GetDecoderKey() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("decoder-key")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// The X509 certificate received in the DTLS handshake, in PEM format
// Default: NULL
func (e *GstDtlsDec) GetPeerPem() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("peer-pem")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// A string containing a X509 certificate and RSA private key in PEM format
// Default: NULL
func (e *GstDtlsDec) SetPem(value string) error {
	return e.Element.SetProperty("pem", value)
}

func (e *GstDtlsDec) GetPem() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pem")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// The SRTP authentication selected in the DTLS handshake. The value will be set to an GstDtlsSrtpAuth.
// Default: 0, Min: 0, Max: 2
func (e *GstDtlsDec) GetSrtpAuth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("srtp-auth")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The SRTP cipher selected in the DTLS handshake. The value will be set to an GstDtlsSrtpCipher.
// Default: 0, Min: 0, Max: 1
func (e *GstDtlsDec) GetSrtpCipher() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("srtp-cipher")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Every encoder/decoder pair should have the same, unique, connection-id
// Default: NULL
func (e *GstDtlsDec) SetConnectionId(value string) error {
	return e.Element.SetProperty("connection-id", value)
}

func (e *GstDtlsDec) GetConnectionId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("connection-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Current connection state
// Default: closed (1)
func (e *GstDtlsDec) GetConnectionState() (GstDtlsConnectionState, error) {
	var value GstDtlsConnectionState
	var ok bool
	v, err := e.Element.GetProperty("connection-state")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDtlsConnectionState)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDtlsConnectionState")
	}
	return value, nil
}

// Encodes packets with DTLS
type GstDtlsEnc struct {
	*gst.Element
}

func NewDtlsEnc() (*GstDtlsEnc, error) {
	e, err := gst.NewElement("dtlsenc")
	if err != nil {
		return nil, err
	}
	return &GstDtlsEnc{Element: e}, nil
}

func NewDtlsEncWithName(name string) (*GstDtlsEnc, error) {
	e, err := gst.NewElementWithName("dtlsenc", name)
	if err != nil {
		return nil, err
	}
	return &GstDtlsEnc{Element: e}, nil
}

// Set to true if the decoder should act as client and initiate the handshake
// Default: false
func (e *GstDtlsEnc) SetIsClient(value bool) error {
	return e.Element.SetProperty("is-client", value)
}

func (e *GstDtlsEnc) GetIsClient() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-client")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// The SRTP authentication selected in the DTLS handshake. The value will be set to an GstDtlsSrtpAuth.
// Default: 0, Min: 0, Max: 2
func (e *GstDtlsEnc) GetSrtpAuth() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("srtp-auth")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// The SRTP cipher selected in the DTLS handshake. The value will be set to an GstDtlsSrtpCipher.
// Default: 0, Min: 0, Max: 1
func (e *GstDtlsEnc) GetSrtpCipher() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("srtp-cipher")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Every encoder/decoder pair should have the same, unique, connection-id
// Default: NULL
func (e *GstDtlsEnc) SetConnectionId(value string) error {
	return e.Element.SetProperty("connection-id", value)
}

func (e *GstDtlsEnc) GetConnectionId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("connection-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Current connection state
// Default: closed (1)
func (e *GstDtlsEnc) GetConnectionState() (GstDtlsConnectionState, error) {
	var value GstDtlsConnectionState
	var ok bool
	v, err := e.Element.GetProperty("connection-state")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDtlsConnectionState)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDtlsConnectionState")
	}
	return value, nil
}

// Master key that should be used by the SRTP encoder

func (e *GstDtlsEnc) GetEncoderKey() (*gst.Buffer, error) {
	var value *gst.Buffer
	var ok bool
	v, err := e.Element.GetProperty("encoder-key")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Buffer)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Buffer")
	}
	return value, nil
}

// Decodes SRTP packets with a key received from DTLS
type GstDtlsSrtpDec struct {
	*GstDtlsSrtpBin
}

func NewDtlsSrtpDec() (*GstDtlsSrtpDec, error) {
	e, err := gst.NewElement("dtlssrtpdec")
	if err != nil {
		return nil, err
	}
	return &GstDtlsSrtpDec{GstDtlsSrtpBin: &GstDtlsSrtpBin{Bin: &gst.Bin{Element: e}}}, nil
}

func NewDtlsSrtpDecWithName(name string) (*GstDtlsSrtpDec, error) {
	e, err := gst.NewElementWithName("dtlssrtpdec", name)
	if err != nil {
		return nil, err
	}
	return &GstDtlsSrtpDec{GstDtlsSrtpBin: &GstDtlsSrtpBin{Bin: &gst.Bin{Element: e}}}, nil
}

// Current connection state
// Default: closed (1)
func (e *GstDtlsSrtpDec) GetConnectionState() (GstDtlsConnectionState, error) {
	var value GstDtlsConnectionState
	var ok bool
	v, err := e.Element.GetProperty("connection-state")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDtlsConnectionState)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDtlsConnectionState")
	}
	return value, nil
}

// The X509 certificate received in the DTLS handshake, in PEM format
// Default: NULL
func (e *GstDtlsSrtpDec) GetPeerPem() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("peer-pem")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// A string containing a X509 certificate and RSA private key in PEM format
// Default: NULL
func (e *GstDtlsSrtpDec) SetPem(value string) error {
	return e.Element.SetProperty("pem", value)
}

func (e *GstDtlsSrtpDec) GetPem() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pem")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Demultiplexes DTLS and SRTP packets
type GstDtlsSrtpDemux struct {
	*gst.Element
}

func NewDtlsSrtpDemux() (*GstDtlsSrtpDemux, error) {
	e, err := gst.NewElement("dtlssrtpdemux")
	if err != nil {
		return nil, err
	}
	return &GstDtlsSrtpDemux{Element: e}, nil
}

func NewDtlsSrtpDemuxWithName(name string) (*GstDtlsSrtpDemux, error) {
	e, err := gst.NewElementWithName("dtlssrtpdemux", name)
	if err != nil {
		return nil, err
	}
	return &GstDtlsSrtpDemux{Element: e}, nil
}

// Encodes SRTP packets with a key received from DTLS
type GstDtlsSrtpEnc struct {
	*GstDtlsSrtpBin
}

func NewDtlsSrtpEnc() (*GstDtlsSrtpEnc, error) {
	e, err := gst.NewElement("dtlssrtpenc")
	if err != nil {
		return nil, err
	}
	return &GstDtlsSrtpEnc{GstDtlsSrtpBin: &GstDtlsSrtpBin{Bin: &gst.Bin{Element: e}}}, nil
}

func NewDtlsSrtpEncWithName(name string) (*GstDtlsSrtpEnc, error) {
	e, err := gst.NewElementWithName("dtlssrtpenc", name)
	if err != nil {
		return nil, err
	}
	return &GstDtlsSrtpEnc{GstDtlsSrtpBin: &GstDtlsSrtpBin{Bin: &gst.Bin{Element: e}}}, nil
}

// Current connection state
// Default: closed (1)
func (e *GstDtlsSrtpEnc) GetConnectionState() (GstDtlsConnectionState, error) {
	var value GstDtlsConnectionState
	var ok bool
	v, err := e.Element.GetProperty("connection-state")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstDtlsConnectionState)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstDtlsConnectionState")
	}
	return value, nil
}

// Set to true if the decoder should act as client and initiate the handshake
// Default: false
func (e *GstDtlsSrtpEnc) SetIsClient(value bool) error {
	return e.Element.SetProperty("is-client", value)
}

func (e *GstDtlsSrtpEnc) GetIsClient() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-client")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Synchronize RTP to the pipeline clock before merging with RTCP
// Default: false
func (e *GstDtlsSrtpEnc) SetRtpSync(value bool) error {
	return e.Element.SetProperty("rtp-sync", value)
}

func (e *GstDtlsSrtpEnc) GetRtpSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("rtp-sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstDtlsConnectionState string

const (
	GstDtlsConnectionStateNew        GstDtlsConnectionState = "new"        // New connection
	GstDtlsConnectionStateClosed     GstDtlsConnectionState = "closed"     // Closed connection on either side
	GstDtlsConnectionStateFailed     GstDtlsConnectionState = "failed"     // Failed connection
	GstDtlsConnectionStateConnecting GstDtlsConnectionState = "connecting" // Connecting
	GstDtlsConnectionStateConnected  GstDtlsConnectionState = "connected"  // Successfully connected
)

type GstDtlsSrtpBin struct {
	*gst.Bin
}

// SRTP master key, if this property is set, DTLS will be disabled

func (e *GstDtlsSrtpBin) SetKey(value *gst.Buffer) error {
	return e.Element.SetProperty("key", value)
}

func (e *GstDtlsSrtpBin) GetKey() (*gst.Buffer, error) {
	var value *gst.Buffer
	var ok bool
	v, err := e.Element.GetProperty("key")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Buffer)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Buffer")
	}
	return value, nil
}

// SRTCP auth name, should be 'null', 'hmac-sha1-32' or 'hmac-sha1-80', if this property is set, DTLS will be disabled
// Default: NULL
func (e *GstDtlsSrtpBin) SetSrtcpAuth(value string) error {
	return e.Element.SetProperty("srtcp-auth", value)
}

func (e *GstDtlsSrtpBin) GetSrtcpAuth() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("srtcp-auth")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// SRTCP cipher name, should be 'null' or 'aes-128-icm', if this property is set, DTLS will be disabled
// Default: NULL
func (e *GstDtlsSrtpBin) SetSrtcpCipher(value string) error {
	return e.Element.SetProperty("srtcp-cipher", value)
}

func (e *GstDtlsSrtpBin) GetSrtcpCipher() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("srtcp-cipher")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// SRTP auth name, should be 'null', 'hmac-sha1-32' or 'hmac-sha1-80', if this property is set, DTLS will be disabled
// Default: NULL
func (e *GstDtlsSrtpBin) SetSrtpAuth(value string) error {
	return e.Element.SetProperty("srtp-auth", value)
}

func (e *GstDtlsSrtpBin) GetSrtpAuth() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("srtp-auth")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// SRTP cipher name, should be 'null' or 'aes-128-icm', if this property is set, DTLS will be disabled
// Default: NULL
func (e *GstDtlsSrtpBin) SetSrtpCipher(value string) error {
	return e.Element.SetProperty("srtp-cipher", value)
}

func (e *GstDtlsSrtpBin) GetSrtpCipher() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("srtp-cipher")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Every encoder/decoder pair should have the same, unique, connection-id
// Default: NULL
func (e *GstDtlsSrtpBin) SetConnectionId(value string) error {
	return e.Element.SetProperty("connection-id", value)
}

func (e *GstDtlsSrtpBin) GetConnectionId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("connection-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}
