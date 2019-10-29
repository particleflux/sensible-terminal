// Package sensible_terminal auto-detects installed terminal emulators
package sensible_terminal

import (
	"os/exec"
)

var candidates = [...]string{
	"x-terminal-emulator", // this is a symlink on debian based systems
	"urxvt",
	"rxvt",
	"termit",
	"terminator",
	"Eterm",
	"aterm",
	"uxterm",
	"gnome-terminal",
	"roxterm",
	"xfce4-terminal",
	"termite",
	"lxterminal",
	"mate-terminal",
	"terminology",
	"qterminal",
	"lilyterm",
	"tilix",
	"terminix",
	"konsole",
	"guake",
	"tilda",
	"alacritty",
	"xterm",
	"st",
	"kitty",
	"hyper",
}

// Detect and return all found terminal emulators
func All() ([]string, error) {
	emulators := make([]string, 0)

	for _, emulator := range candidates {
		if commandExists(emulator) {
			emulators = append(emulators, emulator)
		}
	}

	if len(emulators) == 0 {
		return nil, &NotFoundError{}
	}

	return emulators, nil
}

// Detect and return the first found terminal emulator
func First() (string, error) {
	for _, emulator := range candidates {
		if commandExists(emulator) {
			return emulator, nil
		}
	}

	return "", &NotFoundError{}
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
