// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package rewrite // import "miniflux.app/v2/internal/reader/rewrite"

import (
	"encoding/json"
	"net/url"
	"os"
	"strings"
)

var proxyMappings = map[string]string{}

func LoadProxyOverrides(externalFilePath string) error {
	externalData, err := os.ReadFile(externalFilePath)
	if err != nil {
		return err
	}
	var externalMappings map[string]string
	if err = json.Unmarshal(externalData, &externalMappings); err != nil {
		return err
	}
	for k, v := range externalMappings {
		proxyMappings[k] = v
	}

	return nil
}

// GetProxyForURL returns the proxy URL for the given URL if it exists, otherwise an empty string.
func GetProxyForURL(u string) string {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return ""
	}

	hostname := parsedUrl.Hostname()

	if proxy, ok := proxyMappings[hostname]; ok {
		return proxy
	}

	for suffix, proxy := range proxyMappings {
		if strings.HasSuffix(hostname, suffix) {
			return proxy
		}
	}

	return ""
}
