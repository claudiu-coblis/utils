// Copyright 2015 Canonical Ltd.
// Copyright 2015 Cloudbase Solutions SRL
// Licensed under the AGPLv3, see LICENCE file for details.

package configurer

const (
	// AptSourcesFile is the default file which list all core
	// sources for apt packages on an apt-based system.
	AptSourcesFile = "/etc/apt/sources.list"

	// AptListsDirectory is the location of the APT sources list.
	AptListsDirectory = "/var/lib/apt/lists"

	// AptConfigDirectory is the default directory in which
	// apt configuration files are stored.
	AptConfigFile = "/etc/apt/apt.conf.d"

	// AptProxyConfigFile is the full file path for the proxy settings that are
	// written by cloudinit and the machine environ worker.
	AptProxyConfigFile = AptConfigFile + "/42-juju-proxy-settings"

	// ExtractAptSource is a shell command that will extract the
	// currently configured APT source location. We assume that
	// the first source for "main" in the file is the one that
	// should be replaced throughout the file.
	ExtractAptSource = `awk "/^deb .* $(lsb_release -sc) .*main.*\$/{print \$2;exit}" ` + AptSourcesFile

	// AptSourceListPrefix is a shell program that translates an
	// APT source (piped from stdin) to a file prefix. The algorithm
	// involves stripping up to one trailing slash, stripping the
	// URL scheme prefix, and finally translating slashes to
	// underscores.
	AptSourceListPrefix = `sed 's,.*://,,' | sed 's,/$,,' | tr / _`
)

// AptPreferenceTemplate is the template specific to an apt preference file.
var AptPreferenceTemplate = `
Explanation: {{.Explanation}}
Package: {{.Package}}
Pin: {{.Pin}}
Pin-Priority: {{.PinPriority}}
`[1:]

// AptSourceTemplate is the template specific to an apt source file.
var AptSourceTemplate = `
# {{.Name}} (added by Juju)
deb {{.Url}} %s main
# deb-src {{.Url}} %s main
`[1:]
