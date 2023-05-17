package detection

func CheckSensitiveWord(str string) bool {
	for _, v := range str {
		if v == ' ' || v == ';' || v == '\n' || v == '\r' || v == '\t' || v == '&' || v == '|' || v == '!' {
			return false
		}
	}
	return true
}
