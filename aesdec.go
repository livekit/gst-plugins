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

type Aesdec struct {
	*base.GstBaseTransform
}

func NewAesdec() (*Aesdec, error) {
	e, err := gst.NewElement("aesdec")
	if err != nil {
		return nil, err
	}
	return &Aesdec{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAesdecWithName(name string) (*Aesdec, error) {
	e, err := gst.NewElementWithName("aesdec", name)
	if err != nil {
		return nil, err
	}
	return &Aesdec{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// cipher (GstAesCipher)
//
// AES cipher mode (key length and mode)
// Currently, 128 and 256 bit keys are supported,
// in "cipher block chaining" (CBC) mode

func (e *Aesdec) GetCipher() (interface{}, error) {
	return e.Element.GetProperty("cipher")
}

func (e *Aesdec) SetCipher(value interface{}) error {
	return e.Element.SetProperty("cipher", value)
}

// iv (string)
//
// AES encryption initialization vector (hexadecimal)

func (e *Aesdec) GetIv() (string, error) {
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

func (e *Aesdec) SetIv(value string) error {
	return e.Element.SetProperty("iv", value)
}

// key (string)
//
// AES encryption key (hexadecimal)

func (e *Aesdec) GetKey() (string, error) {
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

func (e *Aesdec) SetKey(value string) error {
	return e.Element.SetProperty("key", value)
}

// per-buffer-padding (bool)
//
// If true, each buffer will be padded using PKCS7 padding
// If false, only the final buffer in the stream will be padded
// (by OpenSSL) using PKCS7

func (e *Aesdec) GetPerBufferPadding() (bool, error) {
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

func (e *Aesdec) SetPerBufferPadding(value bool) error {
	return e.Element.SetProperty("per-buffer-padding", value)
}

// serialize-iv (bool)
//
// If true, read initialization vector from first 16 bytes of first buffer

func (e *Aesdec) GetSerializeIv() (bool, error) {
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

func (e *Aesdec) SetSerializeIv(value bool) error {
	return e.Element.SetProperty("serialize-iv", value)
}

