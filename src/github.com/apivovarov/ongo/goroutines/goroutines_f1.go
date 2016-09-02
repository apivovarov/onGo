package main

import "fmt"
import "time"

func f1() {
  for i := 0; i < 5; i++ {
    if i > 0 {time.Sleep(100 * time.Millisecond)}
    fmt.Println(i)
  }
}

func f1_run() {
  go f1()
  time.Sleep(220 * time.Millisecond)
	fmt.Println("Goroutines")
  time.Sleep(600 * time.Millisecond)
  fmt.Println("End f1")
}
