package bcrypt

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMd5(input string, salt string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum([]byte(salt)))

}
