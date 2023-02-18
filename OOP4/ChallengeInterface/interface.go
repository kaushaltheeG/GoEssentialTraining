package main

import (
	"fmt"
	"io"
	"os"
)

//Capper implents io.Writer and turns everything to uppercase
type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	diff := byte('a'-'A')

	out := make([]byte, len(p)) //making new byte slice to not override og 
	for i,c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c 
	}

	return c.wtr.Write(out)
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "hello there")
}