package agent

import (
	"fmt"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

// Processor test processor
type Processor struct {
}

// Process test process function
func (p Processor) Process(mts []plugin.Metric, cfg plugin.Config) ([]plugin.Metric, error) {
	for _, mt := range mts {
		switch mt.Namespace.Strings()[len(mt.Namespace.Strings())-1] {

		case "http_requests_total":
			// do something
			fmt.Printf("http_requests_total is %v\n", mt.Data)
		}
	}

	return mts, nil
}

/*
	GetConfigPolicy() returns the configPolicy for your plugin.

	A config policy is how users can provide configuration info to
	plugin. Here you define what sorts of config info your plugin
	needs and/or requires.
*/
func (p Processor) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()

	//policy.AddNewBoolRule([]string{"random"},"testbool",false)

	return *policy, nil
}
