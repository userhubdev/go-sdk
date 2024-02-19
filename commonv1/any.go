package commonv1

// Any contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type Any struct {
	// The type of the serialized message.
	ObjectType string `json:"@type"`
}
