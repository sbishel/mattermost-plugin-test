package main

import (
	"strings"
	"sync"

	"github.com/gorilla/mux"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
	"github.com/mattermost/mattermost/server/public/pluginapi"
	"github.com/pkg/errors"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// client is the Mattermost server API client.
	client *pluginapi.Client

	// router is the HTTP router for handling API requests.
	router *mux.Router

	// botID is the user ID of the bot account created by this plugin.
	botID string

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

// OnActivate is invoked when the plugin is activated. If an error is returned, the plugin will be deactivated.
func (p *Plugin) OnActivate() error {
	p.client = pluginapi.NewClient(p.API, p.Driver)

	botID, err := p.client.Bot.EnsureBot(&model.Bot{
		Username:    "test-plugin",
		DisplayName: "Test Plugin Bot",
		Description: "Bot for the Test Plugin.",
	})
	if err != nil {
		return errors.Wrap(err, "failed to ensure bot")
	}
	p.botID = botID

	if err := p.client.SlashCommand.Register(&model.Command{
		Trigger:          "e2e-dialog",
		AutoComplete:     true,
		AutoCompleteDesc: "Open interactive test dialogs",
		AutoCompleteHint: "[subcommand]",
		AutocompleteData: getDialogAutocompleteData(),
	}); err != nil {
		return errors.Wrap(err, "failed to register e2e-dialog command")
	}

	p.router = p.initRouter()

	return nil
}

func (p *Plugin) ExecuteCommand(c *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	fields := strings.Fields(args.Command)
	if len(fields) == 0 {
		return &model.CommandResponse{
			ResponseType: model.CommandResponseTypeEphemeral,
			Text:         "Empty command",
		}, nil
	}

	trigger := strings.TrimPrefix(fields[0], "/")
	if trigger == "e2e-dialog" {
		return p.executeDialogCommand(args)
	}

	return &model.CommandResponse{
		ResponseType: model.CommandResponseTypeEphemeral,
		Text:         "Unknown command: " + args.Command,
	}, nil
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
