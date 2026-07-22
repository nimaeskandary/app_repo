//go:build production && !staging

package config

import _ "embed"

//go:embed production.json
var Bytes []byte
