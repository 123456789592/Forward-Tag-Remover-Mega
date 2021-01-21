
/*  Copyright (C) 2020 by 123456789592@Github, < https://github.com/123456789592 >.
 *
 * This file is part of < https://github.com/123456789592/Forward-Tag-Remover-Mega > project,
 * and is released under the "GNU v3.0 License Agreement".
 * Please see < https://github.com/123456789592/Forward-Tag-Remover-Mega/blob/master/LICENSE >
 *
 * All rights reserved.
 */

package main

import (
	"fmt"
	"os"

	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/handlers"
	"github.com/PaulSonOfLars/gotgbot/handlers/Filters"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/123456789592/Forward-Tag-Remover-Mega/captions"
	"github.com/123456789592/Forward-Tag-Remover-Mega/commands"
	"github.com/123456789592/Forward-Tag-Remover-Mega/functions"
)

func main() {
	cfg := zap.NewProductionEncoderConfig()

	cfg.EncodeLevel = zapcore.CapitalLevelEncoder

	cfg.EncodeTime = zapcore.RFC3339TimeEncoder

	logger := zap.New(zapcore.NewCore(zapcore.NewConsoleEncoder(cfg), os.Stdout, zap.InfoLevel))

	defer logger.Sync()

	l := logger.Sugar()

	token := os.Getenv("BOT_TOKEN")
	u, err := gotgbot.NewUpdater(logger, token)
	if err != nil {
		l.Fatalw("Updater failed starting", zap.Error(err))
		return
	}

	bot := u.Bot.FirstName
	fmt.Printf("Successfully logged as %s", bot)

	u.Dispatcher.AddHandler(handlers.NewCommand("start", commands.Start))
	u.Dispatcher.AddHandler(handlers.NewCommand("help", commands.Help))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Text, captions.SetCaption))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Document, functions.ForwardDocument))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Video, functions.ForwardVideo))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Photo, functions.ForwardPhoto))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Voice, functions.ForwardVoice))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Audio, functions.ForwardAudio))
	u.Dispatcher.AddHandler(handlers.NewMessage(Filters.Sticker, functions.ForwardSticker))

	err = u.StartPolling()
	if err != nil {
		l.Fatalw("Polling failed", zap.Error(err))
		return
	}
	u.Idle()

}
