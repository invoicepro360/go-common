package templates

import (
	"database/sql/driver"
	"encoding/json"
)

type SalesTaxField []SalesTax
type CustomerField map[string]interface{}
type InvoiceColumnField map[string]interface{}

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

func (c *CustomerField) Value() (driver.Value, error) {
	return json.Marshal(c)
}

func (c *CustomerField) Scan(val interface{}) error {

	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &c)
		return nil
	case string:
		json.Unmarshal([]byte(v), &c)
		return nil
	default:
		json.Unmarshal([]byte(nil), &c)
		return nil
		// return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}

func (ic *InvoiceColumnField) Value() (driver.Value, error) {
	return json.Marshal(ic)
}

func (ic *InvoiceColumnField) Scan(val interface{}) error {
	switch v := val.(type) {
	case []byte:
		json.Unmarshal(v, &ic)
		return nil
	case string:
		json.Unmarshal([]byte(v), &ic)
		return nil
	default:
		json.Unmarshal([]byte(nil), &ic)
		return nil
		// return errors.New(fmt.Sprintf("Unsupported type: %T", v))
	}
}
