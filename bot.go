package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const API_URL = `https://api.unsplash.com/`

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load("config.env"); err != nil {
		log.Println("No file named config.env found")
	}
}

func main() {
	var botToken = os.Getenv("API_KEY")
	var unsplashAccess = os.Getenv("ACCESS_KEY")
	var unsplashSecret = os.Getenv("SECRET_KEY")

	// Checks if all variables are present
	if botToken == "" && unsplashAccess == "" && unsplashSecret == "" {
		log.Println("One or more variables missing in config")
	}

	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeLevel = zapcore.CapitalLevelEncoder
	cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), os.Stdout, zap.InfoLevel))
	defer logger.Sync()
	l := logger.Sugar()

	l.Info("Starting gotgbot...")
	ubot, err := gotgbot.NewUpdater(botToken, logger)
	if err != nil {
		l.Fatalw("Failed to start updater", zap.Error(err))
	}

	// Reply to /start messages
	ubot.Dispatcher.AddHandler(handlers.NewArgsCommand("start", startHandler))

	// Reply to /random messages
	ubot.Dispatcher.AddHandler(handlers.NewCommand("random", randomHandler))

	ubot.StartPolling()
	// wait
	ubot.Idle()
}

func startHandler(_ ext.Bot, u *gotgbot.Update, args []string) error {
	msg := u.EffectiveMessage
	_, err := msg.ReplyTextf("Hi there! I'm a telegram bot, written in Go and based on Unsplash's API")
	return err
}

func randomHandler(b ext.Bot, u *gotgbot.Update) error {
	unsplash := random()
	caption := fmt.Sprintf("ID: %s\nBy %s\nLink: %s", unsplash.ID, unsplash.User.Name, unsplash.Links.HTML)
	_, err := b.ReplyPhotoCaptionStr(u.EffectiveChat.Id, unsplash.Urls.Thumb, caption, u.EffectiveMessage.MessageId)
	if err != nil {
		b.Logger.Warnw("Error sending V2", zap.Error(err))
	}
	return nil
}
