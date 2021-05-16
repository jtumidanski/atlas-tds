// Package classification of Topic Discovery Service API
//
// Documentation for Topic Discovery Service API
//
// Schemes: http
// BasePath: /ms/tds/topics
// Version: 1.0.0
//
// Consumes:
// -application/json
//
// Produces:
// -application/json
// swagger:meta
package handlers

import (
	"atlas-tds/attributes"
	"atlas-tds/configurations"
	"log"
	"net/http"
)

// Topics handler for getting topic information
type Topic struct {
	l *log.Logger
}

func NewTopic(l *log.Logger) *Topic {
	return &Topic{l}
}

// swagger:route GET /topics topics getTopics
// Return a list of topics in the registry
// responses:
//	200: TopicsResponse

// GetTopics handles GET requests
func (t *Topic) GetTopics(rw http.ResponseWriter, _ *http.Request) {
	var response attributes.TopicListDataContainer
	response.Data = make([]attributes.TopicData, 0)

	c, _ := configurations.GetConfiguration()
	for _, x := range c.Topics {
		var serverData = getTopicResponseObject(x)
		response.Data = append(response.Data, serverData)
	}

	err := attributes.ToJSON(response, rw)
	if err != nil {
		t.l.Println("Error encoding GetTopics response")
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

// swagger:route GET /topics/{topicId} topics getTopic
// Return a topic in the registry
// responses:
//	200: TopicResponse

// GetTopic handles GET requests
func (t *Topic) GetTopic(rw http.ResponseWriter, r *http.Request) {
	topicId := readString(r, "topicId")

	c, _ := configurations.GetConfiguration()
	tc, _ := c.GetTopicConfiguration(topicId)

	var response attributes.TopicDataContainer
	response.Data = getTopicResponseObject(*tc)

	err := attributes.ToJSON(response, rw)
	if err != nil {
		t.l.Println("Error encoding GetTopics response")
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

func getTopicResponseObject(c configurations.TopicConfiguration) attributes.TopicData {
	return attributes.TopicData{
		Id:   c.Id,
		Type: "com.atlas.tds.rest.attribute.TopicAttributes",
		Attributes: attributes.TopicAttributes{
			Name: c.Name,
		},
	}
}
