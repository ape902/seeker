package encryptions

import (
	"errors"
	"testing"
)

func TestNewAES(t *testing.T) {
	type args struct {
		Password   string
		EncryptKey string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			"NewAES",
			args{
				Password: "123456",
				// 密钥长度必须为16、24或32个字节
				EncryptKey: "12345678abcdefgh",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ciphertext, err := Base64AESCBCEncrypt([]byte(tt.args.Password), []byte(tt.args.EncryptKey))
			if err != nil {
				t.Error(err)
				return
			}

			plaintext, err := Base64AESCBCDecrypt(ciphertext, []byte(tt.args.EncryptKey))
			if err != nil {
				t.Error(err)
				return
			}

			if string(plaintext) != tt.args.Password {
				t.Error(errors.New("密码验证错误"))
			}
		})
	}
}
