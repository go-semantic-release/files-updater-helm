package updater

import (
	helmutils "helm.sh/helm/v3/pkg/chartutil"
)

var FUVERSION = "dev"

type Updater struct {
}

func (u *Updater) Init(m map[string]string) error {
	return nil
}

func (u *Updater) Name() string {
	return "helm"
}

func (u *Updater) Version() string {
	return FUVERSION
}

func (u *Updater) ForFiles() string {
	return "Chart\\.yaml"
}

func (u *Updater) Apply(file, newVersion string) error {
	metadata, err := helmutils.LoadChartfile(file)
	if err != nil {
		return err
	}

	metadata.Version = newVersion

	helmutils.SaveChartfile(file, metadata)

	return nil
}
