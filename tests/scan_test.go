package tests

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestScan1(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lineParts := strings.Split(scanner.Text(), "~")
		fmt.Println(lineParts)
	}

	fmt.Println("gggg")
}
