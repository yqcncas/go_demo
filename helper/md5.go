package helper

import (
	"crypto/md5"
	"fmt"
)

func GetMD5(pwd string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(pwd)))
}
