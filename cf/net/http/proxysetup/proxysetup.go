package proxysetup

import (
	"net/http"
	"os"

	"github.com/anynines/go-proxy-setup-ntlm/proxysetup/ntlm"
)

func FromEnvironment() func(http.ProxySetupContext) error {
	if os.Getenv("NTLM_PROXY") == "true" {
		return ntlm.ProxySetup
	}

	return http.DefaultProxySetup
}
