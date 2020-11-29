package module

import "os"

func GetSentence(alpha *Alphabet) (result string) {
	args := os.Args[1]
	sentence := make([][8]string, len(args))
	for i := 0; i < len(sentence); i++ {
		sentence[i] = alpha.LetterAscii[string(args[i])]
	}
	for i := 0; i < len(sentence[i]); i++ {
		for j := 0; j < len(sentence); j++ {
			for k := 0; k < len(sentence[j][i]); k++ {
				if sentence[j][i][k] != '\n' && sentence[j][i][k] != '\r' {
					result += string(sentence[j][i][k])
				}
			}
		}
		result += string('\n')
	}
	return
}
