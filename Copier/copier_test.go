package Copier

import (
	"log"
	"testing"
	"time"

	"github.com/sjsdfg/common-lang-in-go/JsonUtils"
	"github.com/stretchr/testify/assert"
)

func TestDeepCopy(t *testing.T) {
	source := &UserInfo{
		UserName:    "1",
		UserImage:   "2",
		RoomId:      "3",
		LiveId:      "4",
		LiveState:   "5",
		CreatedTime: time.Now(), // TODO copier 框架对于 time.Time 复制操作失效
	}

	dst := new(UserInfo)
	t.Log(source, dst)
	_ = DeepCopy(dst, source)
	t.Log(dst)
	assert.Equal(t, source, dst)
}

type A struct {
	A string `copier:"B"`
}

type B struct {
	B string
}

func TestDeepCopy2(t *testing.T) {
	a := &A{A: "a"}
	b := new(B)
	err := DeepCopy(b, a)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(a)
	t.Log(b)
}

func TestJsonCopy(t *testing.T) {
	source := &UserInfo{
		UserName:  "1",
		UserImage: "2",
		RoomId:    "3",
		LiveId:    "4",
		LiveState: "5",
		// CreatedTime: time.Now(),
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
	UserName    string    `json:"userName,omitempty"`
	UserImage   string    `json:"userImage,omitempty"`
	RoomId      string    `json:"roomId,omitempty"`
	LiveId      string    `json:"liveId,omitempty"`
	LiveState   string    `json:"liveState,omitempty"`
	Time        int64     `json:"time,string,omitempty"`
	CreatedTime time.Time `json:"createdTime"`
}

type NestedUserInfo struct {
	UserInfo UserInfo         `json:"userInfo,string,omitempty"`
	Map      map[string]int64 `json:"map,string,omitempty"`
}

func TestJsonCopy3(t *testing.T) {
	info := &NestedUserInfo{
		UserInfo: UserInfo{
			UserName:  "123",
			UserImage: "456",
			RoomId:    "879",
			LiveId:    "adsf",
			LiveState: "aga",
			Time:      10,
		},
		Map: map[string]int64{"1": 3},
	}
	jsonStr, err := JsonUtils.MarshalToString(info)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(jsonStr)

	info2 := &NestedUserInfo{}
	err = JsonUtils.Unmarshal([]byte(jsonStr), info2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(info2)
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
