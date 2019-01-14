package util

import "golang.org/x/crypto/bcrypt"

// 生成密码
func GeneratePassword(password string, cost int) (string, error) {
	if cost < 1 {
		cost = 1
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// 校验密码
func ValidatePassword(hashPwd, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}
