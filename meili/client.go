package meili

import (
	"github.com/long2ice/awesome/conf"
	"github.com/meilisearch/meilisearch-go"
	log "github.com/sirupsen/logrus"
)

var (
	Client     *meilisearch.Client
	TopicIndex *meilisearch.Index
	RepoIndex  *meilisearch.Index
)

func init() {
	Client = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   conf.MeiliSearchConfig.Server,
		APIKey: conf.MeiliSearchConfig.MasterKey,
	})
	TopicIndex = Client.Index("awesome-topic")
	_, err := TopicIndex.UpdateFilterableAttributes(&[]string{"platform_id"})
	if err != nil {
		log.Panic(err)
	}
	RepoIndex = Client.Index("awesome-repo")
	_, err = RepoIndex.UpdateFilterableAttributes(&[]string{"topic_id", "type", "sub_topic"})
	if err != nil {
		log.Panic(err)
	}
}
func GetIds(searchRes *meilisearch.SearchResponse) []int {
	var ids []int
	for _, item := range searchRes.Hits {
		id := item.(map[string]interface{})["id"].(float64)
		ids = append(ids, int(id))
	}
	return ids
}
