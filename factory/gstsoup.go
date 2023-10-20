// source: gst-plugins-good

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// Sends streams to HTTP server via PUT
type GstSoupHttpClientSink struct {
	*base.GstBaseSink
}

func NewSoupHttpClientSink() (*GstSoupHttpClientSink, error) {
	e, err := gst.NewElement("souphttpclientsink")
	if err != nil {
		return nil, err
	}
	return &GstSoupHttpClientSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

func NewSoupHttpClientSinkWithName(name string) (*GstSoupHttpClientSink, error) {
	e, err := gst.NewElementWithName("souphttpclientsink", name)
	if err != nil {
		return nil, err
	}
	return &GstSoupHttpClientSink{GstBaseSink: &base.GstBaseSink{Element: e}}, nil
}

// HTTP proxy server URI

func (e *GstSoupHttpClientSink) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

func (e *GstSoupHttpClientSink) GetProxy() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Maximum number of retries, zero to disable, -1 to retry forever
// Default: 0, Min: -1, Max: 2147483647
func (e *GstSoupHttpClientSink) SetRetries(value int) error {
	return e.Element.SetProperty("retries", value)
}

func (e *GstSoupHttpClientSink) GetRetries() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("retries")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// Delay in seconds between retries after a failure
// Default: 5, Min: 1, Max: 2147483647
func (e *GstSoupHttpClientSink) SetRetryDelay(value int) error {
	return e.Element.SetProperty("retry-delay", value)
}

func (e *GstSoupHttpClientSink) GetRetryDelay() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("retry-delay")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// SoupSession object to use for communication

func (e *GstSoupHttpClientSink) SetSession(value interface{}) error {
	return e.Element.SetProperty("session", value)
}

func (e *GstSoupHttpClientSink) GetSession() (interface{}, error) {
	return e.Element.GetProperty("session")
}

// Automatically follow HTTP redirects (HTTP Status Code 3xx)
// Default: true
func (e *GstSoupHttpClientSink) SetAutomaticRedirect(value bool) error {
	return e.Element.SetProperty("automatic-redirect", value)
}

func (e *GstSoupHttpClientSink) GetAutomaticRedirect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("automatic-redirect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// HTTP request cookies

func (e *GstSoupHttpClientSink) SetCookies(value interface{}) error {
	return e.Element.SetProperty("cookies", value)
}

func (e *GstSoupHttpClientSink) GetCookies() (interface{}, error) {
	return e.Element.GetProperty("cookies")
}

// Set log level for soup's HTTP session log
// Default: none (0)
func (e *GstSoupHttpClientSink) SetHttpLogLevel(value interface{}) error {
	return e.Element.SetProperty("http-log-level", value)
}

func (e *GstSoupHttpClientSink) GetHttpLogLevel() (interface{}, error) {
	return e.Element.GetProperty("http-log-level")
}

// URI to send to
// Default: NULL
func (e *GstSoupHttpClientSink) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstSoupHttpClientSink) GetLocation() (string, error) {
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

// Value of the User-Agent HTTP request header field
// Default: GStreamer souphttpclientsink
func (e *GstSoupHttpClientSink) SetUserAgent(value string) error {
	return e.Element.SetProperty("user-agent", value)
}

func (e *GstSoupHttpClientSink) GetUserAgent() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-agent")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// user password for authentication
// Default: NULL
func (e *GstSoupHttpClientSink) SetUserPw(value string) error {
	return e.Element.SetProperty("user-pw", value)
}

func (e *GstSoupHttpClientSink) GetUserPw() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-pw")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// user id for proxy authentication
// Default: NULL
func (e *GstSoupHttpClientSink) SetProxyId(value string) error {
	return e.Element.SetProperty("proxy-id", value)
}

func (e *GstSoupHttpClientSink) GetProxyId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// user password for proxy authentication
// Default: NULL
func (e *GstSoupHttpClientSink) SetProxyPw(value string) error {
	return e.Element.SetProperty("proxy-pw", value)
}

func (e *GstSoupHttpClientSink) GetProxyPw() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-pw")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// user id for authentication
// Default: NULL
func (e *GstSoupHttpClientSink) SetUserId(value string) error {
	return e.Element.SetProperty("user-id", value)
}

func (e *GstSoupHttpClientSink) GetUserId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Receive data as a client over the network via HTTP using SOUP
type GstSoupHTTPSrc struct {
	*base.GstPushSrc
}

func NewSoupHTTPSrc() (*GstSoupHTTPSrc, error) {
	e, err := gst.NewElement("souphttpsrc")
	if err != nil {
		return nil, err
	}
	return &GstSoupHTTPSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

func NewSoupHTTPSrcWithName(name string) (*GstSoupHTTPSrc, error) {
	e, err := gst.NewElementWithName("souphttpsrc", name)
	if err != nil {
		return nil, err
	}
	return &GstSoupHTTPSrc{GstPushSrc: &base.GstPushSrc{GstBaseSrc: &base.GstBaseSrc{Element: e}}}, nil
}

// Set log level for soup's HTTP session log
// Default: headers (2)
func (e *GstSoupHTTPSrc) SetHttpLogLevel(value interface{}) error {
	return e.Element.SetProperty("http-log-level", value)
}

func (e *GstSoupHTTPSrc) GetHttpLogLevel() (interface{}, error) {
	return e.Element.GetProperty("http-log-level")
}

// Location to read from
// Default: NULL
func (e *GstSoupHTTPSrc) SetLocation(value string) error {
	return e.Element.SetProperty("location", value)
}

func (e *GstSoupHTTPSrc) GetLocation() (string, error) {
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

// HTTP location URI user id for authentication
// Default: NULL
func (e *GstSoupHTTPSrc) SetUserId(value string) error {
	return e.Element.SetProperty("user-id", value)
}

func (e *GstSoupHTTPSrc) GetUserId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Extra headers to append to the HTTP request

func (e *GstSoupHTTPSrc) SetExtraHeaders(value *gst.Structure) error {
	return e.Element.SetProperty("extra-headers", value)
}

func (e *GstSoupHTTPSrc) GetExtraHeaders() (*gst.Structure, error) {
	var value *gst.Structure
	var ok bool
	v, err := e.Element.GetProperty("extra-headers")
	if err != nil {
		return value, err
	}
	value, ok = v.(*gst.Structure)
	if !ok {
		return value, fmt.Errorf("could not cast value to *gst.Structure")
	}
	return value, nil
}

// The HTTP method to use (GET, HEAD, OPTIONS, etc)
// Default: NULL
func (e *GstSoupHTTPSrc) SetMethod(value string) error {
	return e.Element.SetProperty("method", value)
}

func (e *GstSoupHTTPSrc) GetMethod() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("method")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Location of a SSL anchor CA file to use
// Default: NULL
func (e *GstSoupHTTPSrc) SetSslCaFile(value string) error {
	return e.Element.SetProperty("ssl-ca-file", value)
}

func (e *GstSoupHTTPSrc) GetSslCaFile() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("ssl-ca-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Use system CA file
// Default: true
func (e *GstSoupHTTPSrc) SetSslUseSystemCaFile(value bool) error {
	return e.Element.SetProperty("ssl-use-system-ca-file", value)
}

func (e *GstSoupHTTPSrc) GetSslUseSystemCaFile() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssl-use-system-ca-file")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Value in seconds to timeout a blocking I/O (0 = No timeout).
// Default: 15, Min: 0, Max: 3600
func (e *GstSoupHTTPSrc) SetTimeout(value uint) error {
	return e.Element.SetProperty("timeout", value)
}

func (e *GstSoupHTTPSrc) GetTimeout() (uint, error) {
	var value uint
	var ok bool
	v, err := e.Element.GetProperty("timeout")
	if err != nil {
		return value, err
	}
	value, ok = v.(uint)
	if !ok {
		return value, fmt.Errorf("could not cast value to uint")
	}
	return value, nil
}

// Act like a live source
// Default: false
func (e *GstSoupHTTPSrc) SetIsLive(value bool) error {
	return e.Element.SetProperty("is-live", value)
}

func (e *GstSoupHTTPSrc) GetIsLive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("is-live")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Allow compressed content encodings
// Default: false
func (e *GstSoupHTTPSrc) SetCompress(value bool) error {
	return e.Element.SetProperty("compress", value)
}

func (e *GstSoupHTTPSrc) GetCompress() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("compress")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// HTTP request cookies

func (e *GstSoupHTTPSrc) SetCookies(value interface{}) error {
	return e.Element.SetProperty("cookies", value)
}

func (e *GstSoupHTTPSrc) GetCookies() (interface{}, error) {
	return e.Element.GetProperty("cookies")
}

// Enable internet radio mode (ask server to send shoutcast/icecast metadata interleaved with the actual stream data)
// Default: true
func (e *GstSoupHTTPSrc) SetIradioMode(value bool) error {
	return e.Element.SetProperty("iradio-mode", value)
}

func (e *GstSoupHTTPSrc) GetIradioMode() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("iradio-mode")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// HTTP proxy URI user id for authentication
// Default: NULL
func (e *GstSoupHTTPSrc) SetProxyId(value string) error {
	return e.Element.SetProperty("proxy-id", value)
}

func (e *GstSoupHTTPSrc) GetProxyId() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-id")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// HTTP proxy URI user password for authentication
// Default: NULL
func (e *GstSoupHTTPSrc) SetProxyPw(value string) error {
	return e.Element.SetProperty("proxy-pw", value)
}

func (e *GstSoupHTTPSrc) GetProxyPw() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy-pw")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Maximum number of retries until giving up (-1=infinite)
// Default: 3, Min: -1, Max: 2147483647
func (e *GstSoupHTTPSrc) SetRetries(value int) error {
	return e.Element.SetProperty("retries", value)
}

func (e *GstSoupHTTPSrc) GetRetries() (int, error) {
	var value int
	var ok bool
	v, err := e.Element.GetProperty("retries")
	if err != nil {
		return value, err
	}
	value, ok = v.(int)
	if !ok {
		return value, fmt.Errorf("could not cast value to int")
	}
	return value, nil
}

// TLS database with anchor certificate authorities used to validate the server certificate

func (e *GstSoupHTTPSrc) SetTlsDatabase(value interface{}) error {
	return e.Element.SetProperty("tls-database", value)
}

func (e *GstSoupHTTPSrc) GetTlsDatabase() (interface{}, error) {
	return e.Element.GetProperty("tls-database")
}

// Automatically follow HTTP redirects (HTTP Status Code 3xx)
// Default: true
func (e *GstSoupHTTPSrc) SetAutomaticRedirect(value bool) error {
	return e.Element.SetProperty("automatic-redirect", value)
}

func (e *GstSoupHTTPSrc) GetAutomaticRedirect() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("automatic-redirect")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Value of the User-Agent HTTP request header field
// Default: GStreamer souphttpsrc {VERSION}
func (e *GstSoupHTTPSrc) SetUserAgent(value string) error {
	return e.Element.SetProperty("user-agent", value)
}

func (e *GstSoupHTTPSrc) GetUserAgent() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-agent")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// HTTP proxy server URI

func (e *GstSoupHTTPSrc) SetProxy(value string) error {
	return e.Element.SetProperty("proxy", value)
}

func (e *GstSoupHTTPSrc) GetProxy() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("proxy")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Strict SSL certificate checking
// Default: true
func (e *GstSoupHTTPSrc) SetSslStrict(value bool) error {
	return e.Element.SetProperty("ssl-strict", value)
}

func (e *GstSoupHTTPSrc) GetSslStrict() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("ssl-strict")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// A GTlsInteraction object to be used when the connection or certificate database need to interact with the user.

func (e *GstSoupHTTPSrc) SetTlsInteraction(value interface{}) error {
	return e.Element.SetProperty("tls-interaction", value)
}

func (e *GstSoupHTTPSrc) GetTlsInteraction() (interface{}, error) {
	return e.Element.GetProperty("tls-interaction")
}

// HTTP location URI user password for authentication
// Default: NULL
func (e *GstSoupHTTPSrc) SetUserPw(value string) error {
	return e.Element.SetProperty("user-pw", value)
}

func (e *GstSoupHTTPSrc) GetUserPw() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("user-pw")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// Use HTTP persistent connections
// Default: true
func (e *GstSoupHTTPSrc) SetKeepAlive(value bool) error {
	return e.Element.SetProperty("keep-alive", value)
}

func (e *GstSoupHTTPSrc) GetKeepAlive() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("keep-alive")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}
