package main

import (
	"fmt"

	"github.com/mattermost/mattermost/server/public/model"
)

const (
	dialogElementNameNumber = "somenumber"
	dialogElementNameEmail  = "someemail"
	dialogStateSome         = "somestate"
	dialogIntroductionText  = "**Some** _introductory_ paragraph in Markdown formatted text with [a link](https://example.com)"
	postTypeE2ETestPlugin   = "custom_e2etest_plugin"
)

func interfaceToString(val any) string {
	switch v := val.(type) {
	case string:
		return v
	case bool:
		if v {
			return "true"
		}
		return "false"
	case float64:
		return fmt.Sprintf("%g", v)
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", v)
	}
}

func getDialogWithSampleElements(pluginURL string) model.Dialog {
	return model.Dialog{
		CallbackId:       "somecallbackid",
		Title:            "Test Dialog",
		IntroductionText: dialogIntroductionText,
		IconURL:          "http://www.mattermost.org/wp-content/uploads/2016/04/icon.png",
		SubmitLabel:      "Submit Test Dialog",
		NotifyOnCancel:   true,
		State:            dialogStateSome,
		Elements: []model.DialogElement{
			{
				DisplayName: "Display Name",
				Name:        "realname",
				Type:        "text",
				SubType:     "",
				Default:     "default text",
				Placeholder: "placeholder",
				HelpText:    "This a regular input in a dialog triggered by a bot.",
			},
			{
				DisplayName: "Email",
				Name:        dialogElementNameEmail,
				Type:        "text",
				SubType:     "email",
				Placeholder: "placeholder@bladekick.com",
				HelpText:    "This a regular email input.",
			},
			{
				DisplayName: "Password",
				Name:        "somepassword",
				Type:        "text",
				SubType:     "password",
				Placeholder: "Password",
				HelpText:    "This is a password input.",
			},
			{
				DisplayName: "Number",
				Name:        dialogElementNameNumber,
				Type:        "text",
				SubType:     "number",
				Default:     "7",
				HelpText:    "This a regular number input.",
			},
			{
				DisplayName: "Textarea",
				Name:        "sometextarea",
				Type:        "textarea",
				Placeholder: "placeholder",
				Optional:    true,
				HelpText:    "This is a textarea.",
				MinLength:   5,
				MaxLength:   100,
			},
			{
				DisplayName: "Option Selector",
				Name:        "someoptionselector",
				Type:        "select",
				Placeholder: "Choose an option...",
				HelpText:    "Choose an option from the list.",
				Options: []*model.PostActionOptions{
					{Text: "Option1", Value: "opt1"},
					{Text: "Option2", Value: "opt2"},
					{Text: "Option3", Value: "opt3"},
				},
			},
			{
				DisplayName: "Option Selector with default",
				Name:        "someoptionselectorwithdefault",
				Type:        "select",
				Default:     "opt2",
				Placeholder: "Choose an option...",
				HelpText:    "Choose an option (default opt2).",
				Options: []*model.PostActionOptions{
					{Text: "Option1", Value: "opt1"},
					{Text: "Option2", Value: "opt2"},
					{Text: "Option3", Value: "opt3"},
				},
			},
			{
				DisplayName: "User Selector",
				Name:        "someuserselector",
				Type:        "select",
				DataSource:  "users",
				Placeholder: "Choose a user...",
				HelpText:    "Choose a user from the list.",
			},
			{
				DisplayName: "Channel Selector",
				Name:        "somechannelselector",
				Type:        "select",
				DataSource:  "channels",
				Placeholder: "Choose a channel...",
				HelpText:    "Choose a channel.",
				Optional:    true,
			},
			{
				DisplayName: "Boolean Selector (required)",
				Name:        "someboolean",
				Type:        "bool",
				Placeholder: "Agrees to the terms",
			},
			{
				DisplayName: "Boolean Selector (optional)",
				Name:        "someboolean_optional",
				Type:        "bool",
				Placeholder: "Agrees to the terms",
				Optional:    true,
			},
			{
				DisplayName: "Boolean Selector (default true)",
				Name:        "someboolean_default_true",
				Type:        "bool",
				Default:     "true",
				Placeholder: "Accepts license",
			},
			{
				DisplayName: "Boolean Selector (default false)",
				Name:        "someboolean_default_false",
				Type:        "bool",
				Default:     "false",
				Placeholder: "Send notifications",
			},
			{
				DisplayName: "Boolean (optional default true)",
				Name:        "someboolean_optional_default_true",
				Type:        "bool",
				Default:     "true",
				Optional:    true,
				Placeholder: "Enable feature",
			},
			{
				DisplayName: "Boolean (optional default false)",
				Name:        "someboolean_optional_default_false",
				Type:        "bool",
				Default:     "false",
				Optional:    true,
				Placeholder: "Beta mode",
			},
			{
				DisplayName: "Radio Option Selector",
				Name:        "someradiooptionselector",
				Type:        "radio",
				Options: []*model.PostActionOptions{
					{Text: "Option1", Value: "opt1"},
					{Text: "Option2", Value: "opt2"},
					{Text: "Option3", Value: "opt3"},
				},
			},
		},
	}
}

