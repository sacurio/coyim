package gui

import (
	"fmt"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/coyim/i18n"
	"github.com/twstrike/coyim/ui"
)

type conversationWindow struct {
	to            string
	account       *account
	win           *gtk.Window
	history       *gtk.TextView
	scrollHistory *gtk.ScrolledWindow
}

func (conv *conversationWindow) conversationMenu() *gtk.MenuBar {
	menubar, _ := gtk.MenuBarNew()
	conversationMenu, _ := gtk.MenuItemNewWithMnemonic(i18n.Local("Conversation"))
	menubar.Append(conversationMenu)

	submenu, _ := gtk.MenuNew()
	conversationMenu.SetSubmenu(submenu)

	startAKE, _ := gtk.MenuItemNewWithMnemonic(i18n.Local("Start encrypted chat"))
	submenu.Append(startAKE)

	//TODO: enable/disable depending on the conversation's encryption state
	startAKE.Connect("activate", func() {
		//TODO: errors
		err := conv.account.StartEncryptedChatWith(conv.to)
		if err != nil {
			fmt.Printf(i18n.Local("Failed to start the encrypted chat: %s\n"), err.Error())
		}
	})

	//TODO: enable/disable depending on the conversation's encryption state
	endOTR, _ := gtk.MenuItemNewWithMnemonic(i18n.Local("End encrypted chat"))
	submenu.Append(endOTR)

	endOTR.Connect("activate", func() {
		//TODO: errors
		err := conv.account.TerminateConversationWith(conv.to)
		if err != nil {
			fmt.Printf(i18n.Local("Failed to terminate the encrypted chat: %s\n"), err.Error())
		}
	})

	verify, _ := gtk.MenuItemNewWithMnemonic(i18n.Local("_Verify fingerprint..."))
	submenu.Append(verify)

	verify.Connect("activate", func() {
		verifyFingerprintDialog(conv.account, conv.to)
	})

	return menubar
}

func newConversationWindow(account *account, uid string) *conversationWindow {
	win, _ := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	history, _ := gtk.TextViewNew()
	scrollHistory, _ := gtk.ScrolledWindowNew(nil, nil)

	conv := &conversationWindow{
		to:            uid,
		account:       account,
		win:           win,
		history:       history,
		scrollHistory: scrollHistory,
	}

	// Unlike the GTK version, this is not supposed to be used as a callback but
	// it attaches the callback to the widget
	conv.win.HideOnDelete()

	conv.win.SetPosition(gtk.WIN_POS_CENTER)
	conv.win.SetDefaultSize(500, 440)
	conv.win.SetDestroyWithParent(true)
	conv.win.SetTitle(uid)

	//TODO: Load recent messages
	conv.history.SetWrapMode(gtk.WRAP_WORD)
	conv.history.SetEditable(false)
	conv.history.SetCursorVisible(false)

	vbox, _ := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 1)

	vbox.SetHomogeneous(false)
	vbox.SetBorderWidth(3)

	text, _ := gtk.EntryNew()
	text.Connect("activate", func() {
		text.SetEditable(false)

		msg, _ := text.GetText()
		text.SetText("")

		text.SetEditable(true)

		err := conv.sendMessage(msg)
		if err != nil {
			fmt.Printf(i18n.Local("Failed to generate OTR message: %s\n"), err.Error())
		}
	})

	conv.scrollHistory.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
	conv.scrollHistory.Add(conv.history)

	vbox.PackStart(conv.conversationMenu(), false, false, 0)
	vbox.PackEnd(text, false, false, 0)
	vbox.PackStart(conv.scrollHistory, true, true, 0)

	conv.win.Add(vbox)

	text.GrabFocus()

	return conv
}

func (conv *conversationWindow) Hide() {
	conv.win.Hide()
}

func (conv *conversationWindow) Show() {
	conv.win.ShowAll()
}

func (conv *conversationWindow) sendMessage(message string) error {
	err := conv.account.EncryptAndSendTo(conv.to, message)
	if err != nil {
		return err
	}

	//TODO: this should not be in both GUI and roster
	conversation := conv.account.GetConversationWith(conv.to)
	encrypted := conversation.IsEncrypted()
	glib.IdleAdd(func() bool {
		conv.appendMessage(conv.account.Config.Account, time.Now(), encrypted, ui.StripHTML([]byte(message)))
		return false
	})

	return nil
}

const timeDisplay = "15:04:05"

func (conv *conversationWindow) appendMessage(from string, timestamp time.Time, encrypted bool, message []byte) {
	glib.IdleAdd(func() bool {
		buff, _ := conv.history.GetBuffer()
		buff.InsertAtCursor("[")
		buff.InsertAtCursor(timestamp.Format(timeDisplay))
		buff.InsertAtCursor("] ")
		buff.InsertAtCursor(from)
		buff.InsertAtCursor(":  ")
		buff.InsertAtCursor(string(message))
		buff.InsertAtCursor("\n")

		adj := conv.scrollHistory.GetVAdjustment()
		adj.SetValue(adj.GetUpper())

		return false
	})
}
