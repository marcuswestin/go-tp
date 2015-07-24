package tokens

type Hash uint32

const (
	offset32 = Hash(2166136261)
	prime32  = Hash(16777619)
)

// https://golang.org/src/hash/fnv/fnv.go
func HashTokens(tokenIDs Ids) Hash {
	hash := offset32
	for _, tokenID := range tokenIDs {
		hash ^= Hash(tokenID)
		hash *= prime32
	}
	return hash
}

func Equal(tokenIDsA, tokenIDsB Ids) bool {
	if len(tokenIDsA) != len(tokenIDsB) {
		return false
	}
	for i := range tokenIDsA {
		if tokenIDsA[i] != tokenIDsB[i] {
			return false
		}
	}
	return true
}

// func tokensHashFnv1(tokenIDs Ids) Hash {
// 	hash := fnv.New32()
// 	hash.Write(bytes())
// 	return Hash(hash.Sum32())
// }

// const SIZEOF_INT32 = 4 // bytes
// func tokensBytes(tokens.IDs) []byte {
// 	if uint(unsafe.Sizeof(tokens[0])) != SIZEOF_INT32 {
// 		panic("SIZE MISMATCH")
// 	}
// 	buf := make([]byte, len(tokens)*size)
// 	for i := range tokens {
// 		binary.LittleEndian.PutUint32(buf[i*SIZEOF_INT32 : (i+1)*SIZEOF_INT32], tokens[i])
// 	}
// 	return buf
// }

// func ngramCantorPairId(tokenIDs Ids) (result NGramId) {
// 	result = ngramIdCantorPair(NGramId(tokenIDs[0]), tokenIDs[1])
// 	for i := 2; i < len(tokenIDs); i++ {
// 		result = ngramIdCantorPair(result, tokenIDs[i])
// 	}
// 	return NGramId(result)
// }

// func ngramIdCantorPair(x NGramId, _y Id) NGramId {
// 	y := NGramId(_y)
// 	return NGramId(((x+y)*(x+y+1))/2 + y)
// }
