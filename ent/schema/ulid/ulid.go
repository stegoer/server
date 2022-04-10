// inspired by https://github.com/manakuro/golang-clean-architecture-ent-gqlgen/blob/main/ent/schema/ulid/ulid.go

package ulid

import (
	"crypto/rand"
	"database/sql/driver"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/oklog/ulid/v2"
)

// ID implements a Universally unique Lexicographically sortable Identifier.
type ID string

var defaultEntropySource = ulid.Monotonic(rand.Reader, 0) //nolint:lll,gochecknoglobals

// newULID returns a new ID for time.Now() using the default entropy source.
func newULID() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(time.Now()), defaultEntropySource)
}

// MustNew returns a new ID for time.Now() given a prefix.
//
// It uses the default entropy source.
func MustNew(prefix string) ID {
	return ID(fmt.Sprintf(`%s%s`, prefix, newULID()))
}

// UnmarshalGQL implements the graphql.Unmarshaler interface.
func (i *ID) UnmarshalGQL(v interface{}) error {
	return i.Scan(v)
}

// MarshalGQL implements the graphql.Marshaler interface.
func (i ID) MarshalGQL(w io.Writer) {
	_, _ = io.WriteString(w, strconv.Quote(string(i)))
}

// Scan implements the Scanner interface.
func (i *ID) Scan(src interface{}) error {
	if src == nil {
		return errors.New("expected a value, got nil")
	}

	switch srcType := src.(type) {
	case string:
		*i = ID(srcType)
	case []byte:
		str := string(srcType)
		*i = ID(str)
	default:
		return fmt.Errorf(`expected a string or []byte, got: %v`, srcType)
	}

	return nil
}

// Value implements the driver.Value interface.
func (i ID) Value() (driver.Value, error) {
	return string(i), nil
}
