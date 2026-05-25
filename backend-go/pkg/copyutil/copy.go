package copy

import "encoding/json"

func JSON[T any](src any, dst *T) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, dst); err != nil {
		return err
	}

	return nil
}
