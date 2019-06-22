// automatically generated by stateify.

package tcp

import (
	"gvisor.dev/gvisor/pkg/state"
	"gvisor.dev/gvisor/pkg/tcpip/buffer"
)

func (x *cubicState) beforeSave() {}
func (x *cubicState) save(m state.Map) {
	x.beforeSave()
	var t unixTime = x.saveT()
	m.SaveValue("t", t)
	m.Save("wLastMax", &x.wLastMax)
	m.Save("wMax", &x.wMax)
	m.Save("numCongestionEvents", &x.numCongestionEvents)
	m.Save("c", &x.c)
	m.Save("k", &x.k)
	m.Save("beta", &x.beta)
	m.Save("wC", &x.wC)
	m.Save("wEst", &x.wEst)
	m.Save("s", &x.s)
}

func (x *cubicState) afterLoad() {}
func (x *cubicState) load(m state.Map) {
	m.Load("wLastMax", &x.wLastMax)
	m.Load("wMax", &x.wMax)
	m.Load("numCongestionEvents", &x.numCongestionEvents)
	m.Load("c", &x.c)
	m.Load("k", &x.k)
	m.Load("beta", &x.beta)
	m.Load("wC", &x.wC)
	m.Load("wEst", &x.wEst)
	m.Load("s", &x.s)
	m.LoadValue("t", new(unixTime), func(y interface{}) { x.loadT(y.(unixTime)) })
}

func (x *SACKInfo) beforeSave() {}
func (x *SACKInfo) save(m state.Map) {
	x.beforeSave()
	m.Save("Blocks", &x.Blocks)
	m.Save("NumBlocks", &x.NumBlocks)
}

func (x *SACKInfo) afterLoad() {}
func (x *SACKInfo) load(m state.Map) {
	m.Load("Blocks", &x.Blocks)
	m.Load("NumBlocks", &x.NumBlocks)
}

func (x *rcvBufAutoTuneParams) beforeSave() {}
func (x *rcvBufAutoTuneParams) save(m state.Map) {
	x.beforeSave()
	var measureTime unixTime = x.saveMeasureTime()
	m.SaveValue("measureTime", measureTime)
	var rttMeasureTime unixTime = x.saveRttMeasureTime()
	m.SaveValue("rttMeasureTime", rttMeasureTime)
	m.Save("copied", &x.copied)
	m.Save("prevCopied", &x.prevCopied)
	m.Save("rtt", &x.rtt)
	m.Save("rttMeasureSeqNumber", &x.rttMeasureSeqNumber)
	m.Save("disabled", &x.disabled)
}

func (x *rcvBufAutoTuneParams) afterLoad() {}
func (x *rcvBufAutoTuneParams) load(m state.Map) {
	m.Load("copied", &x.copied)
	m.Load("prevCopied", &x.prevCopied)
	m.Load("rtt", &x.rtt)
	m.Load("rttMeasureSeqNumber", &x.rttMeasureSeqNumber)
	m.Load("disabled", &x.disabled)
	m.LoadValue("measureTime", new(unixTime), func(y interface{}) { x.loadMeasureTime(y.(unixTime)) })
	m.LoadValue("rttMeasureTime", new(unixTime), func(y interface{}) { x.loadRttMeasureTime(y.(unixTime)) })
}

