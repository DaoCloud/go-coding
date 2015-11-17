package coding

import (
    "fmt"
)

type Commit struct {
    FullMessage  string     `json:"fullMessage"`
    ShortMessage string     `json:"shortMessage"`
    AllMessage   string     `json:"allMessage"`
    CommitID     string     `json:"commitId"`
    CommitTime   int64      `json:"commitTime"`
    Committer    *Committer `json:"committer"`
    NotesCount   int        `json:"notesCount"`
}

type Committer struct {
    Name   string `json:"name"`
    Email  string `json:"email"`
    Avatar string `json:"avatar"`
    Link   string `json:"link"`
}

// 从 branch 名称查询 commit ( scope=project, project:depot )
// API: GET  /api/user/{user_name}/project/{project_name}/git/tree/{branch_name}
func (s *ProjectService) GetCommit(owner, project string, branch string) (*Commit, error) {
    url := fmt.Sprintf("/user/%s/project/%s/git/tree/%s", owner, project, branch)
    url = s.client.addToken(url)
    req, err := s.client.NewRequest("GET",
        url,
        nil,
    )
    if err != nil {
        return nil, err
    }

    out := struct {
        Code int `json:"code"`
        Data struct {
            Ref    string  `json:"ref"`
            Commit *Commit `json:"lastCommit"`
        }   `json:"data"`
    }{}

    _, err = s.client.Do(req, &out)
    return out.Data.Commit, err
}
