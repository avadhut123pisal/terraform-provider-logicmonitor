// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/logicmonitor/lm-sdk-go/client/lm"
)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/santaba/rest"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// Config information for LMSdkGo client
type Config struct {
	AccessKey    *string
	AccessID     *string
	TransportCfg *TransportConfig
}

// NewConfig create a new empty client Config
func NewConfig() *Config {
	return &Config{
		TransportCfg: DefaultTransportConfig(),
	}
}

// SetAccessID for the client Config
func (c *Config) SetAccessID(accessID *string) {
	c.AccessID = accessID
}

// SetAccessKey for the client Config
func (c *Config) SetAccessKey(accessKey *string) {
	c.AccessKey = accessKey
}

// SetAccountDomain for the client Config
func (c *Config) SetAccountDomain(accountDomain *string) {
	if c.TransportCfg == nil {
		c.TransportCfg = DefaultTransportConfig()
	}
	domain := *accountDomain
	c.TransportCfg = c.TransportCfg.WithHost(domain)
}

// New creates a new LM sdk go client
func New(c *Config) *LMSdkGo {
	transport := httptransport.New(c.TransportCfg.Host, c.TransportCfg.BasePath, c.TransportCfg.Schemes)
	transport.Producers["application/version4+json"] = runtime.JSONProducer()
	transport.Producers["application/json; charset=utf-8"] = runtime.JSONProducer()

	authInfo := LMv1Auth(*c.AccessID, *c.AccessKey)

	cli := new(LMSdkGo)
	cli.Transport = transport

	cli.LM = lm.New(transport, strfmt.Default, authInfo)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// LMSdkGo is a client for LM sdk go
type LMSdkGo struct {
	LM *lm.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *LMSdkGo) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.LM.SetTransport(transport)

}

func LMv1Auth(accessId, accessKey string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		// get epoch
		now := time.Now()
		nanos := now.UnixNano()
		epoch := strconv.FormatInt(nanos/1000000, 10)

		// build the signature
		h := hmac.New(sha256.New, []byte(accessKey))
		h.Write([]byte(r.GetMethod() + epoch))

		if r.GetBodyParam() != nil {
			buf := new(bytes.Buffer)
			enc := json.NewEncoder(buf)
			enc.SetEscapeHTML(false)
			_ = enc.Encode(r.GetBodyParam())
			h.Write(buf.Bytes())
		}

		if r.GetFileParam() != nil {
			for _, files := range r.GetFileParam() {
				for i, file := range files {
					buf := bytes.NewBuffer(nil)
					buf.ReadFrom(file)
					h.Write(buf.Bytes())
					file = runtime.NamedReader(file.Name(), bytes.NewReader(buf.Bytes()))
					files[i] = file
				}
			}
		}

		h.Write([]byte(r.GetPath()))
		hexDigest := hex.EncodeToString(h.Sum(nil))
		signature := base64.StdEncoding.EncodeToString([]byte(hexDigest))
		r.SetHeaderParam("Authorization", fmt.Sprintf("LMv1 %s:%s:%s", accessId, signature, epoch))
		return r.SetHeaderParam("X-version", "4")
	})
}
