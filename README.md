# tulipindicators-go

An attempt to wrap
[tulipindicators](https://github.com/TulipCharts/tulipindicators)
in go.  Each indicator is wrapped in its auto-generated source file
under `indicators/`, with matching inputs, options, and outputs as
its underlying C function.

As go does not support `go get` with `git submodule`, the repository
include a checked out version `tulipindicators` source files.
The included files are from commit
[cffa15e](https://github.com/TulipCharts/tulipindicators/commit/cffa15e389b7a0c472588f22fa326a78bd734391).