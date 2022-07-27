package load_balance

import (
	"fmt"
	"testing"
)

func TestRandomBalance(t *testing.T) {
	rb := RandomBalance{}
	rb.Add("127.0.0.1:2003", "127.0.0.1:2004", "127.0.0.1:2005", "127.0.0.1:2006", "127.0.0.1:2007")
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
}
