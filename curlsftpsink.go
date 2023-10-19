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

type Curlsftpsink struct {
	*base.GstBaseSink
}

func NewCurlsftpsink() (*Curlsftpsink, error) {
	e, err := gst.NewElement("curlsftpsink")
	if err != nil {
		return nil, err
	}
	return &Curlsftpsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewCurlsftpsinkWithName(name string) (*Curlsftpsink, error) {
	e, err := gst.NewElementWithName("curlsftpsink", name)
	if err != nil {
		return nil, err
	}
	return &Curlsftpsink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// ----- Properties -----

// async (bool)
//
// Go asynchronously to PAUSED

func (e *Curlsftpsink) GetAsync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("async")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Curlsftpsink) SetAsync(value bool) error {
	return e.Element.SetProperty("async", value)
}

// blocksize (uint)
//
// Size in bytes to pull per buffer (0 = default)

func (e *Curlsftpsink) GetBlocksize() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("blocksize")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

func (e *Curlsftpsink) SetBlocksize(value uint) error {
	return e.Element.SetProperty("blocksize", value)
}

// create-dirs (bool)
//
// Attempt to create missing directories

func (e *Curlsftpsink) GetCreateDirs() (bool, error) {
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

func (e *Curlsftpsink) SetCreateDirs(value bool) error {
	return e.Element.SetProperty("create-dirs", value)
}

// enable-last-sample (bool)
//
// Enable the last-sample property

func (e *Curlsftpsink) GetEnableLastSample() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("enable-last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Curlsftpsink) SetEnableLastSample(value bool) error {
	return e.Element.SetProperty("enable-last-sample", value)
}

// file-name (string)
//
// The base file name for the uploaded images

func (e *Curlsftpsink) GetFileName() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("file-name")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsftpsink) SetFileName(value string) error {
	return e.Element.SetProperty("file-name", value)
}

// last-sample (GstSample)
//
// The last sample received in the sink

func (e *Curlsftpsink) GetLastSample() (*gst.Sample, error) {
	var value *gst.Sample
	var ok bool
	v, err := e.Element.GetProperty("last-sample")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Sample)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Sample")
	}
	return value, nil
}

// location (string)
//
// URI location to write to

func (e *Curlsftpsink) GetLocation() (string, error) {
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

func (e *Curlsftpsink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

// max-bitrate (uint64)
//
// The maximum bits per second to render (0 = disabled)

func (e *Curlsftpsink) GetMaxBitrate() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("max-bitrate")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Curlsftpsink) SetMaxBitrate(value uint64) error {
	return e.Element.SetProperty("max-bitrate", value)
}

// max-lateness (int64)
//
// Maximum number of nanoseconds that a buffer can be late before it is dropped (-1 unlimited)

func (e *Curlsftpsink) GetMaxLateness() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("max-lateness")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Curlsftpsink) SetMaxLateness(value int64) error {
	return e.Element.SetProperty("max-lateness", value)
}

// passwd (string)
//
// User password to use for server authentication

func (e *Curlsftpsink) GetPasswd() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("passwd")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsftpsink) SetPasswd(value string) error {
	return e.Element.SetProperty("passwd", value)
}

// processing-deadline (uint64)
//
// Maximum processing time for a buffer in nanoseconds

func (e *Curlsftpsink) GetProcessingDeadline() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("processing-deadline")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Curlsftpsink) SetProcessingDeadline(value uint64) error {
	return e.Element.SetProperty("processing-deadline", value)
}

// qos (bool)
//
// Generate Quality-of-Service events upstream

func (e *Curlsftpsink) GetQos() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("qos")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Curlsftpsink) SetQos(value bool) error {
	return e.Element.SetProperty("qos", value)
}

// qos-dscp (int)
//
// Quality of Service, differentiated services code point (0 default)

func (e *Curlsftpsink) GetQosDscp() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("qos-dscp")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Curlsftpsink) SetQosDscp(value int) error {
	return e.Element.SetProperty("qos-dscp", value)
}

// render-delay (uint64)
//
// Additional render delay of the sink in nanoseconds

func (e *Curlsftpsink) GetRenderDelay() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("render-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Curlsftpsink) SetRenderDelay(value uint64) error {
	return e.Element.SetProperty("render-delay", value)
}

// ssh-accept-unknownhost (bool)
//
// Accept an unknown remote public host key

func (e *Curlsftpsink) GetSshAcceptUnknownhost() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssh-accept-unknownhost")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Curlsftpsink) SetSshAcceptUnknownhost(value bool) error {
	return e.Element.SetProperty("ssh-accept-unknownhost", value)
}

