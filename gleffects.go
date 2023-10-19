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

type Gleffects struct {
	*base.GstBaseTransform
}

func NewGleffects() (*Gleffects, error) {
	e, err := gst.NewElement("gleffects")
	if err != nil {
		return nil, err
	}
	return &Gleffects{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewGleffectsWithName(name string) (*Gleffects, error) {
	e, err := gst.NewElementWithName("gleffects", name)
	if err != nil {
		return nil, err
	}
	return &Gleffects{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// ----- Properties -----

// effect (GstGleffectsEffect)
//
// Select which effect apply to GL video texture

func (e *Gleffects) GetEffect() (GstGleffectsEffect, error) {
	var value GstGleffectsEffect
	var ok bool
	v, err := e.Element.GetProperty("effect")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstGleffectsEffect)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstGleffectsEffect")
	}
	return value, nil
}

func (e *Gleffects) SetEffect(value GstGleffectsEffect) error {
	e.Element.SetArg("effect", string(value))
	return nil
}

// hswap (bool)
//
// Switch video texture left to right, useful with webcams

func (e *Gleffects) GetHswap() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("hswap")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Gleffects) SetHswap(value bool) error {
	return e.Element.SetProperty("hswap", value)
}

// invert (bool)
//
// Invert colors to get dark edges on bright background when using sobel effect

func (e *Gleffects) GetInvert() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("invert")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Gleffects) SetInvert(value bool) error {
	return e.Element.SetProperty("invert", value)
}

// ----- Constants -----

type GstGleffectsEffect string

const (
	GstGleffectsEffectIdentity GstGleffectsEffect = "identity" // identity (0) – Do nothing Effect
	GstGleffectsEffectMirror GstGleffectsEffect = "mirror" // mirror (1) – Mirror Effect
	GstGleffectsEffectSqueeze GstGleffectsEffect = "squeeze" // squeeze (2) – Squeeze Effect
	GstGleffectsEffectStretch GstGleffectsEffect = "stretch" // stretch (3) – Stretch Effect
	GstGleffectsEffectTunnel GstGleffectsEffect = "tunnel" // tunnel (4) – Light Tunnel Effect
	GstGleffectsEffectFisheye GstGleffectsEffect = "fisheye" // fisheye (5) – FishEye Effect
	GstGleffectsEffectTwirl GstGleffectsEffect = "twirl" // twirl (6) – Twirl Effect
	GstGleffectsEffectBulge GstGleffectsEffect = "bulge" // bulge (7) – Bulge Effect
	GstGleffectsEffectSquare GstGleffectsEffect = "square" // square (8) – Square Effect
	GstGleffectsEffectHeat GstGleffectsEffect = "heat" // heat (9) – Heat Signature Effect
	GstGleffectsEffectSepia GstGleffectsEffect = "sepia" // sepia (10) – Sepia Toning Effect
	GstGleffectsEffectXpro GstGleffectsEffect = "xpro" // xpro (11) – Cross Processing Effect
	GstGleffectsEffectLumaxpro GstGleffectsEffect = "lumaxpro" // lumaxpro (12) – Luma Cross Processing Effect
	GstGleffectsEffectXray GstGleffectsEffect = "xray" // xray (13) – Glowing negative effect
	GstGleffectsEffectSin GstGleffectsEffect = "sin" // sin (14) – All Grey but Red Effect
	GstGleffectsEffectGlow GstGleffectsEffect = "glow" // glow (15) – Glow Lighting Effect
	GstGleffectsEffectSobel GstGleffectsEffect = "sobel" // sobel (16) – Sobel edge detection Effect
	GstGleffectsEffectBlur GstGleffectsEffect = "blur" // blur (17) – Blur with 9x9 separable convolution Effect
	GstGleffectsEffectLaplacian GstGleffectsEffect = "laplacian" // laplacian (18) – Laplacian Convolution Demo Effect
)

