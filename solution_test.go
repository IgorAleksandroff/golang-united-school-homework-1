package solution

import (
	"testing"

	"github.com/kyokomi/emoji/v2"
)

func TestGetMessage(t *testing.T) {
	s := emoji.Sprint("Hello :world_map:!")
	v := GetMessage()
	if v != s {
			t.Error("Expected ", s,", got ", v)
	}
}