package main

import (
	"container/list"
	"fmt"
)

type ICommand interface {
	Execute()
	Undo()
}

type NoCommand struct{}

func (c *NoCommand) Execute() {
	fmt.Println("[пульт] слот не назначен.")
}
func (c *NoCommand) Undo() {
	fmt.Println("[пульт] слот не назначен (отмена).")
}

type Light struct {
	Location string
}

func (l *Light) On() {
	fmt.Printf("свет в '%s' включен.\n", l.Location)
}
func (l *Light) Off() {
	fmt.Printf("свет в '%s' выключен.\n", l.Location)
}

type Television struct {
	Location string
	Channel  int
}

func (tv *Television) On() {
	fmt.Printf("телевизор в '%s' включен.\n", tv.Location)
}
func (tv *Television) Off() {
	fmt.Printf("телевизор в '%s' выключен.\n", tv.Location)
}
func (tv *Television) GetChannel() int { return tv.Channel }
func (tv *Television) SetChannel(ch int) {
	tv.Channel = ch
	fmt.Printf("телевизор в '%s' переключен на канал %d.\n", tv.Location, ch)
}

type AirConditioner struct {
	Location    string
	Temperature int
}

func (ac *AirConditioner) On() {
	fmt.Printf("кондиционер в '%s' включен.\n", ac.Location)
}
func (ac *AirConditioner) Off() {
	fmt.Printf("кондиционер в '%s' выключен.\n", ac.Location)
}
func (ac *AirConditioner) GetTemperature() int { return ac.Temperature }
func (ac *AirConditioner) SetTemperature(temp int) {
	ac.Temperature = temp
	fmt.Printf("температура в '%s' установлена на %d°c.\n", ac.Location, ac.Temperature)
}

type SmartBlinds struct {
	Location string
}

func (b *SmartBlinds) Open() {
	fmt.Printf("шторы в '%s' открыты.\n", b.Location)
}
func (b *SmartBlinds) Close() {
	fmt.Printf("шторы в '%s' закрыты.\n", b.Location)
}

type MusicPlayer struct {
	Volume int
}

func (p *MusicPlayer) Play() {
	fmt.Println("музыка включена.")
}
func (p *MusicPlayer) Stop() {
	fmt.Println("музыка выключена.")
}
func (p *MusicPlayer) GetVolume() int { return p.Volume }
func (p *MusicPlayer) SetVolume(vol int) {
	p.Volume = vol
	fmt.Printf("громкость установлена на %d.\n", vol)
}

type LightOnCommand struct{ light *Light }

func (c *LightOnCommand) Execute() { c.light.On() }
func (c *LightOnCommand) Undo()    { c.light.Off() }

type LightOffCommand struct{ light *Light }

func (c *LightOffCommand) Execute() { c.light.Off() }
func (c *LightOffCommand) Undo()    { c.light.On() }

type TVOnCommand struct{ tv *Television }

func (c *TVOnCommand) Execute() { c.tv.On() }
func (c *TVOnCommand) Undo()    { c.tv.Off() }

type TVOffCommand struct{ tv *Television }

func (c *TVOffCommand) Execute() { c.tv.Off() }
func (c *TVOffCommand) Undo()    { c.tv.On() }

type TVSetChannelCommand struct {
	tv             *Television
	newChannel     int
	previousChannel int
}

func (c *TVSetChannelCommand) Execute() {
	c.previousChannel = c.tv.GetChannel()
	c.tv.SetChannel(c.newChannel)
}
func (c *TVSetChannelCommand) Undo() {
	c.tv.SetChannel(c.previousChannel)
}

type ACSetTempCommand struct {
	ac           *AirConditioner
	newTemp      int
	previousTemp int
}

func (c *ACSetTempCommand) Execute() {
	c.previousTemp = c.ac.GetTemperature()
	c.ac.SetTemperature(c.newTemp)
}
func (c *ACSetTempCommand) Undo() {
	c.ac.SetTemperature(c.previousTemp)
}

