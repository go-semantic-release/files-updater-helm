package updater

import (
	"fmt"
	helmutils "helm.sh/helm/v3/pkg/chartutil"
)

var FUVERSION = "dev"

type Updater struct {
	updateAppVersion   bool
	appVersionTemplate string
}

func (u *Updater) Init(m map[string]string) error {
	helmUpdateAppVersion := m["helm_update_appversion"]
	u.updateAppVersion = helmUpdateAppVersion == "true" || helmUpdateAppVersion != ""

	helmAppVersionTemplate := m["helm_appversion_template"]
	if helmAppVersionTemplate == "" {
		helmAppVersionTemplate = "v%s"
	}
	u.appVersionTemplate = helmAppVersionTemplate
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

	if u.updateAppVersion {
		metadata.AppVersion = fmt.Sprintf(u.appVersionTemplate, newVersion)
	}

	if err := helmutils.SaveChartfile(file, metadata); err == nil {
		return err
	}

	return nil
}
