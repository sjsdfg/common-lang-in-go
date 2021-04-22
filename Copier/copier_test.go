package Copier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeepCopy(t *testing.T) {
	source := &UserInfo{
		UserName:  "1",
		UserImage: "2",
		RoomId:    "3",
		LiveId:    "4",
		LiveState: "5",
	}

	dst := new(UserInfo)
	t.Log(dst)
	_ = DeepCopy(dst, source)
	t.Log(dst)
	assert.Equal(t, source, dst)
}

func TestJsonCopy(t *testing.T) {
	source := &UserInfo{
		UserName:  "1",
		UserImage: "2",
		RoomId:    "3",
		LiveId:    "4",
		LiveState: "5",
	}

	dst := new(UserInfo)
	t.Log(dst)
	_ = JsonCopy(dst, source)
	t.Log(dst)
	assert.Equal(t, source, dst)
}

type UserInfo struct {
	UserName  string `json:"userName"`
	UserImage string `json:"userImage"`
	RoomId    string `json:"roomId"`
	LiveId    string `json:"liveId"`
	LiveState string `json:"liveState"`
}
