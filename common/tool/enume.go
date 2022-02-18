package tool

var GenderInfo = map[int64]string{
	1: "男",
	2: "女",
	3: "保密",
}

func GetGenderInfo(gender int64) string {
	if info, ok := GenderInfo[gender]; ok {
		return info
	}
	return GenderInfo[3]
}
