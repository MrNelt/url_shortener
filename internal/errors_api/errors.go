package errorsApi

import (
	"fmt"
)

var (
	ErrSuffixAlreadyExists = fmt.Errorf("Suffix already exists")
	ErrSuffixNotFound      = fmt.Errorf("Suffix not found")
	ErrOwnSuffixWithoutTTL = fmt.Errorf("Own Suffix without TTL")
	ErrLinkNotFound        = fmt.Errorf("Link not found")
	ErrIDNotFound          = fmt.Errorf("ID not found")
)