func (x *endpoint) save(m state.Map) {
	x.beforeSave()
	var lastError string = x.saveLastError()
	m.SaveValue("lastError", lastError)
	var state EndpointState = x.saveState()
	m.SaveValue("state", state)
	var hardError string = x.saveHardError()
	m.SaveValue("hardError", hardError)
	var acceptedChan []*endpoint = x.saveAcceptedChan()
	m.SaveValue("acceptedChan", acceptedChan)
	m.Save("netProto", &x.netProto)
	m.Save("waiterQueue", &x.waiterQueue)
	m.Save("rcvList", &x.rcvList)
	m.Save("rcvClosed", &x.rcvClosed)
	m.Save("rcvBufSize", &x.rcvBufSize)
	m.Save("rcvBufUsed", &x.rcvBufUsed)
	m.Save("rcvAutoParams", &x.rcvAutoParams)
	m.Save("zeroWindow", &x.zeroWindow)
	m.Save("id", &x.id)
	m.Save("isRegistered", &x.isRegistered)
	m.Save("v6only", &x.v6only)
	m.Save("isConnectNotified", &x.isConnectNotified)
	m.Save("broadcast", &x.broadcast)
	m.Save("workerRunning", &x.workerRunning)
	m.Save("workerCleanup", &x.workerCleanup)
	m.Save("sendTSOk", &x.sendTSOk)
	m.Save("recentTS", &x.recentTS)
	m.Save("tsOffset", &x.tsOffset)
	m.Save("shutdownFlags", &x.shutdownFlags)
	m.Save("sackPermitted", &x.sackPermitted)
	m.Save("sack", &x.sack)
	m.Save("reusePort", &x.reusePort)
	m.Save("delay", &x.delay)
	m.Save("cork", &x.cork)
	m.Save("scoreboard", &x.scoreboard)
	m.Save("reuseAddr", &x.reuseAddr)
	m.Save("slowAck", &x.slowAck)
	m.Save("segmentQueue", &x.segmentQueue)
	m.Save("synRcvdCount", &x.synRcvdCount)
	m.Save("sndBufSize", &x.sndBufSize)
	m.Save("sndBufUsed", &x.sndBufUsed)
	m.Save("sndClosed", &x.sndClosed)
	m.Save("sndBufInQueue", &x.sndBufInQueue)
	m.Save("sndQueue", &x.sndQueue)
	m.Save("cc", &x.cc)
	m.Save("packetTooBigCount", &x.packetTooBigCount)
	m.Save("sndMTU", &x.sndMTU)
	m.Save("keepalive", &x.keepalive)
	m.Save("rcv", &x.rcv)
	m.Save("snd", &x.snd)
	m.Save("bindAddress", &x.bindAddress)
	m.Save("connectingAddress", &x.connectingAddress)
	m.Save("amss", &x.amss)
	m.Save("gso", &x.gso)
}

func (x *endpoint) load(m state.Map) {
	m.Load("netProto", &x.netProto)
	m.LoadWait("waiterQueue", &x.waiterQueue)
	m.LoadWait("rcvList", &x.rcvList)
	m.Load("rcvClosed", &x.rcvClosed)
	m.Load("rcvBufSize", &x.rcvBufSize)
	m.Load("rcvBufUsed", &x.rcvBufUsed)
	m.Load("rcvAutoParams", &x.rcvAutoParams)
	m.Load("zeroWindow", &x.zeroWindow)
	m.Load("id", &x.id)
	m.Load("isRegistered", &x.isRegistered)
	m.Load("v6only", &x.v6only)
	m.Load("isConnectNotified", &x.isConnectNotified)
	m.Load("broadcast", &x.broadcast)
	m.Load("workerRunning", &x.workerRunning)
	m.Load("workerCleanup", &x.workerCleanup)
	m.Load("sendTSOk", &x.sendTSOk)
	m.Load("recentTS", &x.recentTS)
	m.Load("tsOffset", &x.tsOffset)
	m.Load("shutdownFlags", &x.shutdownFlags)
	m.Load("sackPermitted", &x.sackPermitted)
	m.Load("sack", &x.sack)
	m.Load("reusePort", &x.reusePort)
	m.Load("delay", &x.delay)
	m.Load("cork", &x.cork)
	m.Load("scoreboard", &x.scoreboard)
	m.Load("reuseAddr", &x.reuseAddr)
	m.Load("slowAck", &x.slowAck)
	m.LoadWait("segmentQueue", &x.segmentQueue)
	m.Load("synRcvdCount", &x.synRcvdCount)
	m.Load("sndBufSize", &x.sndBufSize)
	m.Load("sndBufUsed", &x.sndBufUsed)
	m.Load("sndClosed", &x.sndClosed)
	m.Load("sndBufInQueue", &x.sndBufInQueue)
	m.LoadWait("sndQueue", &x.sndQueue)
	m.Load("cc", &x.cc)
	m.Load("packetTooBigCount", &x.packetTooBigCount)
	m.Load("sndMTU", &x.sndMTU)
	m.Load("keepalive", &x.keepalive)
	m.LoadWait("rcv", &x.rcv)
	m.LoadWait("snd", &x.snd)
	m.Load("bindAddress", &x.bindAddress)
	m.Load("connectingAddress", &x.connectingAddress)
	m.Load("amss", &x.amss)
	m.Load("gso", &x.gso)
	m.LoadValue("lastError", new(string), func(y interface{}) { x.loadLastError(y.(string)) })
	m.LoadValue("state", new(EndpointState), func(y interface{}) { x.loadState(y.(EndpointState)) })
	m.LoadValue("hardError", new(string), func(y interface{}) { x.loadHardError(y.(string)) })
	m.LoadValue("acceptedChan", new([]*endpoint), func(y interface{}) { x.loadAcceptedChan(y.([]*endpoint)) })
	m.AfterLoad(x.afterLoad)
}

