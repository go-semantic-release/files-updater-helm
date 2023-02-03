package updater

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestHelmUpdater(t *testing.T) {
	require := require.New(t)

	updater := &Updater{}

	nVer := "1.2.3"
	chartPath := "../../test/Chart.yaml"

	err := updater.Apply(chartPath, nVer)
	require.NoError(err)
	f, err := os.OpenFile(chartPath, os.O_RDONLY, 0)
	require.NoError(err)
	defer f.Close()
}

func TestHelmUpdaterAppVersion(t *testing.T) {
	require := require.New(t)

	updater := &Updater{}

	conf := map[string]string{
		"helm_update_appversion": "true",
	}

	require.NoError(updater.Init(conf))

	nVer := "1.2.3"
	chartPath := "../../test/Chart.yaml"

	err := updater.Apply(chartPath, nVer)
	require.NoError(err)
	f, err := os.ReadFile(chartPath)
	require.NoError(err)

	data := make(map[interface{}]interface{})
	err = yaml.Unmarshal(f, &data)
	require.NoError(err)
	require.Equal("v1.2.3", data["appVersion"])
}
