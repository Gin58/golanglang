package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("newno modorichi is nil")
	} else {
		tracer.Trace("こんにちは")
		if buf.String() != "こんにちは\n" {
			t.Errorf("'%s'という誤った値が出力されてます", buf.String())
		}
	}
}