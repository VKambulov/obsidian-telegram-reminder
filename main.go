package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var timezone *time.Location
var template string

var dateRegex = regexp.MustCompile(`@(\d{4}-\d{2}-\d{2} \d{2}:\d{2}|\d{4}-\d{2}-\d{2})`)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	log.SetOutput(file)

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	timezone, err = time.LoadLocation(os.Getenv("TIMEZONE"))
	if err != nil {
		log.Fatal("Error loading timezone:", err)
	}

	content, err := os.ReadFile(os.Getenv("MESSAGE_TEMPLATE_PATH"))
	if err != nil {
		log.Print("Error loading template or template not found: ", err)
		template = "Obsidian Reminder\n{{filename}}\n{{datetime}}\n\n{{message}}"
	} else {
		template = string(content)
	}

	checkMarkdownFiles()
}

func checkMarkdownFiles() {
	err := filepath.Walk(
		os.Getenv("OBSIDIAN_VAULT_PATH"),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.HasSuffix(info.Name(), ".md") {
				return processMarkdownFile(path)
			}

			return nil
		})
	if err != nil {
		log.Fatal("Error walking the path:", err)
	}
}

func processMarkdownFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if matches := dateRegex.FindStringSubmatch(line); matches != nil {
			dateStr := matches[1]

			if !strings.Contains(dateStr, ":") {
				remindTime := os.Getenv("REMIND_TIME")

				if remindTime == "" {
					remindTime = "09:00"
				}

				dateStr = dateStr + " " + remindTime
			}

			if date, err := time.ParseInLocation("2006-01-02 15:04", dateStr, timezone); err == nil {
				now := time.Now()
				if date.After(now.Add(-5*time.Minute)) && date.Before(time.Now().In(timezone)) {
					log.Print("Found reminder in markdown with path: ", path)
					sendTelegramReminder(date, line, path)
				}
			}
		}
	}

	return nil
}

func sendTelegramReminder(date time.Time, message string, path string) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal("Error creating Telegram bot:", err)
	}

	chatId, err := strconv.Atoi(os.Getenv("TELEGRAM_CHAT_ID"))
	if err != nil {
		log.Fatal("Error convert Telegram Chat ID to int:", err)
	}

	datetime := date.Format("2006-01-02 15:04")

	reminderMessage := strings.ReplaceAll(template, "{{datetime}}", datetime)
	reminderMessage = strings.ReplaceAll(reminderMessage, "{{filename}}", strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))
	reminderMessage = strings.ReplaceAll(reminderMessage, "{{message}}", message)

	msg := tgbotapi.NewMessage(int64(chatId), reminderMessage)
	if _, err := bot.Send(msg); err != nil {
		log.Fatal("Error sending Telegram message:", err)
	}
}
