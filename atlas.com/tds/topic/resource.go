package topic

import (
	"atlas-tds/configurations"
	"atlas-tds/json"
	"atlas-tds/rest"
	"github.com/gorilla/mux"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"net/http"
)

const (
	GetTopics = "get_topics"
	GetTopic  = "get_topic"
)

func InitResource(router *mux.Router, l logrus.FieldLogger) {
	r := router.PathPrefix("/topics").Subrouter()
	r.HandleFunc("/", registerGetTopics(l)).Methods(http.MethodGet)
	r.HandleFunc("/{id}", registerGetTopic(l)).Methods(http.MethodGet)
}

type IdHandler func(topicId string) http.HandlerFunc

func ParseId(next IdHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(mux.Vars(r)["id"])(w, r)
	}
}

func registerGetTopics(l logrus.FieldLogger) http.HandlerFunc {
	return rest.RetrieveSpan(GetTopics, handleGetTopics(l))
}

func curriedGetTopics(l logrus.FieldLogger) func(w http.ResponseWriter) {
	return func(w http.ResponseWriter) {
		var response DataListContainer
		response.Data = make([]DataBody, 0)

		c, _ := configurations.GetConfiguration()
		for _, x := range c.Topics {
			var serverData = getTopicResponseObject(x)
			response.Data = append(response.Data, serverData)
		}

		err := json.ToJSON(response, w)
		if err != nil {
			l.WithError(err).Errorf("Error encoding response")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func handleGetTopics(l logrus.FieldLogger) rest.SpanHandler {
	return func(_ opentracing.Span) http.HandlerFunc {
		return func(w http.ResponseWriter, _ *http.Request) {
			curriedGetTopics(l)(w)
		}
	}
}

func registerGetTopic(l logrus.FieldLogger) http.HandlerFunc {
	return rest.RetrieveSpan(GetTopic, handleGetTopic(l))
}

func handleGetTopic(l logrus.FieldLogger) rest.SpanHandler {
	return func(_ opentracing.Span) http.HandlerFunc {
		return ParseId(func(topicId string) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				curriedGetTopic(l)(topicId)(w)
			}
		})
	}
}

func curriedGetTopic(l logrus.FieldLogger) func(topicId string) func(w http.ResponseWriter) {
	return func(topicId string) func(w http.ResponseWriter) {
		return func(w http.ResponseWriter) {
			c, _ := configurations.GetConfiguration()
			tc, _ := c.GetTopicConfiguration(topicId)

			var response DataContainer
			response.Data = getTopicResponseObject(*tc)

			err := json.ToJSON(response, w)
			if err != nil {
				l.WithError(err).Errorf("Error encoding response")
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	}
}

func getTopicResponseObject(c configurations.TopicConfiguration) DataBody {
	return DataBody{
		Id:   c.Id,
		Type: "com.atlas.tds.rest.attribute.TopicAttributes",
		Attributes: Attributes{
			Name: c.Name,
		},
	}
}
