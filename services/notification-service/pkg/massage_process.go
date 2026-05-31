package pkg

import (
	"fmt"
)


func Process(body string) error {
	fmt.Println("processing:", body)

	
	return fmt.Errorf("fail")
}
