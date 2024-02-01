#!/bin/bash

# I do not want my screensaver to auto-activate after a timeout.
# For some reason, you *have* to select a nonzero, finite timeout value for xscreensaver (there's no "infinity").
# I only want my screensaver to active when I lock my system by pressing [Meta]+[L].
# This hotkey is set up to launch this script, which 
# - starts the xscreensaver daemon,
# - activates the screensaver (with lockscreen enabled),
# - waits for it to terminate (unlocking logs an "UNBLANK" event),
# - then kills the xscreensaver daemon.

xscreensaver -no-splash &
sleep .1
xscreensaver-command -activate &
sleep .1
xscreensaver-command -watch | { grep -m 1 UNBLANK && killall xscreensaver-co ; }
pkill xscreensaver
