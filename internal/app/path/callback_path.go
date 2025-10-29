package path

import (
	"errors"
	"fmt"
	"strings"
)

type CallbackPath struct {
	Domain       string
	Subdomain    string
	CallbackName string
	CallbackData string
}

var ErrUknownCallback = errors.New("unknown callback")

func ParseCallback(callbackData string) (CallbackPath, error) {
	callbackParts := strings.SplitN(callbackData, "__", 4)
	if len(callbackParts) != 4 {
		return CallbackPath{}, ErrUknownCallback
	}

	return CallbackPath{
		Domain:       callbackParts[0],
		Subdomain:    callbackParts[1],
		CallbackName: callbackParts[2],
		CallbackData: callbackParts[3],
	}, nil
}

func (c CallbackPath) String() string {
	return fmt.Sprintf("%s__%s__%s__%s", c.Domain, c.Subdomain, c.CallbackName, c.CallbackData)
}
