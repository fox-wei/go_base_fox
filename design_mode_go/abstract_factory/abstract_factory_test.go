package abstract_factory

import (
	"testing"
)

func TestNewlunchFactory(t *testing.T) {
	var factory AbstractFactory
	var tv ITelevision
	var air IAirConditioner

	factory = &HuaweiFactory{}
	tv = factory.CreateTelevision()
	air = factory.CreateAirConditioner()
	tv.Wathch()
	air.SetTemperature(26)
}

func TestMiAirConditioner_SetTemperature(t *testing.T) {
	type args struct {
		temperature int
	}
	tests := []struct {
		name string
		ma   *MiAirConditioner
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ma.SetTemperature(tt.args.temperature)
		})
	}
}
