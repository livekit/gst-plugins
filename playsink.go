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

type Playsink struct {
	Element *gst.Element
}

func NewPlaysink() (*Playsink, error) {
	e, err := gst.NewElement("playsink")
	if err != nil {
		return nil, err
	}
	return &Playsink{Element: e}, nil
}

func NewPlaysinkWithName(name string) (*Playsink, error) {
	e, err := gst.NewElementWithName("playsink", name)
	if err != nil {
		return nil, err
	}
	return &Playsink{Element: e}, nil
}

// ----- Properties -----

// audio-filter (GstElement)
//
// Set the audio filter element/bin to use. Will apply on a best-effort basis
// unless GST_PLAY_FLAG_FORCE_FILTERS is set. playsink must be in
// GST_STATE_NULL

func (e *Playsink) GetAudioFilter() (interface{}, error) {
	return e.Element.GetProperty("audio-filter")
}

func (e *Playsink) SetAudioFilter(value interface{}) error {
	return e.Element.SetProperty("audio-filter", value)
}

// audio-sink (GstElement)
//
// Set the used audio sink element. NULL will use the default sink. playsink
// must be in GST_STATE_NULL

func (e *Playsink) GetAudioSink() (interface{}, error) {
	return e.Element.GetProperty("audio-sink")
}

func (e *Playsink) SetAudioSink(value interface{}) error {
	return e.Element.SetProperty("audio-sink", value)
}

// av-offset (int64)
//
// Control the synchronisation offset between the audio and video streams.
// Positive values make the audio ahead of the video and negative values make
// the audio go behind the video.

func (e *Playsink) GetAvOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("av-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Playsink) SetAvOffset(value int64) error {
	return e.Element.SetProperty("av-offset", value)
}

// flags (GstPlayFlags)
//
// Control the behaviour of playsink.

func (e *Playsink) GetFlags() (GstPlayFlags, error) {
	var value GstPlayFlags
	var ok bool
	v, err := e.Element.GetProperty("flags")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstPlayFlags)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstPlayFlags")
	}
	return value, nil
}

func (e *Playsink) SetFlags(value GstPlayFlags) error {
	e.Element.SetArg("flags", string(value))
	return nil
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Playsink) GetForceAspectRatio() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("force-aspect-ratio")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Playsink) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// mute (bool)
//
// Mute the audio channel without changing the volume

func (e *Playsink) GetMute() (bool, error) {
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

func (e *Playsink) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// sample (GstSample)
//
// Get the currently rendered or prerolled sample in the video sink.
// The GstCaps in the sample will describe the format of the buffer.

func (e *Playsink) GetSample() (*gst.Sample, error) {
	var value *gst.Sample
	var ok bool
	v, err := e.Element.GetProperty("sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Sample)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Sample")
	}
	return value, nil
}

// send-event-mode (GstPlaySinkSendEventMode)
//
// How to send events received in send_event function

func (e *Playsink) GetSendEventMode() (GstPlaySinkSendEventMode, error) {
	var value GstPlaySinkSendEventMode
	var ok bool
	v, err := e.Element.GetProperty("send-event-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstPlaySinkSendEventMode)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstPlaySinkSendEventMode")
	}
	return value, nil
}

func (e *Playsink) SetSendEventMode(value GstPlaySinkSendEventMode) error {
	e.Element.SetArg("send-event-mode", string(value))
	return nil
}

// subtitle-encoding (string)
//
// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.

func (e *Playsink) GetSubtitleEncoding() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-encoding")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Playsink) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

// subtitle-font-desc (string)
//
// Pango font description of font to be used for subtitle rendering

func (e *Playsink) GetSubtitleFontDesc() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("subtitle-font-desc")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Playsink) SetSubtitleFontDesc(value string) error {
	return e.Element.SetProperty("subtitle-font-desc", value)
}

// text-offset (int64)
//
// Control the synchronisation offset between the text and video streams.
// Positive values make the text ahead of the video and negative values make
// the text go behind the video.

func (e *Playsink) GetTextOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("text-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Playsink) SetTextOffset(value int64) error {
	return e.Element.SetProperty("text-offset", value)
}

