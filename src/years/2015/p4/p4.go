package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	// "time"
)

func main() {
	const secret = "yzbqklnj"
	const prefix_five = "00000"
	const prefix_six = "000000"
	five_found, six_found := false, false

	i := 0
	for {
		// fmt.Println(i)

		key_str := secret + strconv.Itoa(i)
		// fmt.Println("\t", key_str)

		key_md5 := md5.Sum([]byte(key_str))
		// fmt.Println("\t", key_md5)

		hash := hex.EncodeToString(key_md5[:]) // just slicing was not enough.
		// fmt.Println("\t", hash)

		if !five_found && strings.HasPrefix(hash, prefix_five) {
			fmt.Println("five zeros --->", i)
			five_found = true
		}

		if !six_found && strings.HasPrefix(hash, prefix_six) {
			fmt.Println("six zeros --->", i)
			six_found = true
		}

		if five_found && six_found {
			break
		}

		// time.Sleep(time.Millisecond * 100)
		i++
	}
}
