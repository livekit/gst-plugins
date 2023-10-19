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

type Checksumsink struct {
	*base.GstBaseSink
}

func NewChecksumsink() (*Checksumsink, error) {
	e, err := gst.NewElement("checksumsink")
	if err != nil {
		return nil, err
	}
	return &Checksumsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewChecksumsinkWithName(name string) (*Checksumsink, error) {
	e, err := gst.NewElementWithName("checksumsink", name)
	if err != nil {
		return nil, err
	}
	return &Checksumsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// hash (GstChecksumSinkHash)
//
// Checksum type

func (e *Checksumsink) GetHash() (GstChecksumSinkHash, error) {
	var value GstChecksumSinkHash
	var ok bool
	v, err := e.Element.GetProperty("hash")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstChecksumSinkHash)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstChecksumSinkHash")
	}
	return value, nil
}

func (e *Checksumsink) SetHash(value GstChecksumSinkHash) error {
	e.Element.SetArg("hash", string(value))
	return nil
}

// ----- Constants -----

type GstChecksumSinkHash string

const (
	GstChecksumSinkHashMd5 GstChecksumSinkHash = "md5" // md5 (0) – MD5
	GstChecksumSinkHashSha1 GstChecksumSinkHash = "sha1" // sha1 (1) – SHA-1
	GstChecksumSinkHashSha256 GstChecksumSinkHash = "sha256" // sha256 (2) – SHA-256
	GstChecksumSinkHashSha512 GstChecksumSinkHash = "sha512" // sha512 (3) – SHA-512
)