// text-sink (GstElement)
//
// Set the used text sink element. NULL will use the default sink. playsink
// must be in GST_STATE_NULL

func (e *Playsink) GetTextSink() (interface{}, error) {
	return e.Element.GetProperty("text-sink")
}

func (e *Playsink) SetTextSink(value interface{}) error {
	return e.Element.SetProperty("text-sink", value)
}

// video-filter (GstElement)
//
// Set the video filter element/bin to use. Will apply on a best-effort basis
// unless GST_PLAY_FLAG_FORCE_FILTERS is set. playsink must be in
// GST_STATE_NULL

func (e *Playsink) GetVideoFilter() (interface{}, error) {
	return e.Element.GetProperty("video-filter")
}

func (e *Playsink) SetVideoFilter(value interface{}) error {
	return e.Element.SetProperty("video-filter", value)
}

// video-sink (GstElement)
//
// Set the used video sink element. NULL will use the default sink. playsink
// must be in GST_STATE_NULL

func (e *Playsink) GetVideoSink() (interface{}, error) {
	return e.Element.GetProperty("video-sink")
}

func (e *Playsink) SetVideoSink(value interface{}) error {
	return e.Element.SetProperty("video-sink", value)
}

// vis-plugin (GstElement)
//
// the visualization element to use (NULL = default)

func (e *Playsink) GetVisPlugin() (interface{}, error) {
	return e.Element.GetProperty("vis-plugin")
}

func (e *Playsink) SetVisPlugin(value interface{}) error {
	return e.Element.SetProperty("vis-plugin", value)
}

// volume (float64)
//
// Get or set the current audio stream volume. 1.0 means 100%,
// 0.0 means mute. This uses a linear volume scale.

func (e *Playsink) GetVolume() (float64, error) {
	var value float64
	var ok bool
	v, err := e.Element.GetProperty("volume")
	if err != nil {
		return value, err
	}
	value, ok = v.(float64)
	if !ok {
		return value, fmt.Errorf("could not cast value to float64")
	}
	return value, nil
}

func (e *Playsink) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

// ----- Constants -----

type GstPlayFlags string

const (
	GstPlayFlagsVideo GstPlayFlags = "video" // video (0x00000001) – Render the video stream
	GstPlayFlagsAudio GstPlayFlags = "audio" // audio (0x00000002) – Render the audio stream
	GstPlayFlagsText GstPlayFlags = "text" // text (0x00000004) – Render subtitles
	GstPlayFlagsVis GstPlayFlags = "vis" // vis (0x00000008) – Render visualisation when no video is present
	GstPlayFlagsSoftVolume GstPlayFlags = "soft-volume" // soft-volume (0x00000010) – Use software volume
	GstPlayFlagsNativeAudio GstPlayFlags = "native-audio" // native-audio (0x00000020) – Only use native audio formats
	GstPlayFlagsNativeVideo GstPlayFlags = "native-video" // native-video (0x00000040) – Only use native video formats
	GstPlayFlagsDownload GstPlayFlags = "download" // download (0x00000080) – Attempt progressive download buffering
	GstPlayFlagsBuffering GstPlayFlags = "buffering" // buffering (0x00000100) – Buffer demuxed/parsed data
	GstPlayFlagsDeinterlace GstPlayFlags = "deinterlace" // deinterlace (0x00000200) – Deinterlace video if necessary
	GstPlayFlagsSoftColorbalance GstPlayFlags = "soft-colorbalance" // soft-colorbalance (0x00000400) – Use software color balance
	GstPlayFlagsForceFilters GstPlayFlags = "force-filters" // force-filters (0x00000800) – Force audio/video filter(s) to be applied
	GstPlayFlagsForceSwDecoders GstPlayFlags = "force-sw-decoders" // force-sw-decoders (0x00001000) – Force only software-based decoders (no effect for playbin3)
)

type GstPlaySinkSendEventMode string

const (
	GstPlaySinkSendEventModeDefault GstPlaySinkSendEventMode = "default" // default (0) – Default GstBin's send_event handling (default)
	GstPlaySinkSendEventModeFirst GstPlaySinkSendEventMode = "first" // first (1) – Sends the event to sinks until the first one handles it
)

