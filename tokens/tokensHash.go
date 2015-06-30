package tokens

type Hash uint32

const (
	offset32 = Hash(2166136261)
	prime32  = Hash(16777619)
)

// https://golang.org/src/hash/fnv/fnv.go
func HashTokens(tokenIds []TokenId) Hash {
	hash := offset32
	for _, tokenId := range tokenIds {
		hash ^= Hash(tokenId)
		hash *= prime32
	}
	return hash
}

func Equal(tokenIdsA, tokenIdsB []TokenId) bool {
	if len(tokenIdsA) != len(tokenIdsB) {
		return false
	}
	for i := range tokenIdsA {
		if tokenIdsA[i] != tokenIdsB[i] {
			return false
		}
	}
	return true
}

// func tokensHashFnv1(tokenIds []TokenId) Hash {
// 	hash := fnv.New32()
// 	hash.Write(bytes())
// 	return Hash(hash.Sum32())
// }

// const SIZEOF_INT32 = 4 // bytes
// func tokensBytes(tokens []TokenId) []byte {
// 	if uint(unsafe.Sizeof(tokens[0])) != SIZEOF_INT32 {
// 		panic("SIZE MISMATCH")
// 	}
// 	buf := make([]byte, len(tokens)*size)
// 	for i := range tokens {
// 		binary.LittleEndian.PutUint32(buf[i*SIZEOF_INT32 : (i+1)*SIZEOF_INT32], tokens[i])
// 	}
// 	return buf
// }

// func ngramCantorPairId(tokenIds []TokenId) (result NGramId) {
// 	result = ngramIdCantorPair(NGramId(tokenIds[0]), tokenIds[1])
// 	for i := 2; i < len(tokenIds); i++ {
// 		result = ngramIdCantorPair(result, tokenIds[i])
// 	}
// 	return NGramId(result)
// }

// func ngramIdCantorPair(x NGramId, _y TokenId) NGramId {
// 	y := NGramId(_y)
// 	return NGramId(((x+y)*(x+y+1))/2 + y)
// }
