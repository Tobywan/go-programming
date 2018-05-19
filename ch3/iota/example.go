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

// KB = 10^3
// YB = 10^24
// How can we generate consts compactly?
// given that 10^(3*iota) won't work
const (
	KB = 1000    // assumes next expression of 1 << ( 10) = 2^ 10 = 1024
	MB = KB * KB // 2 ^ 20
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB // 2 ^ 80
)
