package chargebee

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Environment struct {
	Key             string
	SiteName        string
	ChargebeeDomain string
	Protocol        string
}

var (
	TotalHTTPTimeout    = 80 * time.Second
	ExportWaitInSecs      = 3 * time.Second
	TimeMachineWaitInSecs = 3 * time.Second
	DefaultEnv            Environment
)

const (
	APIVersion = "v2"
	Charset    = "UTF-8"
)

func Configure(key string, siteName string) {
	if key == "" || siteName == "" {
		panic(errors.New("Key/siteName cannot be empty"))
	}
	DefaultEnv = Environment{Key: key, SiteName: siteName}
}
func WithHTTPClient(c *http.Client) {
	httpClient = c
}
func (env *Environment) apiBaseUrl() string {
	if env.Protocol == "" {
		env.Protocol = "https"
	}
	if env.ChargebeeDomain != "" {
		return fmt.Sprintf("%v://%v.%v/api/%v", env.Protocol, env.SiteName, env.ChargebeeDomain, APIVersion)
	}
	return fmt.Sprintf("https://%v.chargebee.com/api/%v", env.SiteName, APIVersion)
}

func DefaultConfig() Environment {
	if DefaultEnv == (Environment{}) {
		panic(errors.New("The default environment has not been configured"))
	}
	return DefaultEnv
}

func UpdateTotalHTTPTimeout(timeout time.Duration) {
	TotalHTTPTimeout = timeout
}
