Emote in Zoom with Keyboard Shortcuts
-------------------------------------

This program:

	- finds a window with "Zoom Meeting" in its title
	- looks for the 'Reactions' icon at the bottom of the screen
	- clicks the 'Reactions' icon
	- depending on the arguments, clicks one of the emoji icons

The buttons at the bottom must be visible, press the Alt key to toggle the
buttons.

Arguments you can pass to this program are:

	clap
	thumbs up
	cry laughing
	open mouth
	heart
	party
	yes
	no
	slow down
	speed up
	away
	raise hand

which correspond to the respective emoji. Pass these words directly, no '-' or
'--'.

Installation
------------

	go install github.com/gonutz/zoom_emote@latest

On Windows

	- Go to your %GOBIN% path and find this program.
	- Create links to this program on your Desktop.
	- Go to a link's Properties.
	- In the 'Target' field, add arguments at the end (e.g. "thumbs up").
	- In the field 'Shortcut key' press e.g. U to create shortcut Ctrl+Alt+U.

You can create multiple links for different emojis and give each a custom
shortcut.

Caveat Emptor
-------------

The program only works on Windows.

This program may or may not work on your machine, it was only tested on my
machine. The screenshot of the 'Reactions' icon was taken at 100% scale, some
Windows machines come with 125% scale which might cause problems.

Zoom might have updated their icons between this program's last commit and you
using it.

You might have different settings on your machine of where your icons are, this
program assumes they are at the bottom of the Zoom window.
