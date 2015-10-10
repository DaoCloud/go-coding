package coding

import (
	"fmt"
	"strings"

	"github.com/google/go-querystring/query"
)

type Key struct {
	Title       string `url:"title"     json:"title"`
	Content     string `url:"content"   json:"content"`
	ID          int    `url:"-"         json:"id"`           // "id":2781,
	OwnerID     int    `url:"-"         json:"owner_id"`     // "owner_id":31828,
	FingerPrint string `url:"-"         json:"finger_print"` // "finger_print":"ef:03:d4:9d:d3:7d:bb:6d:ac:c3:99:9a:10:a9:32:4d",
	Type        string `url:"-"         json:"type"`         // "type":"deploy",
	CreatedAt   int64  `url:"-"         json:"created_at"`   // "created_at":1431509572347
}

// 生成部署公钥 ( scope=project, project:depot )
// API: POST /api/user/{user_name}/project/{project_name}/git/deploy_key
func (s *ProjectService) CreateDeployKey(owner, project string, key *Key) error {
	v, err := query.Values(key)
	if err != nil {
		return err
	}

	req, err := s.client.NewRequest("POST",
		fmt.Sprintf("/user/%s/project/%s/git/deploy_key", owner, project),
		strings.NewReader(v.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	_, err = s.client.Do(req, nil)
	return err
}

// 删除部署公钥 ( scope=project, project:depot )
// API: DELETE /api/user/{user_name}/key/{id}
func (s *ProjectService) DeleteKey(owner string, keyID int) error {
	req, err := s.client.NewRequest("DELETE",
		fmt.Sprintf("/user/%s/key/%d", owner, keyID),
		nil)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req, nil)
	return err
}

// API: GET /user/{user_name}/project/{project_name}/git/deploy_keys
func (s *ProjectService) ListDeployKeys(owner, project string) ([]*Key, error) {
	out := struct {
		Code int    `json:"code"`
		Data []*Key `json:"data"`
	}{}

	req, err := s.client.NewRequest("DELETE",
		fmt.Sprintf("/user/%s/project/%s/git/deploy_keys", owner, project),
		nil)
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &out.Data)
	return out.Data, err
}
