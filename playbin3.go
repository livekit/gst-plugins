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

type Playbin3 struct {
	Element *gst.Element
}

func NewPlaybin3() (*Playbin3, error) {
	e, err := gst.NewElement("playbin3")
	if err != nil {
		return nil, err
	}
	return &Playbin3{Element: e}, nil
}

func NewPlaybin3WithName(name string) (*Playbin3, error) {
	e, err := gst.NewElementWithName("playbin3", name)
	if err != nil {
		return nil, err
	}
	return &Playbin3{Element: e}, nil
}

// ----- Properties -----

// audio-filter (GstElement)
//
// the audio filter(s) to apply, if possible

func (e *Playbin3) GetAudioFilter() (interface{}, error) {
	return e.Element.GetProperty("audio-filter")
}

func (e *Playbin3) SetAudioFilter(value interface{}) error {
	return e.Element.SetProperty("audio-filter", value)
}

// audio-sink (GstElement)
//
// Get or set the audio sink to use for audio output. If set to
// NULL, one will be auto-selected. To disable audio entirely, unset
// the AUDIO flag in the flags property.

func (e *Playbin3) GetAudioSink() (interface{}, error) {
	return e.Element.GetProperty("audio-sink")
}

func (e *Playbin3) SetAudioSink(value interface{}) error {
	return e.Element.SetProperty("audio-sink", value)
}

// audio-stream-combiner (GstElement)
//
// Get or set the current audio stream combiner. By default, no
// element is used and the selected stream is used directly.

func (e *Playbin3) GetAudioStreamCombiner() (interface{}, error) {
	return e.Element.GetProperty("audio-stream-combiner")
}

func (e *Playbin3) SetAudioStreamCombiner(value interface{}) error {
	return e.Element.SetProperty("audio-stream-combiner", value)
}

// av-offset (int64)
//
// Control the synchronisation offset between the audio and video streams.
// Positive values make the audio ahead of the video and negative values make
// the audio go behind the video.

func (e *Playbin3) GetAvOffset() (int64, error) {
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

func (e *Playbin3) SetAvOffset(value int64) error {
	return e.Element.SetProperty("av-offset", value)
}

// buffer-duration (int64)
//
// Buffer duration when buffering network streams

func (e *Playbin3) GetBufferDuration() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("buffer-duration")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Playbin3) SetBufferDuration(value int64) error {
	return e.Element.SetProperty("buffer-duration", value)
}

// buffer-size (int)
//
// Buffer size when buffering network streams

func (e *Playbin3) GetBufferSize() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("buffer-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Playbin3) SetBufferSize(value int) error {
	return e.Element.SetProperty("buffer-size", value)
}

// connection-speed (uint64)
//
// Network connection speed in kbps (0 = unknown)

func (e *Playbin3) GetConnectionSpeed() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("connection-speed")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Playbin3) SetConnectionSpeed(value uint64) error {
	return e.Element.SetProperty("connection-speed", value)
}

// current-suburi (string)
//
// The currently playing subtitle uri.

func (e *Playbin3) GetCurrentSuburi() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-suburi")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// current-uri (string)
//
// The currently playing uri.

