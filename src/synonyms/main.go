package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Gin58/golanglang/src/thesaurus"
	"github.com/joho/godotenv"
)

func main() {
		err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurus := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("%qの類語検索エラー: %v", word, err)
		}
		if len(syns) == 0 {
			log.Fatalln("類語なし")
		}
		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}