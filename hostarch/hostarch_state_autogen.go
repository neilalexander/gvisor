// automatically generated by stateify.

package hostarch

import (
	"inet.af/netstack/state"
)

func (a *AccessType) StateTypeName() string {
	return "pkg/hostarch.AccessType"
}

func (a *AccessType) StateFields() []string {
	return []string{
		"Read",
		"Write",
		"Execute",
	}
}

func (a *AccessType) beforeSave() {}

// +checklocksignore
func (a *AccessType) StateSave(stateSinkObject state.Sink) {
	a.beforeSave()
	stateSinkObject.Save(0, &a.Read)
	stateSinkObject.Save(1, &a.Write)
	stateSinkObject.Save(2, &a.Execute)
}

func (a *AccessType) afterLoad() {}

// +checklocksignore
func (a *AccessType) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &a.Read)
	stateSourceObject.Load(1, &a.Write)
	stateSourceObject.Load(2, &a.Execute)
}

func (v *Addr) StateTypeName() string {
	return "pkg/hostarch.Addr"
}

func (v *Addr) StateFields() []string {
	return nil
}

func (r *AddrRange) StateTypeName() string {
	return "pkg/hostarch.AddrRange"
}

func (r *AddrRange) StateFields() []string {
	return []string{
		"Start",
		"End",
	}
}

func (r *AddrRange) beforeSave() {}

// +checklocksignore
func (r *AddrRange) StateSave(stateSinkObject state.Sink) {
	r.beforeSave()
	stateSinkObject.Save(0, &r.Start)
	stateSinkObject.Save(1, &r.End)
}

func (r *AddrRange) afterLoad() {}

// +checklocksignore
func (r *AddrRange) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &r.Start)
	stateSourceObject.Load(1, &r.End)
}

func init() {
	state.Register((*AccessType)(nil))
	state.Register((*Addr)(nil))
	state.Register((*AddrRange)(nil))
}
