package builder

import "testing"

func TestBuilder1(t *testing.T) {
	builder := &Builder1{}
	director := NewDirector(builder)
	director.Construct()
	result := builder.GetResult()
	if result != "123" {
		t.Fatalf("Build1 fail expect 123 acture %s", result)
	}
}

func TestBuilder2(t *testing.T) {
	builder := &Builder2{}
	director := NewDirector(builder)
	director.Construct()
	result := builder.GetResult()
	if result != 6 {
		t.Fatalf("Build2 fail expect 6 acture %d", result)
	}
}