func (x *keepalive) beforeSave() {}
func (x *keepalive) save(m state.Map) {
	x.beforeSave()
	m.Save("enabled", &x.enabled)
	m.Save("idle", &x.idle)
	m.Save("interval", &x.interval)
	m.Save("count", &x.count)
	m.Save("unacked", &x.unacked)
}

func (x *keepalive) afterLoad() {}
func (x *keepalive) load(m state.Map) {
	m.Load("enabled", &x.enabled)
	m.Load("idle", &x.idle)
	m.Load("interval", &x.interval)
	m.Load("count", &x.count)
	m.Load("unacked", &x.unacked)
}

func (x *receiver) beforeSave() {}
func (x *receiver) save(m state.Map) {
	x.beforeSave()
	m.Save("ep", &x.ep)
	m.Save("rcvNxt", &x.rcvNxt)
	m.Save("rcvAcc", &x.rcvAcc)
	m.Save("rcvWnd", &x.rcvWnd)
	m.Save("rcvWndScale", &x.rcvWndScale)
	m.Save("closed", &x.closed)
	m.Save("pendingRcvdSegments", &x.pendingRcvdSegments)
	m.Save("pendingBufUsed", &x.pendingBufUsed)
	m.Save("pendingBufSize", &x.pendingBufSize)
}

func (x *receiver) afterLoad() {}
func (x *receiver) load(m state.Map) {
	m.Load("ep", &x.ep)
	m.Load("rcvNxt", &x.rcvNxt)
	m.Load("rcvAcc", &x.rcvAcc)
	m.Load("rcvWnd", &x.rcvWnd)
	m.Load("rcvWndScale", &x.rcvWndScale)
	m.Load("closed", &x.closed)
	m.Load("pendingRcvdSegments", &x.pendingRcvdSegments)
	m.Load("pendingBufUsed", &x.pendingBufUsed)
	m.Load("pendingBufSize", &x.pendingBufSize)
}

func (x *renoState) beforeSave() {}
func (x *renoState) save(m state.Map) {
	x.beforeSave()
	m.Save("s", &x.s)
}

func (x *renoState) afterLoad() {}
func (x *renoState) load(m state.Map) {
	m.Load("s", &x.s)
}

func (x *SACKScoreboard) beforeSave() {}
func (x *SACKScoreboard) save(m state.Map) {
	x.beforeSave()
	m.Save("smss", &x.smss)
	m.Save("maxSACKED", &x.maxSACKED)
}

func (x *SACKScoreboard) afterLoad() {}
func (x *SACKScoreboard) load(m state.Map) {
	m.Load("smss", &x.smss)
	m.Load("maxSACKED", &x.maxSACKED)
}

func (x *segment) beforeSave() {}
func (x *segment) save(m state.Map) {
	x.beforeSave()
	var data buffer.VectorisedView = x.saveData()
	m.SaveValue("data", data)
	var options []byte = x.saveOptions()
	m.SaveValue("options", options)
	var rcvdTime unixTime = x.saveRcvdTime()
	m.SaveValue("rcvdTime", rcvdTime)
	var xmitTime unixTime = x.saveXmitTime()
	m.SaveValue("xmitTime", xmitTime)
	m.Save("segmentEntry", &x.segmentEntry)
	m.Save("refCnt", &x.refCnt)
	m.Save("viewToDeliver", &x.viewToDeliver)
	m.Save("sequenceNumber", &x.sequenceNumber)
	m.Save("ackNumber", &x.ackNumber)
	m.Save("flags", &x.flags)
	m.Save("window", &x.window)
	m.Save("csum", &x.csum)
	m.Save("csumValid", &x.csumValid)
	m.Save("parsedOptions", &x.parsedOptions)
	m.Save("hasNewSACKInfo", &x.hasNewSACKInfo)
}

func (x *segment) afterLoad() {}
func (x *segment) load(m state.Map) {
	m.Load("segmentEntry", &x.segmentEntry)
	m.Load("refCnt", &x.refCnt)
	m.Load("viewToDeliver", &x.viewToDeliver)
	m.Load("sequenceNumber", &x.sequenceNumber)
	m.Load("ackNumber", &x.ackNumber)
	m.Load("flags", &x.flags)
	m.Load("window", &x.window)
	m.Load("csum", &x.csum)
	m.Load("csumValid", &x.csumValid)
	m.Load("parsedOptions", &x.parsedOptions)
	m.Load("hasNewSACKInfo", &x.hasNewSACKInfo)
	m.LoadValue("data", new(buffer.VectorisedView), func(y interface{}) { x.loadData(y.(buffer.VectorisedView)) })
	m.LoadValue("options", new([]byte), func(y interface{}) { x.loadOptions(y.([]byte)) })
	m.LoadValue("rcvdTime", new(unixTime), func(y interface{}) { x.loadRcvdTime(y.(unixTime)) })
	m.LoadValue("xmitTime", new(unixTime), func(y interface{}) { x.loadXmitTime(y.(unixTime)) })
}

