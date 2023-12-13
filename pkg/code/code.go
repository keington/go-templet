package code

import "fmt"

/**
 * @author: x.gallagher.anderson@gmail.com
 * @time: 2023/11/20 21:50
 * @file: code.go
 * @description: code
 */

// statusCode is an implementer for interface Code for internal usage only.
type statusCode struct {
	code    int    // Error code, usually an integer.
	message string // Brief message for this error code.
}

// Code returns the integer number of current error code.
func (c statusCode) Code() int {
	return c.code
}

// Message returns the brief message for current error code.
func (c statusCode) Message() string {
	return c.message
}

// String returns current error code as a string.
func (c statusCode) String() string {

	if c.message != "" {
		return fmt.Sprintf(`%d:%s`, c.code, c.message)
	}
	return fmt.Sprintf(`%d`, c.code)
}
