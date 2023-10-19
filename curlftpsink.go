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

type Curlftpsink struct {
	*base.GstBaseSink
}

func NewCurlftpsink() (*Curlftpsink, error) {
	e, err := gst.NewElement("curlftpsink")
	if err != nil {
		return nil, err
	}
	return &Curlftpsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewCurlftpsinkWithName(name string) (*Curlftpsink, error) {
	e, err := gst.NewElementWithName("curlftpsink", name)
	if err != nil {
		return nil, err
	}
	return &Curlftpsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// create-dirs (bool)
//
// Attempt to create missing directory included in the path

func (e *Curlftpsink) GetCreateDirs() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("create-dirs")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Curlftpsink) SetCreateDirs(value bool) error {
	return e.Element.SetProperty("create-dirs", value)
}

// create-tmp-file (bool)
//
// Use a temporary file name when uploading a a file. When the transfer is complete,            this temporary file is renamed to the final file name. This is useful for ensuring           that remote systems do not read a partially uploaded file

func (e *Curlftpsink) GetCreateTmpFile() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("create-tmp-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Curlftpsink) SetCreateTmpFile(value bool) error {
	return e.Element.SetProperty("create-tmp-file", value)
}

// epsv-mode (bool)
//
// Enable the use of the EPSV command when doing passive FTP transfers

func (e *Curlftpsink) GetEpsvMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("epsv-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Curlftpsink) SetEpsvMode(value bool) error {
	return e.Element.SetProperty("epsv-mode", value)
}

// ftp-port (string)
//
// The PORT instruction tells the remote server to connect to the IP address

func (e *Curlftpsink) GetFtpPort() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ftp-port")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlftpsink) SetFtpPort(value string) error {
	return e.Element.SetProperty("ftp-port", value)
}

// temp-file-name (string)
//
// Filename pattern to use when generating a temporary filename for uploads

func (e *Curlftpsink) GetTempFileName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("temp-file-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlftpsink) SetTempFileName(value string) error {
	return e.Element.SetProperty("temp-file-name", value)
}

