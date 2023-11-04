package decrypt

func Decrypt(data []byte) []byte {
	mask := []byte("Growatt")
	out := make([]byte, len(data))
	copy(out, data)

	maskIndex := 0
	for i := 8; i < len(data); i++ {
		out[i] ^= mask[maskIndex]
		maskIndex = (maskIndex + 1) % len(mask)
	}
	return out
}
