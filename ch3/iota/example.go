package iota

const (
	_   = 1 << (10 * iota) // discard 1 when iota is zero
	kiB                    // assumes next expression of 1 << ( 10) = 2^ 10 = 1024
	miB                    // 2 ^ 20
	giB
	tiB
	piB
	eiB
	ziB
	yiB // 2 ^ 80
)
