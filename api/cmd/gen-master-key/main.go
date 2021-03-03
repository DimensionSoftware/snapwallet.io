package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"github.com/google/tink/go/aead"
	"github.com/google/tink/go/insecurecleartextkeyset"
	"github.com/google/tink/go/keyset"
)

func main() {
	var buf bytes.Buffer

	kekH := newKey()
	kek, err := aead.New(kekH)
	if err != nil {
		log.Fatal(err)
	}

	dekH1 := newKey()

	buf.Reset()
	dekH1.Write(keyset.NewJSONWriter(&buf), kek)
	if err != nil {
		log.Fatal(err)
	}

	dekH2, err := keyset.Read(keyset.NewJSONReader(&buf), kek)
	if err != nil {
		log.Fatal(err)
	}

	dek1, err := aead.New(dekH1)
	if err != nil {
		log.Fatal(err)
	}
	dek2, err := aead.New(dekH2)
	if err != nil {
		log.Fatal(err)
	}

	encData, err := dek1.Encrypt([]byte("hola"), []byte{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(encData)

	data, err := dek2.Decrypt(encData, []byte{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	//printKey(kek, false)
	//printKey(dek, false)
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
