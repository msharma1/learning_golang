package main

import "fmt"

type PowerCompoment struct {
	IsOn bool
}

func (p *PowerCompoment) TurnOn() {
	p.IsOn = true
}
func (p *PowerCompoment) TurnOff() {
	p.IsOn = false
}

type WifiConnector struct {
	IsConnected bool
	NetworkName string
}

func (w *WifiConnector) ConnectToWiFi(network string) {
	w.IsConnected = true
	w.NetworkName = network
}

type MusicPlayer struct {
	IsPlaying   bool
	CurrentSong string
}

func (m *MusicPlayer) PlayMusic(songName string) {
	m.IsPlaying = true
	m.CurrentSong = songName

}

type OldLamp struct {
	PowerCompoment
}

type MusicBox struct {
	WifiConnector
	MusicPlayer
}

type SmartBulb struct {
	PowerCompoment
	WifiConnector
}

func main() {
	sb := SmartBulb{}
	sb.TurnOn()
	fmt.Printf("is on? %v\n", sb.IsOn)
	sb.TurnOff()
	fmt.Printf("is on? %v\n", sb.IsOn)
	sb.ConnectToWiFi("Home-5G")
	fmt.Printf("connected to wifi? %v, network name: %s\n", sb.IsConnected, sb.NetworkName)

	mb := MusicBox{}
	mb.PlayMusic("ABCD")
	fmt.Printf("music on? %v\n", mb.IsPlaying)
	fmt.Printf("current song? %s\n", mb.CurrentSong)
}
