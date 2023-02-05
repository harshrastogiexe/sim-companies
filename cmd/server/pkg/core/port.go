package core

import "fmt"

type port uint

func (p port) String() string {
	return fmt.Sprintf(":%d", p)
}
