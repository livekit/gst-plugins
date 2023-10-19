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

type Netsim struct {
	Element *gst.Element
}

func NewNetsim() (*Netsim, error) {
	e, err := gst.NewElement("netsim")
	if err != nil {
		return nil, err
	}
	return &Netsim{Element: e}, nil
}

func NewNetsimWithName(name string) (*Netsim, error) {
	e, err := gst.NewElementWithName("netsim", name)
	if err != nil {
		return nil, err
	}
	return &Netsim{Element: e}, nil
}

// ----- Properties -----

// allow-reordering (bool)
//
// When delaying packets, are they allowed to be reordered or not. By
// default this is enabled, but in the real world packet reordering is
// fairly uncommon, yet the delay functions will always introduce reordering
// if delay > packet-spacing, This property allows switching that off.

func (e *Netsim) GetAllowReordering() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("allow-reordering")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Netsim) SetAllowReordering(value bool) error {
	return e.Element.SetProperty("allow-reordering", value)
}

// delay-distribution (GstNetSimDistribution)
//
// Distribution for the amount of delay.

func (e *Netsim) GetDelayDistribution() (GstNetSimDistribution, error) {
	var value GstNetSimDistribution
	var ok bool
	v, err := e.Element.GetProperty("delay-distribution")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstNetSimDistribution)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstNetSimDistribution")
	}
	return value, nil
}

func (e *Netsim) SetDelayDistribution(value GstNetSimDistribution) error {
	e.Element.SetArg("delay-distribution", string(value))
	return nil
}

// delay-probability (float32)
//
// The Probability a buffer is delayed

func (e *Netsim) GetDelayProbability() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("delay-probability")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Netsim) SetDelayProbability(value float32) error {
	return e.Element.SetProperty("delay-probability", value)
}

// drop-packets (uint)
//
// Drop the next n packets

func (e *Netsim) GetDropPackets() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("drop-packets")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Netsim) SetDropPackets(value uint) error {
	return e.Element.SetProperty("drop-packets", value)
}

// drop-probability (float32)
//
// The Probability a buffer is dropped

func (e *Netsim) GetDropProbability() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("drop-probability")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Netsim) SetDropProbability(value float32) error {
	return e.Element.SetProperty("drop-probability", value)
}

// duplicate-probability (float32)
//
// The Probability a buffer is duplicated

func (e *Netsim) GetDuplicateProbability() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("duplicate-probability")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Netsim) SetDuplicateProbability(value float32) error {
	return e.Element.SetProperty("duplicate-probability", value)
}

// max-bucket-size (int)
//
// The size of the token bucket, related to burstiness resilience.

func (e *Netsim) GetMaxBucketSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-bucket-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Netsim) SetMaxBucketSize(value int) error {
	return e.Element.SetProperty("max-bucket-size", value)
}

// max-delay (int)
//
// The maximum delay (inclusive) in ms to apply to buffers

func (e *Netsim) GetMaxDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Netsim) SetMaxDelay(value int) error {
	return e.Element.SetProperty("max-delay", value)
}

// max-kbps (int)
//
// The maximum number of kilobits to let through per second. Setting this
// property to a positive value enables network congestion simulation using
// a token bucket algorithm. Also see the "max-bucket-size" property,

func (e *Netsim) GetMaxKbps() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("max-kbps")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Netsim) SetMaxKbps(value int) error {
	return e.Element.SetProperty("max-kbps", value)
}

// min-delay (int)
//
// The minimum delay in ms to apply to buffers

func (e *Netsim) GetMinDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("min-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Netsim) SetMinDelay(value int) error {
	return e.Element.SetProperty("min-delay", value)
}

// ----- Constants -----

type GstNetSimDistribution string

const (
	GstNetSimDistributionUniform GstNetSimDistribution = "uniform" // uniform (0) – uniform
	GstNetSimDistributionNormal GstNetSimDistribution = "normal" // normal (1) – normal
	GstNetSimDistributionGamma GstNetSimDistribution = "gamma" // gamma (2) – gamma
)

