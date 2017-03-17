package agent

import (
	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

// RProcessor test processor
type Processor struct {
}

// Process test process function
func (p Processor) Process(mts []plugin.Metric, cfg plugin.Config) ([]plugin.Metric, error) {
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
