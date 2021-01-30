package configs

import (
	"reflect"
	"testing"
)

func TestConfigIsSingleton(t *testing.T) {
	cfg1 := New()
	cfg2 := New()

	if !reflect.DeepEqual(cfg1, cfg2) {
		t.Errorf("Config file isn't singleton.")
	}
}
