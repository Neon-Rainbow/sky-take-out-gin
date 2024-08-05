package encrypt

import (
	"crypto/md5"
	"encoding/hex"
	"sky-take-out-gin/pkg/common/config"
)

// EncryptPassword 加密密码
// @Param password: 密码
// @Return string: 加密后的密码
func EncryptPassword(password string) string {
	var secret = config.GetConfig().SecretConfig.PasswordSecret
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}
