package xsd

import "fmt"

func sfmt(s string, a ...interface{}) string {
	return fmt.Sprintf(s, a...)
}
