package coding

import (
	"fmt"
)

type Commit struct {
	FullMessage  string     `json:"fullMessage"`  // "fullMessage":"new file abs.client.md",
	ShortMessage string     `json:"shortMessage"` // "shortMessage":"new file abs.client.md\n",
	AllMessage   string     `json:"allMessage"`   // "allMessage":"",
	CommitID     string     `json:"commitId"`     // "commitId":"8a305b5304a0eab0c58ea861dbb1b760a72ea9a8",
	CommitTime   int64      `json:"commitTime"`   // "commitTime":1431057342000,
	Committer    *Committer `json:"committer"`    // "committer":
	NotesCount   int        `json:"notesCount"`   // "notesCount":0
}

type Committer struct {
	Name   string `json:"name"`   //     "name":"ZXX_ABC",
	Email  string `json:"email"`  //     "email":"kcccss111@gmail.com",
	Avatar string `json:"avatar"` //     "avatar":"https://dn-coding-net-production-statis.client.qbox.me/0174e523-963b-4cfb-a2ef-d4f0efea3465.jpg?imageMogr2/auto-orient/format/jpeg/crop/!128x128a0a0",
	Link   string `json:"link"`   //     "link":"/u/zxx_sse"
}

// 从 branch 名称查询 commit ( scope=project, project:depot )
// API: GET  /api/user/{user_name}/project/{project_name}/git/tree/{branch_name}
func (s *ProjectService) GetCommit(owner, project string, branch string) (*Commit, error) {
	req, err := s.client.NewRequest("GET",
		fmt.Sprintf("/user/%s/project/%s/git/tree/%s", owner, project, branch),
		nil)
	if err != nil {
		return nil, err
	}

	out := struct {
		Code int `json:"code"`
		Data struct {
			Ref    string  `json:"ref"`
			Commit *Commit `json:"lastCommit"`
		} `json:"data"`
	}{}

	_, err = s.client.Do(req, &out)
	return out.Data.Commit, err
}
