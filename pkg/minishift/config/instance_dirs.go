/*
Copyright (C) 2017 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"github.com/minishift/minishift/pkg/minikube/constants"
)

type MinishiftDirs struct {
	Home       string
	Config     string
	Machines   string
	Certs      string
	Cache      string
	IsoCache   string
	OcCache    string
	ImageCache string
	Addons     string
	Logs       string
	Tmp        string
}

var InstanceDirs *MinishiftDirs

// Constructor for MinishiftDirs
func NewMinishiftDirs() *MinishiftDirs {
	constants.Minipath = constants.GetProfileHomeDir()

	return &MinishiftDirs{
		Home:       constants.Minipath,
		Certs:      constants.MakeMiniPath("certs"),
		Machines:   constants.MakeMiniPath("machines"),
		Cache:      constants.MakeMiniPath("cache"),
		IsoCache:   constants.MakeMiniPath("cache", "iso"),
		OcCache:    constants.MakeMiniPath("cache", "oc"),
		ImageCache: constants.MakeMiniPath("cache", "images"),
		Addons:     constants.MakeMiniPath("addons"),
		Logs:       constants.MakeMiniPath("logs"),
		Tmp:        constants.MakeMiniPath("tmp"),
		Config:     constants.MakeMiniPath("config"),
	}
}
