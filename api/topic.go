package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/long2ice/awesome/db"
	"github.com/long2ice/awesome/ent/topic"
	"github.com/long2ice/awesome/meili"
	"github.com/meilisearch/meilisearch-go"
)

type TopicSearch struct {
	Keyword string `query:"keyword" validate:"required" example:"mysql"`
	Limit   int64  `query:"limit"   validate:"required" example:"20"`
	Offset  int64  `query:"offset"                      example:"0"`
}

func (t *TopicSearch) Handler(c *fiber.Ctx) error {
	searchRes, err := meili.TopicIndex.Search(t.Keyword, &meilisearch.SearchRequest{
		Offset:               t.Offset,
		Limit:                t.Limit,
		AttributesToRetrieve: []string{"id"},
	})
	if err != nil {
		return err
	}
	ids := meili.GetIds(searchRes)
	topics := db.Client.Topic.Query().Where(topic.IDIn(ids...)).AllX(c.Context())
	return c.JSON(fiber.Map{
		"data":  topics,
		"total": searchRes.NbHits,
	})
}
