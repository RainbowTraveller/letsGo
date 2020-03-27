/*This is demo for implementing interface
 */

package main

import (
	"fmt"
	"io"
	"os"
)

type Capper struct {
	Writer io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	//byte distance between upper and lower case
	diff := byte('a' - 'A')
	//Allocate memory
	out := make([]byte, len(p))
	for i, c := range p {
		if c >= 'a' && c <= 'z' {
			c -= diff
		}
		out[i] = c
	}
	return c.Writer.Write(out)
}

func main() {

	up := &Capper{os.Stdout}
	fmt.Fprintln(up, "please convert to upper case")
}
