package helpers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

// NullString is a wrapper around sql.NullString
type NullString sql.NullString

// MarshalJSON method is called by json.Marshal,
// whenever it is of type NullString
func (x *NullString) MarshalJSON() ([]byte, error) {
	if !x.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(x.String)
}

// Pretty JSON
func PrettyJSON(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(string(b))
}
