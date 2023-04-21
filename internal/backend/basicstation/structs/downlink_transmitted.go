package structs

import (
	"github.com/fraanx/lorawan"
	"github.com/fraanx/chirpstack/api/go/v4/gw"
)

// DownlinkTransmitted implements the downlink transmitted message.
type DownlinkTransmitted struct {
	MessageType MessageType `json:"msgtype"`

	DIID uint32 `json:"diid"`
}

// DownlinkTransmittedToProto converts the DownlinkTransmitted to the protobuf struct.
func DownlinkTransmittedToProto(gatewayID lorawan.EUI64, dt DownlinkTransmitted) (gw.DownlinkTxAck, error) {
	return gw.DownlinkTxAck{
		GatewayId:  gatewayID.String(),
		DownlinkId: dt.DIID,
		Items: []*gw.DownlinkTxAckItem{
			{
				Status: gw.TxAckStatus_OK,
			},
		},
	}, nil
}
