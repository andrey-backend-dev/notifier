package notify

import "testing"

func TestParseChannelsOneChannel(t *testing.T) {
	var input = "email"
	var expected = ChannelEmail

	result, err := ParseChannels(input)

	if err != nil {
		t.Errorf("Error should not be returned: %q", err)
	}

	if expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestParseChannelsOneChannelWithDelimiter(t *testing.T) {
	var input = "email,"
	var expected = ChannelEmail

	result, err := ParseChannels(input)

	if err != nil {
		t.Errorf("Error should not be returned: %q", err)
	}

	if expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestParseChannelsOneChannelWithDelimiterAndNotTrimmed(t *testing.T) {
	var input = "email,     "
	var expected = ChannelEmail

	result, err := ParseChannels(input)

	if err != nil {
		t.Errorf("Error should not be returned: %q", err)
	}

	if expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestParseChannelsTwoChannels(t *testing.T) {
	var input = "email,push"
	var expected = Channel(3)

	result, err := ParseChannels(input)

	if err != nil {
		t.Errorf("Error should not be returned: %q", err)
	}

	if expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestParseChannelsThreeChannels(t *testing.T) {
	var input = "email,push,sms"
	var expected = Channel(7)

	result, err := ParseChannels(input)

	if err != nil {
		t.Errorf("Error should not be returned: %q", err)
	}

	if expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestParseChannelsOneUnknownChannel(t *testing.T) {
	var input = "unknown"
	var expected = Channel(0)

	result, err := ParseChannels(input)

	if err == nil {
		t.Error("Error should be returned")
	}

	if expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestParseChannelsTwoChannelsWhereOneIsUnknown(t *testing.T) {
	var input = "sms,unknown"
	var expected = Channel(0)

	result, err := ParseChannels(input)

	if err == nil {
		t.Error("Error should be returned")
	}

	if expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestParseChannelsEmptyString(t *testing.T) {
	var input = ""
	var expected = Channel(0)

	result, err := ParseChannels(input)

	if err != nil {
		t.Errorf("Error should not be returned: %q", err)
	}

	if expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestParseChannelsNotTrimmedEmptyString(t *testing.T) {
	var input = "        "
	var expected = Channel(0)

	result, err := ParseChannels(input)

	if err != nil {
		t.Errorf("Error should not be returned: %q", err)
	}

	if expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestParseChannelsDuplicateChannels(t *testing.T) {
	var input = "sms,sms,sms"
	var expected = ChannelSms

	result, err := ParseChannels(input)

	if err != nil {
		t.Errorf("Error should not be returned: %q", err)
	}

	if expected != result {
		t.Errorf("Expected: %v, actual: %v", expected, result)
	}
}

func TestChannelHasEmailPushSms(t *testing.T) {
	channel, _ := ParseChannels("email,push,sms")
	emailChannel := ChannelEmail
	pushChannel := ChannelPush
	smsChannel := ChannelSms

	if !channel.Has(emailChannel) {
		t.Errorf("Expected: true, actual: false")
	}

	if !channel.Has(pushChannel) {
		t.Errorf("Expected: true, actual: false")
	}

	if !channel.Has(smsChannel) {
		t.Errorf("Expected: true, actual: false")
	}
}

func TestChannelHasEmailPush(t *testing.T) {
	channel, _ := ParseChannels("email,push")
	emailChannel := ChannelEmail
	pushChannel := ChannelPush
	smsChannel := ChannelSms

	if !channel.Has(emailChannel) {
		t.Errorf("Expected: true, actual: false")
	}

	if !channel.Has(pushChannel) {
		t.Errorf("Expected: true, actual: false")
	}

	if channel.Has(smsChannel) {
		t.Errorf("Expected: false, actual: true")
	}
}

func TestChannelHasEmailSms(t *testing.T) {
	channel, _ := ParseChannels("email,sms")
	emailChannel := ChannelEmail
	pushChannel := ChannelPush
	smsChannel := ChannelSms

	if !channel.Has(emailChannel) {
		t.Errorf("Expected: true, actual: false")
	}

	if channel.Has(pushChannel) {
		t.Errorf("Expected: false, actual: true")
	}

	if !channel.Has(smsChannel) {
		t.Errorf("Expected: true, actual: false")
	}
}
