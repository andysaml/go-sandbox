package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func sort(arr []string) {

}

func main() {
	fmt.Println("Enter digits:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	digits := strings.Fields(text)
	sort(digits)
}
