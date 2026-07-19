# tools

## Go tools

This module is so that go tools installed via go get -tool, e.g. `go get -tool github.com/golangci/golangci-lint/cmd/golangci-lint@latest`, do not pollute dependencies in other modules. Often installing go tools can be finkicky dependncy wise, and then gp module updates can break things. This allos each go tool to have its own deps.

* cd to `<root>/<my_tool>/`
* `go get -tool my_tool@version|latest`

## running a tool

You can run one of these tools from the root directory with the -modfile flag, e.g. `go tool -modfile=./tools/my_tool/go.mod golangci-lint run`
