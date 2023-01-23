package key

type Key interface {
	Bytes() []byte
	Public() (Key, error)
}
