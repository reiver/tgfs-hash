package hash

import (
	"golang.org/x/crypto/sha3"

	"errors"
	"fmt"
	"io"
)

func Func(r io.Reader) (digest []byte, n int64, err error) {

	hasher := sha3.New512()
	if nil == hasher {
		return nil, 0, errors.New("ERROR: Could not initialize hash function")
	}

	n, err = io.Copy(hasher, r)
	if nil != err {
		return nil, n, fmt.Errorf("ERROR: Could not generate hash function digest: %s", err)
	}

	digest = hasher.Sum(nil)

	return digest, n, nil
}
