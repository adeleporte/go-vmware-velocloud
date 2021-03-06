/*
 * VeloCloud Orchestrator API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 3.3.2
 * Contact: support@velocloud.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LiveModeDataLinkStatsDataData struct {
	BackupOnly int32 `json:"backupOnly,omitempty"`
	BestJitterMsRx int32 `json:"bestJitterMsRx,omitempty"`
	BestJitterMsTx int32 `json:"bestJitterMsTx,omitempty"`
	BestLatencyMsRx int32 `json:"bestLatencyMsRx,omitempty"`
	BestLatencyMsTx int32 `json:"bestLatencyMsTx,omitempty"`
	BestLossPctRx float32 `json:"bestLossPctRx,omitempty"`
	BestLossPctTx float32 `json:"bestLossPctTx,omitempty"`
	BpsOfBestPathRx int32 `json:"bpsOfBestPathRx,omitempty"`
	BpsOfBestPathTx int32 `json:"bpsOfBestPathTx,omitempty"`
	BytesRx int64 `json:"bytesRx,omitempty"`
	BytesTx int64 `json:"bytesTx,omitempty"`
	ControlBytesRx int64 `json:"controlBytesRx,omitempty"`
	ControlBytesTx int64 `json:"controlBytesTx,omitempty"`
	ControlPacketsRx int64 `json:"controlPacketsRx,omitempty"`
	ControlPacketsTx int64 `json:"controlPacketsTx,omitempty"`
	Destinations []Destinations `json:"destinations,omitempty"`
	IcmpBytesRx int64 `json:"icmpBytesRx,omitempty"`
	IcmpBytesTx int64 `json:"icmpBytesTx,omitempty"`
	IcmpPacketsRx int64 `json:"icmpPacketsRx,omitempty"`
	IcmpPacketsTx int64 `json:"icmpPacketsTx,omitempty"`
	Interface string `json:"interface,omitempty"`
	InternalId string `json:"internalId,omitempty"`
	LocalIpAddress string `json:"localIpAddress,omitempty"`
	LogicalId string `json:"logicalId,omitempty"`
	Mode string `json:"mode,omitempty"`
	Mtu int32 `json:"mtu,omitempty"`
	Name string `json:"name,omitempty"`
	OtherBytesRx int64 `json:"otherBytesRx,omitempty"`
	OtherBytesTx int64 `json:"otherBytesTx,omitempty"`
	OtherPacketsRx int64 `json:"otherPacketsRx,omitempty"`
	OtherPacketsTx int64 `json:"otherPacketsTx,omitempty"`
	P1BytesRx int64 `json:"p1BytesRx,omitempty"`
	P1BytesTx int64 `json:"p1BytesTx,omitempty"`
	P1PacketsRx int64 `json:"p1PacketsRx,omitempty"`
	P1PacketsTx int64 `json:"p1PacketsTx,omitempty"`
	P2BytesRx int64 `json:"p2BytesRx,omitempty"`
	P2BytesTx int64 `json:"p2BytesTx,omitempty"`
	P2PacketsRx int64 `json:"p2PacketsRx,omitempty"`
	P2PacketsTx int64 `json:"p2PacketsTx,omitempty"`
	P3BytesRx int64 `json:"p3BytesRx,omitempty"`
	P3BytesTx int64 `json:"p3BytesTx,omitempty"`
	P3PacketsRx int64 `json:"p3PacketsRx,omitempty"`
	P3PacketsTx int64 `json:"p3PacketsTx,omitempty"`
	PublicIpAddress string `json:"publicIpAddress,omitempty"`
	ScoreRx float32 `json:"scoreRx,omitempty"`
	ScoreTx float32 `json:"scoreTx,omitempty"`
	SignalStrength float32 `json:"signalStrength,omitempty"`
	State string `json:"state,omitempty"`
	TcpBytesRx int64 `json:"tcpBytesRx,omitempty"`
	TcpBytesTx int64 `json:"tcpBytesTx,omitempty"`
	TcpPacketsRx int64 `json:"tcpPacketsRx,omitempty"`
	TcpPacketsTx int64 `json:"tcpPacketsTx,omitempty"`
	Type string `json:"type,omitempty"`
	UdpBytesRx int64 `json:"udpBytesRx,omitempty"`
	UdpBytesTx int64 `json:"udpBytesTx,omitempty"`
	UdpHolePunching int32 `json:"udpHolePunching,omitempty"`
	UdpPacketsRx int64 `json:"udpPacketsRx,omitempty"`
	UdpPacketsTx int64 `json:"udpPacketsTx,omitempty"`
	VlanId int32 `json:"vlanId,omitempty"`
	VpnState string `json:"vpnState,omitempty"`
}
