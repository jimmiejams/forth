package interpreter

import (
    "bufio"
    "fmt"
    "os"
)

func Repl() {
    scanner := bufio.NewScanner (os.Stdin)
    for scanner.Scan() {
        input := scanner.Text()
        fmt.Printf ("output: %s\n", input)
    }
}