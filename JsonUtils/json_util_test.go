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

func TestName(t *testing.T) {
	m2 := map[string]string{
		"id":                "123",
		"successTeamMember": `"{"aaa":123}"`,
	}
	json2 := `{"id":"123","successTeamMember":"\"{\"aaaa\":100,\"bbbb\":200}\""}`
	json, err := MarshalToString(m2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(json)
	t.Log(json2)
	// assert.EqualValues(t, json2, json)
	m := make(map[string]string)
	err = Unmarshal(([]byte)(json2), &m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(m)

	// s := struct {
	// 	Id                string           `json:"id"`
	// 	SuccessTeamMember map[string]int64 `json:"successTeamMember"`
	// }{}
	// err = Unmarshal(([]byte)(json), &s)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// t.Log(s)
}
