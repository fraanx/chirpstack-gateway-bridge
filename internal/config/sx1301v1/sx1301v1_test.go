package sx1301v1

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/fraanx/chirpstack/api/go/v4/gw"
)

func TestGetRadioFrequencies(t *testing.T) {
	tests := []struct {
		Name             string
		Channels         []*gw.ChannelConfiguration
		RadioFrequencies [2]uint32
		Error            error
	}{
		{
			Name: "one channel",
			Channels: []*gw.ChannelConfiguration{
				{
					Frequency: 868100000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth: 125000,
						},
					},
				},
			},
			RadioFrequencies: [2]uint32{868500000},
		},
		{
			Name: "channels don't fit",
			Channels: []*gw.ChannelConfiguration{
				{
					Frequency: 867100000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth: 125000,
						},
					},
				},
				{
					Frequency: 868100000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth: 125000,
						},
					},
				},
				{
					Frequency: 869100000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth: 125000,
						},
					},
				},
			},
			Error: errors.New("channel 869100000 does not fit in radio bandwidth"),
		},
		{
			Name: "EU868 three channels",
			Channels: []*gw.ChannelConfiguration{
				{
					Frequency: 868100000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth: 125000,
						},
					},
				},
				{
					Frequency: 868300000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth: 125000,
						},
					},
				},
				{
					Frequency: 868500000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth: 125000,
						},
					},
				},
			},
			RadioFrequencies: [2]uint32{868500000},
		},
		{
			Name: "EU868 8 channels + single SF + FSK",
			Channels: []*gw.ChannelConfiguration{
				{
					Frequency: 868100000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth: 125000,
						},
					},
				},
				{
					Frequency: 868300000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth: 125000,
						},
					},
				},
				{
					Frequency: 868500000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth: 125000,
						},
					},
				},
				{
					Frequency: 867100000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10, 11, 12},
						},
					},
				},
				{
					Frequency: 867300000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10, 11, 12},
						},
					},
				},
				{
					Frequency: 867500000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10, 11, 12},
						},
					},
				},
				{
					Frequency: 867700000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10, 11, 12},
						},
					},
				},
				{
					Frequency: 867900000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10, 11, 12},
						},
					},
				},
				{
					Frequency: 868300000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        250000,
							SpreadingFactors: []uint32{7},
						},
					},
				},
				{
					Frequency: 868800000,
					ModulationConfig: &gw.ChannelConfiguration_FskModulationConfig{
						FskModulationConfig: &gw.FskModulationConfig{
							Bandwidth: 125000,
							Bitrate:   50000,
						},
					},
				},
			},
			RadioFrequencies: [2]uint32{867500000, 868500000},
		},
		{
			Name: "US915 0-7 + 64",
			Channels: []*gw.ChannelConfiguration{
				{
					Frequency: 902300000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10},
						},
					},
				},
				{
					Frequency: 902500000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10},
						},
					},
				},
				{
					Frequency: 902700000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10},
						},
					},
				},
				{
					Frequency: 902900000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10},
						},
					},
				},
				{
					Frequency: 903100000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10},
						},
					},
				},
				{
					Frequency: 903300000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10},
						},
					},
				},
				{
					Frequency: 903500000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10},
						},
					},
				},
				{
					Frequency: 903700000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        125000,
							SpreadingFactors: []uint32{7, 8, 9, 10},
						},
					},
				},
				{
					Frequency: 903000000,
					ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
						LoraModulationConfig: &gw.LoraModulationConfig{
							Bandwidth:        500000,
							SpreadingFactors: []uint32{8},
						},
					},
				},
			},
			RadioFrequencies: [2]uint32{902700000, 903700000},
		},
	}

	for _, tst := range tests {
		t.Run(tst.Name, func(t *testing.T) {
			assert := require.New(t)

			rf, err := GetRadioFrequencies(tst.Channels)
			assert.Equal(tst.Error, err)

			if err == nil {
				assert.Equal(tst.RadioFrequencies, rf)
			}
		})
	}
}

func TestGetRadioForChannel(t *testing.T) {
	tests := []struct {
		Name             string
		RadioFrequencies [2]uint32
		Channel          *gw.ChannelConfiguration
		Radio            int
		Error            error
	}{
		{
			Name:             "Radio 0",
			RadioFrequencies: [2]uint32{868500000},
			Channel: &gw.ChannelConfiguration{
				Frequency: 868100000,
				ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
					LoraModulationConfig: &gw.LoraModulationConfig{
						Bandwidth: 125000,
					},
				},
			},
			Radio: 0,
		},
		{
			Name:             "Out of bandwidth",
			RadioFrequencies: [2]uint32{868500000},
			Channel: &gw.ChannelConfiguration{
				Frequency: 869100000,
				ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
					LoraModulationConfig: &gw.LoraModulationConfig{
						Bandwidth: 125000,
					},
				},
			},
			Error: errors.New("channel 869100000 does not fit in radio bandwidth"),
		},
		{
			Name:             "Radio 1",
			RadioFrequencies: [2]uint32{867500000, 868500000},
			Channel: &gw.ChannelConfiguration{
				Frequency: 868100000,
				ModulationConfig: &gw.ChannelConfiguration_LoraModulationConfig{
					LoraModulationConfig: &gw.LoraModulationConfig{
						Bandwidth: 125000,
					},
				},
			},
			Radio: 1,
		},
	}

	for _, tst := range tests {
		t.Run(tst.Name, func(t *testing.T) {
			assert := require.New(t)

			r, err := GetRadioForChannel(tst.RadioFrequencies, tst.Channel)
			assert.Equal(tst.Error, err)

			if err == nil {
				assert.Equal(tst.Radio, r)
			}
		})
	}
}
