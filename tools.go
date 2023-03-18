package toolkit

import "crypto/rand"

const RANDOM_STRING_SOURCE = "ABCDEFGNOPFQEASDFASabcdedifkgnasoiereasfdadf0347235872-98765489"

//Type used to instantiate the module. Any variable of this type will have all access to all the methods with receiver *Tools
type Tools struct {
}

type UploadedFile struct {
	NewFileName      string
	OriginalFileName string
	FileSize         int
}

//This function returns a string of random characters of length n
func (recv *Tools) RandomString(n int) string {
	s, r := make([]rune, n), []rune(RANDOM_STRING_SOURCE)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}
	return string(s)
}

func TestIng() string {
	return "ABC"
}
