package main

import (
	"code.google.com/p/go-uuid/uuid"
	"github.com/mpbsbr/fnv"

	"crypto/md5"
	"encoding/hex"
	gfnv "hash/fnv"
	"log"
	"runtime"
	"strings"
	"time"
)

const MaxClearTextIterations = 20
const Margin = 2

func main() {

	logBlankLines(Margin)

	clearTexts := make([]string, MaxClearTextIterations)
	for i := range clearTexts {
		clearTexts[i] = uuid.New()
	}

	validate := clearTexts[len(clearTexts)-1]
	log.Printf("clearText:\t\t\t %s", validate)
	log.Printf("stdlib fnv64a digest:\t %s", hex.EncodeToString(gfnv64aOne(validate)))
	log.Printf("stdlib md5 digest:\t\t %s\n", hex.EncodeToString(md5One(validate)))
	log.Printf("fnv128a digest:\t\t %s", hex.EncodeToString(fnv128aOne(validate)))
	log.Printf("fnv128 digest:\t\t %s", hex.EncodeToString(fnv128One(validate)))

	logBlankLines(2)

	gfnv64aAll(clearTexts)
	md5All(clearTexts)
	logBlankLines(1)
	fnv128aAll(clearTexts)
	fnv128All(clearTexts)

	logBlankLines(Margin)
}

func fnv128All(clearTexts []string) {
	defer timeTrack(time.Now(), funcName(1))
	for _, value := range clearTexts {
		fnv128One(value)
	}
}

func fnv128One(clearText string) []byte {
	hash := fnv.New128()
	hash.Write([]byte(clearText))
	return hash.Sum(nil)
}

func fnv128aAll(clearTexts []string) {
	defer timeTrack(time.Now(), funcName(1))
	for _, value := range clearTexts {
		fnv128aOne(value)
	}
}

func fnv128aOne(clearText string) []byte {
	hash := fnv.New128a()
	hash.Write([]byte(clearText))
	return hash.Sum(nil)
}

func gfnv64aAll(clearTexts []string) {
	defer timeTrack(time.Now(), funcName(1))
	for _, value := range clearTexts {
		gfnv64aOne(value)
	}
}

func gfnv64aOne(clearText string) []byte {
	hash := gfnv.New64a()
	hash.Write([]byte(clearText))
	return hash.Sum(nil)
}

func md5All(clearTexts []string) {
	defer timeTrack(time.Now(), funcName(1))
	for _, value := range clearTexts {
		md5One(value)
	}
}

func md5One(clearText string) []byte {
	hash := md5.New()
	hash.Write([]byte(clearText))
	return hash.Sum(nil)
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%-20s %s", name, elapsed/MaxClearTextIterations*10)
}

func funcName(depth int) string {
	pc, _, _, _ := runtime.Caller(depth)
	return runtime.FuncForPC(pc).Name()
}

func logBlankLines(count int) {
	log.SetFlags(0)
	log.Printf(strings.Repeat("\n", count))
	log.SetFlags(log.LstdFlags)
}
