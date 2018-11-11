package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

const (
	// alphabet is the modified base58 alphabet used by Bitcoin.
	alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	alphabetIdx0 = '1'
)

var b58 = [256]byte{
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 0, 1, 2, 3, 4, 5, 6,
	7, 8, 255, 255, 255, 255, 255, 255,
	255, 9, 10, 11, 12, 13, 14, 15,
	16, 255, 17, 18, 19, 20, 21, 255,
	22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 255, 255, 255, 255, 255,
	255, 33, 34, 35, 36, 37, 38, 39,
	40, 41, 42, 43, 255, 44, 45, 46,
	47, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
	255, 255, 255, 255, 255, 255, 255, 255,
}

//go:generate go run genalphabet.go

var bigRadix = big.NewInt(58)
var bigZero = big.NewInt(0)

// Decode decodes a modified base58 string to a byte slice.
func Decode(b string) []byte {
	answer := big.NewInt(0)
	j := big.NewInt(1)

	scratch := new(big.Int)
	for i := len(b) - 1; i >= 0; i-- {
		//字符，ascii码表的简版-->得到字符代表的值(0，1,2，..57)
		tmp := b58[b[i]]
		//出现不该出现的字符
		if tmp == 255 {
			return []byte("")
		}
		scratch.SetInt64(int64(tmp))

		//scratch = j*scratch
		scratch.Mul(j, scratch)

		answer.Add(answer, scratch)
		//每次进位都要乘上58
		j.Mul(j, bigRadix)
	}

	//得到大端的字节序
	tmpval := answer.Bytes()

	var numZeros int
	for numZeros = 0; numZeros < len(b); numZeros++ {
		//得到高位0的位数
		if b[numZeros] != alphabetIdx0 {
			break
		}
	}
	//得到原来数字的长度
	flen := numZeros + len(tmpval)
	//构造一个新地存放结果的空间
	val := make([]byte, flen, flen)
	copy(val[numZeros:], tmpval)

	return val
}

// Encode encodes a byte slice to a modified base58 string.
func Encode(b []byte) string {
	x := new(big.Int)
	//将b解释为大端存储
	x.SetBytes(b)

	//Base58编码可以表示的比特位数为Log258 {\displaystyle \approx } \approx5.858bit。经过Base58编码的数据为原始的数据长度的1.37倍
	answer := make([]byte, 0, len(b)*136/100)

	for x.Cmp(bigZero) > 0 {
		mod := new(big.Int)
		//x除于58的余数mod，并将商赋值给x
		x.DivMod(x, bigRadix, mod)
		answer = append(answer, alphabet[mod.Int64()])
	}

	// leading zero bytes
	//因为如果高位为0，0除任何数为0，可以直接设置为‘1’
	for _, i := range b {
		if i != 0 {
			break
		}
		answer = append(answer, alphabetIdx0)
	}

	// reverse
	//因为之前先附加低位的，后附加高位的，所以需要翻转
	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)
}
func main() {
	if true {
		fmt.Println("--1.比特币base58---------------")
		//[201 204 160 110 48 128 185 97 248 106 229 16 111 62 163 92 162 61 160 114 98 148 125 61 10 158 246 26 81 145 51 204 132 1 149 208 141 255]
		//bytes := Decode("WrQWSQZwCCsuBbbyxfhmj3RaKVND1sWtUvUT7J8aAFKEEJ43RNaJ")
		bytes := Decode("5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreAnchuDf")
		fmt.Println(bytes)
		str := Encode(bytes)
		fmt.Println(str)
	}
	if true {
		fmt.Println("--2.比特币base58---------------")
		//[201 204 160 110 48 128 185 97 248 106 229 16 111 62 163 92 162 61 160 114 98 148 125 61 10 158 246 26 81 145 51 204 132 1 149 208 141 255]
		bytes := Decode("KwDiBf89QgGbjEhKnhXJuH7LrciVrZi3qYjgd9M7rFU73sVHnoWn")
		fmt.Println(bytes)
		str := Encode(bytes)
		fmt.Println(str)
	}
	if true {
		fmt.Println("--3.比特币sha256---------------")
		h := sha256.New()
		h.Write([]byte{128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1})
		bs := h.Sum(nil)
		fmt.Println(bs)
		//h2 := sha256.New()
		h2 := h
		h2.Reset()
		h2.Write(bs)
		bs2 := h2.Sum(nil)
		fmt.Println(bs2)
	}
}
