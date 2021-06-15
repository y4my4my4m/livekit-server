package types

type ProtocolVersion int

const DefaultProtocol = 2
const LatestProtocol = 3

func (v ProtocolVersion) SupportsPackedStreamId() bool {
	return v > 0
}

// data packets was introduced in v2
func (v ProtocolVersion) HandlesDataPackets() bool {
	return v > 1
}

// subscription map is introduced in v3
func (v ProtocolVersion) SupportsSubscriptionMap() bool {
	return v > 2
}
