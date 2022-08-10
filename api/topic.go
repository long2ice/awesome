package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/long2ice/awesome/db"
	"github.com/long2ice/awesome/ent"
	"github.com/long2ice/awesome/ent/predicate"
	"github.com/long2ice/awesome/ent/repo"
	"github.com/long2ice/awesome/ent/topic"
	"github.com/long2ice/awesome/meili"
	"github.com/meilisearch/meilisearch-go"
)

type TopicSearchReq struct {
	Keyword    string `query:"keyword"     example:"mysql"`
	PlatformID int    `query:"platform_id"`
}

func TopicSearch(c *fiber.Ctx, req TopicSearchReq) error {
	var topics []*ent.Topic
	var total int64
	if req.Keyword != "" {
		searchReq := &meilisearch.SearchRequest{
			AttributesToRetrieve: []string{"id"},
		}
		if req.PlatformID != 0 {
			searchReq.Filter = fmt.Sprintf("platform_id = %d", req.PlatformID)
		}
		searchRes, err := meili.TopicIndex.Search(req.Keyword, searchReq)
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
		total = searchRes.EstimatedTotalHits
	} else {
		var where []predicate.Topic
		if req.PlatformID != 0 {
			where = append(where, topic.PlatformID(req.PlatformID))
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

type RepoSubTopicReq struct {
	TopicID int    `uri:"topic_id" validate:"required" example:"1"`
	Type    string `                                   example:"repo" query:"type"`
}

func RepoSubTopic(c *fiber.Ctx, req RepoSubTopicReq) error {
	where := []predicate.Repo{repo.TopicID(req.TopicID)}
	if req.Type != "" {
		where = append(where, repo.TypeEQ(repo.Type(req.Type)))
	}
	topics := db.Client.Repo.Query().
		Where(where...).Unique(true).
		Select(repo.FieldSubTopic).
		StringsX(c.Context())
	return c.JSON(topics)
}
