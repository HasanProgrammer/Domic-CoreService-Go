package contracts

type IInternalMessageBroker interface {
	Subscribe(queue string, closure func(body []byte) error) error
	Publish(event interface{}, exchange string, routeKey *string) error
}
