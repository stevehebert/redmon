package parser

import (
	"errors"
	"strings"
)

type MarketMessage struct {
	Location string
	Have     string
	Want     string
}

func parseMessage(rawMessage string) (*MarketMessage, error) {

	_, _, err := separate(rawMessage, "[H]")

	return nil, err
}

func Parse(message string) (*MarketMessage, error) {
	left, have, error := separate(message, "[H]")

	if error != nil {
		return nil, errors.New("unable to locate [HAVE] portion of message")
	}

	// we know have is in the right hand side.  therefore, if want is also in the right hand
	// side, then we know location is on the left

	have1, want, err := separate(have, "[W]")
	if err != nil {
		// here we know that WANT must be in the left hand with location
		location, want, err := separate(left, "[W]")
		if err != nil {
			// couldn't find want, therefore the structure is wrong
			return nil, errors.New("unable to locate [W] portion of message")
		}

		// check to see that location has a value

		if location == "" {
			return nil, errors.New("location is undefined")
		}

		return &MarketMessage{Location: location, Have: have, Want: want}, nil
	}

	if left == "" {
		return nil, errors.New("location is undefined")
	}

	return &MarketMessage{Location: left, Have: have1, Want: want}, nil
}

func separate(search string, separator string) (string, string, error) {
	results := strings.SplitN(search, separator, 2)

	if len(results) == 2 {
		return strings.Trim(results[0], " "), strings.Trim(results[1], " "), nil
	}

	return "", "", errors.New("missing seperator in string")

}
