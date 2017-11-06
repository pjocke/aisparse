package main

import (
	"fmt"
)

type Position struct{
	MMSI		uint32
	Status		uint32
	Speed		uint32
	Longitude	uint32
	Latitude	uint32
}

func main() {
	p := Position{
		MMSI: parsePayload("13u?etPv2;0n:dDPwUM1U1Cb069D", 8, 37),
		Status: parsePayload("13u?etPv2;0n:dDPwUM1U1Cb069D", 38, 41),
		Speed: parsePayload("13u?etPv2;0n:dDPwUM1U1Cb069D", 50, 59),
		Longitude: parsePayload("13u?etPv2;0n:dDPwUM1U1Cb069D", 61, 88),
		Latitude: parsePayload("13u?etPv2;0n:dDPwUM1U1Cb069D", 89, 115),
	}

	Statuses := map[uint32]string{
		0: "under way using engine",
		1: "at anchor",
		5: "moored",
		8: "under way sailing",
	}

	fmt.Printf("Vessel with MMSI %d with position %f N, %f E(?) is %s going %d knots.\n", p.MMSI, float32(p.Latitude)/600000, float32(p.Longitude)/600000, Statuses[p.Status], p.Speed)
}


func parsePayload(s string, first, last int) uint32 {
	payload := []byte(s)

	start := first / 6
	loops := (last - first + 1) / 6

	size := last - first + 1
	position := 0

	var result uint32

	for i := 0; i <= loops; i++ {
		temp := uint32(payload[start + i]) - 48
		if temp > 40 {
			temp -= 8
		}

		if i == 0 {
			// Clear leftmost bits if needed.
			position = 6 - (first %6)
			temp = temp << uint(32 - position) >> uint(32 - size)
		} else if i < loops {
			temp = temp << uint(size - position)
		} else if i == loops {
			// Clear rightmost bits if needed.
			temp = temp >> uint(6 - (last % 6) - 1)
		}
		position = position + 6

		result = result | temp
	}
	return(result)
}
