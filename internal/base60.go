package internal

import (
	"fmt"
	"log"
	"slices"
	"strings"
)

// chars represents the set of valid characters for Base60 encoding
var chars = []rune("䷁䷖䷇䷓䷏䷢䷬䷋" +
	"䷎䷳䷦䷴䷽䷷䷞䷠" +
	"䷆䷃䷜䷺䷧䷿䷮䷅" +
	"䷭䷑䷯䷸䷟䷱䷛䷫" +
	"䷗䷚䷂䷩䷲䷔䷐䷘" +
	"䷣䷕䷾䷤䷶䷝䷰䷌" +
	"䷒䷨䷻䷼䷵䷥䷹䷉" +
	"䷊䷙䷄䷈")

var gan = []rune("甲乙丙丁戊己庚辛壬癸")
var zhi = []rune("子丑寅卯辰巳午未申酉戌亥")

var base60Lookup = func() map[rune]uint64 {
	m := make(map[rune]uint64)
	for i, c := range chars {
		m[c] = uint64(i)
	}
	return m
}()

var ganzhiLookup = func() map[rune]string {
	m := make(map[rune]string)
	for i := 0; i < 60; i++ {
		m[chars[i]] = string(gan[i%10]) + string(zhi[i%12])
	}
	return m
}()

type Base60 interface {
	Encode(in []byte) (out string)
	Decode(in string) (out []byte, err error)
	Verify(in string) bool
	Human(in string) (out string)
}

type base60 struct{}

func NewBase60() Base60 {
	g := new(base60)
	return g
}

// bytesToUint64 converts a byte slice to uint64
func bytesToUint64(bytes []byte) uint64 {
	var num uint64
	for _, b := range bytes {
		num = (num << 8) | uint64(b)
	}
	return num
}

// uint64ToBase60 converts an uint64 to a Base60 encoded string
func uint64ToBase60(num uint64) string {
	if num == 0 {
		return string(chars[0])
	}

	var runes []rune
	for num > 0 {
		remainder := num % 60
		runes = append(runes, chars[remainder])
		num /= 60
	}

	// Reverse the runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func (g *base60) Encode(in []byte) (out string) {
	var encoded []string
	inLen := len(in)
	chunkSize := 8 // 8 bytes for uint64

	for i := 0; i < inLen; i += chunkSize {
		end := i + chunkSize
		if end > inLen {
			end = inLen
		}
		chunk := in[i:end]
		// Convert chunk to uint64
		num := bytesToUint64(chunk)
		log.Println(num)
		// Convert uint64 to Base60
		encoded = append(encoded, uint64ToBase60(num))
	}
	return strings.Join(encoded, "")
}

// base60ToUint64 converts a Base60 encoded runes to an uint64
func base60ToUint64(s []rune) (uint64, error) {
	var num uint64
	for _, c := range s {
		val, exists := base60Lookup[c]
		if !exists {
			return 0, fmt.Errorf("invalid character %q in Base60 string", c)
		}
		num = num*60 + val
	}
	return num, nil
}

// uint64ToBytes converts an uint64 to a byte slice
func uint64ToBytes(num uint64) []byte {
	var byteCount int
	for i := uint(0); i < 8; i++ {
		if (num >> (i * 8)) == 0 {
			break
		}
		byteCount = int(i)
	}
	b := make([]byte, byteCount+1)
	for i := byteCount; i >= 0; i-- {
		b[i] = byte(num & 0xFF)
		num >>= 8
	}
	return b
}

// Decode converts a Base60 encoded string back to the original byte slice
func (g *base60) Decode(in string) (out []byte, err error) {
	// Split the encoded string into chunks of Base60 encoded uint64
	chunkSize := 11 // Each uint64 can be represented by up to 11 Base60 characters
	inRunes := []rune(in)
	inLen := len(inRunes)

	for i := 0; i < inLen; i += chunkSize {
		end := i + chunkSize
		if end > inLen {
			end = inLen
		}
		chunk := inRunes[i:end]
		// Convert Base60 chunk to uint64
		var num uint64
		num, err = base60ToUint64(chunk)
		if err != nil {
			return nil, err
		}
		// Convert uint64 to byte slice and append to output
		out = append(out, uint64ToBytes(num)...)
	}
	return out, nil
}

// Verify checks if a given string is a valid Base60 encoded string
func (g *base60) Verify(str string) bool {
	for _, i := range []rune(str) {
		if !slices.Contains(chars, i) {
			return false
		}
	}
	return true
}

func (g *base60) Human(in string) (out string) {
	var encoded []string
	ganzhi := ganzhiLookup
	for _, i := range []rune(in) {
		encoded = append(encoded, ganzhi[i])
	}
	return strings.Join(encoded, "")
}
