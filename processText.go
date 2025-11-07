package main

import (
	"fmt"
	"strings"
)

// processText is the main text-processing pipeline
// it sequentially applies all transformation stages
// func processText(text string) string {

// 	words := tokenize(text)

// 	fmt.Printf("debug raw tokens: %#v", words)

// 	words = convertHexAndBin(words)

// 	words = fixArticles(words)

// 	words = applyCaseRules(words)

// 	formattedText := applyPunctuationRules(words)

// 	finalText := fixQuotes(formattedText)

// 	fmt.Println("DEBUG tokens: ", finalText)

// 	return finalText

// }

func ProcessText(text string) string {
    // 1️⃣ Pre-cleaning
    prepped := fixQuotes(text)
    prepped = applyPunctuationRules(prepped)

    // 2️⃣ Tokenization
    words := tokenize(prepped)
    fmt.Printf("debug raw tokens: %#v", words)

    // 3️⃣ Transformations
    words = convertHexAndBin(words)
    words = fixArticles(words)
    words = applyCaseRules(words)

    // 4️⃣ Rebuild text
    finalText := strings.Join(words, " ")

    fmt.Println("DEBUG tokens:", finalText)
    return finalText
}