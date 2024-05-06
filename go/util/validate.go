package util

import "regexp"

// 检查邮箱格式
func CheckEmail(email string) bool {
	if email == "" {
		return false
	}
	if matched, _ := regexp.MatchString("^([a-z0-9_\\.-]+)@([\\da-z\\.-]+)\\.([a-z\\.]{2,6})$", email); matched {
		return true
	}
	return false
}

// 检查密码格式 包含大小写字母、数字、@符号和句点（.）的字符串，长度在6到30个字符之间。
func CheckPassword(password string) bool {
	if password == "" {
		return false
	}
	if matched, _ := regexp.MatchString("^[0-9a-zA-Z@.]{6,30}$", password); matched {
		return true
	}
	return false
}

// 检查手机号格式
func CheckPhoneNumber(phoneNumber string) bool {
	if phoneNumber == "" {
		return false
	}
	if matched, _ := regexp.MatchString("^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])d{8}$", phoneNumber); matched {
		return true
	}
	return false
}

// 检查用户名格式 字母开头，允许5-16字节，允许字母数字下划线
func CheckUsername(username string) bool {
	if username == "" {
		return false
	}
	// 字母开头，允许5-16字节，允许字母数字下划线
	if matched, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_]{4,15}$", username); matched {
		return true
	}
	return false
}

// 检查内容是否为空
func CheckContent(str string) bool {
	return !AllIsInvisibleCharacter(str)
}

// 检查是否全是不可见字符 是返回true，否则返回false
func AllIsInvisibleCharacter(str string) bool {
	strList := []byte(str)
	for _, c := range strList {
		if c != ' ' && c != '\n' && c != '\t' {
			return false
		}
	}
	return true
}

// 检查是否为空字符串
func IsBlank(str string) bool {
	return DeletePreAndSufSpace(str) == ""
}

// 删除字符串前后空格
func DeletePreAndSufSpace(str string) string {
	strList := []byte(str)
	spaceCount, count := 0, len(strList)
	for i := 0; i <= count-1; i++ {
		if strList[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}

	strList = strList[spaceCount:]
	spaceCount, count = 0, len(strList)
	for i := count - 1; i >= 0; i-- {
		if strList[i] == 32 {
			spaceCount++
		} else {
			break
		}
	}

	return string(strList[:count-spaceCount])
}
