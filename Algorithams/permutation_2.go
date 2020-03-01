package main

import "fmt"
import "sync"

func main() {
	const alphabet = "abcde"

	for str := range GeneratePermutation(alphabet) {
		fmt.Println(str)
	}
}
    "math;func GenPermutation": {
        "prefix": "math;func GeneratePermutation",
        "body": [
            "",
            "/*",
            "GeneratePermutation will generate and retruns the premuation of the given value",
            "NOTE: input must be in sorted order",
            "Example",
            "",
            "// Make sure the alphabet is sorted.",
            "const input = \"abcde\"",
            "for str := range GeneratePermutation(input) {",
            "fmt.Println(str)",
            "*/",
            "func GeneratePermutation(input string) <-chan string {",
            "c := make(chan string, len(input))",
            "go func() {",
            "defer close(c)",
            "if len(input) == 0 {",
            "return",
            "}",
            "",
            "var wg sync.WaitGroup",
            "wg.Add(len(input))",
            "for i := 1; i <= len(input); i++ {",
            "go func(i int) {",
            "",
            "// Perform permutation on a subset of the alphabet.",
            "Word(input[:i]).permute(c)",
            "",
            "// Signal Waitgroup that we are done.",
            "wg.Done()",
            "}(i)",
            "}",
            "",
            "wg.Wait()",
            "}()",
            "",
            "return c",
            "}",
            "",
            "type Word []rune",
            "",
            "// Permute generates all possible combinations of",
            "// the current word. This assumes the runes are sorted.",
            "func (w Word) permute(out chan<- string) {",
            "if len(w) <= 1 {",
            "out <- string(w)",
            "return",
            "}",
            "",
            "// Write first result manually.",
            "out <- string(w)",
            "",
            "// Find and print all remaining permutations.",
            "for w.next() {",
            "out <- string(w)",
            "}",
            "}",
            "",
            "// next performs a single permutation by shuffling characters around.",
            "// Returns false if there are no more new permutations.",
            "func (w Word) next() bool {",
            "var left, right int",
            "left = len(w) - 2",
            "for w[left] >= w[left+1] && left >= 1 {",
            "left--",
            "}",
            "",
            "if left == 0 && w[left] >= w[left+1] {",
            "return false",
            "}",
            "",
            "right = len(w) - 1",
            "for w[left] >= w[right] {",
            "right--",
            "}",
            "",
            "w[left], w[right] = w[right], w[left]",
            "left++",
            "right = len(w) - 1",
            "for left < right {",
            "w[left], w[right] = w[right], w[left]",
            "left++",
            "right--",
            "}",
            "",
            "return true",
            "}",
            ""
        ],
        "description": "func GeneratePermutation "
    }