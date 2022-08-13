package utils

import (
	"fmt"
	"os"
)

type UrlInfo struct {
	UseSSL      bool
	Locate      bool
	UseInternal bool
	Domain      string
}

type ServiceInfo struct {
	Service string
	Region  string
}

const (
	urlTemplate         = ".api.%s"
	internalUrlTemplate = "internal.api.%s"
	http                = "http"
	https               = "https"
	EnvEndpointDomain   = "KSC_ENDPOINT_DOMAIN"
)

func Url(urlInfo *UrlInfo, info ServiceInfo) string {
	p := Protocol(urlInfo.UseSSL)
	domain := urlInfo.Domain
	if domain == "" {
		if v, ok := os.LookupEnv(EnvEndpointDomain); ok {
			domain = v
		}
		if domain == "" {
			domain = "ksyun.com"
		}
	}
	if urlInfo.UseInternal {
		return p + "://" + fmt.Sprintf(internalUrlTemplate, domain)
	}
	url := fmt.Sprintf(urlTemplate, domain)
	if urlInfo.Locate && &info.Region != nil {
		return p + "://" + info.Service + "." + info.Region + url
	}
	return p + "://" + info.Service + url
}

func Protocol(useSSL bool) string {
	if useSSL {
		return https
	}
	return http
}
