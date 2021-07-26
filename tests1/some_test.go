package tests1

import (
	"testing"
	"time"
)

func TestOne(t *testing.T) {
	t.Log("sleep 1")
	time.Sleep(time.Second)
	t.Log("sleep 2")
	time.Sleep(time.Second)
	t.Log("sleep 3")
	time.Sleep(time.Second)
	t.Log("done")
}
