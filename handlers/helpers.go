package handlers

import (
	"fmt"
)

// Helper Func
func CheckError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}
}
