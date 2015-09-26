package jwt
import "testing"

func BenchmarkCreateAndSignJWT(b *testing.B) {

	key := PrivateKeyFromPEM()

	for i := 0; i < b.N; i++ {
		CreateAndSignToken(key)
	}
}
