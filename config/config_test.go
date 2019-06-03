package config

import "testing"

func TestRead(t *testing.T) {
	BindServerConfig("server.json")
}
