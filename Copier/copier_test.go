package Copier

import (
	"log"
	"testing"
	"time"

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

func TestJsonCopy2(t *testing.T) {
	source := &UserInfo{
		UserName:  "123",
		UserImage: "456",
		RoomId:    "",
		LiveId:    "",
		LiveState: "",
	}

	dst := &UserInfo{
		UserName:  "",
		UserImage: "789",
		RoomId:    "789",
		LiveId:    "456",
		LiveState: "231",
		Time:      time.Now().Unix(),
	}

	_ = JsonCopy(dst, source)
	t.Log(dst)
}

type UserInfo struct {
	UserName  string `json:"userName,omitempty"`
	UserImage string `json:"userImage,omitempty"`
	RoomId    string `json:"roomId,omitempty"`
	LiveId    string `json:"liveId,omitempty"`
	LiveState string `json:"liveState,omitempty"`
	Time      int64  `json:"time,omitempty"`
}

func TestRefer(t *testing.T) {
	log.Println("Point1:", runSomething())
}

func runSomething() int {
	x := 0
	defer func(y *int) {
		log.Println("Point2:", *y)
		*y = 7
	}(&x)
	x = 2
	return x
}
