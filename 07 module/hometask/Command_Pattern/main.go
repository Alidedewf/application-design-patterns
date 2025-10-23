package main

import (
	"fmt"
)

type ICommand interface {
	Execute()
	Undo()
}

type Light struct {
	state string
}

func (l *Light) On() {
	l.state = "Включен"
	fmt.Println("Свет включен")
}

func (l *Light) Off() {
	l.state = "Выключен"
	fmt.Println("Свет выключен")
}

type Door struct {
	state string
}

func (d *Door) Open() {
	d.state = "open"
	fmt.Println("Дверь открыта")
}

func (d *Door) Close() {
	d.state = "Закрыто"
	fmt.Println("Дверь закрыта")
}

type Thermostat struct {
	temperature int
}

func (t *Thermostat) IncreaseTemp() {
	t.temperature++
	fmt.Printf("Температура увеличена до %d°C\n", t.temperature)
}

func (t *Thermostat) DecreaseTemp() {
	t.temperature--
	fmt.Printf("Температура уменьшена до %d°C\n", t.temperature)
}

type TV struct {
	state string
}

func (tv *TV) On() {
	tv.state = "Включен"
	fmt.Println("Телевизор включен")
}

func (tv *TV) Off() {
	tv.state = "Выключено"
	fmt.Println("Телевизор выключен")
}

type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()

}
func (c *LightOnCommand) Undo() {
	c.light.Off()
}

type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()

}
func (c *LightOffCommand) Undo() {
	c.light.On()
}

type DoorOpenCommand struct {
	door *Door
}

func (c *DoorOpenCommand) Execute() {
	c.door.Open()
}

func (c *DoorOpenCommand) Undo() {
	c.door.Close()
}

type DoorCloseCommand struct {
	door *Door
}

func (c *DoorCloseCommand) Execute() {
	c.door.Close()
}

func (c *DoorCloseCommand) Undo() {
	c.door.Open()
}

type TempIncreaseCommand struct {
	thermostat *Thermostat
}

func (c *TempIncreaseCommand) Execute() {
	c.thermostat.IncreaseTemp()
}

func (c *TempIncreaseCommand) Undo() {
	c.thermostat.DecreaseTemp()
}

type TempDecreaseCommand struct {
	thermostat *Thermostat
}

func (c *TempDecreaseCommand) Execute() {
	c.thermostat.DecreaseTemp()
}

func (c *TempDecreaseCommand) Undo() {
	c.thermostat.IncreaseTemp()
}

type TVOnCommand struct {
	tv *TV
}

func (c *TVOnCommand) Execute() {
	c.tv.On()
}

func (c *TVOnCommand) Undo() {
	c.tv.Off()
}

type RemoteControl struct {
	commandHistory []ICommand
}

func NewRemoteControl() *RemoteControl {
	return &RemoteControl{
		commandHistory: make([]ICommand, 0),
	}
}

func (r *RemoteControl) ExecuteCommand(cmd ICommand) {
	fmt.Println("-- Выполнение команды -- ")
	cmd.Execute()
	r.commandHistory = append(r.commandHistory, cmd)
}

func (r *RemoteControl) PressUndo() {
	fmt.Println("-- Отмена команды -- ")
	if len(r.commandHistory) == 0 {
		fmt.Println("Нет команд для отмены")
		return
	}
	lastIndex := len(r.commandHistory) - 1
	lastCommand := r.commandHistory[lastIndex]

	lastCommand.Undo()

	r.commandHistory = r.commandHistory[:lastIndex]
}

func main() {
	remote := NewRemoteControl()

	livingRoomLight := &Light{}
	mainDoor := &Door{}
	mainThermostat := &Thermostat{temperature: 20}
	livingRoomTV := &TV{}

	lightOn := &LightOnCommand{light: livingRoomLight}
	lightOff := &LightOffCommand{light: livingRoomLight}
	doorOpen := &DoorOpenCommand{door: mainDoor}
	tempUp := &TempIncreaseCommand{thermostat: mainThermostat}
	tvOn := &TVOnCommand{tv: livingRoomTV}

	// --- 2. Тестирование ---
	remote.ExecuteCommand(lightOn)
	remote.ExecuteCommand(doorOpen)
	remote.ExecuteCommand(tempUp)
	remote.ExecuteCommand(tvOn)

	fmt.Println("\n--- Начинаем отмену (Undo) ---")

	// (Тест отмены)
	remote.PressUndo()
	remote.PressUndo()
	remote.PressUndo()

	remote.ExecuteCommand(lightOff)

	remote.PressUndo()
	remote.PressUndo()

	remote.PressUndo()
}
