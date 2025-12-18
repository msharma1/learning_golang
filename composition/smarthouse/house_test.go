package main

import (
	"testing"
)

func TestSmartBulbPower(t *testing.T) {
	sb := SmartBulb{}

	if sb.IsOn {
		t.Errorf("New bulb should be off")
	}

	sb.TurnOn()
	if !sb.IsOn {
		t.Errorf("Bulb should now be on")
	}

	sb.TurnOff()
	if sb.IsOn {
		t.Errorf("Bulb should now be off")
	}
}

func TestBulbWifiConnector(t *testing.T) {
	sb := SmartBulb{}
	expectedNetwork := "Office wifi"

	if sb.IsConnected {
		t.Error("New bulb should not be connected")
	}

	sb.ConnectToWiFi("Office wifi")

	if !sb.IsConnected || sb.NetworkName != expectedNetwork {
		t.Error("New bulb should be connected to Office wifi")
	}
}
