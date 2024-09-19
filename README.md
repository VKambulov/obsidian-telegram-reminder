# Obsidian Telegram Reminder

A simple service for reminders in Telegram from Obsidian notes.
This service must be running on a server or PC that will be running continuously.

You can also use this service for other programs that work with markdown files, such as Logseq.

Currently, only the 24-hour format is supported.

Tested on Windows and Linux.

[Описание на русском языке](README-RU.md)

## How it's working

If you use Syncthing to sync Obsidian Vault between devices,
or if you have a PC that will run continuously, you may use this service.

The service checks your vault every 5 minutes using a cron scheduler task and looks for this patterns
in markdown files: `@2024-09-01 13:00`, `@2024-09-01`.

## Quick start

1. Download the latest release for your OS.
2. Place the binary file in the directory from which the scheduled task will be launched. For example: `/opt/obsidian-telegram-reminder`.
3. Place a `.env` file in the directory and fill it out.
4. Create a cron scheduled task in your system. The task must run every 5 minutes. For example: `*/5 * * * * /opt/obsidian-telegram-reminder/obsidian-telegram-reminder`.
5. Create a new note or write in an exiting note this text (replace date and time with your current): `Remind me this! @2024-09-01 20:00`
6. If you have set everything up correctly, in 5 minutes you will receive a message in the Telegram bot!

## Env variables description

### `OBSIDIAN_VAULT_PATH`

Full path to your Obsidian Vault.

Ex. `"/home/user/syncthing/My Vault"`

### `MESSAGE_TEMPLATE_PATH`

Template for your Telegram messages.

You can create an Obsidian note in your vault and set its path to this variable.

Example template:

```text
Obsidian Reminder

File name: {{filename}}

Reminder date: {{datetime}}

{{message}}
```

Markdown is not supported correctly in Telegram, so its template must be in plain text format.

If you do not set a template, the following template will be used by default:

```text
Obsidian Reminder
{{filename}}
{{datetime}}

{{message}}
```

### `TIMEZONE`

The timezone of your reminder notes.

Your server may have a different timezone, so you need to set it.

### `REMIND_TIME`

If you do not want to specify the notification time each time,
you can use a notification template without a time: `@2024-09-01`.
In this case, the notification will be sent by default at 9:00 in your timezone.
With this variable, you can override the notification time; to do this,
specify the time in the `HH:MM` format.

### `TELEGRAM_BOT_TOKEN`

Your Telegram Bot API key from [@BotFather](https://t.me/BotFather) bot.

### `TELEGRAM_CHAT_ID`

Telegram Chat ID.

You can get it from the [@getmyid_bot](https://t.me/getmyid_bot) bot.

## Building from source

To build from source code run this command:

```bash
go build
```
