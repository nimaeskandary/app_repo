//go:build !production && !staging

package gordle_config

import _ "embed"

//go:embed development.json
var Bytes []byte
