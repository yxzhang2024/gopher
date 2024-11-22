package library3

type MyTLSConfig struct {
	InsecureSkipVerify bool
}

func (MyTLSConfig) Conn() {
	println("Assuming a TLS connection is established")
}
