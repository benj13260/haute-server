package logz

import (
	"strings"
	"time"

	"github.com/slack-go/slack"
	"go.uber.org/zap/zapcore"
)

const (
	ChannelSlowQueries = "zlogz_slow_queries"
	ChannelProcessing  = "zlogz_processing"
)

// SlackHook is a zap Hook for dispatching messages to the specified
// channel on Slack.
type SlackHook struct {
	// Messages with a log level not contained in this array
	// will not be dispatched. If nil, all messages will be dispatched.
	AcceptedLevels []zapcore.Level
	HookURL        string // Webhook URL
	ChannelSlow    string

	// slack post parameters
	Username  string // display name
	Channel   string // `#channel-name`
	IconEmoji string // emoji string ex) ":ghost:":
	IconURL   string // icon url

	FieldHeader string        // a header above field data
	Timeout     time.Duration // request timeout
	Async       bool          // if async is true, send a message asynchronously.

}

func NewSlackHook(hookURL, channel, channelSlow, username string, level *zapcore.Level) *SlackHook {
	rt := &SlackHook{
		HookURL:     hookURL,
		Username:    username,
		Channel:     channel,
		ChannelSlow: channelSlow,
		IconEmoji:   ":parrot:",
	}
	if level != nil {
		var levelIndex = 0
		rt.AcceptedLevels = []zapcore.Level{}
		for i := range AllLevels {
			if AllLevels[i] == *level {
				levelIndex = i
			}
		}
		for i := levelIndex; i < len(AllLevels); i++ {
			rt.AcceptedLevels = append(rt.AcceptedLevels, AllLevels[i])
		}
	}
	return rt
}

func (sh *SlackHook) channel(loggerName string) string {
	switch {
	case strings.Contains(loggerName, "Query_trace"):
		return sh.ChannelSlow
	default:
		return sh.Channel
	}
}

func (sh *SlackHook) GetHook() func(zapcore.Entry) error {
	return func(e zapcore.Entry) error {
		if !strings.Contains(e.LoggerName, "Query_trace") &&
			!sh.isAcceptedLevel(e.Level) {
			return nil
		}

		payload := &slack.WebhookMessage{
			Username:  sh.Username,
			Channel:   sh.channel(e.LoggerName),
			IconEmoji: sh.IconEmoji,
			IconURL:   sh.IconURL,
		}

		color := LevelColorMap[e.Level]
		attachment := slack.Attachment{
			Title: e.Level.CapitalString() + "_" + e.Caller.TrimmedPath(),
			Text:  e.Message,
			Color: color,
			// AuthorName: ,
		}
		payload.Attachments = []slack.Attachment{attachment}
		attachment.Color = color

		return sh.postMessage(payload)
	}
}

func (sh *SlackHook) postMessage(payload *slack.WebhookMessage) error {
	return slack.PostWebhook(sh.HookURL, payload)
}

// Levels sets which levels to sent to slack
func (sh *SlackHook) Levels() []zapcore.Level {
	if sh.AcceptedLevels == nil {
		return AllLevels
	}
	return sh.AcceptedLevels
}

func (sh *SlackHook) isAcceptedLevel(level zapcore.Level) bool {
	for _, lv := range sh.Levels() {
		if lv == level {
			return true
		}
	}
	return false
}

// Supported log levels
var AllLevels = []zapcore.Level{
	zapcore.DebugLevel,
	zapcore.InfoLevel,
	zapcore.WarnLevel,
	zapcore.ErrorLevel,
	zapcore.DPanicLevel,
	zapcore.PanicLevel,
	zapcore.FatalLevel,
}

var LevelColorMap = map[zapcore.Level]string{
	zapcore.DebugLevel:  "#616161",
	zapcore.InfoLevel:   "#1976D2",
	zapcore.WarnLevel:   "#FBC02D",
	zapcore.ErrorLevel:  "#D32F2F",
	zapcore.DPanicLevel: "#7B1FA2",
	zapcore.PanicLevel:  "#7B1FA2",
	zapcore.FatalLevel:  "#7B1FA2",
}

// LevelThreshold - Returns every logging level above and including the given parameter.
func LevelThreshold(l zapcore.Level) []zapcore.Level {
	for i := range AllLevels {
		if AllLevels[i] == l {
			return AllLevels[i:]
		}
	}
	return []zapcore.Level{}
}
