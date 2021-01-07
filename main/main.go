package main

import (
	"fmt"
	"github.com/sjsdfg/common-lang-in-go/TimeUtils"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339) == now.In(TimeUtils.UTC8).Format(time.RFC3339))
}
