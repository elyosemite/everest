package domain

import (
	"errors"
	"fmt"
	"strings"
)

type Address struct {
	street     string
	number     string
	complement string
	city       string
	state      string
	zipCode    string
}

func NewAddress(street, number, city, state, zipCode string) (*Address, error) {
	if len(street) < 3 {
		return nil, errors.New("street must be at least 3 characters")
	}

	if len(state) != 2 {
		return nil, errors.New("state must be 2 characters (UF)")
	}

	return &Address{
		street:  strings.TrimSpace(street),
		number:  strings.TrimSpace(number),
		city:    strings.TrimSpace(city),
		state:   strings.ToUpper(strings.TrimSpace(state)),
		zipCode: strings.ReplaceAll(zipCode, "[^0-9]", ""),
	}, nil
}

func (a *Address) FullAddress() string {
	return fmt.Sprintf("%s, %s - %s/%s, CEP: %s",
		a.street,
		a.number,
		a.city,
		a.state,
		a.zipCode[:5]+"-"+a.zipCode[5:],
	)
}