// ssh-auth-type (GstCurlSshAuthType)
//
// SSH authentication method to authenticate on the SSH/SFTP server

func (e *Curlsftpsink) GetSshAuthType() (GstCurlSshAuthType, error) {
	var value GstCurlSshAuthType
	var ok bool
	v, err := e.Element.GetProperty("ssh-auth-type")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstCurlSshAuthType)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstCurlSshAuthType")
	}
	return value, nil
}

func (e *Curlsftpsink) SetSshAuthType(value GstCurlSshAuthType) error {
	e.Element.SetArg("ssh-auth-type", string(value))
	return nil
}

// ssh-host-pubkey-md5 (string)
//
// MD5 checksum (32 hexadecimal digits, case-insensitive) of the remote host's public key

func (e *Curlsftpsink) GetSshHostPubkeyMd5() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-host-pubkey-md5")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsftpsink) SetSshHostPubkeyMd5(value string) error {
	return e.Element.SetProperty("ssh-host-pubkey-md5", value)
}

// ssh-key-passphrase (string)
//
// The passphrase used to protect the SSH private key file

func (e *Curlsftpsink) GetSshKeyPassphrase() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-key-passphrase")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsftpsink) SetSshKeyPassphrase(value string) error {
	return e.Element.SetProperty("ssh-key-passphrase", value)
}

// ssh-knownhosts (string)
//
// The complete path & filename of the SSH 'known_hosts' file

func (e *Curlsftpsink) GetSshKnownhosts() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-knownhosts")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsftpsink) SetSshKnownhosts(value string) error {
	return e.Element.SetProperty("ssh-knownhosts", value)
}

// ssh-priv-keyfile (string)
//
// The complete path & filename of the SSH private key file

func (e *Curlsftpsink) GetSshPrivKeyfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-priv-keyfile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsftpsink) SetSshPrivKeyfile(value string) error {
	return e.Element.SetProperty("ssh-priv-keyfile", value)
}

// ssh-pub-keyfile (string)
//
// The complete path & filename of the SSH public key file

func (e *Curlsftpsink) GetSshPubKeyfile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssh-pub-keyfile")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsftpsink) SetSshPubKeyfile(value string) error {
	return e.Element.SetProperty("ssh-pub-keyfile", value)
}

// stats (GstStructure)
//
// Sink Statistics

func (e *Curlsftpsink) GetStats() (interface{}, error) {
	return e.Element.GetProperty("stats")
}

// sync (bool)
//
// Sync on the clock

func (e *Curlsftpsink) GetSync() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("sync")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

func (e *Curlsftpsink) SetSync(value bool) error {
	return e.Element.SetProperty("sync", value)
}

// throttle-time (uint64)
//
// The time to keep between rendered buffers (0 = disabled)

func (e *Curlsftpsink) GetThrottleTime() (uint64, error) {
	var value uint64
	var ok bool
	v, err := e.Element.GetProperty("throttle-time")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint64)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint64")
	}
	return value, nil
}

func (e *Curlsftpsink) SetThrottleTime(value uint64) error {
	return e.Element.SetProperty("throttle-time", value)
}

// timeout (int)
//
// Number of seconds waiting to write before timeout

func (e *Curlsftpsink) GetTimeout() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

func (e *Curlsftpsink) SetTimeout(value int) error {
	return e.Element.SetProperty("timeout", value)
}

// ts-offset (int64)
//
// Timestamp offset in nanoseconds

func (e *Curlsftpsink) GetTsOffset() (int64, error) {
	var value int64
	var ok bool
	v, err := e.Element.GetProperty("ts-offset")
	if err != nil {
		return value, err
	}
	value, ok = v.(int64)
	if !ok {
		return value, fmt.Errorf("could not cast value to int64")
	}
	return value, nil
}

func (e *Curlsftpsink) SetTsOffset(value int64) error {
	return e.Element.SetProperty("ts-offset", value)
}

// user (string)
//
// User name to use for server authentication

func (e *Curlsftpsink) GetUser() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

func (e *Curlsftpsink) SetUser(value string) error {
	return e.Element.SetProperty("user", value)
}

// ----- Constants -----

type GstCurlSshAuthType string

const (
	GstCurlSshAuthTypeNone GstCurlSshAuthType = "none" // none (0) – Not allowed
	GstCurlSshAuthTypePubkey GstCurlSshAuthType = "pubkey" // pubkey (1) – Public/private key files
	GstCurlSshAuthTypePassword GstCurlSshAuthType = "password" // password (2) – Password authentication
)

