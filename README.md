# Mattermost Test Plugin

A Mattermost plugin that provides test infrastructure for E2E tests. It replaces the external Node.js webhook server and supplements the demo plugin with consolidated interactive dialog testing and channel header button testing.

## Features

### Channel Header Buttons

Registers 15 channel header button icons on activation. Used by the Cypress E2E test MM-T1649 to verify that 16+ plugin icons collapse into a single dropdown.

### `/e2e-dialog` Slash Command

Opens interactive test dialogs covering all dialog element types. Subcommands:

| Subcommand | Description |
|------------|-------------|
| *(none)* | Full dialog with one of every element type |
| `text` | Text variants: plain, email, number, password, textarea |
| `boolean` | Required, optional, default-true, default-false |
| `select` | Radio, static select, user selector, channel selector |
| `multi-select` | Multi-select options, multi-select users, dynamic search |
| `date` | Basic date, future-only, relative default, date range |
| `datetime` | Intervals, timezones, manual entry, constrained |
| `no-elements` | Confirmation dialog (no form fields) |
| `field-refresh` | Conditional fields that update on selection change |
| `multi-step` | Three-step registration workflow |
| `error` | Dialog that returns validation errors on submit |
| `error-no-elements` | Confirmation dialog that returns errors |

### Dialog Submission Handlers

| Endpoint | Purpose |
|----------|---------|
| `/dialog/submit-with-validation` | Full dialog: validates number field equals 42 |
| `/dialog/submit-confirm` | Simple confirmation post |
| `/dialog/submit-generic` | Formats all submission fields as a post |
| `/dialog/error` | Always returns a dialog error |
| `/dialog/field-refresh` | Returns updated form based on selection |
| `/dialog/multistep` | Multi-step state machine (3 steps) |
| `/dialog/roles` | Dynamic select endpoint for role search |

## Configuration

| Setting | Type | Default | Description |
|---------|------|---------|-------------|
| `EnableDialogCommands` | bool | `true` | Enable the `/e2e-dialog` slash command and dialog handlers |

Configure in **System Console > Plugins > Test Plugin**.

## Building

```bash
make dist
```

Produces `dist/com.mattermost.test-plugin.tar.gz`.

## Deploying

```bash
make deploy
```

Requires `MM_SERVICESETTINGS_SITEURL` and either `MM_ADMIN_TOKEN` or `MM_ADMIN_USERNAME`/`MM_ADMIN_PASSWORD`.

## For E2E Test Authors

Use the `/e2e-dialog` slash command in your tests instead of the webhook server. The bot posts submission data to the channel as a `custom_e2etest_plugin` post type, which you can assert against.

### Cypress Migration

| Old trigger (webhook) | New command |
|----------------------|-------------|
| `/dialog_request` | `/e2e-dialog` |
| `/simple_dialog_request` | `/e2e-dialog no-elements` |
| `/boolean_dialog_request` | `/e2e-dialog boolean` |
| `/multiselect_dialog_request` | `/e2e-dialog multi-select` |
| `/dynamic_select_dialog_request` | `/e2e-dialog multi-select` |
| `/datetime_dialog_request` (date) | `/e2e-dialog date` |
| `/datetime_dialog_request` (datetime) | `/e2e-dialog datetime` |
| `/dialog/field-refresh` | `/e2e-dialog field-refresh` |
| `/dialog/multistep` | `/e2e-dialog multi-step` |

### Mobile Migration

| Old command | New command |
|-------------|-------------|
| `/dialog basic` | `/e2e-dialog text` |
| `/dialog error` | `/e2e-dialog error` |
| `/dialog boolean` | `/e2e-dialog boolean` |
| `/dialog selectfields` | `/e2e-dialog select` |
| `/dialog textfields` | `/e2e-dialog text` |
| `/dialog multi-select` | `/e2e-dialog multi-select` |
| `/dialog dynamic-select` | `/e2e-dialog multi-select` |
| `/dialog multistep` | `/e2e-dialog multi-step` |
| `/dialog field-refresh` | `/e2e-dialog field-refresh` |
| `/dialog datetime-basic` | `/e2e-dialog date` or `/e2e-dialog datetime` |
| `/dialog datetime-timezone` | `/e2e-dialog datetime` |

## Releasing this plugin

A new minor version of this plugin is released with every feature release of Mattermost. The new version should be cut until Code complete.

## How to Release

To trigger a release, follow these steps:

1. **For Patch Release:** Run the following command:
    ```
    make patch
    ```
   This will release a patch change.

2. **For Minor Release:** Run the following command:
    ```
    make minor
    ```
   This will release a minor change.

3. **For Major Release:** Run the following command:
    ```
    make major
    ```
   This will release a major change.

4. **For Patch Release Candidate (RC):** Run the following command:
    ```
    make patch-rc
    ```
   This will release a patch release candidate.

5. **For Minor Release Candidate (RC):** Run the following command:
    ```
    make minor-rc
    ```
   This will release a minor release candidate.

6. **For Major Release Candidate (RC):** Run the following command:
    ```
    make major-rc
    ```
   This will release a major release candidate.
