// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package extensions

type MetadataAPI struct {
	Version string `yaml:"version"`
}

type MetadataConfig struct {
	Metadata Metadata `yaml:"metadata"`
}

type Metadata struct {
	Author        string     `yaml:"author"`
	Compatibility CompatInfo `yaml:"compatibility"`
	Description   string     `yaml:"description"`
	Name          string     `yaml:"name"`
	Versions      Versions   `yaml:"versions"`
}

type Versions struct {
	ExtensionVersion string `yaml:"extension"`
	PkgVersion       string `yaml:"pkg"`
}

type CompatInfo struct {
	Talos  TalosCompat  `yaml:"talos"`
	Kernel KernelCompat `yaml:"kernel"`
}

type TalosCompat struct {
	Version string `yaml:"version"`
}

type KernelCompat struct {
	Version string `yaml:"version"`
}
