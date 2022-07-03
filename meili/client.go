package meili

import (
	"github.com/long2ice/awesome/conf"
	"github.com/meilisearch/meilisearch-go"
)

var (
	Client *meilisearch.Client
	Index  *meilisearch.Index
)

func init() {
	Client = meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   conf.MeiliConfig.Server,
		APIKey: conf.MeiliConfig.MasterKey,
	})
	Index = Client.Index("awesome")
}
