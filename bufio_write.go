package main

import (
	"bufio"
	"os"
)

func main()  {
	writer := bufio.NewWriter(os.Stdout)

	writer.WriteString("Hello world\n")
	writer.WriteString("Selamat Belajar\n")
	writer.Flush()
}