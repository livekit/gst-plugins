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

	"github.com/livekit/gstplugins/base"
)

type Aesenc struct {
	*base.GstBaseTransform
}

func NewAesenc() (*Aesenc, error) {
	e, err := gst.NewElement("aesenc")
	if err != nil {
		return nil, err
	}
	return &Aesenc{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAesencWithName(name string) (*Aesenc, error) {
	e, err := gst.NewElementWithName("aesenc", name)
	if err != nil {
		return nil, err
	}
	return &Aesenc{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// cipher (GstAesCipher)
//
// AES cipher mode (key length and mode)
// Currently, 128 and 256 bit keys are supported,
// in "cipher block chaining" (CBC) mode

func (e *Aesenc) GetCipher() (GstAesCipher, error) {
	var value GstAesCipher
	var ok bool
	v, err := e.Element.GetProperty("cipher")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAesCipher)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAesCipher")
	}
	return value, nil
}

func (e *Aesenc) SetCipher(value GstAesCipher) error {
	e.Element.SetArg("cipher", string(value))
	return nil
}

// iv (string)
//
// AES encryption initialization vector (hexadecimal)

func (e *Aesenc) GetIv() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("iv")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Aesenc) SetIv(value string) error {
	return e.Element.SetProperty("iv", value)
}

// key (string)
//
// AES encryption key (hexadecimal)

func (e *Aesenc) GetKey() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("key")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Aesenc) SetKey(value string) error {
	return e.Element.SetProperty("key", value)
}

// per-buffer-padding (bool)
//
// If true, each buffer will be padded using PKCS7 padding
// If false, only the final buffer in the stream will be padded
// (by OpenSSL) using PKCS7

func (e *Aesenc) GetPerBufferPadding() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("per-buffer-padding")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Aesenc) SetPerBufferPadding(value bool) error {
	return e.Element.SetProperty("per-buffer-padding", value)
}

// serialize-iv (bool)
//
// If true, store initialization vector in first 16 bytes of first buffer

func (e *Aesenc) GetSerializeIv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("serialize-iv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Aesenc) SetSerializeIv(value bool) error {
	return e.Element.SetProperty("serialize-iv", value)
}

// ----- Constants -----

type GstAesCipher string

const (
	GstAesCipherAes128Cbc GstAesCipher = "aes-128-cbc" // aes-128-cbc (0) – AES 128 bit cipher key using CBC method
	GstAesCipherAes256Cbc GstAesCipher = "aes-256-cbc" // aes-256-cbc (1) – AES 256 bit cipher key using CBC method
)

