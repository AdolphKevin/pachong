package main

import (
	"fmt"
	"testing"
)

func TestRemove(t *testing.T) {
	s := removeUnnecessaryString(" 璧山县(重庆市重庆市郊县下辖的县)")
	fmt.Println(s)
}
