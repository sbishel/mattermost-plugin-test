package main

import (
	"context"
	"testing"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/require"

	"github.com/mattermost/mattermost-plugin-test/server/testhelper"
)

// TestPluginActivation verifies that the plugin (built from this repository via `make dist`)
// starts successfully on a real Mattermost server. Setup() spins up Postgres + Mattermost
// containers, resets the database, uploads the plugin bundle, and enables it. This test
// then confirms the plugin reaches the Running state by querying the plugin status API.
func TestPluginActivation(t *testing.T) {
	th := testhelper.Setup(t)

	ctx := context.Background()
	statuses, _, err := th.AdminClient.GetPluginStatuses(ctx)
	require.NoError(t, err)

	found := false
	for _, s := range statuses {
		if s.PluginId == testhelper.PluginID() && s.State == model.PluginStateRunning {
			found = true
			break
		}
	}
	require.True(t, found, "plugin %s should be running", testhelper.PluginID())
}
