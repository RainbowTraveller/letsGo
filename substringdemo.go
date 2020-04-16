package main

import (
	"fmt"
	"strings"
)

func main() {
	longstr := "test_data/indentationError.yaml:7:0:error:yaml: unmarshal errors: field nodes not found in type entities.ContextStateMapping,Key: 'ContextState.Mappings[0].ContextStateNodeSets' Error:Field validation for 'ContextStateNodeSets' failed on the 'required' tag,Key: 'ContextState.Mappings[1].AppName' Error:Field validation for 'AppName' failed on the 'required' tag,Key: 'ContextState.Mappings[1].Os' Error:Field validation for 'Os' failed on the 'required' tag,Key: 'ContextState.Mappings[1].ContextStateNodeSets' Error:Field validation for 'ContextStateNodeSets' failed on the 'required' tag"
	shortstr := "Key: 'ContextState.Mappings[0].ContextStateNodeSets' Error:Field validation for 'ContextStateNodeSets' failed on the 'required' tag"

	fmt.Println(strings.Contains(longstr, shortstr))
	fmt.Println(strings.Contains(longstr, "Key: 'ContextState.Mappings[1].AppName' Error:Field validation for 'AppName' failed on the 'required' tag"))
	fmt.Println(strings.Contains(longstr, "Key: 'ContextState.Mappings[1].Os' Error:Field validation for 'Os' failed on the 'required' tag"))
	fmt.Println(strings.Contains(longstr, "Key: 'ContextState.Mappings[1].ContextStateNodeSets' Error:Field validation for 'ContextStateNodeSets' failed on the 'required' tag"))

}
