package plug1

import "plugin"

func Open(path string) (*plugin.Plugin, error) {
	return plugin.Open(path)
}