func getDialogText() model.Dialog {
	return model.Dialog{
		Title:       "Text Fields",
		SubmitLabel: "Submit",
		Elements: []model.DialogElement{
			{
				DisplayName: "Plain Text",
				Name:        "optional_text",
				Type:        "text",
				Placeholder: "Optional text here",
				Optional:    true,
				HelpText:    "This is an optional text field.",
			},
			{
				DisplayName: "Required Text",
				Name:        "required_text",
				Type:        "text",
				Placeholder: "Required (3-128 chars)",
				MinLength:   3,
				MaxLength:   128,
				HelpText:    "Required text field with length constraints.",
			},
			{
				DisplayName: "Email",
				Name:        "email_field",
				Type:        "text",
				SubType:     "email",
				Placeholder: "user@example.com",
				HelpText:    "Email address field.",
			},
			{
				DisplayName: "Number",
				Name:        "number_field",
				Type:        "text",
				SubType:     "number",
				Default:     "42",
				HelpText:    "Number input field.",
			},
			{
				DisplayName: "Password",
				Name:        "password_field",
				Type:        "text",
				SubType:     "password",
				Placeholder: "Enter password",
				HelpText:    "Password field.",
			},
			{
				DisplayName: "Textarea",
				Name:        "textarea_field",
				Type:        "textarea",
				Placeholder: "Enter multiline text...",
				Optional:    true,
				MinLength:   5,
				MaxLength:   500,
				HelpText:    "Optional textarea with length constraints.",
			},
		},
	}
}

func getDialogBoolean() model.Dialog {
	return model.Dialog{
		Title:       "Boolean Fields",
		SubmitLabel: "Submit",
		Elements: []model.DialogElement{
			{
				DisplayName: "Required Boolean",
				Name:        "required_bool",
				Type:        "bool",
				Placeholder: "You must choose",
			},
			{
				DisplayName: "Optional Boolean",
				Name:        "optional_bool",
				Type:        "bool",
				Optional:    true,
				Placeholder: "Your choice",
			},
			{
				DisplayName: "Default True",
				Name:        "default_true_bool",
				Type:        "bool",
				Default:     "true",
				Placeholder: "Pre-checked",
			},
			{
				DisplayName: "Default False",
				Name:        "default_false_bool",
				Type:        "bool",
				Default:     "false",
				Placeholder: "Not pre-checked",
			},
		},
	}
}

func getDialogSelect() model.Dialog {
	return model.Dialog{
		Title:       "Select Fields",
		SubmitLabel: "Submit",
		Elements: []model.DialogElement{
			{
				DisplayName: "Radio Selector",
				Name:        "radio_field",
				Type:        "radio",
				Options: []*model.PostActionOptions{
					{Text: "Option A", Value: "optA"},
					{Text: "Option B", Value: "optB"},
					{Text: "Option C", Value: "optC"},
				},
			},
			{
				DisplayName: "Static Select",
				Name:        "static_select",
				Type:        "select",
				Placeholder: "Pick one...",
				HelpText:    "Pick one option.",
				Options: []*model.PostActionOptions{
					{Text: "Option 1", Value: "opt1"},
					{Text: "Option 2", Value: "opt2"},
					{Text: "Option 3", Value: "opt3"},
				},
			},
			{
				DisplayName: "User Selector",
				Name:        "user_select",
				Type:        "select",
				DataSource:  "users",
				Placeholder: "Choose a user...",
				HelpText:    "Select a user.",
			},
			{
				DisplayName: "Channel Selector",
				Name:        "channel_select",
				Type:        "select",
				DataSource:  "channels",
				Placeholder: "Choose a channel...",
				Optional:    true,
				HelpText:    "Select a channel.",
			},
		},
	}
}

