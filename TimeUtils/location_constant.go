package TimeUtils

import "time"

// 命名由来 https://zh.wikipedia.org/wiki/%E6%97%B6%E5%8C%BA%E5%88%97%E8%A1%A8
var (
	// 标准时区
	UTC = time.FixedZone("UTC", 0)
	// 东部时区
	CET = time.FixedZone("UTC+1", int(SecondsPerHour)*1)
	EET = time.FixedZone("UTC+2", int(SecondsPerHour)*2)
	MSK = time.FixedZone("UTC+3", int(SecondsPerHour)*3)
	GST = time.FixedZone("UTC+4", int(SecondsPerHour)*4)
	PKT = time.FixedZone("UTC+5", int(SecondsPerHour)*5)
	BHT = time.FixedZone("UTC+6", int(SecondsPerHour)*6)
	ICT = time.FixedZone("UTC+7", int(SecondsPerHour)*7)
	// 中国标准时间
	CT   = time.FixedZone("UTC+8", int(SecondsPerHour)*8)
	JST  = time.FixedZone("UTC+9", int(SecondsPerHour)*9)
	AEST = time.FixedZone("UTC+10", int(SecondsPerHour)*10)
	VUT  = time.FixedZone("UTC+11", int(SecondsPerHour)*11)
	// 东西十二区
	IDLW = time.FixedZone("UTC+12", int(SecondsPerHour)*12)
	// 西部时区
	SST  = time.FixedZone("UTC-11", int(SecondsPerHour)*-11)
	HST  = time.FixedZone("UTC-10", int(SecondsPerHour)*-10)
	AKST = time.FixedZone("UTC-9", int(SecondsPerHour)*-9)
	PST  = time.FixedZone("UTC-8", int(SecondsPerHour)*-8)
	MST  = time.FixedZone("UTC-7", int(SecondsPerHour)*-7)
	CST  = time.FixedZone("UTC-6", int(SecondsPerHour)*-6)
	EST  = time.FixedZone("UTC-5", int(SecondsPerHour)*-5)
	AST  = time.FixedZone("UTC-4", int(SecondsPerHour)*-4)
	BRT  = time.FixedZone("UTC-3", int(SecondsPerHour)*-3)
	FNT  = time.FixedZone("UTC-2", int(SecondsPerHour)*-2)
	CVT  = time.FixedZone("UTC-1", int(SecondsPerHour)*-1)
)
