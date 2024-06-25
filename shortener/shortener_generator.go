package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha256Of(input string) []byte {
	hash := sha256.New()
	hash.Write([]byte(input))
	return hash.Sum(nil)
}

func base58Endcoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(encoded)
}

func GenerateShortUrl(initialLink string, userId string) string {
	urlHashBytes := sha256Of(initialLink + userId)
	generatedBigNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58Endcoded([]byte(fmt.Sprint(generatedBigNumber)))

	return finalString[:8]
}