func getDialogMultiSelect(pluginURL string) model.Dialog {
	return model.Dialog{
		Title:       "Multi-Select Fields",
		SubmitLabel: "Submit",
		Elements: []model.DialogElement{
			{
				DisplayName: "Multi-Select Options",
				Name:        "multi_options",
				Type:        "select",
				MultiSelect: true,
				Default:     "opt1,opt3",
				HelpText:    "Select multiple options.",
				Options: []*model.PostActionOptions{
					{Text: "Option 1", Value: "opt1"},
					{Text: "Option 2", Value: "opt2"},
					{Text: "Option 3", Value: "opt3"},
					{Text: "Option 4", Value: "opt4"},
				},
			},
			{
				DisplayName: "Multi-Select Users",
				Name:        "multi_users",
				Type:        "select",
				DataSource:  "users",
				MultiSelect: true,
				Placeholder: "Choose users...",
				HelpText:    "Select multiple users.",
				Optional:    true,
			},
			{
				DisplayName:   "Dynamic Role Search",
				Name:          "dynamic_roles",
				Type:          "select",
				DataSource:    "dynamic",
				DataSourceURL: fmt.Sprintf("%s/dialog/roles", pluginURL),
				Placeholder:   "Search roles...",
				HelpText:      "Dynamic search from plugin endpoint.",
			},
		},
	}
}

func getDialogDate() model.Dialog {
	return model.Dialog{
		Title:       "Date Fields",
		SubmitLabel: "Submit",
		Elements: []model.DialogElement{
			{
				DisplayName: "Basic Date",
				Name:        "basic_date",
				Type:        "date",
				HelpText:    "Pick any date.",
				Optional:    true,
			},
			{
				DisplayName: "Future Date Only",
				Name:        "future_date",
				Type:        "date",
				HelpText:    "Only future dates allowed.",
			},
			{
				DisplayName: "Default Today",
				Name:        "default_date",
				Type:        "date",
				Default:     "today",
				HelpText:    "Defaults to today.",
				Optional:    true,
			},
		},
	}
}

func getDialogDateTime() model.Dialog {
	return model.Dialog{
		Title:       "DateTime Fields",
		SubmitLabel: "Submit",
		Elements: []model.DialogElement{
			{
				DisplayName: "Basic DateTime",
				Name:        "basic_datetime",
				Type:        "datetime",
				DateTimeConfig: &model.DialogDateTimeConfig{
					TimeInterval: 30,
				},
				HelpText: "30-minute intervals.",
				Optional: true,
			},
			{
				DisplayName: "Chicago DateTime",
				Name:        "tz_datetime",
				Type:        "datetime",
				DateTimeConfig: &model.DialogDateTimeConfig{
					TimeInterval:     60,
					LocationTimezone: "America/Chicago",
				},
				HelpText: "Times shown in America/Chicago timezone.",
			},
			{
				DisplayName: "London Manual Time",
				Name:        "london_manual_datetime",
				Type:        "datetime",
				DateTimeConfig: &model.DialogDateTimeConfig{
					TimeInterval:         30,
					LocationTimezone:     "Europe/London",
					AllowManualTimeEntry: true,
				},
				HelpText: "London timezone with manual entry.",
			},
		},
	}
}

func getDialogDateTimeBasic() model.Dialog {
	return model.Dialog{
		CallbackId: "datetime_basic",
		Title:      "Date & DateTime Basics",
		Elements: []model.DialogElement{
			{
				DisplayName: "Event Date",
				Name:        "event_date",
				Type:        "date",
				HelpText:    "Select the date for your event",
				Placeholder: "Select a date",
			},
			{
				DisplayName:  "Meeting Time",
				Name:         "meeting_time",
				Type:         "datetime",
				HelpText:     "Select the date and time for your meeting",
				Placeholder:  "Select date and time",
				TimeInterval: 60,
			},
			{
				DisplayName: "Future Date Only",
				Name:        "future_date",
				Type:        "date",
				HelpText:    "Must be today or later",
				Placeholder: "Select a future date",
				MinDate:     "today",
				Optional:    true,
			},
			{
				DisplayName:  "Custom Interval Time",
				Name:         "interval_time",
				Type:         "datetime",
				HelpText:     "Time picker with 30-minute intervals",
				Placeholder:  "Select time (30min intervals)",
				TimeInterval: 30,
				Optional:     true,
			},
			{
				DisplayName: "Relative Date Example",
				Name:        "relative_date",
				Type:        "date",
				HelpText:    "Defaults to today using relative date",
				Placeholder: "Today by default",
				Default:     "today",
				Optional:    true,
			},
			{
				DisplayName: "Relative DateTime Example",
				Name:        "relative_datetime",
				Type:        "datetime",
				HelpText:    "Defaults to tomorrow using relative date",
				Placeholder: "Tomorrow by default",
				Default:     "+1d",
				Optional:    true,
			},
		},
		SubmitLabel:    "Submit",
		NotifyOnCancel: true,
		State:          "datetime_basic",
		IntroductionText: "**Date & DateTime Basics Demo**\n\n" +
			"This dialog demonstrates core date and datetime functionality:\n" +
			"- Basic date and datetime fields\n" +
			"- Min date constraints\n" +
			"- Custom time intervals\n" +
			"- Relative date defaults",
	}
}

