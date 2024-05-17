package b

import (
	"os"
	"os/exec"
	"path/filepath"
	"plugin"
	"testing"

	"github.com/pfi79/go-test-cover-error/a"
	"github.com/stretchr/testify/require"
)

func TestEndorsementPlugin(t *testing.T) {
	testDir := t.TempDir()
	pluginPath := filepath.Join(testDir, "a.so")
	aPluginPath := "github.com/pfi79/go-test-cover-error/testdata/"

	cmd := exec.Command("go", "build", "-o", pluginPath, "-buildmode=plugin")
	cmd.Args = append(cmd.Args, aPluginPath)
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "Could not build plugin: "+string(output))

	testReg := make(map[string]a.PluginFactory)
	t.Log(len(testReg))

	_, err = os.Stat(pluginPath)
	require.NoError(t, err, "Could not stat plugin: "+pluginPath)

	p, err := plugin.Open(pluginPath)
	require.NoError(t, err, "Could not open plugin: "+pluginPath)
	t.Log(p)
}
