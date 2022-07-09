package tasks

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/hibiken/asynq"
	"github.com/long2ice/awesome/conf"
	"github.com/long2ice/awesome/db"
	"github.com/long2ice/awesome/ent"
	"github.com/long2ice/awesome/ent/repo"
	"github.com/long2ice/awesome/meili"
	"github.com/long2ice/awesome/schema"
	"github.com/meilisearch/meilisearch-go"
	log "github.com/sirupsen/logrus"
)

const (
	TypeGetTopicPeriodic  = "GetTopicPeriodic"
	TypeSyncTopicPeriodic = "SyncTopicPeriodic"
	TypeSyncRepoPeriodic  = "SyncRepoPeriodic"
)

func GetTopicPeriodic(ctx context.Context, _ *asynq.Task) error {
	url := "https://awesome.digitalbunker.dev"
	resp, err := Resty.R().Get(url)
	if err != nil {
		return err
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(resp.Body()))
	topics := make(map[string][]schema.Topic)
	doc.Find(".card").Each(func(i int, selection *goquery.Selection) {
		body := selection.Find(".card-body").First()
		platformName := body.Find(".small").Text()
		var subTopic string
		topicName := body.Find(".card-title").Text()
		description := body.Find(".card-text").Text()
		var githubURL, URL string
		body.Find("a").Each(func(i int, selection *goquery.Selection) {
			if i == 0 {
				githubURL, _ = selection.Attr("href")
			} else {
				URL, _ = selection.Attr("href")
				URL = strings.Replace(URL, "https://github.com/", "", -1)
				URL = url + strings.TrimLeft(URL, ".")
			}
		})
		footer := selection.Find(".card-footer").First()
		if footer != nil {
			footer.Find("small").Each(func(i int, selection *goquery.Selection) {
				if i == 2 {
					subTopic = selection.Text()
				}
			})
		}
		if subTopic != "" {
			tmp := subTopic
			subTopic = topicName
			topicName = tmp
		}
		topics[platformName] = append(topics[platformName], schema.Topic{
			Name:        topicName,
			SubName:     subTopic,
			URL:         URL,
			GithubURL:   githubURL,
			Description: description,
		})
	})
	for k, v := range topics {
		pid := db.Client.Platform.Create().SetName(k).OnConflict().Ignore().IDX(ctx)
		for _, t := range v {
			id := db.Client.Topic.Create().
				SetName(t.Name).
				SetSubName(t.SubName).
				SetPlatformID(pid).
				SetURL(t.URL).
				SetGithubURL(t.GithubURL).
				SetDescription(t.Description).OnConflict().UpdateNewValues().
				IDX(ctx)
			var info *asynq.TaskInfo
			info, err = Client.Enqueue(
				asynq.NewTask(TypeGetRepos, []byte(fmt.Sprintf("%d", id))),
			)
			if err != nil {
				return err
			}
			log.Infof("enqueued task: id=%s queue=%s topicID=%d", info.ID, info.Queue, id)
		}
	}

	return err
}
func SyncTopicPeriodic(ctx context.Context, _ *asynq.Task) error {
	topics := db.Client.Topic.Query().AllX(ctx)
	var documents []map[string]interface{}
	for _, tp := range topics {
		documents = append(documents, map[string]interface{}{
			"id":          tp.ID,
			"name":        tp.Name,
			"sub_name":    tp.SubName,
			"description": tp.Description,
			"platform_id": tp.PlatformID,
		})
	}
	if len(documents) > 0 {
		var t *meilisearch.Task
		t, err := meili.TopicIndex.AddDocuments(documents)
		if err != nil {
			return err
		}
		log.Infof("success add %d topics to meilisearch, taskID: %d", len(documents), t.UID)
	}
	return nil
}
func SyncRepoPeriodic(ctx context.Context, _ *asynq.Task) error {
	offset := 0
	for {
		repos := db.Client.Repo.Query().
			Limit(conf.MeiliSearchConfig.BatchSize).
			Offset(offset).
			Order(ent.Asc(repo.FieldID)).
			AllX(ctx)
		if len(repos) == 0 {
			break
		}
		var documents []map[string]interface{}
		for _, r := range repos {
			documents = append(documents, map[string]interface{}{
				"id":          r.ID,
				"name":        r.Name,
				"description": r.Description,
				"sub_topic":   r.SubTopic,
				"topic_id":    r.TopicID,
				"type":        r.Type,
			})
		}
		if len(documents) > 0 {
			var task *meilisearch.Task
			task, err := meili.RepoIndex.AddDocuments(documents)
			if err != nil {
				return err
			}
			log.Infof("success add %d repos to meilisearch, taskID: %d", len(documents), task.UID)
		}
		offset += len(repos)
	}

	return nil
}
