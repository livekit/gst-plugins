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

type Gssink struct {
	*base.GstBaseSink
}

func NewGssink() (*Gssink, error) {
	e, err := gst.NewElement("gssink")
	if err != nil {
		return nil, err
	}
	return &Gssink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewGssinkWithName(name string) (*Gssink, error) {
	e, err := gst.NewElementWithName("gssink", name)
	if err != nil {
		return nil, err
	}
	return &Gssink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// bucket-name (string)
//
// Name of the Google Cloud Storage bucket.

func (e *Gssink) GetBucketName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("bucket-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Gssink) SetBucketName(value string) error {
	return e.Element.SetProperty("bucket-name", value)
}

// index (int)
//
// Index to use with location property to create file names.

func (e *Gssink) GetIndex() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("index")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Gssink) SetIndex(value int) error {
	return e.Element.SetProperty("index", value)
}

// next-file (GstGsSinkNext)
//
// A Gs-sink-next that specifies when to start a new file.

func (e *Gssink) GetNextFile() (GstGsSinkNext, error) {
	var value GstGsSinkNext
	var ok bool
	v, err := e.Element.GetProperty("next-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGsSinkNext)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGsSinkNext")
	}
	return value, nil
}

func (e *Gssink) SetNextFile(value GstGsSinkNext) error {
	e.Element.SetArg("next-file", string(value))
	return nil
}

// object-name (string)
//
// Full path name of the remote file.

func (e *Gssink) GetObjectName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("object-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Gssink) SetObjectName(value string) error {
	return e.Element.SetProperty("object-name", value)
}

// post-messages (bool)
//
// Post a message on the GstBus for each file.

func (e *Gssink) GetPostMessages() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-messages")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Gssink) SetPostMessages(value bool) error {
	return e.Element.SetProperty("post-messages", value)
}

// service-account-email (string)
//
// Service Account Email to use for credentials.

func (e *Gssink) GetServiceAccountEmail() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("service-account-email")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Gssink) SetServiceAccountEmail(value string) error {
	return e.Element.SetProperty("service-account-email", value)
}

// start-date (string)
//
// Start date in iso8601 format.

func (e *Gssink) GetStartDate() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("start-date")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Gssink) SetStartDate(value string) error {
	return e.Element.SetProperty("start-date", value)
}

// ----- Constants -----

type GstGsSinkNext string

const (
	GstGsSinkNextBuffer GstGsSinkNext = "buffer" // buffer (0) – New file for each buffer
	GstGsSinkNextOnlyOneFileNoNextFile GstGsSinkNext = "Only one file, no next file" // Only one file, no next file (1) – New file for each buffer
)

