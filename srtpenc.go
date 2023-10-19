// Copyright 2023 LiveKit, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gstplugins

import (
	"fmt"

	"github.com/tinyzimmer/go-gst/gst"
)

type Srtpenc struct {
	Element *gst.Element
}

func NewSrtpenc() (*Srtpenc, error) {
	e, err := gst.NewElement("srtpenc")
	if err != nil {
		return nil, err
	}
	return &Srtpenc{Element: e}, nil
}

func NewSrtpencWithName(name string) (*Srtpenc, error) {
	e, err := gst.NewElementWithName("srtpenc", name)
	if err != nil {
		return nil, err
	}
	return &Srtpenc{Element: e}, nil
}

// ----- Properties -----

// allow-repeat-tx (bool)
//
// Whether retransmissions of packets with the same sequence number are allowed(Note that such repeated transmissions must have the same RTP payload, or a severe security weakness is introduced!)

func (e *Srtpenc) GetAllowRepeatTx() (bool, error) {
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

func (e *Srtpenc) SetAllowRepeatTx(value bool) error {
	return e.Element.SetProperty("allow-repeat-tx", value)
}

// key (GstBuffer)
//
// Master key (minimum of 30 and maximum of 46 bytes)

func (e *Srtpenc) GetKey() (interface{}, error) {
	return e.Element.GetProperty("key")
}

func (e *Srtpenc) SetKey(value interface{}) error {
	return e.Element.SetProperty("key", value)
}

// mki (GstBuffer)
//
// Master key Identifier (NULL means no MKI)

func (e *Srtpenc) GetMki() (interface{}, error) {
	return e.Element.GetProperty("mki")
}

func (e *Srtpenc) SetMki(value interface{}) error {
	return e.Element.SetProperty("mki", value)
}

// random-key (bool)
//
// Generate a random key if TRUE

func (e *Srtpenc) GetRandomKey() (bool, error) {
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

func (e *Srtpenc) SetRandomKey(value bool) error {
	return e.Element.SetProperty("random-key", value)
}

// replay-window-size (uint)
//
// Size of the replay protection window

func (e *Srtpenc) GetReplayWindowSize() (uint, error) {
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

func (e *Srtpenc) SetReplayWindowSize(value uint) error {
	return e.Element.SetProperty("replay-window-size", value)
}

// rtcp-auth (GstSrtpAuthType)
//
// RTCP Authentication

func (e *Srtpenc) GetRtcpAuth() (GstSrtpAuthType, error) {
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

func (e *Srtpenc) SetRtcpAuth(value GstSrtpAuthType) error {
	e.Element.SetArg("rtcp-auth", string(value))
	return nil
}

// rtcp-cipher (GstSrtpCipherType)
//
// RTCP Cipher

func (e *Srtpenc) GetRtcpCipher() (GstSrtpCipherType, error) {
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

func (e *Srtpenc) SetRtcpCipher(value GstSrtpCipherType) error {
	e.Element.SetArg("rtcp-cipher", string(value))
	return nil
}

// rtp-auth (GstSrtpAuthType)
//
// RTP Authentication

func (e *Srtpenc) GetRtpAuth() (GstSrtpAuthType, error) {
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

func (e *Srtpenc) SetRtpAuth(value GstSrtpAuthType) error {
	e.Element.SetArg("rtp-auth", string(value))
	return nil
}

// rtp-cipher (GstSrtpCipherType)
//
// RTP Cipher

func (e *Srtpenc) GetRtpCipher() (GstSrtpCipherType, error) {
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

func (e *Srtpenc) SetRtpCipher(value GstSrtpCipherType) error {
	e.Element.SetArg("rtp-cipher", string(value))
	return nil
}

// stats (GstStructure)
//
// Various statistics

func (e *Srtpenc) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// ----- Constants -----

type GstSrtpAuthType string

const (
	GstSrtpAuthTypeNull GstSrtpAuthType = "null" // null (0) – GST_SRTP_AUTH_NULL
	GstSrtpAuthTypeHmacSha132 GstSrtpAuthType = "hmac-sha1-32" // hmac-sha1-32 (1) – GST_SRTP_AUTH_HMAC_SHA1_32
	GstSrtpAuthTypeHmacSha180 GstSrtpAuthType = "hmac-sha1-80" // hmac-sha1-80 (2) – GST_SRTP_AUTH_HMAC_SHA1_80
)

type GstSrtpCipherType string

const (
	GstSrtpCipherTypeNull GstSrtpCipherType = "null" // null (0) – GST_SRTP_CIPHER_NULL
	GstSrtpCipherTypeAes128Icm GstSrtpCipherType = "aes-128-icm" // aes-128-icm (1) – GST_SRTP_CIPHER_AES_128_ICM
	GstSrtpCipherTypeAes256Icm GstSrtpCipherType = "aes-256-icm" // aes-256-icm (2) – GST_SRTP_CIPHER_AES_256_ICM
	GstSrtpCipherTypeAes128Gcm GstSrtpCipherType = "aes-128-gcm" // aes-128-gcm (3) – GST_SRTP_CIPHER_AES_128_GCM
	GstSrtpCipherTypeAes256Gcm GstSrtpCipherType = "aes-256-gcm" // aes-256-gcm (4) – GST_SRTP_CIPHER_AES_256_GCM
)

