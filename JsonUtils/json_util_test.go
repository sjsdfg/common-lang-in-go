package JsonUtils

import (
	"testing"

	"github.com/sjsdfg/common-lang-in-go/Copier"
)

func TestToMap(t *testing.T) {
	toMap, err := ToMap(&UserInfo{
		UserName:  "1",
		UserImage: "2",
		RoomId:    "3",
		LiveId:    "4",
		LiveState: "5",
		OtherStruct: &OtherStruct{
			Name:  "6",
			Value: "7",
		},
	})
	t.Log(toMap)
	t.Log(err)
	var value UserInfo
	err = Copier.JsonCopy(&value, &toMap)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log(value)
	}
}

type UserInfo struct {
	UserName    string       `json:"userName"`
	UserImage   string       `json:"userImage"`
	RoomId      string       `json:"roomId"`
	LiveId      string       `json:"liveId"`
	LiveState   string       `json:"liveState"`
	OtherStruct *OtherStruct `json:"otherStruct,omitempty"`
}

type OtherStruct struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
