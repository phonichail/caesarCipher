package caesarCipher

func CaesarCipher(input string) string {
	cipher := []rune("abcdefghijklmnopqrstuvwxyzæøå,. ")
	message := []rune(input)
	output := ""

	for i := 0; i < len(message); i++ {
		for j := 0; j < len(cipher); j++ {
			if message[i] == cipher[j] {
				if j+4 < len(cipher) {
					output += string(cipher[j+4])
				} else {
					output += string(cipher[j+4-len(cipher)])
				}
				break
			}
		}
	}
	return output
}
