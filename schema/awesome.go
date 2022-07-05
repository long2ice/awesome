package schema

import (
	"strconv"
	"strings"
	"time"
)

type Topic struct {
	Name        string
	SubName     string
	URL         string
	GithubURL   string
	Description string
}
type Time struct {
	*time.Time
}

func (r *Time) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		r.Time = nil
	} else {
		timeStr := strings.Trim(string(b), "\"")
		parse, err := time.Parse("2006-01-02 15:04:05", timeStr)
		if err != nil {
			return err
		}
		r.Time = &parse
	}
	return nil
}

type Repo struct {
	ParentListURL       string `json:"parent_list_url"`
	Type                string `json:"type"`
	Category            string `json:"category"`
	RepoURL             string `json:"repo_url"`
	FullName            string `json:"full_name"`
	Description         string `json:"description"`
	Forks               string `json:"forks"`
	Stars               string `json:"stars"`
	Watchers            string `json:"watchers"`
	UpdatedAt           Time   `json:"updated_at"`
	Title               string `json:"title"`
	ResourceDescription string `json:"resource_description"`
}

func (r Repo) ForkCount() int {
	count, _ := strconv.Atoi(r.Forks)
	return count
}
func (r Repo) StarCount() int {
	count, _ := strconv.Atoi(r.Stars)
	return count
}
func (r Repo) WatchCount() int {
	count, _ := strconv.Atoi(r.Watchers)
	return count
}
