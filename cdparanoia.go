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

type Cdparanoia struct {
	*base.GstPushSrc
}

func NewCdparanoia() (*Cdparanoia, error) {
	e, err := gst.NewElement("cdparanoia")
	if err != nil {
		return nil, err
	}
	return &Cdparanoia{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

func NewCdparanoiaWithName(name string) (*Cdparanoia, error) {
	e, err := gst.NewElementWithName("cdparanoia", name)
	if err != nil {
		return nil, err
	}
	return &Cdparanoia{GstPushSrc: &base.GstPushSrc{Element: e}}, nil
}

// ----- Properties -----

// cache-size (int)
//
// Set CD cache size to n sectors (-1 = auto)

func (e *Cdparanoia) GetCacheSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("cache-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cdparanoia) SetCacheSize(value int) error {
	return e.Element.SetProperty("cache-size", value)
}

// generic-device (string)
//
// Use specified generic scsi device

func (e *Cdparanoia) GetGenericDevice() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("generic-device")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Cdparanoia) SetGenericDevice(value string) error {
	return e.Element.SetProperty("generic-device", value)
}

// paranoia-mode (GstCdParanoiaMode)
//
// Type of checking to perform

func (e *Cdparanoia) GetParanoiaMode() (GstCdParanoiaMode, error) {
	var value GstCdParanoiaMode
	var ok bool
	v, err := e.Element.GetProperty("paranoia-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCdParanoiaMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCdParanoiaMode")
	}
	return value, nil
}

func (e *Cdparanoia) SetParanoiaMode(value GstCdParanoiaMode) error {
	e.Element.SetArg("paranoia-mode", string(value))
	return nil
}

// read-speed (int)
//
// Read from device at specified speed (-1 and 0 = full speed)

func (e *Cdparanoia) GetReadSpeed() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("read-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cdparanoia) SetReadSpeed(value int) error {
	return e.Element.SetProperty("read-speed", value)
}

// search-overlap (int)
//
// Force minimum overlap search during verification to n sectors

func (e *Cdparanoia) GetSearchOverlap() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("search-overlap")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Cdparanoia) SetSearchOverlap(value int) error {
	return e.Element.SetProperty("search-overlap", value)
}

// ----- Constants -----

type GstCdParanoiaMode string

const (
	GstCdParanoiaModeDisable GstCdParanoiaMode = "disable" // disable (0x00000000) – PARANOIA_MODE_DISABLE
	GstCdParanoiaModeFragment GstCdParanoiaMode = "fragment" // fragment (0x00000002) – PARANOIA_MODE_FRAGMENT
	GstCdParanoiaModeOverlap GstCdParanoiaMode = "overlap" // overlap (0x00000004) – PARANOIA_MODE_OVERLAP
	GstCdParanoiaModeScratch GstCdParanoiaMode = "scratch" // scratch (0x00000008) – PARANOIA_MODE_SCRATCH
	GstCdParanoiaModeRepair GstCdParanoiaMode = "repair" // repair (0x00000010) – PARANOIA_MODE_REPAIR
	GstCdParanoiaModeFull GstCdParanoiaMode = "full" // full (0x000000ff) – PARANOIA_MODE_FULL
)