func (x *segmentQueue) beforeSave() {}
func (x *segmentQueue) save(m state.Map) {
	x.beforeSave()
	m.Save("list", &x.list)
	m.Save("limit", &x.limit)
	m.Save("used", &x.used)
}

func (x *segmentQueue) afterLoad() {}
func (x *segmentQueue) load(m state.Map) {
	m.LoadWait("list", &x.list)
	m.Load("limit", &x.limit)
	m.Load("used", &x.used)
}

func (x *sender) beforeSave() {}
func (x *sender) save(m state.Map) {
	x.beforeSave()
	var lastSendTime unixTime = x.saveLastSendTime()
	m.SaveValue("lastSendTime", lastSendTime)
	var rttMeasureTime unixTime = x.saveRttMeasureTime()
	m.SaveValue("rttMeasureTime", rttMeasureTime)
	m.Save("ep", &x.ep)
	m.Save("dupAckCount", &x.dupAckCount)
	m.Save("fr", &x.fr)
	m.Save("sndCwnd", &x.sndCwnd)
	m.Save("sndSsthresh", &x.sndSsthresh)
	m.Save("sndCAAckCount", &x.sndCAAckCount)
	m.Save("outstanding", &x.outstanding)
	m.Save("sndWnd", &x.sndWnd)
	m.Save("sndUna", &x.sndUna)
	m.Save("sndNxt", &x.sndNxt)
	m.Save("sndNxtList", &x.sndNxtList)
	m.Save("rttMeasureSeqNum", &x.rttMeasureSeqNum)
	m.Save("closed", &x.closed)
	m.Save("writeNext", &x.writeNext)
	m.Save("writeList", &x.writeList)
	m.Save("rtt", &x.rtt)
	m.Save("rto", &x.rto)
	m.Save("maxPayloadSize", &x.maxPayloadSize)
	m.Save("gso", &x.gso)
	m.Save("sndWndScale", &x.sndWndScale)
	m.Save("maxSentAck", &x.maxSentAck)
	m.Save("cc", &x.cc)
}

func (x *sender) load(m state.Map) {
	m.Load("ep", &x.ep)
	m.Load("dupAckCount", &x.dupAckCount)
	m.Load("fr", &x.fr)
	m.Load("sndCwnd", &x.sndCwnd)
	m.Load("sndSsthresh", &x.sndSsthresh)
	m.Load("sndCAAckCount", &x.sndCAAckCount)
	m.Load("outstanding", &x.outstanding)
	m.Load("sndWnd", &x.sndWnd)
	m.Load("sndUna", &x.sndUna)
	m.Load("sndNxt", &x.sndNxt)
	m.Load("sndNxtList", &x.sndNxtList)
	m.Load("rttMeasureSeqNum", &x.rttMeasureSeqNum)
	m.Load("closed", &x.closed)
	m.Load("writeNext", &x.writeNext)
	m.Load("writeList", &x.writeList)
	m.Load("rtt", &x.rtt)
	m.Load("rto", &x.rto)
	m.Load("maxPayloadSize", &x.maxPayloadSize)
	m.Load("gso", &x.gso)
	m.Load("sndWndScale", &x.sndWndScale)
	m.Load("maxSentAck", &x.maxSentAck)
	m.Load("cc", &x.cc)
	m.LoadValue("lastSendTime", new(unixTime), func(y interface{}) { x.loadLastSendTime(y.(unixTime)) })
	m.LoadValue("rttMeasureTime", new(unixTime), func(y interface{}) { x.loadRttMeasureTime(y.(unixTime)) })
	m.AfterLoad(x.afterLoad)
}

func (x *rtt) beforeSave() {}
func (x *rtt) save(m state.Map) {
	x.beforeSave()
	m.Save("srtt", &x.srtt)
	m.Save("rttvar", &x.rttvar)
	m.Save("srttInited", &x.srttInited)
}

