package gitlab

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type ObjectAttributes struct {
	Id          int
	Status      string
	Created_at  GitlabTime
	Finished_at GitlabTime
	Duration    int
	Url         string
	Ref         string
}

type Project struct {
	Name      string
	Namespace string
	Web_url   string
}

type Build struct {
	Stage           string
	Name            string
	Status          string
	Duration        float32
	Queued_duration float32
}

type GitlabRequest struct {
	Object_kind       string
	Object_attributes ObjectAttributes
	Project           Project
	Builds            []Build
}

type GitlabTime time.Time

func (j *GitlabTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02 15:04:05 -0700", s)
	if err != nil {
		return err
	}
	*j = GitlabTime(t)
	return nil
}

func (j GitlabTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Maybe a Format function for printing your date
func (j GitlabTime) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (j GitlabTime) Sec() int64 {
	return time.Time(j).Unix()
}

func ParseGitlabHook(req *http.Request) (*GitlabRequest, error) {
	var r GitlabRequest

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}
