package notify

import (
	"fmt"
	"strings"
)

type Channel uint32

const (
	ChannelEmail Channel = 1 << iota // 0000 ... 0000 0001
	ChannelPush                      // 0000 ... 0000 0010
	ChannelSms                       // 0000 ... 0000 0100
)

func ParseChannels(input string) (Channel, error) {
	if strings.TrimSpace(input) == "" {
		return 0, nil
	}

	var resultChannel Channel

	for _, strChannel := range strings.Split(input, ",") {
		channel, err := parseChannel(strings.TrimSpace(strChannel))
		if err != nil {
			return 0, err
		}

		resultChannel |= channel
	}

	return resultChannel, nil
}

func parseChannel(input string) (Channel, error) {
	switch input {
	case "email":
		return ChannelEmail, nil
	case "push":
		return ChannelPush, nil
	case "sms":
		return ChannelSms, nil
	case "":
		return 0, nil
	default:
		return 0, fmt.Errorf("Channel %q not found", input)
	}
}

func (channel Channel) Has(anotherChannel Channel) bool {
	return anotherChannel&channel == anotherChannel
}
