package api

import (
	"fmt"

	"github.com/long2ice/awesome/ent/repo"

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
	PlatformID int    `query:"platform_id"`
}

func (t *TopicSearch) Handler(c *fiber.Ctx) error {
	var topics []*ent.Topic
	var total int64
	if t.Keyword != "" {
		searchReq := &meilisearch.SearchRequest{
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
			AllX(c.Context())
	}

	return c.JSON(fiber.Map{
		"data":  topics,
		"total": total,
	})
}

type RepoSubTopic struct {
	TopicID int    `uri:"topic_id" validate:"required" example:"1"`
	Type    string `                                   example:"repo" query:"type"`
}

func (r *RepoSubTopic) Handler(c *fiber.Ctx) error {
	where := []predicate.Repo{repo.TopicID(r.TopicID)}
	if r.Type != "" {
		where = append(where, repo.TypeEQ(repo.Type(r.Type)))
	}
	topics := db.Client.Repo.Query().
		Where(where...).Unique(true).
		Select(repo.FieldSubTopic).
		StringsX(c.Context())
	return c.JSON(topics)
}
