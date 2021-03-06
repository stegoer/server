package util

// Channel defines the type for the "channel" enum field.
type Channel string

// Channel values.
const (
	ChannelRed          Channel = "RED"
	ChannelGreen        Channel = "GREEN"
	ChannelBlue         Channel = "BLUE"
	ChannelRedGreen     Channel = "RED_GREEN"
	ChannelRedBlue      Channel = "RED_BLUE"
	ChannelGreenBlue    Channel = "GREEN_BLUE"
	ChannelRedGreenBlue Channel = "RED_GREEN_BLUE"
)

// String returns a string representation of the Channel.
func (c Channel) String() string {
	return string(c)
}

// IncludesRed returns whether the Channel includes Channel with red value.
func (c Channel) IncludesRed() bool {
	switch c {
	case ChannelRed, ChannelRedGreen, ChannelRedBlue, ChannelRedGreenBlue:
		return true
	default:
		return false
	}
}

// IncludesGreen returns whether the Channel includes Channel with green value.
func (c Channel) IncludesGreen() bool {
	switch c {
	case ChannelGreen, ChannelRedGreen, ChannelGreenBlue, ChannelRedGreenBlue:
		return true
	default:
		return false
	}
}

// IncludesBlue returns whether the Channel includes Channel with blue value.
func (c Channel) IncludesBlue() bool {
	switch c {
	case ChannelBlue, ChannelRedBlue, ChannelGreenBlue, ChannelRedGreenBlue:
		return true
	default:
		return false
	}
}

// Count returns the number of color channels included in the Channel.
func (c Channel) Count() int {
	switch c {
	case ChannelRed, ChannelGreen, ChannelBlue:
		return 1
	case ChannelRedGreen, ChannelRedBlue, ChannelGreenBlue:
		return 2
	case ChannelRedGreenBlue:
		return 3
	default:
		// should be unreachable
		return 0
	}
}
