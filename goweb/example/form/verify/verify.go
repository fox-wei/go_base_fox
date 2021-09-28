package verify

import "regexp"

/*
*verify：Chinese，English，telephone，email，id
 */

//^The Chinese
func IsChinese(value string) bool {
	if m, _ := regexp.MatchString("^\\p{Han}+$", value); !m {
		return false
	}
	return true
}

//^The English
func IsEnglish(value string) bool {
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", value); !m {
		return false
	}
	return true
}

//^The Email
func IsEmail(value string) bool {
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, value); !m {
		return false
	}
	return true
}

//^The id
func IsID(value string) bool {
	if len(value) <= 15 {
		if m, _ := regexp.MatchString(`^(\d{15})$`, value); !m {
			return false
		}
	} else {
		//验证 18 位身份证，18 位前 17 位为数字，最后一位是校验位，可能为数字或字符 X。
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, value); !m {
			return false
		}
	}
	return true
}

func IsMobile(value string) bool {
	//*Chinese number
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, value); !m {
		return false
	}
	return true
}
