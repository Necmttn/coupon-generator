package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func main() {
	lengthOfCoupon := flag.Int("length", 8, "length of each coupon")
	amountOfCoupon := flag.Int("amount", 1000, "amount of coupons")
	pathForFile := flag.String("path", "./coupons.txt", "Path for Coupon txt file")
	flag.Parse()

	f, err := os.Create(*pathForFile)
	check(err)
	defer f.Close()
	coupons := make(chan string)
	fmt.Println("Amount:", *amountOfCoupon)
	fmt.Println("Length:", *lengthOfCoupon)
	for i := 1; i <= *amountOfCoupon; i++ {
		go func(i int) {
			coupons <- generateRandomString(*lengthOfCoupon)
		}(i)
		writeToFile(f, <-coupons)
	}
	fmt.Println("finished check the path", *pathForFile)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func writeToFile(f *os.File, coupon string) {
	_, err := f.WriteString(coupon + "\n")
	check(err)
	f.Sync()
}

func generateRandomString(n int) string {
	b := make([]byte, n)

	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)

}
