package TimeUtils

import "time"

var (
	// 标准时区
	UTC = time.FixedZone("UTC", 0)
	// 东部时区
	UTC1  = time.FixedZone("UTC", int(SecondsPerHour)*1)
	UTC2  = time.FixedZone("UTC", int(SecondsPerHour)*2)
	UTC3  = time.FixedZone("UTC", int(SecondsPerHour)*3)
	UTC4  = time.FixedZone("UTC", int(SecondsPerHour)*4)
	UTC5  = time.FixedZone("UTC", int(SecondsPerHour)*5)
	UTC6  = time.FixedZone("UTC", int(SecondsPerHour)*6)
	UTC7  = time.FixedZone("UTC", int(SecondsPerHour)*7)
	UTC8  = time.FixedZone("UTC", int(SecondsPerHour)*8)
	UTC9  = time.FixedZone("UTC", int(SecondsPerHour)*9)
	UTC10 = time.FixedZone("UTC", int(SecondsPerHour)*10)
	UTC11 = time.FixedZone("UTC", int(SecondsPerHour)*11)
	// 东西十二区
	UTC12 = time.FixedZone("UTC", int(SecondsPerHour)*12)
	// 西部时区
	UTCNeg11 = time.FixedZone("UTC", int(SecondsPerHour)*-11)
	UTCNeg10 = time.FixedZone("UTC", int(SecondsPerHour)*-10)
	UTCNeg9  = time.FixedZone("UTC", int(SecondsPerHour)*-9)
	UTCNeg8  = time.FixedZone("UTC", int(SecondsPerHour)*-8)
	UTCNeg7  = time.FixedZone("UTC", int(SecondsPerHour)*-7)
	UTCNeg6  = time.FixedZone("UTC", int(SecondsPerHour)*-6)
	UTCNeg5  = time.FixedZone("UTC", int(SecondsPerHour)*-5)
	UTCNeg4  = time.FixedZone("UTC", int(SecondsPerHour)*-4)
	UTCNeg3  = time.FixedZone("UTC", int(SecondsPerHour)*-3)
	UTCNeg2  = time.FixedZone("UTC", int(SecondsPerHour)*-2)
	UTCNeg1  = time.FixedZone("UTC", int(SecondsPerHour)*-1)
)
