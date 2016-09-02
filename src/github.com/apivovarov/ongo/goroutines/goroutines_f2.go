package main

import "fmt"
import "time"

func prod(c, m chan int) {
  for i := 0; i < 20; i++ {
    //time.Sleep(200 * time.Millisecond)
    fmt.Printf("prod: %d\n", i)
    c <- i
  }
  c <- -1 // EOF to consumer
  m <- -1 // EOF to main
}

func cons(c chan int) {
  i := 0;
  for i >= 0 {
    i = <-c
    fmt.Printf("cons: %d\n", i)
  }
}

func f2_run() {
  c := make(chan int)
  m := make(chan int)
  go cons(c)
  go prod(c, m)

  m_val := 0
  // wait for prod to finish
  for m_val != -1 {
    m_val = <-m
  }
  time.Sleep(1 * time.Millisecond)
  fmt.Println("End prod->cons")
}
