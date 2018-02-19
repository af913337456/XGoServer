package encrypt

import (
	"testing"
	"fmt"
)

/**

作者(Author): 林冠宏 / 指尖下的幽灵

Created on : 2018/2/19

*/

func TestEncAes(t *testing.T) {

	testStr :=
		"在CBC模式中，每个平文块先与前一个密文块进行异或后，再进行加密。在这种方法中，每个密文块都依赖于它前面的所有平文块。" +
			"同时，为了保证每条消息的唯一性，在第一个块中需要使用初始化向量。"+
	"CBC是最为常用的工作模式。它的主要缺点在于加密过程是串行的，无法被并行化，而且消息必须被填充到块大小的整数倍。" +
		"解决后一个问题的一种方法是利用密文窃取。"+
	"注意在加密时，平文中的微小改变会导致其后的全部密文块发生改变，而在解密时，从两个邻接的密文块中即可得到一个平文块。" +
		"因此，解密过程可以被并行化，而解密时，密文中一位的改变只会导致其对应的平文块完全改变和下一个平文块中对应位发生改变，" +
			"不会影响到其它平文的内容。"
	testStr = testStr + testStr + testStr + testStr + testStr + testStr

	e := DefaultAES{}
	data := e.AesEncryptStr(testStr)

	fmt.Println(data)
	decAes(data)
}

func Test2(t *testing.T) {
	e := DefaultAES{}

	data := e.AesEncryptStr("狗年平安")

	fmt.Println(data)
//\ufffd\u0002\ufffd\u001eu\ufffd\u001c\ufffd\u0013m\ufffdM\\\ufffd\ufffd\u0011
	decAes(data)
}

func decAes(encData string) {
	e := DefaultAES{}
	data := e.AesDecryptStr(encData)

	fmt.Println(data)
}


