package bcrypt

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/joho/godotenv"
	"github.com/kianooshaz/clean_service/core/pkg/logs"
	"os"
)

func GetMd5(input string) string {
	if err := godotenv.Load(); err != nil {
		logs.ErrorLogger.Fatalln("Error loading .env file")
	}
	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum([]byte(os.Getenv("CRYPTO_SECRET"))))

}
