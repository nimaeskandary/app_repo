# tools

## Go tools

This directory can be used so so that go tools installed via go get -tool, e.g. `go get -tool github.com/golangci/golangci-lint/cmd/golangci-lint@latest`, do not pollute dependencies in other modules. Often installing go tools can be finkicky dependncy wise, because all the dependencies of the tool can conflict with your code bases actual dependencies.

### adding a new go tool

* mkdir `tools/some_tool && cd tools/some_tool`
* create a simple `go.mod`, e.g.

```
module github.com/nimaeskandary/app_repo/tools/some_tool

go 1.26.5
```

* `go get -tool some_tool@version|latest`

This will give this go tool a unique `go.mod` to track its own dependencies, and when you upgrade its deps or your main apps deps it won't cause conflicts

### running a tool

You can run one of these tools from the root directory with the -modfile flag, e.g. `go tool -modfile=./tools/some_tool/go.mod golangci-lint run`
