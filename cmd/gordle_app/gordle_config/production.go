//go:build production && !staging

package gordle_config

import _ "embed"

//go:embed production.json
var Bytes []byte
