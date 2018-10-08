package config

import (
    "testing"
)

func TestYAML(t *testing.T) {
	var data = `
A: Easy!
B:
    A: 2
    B: [3, 4]
    C: 
        D: "value1"
    E:
    -   E1: value_e1
    -   E2: value_e2
`
    configs = loadConfigs(data)
    t.Log(GetStringOrEmpty(`B.E`))
}
