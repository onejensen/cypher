package cipher_test

import (
	cipher "cypher/pkg"
	"testing"
)

// 测试ROT13加密和解密
func TestROT13(t *testing.T) {
	original := "Hello, World!"
	encrypted := cipher.EncryptROT13(original)
	decrypted := cipher.DecryptROT13(encrypted)

	if decrypted != original {
		t.Errorf("DecryptROT13(EncryptROT13(%q)) = %q, want %q", original, decrypted, original)
	}
}

// 测试Reverse加密和解密
func TestReverse(t *testing.T) {
	original := "Hello, World!"
	encrypted := cipher.EncryptReverse(original)
	decrypted := cipher.DecryptReverse(encrypted)

	if decrypted != original {
		t.Errorf("DecryptReverse(EncryptReverse(%q)) = %q, want %q", original, decrypted, original)
	}
}
