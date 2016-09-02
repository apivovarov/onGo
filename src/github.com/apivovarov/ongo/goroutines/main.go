package main

import "fmt"
import "os"

func usage() {
  fmt.Println("Usage: bin/goroutines f1|f2")
}

func main() {
  if len(os.Args) != 2 {
    usage()
    return
  }

  switch os.Args[1] {
    case "f1":
      f1_run()
    case "f2":
      f2_run()
    default:
      usage()
  }
}
