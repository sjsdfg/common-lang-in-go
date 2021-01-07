package TimeUtils

import "time"

func init() {
	LocationChina, _ = time.LoadLocation("Asia/Shanghai")
	LocationIndonesia, _ = time.LoadLocation("Asia/Jakarta")
}
