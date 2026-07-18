package bridge

import (
	greet_types "github.com/nimaeskandary/wails3-react-polylith/components/greet/go/types"
)

// structs in this package embed the interface we want to expose to the frontend. Wails will generate bindings
// for methods on these structs that the frontend can use. We can't use the interfaces directly, wails requires a
// concrete struct to make bindings, so we are embedding our types into these struct.
//
// This project generally uses go.uber.org/fx for dependency injection. We often define interfaces as our domain types
// then when initializing our dependency system, we initialize the concrete struct implementations of those type into our
// dependency graph. Those concrete struct implementations are usually private though, and the interface
// definition can't be used by wails, so this is glue.

type GreetService struct {
	greet_types.GreetService
}
