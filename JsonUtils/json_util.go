package JsonUtils

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/tidwall/pretty"
)

var instance = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(obj interface{}) ([]byte, error) {
	return instance.Marshal(obj)
}

func MarshalToString(obj interface{}) (string, error) {
	return instance.MarshalToString(obj)
}

func Unmarshal(data []byte, v interface{}) error {
	return instance.Unmarshal(data, v)
}

func PrettyJson(jsonBytes []byte) []byte {
	return pretty.Pretty(jsonBytes)
}

func Pretty(v interface{}) []byte {
	bytes, _ := Marshal(v)
	return PrettyJson(bytes)
}

func PrettyWithColor(v interface{}) []byte {
	return pretty.Color(Pretty(v), pretty.TerminalStyle)
}
