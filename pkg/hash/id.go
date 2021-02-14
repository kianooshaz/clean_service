package hash

import (
	"github.com/kianooshaz/clean_service/config"
	"github.com/speps/go-hashids"
)

func EncodeId(id uint) string {
	hd := hashids.NewData()
	hd.Salt = config.HashIdSalt()
	hd.MinLength = config.HashIdMinLength()
	h, _ := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{int(id)})
	return e
}

func DecodeId(id string) uint {
	hd := hashids.NewData()
	hd.Salt = config.HashIdSalt()
	hd.MinLength = config.HashIdMinLength()
	h, _ := hashids.NewWithData(hd)
	d, _ := h.DecodeWithError(id)
	return uint(d[0])
}
