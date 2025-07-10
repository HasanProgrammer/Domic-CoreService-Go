package contracts

type ISerializer interface {
	Serialize(object interface{}) (string, error)
	Deserialize(stringifyPayload string, target interface{}) error
}
