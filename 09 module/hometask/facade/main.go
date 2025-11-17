package main

import "fmt"

type TV struct {
	channel int
}

func (t *TV) On() {
	fmt.Println("TV включен")
}

func (t *TV) Off() {
	fmt.Println("TV выключен")
}

func (t *TV) SetChannel(channel int) {
	t.channel = channel
	fmt.Printf("TV: установлен канал %d\n", channel)
}

// AudioSystem

type AudioSystem struct {
	volume int
}

func (a *AudioSystem) On() {
	fmt.Println("AudioSystem: включена")
}

func (a *AudioSystem) Off() {
	fmt.Println("AudioSystem: включена")
}

func (a *AudioSystem) SetVolume(v int) {
	a.volume = v
	fmt.Printf("AudioSystem: громкость установлена на %d\n", v)
}

// DVDPlayer

type DVDPlayer struct{}

func (d *DVDPlayer) Play() {
	fmt.Println("DVDPlayer: воспроизведение")
}

func (d *DVDPlayer) Pause() {
	fmt.Println("DVDPlayer: пауза")
}

func (d *DVDPlayer) Stop() {
	fmt.Println("DVDPlayer: остановка")
}

// GameConsole

type GameConsole struct{}

func (g *GameConsole) On() {
	fmt.Println("GameConsole: включена")
}

func (g *GameConsole) StartGame() {
	fmt.Println("GameConsole: запуск игры...")
}

// Фасад

type HomeTheaterFacade struct {
	tv          *TV
	audio       *AudioSystem
	dvd         *DVDPlayer
	gameConsole *GameConsole
}

func NewHomeTheaterFacade(tv *TV, audio *AudioSystem, dvd *DVDPlayer, gc *GameConsole) *HomeTheaterFacade {
	return &HomeTheaterFacade{tv, audio, dvd, gc}
}

func (h *HomeTheaterFacade) WatchMovie() {
	fmt.Println("\n Просмотр Фильма")
	h.tv.On()
	h.tv.SetChannel(1)
	h.audio.On()
	h.audio.SetVolume(15)
	h.dvd.Play()
}

func (h *HomeTheaterFacade) PlayGame() {
	fmt.Println("\n Запуск игры")
	h.tv.On()
	h.audio.On()
	h.audio.SetVolume(20)
	h.gameConsole.On()
	h.gameConsole.StartGame()
}

func(h *HomeTheaterFacade) PlayMusic() {
	fmt.Println("\n Включение музыки")
	h.tv.On()
	h.audio.On()
	h.audio.SetVolume(20)
}

func (h *HomeTheaterFacade) Shutdown() {
	fmt.Println("\n Отключение системы")
	h.dvd.Stop()
	h.audio.Off()
	h.tv.Off()
}

func main() {
	tv := &TV{}
	audio := &AudioSystem{}
	dvd := &DVDPlayer{}
	gc := &GameConsole{}

	home := NewHomeTheaterFacade(tv, audio, dvd, gc)

	home.WatchMovie()
	home.Shutdown()

	home.PlayGame()
	home.Shutdown()

	home.PlayMusic()
	home.Shutdown()
}