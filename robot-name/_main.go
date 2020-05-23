/* package main

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	math_rand "math/rand"
)

func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	math_rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}
*/

package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// Adapted from https://elithrar.github.io/article/generating-secure-random-numbers-crypto-rand/

func init() {
	assertAvailablePRNG()
}

func assertAvailablePRNG() {
	// Assert that a cryptographically secure PRNG is available.
	// Panic otherwise.
	buf := make([]byte, 1)

	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		panic(fmt.Sprintf("crypto/rand is unavailable: Read() failed with %#v", err))
	}
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	bytes, err := GenerateRandomBytes(n)
	if err != nil {
		return "", err
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes), nil
}

// GenerateRandomStringURLSafe returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomStringURLSafe(n int) (string, error) {
	b, err := GenerateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}

func main() {
	// Example: this will give us a 44 byte, base64 encoded output
	/* token, err := GenerateRandomStringURLSafe(32)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		panic(err)
	}
	fmt.Println(token) */

	// Example: this will give us a 32 byte output
	token, err := GenerateRandomString(2)
	if err != nil {
		// Serve an appropriately vague error to the
		// user, but log the details internally.
		panic(err)
	}
	fmt.Println(token)
}


	/* var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	// restrictions on first uppercase characters
	const charConstrain = 2

	for i := 0; i < charConstrain; i++ {
		// A = ascii decimal(65) & Z = ascii decimal(90)
		// this can be avoided by an array of all uppercase alphabet
		// but i like playing with ascii values.
		intVal := seededRand.Intn(90-65+1) + 65
		name = name + string([]byte{byte(intVal)})
	}

	// restrictions on numerals = 3 digit
	randomNum := seededRand.Intn(999-100+1) + 100
	name = name + strconv.Itoa(randomNum) */