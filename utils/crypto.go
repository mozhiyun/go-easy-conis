package utils

import (
	"crypto/sha1"
	"fmt"
)

func Sha1Hex(s string) string {
	//产生一个散列值得方式是 sha1.New()，sha1.Write(bytes)，然后 sha1.Sum([]byte{})。这里我们从一个新的散列开始。
	h := sha1.New() // md5加密类似md5.New()
	//写入要处理的字节。如果是一个字符串，需要使用[]byte(s) 来强制转换成字节数组。
	h.Write([]byte(s))
	//这个用来得到最终的散列值的字符切片。Sum 的参数可以用来对现有的字符切片追加额外的字节切片：一般不需要要。
	bs := h.Sum(nil)
	//如果需要对另一个字符串加密，要么重新生成一个新的散列，要么一定要调用h.Reset()方法，不然生成的加密字符串会是拼接第一个字符串之后进行加密
	h.Reset() //重要！！！
	//SHA1 值经常以 16 进制输出，使用%x 来将散列结果格式化为 16 进制字符串。
	return fmt.Sprintf("%x", bs)
	//h.Write([byte]("string2"))
	//fmt.Printf("%x\n", fmt.Sprintf("%x", h.Sum(nil)))
}
