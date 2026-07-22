//go:build !production && !staging

package config

import _ "embed"

//go:embed development.json
var Bytes []byte
