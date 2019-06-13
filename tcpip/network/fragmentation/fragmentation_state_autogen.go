// automatically generated by stateify.

package fragmentation

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *reassemblerList) beforeSave() {}
func (x *reassemblerList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *reassemblerList) afterLoad() {}
func (x *reassemblerList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *reassemblerEntry) beforeSave() {}
func (x *reassemblerEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *reassemblerEntry) afterLoad() {}
func (x *reassemblerEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func init() {
	state.Register("fragmentation.reassemblerList", (*reassemblerList)(nil), state.Fns{Save: (*reassemblerList).save, Load: (*reassemblerList).load})
	state.Register("fragmentation.reassemblerEntry", (*reassemblerEntry)(nil), state.Fns{Save: (*reassemblerEntry).save, Load: (*reassemblerEntry).load})
}
