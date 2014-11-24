package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
)

type byteCount struct {
	r byte // the byte to be counted
	c int  // a counter
}

type byteCounts []byteCount // method slave

func (r byteCounts) Len() int {
	return len(r)
}

func (r byteCounts) Less(i, j int) bool {
	if r[i].c < r[j].c {
		return true
	}
	return false
}

func (r byteCounts) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

const keySize = 3

func main() {
	//This program takes a CSV file
	file, err := ioutil.ReadFile("p059_cipher.txt")
	if err != nil {
		panic(err)
	}
	//This "CSV" file only has commas, no need for anything heavy handed.
	noComma := bytes.Split(file, []byte(","))
	//We need to remove the newline character.
	noComma[len(noComma)-1] = bytes.TrimSuffix(noComma[len(noComma)-1], []byte("\n"))
	//We need to convert the ASCII characters to integers in this case, the input file is a bit stupid
	noAscii := make([]byte, len(noComma))
	for i, v := range noComma {
		tmp, err := strconv.ParseUint(string(v), 0, 8)
		//ie value is bigger than a byte
		if err != nil {
			panic(err)
		}
		noAscii[i] = byte(uint8(tmp))
	}
	//Turn the slice in to a bytes.Buffer
	buffer := bytes.NewBuffer(noAscii)

	//Buffers for the individual Xor testing
	splits := make([]bytes.Buffer, keySize)

	// Break label exists to not have a multi level break manually
Loop:
	//While there are bytes in the buffer keep splitting
	for {
		for i := range splits {
			tmp, err := buffer.ReadByte()
			if err != nil {
				//We expect to hit EOF at some point.
				if err != io.EOF {
					panic(err)
				} else {
					break Loop
				}
			}
			err = splits[i].WriteByte(tmp)
			if err != nil {
				panic(err)
			}
		}
	}

	frequencies := make([]byteCounts, len(splits))
	for i := range frequencies {
		frequencies[i] = byteFrequencies(splits[i].Bytes())
		sort.Sort(sort.Reverse(frequencies[i]))
	}
	//unXor holds the unxored values
	unXor := make([]bytes.Buffer, keySize)
	for i := range splits {
		decoded := xorTest(splits[i].Bytes(), frequencies[i][0].r)
		unXor[i].Write(decoded)
	}
	// Break label exists to not have a multi level break manually
Loop2:
	//While there are bytes in the buffer keep splitting
	for {
		for i := range unXor {
			tmp, err := unXor[i].ReadByte()
			if err != nil {
				//We expect to hit EOF at some point.
				if err != io.EOF {
					panic(err)
				} else {
					break Loop2
				}
			}
			err = buffer.WriteByte(tmp)
			if err != nil {
				panic(err)
			}
		}
	}
	fmt.Println(string(buffer.Bytes()))
}

func byteFrequencies(s []byte) byteCounts {
	//Only take 7 bit ASCII in to account.
	ret := make([]byteCount, 128)
	for i := range ret {
		ret[i].r = byte(i)
	}
	for _, v := range s {
		ret[v].c++
	}
	return ret
}

func xorTest(s []byte, cmp byte) []byte {
	x := make([]byte, len(s))
	for i, v := range s {
		x[i] = cmp ^ v
	}
	return x
}