func (x *rtt) afterLoad() {}
func (x *rtt) load(m state.Map) {
	m.Load("srtt", &x.srtt)
	m.Load("rttvar", &x.rttvar)
	m.Load("srttInited", &x.srttInited)
}

func (x *fastRecovery) beforeSave() {}
func (x *fastRecovery) save(m state.Map) {
	x.beforeSave()
	m.Save("active", &x.active)
	m.Save("first", &x.first)
	m.Save("last", &x.last)
	m.Save("maxCwnd", &x.maxCwnd)
	m.Save("highRxt", &x.highRxt)
	m.Save("rescueRxt", &x.rescueRxt)
}

func (x *fastRecovery) afterLoad() {}
func (x *fastRecovery) load(m state.Map) {
	m.Load("active", &x.active)
	m.Load("first", &x.first)
	m.Load("last", &x.last)
	m.Load("maxCwnd", &x.maxCwnd)
	m.Load("highRxt", &x.highRxt)
	m.Load("rescueRxt", &x.rescueRxt)
}

func (x *unixTime) beforeSave() {}
func (x *unixTime) save(m state.Map) {
	x.beforeSave()
	m.Save("second", &x.second)
	m.Save("nano", &x.nano)
}

func (x *unixTime) afterLoad() {}
func (x *unixTime) load(m state.Map) {
	m.Load("second", &x.second)
	m.Load("nano", &x.nano)
}

func (x *segmentList) beforeSave() {}
func (x *segmentList) save(m state.Map) {
	x.beforeSave()
	m.Save("head", &x.head)
	m.Save("tail", &x.tail)
}

func (x *segmentList) afterLoad() {}
func (x *segmentList) load(m state.Map) {
	m.Load("head", &x.head)
	m.Load("tail", &x.tail)
}

func (x *segmentEntry) beforeSave() {}
func (x *segmentEntry) save(m state.Map) {
	x.beforeSave()
	m.Save("next", &x.next)
	m.Save("prev", &x.prev)
}

func (x *segmentEntry) afterLoad() {}
func (x *segmentEntry) load(m state.Map) {
	m.Load("next", &x.next)
	m.Load("prev", &x.prev)
}

func init() {
	state.Register("tcp.cubicState", (*cubicState)(nil), state.Fns{Save: (*cubicState).save, Load: (*cubicState).load})
	state.Register("tcp.SACKInfo", (*SACKInfo)(nil), state.Fns{Save: (*SACKInfo).save, Load: (*SACKInfo).load})
	state.Register("tcp.rcvBufAutoTuneParams", (*rcvBufAutoTuneParams)(nil), state.Fns{Save: (*rcvBufAutoTuneParams).save, Load: (*rcvBufAutoTuneParams).load})
	state.Register("tcp.endpoint", (*endpoint)(nil), state.Fns{Save: (*endpoint).save, Load: (*endpoint).load})
	state.Register("tcp.keepalive", (*keepalive)(nil), state.Fns{Save: (*keepalive).save, Load: (*keepalive).load})
	state.Register("tcp.receiver", (*receiver)(nil), state.Fns{Save: (*receiver).save, Load: (*receiver).load})
	state.Register("tcp.renoState", (*renoState)(nil), state.Fns{Save: (*renoState).save, Load: (*renoState).load})
	state.Register("tcp.SACKScoreboard", (*SACKScoreboard)(nil), state.Fns{Save: (*SACKScoreboard).save, Load: (*SACKScoreboard).load})
	state.Register("tcp.segment", (*segment)(nil), state.Fns{Save: (*segment).save, Load: (*segment).load})
	state.Register("tcp.segmentQueue", (*segmentQueue)(nil), state.Fns{Save: (*segmentQueue).save, Load: (*segmentQueue).load})
	state.Register("tcp.sender", (*sender)(nil), state.Fns{Save: (*sender).save, Load: (*sender).load})
	state.Register("tcp.rtt", (*rtt)(nil), state.Fns{Save: (*rtt).save, Load: (*rtt).load})
	state.Register("tcp.fastRecovery", (*fastRecovery)(nil), state.Fns{Save: (*fastRecovery).save, Load: (*fastRecovery).load})
	state.Register("tcp.unixTime", (*unixTime)(nil), state.Fns{Save: (*unixTime).save, Load: (*unixTime).load})
	state.Register("tcp.segmentList", (*segmentList)(nil), state.Fns{Save: (*segmentList).save, Load: (*segmentList).load})
	state.Register("tcp.segmentEntry", (*segmentEntry)(nil), state.Fns{Save: (*segmentEntry).save, Load: (*segmentEntry).load})
}
