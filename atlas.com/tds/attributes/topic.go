package attributes

type InputTopic struct {
	Data TopicData `json:"data"`
}

// A list of Topics
// swagger:response TopicsResponse
type TopicListDataContainer struct {
	// All topics
	// in: body
	Data []TopicData `json:"data"`
}

// A Topic
// swagger:response TopicResponse
type TopicDataContainer struct {
	// A topic
	// in: body
	Data TopicData `json:"data"`
}

type TopicData struct {
	Id         string          `json:"id"`
	Type       string          `json:"type"`
	Attributes TopicAttributes `json:"attributes"`
}

type TopicAttributes struct {
	Name   string   `json:"name"`
}
