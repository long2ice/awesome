package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/long2ice/awesome/db"
	"github.com/long2ice/awesome/ent/repo"
	"github.com/long2ice/awesome/meili"
	"github.com/meilisearch/meilisearch-go"
)

type Repo struct {
	Keyword string `query:"keyword" validate:"required" example:"mysql"`
	Limit   int64  `query:"limit"   validate:"required" example:"20"`
	Offset  int64  `query:"offset"                      example:"0"`
}

func (p *Repo) Handler(c *fiber.Ctx) error {
	searchRes, err := meili.RepoIndex.Search(p.Keyword, &meilisearch.SearchRequest{
		Offset:               p.Offset,
		Limit:                p.Limit,
		AttributesToRetrieve: []string{"id"},
	})
	if err != nil {
		return err
	}
	ids := meili.GetIds(searchRes)
	repos := db.Client.Repo.Query().Where(repo.IDIn(ids...)).AllX(c.Context())
	return c.JSON(fiber.Map{
		"data":  repos,
		"total": searchRes.NbHits,
	})
}
