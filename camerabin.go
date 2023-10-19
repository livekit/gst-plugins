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

type Camerabin struct {
	Element *gst.Element
}

func NewCamerabin() (*Camerabin, error) {
	e, err := gst.NewElement("camerabin")
	if err != nil {
		return nil, err
	}
	return &Camerabin{Element: e}, nil
}

func NewCamerabinWithName(name string) (*Camerabin, error) {
	e, err := gst.NewElementWithName("camerabin", name)
	if err != nil {
		return nil, err
	}
	return &Camerabin{Element: e}, nil
}

// ----- Properties -----

// audio-capture-caps (GstCaps)
//
// Format to capture audio for video recording represented as GstCaps

func (e *Camerabin) GetAudioCaptureCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("audio-capture-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Camerabin) SetAudioCaptureCaps(value *gst.Caps) error {
	return e.Element.SetProperty("audio-capture-caps", value)
}

// audio-capture-supported-caps (GstCaps)
//
// Formats supported for capturing audio represented as GstCaps

func (e *Camerabin) GetAudioCaptureSupportedCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("audio-capture-supported-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// audio-filter (GstElement)
//
// The element that will process captured audio buffers when recording. (Should be set on NULL state)

func (e *Camerabin) GetAudioFilter() (interface{}, error) {
	return e.Element.GetProperty("audio-filter")
}

func (e *Camerabin) SetAudioFilter(value interface{}) error {
	return e.Element.SetProperty("audio-filter", value)
}

// audio-source (GstElement)
//
// The audio source element to be used on video recordings. It is only taken into use on the next null to ready transition

func (e *Camerabin) GetAudioSource() (interface{}, error) {
	return e.Element.GetProperty("audio-source")
}

func (e *Camerabin) SetAudioSource(value interface{}) error {
	return e.Element.SetProperty("audio-source", value)
}

// camera-source (GstElement)
//
// The camera source element to be used. It is only taken into use on the next null to ready transition

func (e *Camerabin) GetCameraSource() (interface{}, error) {
	return e.Element.GetProperty("camera-source")
}

func (e *Camerabin) SetCameraSource(value interface{}) error {
	return e.Element.SetProperty("camera-source", value)
}

// flags (GstCamFlags)
//
// Control the behaviour of encodebin.

func (e *Camerabin) GetFlags() (GstCamFlags, error) {
	var value GstCamFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCamFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCamFlags")
	}
	return value, nil
}

func (e *Camerabin) SetFlags(value GstCamFlags) error {
	e.Element.SetArg("flags", string(value))
	return nil
}

// idle (bool)
//
// If camerabin2 is idle (not doing captures).

func (e *Camerabin) GetIdle() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("idle")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// image-capture-caps (GstCaps)
//
// Caps for image capture

func (e *Camerabin) GetImageCaptureCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("image-capture-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Camerabin) SetImageCaptureCaps(value *gst.Caps) error {
	return e.Element.SetProperty("image-capture-caps", value)
}

// image-capture-supported-caps (GstCaps)
//
// Formats supported for capturing images represented as GstCaps

