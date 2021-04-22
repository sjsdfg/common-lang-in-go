package Copier

import (
	"fmt"

	"github.com/jinzhu/copier"
	jsoniter "github.com/json-iterator/go"
)

// DeepCopy Make a deep copy from src into dst.
func DeepCopy(dst interface{}, src interface{}) error {
	return copier.CopyWithOption(dst, src, copier.Option{
		DeepCopy:    true,
		IgnoreEmpty: true,
	})
}

// JsonCopy using jsoniter to copy value
func JsonCopy(dst interface{}, src interface{}) error {
	if dst == nil {
		return fmt.Errorf("dst cannot be nil")
	}
	if src == nil {
		return fmt.Errorf("src cannot be nil")
	}
	bytes, err := jsoniter.Marshal(src)
	if err != nil {
		return fmt.Errorf("unable to marshal src: %v", err)
	}
	err = jsoniter.Unmarshal(bytes, dst)
	if err != nil {
		return fmt.Errorf("unable to unmarshal into dst: %v", err)
	}
	return nil
}
