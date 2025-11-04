package main

import (
	"fmt"
)

type IInternalDeliveryService interface {
	DeliverOrder(orderId string)
	GetDeliveryStatus(orderId string) string
}

type InternalDeliveryService struct{}

func (s *InternalDeliveryService) DeliverOrder(orderId string) {
	fmt.Println("Внутренняя доставка: доставка заказа", orderId)
}

func (s *InternalDeliveryService) GetDeliveryStatus(orderId string) string {
	return "Статус внутренней доставки: доставлено"
}

type ExternalLogisticsServiceA struct{}
func (a *ExternalLogisticsServiceA) ShipItem(itemId int) {
	fmt.Println("Внешний A: отправка товара", itemId)
}
func (a *ExternalLogisticsServiceA) TrackShipment(shipmentId int) string {
	return "Отслеживание внешнего A: в пути"
}

type ExternalLogisticsServiceB struct{}
func (b *ExternalLogisticsServiceB) SendPackage(info string) {
	fmt.Println("Внешний B: отправка пакета", info)
}
func (b *ExternalLogisticsServiceB) CheckPackageStatus(code string) string {
	return "Внешний статус B: доставлено"
}

type ExternalLogisticsServiceC struct{}
func (c *ExternalLogisticsServiceC) StartDelivery(address string, weight float64) {
	fmt.Printf("Внешний C: доставка в %s (%.2f kg)\n", address, weight)
}
func (c *ExternalLogisticsServiceC) StatusUpdate() string {
	return "Внешний C: ожидание получения"
}

type LogisticsAdapterA struct {
	service *ExternalLogisticsServiceA
}

func (a *LogisticsAdapterA) DeliverOrder(orderId string) {
	a.service.ShipItem(1001)
}

func (a *LogisticsAdapterA) GetDeliveryStatus(orderId string) string {
	return a.service.TrackShipment(1001)
}

type LogisticsAdapterB struct {
	service *ExternalLogisticsServiceB
}

func (b *LogisticsAdapterB) DeliverOrder(orderId string) {
	b.service.SendPackage(orderId)
}

func (b *LogisticsAdapterB) GetDeliveryStatus(orderId string) string {
	return b.service.CheckPackageStatus(orderId)
}

type LogisticsAdapterC struct {
	service *ExternalLogisticsServiceC
}

func (c *LogisticsAdapterC) DeliverOrder(orderId string) {
	c.service.StartDelivery("Алматы, склад №2", 15.7)
}

func (c *LogisticsAdapterC) GetDeliveryStatus(orderId string) string {
	return c.service.StatusUpdate()
}

type DeliveryServiceFactory struct{}

func (f *DeliveryServiceFactory) GetService(serviceType string) IInternalDeliveryService {
	switch serviceType {
	case "internal":
		return &InternalDeliveryService{}
	case "A":
		return &LogisticsAdapterA{&ExternalLogisticsServiceA{}}
	case "B":
		return &LogisticsAdapterB{&ExternalLogisticsServiceB{}}
	case "C":
		return &LogisticsAdapterC{&ExternalLogisticsServiceC{}}
	default:
		fmt.Println("Неизвестный тип службы доставки:", serviceType)
		return nil
	}
}

func main() {
	factory := &DeliveryServiceFactory{}

	services := []string{"internal", "A", "B", "C"}

	for _, s := range services {
		fmt.Println("\n--- Using service:", s, "---")
		service := factory.GetService(s)
		if service == nil {
			continue
		}
		service.DeliverOrder("Заказ_001")
		fmt.Println(service.GetDeliveryStatus("Заказ_001"))
	}
}