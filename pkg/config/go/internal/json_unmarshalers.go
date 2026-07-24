package internal

import (
	"encoding/json/v2"
	"sync"
)

// jsonUnmarshalers is used to extend the a json v2 parser with addition logic for custom types. E.g.
// if you have a config loader that you need to provide a specific transformation to a type, you can
// call RegisterJSONUnmarshaler in the init() function of where you declare that config value type
var jsonUnmarshalers = struct {
	sync.RWMutex
	values []*json.Unmarshalers
}{}

// RegisterJSONUnmarshaler adds a custom type unmarshaler to the config loader.
func RegisterJSONUnmarshaler(unmarshaler *json.Unmarshalers) {
	jsonUnmarshalers.Lock()
	defer jsonUnmarshalers.Unlock()

	jsonUnmarshalers.values = append(jsonUnmarshalers.values, unmarshaler)
}

// registeredJSONUnmarshalers returns a stable snapshot of custom unmarshalers.
func registeredJSONUnmarshalers() *json.Unmarshalers {
	jsonUnmarshalers.RLock()
	defer jsonUnmarshalers.RUnlock()

	return json.JoinUnmarshalers(jsonUnmarshalers.values...)
}
