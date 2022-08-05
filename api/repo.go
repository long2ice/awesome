package api

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/long2ice/awesome/db"
	"github.com/long2ice/awesome/ent/repo"
	"github.com/long2ice/awesome/meili"
	"github.com/meilisearch/meilisearch-go"
)

type RepoReq struct {
	Keyword  string `query:"keyword"  example:"mysql"`
	Limit    int64  `query:"limit"    example:"20"    validate:"required"`
	Offset   int64  `query:"offset"   example:"0"`
	TopicID  int    `                                 validate:"required" uri:"topic_id"`
	Type     string `query:"type"     example:"repo"`
	SubTopic string `query:"subtopic"`
}

func Repo(c *fiber.Ctx, req RepoReq) error {
	searchReq := &meilisearch.SearchRequest{
		Offset:               req.Offset,
		Limit:                req.Limit,
		AttributesToRetrieve: []string{"id"},
	}
	var filters []string
	var repoFilters []string
	f := fmt.Sprintf("topic_id = %d", req.TopicID)
	filters = append(filters, f)
	repoFilters = append(repoFilters, f)
	if req.Type != "" {
		filters = append(filters, fmt.Sprintf("type = '%s'", req.Type))
	} else {
		repoFilters = append(repoFilters, "type = 'repo'")
	}
	if req.SubTopic != "" {
		f = fmt.Sprintf("sub_topic = '%s'", req.SubTopic)
		filters = append(filters, f)
		repoFilters = append(repoFilters, f)
	}
	searchReq.Filter = strings.Join(filters, " AND ")
	searchRes, err := meili.RepoIndex.Search(req.Keyword, searchReq)
	if err != nil {
		return err
	}
	var repoSearchRes *meilisearch.SearchResponse
	searchReq.Filter = strings.Join(repoFilters, " AND ")
	repoSearchRes, err = meili.RepoIndex.Search(req.Keyword, searchReq)
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
	var repoTotal, resourceTotal, total int64

	if req.Type == "repo" {
		total = repoSearchRes.NbHits
		repoTotal = searchRes.NbHits
		resourceTotal = total - repoTotal
	} else if req.Type == "" {
		total = searchRes.NbHits
		repoTotal = repoSearchRes.NbHits
		resourceTotal = total - repoTotal
	} else {
		total = repoSearchRes.NbHits
		resourceTotal = searchRes.NbHits
		repoTotal = total - resourceTotal
	}
	return c.JSON(fiber.Map{
		"data":           repos,
		"total":          total,
		"repo_total":     repoTotal,
		"resource_total": resourceTotal,
	})
}
