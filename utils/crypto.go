/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/7/23
   Description :
-------------------------------------------------
*/

package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

var Crypto = new(cryptoUtil)

type cryptoUtil struct{}

func (*cryptoUtil) Md5(text string) string {
	m := md5.New()
	m.Write([]byte(text))
	return hex.EncodeToString(m.Sum(nil))
}

func (*cryptoUtil) Md532(text string) string {
	m := md5.New()
	m.Write([]byte(text))
	return hex.EncodeToString(m.Sum(nil))
}

func (*cryptoUtil) Md516(text string) string {
	m := md5.New()
	m.Write([]byte(text))
	return hex.EncodeToString(m.Sum(nil))[8:24]
}

func (*cryptoUtil) Sha1(text string) string {
	c := sha1.New()
	c.Write([]byte(text))
	return hex.EncodeToString(c.Sum(nil))
}
