package zj005_空格替换

func replaceSpace(s string) string {
	sbyes := []byte(s)
	repace := ""
	for _, v := range sbyes {
		if v == ' ' {
			repace = repace + "%20"
		} else {
			repace = repace + string(v)
		}
	}
	return repace
}
