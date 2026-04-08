package main

import (
	"fmt"
	"strings"

	"github.com/mattermost/mattermost/server/public/model"
)

func (p *Plugin) executeDialogCommand(args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	if !p.getConfiguration().EnableDialogCommands {
		return &model.CommandResponse{
			ResponseType: model.CommandResponseTypeEphemeral,
			Text:         "Dialog commands are disabled. Enable in System Console > Plugins > Test Plugin.",
		}, nil
	}

	fields := strings.Fields(args.Command)
	var subcommand string
	if len(fields) > 1 {
		subcommand = fields[1]
	}

	pluginURL := fmt.Sprintf("/plugins/%s", manifest.Id)

	var dialog model.Dialog
	var callbackURL string

	switch subcommand {
	case "":
		dialog = getDialogWithSampleElements(pluginURL)
		callbackURL = fmt.Sprintf("%s/dialog/submit-with-validation", pluginURL)
	case "text":
		dialog = getDialogText()
		callbackURL = fmt.Sprintf("%s/dialog/submit-generic", pluginURL)
	case "boolean":
		dialog = getDialogBoolean()
		callbackURL = fmt.Sprintf("%s/dialog/submit-generic", pluginURL)
	case "select":
		dialog = getDialogSelect()
		callbackURL = fmt.Sprintf("%s/dialog/submit-generic", pluginURL)
	case "multi-select":
		dialog = getDialogMultiSelect(pluginURL)
		callbackURL = fmt.Sprintf("%s/dialog/submit-generic", pluginURL)
	case "date":
		dialog = getDialogDate()
		callbackURL = fmt.Sprintf("%s/dialog/submit-generic", pluginURL)
	case "datetime":
		dialog = getDialogDateTime()
		callbackURL = fmt.Sprintf("%s/dialog/submit-generic", pluginURL)
	case "datetime-basic":
		dialog = getDialogDateTimeBasic()
		callbackURL = fmt.Sprintf("%s/dialog/submit-generic", pluginURL)
	case "datetime-timezone":
		dialog = getDialogDateTimeTimezone()
		callbackURL = fmt.Sprintf("%s/dialog/submit-generic", pluginURL)
	case "no-elements":
		dialog = getDialogWithoutElements()
		callbackURL = fmt.Sprintf("%s/dialog/submit-confirm", pluginURL)
	case "field-refresh":
		dialog = getDialogWithFieldRefresh(pluginURL)
		callbackURL = fmt.Sprintf("%s/dialog/submit-generic", pluginURL)
	case "multi-step":
		dialog = getDialogStep1()
		callbackURL = fmt.Sprintf("%s/dialog/multistep", pluginURL)
	case "error":
		dialog = getDialogWithSampleElements(pluginURL)
		callbackURL = fmt.Sprintf("%s/dialog/error", pluginURL)
	case "error-no-elements":
		dialog = getDialogWithoutElements()
		callbackURL = fmt.Sprintf("%s/dialog/error", pluginURL)
	default:
		return &model.CommandResponse{
			ResponseType: model.CommandResponseTypeEphemeral,
			Text:         fmt.Sprintf("Unknown subcommand: %s", subcommand),
		}, nil
	}

	appErr := p.API.OpenInteractiveDialog(model.OpenDialogRequest{
		TriggerId: args.TriggerId,
		URL:       callbackURL,
		Dialog:    dialog,
	})
	if appErr != nil {
		return &model.CommandResponse{
			ResponseType: model.CommandResponseTypeEphemeral,
			Text:         fmt.Sprintf("Failed to open dialog: %s", appErr.Error()),
		}, nil
	}

	return &model.CommandResponse{}, nil
}

func getDialogAutocompleteData() *model.AutocompleteData {
	dialog := model.NewAutocompleteData("e2e-dialog", "[subcommand]", "Open interactive test dialogs")

	dialog.AddCommand(model.NewAutocompleteData("text", "", "Text field variants"))
	dialog.AddCommand(model.NewAutocompleteData("boolean", "", "Boolean field variants"))
	dialog.AddCommand(model.NewAutocompleteData("select", "", "Select field variants"))
	dialog.AddCommand(model.NewAutocompleteData("multi-select", "", "Multi-select and dynamic search"))
	dialog.AddCommand(model.NewAutocompleteData("date", "", "Date field variants"))
	dialog.AddCommand(model.NewAutocompleteData("datetime", "", "DateTime field variants"))
	dialog.AddCommand(model.NewAutocompleteData("datetime-basic", "", "Webhook-aligned date/datetime basics (min date, intervals, relative dates)"))
	dialog.AddCommand(model.NewAutocompleteData("datetime-timezone", "", "Webhook-aligned timezone support and manual time entry"))
	dialog.AddCommand(model.NewAutocompleteData("no-elements", "", "Confirmation dialog without fields"))
	dialog.AddCommand(model.NewAutocompleteData("field-refresh", "", "Dialog with conditional field refresh"))
	dialog.AddCommand(model.NewAutocompleteData("multi-step", "", "Multi-step registration workflow"))
	dialog.AddCommand(model.NewAutocompleteData("error", "", "Dialog that returns errors on submit"))
	dialog.AddCommand(model.NewAutocompleteData("error-no-elements", "", "Confirmation dialog that returns errors"))

	return dialog
}