func getDialogDateTimeTimezone() model.Dialog {
	return model.Dialog{
		CallbackId: "datetime_timezone",
		Title:      "Timezone & Manual Entry Demo",
		Elements: []model.DialogElement{
			{
				DisplayName: "Your Local Time (Manual Entry)",
				Name:        "local_manual",
				Type:        "datetime",
				HelpText:    "Type any time: 9am, 14:30, 3:45pm - no rounding",
				DateTimeConfig: &model.DialogDateTimeConfig{
					AllowManualTimeEntry: true,
				},
				Optional: true,
			},
			{
				DisplayName: "London Office Hours (Dropdown)",
				Name:        "london_dropdown",
				Type:        "datetime",
				HelpText:    "Times shown in GMT - select from 60 min intervals",
				DateTimeConfig: &model.DialogDateTimeConfig{
					LocationTimezone: "Europe/London",
					TimeInterval:    60,
				},
				Optional: true,
			},
			{
				DisplayName: "London Office Hours (Manual Entry)",
				Name:        "london_manual",
				Type:        "datetime",
				HelpText:    "Type time in GMT: 9am, 14:30, 3:45pm - no rounding",
				DateTimeConfig: &model.DialogDateTimeConfig{
					LocationTimezone:     "Europe/London",
					AllowManualTimeEntry: true,
				},
				Optional: true,
			},
		},
		SubmitLabel:    "Submit",
		NotifyOnCancel: true,
		State:          "datetime_timezone",
		IntroductionText: "**Timezone & Manual Entry Demo**\n\n" +
			"This dialog demonstrates timezone support and manual time entry features.",
	}
}

func getDialogWithoutElements() model.Dialog {
	return model.Dialog{
		Title:            "Confirmation Dialog",
		SubmitLabel:      "Confirm",
		IntroductionText: "Are you sure you want to proceed?",
		NotifyOnCancel:   true,
		Elements:         nil,
	}
}

func getDialogWithFieldRefresh(pluginURL string) model.Dialog {
	return model.Dialog{
		Title:       "Field Refresh Dialog",
		SourceURL:   fmt.Sprintf("%s/dialog/field-refresh", pluginURL),
		SubmitLabel: "Submit",
		Elements: []model.DialogElement{
			{
				DisplayName: "Category",
				Name:        "category",
				Type:        "select",
				Refresh:     true,
				HelpText:    "Selecting changes available options below.",
				Options: []*model.PostActionOptions{
					{Text: "General", Value: "general"},
					{Text: "Advanced", Value: "advanced"},
				},
			},
			{
				DisplayName: "Details",
				Name:        "details",
				Type:        "text",
				Optional:    true,
				HelpText:    "Additional details based on category.",
			},
		},
	}
}

func getDialogStep1() model.Dialog {
	return model.Dialog{
		Title:       "Registration Step 1",
		CallbackId:  "step1",
		SubmitLabel: "Next",
		Elements: []model.DialogElement{
			{
				DisplayName: "Full Name",
				Name:        "fullname",
				Type:        "text",
				Placeholder: "First and last name",
				HelpText:    "Enter your full name.",
			},
			{
				DisplayName: "Email",
				Name:        "email",
				Type:        "text",
				SubType:     "email",
				Placeholder: "user@example.com",
			},
		},
	}
}

func getDialogStep2(state string) model.Dialog {
	return model.Dialog{
		Title:       "Registration Step 2",
		CallbackId:  "step2",
		State:       state,
		SubmitLabel: "Next",
		Elements: []model.DialogElement{
			{
				DisplayName: "Role",
				Name:        "role",
				Type:        "select",
				Options: []*model.PostActionOptions{
					{Text: "Developer", Value: "developer"},
					{Text: "Designer", Value: "designer"},
					{Text: "Manager", Value: "manager"},
				},
			},
			{
				DisplayName: "Team",
				Name:        "team",
				Type:        "select",
				DataSource:  "channels",
				Placeholder: "Select team channel...",
				Optional:    true,
			},
		},
	}
}

func getDialogStep3Summary(state string) model.Dialog {
	return model.Dialog{
		Title:            "Registration Step 3",
		CallbackId:       "step3",
		State:            state,
		SubmitLabel:      "Complete",
		IntroductionText: "Review and complete your registration.",
		Elements: []model.DialogElement{
			{
				DisplayName: "Bio",
				Name:        "bio",
				Type:        "textarea",
				Optional:    true,
				Placeholder: "Tell us about yourself...",
				MaxLength:   500,
			},
			{
				DisplayName: "Receive Updates",
				Name:        "updates",
				Type:        "bool",
				Default:     "true",
				Placeholder: "Email me updates",
			},
		},
	}
}
