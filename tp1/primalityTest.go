package main

import (
	"fmt"
	"time"
  "os"
  "log"
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

func isPrime2(n int) bool {
  for i := 2; i < n / 2; i++ {
    if n % i == 0 {
      return false
    }
  }
  return true
}

func main () {
  go func() {
    fileName := "primalityTest1.csv"

    f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {
        log.Fatal(err)
    }

    f.Truncate(0)
    f.WriteString("n, execTime\n")
    for _, n := range numbers {
      start := time.Now()
      isPrime1(n)
      execTime := time.Since(start)
      fmt.Printf("isPrime1(%d) took %s\n", n, execTime)
      f.WriteString(fmt.Sprintf("%d, %f\n", n, execTime.Seconds()))
    }
  }()
  go func(){

    fileName := "primalityTest2.csv"

    f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)

    if err != nil {
        log.Fatal(err)
    }

    f.Truncate(0)
    f.WriteString("n, execTime\n")
    for _, n := range numbers {
      start := time.Now()
      isPrime2(n)
      execTime := time.Since(start)
      fmt.Printf("isPrime2(%d) took %s\n", n, execTime)
      f.WriteString(fmt.Sprintf("%d, %f\n", n, execTime.Seconds()))
    }
  }()
  time.Sleep(10 * time.Second)
}
