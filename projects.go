package coding

import (
	"fmt"
)

type Project struct {
	CreatedAt             int64  `json:"created_at"`               // "created_at": 1426246044000
	BackendProjectPath    string `json:"backend_project_path"`     // "backend_project_path": "/user/baoti/project/Coding-API",
	Description           string `json:"description"`              // "description": "Coding 的 API 是啥样呢？瞧一瞧，看一看。\n注意：由于官方未提供 API 文档，此处 API 整理自 Coding-Android 项目源码。\n尚未对照 IOS 客户端源码。请前往演示, 或转至 http://coding-api.coding.io/ 以查看当前整理的 API.",
	GitURL                string `json:"git_url"`                  // "git_url": "git://coding.net/baoti/Coding-API.git",
	SshURL                string `json:"ssh_url"`                  // "ssh_url": "git@coding.net:baoti/Coding-API.git",
	IsPublic              bool   `json:"is_public"`                // "is_public": true,
	HttpsURL              string `json:"https_url"`                // "https_url": "https://coding.net/baoti/Coding-API.git",
	ID                    int64  `json:"id"`                       // "id": 67965,
	Name                  string `json:"name"`                     // "name": "Coding-API",
	OwnerID               int64  `json:"owner_id"`                 // "owner_id": 84337,
	OwnerUserName         string `json:"owner_user_name"`          // "owner_user_name": "baoti",
	OwnerUserPicture      string `json:"owner_user_picture"`       // "owner_user_picture": "/static/fruit_avatar/Fruit-2.png",
	OwnerUserHome         string `json:"owner_user_home"`          // "owner_user_home": "<a href=\"https://coding.net/u/baoti\">baoti</a>",
	ProjectPath           string `json:"project_path"`             // "project_path": "/u/baoti/p/Coding-API",
	Status                int    `json:"status"`                   // "status": 1,
	Type                  int    `json:"type"`                     // "type": 1,
	UpdatedAt             int64  `json:"updated_at"`               // "updated_at": 1426576642885,
	LastUpdated           int64  `json:"last_updated"`             // "last_updated": 1426576630625,
	ForkCount             int    `json:"fork_count"`               // "fork_count": 0,
	StarCount             int    `json:"star_count"`               // "star_count": 2,
	WatchCount            int    `json:"watch_count"`              // "watch_count": 5,
	Pin                   bool   `json:"pin"`                      // "pin": false,
	DepotPath             string `json:"depot_path"`               // "depot_path": "/u/baoti/p/Coding-API/git",
	Forked                bool   `json:"forked"`                   // "forked": false,
	UnReadActivitiesCount int    `json:"un_read_activities_count"` // "un_read_activities_count": 0,
	Icon                  string `json:"icon"`                     // "icon": "/static/project_icon/scenery-23.png",
	CurrentUserRoleID     int    `json:"current_user_role_id"`     // "current_user_role_id": 100,
	CurrentUserRole       string `json:"current_user_role"`        // "current_user_role": "owner",
	Stared                bool   `json:"stared"`                   // "stared": false,
	Watched               bool   `json:"watched"`                  // "watched": false,
	Recommended           int    `json:"recommended"`              // "recommended": 1,
	MaxMember             int    `json:"max_member"`               // "max_member": 10,
	GroupID               int    `json:"groupId"`                  // "groupId": 0
}

type ProjectService struct {
	client *Client
}

// API:  GET /api/user/{user_name}/project/{project_name}
func (s *ProjectService) GetProject(owner, project string) (*Project, error) {
	req, err := s.client.NewRequest("GET",
		fmt.Sprintf("/user/%s/project/%s", owner, project),
		nil)
	if err != nil {
		return nil, err
	}

	out := struct {
		Code    int      `json:"code"`
		Project *Project `json:"data"`
	}{}

	_, err = s.client.Do(req, &out)
	return out.Project, err
}

// 用户的项目列表 ( scope=project )
// API: GET /api/user/projects
func (s *ProjectService) GetUserProjects() ([]*Project, error) {
	req, err := s.client.NewRequest("GET", "/user/projects", nil)
	if err != nil {
		return nil, err
	}
	out := struct {
		Code int `json:"code"`
		Data struct {
			Page      int        `json:"page"`
			PageSize  int        `json:"pageSize"`
			TotalPage int        `json:"totalPage"`
			TotalRow  int        `json:"totalRow"`
			Projects  []*Project `json:"list"`
		} `json:"data"`
	}{}

	_, err = s.client.Do(req, &out)
	return out.Data.Projects, err
}

type File struct {
	Data string `json:"data"`
	Path string `json:"path"`
	Name string `json:"name"`
}

// 读取 repo 某个文件 ( scope=project, project:depot )
// API: GET /api/user/{user_name}/project/{project_name}/git/bolb/{branch_name}/{file_path}
func (s *ProjectService) GetProjectRaw(owner, project, path, branch string) ([]byte, error) {
	req, err := s.client.NewRequest("GET",
		fmt.Sprintf("/user/%s/project/%s/git/blob/%s/%s", owner, project, branch, path),
		nil)
	if err != nil {
		return nil, err
	}

	out := struct {
		Code int `json:"code"`
		Data struct {
			File File `json:"file"`
		} `json:"data"`
	}{}

	_, err = s.client.Do(req, &out)
	return []byte(out.Data.File.Data), err
}
