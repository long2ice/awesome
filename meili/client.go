package meili

import (
	"github.com/long2ice/awesome/conf"
	"github.com/meilisearch/meilisearch-go"
)

var (
	Client     *meilisearch.Client
	TopicIndex *meilisearch.Index
	RepoIndex  *meilisearch.Index
)

func init() {
	Client = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   conf.MeiliConfig.Server,
		APIKey: conf.MeiliConfig.MasterKey,
	})
	TopicIndex = Client.Index("awesome-topic")
	RepoIndex = Client.Index("awesome-repo")
}
