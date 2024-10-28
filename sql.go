package optional

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// Database Scan
func (o *Option[T]) Scan(value any) error {
	if value == nil {
		o.isSome = false
		return nil
	}
	var v T
	switch val := value.(type) {
	case T:
		v = val
	case []byte:
		if err := json.Unmarshal(val, &v); err != nil {
			return err
		}
	default:
		return errors.New("unsupported type")
	}
	o.value = v
	o.isSome = true
	return nil
}

// Database Value
func (o Option[T]) Value() (driver.Value, error) {
	if o.IsSome() {
		return json.Marshal(o.value)
	}
	return nil, nil
}
