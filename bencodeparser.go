package main

import (
	"fmt"
	"strconv"
)

func main() {
	text := "5:hello"

	for i := 0; i < len(text); i++ {
		if text[i] == ':' {
			n, error := strconv.Atoi(text[:i])
			if error != nil {
				fmt.Println("Error converting length:", error)
				return
			}
			content := text[i+1:]
			fmt.Printf("Length: %d, Content: %s\n", n, content)

		}
	}
}

// func parseString(s string) (string, error) {
// 	parts := strings.SplitN(s, ":", 2)
// 	if len(parts) != 2 {
// 		return "", fmt.Errorf("invalid format")
// 	}

// 	length, err := strconv.Atoi(parts[0])
// 	if err != nil {
// 		return "", err
// 	}

// 	if len(parts[1]) < length {
// 		return "", fmt.Errorf("string too short")
// 	}

// 	return parts[1][:length], nil
// }

// func main() {
// 	result, err := parseString("5:hello")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Parsed:", result)
// 	}
// }
