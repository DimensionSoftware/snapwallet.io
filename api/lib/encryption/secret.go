package encryption

// Secret ..
type Secret interface {
	Unseal() interface{}
	Seal() []byte
}
