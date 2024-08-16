package encryptions

/*
AES（Advanced Encryption Standard）高级加密标准。Rijndael算法是AES标准的一个实现，一般说AES指的就是Rijndael算法。
	AES密钥的长度只有16，24，32字节，分别对应AES-128，AES-192、AES-256
	AES支持5种加密模式：
	1、ECB(Electronic Codebook)模式：最简单的模式，将明文分成固定大小的块，然后分别加密。这种模式的问题是相同的明文块会加密成相同的密文块，因此缺乏随机性和安全性。
	2、CBC(Cipher Block Chaining)模式：每个明文块与前一个密文块进行异或操作，然后分别加密。这增加了随机性，并且在错误传播方面具有良好的性质，但需要一个初始化向阳(IV)
	3、CFB(Cipher Feedback)模式：通过将前一个密文块作为输入与明文块进行异或操作，产生密文块。这种模式对于错误传播具有良好的性质，但也需要一个初始化向量。
	4、CTR(Counter)模式：使用计数器来生成一个密钥流，然后与明文进行异或操作。在并行计算环境中效率较高。
*/

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// PKCS7Padding 将明文填充为块长度的整数倍
func pkcS7Padding(pass []byte, blockSize int) []byte {
	pad := blockSize - len(pass)%blockSize
	padtext := bytes.Repeat([]byte{byte(pad)}, pad)

	return append(pass, padtext...)
}

// PKCS7UnPadding 从明文尾部删除填充数据
func pkcS7UnPadding(p []byte) []byte {
	length := len(p)
	paddLen := int(p[length-1])

	return p[:(length - paddLen)]
}

// aesCBCEncrypt CBC模式下用AES算法加密数据
// 注意：要选择AES-128、AES-192或AES-256，密钥长度必须为16、24或32个字节
// 注意：AES块大小为16字节
func aesCBCEncrypt(p, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	p = pkcS7Padding(p, block.BlockSize())
	ciphertext := make([]byte, len(p))
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(ciphertext, p)

	return ciphertext, nil
}

// aesCBCDecrypt CBC模式下使用AES算法解密密文
func aesCBCDecrypt(c, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(c))
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	blockMode.CryptBlocks(plaintext, c)

	return pkcS7UnPadding(plaintext), nil
}

// Base64AESCBCEncrypt CBC模式下用AES算法加密数据，用base64编码
// key: 密钥长度必须为16、24或32个字节
func Base64AESCBCEncrypt(pass, key []byte) (string, error) {
	c, err := aesCBCEncrypt(pass, key)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(c), nil
}

// Base64AESCBCDecrypt CBC模式下用AES算法解密base64编码的密文
func Base64AESCBCDecrypt(c string, key []byte) ([]byte, error) {
	oriCipher, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		return nil, err
	}
	p, err := aesCBCDecrypt(oriCipher, key)
	if err != nil {
		return nil, err
	}

	return p, nil
}