func (e *Camerabin) GetImageCaptureSupportedCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("image-capture-supported-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// image-filter (GstElement)
//
// The element that will process captured image frames. (Should be set on NULL state)

func (e *Camerabin) GetImageFilter() (interface{}, error) {
	return e.Element.GetProperty("image-filter")
}

func (e *Camerabin) SetImageFilter(value interface{}) error {
	return e.Element.SetProperty("image-filter", value)
}

// image-profile (GstEncodingProfile)
//
// The GstEncodingProfile to use for image captures.

func (e *Camerabin) GetImageProfile() (interface{}, error) {
	return e.Element.GetProperty("image-profile")
}

func (e *Camerabin) SetImageProfile(value interface{}) error {
	return e.Element.SetProperty("image-profile", value)
}

// location (string)
//
// Location to save the captured files. A %%d might be used on thefilename as a placeholder for a numeric index of the capture.Default is cap_%%d

func (e *Camerabin) GetLocation() (string, error) {
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

func (e *Camerabin) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// max-zoom (float32)
//
// Digital zoom factor (e.g. 1.5 means 1.5x)

func (e *Camerabin) GetMaxZoom() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("max-zoom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

// mode (GstCameraBin2Mode)
//
// The capture mode (still image capture or video recording)

func (e *Camerabin) GetMode() (interface{}, error) {
	return e.Element.GetProperty("mode")
}

func (e *Camerabin) SetMode(value interface{}) error {
	return e.Element.SetProperty("mode", value)
}

// mute (bool)
//
// If the audio recording should be muted. Note that this still saves audio data to the resulting file, but they are silent. Use a video-profile without audio to disable audio completely

func (e *Camerabin) GetMute() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("mute")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Camerabin) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// post-previews (bool)
//
// If capture preview images should be posted to the bus

func (e *Camerabin) GetPostPreviews() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("post-previews")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Camerabin) SetPostPreviews(value bool) error {
	return e.Element.SetProperty("post-previews", value)
}

// preview-caps (GstCaps)
//
// The caps of the preview image to be posted

func (e *Camerabin) GetPreviewCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("preview-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Camerabin) SetPreviewCaps(value *gst.Caps) error {
	return e.Element.SetProperty("preview-caps", value)
}

// preview-filter (GstElement)
//
// The element that will process preview buffers. (Should be set on NULL state)

func (e *Camerabin) GetPreviewFilter() (interface{}, error) {
	return e.Element.GetProperty("preview-filter")
}

func (e *Camerabin) SetPreviewFilter(value interface{}) error {
	return e.Element.SetProperty("preview-filter", value)
}

// video-capture-caps (GstCaps)
//
// Caps for video capture

func (e *Camerabin) GetVideoCaptureCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("video-capture-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Camerabin) SetVideoCaptureCaps(value *gst.Caps) error {
	return e.Element.SetProperty("video-capture-caps", value)
}

// video-capture-supported-caps (GstCaps)
//
// Formats supported for capturing videos represented as GstCaps

func (e *Camerabin) GetVideoCaptureSupportedCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("video-capture-supported-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// video-filter (GstElement)
//
// The element that will process captured video frames. (Should be set on NULL state)

func (e *Camerabin) GetVideoFilter() (interface{}, error) {
	return e.Element.GetProperty("video-filter")
}

func (e *Camerabin) SetVideoFilter(value interface{}) error {
	return e.Element.SetProperty("video-filter", value)
}

// video-profile (GstEncodingProfile)
//
// The GstEncodingProfile to use for video recording. Audio is enabled when this profile supports audio.

func (e *Camerabin) GetVideoProfile() (interface{}, error) {
	return e.Element.GetProperty("video-profile")
}

func (e *Camerabin) SetVideoProfile(value interface{}) error {
	return e.Element.SetProperty("video-profile", value)
}

// viewfinder-caps (GstCaps)
//
// Restricts the caps that can be used on the viewfinder

func (e *Camerabin) GetViewfinderCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("viewfinder-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

func (e *Camerabin) SetViewfinderCaps(value *gst.Caps) error {
	return e.Element.SetProperty("viewfinder-caps", value)
}

// viewfinder-filter (GstElement)
//
// The element that will process frames going to the viewfinder. (Should be set on NULL state)

func (e *Camerabin) GetViewfinderFilter() (interface{}, error) {
	return e.Element.GetProperty("viewfinder-filter")
}

func (e *Camerabin) SetViewfinderFilter(value interface{}) error {
	return e.Element.SetProperty("viewfinder-filter", value)
}

// viewfinder-sink (GstElement)
//
// The video sink of the viewfinder. It is only taken into use on the next null to ready transition

func (e *Camerabin) GetViewfinderSink() (interface{}, error) {
	return e.Element.GetProperty("viewfinder-sink")
}

func (e *Camerabin) SetViewfinderSink(value interface{}) error {
	return e.Element.SetProperty("viewfinder-sink", value)
}

// viewfinder-supported-caps (GstCaps)
//
// The caps that the camera source can produce on the viewfinder pad

func (e *Camerabin) GetViewfinderSupportedCaps() (*gst.Caps, error) {
	var value *gst.Caps
	var ok bool
	v, err := e.Element.GetProperty("viewfinder-supported-caps")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Caps)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Caps")
	}
	return value, nil
}

// zoom (float32)
//
// Digital zoom factor (e.g. 1.5 means 1.5x)

func (e *Camerabin) GetZoom() (float32, error) {
	var value float32
	var ok bool
	v, err := e.Element.GetProperty("zoom")
	if err != nil {
		return value, err
	}
	value, ok = v.(float32)
	if !ok {
		return value, fmt.Errorf("could not cast value to float32")
	}
	return value, nil
}

func (e *Camerabin) SetZoom(value float32) error {
	return e.Element.SetProperty("zoom", value)
}

// ----- Constants -----

type GstCamFlags string

const (
	GstCamFlagsNoAudioConversion GstCamFlags = "no-audio-conversion" // no-audio-conversion (0x00000001) – Do not use audio conversion elements
	GstCamFlagsNoVideoConversion GstCamFlags = "no-video-conversion" // no-video-conversion (0x00000002) – Do not use video conversion elements
	GstCamFlagsNoViewfinderConversion GstCamFlags = "no-viewfinder-conversion" // no-viewfinder-conversion (0x00000004) – Do not use viewfinder conversion elements
	GstCamFlagsNoImageConversion GstCamFlags = "no-image-conversion" // no-image-conversion (0x00000008) – Do not use image conversion elements
)

