package main

import (
	"testing"
)

func TestParsePayload(t *testing.T) {
	p := Position{
		MMSI: ParsePayload("13u?etPv2;0n:dDPwUM1U1Cb069D", 8, 37),
		Status: ParsePayload("13u?etPv2;0n:dDPwUM1U1Cb069D", 38, 41),
		Speed: ParsePayload("13u?etPv2;0n:dDPwUM1U1Cb069D", 50, 59),
		Longitude: ParsePayload("13u?etPv2;0n:dDPwUM1U1Cb069D", 61, 88),
		Latitude: ParsePayload("13u?etPv2;0n:dDPwUM1U1Cb069D", 89, 115),
	}

	if p.MMSI != 265547250 && p.Status != 0 && float32(p.Speed) != 13.9 {
		t.Error("Test of ParsePayload() failed")
	}
}
