package templates

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
)

type SalesTaxField []SalesTax
type CustomerField map[string]interface{}
type JsonField map[string]interface{}

type amount sql.NullFloat64

func (s *SalesTaxField) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (s *SalesTaxField) Scan(val interface{}) error {

	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &s)
		return nil
	case string:
		json.Unmarshal([]byte(v), &s)
		return nil
	default:
		json.Unmarshal([]byte(nil), &s)
		return nil
		// return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}



func (js *JsonField) Value() (driver.Value, error) {
	return json.Marshal(js)
}

func (js *JsonField) Scan(val interface{}) error {

	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &js)
		return nil
	case string:
		json.Unmarshal([]byte(v), &js)
		return nil
	default:
		json.Unmarshal([]byte(nil), &js)
		return nil
		// return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (nf *amount) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte("null"), nil
	}

	return json.Marshal(fmt.Sprintf("%.2f", nf.Float64))
}

// Scan implements the Scanner interface for amount
func (nf *amount) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return err
	}

	// if nil then make Valid false
	if reflect.TypeOf(value) == nil {
		*nf = amount{f.Float64, false}

	} else {
		*nf = amount{f.Float64, true}
	}

	return nil
}
