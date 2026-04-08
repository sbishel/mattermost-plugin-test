package main

import (
	"encoding/json"
	"fmt"
	"maps"
	"net/http"
	"sort"
	"strings"

	"github.com/mattermost/mattermost/server/public/model"
)

func (p *Plugin) writeJSON(w http.ResponseWriter, v any) {
	b, err := json.Marshal(v)
	if err != nil {
		p.API.LogError("Failed to marshal JSON response", "error", err.Error())
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
}

func (p *Plugin) writeEmptyJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprint(w, "{}")
}

func (p *Plugin) createBotPost(channelID, message string) {
	if _, appErr := p.API.CreatePost(&model.Post{
		UserId:    p.botID,
		ChannelId: channelID,
		Message:   message,
		Type:      postTypeE2ETestPlugin,
	}); appErr != nil {
		p.API.LogError("Failed to create post", "error", appErr.Error())
	}
}

func (p *Plugin) handleDialogSubmitWithValidation(w http.ResponseWriter, r *http.Request) {
	if !p.getConfiguration().EnableDialogCommands {
		http.Error(w, "Dialog commands are disabled", http.StatusForbidden)
		return
	}

	var request model.SubmitDialogRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if request.Cancelled {
		p.createBotPost(request.ChannelId, "Dialog was cancelled")
		p.writeEmptyJSON(w)
		return
	}

	val := request.Submission[dialogElementNameNumber]
	if val != "42" && val != float64(42) {
		p.writeJSON(w, model.SubmitDialogResponse{
			Errors: map[string]string{
				dialogElementNameNumber: "This must be 42",
			},
		})
		return
	}

	var sb strings.Builder
	sb.WriteString("Full dialog submission:\n")
	keys := make([]string, 0, len(request.Submission))
	for k := range request.Submission {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(&sb, "- **%s**: %s\n", k, interfaceToString(request.Submission[k]))
	}
	fmt.Fprintf(&sb, "- **state**: %s\n", request.State)
	msg := sb.String()

	p.createBotPost(request.ChannelId, msg)
	p.writeEmptyJSON(w)
}

func (p *Plugin) handleDialogSubmitConfirm(w http.ResponseWriter, r *http.Request) {
	if !p.getConfiguration().EnableDialogCommands {
		http.Error(w, "Dialog commands are disabled", http.StatusForbidden)
		return
	}

	var request model.SubmitDialogRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if request.Cancelled {
		p.writeEmptyJSON(w)
		return
	}

	p.createBotPost(request.ChannelId, "Confirmation received.")
	p.writeEmptyJSON(w)
}

func (p *Plugin) handleDialogSubmitGeneric(w http.ResponseWriter, r *http.Request) {
	if !p.getConfiguration().EnableDialogCommands {
		http.Error(w, "Dialog commands are disabled", http.StatusForbidden)
		return
	}

	var request model.SubmitDialogRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	var sb strings.Builder
	sb.WriteString("Dialog submission:\n")
	keys := make([]string, 0, len(request.Submission))
	for k := range request.Submission {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(&sb, "- **%s**: %s\n", k, interfaceToString(request.Submission[k]))
	}
	if request.State != "" {
		fmt.Fprintf(&sb, "- **state**: %s\n", request.State)
	}
	msg := sb.String()

	p.createBotPost(request.ChannelId, msg)
	p.writeEmptyJSON(w)
}

func (p *Plugin) handleDialogWithError(w http.ResponseWriter, r *http.Request) {
	if !p.getConfiguration().EnableDialogCommands {
		http.Error(w, "Dialog commands are disabled", http.StatusForbidden)
		return
	}

	var request model.SubmitDialogRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if request.Cancelled {
		p.writeEmptyJSON(w)
		return
	}

	p.writeJSON(w, model.SubmitDialogResponse{
		Error: "This is an error from the dialog submission.",
	})
}

