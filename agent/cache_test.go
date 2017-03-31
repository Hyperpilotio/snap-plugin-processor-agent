package agent

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type MockMetricsDownloader struct {
}

func TestAgentPlugin(t *testing.T) {
	Convey("Create a QoS Processor", t, func() {
		p := &GodddQoSProcessor{
			Cache: NewCache(),
		}

		So(p.Cache, ShouldNotBeNil)
		Convey("Should be able to keep data", func() {
			key := "TEST_CACHE"
			value := 100.0
			p.Cache.Data[key] = CacheType{Pre: value}
			So(p.Cache.Data[key].Pre, ShouldEqual, 100.0)
		})
	})
}
