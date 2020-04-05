package public

import (
	"fmt"
	"testing"
)

func TestJWTEncode(t *testing.T) {
	tests := []struct {
		name string
		foo  string
	}{
		{name: "test1", foo: "foo",},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := Encode(tt.foo)
			if err != nil {
				t.Fatal(err)
			}
			result,err:=Decode(token)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println("result",result)
			if tt.foo!=result{
				t.Fatal("tt.foo!=result")
			}
		})
	}
}

