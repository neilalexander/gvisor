// automatically generated by stateify.

package icmp

import (
	"gvisor.googlesource.com/gvisor/pkg/state"
	"gvisor.googlesource.com/gvisor/pkg/tcpip/buffer"
)

func (x *icmpPacket) beforeSave() {}
func (x *icmpPacket) save(m state.Map) {
	x.beforeSave()
	var data buffer.VectorisedView = x.saveData()
	m.SaveValue("data", data)
	m.Save("icmpPacketEntry", &x.icmpPacketEntry)
	m.Save("senderAddress", &x.senderAddress)
	m.Save("timestamp", &x.timestamp)
}

func (x *icmpPacket) afterLoad() {}
func (x *icmpPacket) load(m state.Map) {
	m.Load("icmpPacketEntry", &x.icmpPacketEntry)
	m.Load("senderAddress", &x.senderAddress)
	m.Load("timestamp", &x.timestamp)
	m.LoadValue("data", new(buffer.VectorisedView), func(y interface{}) { x.loadData(y.(buffer.VectorisedView)) })
}

func (x *endpoint) save(m state.Map) {
	x.beforeSave()
	var rcvBufSizeMax int = x.saveRcvBufSizeMax()
	m.SaveValue("rcvBufSizeMax", rcvBufSizeMax)
	m.Save("netProto", &x.netProto)
	m.Save("transProto", &x.transProto)
	m.Save("waiterQueue", &x.waiterQueue)
	m.Save("rcvReady", &x.rcvReady)
	m.Save("rcvList", &x.rcvList)
	m.Save("rcvBufSize", &x.rcvBufSize)
	m.Save("rcvClosed", &x.rcvClosed)
	m.Save("sndBufSize", &x.sndBufSize)
	m.Save("shutdownFlags", &x.shutdownFlags)
	m.Save("id", &x.id)
	m.Save("state", &x.state)
	m.Save("bindNICID", &x.bindNICID)
	m.Save("bindAddr", &x.bindAddr)
	m.Save("regNICID", &x.regNICID)
}

func (x *endpoint) load(m state.Map) {
	m.Load("netProto", &x.netProto)
	m.Load("transProto", &x.transProto)
	m.Load("waiterQueue", &x.waiterQueue)
	m.Load("rcvReady", &x.rcvReady)
	m.Load("rcvList", &x.rcvList)
	m.Load("rcvBufSize", &x.rcvBufSize)
	m.Load("rcvClosed", &x.rcvClosed)
	m.Load("sndBufSize", &x.sndBufSize)
	m.Load("shutdownFlags", &x.shutdownFlags)
	m.Load("id", &x.id)
	m.Load("state", &x.state)
	m.Load("bindNICID", &x.bindNICID)
	m.Load("bindAddr", &x.bindAddr)
	m.Load("regNICID", &x.regNICID)
	m.LoadValue("rcvBufSizeMax", new(int), func(y interface{}) { x.loadRcvBufSizeMax(y.(int)) })
	m.AfterLoad(x.afterLoad)
}

func (x *icmpPacketList) beforeSave() {}
func (x *icmpPacketList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *icmpPacketList) afterLoad() {}
func (x *icmpPacketList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *icmpPacketEntry) beforeSave() {}
func (x *icmpPacketEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *icmpPacketEntry) afterLoad() {}
func (x *icmpPacketEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func init() {
	state.Register("icmp.icmpPacket", (*icmpPacket)(nil), state.Fns{Save: (*icmpPacket).save, Load: (*icmpPacket).load})
	state.Register("icmp.endpoint", (*endpoint)(nil), state.Fns{Save: (*endpoint).save, Load: (*endpoint).load})
	state.Register("icmp.icmpPacketList", (*icmpPacketList)(nil), state.Fns{Save: (*icmpPacketList).save, Load: (*icmpPacketList).load})
	state.Register("icmp.icmpPacketEntry", (*icmpPacketEntry)(nil), state.Fns{Save: (*icmpPacketEntry).save, Load: (*icmpPacketEntry).load})
}
