package models

const (
	ClaimKey      = "jwt_key"
	Authorization = "Authorization"
	Bearer        = "bearer"
	DefaultTTL    = 86400e9 * 30
)

var JwtSecret = []byte("xingshidaiyuhang_earthlive!@#$eis")
