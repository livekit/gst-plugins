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
	"github.com/tinyzimmer/go-gst/gst"
)

type DshowvdecDivx5 struct {
	Element *gst.Element
}

func NewDshowvdecDivx5() (*DshowvdecDivx5, error) {
	e, err := gst.NewElement("dshowvdec_divx5")
	if err != nil {
		return nil, err
	}
	return &DshowvdecDivx5{Element: e}, nil
}

func NewDshowvdecDivx5WithName(name string) (*DshowvdecDivx5, error) {
	e, err := gst.NewElementWithName("dshowvdec_divx5", name)
	if err != nil {
		return nil, err
	}
	return &DshowvdecDivx5{Element: e}, nil
}

