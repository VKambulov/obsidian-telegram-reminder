# Obsidian Telegram Reminder

A simple service for reminders in Telegram from Obsidian notes.
This service must be running on a server or PC that will be running continuously.

Tested on Windows and Linux.

## How it's working

If you use Syncthing to sync Obsidian Vault between your devices,
or have a PC that will run continuously, you may use this service.

The service checks your vault every 5 minutes using a cron scheduler task and looks for this pattern
in markdown files: `@2024-09-01 13:00`.

## Quick start

1. Download latest release for you OS.
2. Place the binary file in the directory from which the scheduled task will be launched. Ex. `/opt/obsidian-telegram-reminder`.
3. Place a `.env` file in the directory and fill it.
4. Create cron schedule task in your system. Task must run each 5 minutes. Ex. `*/5 * * * * /opt/obsidian-telegram-reminder/obsidian-telegram-reminder`.
5. Create new note or write in exiting note this text (replace date and time with you current): `Remind me this! @2024-09-01 20:00`
6. If you have set everything up correctly, in 5 minutes you will receive a message in Telegram bot!

## Env variables description

### `OBSIDIAN_VAULT_PATH`

Full path to you Obsidian Vault.

Ex. `"/home/user/syncthing/My Vault"`

### `MESSAGE_TEMPLATE_PATH`

Template for your Telegram messages.

You can create an Obsidian note in your vault and set its path to this variable.

Example template:

```text
Obsidian reminder

Reminder date: {{datetime}}

{{message}}
```

Markdown is not supported correctly in Telegram, so its template must be in plain text format.

If you do not set a template, the following template will be used by default:

```text
Obsidian Reminder
{{datetime}}

{{message}}
```

### `TIMEZONE`

The timezone of your reminder notes.

Your server may have a different timezone so you need to set it.

### `TELEGRAM_BOT_TOKEN`

Your Telegram Bot API key from [@BotFather](https://t.me/BotFather) bot.

### `TELEGRAM_CHAT_ID`

Telegram Chat ID.

You can get it in the [@getmyid_bot](https://t.me/getmyid_bot) bot.

## Building from source

To build from source code run this command:

```bash
go build
```
