package gui

import "github.com/twstrike/go-gtk/gtk"

type checkButton struct {
	text   string
	active bool
	id     string
}

func (wr *widgetRegistry) getActive(id string) bool {
	w := wr.reg[id]

	switch w := w.(type) {
	case *gtk.CheckButton:
		return w.GetActive()
	}

	return false
}

func (c checkButton) create(reg *widgetRegistry) gtk.IWidget {
	entry := gtk.NewCheckButtonWithLabel(c.text)
	entry.SetActive(c.active)
	reg.register(c.id, entry)

	return entry
}
