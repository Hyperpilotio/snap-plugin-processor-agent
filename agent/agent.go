package agent

import (
	"errors"
	"fmt"
	"net/url"
	"path"

	"github.com/go-resty/resty"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

// Processor test processor
type GodddQoSProcessor struct {
	Cache *Cache
}

// NewProcessor generate processor
func NewProcessor() plugin.Processor {
	return &GodddQoSProcessor{
		Cache: NewCache(),
	}
}

func SendMetricToQosDataStore(urlString string, value float64) error {
	u, err := url.Parse(urlString)
	if err != nil {
		return errors.New("Unable to parse url: " + err.Error())
	}

	u.Path = path.Join(u.Path, "v1/apps/goddd/metrics/slack")
	body := make(map[string]float64)
	body["value"] = value
	if response, err := resty.R().SetBody(body).Post(u.String()); err != nil {
		return errors.New("Unable to send metrics to qos data store: " + err.Error())
	} else if response.StatusCode() >= 300 {
		return fmt.Errorf("Unable to send metrics to qos data store: Unexpected response code %d, body: %s", response.StatusCode(), response.String())
	}

	return nil
}

func calculateSlackValue(goal interface{}, current interface{}) float64 {
	switch goal.(type) {
	case float64:
		return (current.(float64) - goal.(float64)) / goal.(float64)
	default:
		return 0.0
	}
}

// Process test process function
func (p *GodddQoSProcessor) Process(mts []plugin.Metric, cfg plugin.Config) ([]plugin.Metric, error) {
	for _, mt := range mts {
		// switch mt.Namespace.Strings()[len(mt.Namespace.Strings())-1] {
		//  case "api_booking_service_request_latency_microseconds":
		// }

		if method, ok := mt.Tags["method"]; ok {
			data, cacheExist := p.Cache.Data[method]
			if !cacheExist {
				p.Cache.Data[method] = CacheType{Pre: mt.Data}
				break
			}
			switch method {
			case "list_cargos":
				// FIXME Instead of giving cache of metric as goal, we should give meaningful benchmark
				_ = calculateSlackValue(data.Pre, mt.Data)
			case "list_locations":
			}
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
func (p *GodddQoSProcessor) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()

	policy.AddNewStringRule([]string{""}, "qos-data-store-url", false, plugin.SetDefaultString("http://qos-data-store:7781"))
	policy.AddNewIntRule([]string{""}, "sla-goal", false, plugin.SetDefaultInt(700))

	return *policy, nil
}
