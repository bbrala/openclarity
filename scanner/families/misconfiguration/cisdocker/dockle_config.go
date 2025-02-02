// Copyright © 2024 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
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

package cisdocker

import (
	"time"

	"github.com/openclarity/openclarity/scanner/common"
	"github.com/openclarity/openclarity/scanner/families/misconfiguration/cisdocker/config"

	dockle_config "github.com/Portshift/dockle/config"
	"github.com/sirupsen/logrus"
)

const (
	DefaultCISDockerTimeout = 2 * time.Minute
)

func createDockleConfig(logger *logrus.Entry, inputType common.InputType, name string, config config.Config) *dockle_config.Config {
	dockleConfig := &dockle_config.Config{
		Debug:      logger.Logger.Level == logrus.DebugLevel,
		Timeout:    DefaultCISDockerTimeout,
		LocalImage: true,
	}

	if config.Registry != nil {
		dockleConfig.LocalImage = false
		dockleConfig.Insecure = config.Registry.SkipVerifyTLS
		dockleConfig.NonSSL = config.Registry.UseHTTP
		if len(config.Registry.Auths) > 0 {
			dockleConfig.AuthURL = config.Registry.Auths[0].Authority
			dockleConfig.Username = config.Registry.Auths[0].Username
			dockleConfig.Password = config.Registry.Auths[0].Password
			dockleConfig.Token = config.Registry.Auths[0].Token
		}
	}

	if config.Timeout != 0 {
		dockleConfig.Timeout = config.Timeout
	}

	// nolint:exhaustive
	switch inputType {
	case common.DOCKERARCHIVE:
		dockleConfig.FilePath = name
	case common.ROOTFS, common.DIR:
		dockleConfig.DirPath = name
	default:
		dockleConfig.ImageName = name
	}

	return dockleConfig
}
