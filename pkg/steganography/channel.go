package steganography

import (
	"strings"

	"github.com/stegoer/server/ent/image"
)

const (
	all       = "RED_GREEN_BLUE"
	separator = "_"
)

func IncludesRedChannel(ch image.Channel) bool {
	switch ch.String() {
	case "RED", "RED_GREEN", "RED_BLUE", all:
		return true
	}

	return false
}

func IncludesGreenChannel(ch image.Channel) bool {
	switch ch.String() {
	case "GREEN", "RED_GREEN", "GREEN_BLUE", all:
		return true
	}

	return false
}

func IncludesBlueChannel(ch image.Channel) bool {
	switch ch.String() {
	case "BLUE", "RED_BLUE", "GREEN_BLUE", all:
		return true
	}

	return false
}

func ChannelCount(ch image.Channel) int {
	return len(strings.Split(ch.String(), separator))
}