type BlindsOpenCommand struct{ blinds *SmartBlinds }

func (c *BlindsOpenCommand) Execute() { c.blinds.Open() }
func (c *BlindsOpenCommand) Undo()    { c.blinds.Close() }

type BlindsCloseCommand struct{ blinds *SmartBlinds }

func (c *BlindsCloseCommand) Execute() { c.blinds.Close() }
func (c *BlindsCloseCommand) Undo()    { c.blinds.Open() }

type PlayerPlayCommand struct{ player *MusicPlayer }

func (c *PlayerPlayCommand) Execute() { c.player.Play() }
func (c *PlayerPlayCommand) Undo()    { c.player.Stop() }

type PlayerStopCommand struct{ player *MusicPlayer }

func (c *PlayerStopCommand) Execute() { c.player.Stop() }
func (c *PlayerStopCommand) Undo()    { c.player.Play() }

type GroupCommand struct {
	commands []ICommand
}

func NewGroupCommand() *GroupCommand {
	return &GroupCommand{commands: make([]ICommand, 0)}
}
func (m *GroupCommand) AddCommand(cmd ICommand) {
	m.commands = append(m.commands, cmd)
}
func (m *GroupCommand) Execute() {
	fmt.Println("--- [группа] выполнение ---")
	for _, cmd := range m.commands {
		cmd.Execute()
	}
	fmt.Println("--- [группа] завершен ---")
}
func (m *GroupCommand) Undo() {
	fmt.Println("--- [группа] отмена ---")
	for i := len(m.commands) - 1; i >= 0; i-- {
		m.commands[i].Undo()
	}
	fmt.Println("--- [группа] отмена завершена ---")
}

type RemoteControl struct {
	onCommands  map[int]ICommand
	offCommands map[int]ICommand

	undoStack *list.List
	redoStack *list.List

	isRecording  bool
	currentGroup *GroupCommand
}

func NewRemoteControl() *RemoteControl {
	noCmd := &NoCommand{}
	on := make(map[int]ICommand)
	off := make(map[int]ICommand)
	for i := 0; i < 7; i++ {
		on[i] = noCmd
		off[i] = noCmd
	}

	return &RemoteControl{
		onCommands:     on,
		offCommands:    off,
		undoStack:      list.New(),
		redoStack:      list.New(),
		isRecording:    false,
	}
}

func (r *RemoteControl) SetCommand(slot int, on ICommand, off ICommand) {
	r.onCommands[slot] = on
	r.offCommands[slot] = off
}

func (r *RemoteControl) executeCommand(cmd ICommand) {
	if r.isRecording {
		fmt.Printf("[запись] добавлена команда в группу...\n")
		r.currentGroup.AddCommand(cmd)
		return
	}

	cmd.Execute()
	r.undoStack.PushBack(cmd)
	r.redoStack.Init()
}

func (r *RemoteControl) PressOnButton(slot int) {
	fmt.Printf("\n--- нажата [вкл] (слот %d) ---\n", slot)
	r.executeCommand(r.onCommands[slot])
}

func (r *RemoteControl) PressOffButton(slot int) {
	fmt.Printf("\n--- нажата [выкл] (слот %d) ---\n", slot)
	r.executeCommand(r.offCommands[slot])
}

func (r *RemoteControl) PressUndoButton() {
	fmt.Println("\n--- нажата [отмена] ---")
	if r.undoStack.Len() == 0 {
		fmt.Println("нечего отменять.")
		return
	}

	lastCmdElement := r.undoStack.Back()
	lastCmd := lastCmdElement.Value.(ICommand)

	lastCmd.Undo()

	r.undoStack.Remove(lastCmdElement)
	r.redoStack.PushBack(lastCmd)
}

func (r *RemoteControl) PressRedoButton() {
	fmt.Println("\n--- нажата [повтор] ---")
	if r.redoStack.Len() == 0 {
		fmt.Println("нечего повторять.")
		return
	}

	lastRedoElement := r.redoStack.Back()
	lastRedoCmd := lastRedoElement.Value.(ICommand)

	lastRedoCmd.Execute()

	r.redoStack.Remove(lastRedoElement)
	r.undoStack.PushBack(lastRedoCmd)
}

