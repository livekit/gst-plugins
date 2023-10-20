// source: gst-plugins-bad

package factory

import (
	"fmt"

	"github.com/go-gst/go-gst/gst"
	"github.com/go-gst/go-gst/gst/base"
)

// AES buffer decryption
type GstAesDec struct {
	*base.GstBaseTransform
}

func NewAesDec() (*GstAesDec, error) {
	e, err := gst.NewElement("aesdec")
	if err != nil {
		return nil, err
	}
	return &GstAesDec{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAesDecWithName(name string) (*GstAesDec, error) {
	e, err := gst.NewElementWithName("aesdec", name)
	if err != nil {
		return nil, err
	}
	return &GstAesDec{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// cipher mode
// Default: aes-128-cbc (0)
func (e *GstAesDec) SetCipher(value GstAesCipher) error {
	e.Element.SetArg("cipher", string(value))
	return nil
}

func (e *GstAesDec) GetCipher() (GstAesCipher, error) {
	var value GstAesCipher
	var ok bool
	v, err := e.Element.GetProperty("cipher")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAesCipher)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAesCipher")
	}
	return value, nil
}

// AES encryption initialization vector (in hexadecimal). Length must equal AES block length (16 bytes)

func (e *GstAesDec) SetIv(value string) error {
	return e.Element.SetProperty("iv", value)
}

func (e *GstAesDec) GetIv() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("iv")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// AES encryption key (in hexadecimal). Length (in bytes) must be equivalent to the number of bits in the key length : 16 bytes for AES 128 and 32 bytes for AES 256

func (e *GstAesDec) SetKey(value string) error {
	return e.Element.SetProperty("key", value)
}

func (e *GstAesDec) GetKey() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("key")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// If true, pad each buffer using PKCS7 padding scheme. Otherwise, onlypad final buffer
// Default: true
func (e *GstAesDec) SetPerBufferPadding(value bool) error {
	return e.Element.SetProperty("per-buffer-padding", value)
}

func (e *GstAesDec) GetPerBufferPadding() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("per-buffer-padding")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Read initialization vector from first 16 bytes of first buffer
// Default: false
func (e *GstAesDec) SetSerializeIv(value bool) error {
	return e.Element.SetProperty("serialize-iv", value)
}

func (e *GstAesDec) GetSerializeIv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("serialize-iv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// AES buffer encryption
type GstAesEnc struct {
	*base.GstBaseTransform
}

func NewAesEnc() (*GstAesEnc, error) {
	e, err := gst.NewElement("aesenc")
	if err != nil {
		return nil, err
	}
	return &GstAesEnc{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

func NewAesEncWithName(name string) (*GstAesEnc, error) {
	e, err := gst.NewElementWithName("aesenc", name)
	if err != nil {
		return nil, err
	}
	return &GstAesEnc{GstBaseTransform: &base.GstBaseTransform{Element: e}}, nil
}

// cipher mode
// Default: aes-128-cbc (0)
func (e *GstAesEnc) SetCipher(value GstAesCipher) error {
	e.Element.SetArg("cipher", string(value))
	return nil
}

func (e *GstAesEnc) GetCipher() (GstAesCipher, error) {
	var value GstAesCipher
	var ok bool
	v, err := e.Element.GetProperty("cipher")
	if err != nil {
		return value, err
	}
	value, ok = v.(GstAesCipher)
	if !ok {
		return value, fmt.Errorf("could not cast value to GstAesCipher")
	}
	return value, nil
}

// AES encryption initialization vector (in hexadecimal). Length must equal AES block length (16 bytes)

func (e *GstAesEnc) SetIv(value string) error {
	return e.Element.SetProperty("iv", value)
}

func (e *GstAesEnc) GetIv() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("iv")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// AES encryption key (in hexadecimal). Length (in bytes) must be equivalent to the number of bits in the key length : 16 bytes for AES 128 and 32 bytes for AES 256

func (e *GstAesEnc) SetKey(value string) error {
	return e.Element.SetProperty("key", value)
}

func (e *GstAesEnc) GetKey() (string, error) {
	var value string
	var ok bool
	v, err := e.Element.GetProperty("key")
	if err != nil {
		return value, err
	}
	value, ok = v.(string)
	if !ok {
		return value, fmt.Errorf("could not cast value to string")
	}
	return value, nil
}

// If true, pad each buffer using PKCS7 padding scheme. Otherwise, onlypad final buffer
// Default: true
func (e *GstAesEnc) SetPerBufferPadding(value bool) error {
	return e.Element.SetProperty("per-buffer-padding", value)
}

func (e *GstAesEnc) GetPerBufferPadding() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("per-buffer-padding")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

// Store initialization vector in first 16 bytes of first buffer
// Default: false
func (e *GstAesEnc) SetSerializeIv(value bool) error {
	return e.Element.SetProperty("serialize-iv", value)
}

func (e *GstAesEnc) GetSerializeIv() (bool, error) {
	var value bool
	var ok bool
	v, err := e.Element.GetProperty("serialize-iv")
	if err != nil {
		return value, err
	}
	value, ok = v.(bool)
	if !ok {
		return value, fmt.Errorf("could not cast value to bool")
	}
	return value, nil
}

type GstAesCipher string

const (
	GstAesCipherAes128Cbc GstAesCipher = "aes-128-cbc" // AES 128 bit cipher key using CBC method
	GstAesCipherAes256Cbc GstAesCipher = "aes-256-cbc" // AES 256 bit cipher key using CBC method
)
