package public

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
)

//MD5 md5加密
func MD5(s string) string {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}


// 获取加盐密码
func GenSaltPassword(passport string, salt string) (string) {
	s_ob1 := sha256.New()
	s_ob1.Write([]byte(passport))
	rs1 := fmt.Sprintf("%x", s_ob1.Sum(nil))
	s_ob2 := sha256.New()
	s_ob2.Write([]byte(rs1 + salt))
	rs2 := fmt.Sprintf("%x", s_ob2.Sum(nil))
	return rs2
}

//InStringList 数组中是否存在某值
func InStringList(t string, list []string) bool {
	for _, s := range list {
		if s == t {
			return true
		}
	}
	return false
}

//对象打印
func Obj2Json(a interface{}) string {
	bs, err := json.Marshal(a)
	if err != nil {
		return ""
	}
	return string(bs)
}