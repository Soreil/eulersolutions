package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"sort"
)

func main() {
	file, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		panic(err)
	}
	ioReader := bytes.NewReader(file)
	csvReader := csv.NewReader(ioReader)
	name, err := csvReader.Read()
	if err != nil {
		panic(err)
	}
	sort.Strings(name)
	length := csvReader.FieldsPerRecord
	var score int64
	for i := 0; i < length; i++ {
		for j, key := 0, 0; j < len(name[i]); j++ {
			key += int(name[i][j] - 'A' + 1)
			if len(name[i])-1 == j {
				score += int64((i + 1) * key)
			}
		}
	}
	fmt.Println(score)
}
