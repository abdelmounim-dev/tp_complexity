package tp1

import (
	"fmt"
	"time"
  "os"
)

var numbers = []int {
  1000003,
  2000003,
  4000037,
  8000009,
  16000057,
  32000011,
  64000031,
  128000003,
  256000001,
  512000009,
  1024000009,
  2048000011,
}

func nbDevisors(n int) int {
  count := 0
  for i := 1; i <= n; i++ {
    if n % i == 0 {
      count++
    }
  }
  return count
}

func isPrime1(n int) bool {
  return nbDevisors(n) == 2
}

func main () {
  os.WriteFile("primalityTest1.csv", []byte(fmt.Sprintln("number, execTime")), 0644)
  for _, n := range numbers {
    start := time.Now()
    isPrime1(n)
    execTime := time.Since(start)
    fmt.Printf("isPrime1(%d) took %s\n", n, execTime)
    os.WriteFile("primalityTest1.csv", []byte(fmt.Sprintf("%d %d\n", n, execTime.Milliseconds())), 0644)  
  }
}
