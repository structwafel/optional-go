// Deserialize JSON to Option[T] and serialize Option[T] to JSON

package optional

import (
	"encoding/json"
)

func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.isSome {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

func (o *Option[T]) UnmarshalJSON(data []byte) error {
	// json value of null will be seen as None
	if string(data) == "null" {
		o.isSome = false
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	o.value = value
	o.isSome = true
	return nil
}

// NullNotNone is a wrapper for Option[T] that allows null to be a valid value
// type NullNotNone[T any] struct {
// 	inner Option[T]
// }

// func (n NullNotNone[T]) ToOption() Option[T] {
// 	return n.inner
// }

// func (o *NullNotNone[T]) MarshalJSON() ([]byte, error) {
// 	if o.inner.isSome {
// 		return json.Marshal(o.inner.value)
// 	}
// 	return json.Marshal(nil)
// }

// func (o *NullNotNone[T]) UnmarshalJSON(data []byte) error {
// 	// json value of null will be seen as None
// 	if string(data) == "null" {
// 		o.inner.isSome = true
// 		var zero T
// 		o.inner.value = zero
// 		return nil
// 	}

// 	var value T
// 	if err := json.Unmarshal(data, &value); err != nil {
// 		return err
// 	}
// 	o.inner.value = value
// 	o.inner.isSome = true
// 	return nil
// }
