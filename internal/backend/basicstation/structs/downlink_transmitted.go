package structs

import (
	"github.com/fraanx/chirpstack-v3-api/go/v3/gw"
	"github.com/fraanx/lorawan"
)

// DownlinkTransmitted implements the downlink transmitted message.
type DownlinkTransmitted struct {
	MessageType MessageType `json:"msgtype"`

	DIID uint32 `json:"diid"`
}

// DownlinkTransmittedToProto converts the DownlinkTransmitted to the protobuf struct.
func DownlinkTransmittedToProto(gatewayID lorawan.EUI64, dt DownlinkTransmitted) (gw.DownlinkTXAck, error) {
	return gw.DownlinkTXAck{
		GatewayId: gatewayID[:],
		Token:     dt.DIID,
		Items: []*gw.DownlinkTXAckItem{
			{
				Status: gw.TxAckStatus_OK,
			},
		},
	}, nil
}
