package config

import "testing"

func TestGet(t *testing.T) {

	t.Run("basic", func(t *testing.T) {
		got, _ := Get()
		t.Logf("Got: %s", got)
	})
}
