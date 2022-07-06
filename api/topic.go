package api

import (
	"fmt"

	"github.com/long2ice/awesome/ent/predicate"

	"github.com/long2ice/awesome/ent"

	"github.com/gofiber/fiber/v2"
	"github.com/long2ice/awesome/db"
	"github.com/long2ice/awesome/ent/topic"
	"github.com/long2ice/awesome/meili"
	"github.com/meilisearch/meilisearch-go"
)

type TopicSearch struct {
	Keyword    string `query:"keyword"     example:"mysql"`
	Limit      int64  `query:"limit"       example:"20"    validate:"required,max=100,min=10"`
	Offset     int64  `query:"offset"      example:"0"`
	PlatformID int    `query:"platform_id"`
}

func (t *TopicSearch) Handler(c *fiber.Ctx) error {
	var topics []*ent.Topic
	var total int64
	if t.Keyword != "" {
		searchReq := &meilisearch.SearchRequest{
			Offset:               t.Offset,
			Limit:                t.Limit,
			AttributesToRetrieve: []string{"id"},
		}
		if t.PlatformID != 0 {
			searchReq.Filter = fmt.Sprintf("platform_id = %d", t.PlatformID)
		}
		searchRes, err := meili.TopicIndex.Search(t.Keyword, searchReq)
		if err != nil {
			return err
		}
		ids := meili.GetIds(searchRes)
		topics = db.Client.Topic.Query().
			Select(
				topic.FieldID,
				topic.FieldName,
				topic.FieldSubName,
				topic.FieldDescription,
				topic.FieldGithubURL).
			Where(topic.IDIn(ids...)).
			AllX(c.Context())
		total = searchRes.NbHits
	} else {
		var where []predicate.Topic
		if t.PlatformID != 0 {
			where = append(where, topic.PlatformID(t.PlatformID))
		}
		topics = db.Client.Topic.Query().
			Select(
				topic.FieldID,
				topic.FieldName,
				topic.FieldSubName,
				topic.FieldDescription,
				topic.FieldGithubURL).Where(where...).
			Limit(int(t.Limit)).Offset(int(t.Offset)).
			AllX(c.Context())
	}

	return c.JSON(fiber.Map{
		"data":  topics,
		"total": total,
	})
}
