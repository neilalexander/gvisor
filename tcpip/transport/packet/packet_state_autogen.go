// automatically generated by stateify.

package packet

import (
	"gvisor.dev/gvisor/pkg/state"
	"gvisor.dev/gvisor/pkg/tcpip/buffer"
)

func (p *packet) StateTypeName() string {
	return "pkg/tcpip/transport/packet.packet"
}

func (p *packet) StateFields() []string {
	return []string{
		"packetEntry",
		"data",
		"timestampNS",
		"senderAddr",
		"packetInfo",
	}
}

func (p *packet) beforeSave() {}

func (p *packet) StateSave(stateSinkObject state.Sink) {
	p.beforeSave()
	var dataValue buffer.VectorisedView = p.saveData()
	stateSinkObject.SaveValue(1, dataValue)
	stateSinkObject.Save(0, &p.packetEntry)
	stateSinkObject.Save(2, &p.timestampNS)
	stateSinkObject.Save(3, &p.senderAddr)
	stateSinkObject.Save(4, &p.packetInfo)
}

func (p *packet) afterLoad() {}

func (p *packet) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &p.packetEntry)
	stateSourceObject.Load(2, &p.timestampNS)
	stateSourceObject.Load(3, &p.senderAddr)
	stateSourceObject.Load(4, &p.packetInfo)
	stateSourceObject.LoadValue(1, new(buffer.VectorisedView), func(y interface{}) { p.loadData(y.(buffer.VectorisedView)) })
}

func (ep *endpoint) StateTypeName() string {
	return "pkg/tcpip/transport/packet.endpoint"
}

func (ep *endpoint) StateFields() []string {
	return []string{
		"TransportEndpointInfo",
		"netProto",
		"waiterQueue",
		"cooked",
		"rcvList",
		"rcvBufSizeMax",
		"rcvBufSize",
		"rcvClosed",
		"sndBufSize",
		"sndBufSizeMax",
		"closed",
		"bound",
		"boundNIC",
		"linger",
		"lastError",
	}
}

func (ep *endpoint) StateSave(stateSinkObject state.Sink) {
	ep.beforeSave()
	var rcvBufSizeMaxValue int = ep.saveRcvBufSizeMax()
	stateSinkObject.SaveValue(5, rcvBufSizeMaxValue)
	var lastErrorValue string = ep.saveLastError()
	stateSinkObject.SaveValue(14, lastErrorValue)
	stateSinkObject.Save(0, &ep.TransportEndpointInfo)
	stateSinkObject.Save(1, &ep.netProto)
	stateSinkObject.Save(2, &ep.waiterQueue)
	stateSinkObject.Save(3, &ep.cooked)
	stateSinkObject.Save(4, &ep.rcvList)
	stateSinkObject.Save(6, &ep.rcvBufSize)
	stateSinkObject.Save(7, &ep.rcvClosed)
	stateSinkObject.Save(8, &ep.sndBufSize)
	stateSinkObject.Save(9, &ep.sndBufSizeMax)
	stateSinkObject.Save(10, &ep.closed)
	stateSinkObject.Save(11, &ep.bound)
	stateSinkObject.Save(12, &ep.boundNIC)
	stateSinkObject.Save(13, &ep.linger)
}

func (ep *endpoint) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &ep.TransportEndpointInfo)
	stateSourceObject.Load(1, &ep.netProto)
	stateSourceObject.Load(2, &ep.waiterQueue)
	stateSourceObject.Load(3, &ep.cooked)
	stateSourceObject.Load(4, &ep.rcvList)
	stateSourceObject.Load(6, &ep.rcvBufSize)
	stateSourceObject.Load(7, &ep.rcvClosed)
	stateSourceObject.Load(8, &ep.sndBufSize)
	stateSourceObject.Load(9, &ep.sndBufSizeMax)
	stateSourceObject.Load(10, &ep.closed)
	stateSourceObject.Load(11, &ep.bound)
	stateSourceObject.Load(12, &ep.boundNIC)
	stateSourceObject.Load(13, &ep.linger)
	stateSourceObject.LoadValue(5, new(int), func(y interface{}) { ep.loadRcvBufSizeMax(y.(int)) })
	stateSourceObject.LoadValue(14, new(string), func(y interface{}) { ep.loadLastError(y.(string)) })
	stateSourceObject.AfterLoad(ep.afterLoad)
}

func (l *packetList) StateTypeName() string {
	return "pkg/tcpip/transport/packet.packetList"
}

func (l *packetList) StateFields() []string {
	return []string{
		"head",
		"tail",
	}
}

func (l *packetList) beforeSave() {}

func (l *packetList) StateSave(stateSinkObject state.Sink) {
	l.beforeSave()
	stateSinkObject.Save(0, &l.head)
	stateSinkObject.Save(1, &l.tail)
}

func (l *packetList) afterLoad() {}

func (l *packetList) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &l.head)
	stateSourceObject.Load(1, &l.tail)
}

func (e *packetEntry) StateTypeName() string {
	return "pkg/tcpip/transport/packet.packetEntry"
}

func (e *packetEntry) StateFields() []string {
	return []string{
		"next",
		"prev",
	}
}

func (e *packetEntry) beforeSave() {}

func (e *packetEntry) StateSave(stateSinkObject state.Sink) {
	e.beforeSave()
	stateSinkObject.Save(0, &e.next)
	stateSinkObject.Save(1, &e.prev)
}

func (e *packetEntry) afterLoad() {}

func (e *packetEntry) StateLoad(stateSourceObject state.Source) {
	stateSourceObject.Load(0, &e.next)
	stateSourceObject.Load(1, &e.prev)
}

func init() {
	state.Register((*packet)(nil))
	state.Register((*endpoint)(nil))
	state.Register((*packetList)(nil))
	state.Register((*packetEntry)(nil))
}
