package fnv

import (
	"hash"
	"math/big"
)

type sum128a struct {
	*big.Int
}

type Hash128 interface {
	hash.Hash
	Sum128() big.Int
}

const (
	offset = "144066263297769815596495629667062367629"
	prime  = "309485009821345068724781371"
)

var offset128, prime128 sum128a

func init() {
	prime128 = sum128a{&big.Int{}}
	offset128 = sum128a{&big.Int{}}

	prime128.SetString("309485009821345068724781371", 0)
	offset128.SetString("144066263297769815596495629667062367629", 0)
}

// New128a returns a new 128-bit FNV-1a big.Int.
func New128a() Hash128 {
	var s sum128a = sum128a{&big.Int{}}
	s.Set(offset128.Int)
	return &s
}

func (s *sum128a) Reset()          { s.Set(offset128.Int) }
func (s *sum128a) Sum128() big.Int { return *s.Int }
func (s *sum128a) Size() int       { return 16 }
func (s *sum128a) BlockSize() int  { return 1 }

func (s *sum128a) Sum(in []byte) []byte {

	sBytes := s.Bytes()
	return append(in, sBytes[len(sBytes)-s.Size():]...)

}

func (s *sum128a) Write(data []byte) (int, error) {

	hash := sum128a{&big.Int{}}
	hash.Set(s.Int)
	for _, c := range data {
		hash.Xor(hash.Int, big.NewInt(int64(c)))
		hash.Mul(hash.Int, prime128.Int)
	}
	s.Set(hash.Int)

	return len(data), nil

}
