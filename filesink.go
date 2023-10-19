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

type Filesink struct {
	*base.GstBaseSink
}

func NewFilesink() (*Filesink, error) {
	e, err := gst.NewElement("filesink")
	if err != nil {
		return nil, err
	}
	return &Filesink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewFilesinkWithName(name string) (*Filesink, error) {
	e, err := gst.NewElementWithName("filesink", name)
	if err != nil {
		return nil, err
	}
	return &Filesink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// append (bool)
//
// Append to an already existing file.

func (e *Filesink) GetAppend() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("append")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Filesink) SetAppend(value bool) error {
	return e.Element.SetProperty("append", value)
}

// buffer-mode (GstFileSinkBufferMode)
//
// The buffering mode to use

func (e *Filesink) GetBufferMode() (GstFileSinkBufferMode, error) {
	var value GstFileSinkBufferMode
	var ok bool
	v, err := e.Element.GetProperty("buffer-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFileSinkBufferMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFileSinkBufferMode")
	}
	return value, nil
}

func (e *Filesink) SetBufferMode(value GstFileSinkBufferMode) error {
	e.Element.SetArg("buffer-mode", string(value))
	return nil
}

// buffer-size (uint)
//
// Size of buffer in number of bytes for line or full buffer-mode

func (e *Filesink) GetBufferSize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Filesink) SetBufferSize(value uint) error {
	return e.Element.SetProperty("buffer-size", value)
}

// file-mode (GstFileSinkFileMode)
//
// Ability to specify file mode.

func (e *Filesink) GetFileMode() (GstFileSinkFileMode, error) {
	var value GstFileSinkFileMode
	var ok bool
	v, err := e.Element.GetProperty("file-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstFileSinkFileMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstFileSinkFileMode")
	}
	return value, nil
}

func (e *Filesink) SetFileMode(value GstFileSinkFileMode) error {
	e.Element.SetArg("file-mode", string(value))
	return nil
}

// location (string)
//
// Location of the file to write

func (e *Filesink) GetLocation() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("location")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Filesink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// max-transient-error-timeout (int)
//
// Retry up to this many ms on transient errors (currently EACCES)

func (e *Filesink) GetMaxTransientErrorTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-transient-error-timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Filesink) SetMaxTransientErrorTimeout(value int) error {
	return e.Element.SetProperty("max-transient-error-timeout", value)
}

// o-sync (bool)
//
// Open the file with O_SYNC for enabling synchronous IO

func (e *Filesink) GetOSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("o-sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Filesink) SetOSync(value bool) error {
	return e.Element.SetProperty("o-sync", value)
}

// ----- Constants -----

type GstFileSinkFileMode string

const (
	GstFileSinkFileModeTruncate GstFileSinkFileMode = "truncate" // truncate (1) – Truncate file (mode wb)
	GstFileSinkFileModeOutput GstFileSinkFileMode = "output" // output (2) – Append file (mode ab)
	GstFileSinkFileModeOverwrite GstFileSinkFileMode = "overwrite" // overwrite (3) – Overwrite file without truncating (mode rb+)
)

type GstFileSinkBufferMode string

const (
	GstFileSinkBufferModeDefault GstFileSinkBufferMode = "default" // default (-1) – Default buffering
	GstFileSinkBufferModeFull GstFileSinkBufferMode = "full" // full (0) – Fully buffered
	GstFileSinkBufferModeLine GstFileSinkBufferMode = "line" // line (1) – Line buffered (deprecated, like full)
	GstFileSinkBufferModeUnbuffered GstFileSinkBufferMode = "unbuffered" // unbuffered (2) – Unbuffered
)

