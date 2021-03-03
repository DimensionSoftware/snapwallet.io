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
	kh, err := keyset.NewHandle(aead.AES256GCMKeyTemplate())
	if err != nil {
		log.Fatal(err)
	}

	w := keyset.NewBinaryWriter(base64.NewEncoder(base64.RawStdEncoding, os.Stdout))
	err = insecurecleartextkeyset.Write(kh, w)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("")

}
