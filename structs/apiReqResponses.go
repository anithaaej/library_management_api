package structs

type Response struct {
	Status  string      `json:"status" bson:"status"`
	Message string      `json:"message" bson:"message"`
	Data    interface{} `json:"body,omitempty" bson:"body,omitempty"`
}
