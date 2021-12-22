package main

import (
	helmUpdater "github.com/go-semantic-release/files-updater-helm/pkg/updater"
	"github.com/go-semantic-release/semantic-release/v2/pkg/plugin"
	"github.com/go-semantic-release/semantic-release/v2/pkg/updater"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		FilesUpdater: func() updater.FilesUpdater {
			return &helmUpdater.Updater{}
		},
	})
}
