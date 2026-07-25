//go:build staging

package config

import _ "embed"

//go:embed staging.json
var Bytes []byte
