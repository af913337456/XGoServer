package encrypt

/**

作者(Author): 林冠宏 / 指尖下的幽灵

Created on : 2018/2/19

*/

const DefaultAESKey = "12345678910Asdfxzvreytggfss78954"

type IEncrypt interface {
	AesEncrypt(origData, key []byte) ([]byte, error)   // 加密
	AesDecrypt(encrypted, key []byte) ([]byte, error)  // 解密
}
