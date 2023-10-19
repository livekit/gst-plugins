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

type Curlsmtpsink struct {
	*base.GstBaseSink
}

func NewCurlsmtpsink() (*Curlsmtpsink, error) {
	e, err := gst.NewElement("curlsmtpsink")
	if err != nil {
		return nil, err
	}
	return &Curlsmtpsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewCurlsmtpsinkWithName(name string) (*Curlsmtpsink, error) {
	e, err := gst.NewElementWithName("curlsmtpsink", name)
	if err != nil {
		return nil, err
	}
	return &Curlsmtpsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// content-type (string)
//
// The mime type of the body of the request

func (e *Curlsmtpsink) GetContentType() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("content-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsmtpsink) SetContentType(value string) error {
	return e.Element.SetProperty("content-type", value)
}

// mail-from (string)
//
// Single address that the given mail should get sent from

func (e *Curlsmtpsink) GetMailFrom() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mail-from")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsmtpsink) SetMailFrom(value string) error {
	return e.Element.SetProperty("mail-from", value)
}

// mail-rcpt (string)
//
// Single address that the given mail should get sent to

func (e *Curlsmtpsink) GetMailRcpt() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("mail-rcpt")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsmtpsink) SetMailRcpt(value string) error {
	return e.Element.SetProperty("mail-rcpt", value)
}

// message-body (string)
//
// Message body

func (e *Curlsmtpsink) GetMessageBody() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("message-body")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsmtpsink) SetMessageBody(value string) error {
	return e.Element.SetProperty("message-body", value)
}

// nbr-attachments (int)
//
// Number attachments to send

func (e *Curlsmtpsink) GetNbrAttachments() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("nbr-attachments")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Curlsmtpsink) SetNbrAttachments(value int) error {
	return e.Element.SetProperty("nbr-attachments", value)
}

// pop-location (string)
//
// URL POP used for authentication

func (e *Curlsmtpsink) GetPopLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pop-location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsmtpsink) SetPopLocation(value string) error {
	return e.Element.SetProperty("pop-location", value)
}

// pop-passwd (string)
//
// User password to use for POP server authentication

func (e *Curlsmtpsink) GetPopPasswd() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pop-passwd")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsmtpsink) SetPopPasswd(value string) error {
	return e.Element.SetProperty("pop-passwd", value)
}

// pop-user (string)
//
// User name to use for POP server authentication

func (e *Curlsmtpsink) GetPopUser() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("pop-user")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsmtpsink) SetPopUser(value string) error {
	return e.Element.SetProperty("pop-user", value)
}

// subject (string)
//
// Mail subject

func (e *Curlsmtpsink) GetSubject() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subject")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsmtpsink) SetSubject(value string) error {
	return e.Element.SetProperty("subject", value)
}

// use-ssl (bool)
//
// Use SSL/TLS for the connection

func (e *Curlsmtpsink) GetUseSsl() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("use-ssl")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Curlsmtpsink) SetUseSsl(value bool) error {
	return e.Element.SetProperty("use-ssl", value)
}

