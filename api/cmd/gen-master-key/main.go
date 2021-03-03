package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/insecurecleartextkeyset"
	"github.com/google/tink/go/keyset"
)

func main() {
	kek := newKey()
	printKey(kek, true)
}

func newKey() *keyset.Handle {
	kh, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		log.Fatal(err)
	}

	return kh
}

func printKey(k *keyset.Handle, binary bool) {
	var w keyset.Writer
	if binary {
		w = keyset.NewBinaryWriter(base64.NewEncoder(base64.RawStdEncoding, os.Stdout))
	} else {
		w = keyset.NewJSONWriter(os.Stdout)
	}

	err := insecurecleartextkeyset.Write(k, w)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("")
}
