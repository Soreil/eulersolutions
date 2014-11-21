package main

import (
	"bytes"
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
	//This "CSV" file only has comma's, no need for anything heavy handed.
	noComma := bytes.Split(file, []byte(","))
	//We need to convert the ASCII characters to integers in this case, the input file is a bit stupid
	noAscii := make([]byte, len(noComma))
	for i, v := range noComma {
		tmp, err := strconv.ParseUint(string(v), 0, 8)
		if err != nil {
			panic(err)
		}
		noAscii[i] = byte(uint8(tmp))
	}
	//Turn a two level slice in to a single slice
	//decoded := bytes.Join(noAscii, nil)
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

	frequencies := frequencies(splits[0].Bytes())
	sort.Sort(sort.Reverse(frequencies))

	//Check every character as an Xor candidate that has at least one hit
	for i := 0; frequencies[i].c != 0; i++ {
		decoded := xorTest(splits[0].Bytes(), frequencies[i].r)
	}
}

func frequencies(s []byte) byteCounts {
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
