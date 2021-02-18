// +build !linux

package kernel

import (
	"github.com/pmalek-sumo/telegraf"
	"github.com/pmalek-sumo/telegraf/plugins/inputs"
)

type Kernel struct {
}

func (k *Kernel) Description() string {
	return "Get kernel statistics from /proc/stat"
}

func (k *Kernel) SampleConfig() string { return "" }

func (k *Kernel) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("kernel", func() telegraf.Input {
		return &Kernel{}
	})
}
