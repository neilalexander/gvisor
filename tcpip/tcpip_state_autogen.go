// automatically generated by stateify.

package tcpip

import (
	"gvisor.dev/gvisor/pkg/state"
)

func (x *FullAddress) beforeSave() {}
func (x *FullAddress) save(m state.Map) {
	x.beforeSave()
	m.Save("NIC", &x.NIC)
	m.Save("Addr", &x.Addr)
	m.Save("Port", &x.Port)
}

func (x *FullAddress) afterLoad() {}
func (x *FullAddress) load(m state.Map) {
	m.Load("NIC", &x.NIC)
	m.Load("Addr", &x.Addr)
	m.Load("Port", &x.Port)
}

func (x *ControlMessages) beforeSave() {}
func (x *ControlMessages) save(m state.Map) {
	x.beforeSave()
	m.Save("HasTimestamp", &x.HasTimestamp)
	m.Save("Timestamp", &x.Timestamp)
}

func (x *ControlMessages) afterLoad() {}
func (x *ControlMessages) load(m state.Map) {
	m.Load("HasTimestamp", &x.HasTimestamp)
	m.Load("Timestamp", &x.Timestamp)
}

func init() {
	state.Register("tcpip.FullAddress", (*FullAddress)(nil), state.Fns{Save: (*FullAddress).save, Load: (*FullAddress).load})
	state.Register("tcpip.ControlMessages", (*ControlMessages)(nil), state.Fns{Save: (*ControlMessages).save, Load: (*ControlMessages).load})
}
