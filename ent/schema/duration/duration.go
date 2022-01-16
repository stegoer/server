package duration

import (
	"StegoLSB/pkg/entity/model"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

const formatIntBase = 10

type Duration time.Duration

// MarshalGQL implements the graphql.Marshaller interface.
func (d Duration) MarshalGQL(t time.Duration) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.FormatInt(int64(t), formatIntBase))
	})
}

// UnmarshalGQL implements the graphql.Unmarshaller interface.
func (d *Duration) UnmarshalGQL(value interface{}) error {
	return d.Scan(value)
}

// Scan implements the Scanner interface.
func (d *Duration) Scan(value interface{}) error {
	switch t := value.(type) {
	case int64:
		*d = Duration(t)
	case string:
		parsed, err := time.ParseDuration(t)
		if err != nil {
			return model.NewInvalidParamError(err, t)
		}
		*d = Duration(parsed)
	case json.Number:
		i, err := t.Int64()
		if err != nil {
			return model.NewInvalidParamError(err, t)
		}
		*d = Duration(i)
	default:
		return model.NewInvalidParamError(errors.Errorf("duration: invalid type %T to unmarshal", t), value)
	}

	return nil
}
