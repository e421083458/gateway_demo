package load_balance

import (
	"fmt"
	"testing"
)

func TestRandomBalance(t *testing.T) {
	rb := &RandomBalance{}
	rb.Add("127.0.0.1:2003") //0
	rb.Add("127.0.0.1:2004") //1
	rb.Add("127.0.0.1:2005") //2
	rb.Add("127.0.0.1:2006") //3
	rb.Add("127.0.0.1:2007") //4

	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
	fmt.Println(rb.Next())
}
