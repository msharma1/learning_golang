package main

import (
	"fmt"
)

type PowerCompoment struct{}

func (ol *PowerCompoment) TurnOn() {
	fmt.Println("Turned On")
}
func (ol *PowerCompoment) TurnOff() {
	fmt.Println("Turned Off")
}

type WifiConnector struct{}

func (mb *WifiConnector) ConnectToWiFi() {
	fmt.Println("Wifi Connected")
}

type MusicPlayer struct{}

func (mb *MusicPlayer) PlayMusic() {
	fmt.Println("Playing Music")

}

type OldLamp struct {
	PowerCompoment
}

type MusicBox struct {
	WifiConnector
	MusicPlayer
}

type SmartBulb struct {
	OldLamp
	WifiConnector
}

func main() {
	sb := SmartBulb{}
	sb.TurnOn()
	sb.TurnOff()
	sb.ConnectToWiFi()
}
