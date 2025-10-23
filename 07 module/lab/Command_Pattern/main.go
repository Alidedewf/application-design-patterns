package main

import (
	"fmt"
	"time"
)

type ICommand interface {
	Execute()
	Undo()
}

type Light struct {
	Location string
}

func (l *Light) On() {
	fmt.Printf("Свет в '%s' включен.\n", l.Location)
}
func (l *Light) Off() {
	fmt.Printf("Свет в '%s' выключен.\n", l.Location)
}

type Television struct {
	Location string
}

func (tv *Television) On() {
	fmt.Printf("Телевизор в '%s' включен.\n", tv.Location)
}
func (tv *Television) Off() {
	fmt.Printf("Телевизор в '%s' выключен.\n", tv.Location)
}

type AirConditioner struct {
	Location    string
	Temperature int
}

func (ac *AirConditioner) On() {
	fmt.Printf("Кондиционер в '%s' включен.\n", ac.Location)
}
func (ac *AirConditioner) Off() {
	fmt.Printf("Кондиционер в '%s' выключен.\n", ac.Location)
}
func (ac *AirConditioner) SetTemperature(temp int) {
	ac.Temperature = temp
	fmt.Printf("Температура в '%s' установлена на %d°C.\n", ac.Location, ac.Temperature)
}
func (ac *AirConditioner) GetTemperature() int {
	return ac.Temperature
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}
func (c *LightOnCommand) Execute() { c.light.On() }
func (c *LightOnCommand) Undo()    { c.light.Off() }

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}
func (c *LightOffCommand) Execute() { c.light.Off() }
func (c *LightOffCommand) Undo()    { c.light.On() }

type TVOnCommand struct {
	tv *Television
}

func NewTVOnCommand(tv *Television) *TVOnCommand {
	return &TVOnCommand{tv: tv}
}
func (c *TVOnCommand) Execute() { c.tv.On() }
func (c *TVOnCommand) Undo()    { c.tv.Off() }

type TVOffCommand struct {
	tv *Television
}

func NewTVOffCommand(tv *Television) *TVOffCommand {
	return &TVOffCommand{tv: tv}
}
func (c *TVOffCommand) Execute() { c.tv.Off() }
func (c *TVOffCommand) Undo()    { c.tv.On() }

type ACOnCommand struct {
	ac *AirConditioner
}

func NewACOnCommand(ac *AirConditioner) *ACOnCommand {
	return &ACOnCommand{ac: ac}
}
func (c *ACOnCommand) Execute() { c.ac.On() }
func (c *ACOnCommand) Undo()    { c.ac.Off() }

type ACOffCommand struct {
	ac *AirConditioner
}

func NewACOffCommand(ac *AirConditioner) *ACOffCommand {
	return &ACOffCommand{ac: ac}
}
func (c *ACOffCommand) Execute() { c.ac.Off() }
func (c *ACOffCommand) Undo()    { c.ac.On() }

type ACSetTempCommand struct {
	ac           *AirConditioner
	newTemp      int
	previousTemp int
}

func NewACSetTempCommand(ac *AirConditioner, temp int) *ACSetTempCommand {
	return &ACSetTempCommand{ac: ac, newTemp: temp}
}

func (c *ACSetTempCommand) Execute() {
	c.previousTemp = c.ac.GetTemperature()
	c.ac.SetTemperature(c.newTemp)
}
func (c *ACSetTempCommand) Undo() {
	c.ac.SetTemperature(c.previousTemp)
}

type NoCommand struct{}

func (c *NoCommand) Execute() {
	fmt.Println("[ПУЛЬТ] Слот не запрограммирован.")
}
func (c *NoCommand) Undo() {
	fmt.Println("[ПУЛЬТ] Слот не запрограммирован (отмена).")
}

type MacroCommand struct {
	commands []ICommand
}

func NewMacroCommand(commands []ICommand) *MacroCommand {
	return &MacroCommand{commands: commands}
}

func (m *MacroCommand) Execute() {
	fmt.Println("--- [МАКРОС] Выполнение ---")
	for _, cmd := range m.commands {
		cmd.Execute()
	}
	fmt.Println("--- [МАКРОС] Завершен ---")
}

func (m *MacroCommand) Undo() {
	fmt.Println("--- [МАКРОС] Отмена ---")
	for i := len(m.commands) - 1; i >= 0; i-- {
		m.commands[i].Undo()
	}
	fmt.Println("--- [МАКРОС] Отмена завершена ---")
}

type Logger struct{}