func (r *RemoteControl) StartRecording() {
	if r.isRecording {
		fmt.Println("[запись] уже идет запись.")
		return
	}
	r.isRecording = true
	r.currentGroup = NewGroupCommand()
	fmt.Println("\n--- [запись] начата запись группы ---")
}

func (r *RemoteControl) StopRecordingAndSet(slot int, on bool) {
	if !r.isRecording {
		fmt.Println("[запись] запись не была начата.")
		return
	}
	r.isRecording = false

	if on {
		r.onCommands[slot] = r.currentGroup
		fmt.Printf("--- [запись] запись завершена. группа назначена на слот %d [вкл] ---\n", slot)
	} else {
		r.offCommands[slot] = r.currentGroup
		fmt.Printf("--- [запись] запись завершена. группа назначена на слот %d [выкл] ---\n", slot)
	}
	r.currentGroup = nil
}

func main() {
	remote := NewRemoteControl()

	livingRoomLight := &Light{Location: "гостиная"}
	mainTV := &Television{Location: "гостиная", Channel: 1}
	mainAC := &AirConditioner{Location: "спальня", Temperature: 22}
	bedroomBlinds := &SmartBlinds{Location: "спальня"}
	mainPlayer := &MusicPlayer{Volume: 10}

	lrLightOn := &LightOnCommand{light: livingRoomLight}
	lrLightOff := &LightOffCommand{light: livingRoomLight}

	tvOn := &TVOnCommand{tv: mainTV}
	tvOff := &TVOffCommand{tv: mainTV}
	tvSetChannel5 := &TVSetChannelCommand{tv: mainTV, newChannel: 5}

	acSet25 := &ACSetTempCommand{ac: mainAC, newTemp: 25}

	blindsOpen := &BlindsOpenCommand{blinds: bedroomBlinds}
	blindsClose := &BlindsCloseCommand{blinds: bedroomBlinds}

	playerPlay := &PlayerPlayCommand{player: mainPlayer}
	playerStop := &PlayerStopCommand{player: mainPlayer}

	remote.SetCommand(0, lrLightOn, lrLightOff)
	remote.SetCommand(1, tvOn, tvOff)
	remote.SetCommand(2, blindsOpen, blindsClose)
	remote.SetCommand(3, playerPlay, playerStop)
	remote.SetCommand(4, tvSetChannel5, acSet25)

	fmt.Println("===== тест 1: команды и undo/redo =====")

	remote.PressOnButton(0)
	remote.PressOnButton(1)
	remote.PressOnButton(4)

	remote.PressUndoButton()
	remote.PressUndoButton()

	remote.PressRedoButton()
	remote.PressRedoButton()

	remote.PressOffButton(0)

	remote.PressRedoButton()

	fmt.Println("\n===== тест 2: группа 'доброе утро' =====")
	groupMorning := NewGroupCommand()
	groupMorning.AddCommand(blindsOpen)
	groupMorning.AddCommand(lrLightOn)
	groupMorning.AddCommand(&ACSetTempCommand{ac: mainAC, newTemp: 23})

	remote.SetCommand(5, groupMorning, &NoCommand{})
	remote.PressOnButton(5)

	remote.PressUndoButton()

	fmt.Println("\n===== тест 3: пустой слот =====")
	remote.PressOnButton(6)
	remote.PressUndoButton()

	fmt.Println("\n===== тест 4: запись группы 'кино' =====")
	remote.StartRecording()

	remote.PressOffButton(0)
	remote.PressOnButton(1)
	remote.PressOnButton(3)

	remote.StopRecordingAndSet(6, true)

	fmt.Println("\n(--- тестируем записанную группу 'кино' ---)")
	remote.PressOnButton(6)

	remote.PressUndoButton()
}