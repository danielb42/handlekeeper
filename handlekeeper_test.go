package handlekeeper

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandlekeeper(t *testing.T) {
	hk, _ := NewHandlekeeper("testfile")
	bytes := make([]byte, 5)

	hk.Handle.WriteString("line1")
	hk.Handle.ReadAt(bytes, 0)
	assert.Equal(t, "line1", string(bytes))
	os.Remove("testfile")

	time.Sleep(time.Millisecond)

	hk.Handle.WriteString("line2")
	hk.Handle.ReadAt(bytes, 0)
	assert.Equal(t, "line2", string(bytes))
	os.Remove("testfile")
}
