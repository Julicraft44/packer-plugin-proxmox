// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package proxmox

import (
	"crypto/tls"
	"log"

	"github.com/Telmate/proxmox-api-go/proxmox"
)

func NewProxmoxClient(config ProxmoxConnectConfig, debug bool) (*proxmox.Client, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: config.SkipCertValidation,
	}

	client, err := proxmox.NewClient(
		config.ProxmoxURL.String(),
		nil,
		"",
		tlsConfig,
		"",
		int(config.TaskTimeout.Seconds()),
	)
	if err != nil {
		return nil, err
	}

	*proxmox.Debug = debug

	if config.Token != "" {
		// configure token auth
		log.Print("using token auth")
		client.SetAPIToken(config.Username, config.Token)
	} else {
		// fallback to login if not using tokens
		log.Print("using password auth")
		err = client.Login(config.Username, config.Password, "")
		if err != nil {
			return nil, err
		}
	}

	return client, nil
}
