package main

import (
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
    // Tokenize first
    words := Tokenize(text)

    // Core transformations on tokens
    words = ConvertHexAndBin(words)
    words = FixArticles(words)
    words = ApplyCaseRules(words)

    // Rebuild text from tokens
    rebuilt := strings.Join(words, " ")

    // Final formatting passes (spacing, quotes) last
    rebuilt = ApplyPunctuationRules(rebuilt)
    finalText := FixQuotes(rebuilt)
    return finalText
}
