package session

import (
	"time"

	"github.com/twstrike/otr3"
)

// EventHandler represents the main notifications that the session can emit
// It's really more an observer than an even handler
type EventHandler interface {
	Debug(string)
	Info(string)
	Warn(string)
	Alert(string)
	NewOTRKeys(from string, conversation *otr3.Conversation)
	MessageReceived(s *Session, from string, timestamp time.Time, encrypted bool, message []byte)
	ProcessPresence(from, to, show, status string, gone bool)
	SubscriptionRequest(s *Session, uid string)
	Subscribed(account, peer string)
	Unsubscribe(account, peer string)
	RegisterCallback(title, instructions string, fields []interface{}) error
}
