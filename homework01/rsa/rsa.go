package rsa

import (
	"errors"
	"math"
	"math/big"
	"math/rand"
)

type Key struct {
	key int
	n   int
}

type KeyPair struct {
	Private Key
	Public  Key
}

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	k := int(math.Sqrt(float64(n)))
	for i := 2; i <= k; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func multiplicativeInverse(e int, phi int) int {
	s1, s2 := 1, 0
	t1, t2 := 0, 1
	for phi != 0 {
		q := e / phi
		r := e % phi
		e, phi = phi, r
		s := s1 - q*s2
		s1, s2 = s2, s
		t := t1 - q*t2
		t1, t2 = t2, t
	}
	result := s1 % s2
	if result < 0 {
		result += s2
	}
	return result
}

func GenerateKeypair(p int, q int) (KeyPair, error) {
	if !isPrime(p) || !isPrime(q) {
		return KeyPair{}, errors.New("Both numbers must be prime.")
	} else if p == q {
		return KeyPair{}, errors.New("p and q can't be equal.")
	}

	// n = pq
	n := p * q

	// phi = (p-1)(q-1)
	phi := (p - 1) * (q - 1)

	e := rand.Intn(phi-1) + 1
	g := gcd(e, phi)
	for g != 1 {
		e = rand.Intn(phi-1) + 1
		g = gcd(e, phi)
	}

	d := multiplicativeInverse(e, phi)
	return KeyPair{Key{e, n}, Key{d, n}}, nil
}

func Encrypt(pk Key, plaintext string) []int {
	cipher := []int{}
	n := new(big.Int)
	for _, ch := range plaintext {
		n = new(big.Int).Exp(
			big.NewInt(int64(ch)), big.NewInt(int64(pk.key)), nil)
		n = new(big.Int).Mod(n, big.NewInt(int64(pk.n)))
		cipher = append(cipher, int(n.Int64()))
	}
	return cipher
}

func Decrypt(pk Key, cipher []int) string {
	plaintext := ""
	n := new(big.Int)
	for _, ch := range cipher {
		n = new(big.Int).Exp(
			big.NewInt(int64(ch)), big.NewInt(int64(pk.key)), nil)
		n = new(big.Int).Mod(n, big.NewInt(int64(pk.n)))
		plaintext += string(rune(int(n.Int64())))
	}
	return plaintext
}
