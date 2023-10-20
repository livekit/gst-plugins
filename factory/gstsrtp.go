// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
)

// A SRTP and SRTCP decoder
type GstSrtpDec struct {
	*gst.Element
}

func NewSrtpDec() (*GstSrtpDec, error) {
	e, err := gst.NewElement("srtpdec")
	if err != nil {
		return nil, err
	}
	return &GstSrtpDec{Element: e}, nil
}

func NewSrtpDecWithName(name string) (*GstSrtpDec, error) {
	e, err := gst.NewElementWithName("srtpdec", name)
	if err != nil {
		return nil, err
	}
	return &GstSrtpDec{Element: e}, nil
}

// Size of the replay protection window
// Default: 128, Min: 64, Max: 32768
func (e *GstSrtpDec) SetReplayWindowSize(value uint) error {
	return e.Element.SetProperty("replay-window-size", value)
}

func (e *GstSrtpDec) GetReplayWindowSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("replay-window-size")
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
// Default: application/x-srtp-decoder-stats, streams=(int)<  >, recv-count=(uint)0, recv-drop-count=(uint)0;
func (e *GstSrtpDec) GetStats() (*gst.Structure, error) {
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

// A SRTP and SRTCP encoder
type GstSrtpEnc struct {
	*gst.Element
}

func NewSrtpEnc() (*GstSrtpEnc, error) {
	e, err := gst.NewElement("srtpenc")
	if err != nil {
		return nil, err
	}
	return &GstSrtpEnc{Element: e}, nil
}

func NewSrtpEncWithName(name string) (*GstSrtpEnc, error) {
	e, err := gst.NewElementWithName("srtpenc", name)
	if err != nil {
		return nil, err
	}
	return &GstSrtpEnc{Element: e}, nil
}

// Whether retransmissions of packets with the same sequence number are allowed(Note that such repeated transmissions must have the same RTP payload, or a severe security weakness is introduced!)
// Default: false
func (e *GstSrtpEnc) SetAllowRepeatTx(value bool) error {
	return e.Element.SetProperty("allow-repeat-tx", value)
}

func (e *GstSrtpEnc) GetAllowRepeatTx() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-repeat-tx")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Master key (minimum of 30 and maximum of 46 bytes)

func (e *GstSrtpEnc) SetKey(value *gst.Buffer) error {
	return e.Element.SetProperty("key", value)
}

func (e *GstSrtpEnc) GetKey() (*gst.Buffer, error) {
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

// Master key Identifier (NULL means no MKI)

func (e *GstSrtpEnc) SetMki(value *gst.Buffer) error {
	return e.Element.SetProperty("mki", value)
}

func (e *GstSrtpEnc) GetMki() (*gst.Buffer, error) {
	var value *gst.Buffer
	var ok bool
	v, err := e.Element.GetProperty("mki")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Buffer)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Buffer")
	}
	return value, nil
}

// Generate a random key if TRUE
// Default: false
func (e *GstSrtpEnc) SetRandomKey(value bool) error {
	return e.Element.SetProperty("random-key", value)
}

func (e *GstSrtpEnc) GetRandomKey() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("random-key")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// RTP Authentication
// Default: hmac-sha1-80 (2)
func (e *GstSrtpEnc) SetRtpAuth(value GstSrtpAuthType) error {
	e.Element.SetArg("rtp-auth", string(value))
	return nil
}

func (e *GstSrtpEnc) GetRtpAuth() (GstSrtpAuthType, error) {
	var value GstSrtpAuthType
	var ok bool
	v, err := e.Element.GetProperty("rtp-auth")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSrtpAuthType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSrtpAuthType")
	}
	return value, nil
}

// Various statistics
// Default: application/x-srtp-encoder-stats, streams=(int)<  >;
func (e *GstSrtpEnc) GetStats() (*gst.Structure, error) {
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

// Size of the replay protection window
// Default: 128, Min: 64, Max: 32768
func (e *GstSrtpEnc) SetReplayWindowSize(value uint) error {
	return e.Element.SetProperty("replay-window-size", value)
}

func (e *GstSrtpEnc) GetReplayWindowSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("replay-window-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// RTCP Authentication
// Default: hmac-sha1-80 (2)
func (e *GstSrtpEnc) SetRtcpAuth(value GstSrtpAuthType) error {
	e.Element.SetArg("rtcp-auth", string(value))
	return nil
}

func (e *GstSrtpEnc) GetRtcpAuth() (GstSrtpAuthType, error) {
	var value GstSrtpAuthType
	var ok bool
	v, err := e.Element.GetProperty("rtcp-auth")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSrtpAuthType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSrtpAuthType")
	}
	return value, nil
}

// RTCP Cipher
// Default: aes-128-icm (1)
func (e *GstSrtpEnc) SetRtcpCipher(value GstSrtpCipherType) error {
	e.Element.SetArg("rtcp-cipher", string(value))
	return nil
}

func (e *GstSrtpEnc) GetRtcpCipher() (GstSrtpCipherType, error) {
	var value GstSrtpCipherType
	var ok bool
	v, err := e.Element.GetProperty("rtcp-cipher")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSrtpCipherType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSrtpCipherType")
	}
	return value, nil
}

// RTP Cipher
// Default: aes-128-icm (1)
func (e *GstSrtpEnc) SetRtpCipher(value GstSrtpCipherType) error {
	e.Element.SetArg("rtp-cipher", string(value))
	return nil
}

func (e *GstSrtpEnc) GetRtpCipher() (GstSrtpCipherType, error) {
	var value GstSrtpCipherType
	var ok bool
	v, err := e.Element.GetProperty("rtp-cipher")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstSrtpCipherType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstSrtpCipherType")
	}
	return value, nil
}

type GstSrtpAuthType string

const (
	GstSrtpAuthTypeNull       GstSrtpAuthType = "null"         // GST_SRTP_AUTH_NULL
	GstSrtpAuthTypeHmacSha132 GstSrtpAuthType = "hmac-sha1-32" // GST_SRTP_AUTH_HMAC_SHA1_32
	GstSrtpAuthTypeHmacSha180 GstSrtpAuthType = "hmac-sha1-80" // GST_SRTP_AUTH_HMAC_SHA1_80
)

type GstSrtpCipherType string

const (
	GstSrtpCipherTypeNull      GstSrtpCipherType = "null"        // GST_SRTP_CIPHER_NULL
	GstSrtpCipherTypeAes128Icm GstSrtpCipherType = "aes-128-icm" // GST_SRTP_CIPHER_AES_128_ICM
	GstSrtpCipherTypeAes256Icm GstSrtpCipherType = "aes-256-icm" // GST_SRTP_CIPHER_AES_256_ICM
	GstSrtpCipherTypeAes128Gcm GstSrtpCipherType = "aes-128-gcm" // GST_SRTP_CIPHER_AES_128_GCM
	GstSrtpCipherTypeAes256Gcm GstSrtpCipherType = "aes-256-gcm" // GST_SRTP_CIPHER_AES_256_GCM
)
