package encrypt

import (
	"crypto/md5"
	"encoding/hex"
)

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/15 23:20
 * @file: md5.go
 * @description: md5工具
 */

// MD5Encrypt 加密
func MD5Encrypt(key string) string {
	d := []byte(key)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// MD5valueCompare 比较
func MD5valueCompare(originalCipher, newCipher string) bool {

	if MD5Encrypt(originalCipher) != newCipher {
		return false
	} else {
		return true
	}
}
