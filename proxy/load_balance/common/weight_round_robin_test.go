package common

import (
	"fmt"
	"testing"
)

func TestLB(t *testing.T) {
	rb := &WeightRoundRobinBalance{}
	rb.Add("127.0.0.1:2003", "10") //0
	rb.Add("127.0.0.1:2004", "20") //1
	rb.Add("127.0.0.1:2005", "40") //2
	rb.Add("127.0.0.1:2006", "20") //3
	rb.Add("127.0.0.1:2007", "10") //4

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
