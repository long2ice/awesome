package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/long2ice/awesome/db"
	"github.com/long2ice/awesome/ent/repo"
	"github.com/long2ice/awesome/meili"
	"github.com/meilisearch/meilisearch-go"
)

type Repo struct {
	Keyword string `query:"keyword"  validate:"required" example:"mysql"`
	Limit   int64  `query:"limit"    validate:"required" example:"20"`
	Offset  int64  `query:"offset"                       example:"0"`
	TopicID int    `query:"topic_id"`
}

func (p *Repo) Handler(c *fiber.Ctx) error {
	searchReq := &meilisearch.SearchRequest{
		Offset:               p.Offset,
		Limit:                p.Limit,
		AttributesToRetrieve: []string{"id"},
	}
	if p.TopicID != 0 {
		searchReq.Filter = fmt.Sprintf("topic_id = %d", p.TopicID)
	}
	searchRes, err := meili.RepoIndex.Search(p.Keyword, searchReq)
	if err != nil {
		return err
	}
	ids := meili.GetIds(searchRes)
	repos := db.Client.Repo.Query().
		Select(
			repo.FieldID,
			repo.FieldName,
			repo.FieldDescription,
			repo.FieldURL,
			repo.FieldSubTopic,
			repo.FieldType,
			repo.FieldStarCount,
			repo.FieldWatchCount,
			repo.FieldForkCount,
			repo.FieldUpdatedAt).
		Where(repo.IDIn(ids...)).
		AllX(c.Context())
	return c.JSON(fiber.Map{
		"data":  repos,
		"total": searchRes.NbHits,
	})
}
