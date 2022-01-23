package globalid

import (
	"fmt"
	"log"
	"reflect"

	"github.com/kucera-lukas/stegoer/ent/image"
	"github.com/kucera-lukas/stegoer/ent/user"
)

type field struct {
	Prefix string
	Table  string
}

// GlobalIDs maps unique string to tables names.
type GlobalIDs struct {
	User  field
	Image field
}

// New generates a map object.
//
// It is intended to be used as global identification for node interface query.
// Prefix should have a max length of 3 characters to encode the entity type.
func New() GlobalIDs {
	return GlobalIDs{
		User: field{
			Prefix: "USR",
			Table:  user.Table,
		},
		Image: field{
			Prefix: "IMG",
			Table:  image.Table,
		},
	}
}

var (
	globalIDS = New()                   //nolint:gochecknoglobals
	maps      = structToMap(&globalIDS) //nolint:gochecknoglobals
)

// FindTableByID returns table name by passed id.
func (GlobalIDs) FindTableByID(id string) (string, error) {
	table, ok := maps[id]
	if !ok {
		return "", fmt.Errorf("could not map '%s' to a table name", id)
	}

	return table, nil
}

func structToMap(data *GlobalIDs) map[string]string {
	elem := reflect.ValueOf(data).Elem()
	size := elem.NumField()
	result := make(map[string]string, size)

	for i := 0; i < size; i++ {
		value := elem.Field(i).Interface()
		valueField, ok := value.(field)

		if !ok {
			log.Panicf("could not convert struct to map")
		}

		result[valueField.Prefix] = valueField.Table
	}

	return result
}
