package updater

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNpmUpdater(t *testing.T) {
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
