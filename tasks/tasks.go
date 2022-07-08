package tasks

import (
	"context"
	"encoding/json"
	"regexp"
	"strconv"

	"github.com/hibiken/asynq"
	"github.com/long2ice/awesome/db"
	"github.com/long2ice/awesome/ent/repo"
	"github.com/long2ice/awesome/schema"
	log "github.com/sirupsen/logrus"
)

const (
	TypeGetRepos = "GetRepos"
)

func GetRepos(ctx context.Context, t *asynq.Task) error {
	topicID, err := strconv.Atoi(string(t.Payload()))
	if err != nil {
		return err
	}
	topic := db.Client.Topic.GetX(ctx, topicID)
	resp, err := Resty.R().Get(topic.URL)
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		log.WithField("status_code", resp.StatusCode()).
			WithField("topicID", topicID).
			Panicf("get topic content error")
	}
	compile, err := regexp.Compile(`let originalResponse = (\{.+\}\]\})`)
	if err != nil {
		return err
	}
	res := compile.FindStringSubmatch(string(resp.Body()))
	if len(res) == 0 {
		log.WithField("topicID", topicID).Error("get repos content err")
		return err
	}
	ret := res[1]
	var repos map[string][]schema.Repo
	err = json.Unmarshal([]byte(ret), &repos)
	if err != nil {
		return err
	}
	for _, v := range repos {
		for _, repoInfo := range v {
			r := db.Client.Repo.Query().Where(repo.URL(repoInfo.RepoURL)).FirstX(ctx)
			var desc string
			var name string
			if repoInfo.Type == "resource" {
				desc = repoInfo.ResourceDescription
				name = repoInfo.GetTitle()
			} else {
				desc = repoInfo.Description
				name = repoInfo.FullName
			}
			if r == nil {
				repoCreate := db.Client.Repo.Create().
					SetName(name).
					SetForkCount(repoInfo.ForkCount()).
					SetStarCount(repoInfo.StarCount()).
					SetWatchCount(repoInfo.WatchCount()).
					SetSubTopic(repoInfo.Category).
					SetURL(repoInfo.RepoURL).
					SetType(repo.Type(repoInfo.Type)).
					SetDescription(desc).
					SetTopics(topic)
				if repoInfo.UpdatedAt.Time != nil {
					repoCreate.SetUpdatedAt(*repoInfo.UpdatedAt.Time)
				}
				r, err = repoCreate.Save(ctx)
				if err != nil {
					return err
				}
			} else {
				r.Name = name
				r.ForkCount = repoInfo.ForkCount()
				r.StarCount = repoInfo.StarCount()
				r.WatchCount = repoInfo.WatchCount()
				r.SubTopic = repoInfo.Category
				r.Type = repo.Type(repoInfo.Type)
				r.Description = desc
				r.TopicID = topicID
				if repoInfo.UpdatedAt.Time != nil {
					r.UpdatedAt = *repoInfo.UpdatedAt.Time
				}
				db.Client.Repo.UpdateOne(r)
			}
		}
	}
	return err
}