func (l *Logger) Log(action string) {
	fmt.Printf("[ЛОГ: %s] Выполнено действие: %s\n", time.Now().Format("15:04:05"), action)
}

type RemoteControl struct {
	onCommands  map[int]ICommand
	offCommands map[int]ICommand
	undoHistory []ICommand
	logger      *Logger
}

func NewRemoteControl(logger *Logger) *RemoteControl {
	noCmd := &NoCommand{}
	on := make(map[int]ICommand)
	off := make(map[int]ICommand)

	for i := 0; i < 5; i++ {
		on[i] = noCmd
		off[i] = noCmd
	}

	return &RemoteControl{
		onCommands:  on,
		offCommands: off,
		undoHistory: make([]ICommand, 0),
		logger:      logger,
	}
}

func (r *RemoteControl) SetCommand(slot int, on ICommand, off ICommand) {
	r.onCommands[slot] = on
	r.offCommands[slot] = off
}

func (r *RemoteControl) PressOnButton(slot int) {
	cmd, ok := r.onCommands[slot]
	if !ok {
		cmd = &NoCommand{}
	}

	fmt.Printf("--- Нажата [ВКЛ] (Слот %d) ---\n", slot)
	cmd.Execute()

	r.undoHistory = append(r.undoHistory, cmd)
	r.logger.Log(fmt.Sprintf("PressOn(Slot %d)", slot))
}

func (r *RemoteControl) PressOffButton(slot int) {
	cmd, ok := r.offCommands[slot]
	if !ok {
		cmd = &NoCommand{}
	}

	fmt.Printf("--- Нажата [ВЫКЛ] (Слот %d) ---\n", slot)
	cmd.Execute()

	r.undoHistory = append(r.undoHistory, cmd)
	r.logger.Log(fmt.Sprintf("PressOff(Slot %d)", slot))
}

func (r *RemoteControl) PressUndoButton() {
	fmt.Println("--- Нажата [ОТМЕНА] ---")

	if len(r.undoHistory) == 0 {
		fmt.Println("Нечего отменять.")
		r.logger.Log("PressUndo (Пусто)")
		return
	}

	lastIndex := len(r.undoHistory) - 1
	lastCommand := r.undoHistory[lastIndex]

	lastCommand.Undo()

	r.undoHistory = r.undoHistory[:lastIndex]
	r.logger.Log("PressUndo (Выполнено)")
}

func main() {
	logger := &Logger{}
	remote := NewRemoteControl(logger)

	livingRoomLight := &Light{Location: "Гостиная"}
	bedroomTV := &Television{Location: "Спальня"}
	kitchenAC := &AirConditioner{Location: "Кухня", Temperature: 20}

	lightOn := NewLightOnCommand(livingRoomLight)
	lightOff := NewLightOffCommand(livingRoomLight)

	tvOn := NewTVOnCommand(bedroomTV)
	tvOff := NewTVOffCommand(bedroomTV)

	acOn := NewACOnCommand(kitchenAC)
	acOff := NewACOffCommand(kitchenAC)
	acTempUp := NewACSetTempCommand(kitchenAC, 24)

	remote.SetCommand(0, lightOn, lightOff)
	remote.SetCommand(1, tvOn, tvOff)
	remote.SetCommand(2, acOn, acOff)
	remote.SetCommand(3, acTempUp, &NoCommand{})

	fmt.Println("===== Тест 1: Свет и Отмена =====")
	remote.PressOnButton(0)
	remote.PressOffButton(0)
	remote.PressUndoButton()

	fmt.Println("\n===== Тест 2: ТВ =====")
	remote.PressOnButton(1)

	fmt.Println("\n===== Тест 3: Кондиционер и Отмена состояния =====")
	remote.PressOnButton(2)
	remote.PressOnButton(3)
	remote.PressUndoButton()
	remote.PressUndoButton()

	fmt.Println("\n===== Тест 4: Макрокоманда (Режим 'Я ухожу') =====")
	macroOff := NewMacroCommand([]ICommand{lightOff, tvOff, acOff})

	remote.SetCommand(4, &NoCommand{}, macroOff)

	fmt.Println("(Включаем все для теста...)")
	remote.PressOnButton(0)
	remote.PressOnButton(1)
	remote.PressOnButton(2)

	remote.PressOffButton(4)

	remote.PressUndoButton()

	fmt.Println("\n===== Тест 5: Обработка ошибок =====")
	remote.PressOnButton(10)
	remote.PressUndoButton()

	fmt.Println("(Очищаем историю...)")
	remote.PressUndoButton()
	remote.PressUndoButton()
	remote.PressUndoButton()
	remote.PressUndoButton()
}