func (e *Playbin3) GetCurrentUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("current-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// flags (GstPlayFlags)
//
// Control the behaviour of playbin.

func (e *Playbin3) GetFlags() (interface{}, error) {
	return e.Element.GetProperty("flags")
}

func (e *Playbin3) SetFlags(value interface{}) error {
	return e.Element.SetProperty("flags", value)
}

// force-aspect-ratio (bool)
//
// When enabled, scaling will respect original aspect ratio

func (e *Playbin3) GetForceAspectRatio() (bool, error) {
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

func (e *Playbin3) SetForceAspectRatio(value bool) error {
	return e.Element.SetProperty("force-aspect-ratio", value)
}

// instant-uri (bool)
//
// Changes to uri are applied immediately, instead of on EOS or when the
// element is set back to PLAYING.

func (e *Playbin3) GetInstantUri() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("instant-uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Playbin3) SetInstantUri(value bool) error {
	return e.Element.SetProperty("instant-uri", value)
}

// mute (bool)
//
// Mute the audio channel without changing the volume

func (e *Playbin3) GetMute() (bool, error) {
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

func (e *Playbin3) SetMute(value bool) error {
	return e.Element.SetProperty("mute", value)
}

// ring-buffer-max-size (uint64)
//
// The maximum size of the ring buffer in bytes. If set to 0, the ring
// buffer is disabled. Default 0.

func (e *Playbin3) GetRingBufferMaxSize() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("ring-buffer-max-size")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Playbin3) SetRingBufferMaxSize(value uint64) error {
	return e.Element.SetProperty("ring-buffer-max-size", value)
}

// sample (GstSample)
//
// Get the currently rendered or prerolled sample in the video sink.
// The GstCaps in the sample will describe the format of the buffer.

func (e *Playbin3) GetSample() (*gst.Sample, error) {
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

// subtitle-encoding (string)
//
// Encoding to assume if input subtitles are not in UTF-8 encoding. If not set, the GST_SUBTITLE_ENCODING environment variable will be checked for an encoding to use. If that is not set either, ISO-8859-15 will be assumed.

func (e *Playbin3) GetSubtitleEncoding() (string, error) {
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

func (e *Playbin3) SetSubtitleEncoding(value string) error {
	return e.Element.SetProperty("subtitle-encoding", value)
}

// subtitle-font-desc (string)
//
// Pango font description of font to be used for subtitle rendering

func (e *Playbin3) GetSubtitleFontDesc() (string, error) {
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

func (e *Playbin3) SetSubtitleFontDesc(value string) error {
	return e.Element.SetProperty("subtitle-font-desc", value)
}

// suburi (string)
//
// Set the next subtitle URI that playbin will play. This property can be
// set from the about-to-finish signal to queue the next subtitle media file.

func (e *Playbin3) GetSuburi() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("suburi")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Playbin3) SetSuburi(value string) error {
	return e.Element.SetProperty("suburi", value)
}

// text-offset (int64)
//
// Control the synchronisation offset between the text and video streams.
// Positive values make the text ahead of the video and negative values make
// the text go behind the video.

func (e *Playbin3) GetTextOffset() (int64, error) {
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

func (e *Playbin3) SetTextOffset(value int64) error {
	return e.Element.SetProperty("text-offset", value)
}

// text-sink (GstElement)
//
// the text output element to use (NULL = default subtitleoverlay)

func (e *Playbin3) GetTextSink() (interface{}, error) {
	return e.Element.GetProperty("text-sink")
}

func (e *Playbin3) SetTextSink(value interface{}) error {
	return e.Element.SetProperty("text-sink", value)
}

// text-stream-combiner (GstElement)
//
// Get or set the current text stream combiner. By default, no
// element is used and the selected stream is used directly.

func (e *Playbin3) GetTextStreamCombiner() (interface{}, error) {
	return e.Element.GetProperty("text-stream-combiner")
}

func (e *Playbin3) SetTextStreamCombiner(value interface{}) error {
	return e.Element.SetProperty("text-stream-combiner", value)
}

// uri (string)
//
// Set the next URI that playbin will play. This property can be set from the
// about-to-finish signal to queue the next media file.

func (e *Playbin3) GetUri() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("uri")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Playbin3) SetUri(value string) error {
	return e.Element.SetProperty("uri", value)
}

// video-filter (GstElement)
//
// the video filter(s) to apply, if possible

func (e *Playbin3) GetVideoFilter() (interface{}, error) {
	return e.Element.GetProperty("video-filter")
}

func (e *Playbin3) SetVideoFilter(value interface{}) error {
	return e.Element.SetProperty("video-filter", value)
}

// video-multiview-flags (GstVideoMultiviewFlags)
//
// Override details of the multiview frame layout

func (e *Playbin3) GetVideoMultiviewFlags() (interface{}, error) {
	return e.Element.GetProperty("video-multiview-flags")
}

func (e *Playbin3) SetVideoMultiviewFlags(value interface{}) error {
	return e.Element.SetProperty("video-multiview-flags", value)
}

// video-multiview-mode (GstVideoMultiviewFramePacking)
//
// Re-interpret a video stream as one of several frame-packed stereoscopic modes.

func (e *Playbin3) GetVideoMultiviewMode() (interface{}, error) {
	return e.Element.GetProperty("video-multiview-mode")
}

func (e *Playbin3) SetVideoMultiviewMode(value interface{}) error {
	return e.Element.SetProperty("video-multiview-mode", value)
}

// video-sink (GstElement)
//
// Get or set the video sink to use for video output. If set to
// NULL, one will be auto-selected. To disable video entirely, unset
// the VIDEO flag in the flags property.

func (e *Playbin3) GetVideoSink() (interface{}, error) {
	return e.Element.GetProperty("video-sink")
}

func (e *Playbin3) SetVideoSink(value interface{}) error {
	return e.Element.SetProperty("video-sink", value)
}

// video-stream-combiner (GstElement)
//
// Get or set the current video stream combiner. By default, no
// element is used and the selected stream is used directly.

func (e *Playbin3) GetVideoStreamCombiner() (interface{}, error) {
	return e.Element.GetProperty("video-stream-combiner")
}

func (e *Playbin3) SetVideoStreamCombiner(value interface{}) error {
	return e.Element.SetProperty("video-stream-combiner", value)
}

// vis-plugin (GstElement)
//
// the visualization element to use (NULL = default)

func (e *Playbin3) GetVisPlugin() (interface{}, error) {
	return e.Element.GetProperty("vis-plugin")
}

func (e *Playbin3) SetVisPlugin(value interface{}) error {
	return e.Element.SetProperty("vis-plugin", value)
}

// volume (float64)
//
// Get or set the current audio stream volume. 1.0 means 100%,
// 0.0 means mute. This uses a linear volume scale.

func (e *Playbin3) GetVolume() (float64, error) {
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

func (e *Playbin3) SetVolume(value float64) error {
	return e.Element.SetProperty("volume", value)
}