func (p *Plugin) handleDialogFieldRefresh(w http.ResponseWriter, r *http.Request) {
	if !p.getConfiguration().EnableDialogCommands {
		http.Error(w, "Dialog commands are disabled", http.StatusForbidden)
		return
	}

	var request model.SubmitDialogRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	pluginURL := fmt.Sprintf("/plugins/%s", manifest.Id)
	category, ok := request.Submission["category"].(string)
	if !ok {
		p.writeJSON(w, model.SubmitDialogResponse{Error: "Invalid category value."})
		return
	}

	var elements []model.DialogElement
	elements = append(elements, model.DialogElement{
		DisplayName: "Category",
		Name:        "category",
		Type:        "select",
		Refresh:     true,
		Default:     category,
		Options: []*model.PostActionOptions{
			{Text: "General", Value: "general"},
			{Text: "Advanced", Value: "advanced"},
		},
		HelpText: "Selecting changes available options below.",
	})

	if category == "advanced" {
		elements = append(elements, model.DialogElement{
			DisplayName: "Advanced Setting",
			Name:        "advanced_setting",
			Type:        "text",
			HelpText:    "This field appears for advanced category.",
		})
	} else {
		elements = append(elements, model.DialogElement{
			DisplayName: "Details",
			Name:        "details",
			Type:        "text",
			Optional:    true,
			HelpText:    "Additional details based on category.",
		})
	}

	resp := model.SubmitDialogResponse{
		Type: string(model.SubmitDialogResponseTypeForm),
		Form: &model.Dialog{
			CallbackId:  "field-refresh",
			Title:       "Field Refresh Dialog",
			SourceURL:   fmt.Sprintf("%s/dialog/field-refresh", pluginURL),
			SubmitLabel: "Submit",
			Elements:    elements,
		},
	}
	p.writeJSON(w, resp)
}

func (p *Plugin) handleDialogMultistep(w http.ResponseWriter, r *http.Request) {
	if !p.getConfiguration().EnableDialogCommands {
		http.Error(w, "Dialog commands are disabled", http.StatusForbidden)
		return
	}

	var request model.SubmitDialogRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if request.Cancelled {
		p.writeEmptyJSON(w)
		return
	}

	switch request.CallbackId {
	case "step1":
		stateJSON, err := json.Marshal(request.Submission)
		if err != nil {
			p.API.LogError("Failed to marshal step1 state", "error", err.Error())
			p.writeJSON(w, model.SubmitDialogResponse{Error: "Internal error saving form state."})
			return
		}
		dialog := getDialogStep2(string(stateJSON))
		p.writeJSON(w, model.SubmitDialogResponse{
			Type: string(model.SubmitDialogResponseTypeForm),
			Form: &dialog,
		})

	case "step2":
		merged := map[string]any{}
		if request.State != "" {
			if err := json.Unmarshal([]byte(request.State), &merged); err != nil {
				p.API.LogError("Failed to unmarshal step2 state", "error", err.Error())
				p.writeJSON(w, model.SubmitDialogResponse{Error: "Internal error loading form state."})
				return
			}
		}
		maps.Copy(merged, request.Submission)
		stateJSON, err := json.Marshal(merged)
		if err != nil {
			p.API.LogError("Failed to marshal step2 state", "error", err.Error())
			p.writeJSON(w, model.SubmitDialogResponse{Error: "Internal error saving form state."})
			return
		}
		dialog := getDialogStep3Summary(string(stateJSON))
		p.writeJSON(w, model.SubmitDialogResponse{
			Type: string(model.SubmitDialogResponseTypeForm),
			Form: &dialog,
		})

	case "step3":
		merged := map[string]any{}
		if request.State != "" {
			if err := json.Unmarshal([]byte(request.State), &merged); err != nil {
				p.API.LogError("Failed to unmarshal step3 state", "error", err.Error())
				p.writeJSON(w, model.SubmitDialogResponse{Error: "Internal error loading form state."})
				return
			}
		}
		maps.Copy(merged, request.Submission)

		var sb strings.Builder
		sb.WriteString("Registration complete!\n")
		keys := make([]string, 0, len(merged))
		for k := range merged {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Fprintf(&sb, "- **%s**: %s\n", k, interfaceToString(merged[k]))
		}
		msg := sb.String()

		p.createBotPost(request.ChannelId, msg)
		p.writeEmptyJSON(w)

	default:
		p.writeJSON(w, model.SubmitDialogResponse{
			Error: "Unknown step",
		})
	}
}

var allRoles = []model.DialogSelectOption{
	{Text: "System Admin", Value: "system_admin"},
	{Text: "System User", Value: "system_user"},
	{Text: "Team Admin", Value: "team_admin"},
	{Text: "Team User", Value: "team_user"},
	{Text: "Channel Admin", Value: "channel_admin"},
	{Text: "Channel User", Value: "channel_user"},
	{Text: "Guest", Value: "guest"},
}

func (p *Plugin) handleDynamicRoles(w http.ResponseWriter, r *http.Request) {
	if !p.getConfiguration().EnableDialogCommands {
		http.Error(w, "Dialog commands are disabled", http.StatusForbidden)
		return
	}

	var req struct {
		Query string `json:"query"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	query := strings.ToLower(req.Query)
	var filtered []model.DialogSelectOption
	for _, role := range allRoles {
		if query == "" || strings.Contains(strings.ToLower(role.Text), query) || strings.Contains(strings.ToLower(role.Value), query) {
			filtered = append(filtered, role)
		}
	}

	p.writeJSON(w, model.LookupDialogResponse{Items: filtered})
}
