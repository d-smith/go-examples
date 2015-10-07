package jwt

import "testing"

func BenchmarkCreateAndSignJWT(b *testing.B) {

	key := PrivateKeyFromPEM()

	for i := 0; i < b.N; i++ {
		CreateAndSignToken(key)
	}
}

func BenchmarkExtractKeyAndCreateSignedToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		key := PrivateKeyFromPEM()
		CreateAndSignToken(key)
	}
}
