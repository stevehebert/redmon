package parser

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestParsingError(t *testing.T) {
	_, err := parseMessage("")

	assert.NotEqual(t, err, nil)

}

func TestSeparate(t *testing.T) {
	left, right, err := separate("[US-MN][H] Paypal [W] Rama m60a", "[H]")

	assert.Equal(t, err, nil)
	assert.Equal(t, left, "[US-MN]")
	assert.Equal(t, right, "Paypal [W] Rama m60a")
}

func TestParse(t *testing.T) {
	mm, err := Parse("[US-MN][H] Paypal [W] Rama m60a")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, mm, nil)
	assert.Equal(t, mm.location, "[US-MN]")
	assert.Equal(t, mm.have, "Paypal")
	assert.Equal(t, mm.want, "Rama m60a")

	mm, err = Parse("[US-MN][W] Paypal [H] Rama m60a")
	assert.Equal(t, err, nil)
	assert.NotEqual(t, mm, nil)
	assert.Equal(t, mm.location, "[US-MN]")
	assert.Equal(t, mm.want, "Paypal")
	assert.Equal(t, mm.have, "Rama m60a")

	mm, err = Parse("[W] Paypal [H] Rama m60a")
	assert.NotEqual(t, err, nil)

	mm, err = Parse("[US-MN][H] Rama m60a")
	assert.NotEqual(t, err, nil)

	mm, err = Parse("[US-MN][H] Rama m60a")
	assert.NotEqual(t, err, nil)
}
