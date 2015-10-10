package gui

import "github.com/twstrike/go-gtk/gtk"

// This is how the gtkrc file should look like on Mac OS
// https://git.gnome.org/browse/gtk+/commit/?id=4ff709c24b8d4b3e26b3d513fde0676e9c43f897
func (u *gtkUI) applyStyle() {
	style := `
gtk-enable-mnemonics = 0

binding "gtk-mac-alt-arrows"
{
  bind "<alt>Right"           { "move-cursor" (words, 1, 0) }
  bind "<alt>KP_Right"        { "move-cursor" (words, 1, 0) }
  bind "<alt>Left"            { "move-cursor" (words, -1, 0) }
  bind "<alt>KP_Left"         { "move-cursor" (words, -1, 0) }
  bind "<shift><alt>Right"    { "move-cursor" (words, 1, 1) }
  bind "<shift><alt>KP_Right" { "move-cursor" (words, 1, 1) }
  bind "<shift><alt>Left"     { "move-cursor" (words, -1, 1) }
  bind "<shift><alt>KP_Left"  { "move-cursor" (words, -1, 1) }
}

class "GtkTextView" binding "gtk-mac-alt-arrows"
class "GtkLabel" binding "gtk-mac-alt-arrows"
class "GtkEntry" binding "gtk-mac-alt-arrows"


binding "gtk-mac-alt-delete"
{
  bind "<alt>Delete" { "delete-from-cursor" (word-ends, 1) }
  bind "<alt>KP_Delete" { "delete-from-cursor" (word-ends, 1) }
  bind "<alt>BackSpace" { "delete-from-cursor" (word-ends, -1) }
}

class "GtkTextView" binding "gtk-mac-alt-delete"
class "GtkEntry" binding "gtk-mac-alt-delete"


binding "gtk-mac-cmd-c"
{
  bind "<meta>x" { "cut-clipboard" () }
  bind "<meta>c" { "copy-clipboard" () }
  bind "<meta>v" { "paste-clipboard" () }
  unbind "<ctrl>x"
  unbind "<ctrl>c"
  unbind "<ctrl>v"
}

class "GtkTextView" binding "gtk-mac-cmd-c"
class "GtkEntry" binding "gtk-mac-cmd-c"


binding "gtk-mac-text-view"
{
  bind "<shift><meta>a" { "select-all" (0) }
  bind "<meta>a" { "select-all" (1) }
  unbind "<shift><ctrl>a"
  unbind "<ctrl>a"
}

class "GtkTextView" binding "gtk-mac-text-view"


binding "gtk-mac-label"
{
  bind "<meta>a" {
    "move-cursor" (paragraph-ends, -1, 0)
    "move-cursor" (paragraph-ends, 1, 1)
  }
  bind "<shift><meta>a" { "move-cursor" (paragraph-ends, 0, 0) }
  bind "<meta>c" { "copy-clipboard" () }
  unbind "<ctrl>a"
  unbind "<shift><ctrl>a"
  unbind "<ctrl>c"
}

class "GtkLabel" binding "gtk-mac-label"


binding "gtk-mac-entry"
{
  bind "<meta>a" {
    "move-cursor" (buffer-ends, -1, 0)
    "move-cursor" (buffer-ends, 1, 1)
  }
  bind "<shift><meta>a" { "move-cursor" (visual-positions, 0, 0) }
  unbind "<ctrl>a"
  unbind "<shift><ctrl>a"
}

class "GtkEntry" binding "gtk-mac-entry"


binding "gtk-mac-cmd-arrows"
{
  bind "<meta>Left" { "move-cursor" (paragraph-ends, -1, 0) }
  bind "<meta>KP_Left" { "move-cursor" (paragraph-ends, -1, 0) }
  bind "<shift><meta>Left" { "move-cursor" (paragraph-ends, -1, 1) }
  bind "<shift><meta>KP_Left" { "move-cursor" (paragraph-ends, -1, 1) }
  bind "<meta>Right" { "move-cursor" (paragraph-ends, 1, 0) }
  bind "<meta>KP_Right" { "move-cursor" (paragraph-ends, 1, 0) }
  bind "<shift><meta>Right" { "move-cursor" (paragraph-ends, 1, 1) }
  bind "<shift><meta>KP_Right" { "move-cursor" (paragraph-ends, 1, 1) }
}

class "GtkTextView" binding "gtk-mac-cmd-arrows"
class "GtkLabel" binding "gtk-mac-cmd-arrows"
class "GtkEntry" binding "gtk-mac-cmd-arrows"


binding "gtk-mac-emacs-like"
{
  bind "<ctrl>a" { "move-cursor" (paragraph-ends, -1, 0) }
  bind "<shift><ctrl>a" { "move-cursor" (paragraph-ends, -1, 1) }
  bind "<ctrl>e" { "move-cursor" (paragraph-ends, 1, 0) }
  bind "<shift><ctrl>e" { "move-cursor" (paragraph-ends, 1, 1) }

  bind "<ctrl>b" { "move-cursor" (logical-positions, -1, 0) }
  bind "<shift><ctrl>b" { "move-cursor" (logical-positions, -1, 1) }
  bind "<ctrl>f" { "move-cursor" (logical-positions, 1, 0) }
  bind "<shift><ctrl>f" { "move-cursor" (logical-positions, 1, 1) }
}

class "GtkTextView" binding "gtk-mac-emacs-like"
class "GtkLabel" binding "gtk-mac-emacs-like"
class "GtkEntry" binding "gtk-mac-emacs-like"


binding "gtk-mac-file-chooser"
{
  bind "<meta>v" { "location-popup-on-paste" () }
  unbind "<ctrl>v"

  bind "<meta><shift>G" { "location-popup" () }
  bind "<meta><shift>H" { "home-folder" () }
  bind "<meta>Up" { "up-folder" () }
}

class "GtkFileChooserDefault" binding "gtk-mac-file-chooser"


binding "gtk-mac-tree-view"
{
  bind "<meta>a" { "select-all" () }
  bind "<shift><meta>a" { "unselect-all" () }
  bind "<meta>f" { "start-interactive-search" () }
  bind "<meta>F" { "start-interactive-search" () }
  unbind "<ctrl>a"
  unbind "<shift><ctrl>a"
  unbind "<ctrl>f"
  unbind "<ctrl>F"
}

class "GtkTreeView" binding "gtk-mac-tree-view"


binding "gtk-mac-icon-view"
{
  bind "<meta>a" { "select-all" () }
  bind "<shift><meta>a" { "unselect-all" () }
  unbind "<ctrl>a"
  unbind "<shift><ctrl>a"
}

class "GtkIconView" binding "gtk-mac-icon-view"
	`

	gtk.ResourceParseString(style)
}
