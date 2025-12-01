package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// MD5 生成MD5哈希值
func MD5(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}

// MD5Upper 生成MD5哈希值（大写）
func MD5Upper(data string) string {
	hash := md5.Sum([]byte(data))
	return fmt.Sprintf("%X", hash)
}

// AESConfig AES加密配置
type AESConfig struct {
	Key  string // 密钥，必须16、24或32字节
	IV   string // 初始化向量，必须16字节
	Mode string // 加密模式，支持ECB、CBC、CFB、OFB
}

// 默认AES配置
var defaultAESConfig = &AESConfig{
	Key:  "default_aes_key_123", // 默认密钥（16字节）
	IV:   "default_iv_123456",   // 默认初始化向量（16字节）
	Mode: "CBC",                 // 默认CBC模式
}

// AESEncrypt AES加密
func AESEncrypt(data string, config *AESConfig) (string, error) {
	if config == nil {
		config = defaultAESConfig
	}

	// 检查密钥长度
	key := []byte(config.Key)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", fmt.Errorf("aes key length must be 16, 24 or 32 bytes")
	}

	// 检查IV长度
	iv := []byte(config.IV)
	if len(iv) != 16 {
		return "", fmt.Errorf("aes iv length must be 16 bytes")
	}

	// 创建AES加密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 填充数据，使用PKCS#7填充
	padding := aes.BlockSize - len(data)%aes.BlockSize
	padtext := make([]byte, len(data)+padding)
	copy(padtext, data)
	for i := len(data); i < len(padtext); i++ {
		padtext[i] = byte(padding)
	}

	// 根据模式加密
	var ciphertext []byte
	switch config.Mode {
	case "CBC":
		ciphertext = make([]byte, len(padtext))
		mode := cipher.NewCBCEncrypter(block, iv)
		mode.CryptBlocks(ciphertext, []byte(padtext))
	case "ECB":
		ciphertext = make([]byte, len(padtext))
		mode := NewECBEncrypter(block)
		mode.CryptBlocks(ciphertext, []byte(padtext))
	default:
		return "", fmt.Errorf("unsupported aes mode: %s", config.Mode)
	}

	// 返回Base64编码
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// AESDecrypt AES解密
func AESDecrypt(ciphertext string, config *AESConfig) (string, error) {
	if config == nil {
		config = defaultAESConfig
	}

	// 检查密钥长度
	key := []byte(config.Key)
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", fmt.Errorf("aes key length must be 16, 24 or 32 bytes")
	}

	// 检查IV长度
	iv := []byte(config.IV)
	if len(iv) != 16 {
		return "", fmt.Errorf("aes iv length must be 16 bytes")
	}

	// 解码Base64
	cipherdata, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	// 创建AES解密块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 根据模式解密
	var plaintext []byte
	switch config.Mode {
	case "CBC":
		plaintext = make([]byte, len(cipherdata))
		mode := cipher.NewCBCDecrypter(block, iv)
		mode.CryptBlocks(plaintext, cipherdata)
	case "ECB":
		plaintext = make([]byte, len(cipherdata))
		mode := NewECBDecrypter(block)
		mode.CryptBlocks(plaintext, cipherdata)
	default:
		return "", fmt.Errorf("unsupported aes mode: %s", config.Mode)
	}

	// 去除PKCS#7填充
	padding := int(plaintext[len(plaintext)-1])
	if padding > aes.BlockSize || padding == 0 {
		return "", fmt.Errorf("invalid padding")
	}
	return string(plaintext[:len(plaintext)-padding]), nil
}

// ECB模式实现（Go标准库不直接支持ECB模式）

// ecbEncrypter ECB加密器
type ecbEncrypter struct {
	cipher.Block
}

// NewECBEncrypter 创建ECB加密器
func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
	return &ecbEncrypter{b}
}

// CryptBlocks ECB加密
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.BlockSize() != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.Encrypt(dst, src[:x.BlockSize()])
		dst = dst[x.BlockSize():]
		src = src[x.BlockSize():]
	}
}

// SetIV 设置IV（ECB模式不需要IV，此方法为空实现）
func (x *ecbEncrypter) SetIV([]byte) {
}

// ecbDecrypter ECB解密器
type ecbDecrypter struct {
	cipher.Block
}

// NewECBDecrypter 创建ECB解密器
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return &ecbDecrypter{b}
}

// CryptBlocks ECB解密
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
	if len(src)%x.BlockSize() != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.Decrypt(dst, src[:x.BlockSize()])
		dst = dst[x.BlockSize():]
		src = src[x.BlockSize():]
	}
}

// SetIV 设置IV（ECB模式不需要IV，此方法为空实现）
func (x *ecbDecrypter) SetIV([]byte) {
}
