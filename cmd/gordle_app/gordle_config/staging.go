//go:build staging

package gordle_config

import _ "embed"

//go:embed staging.json
var Bytes []byte
