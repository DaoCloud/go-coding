package coding

import (
    "fmt"
)

type Project struct {
    CreatedAt             int64  `json:"created_at"`
    BackendProjectPath    string `json:"backend_project_path"`
    Description           string `json:"description"`
    GitURL                string `json:"git_url"`
    SshURL                string `json:"ssh_url"`
    IsPublic              bool   `json:"is_public"`
    HttpsURL              string `json:"https_url"`
    ID                    int64  `json:"id"`
    Name                  string `json:"name"`
    OwnerID               int64  `json:"owner_id"`
    OwnerUserName         string `json:"owner_user_name"`
    OwnerUserPicture      string `json:"owner_user_picture"`
    OwnerUserHome         string `json:"owner_user_home"`
    ProjectPath           string `json:"project_path"`
    Status                int    `json:"status"`
    Type                  int    `json:"type"`
    UpdatedAt             int64  `json:"updated_at"`
    LastUpdated           int64  `json:"last_updated"`
    ForkCount             int    `json:"fork_count"`
    StarCount             int    `json:"star_count"`
    WatchCount            int    `json:"watch_count"`
    Pin                   bool   `json:"pin"`
    DepotPath             string `json:"depot_path"`
    Forked                bool   `json:"forked"`
    UnReadActivitiesCount int    `json:"un_read_activities_count"`
    Icon                  string `json:"icon"`
    CurrentUserRoleID     int    `json:"current_user_role_id"`
    CurrentUserRole       string `json:"current_user_role"`
    Stared                bool   `json:"stared"`
    Watched               bool   `json:"watched"`
    Recommended           int    `json:"recommended"`
    MaxMember             int    `json:"max_member"`
    GroupID               int    `json:"groupId"`
}

type ProjectService struct {
    client *Client
}

// API:  GET /api/user/{user_name}/project/{project_name}
func (s *ProjectService) GetProject(owner, project string) (*Project, error) {
    url := fmt.Sprintf("/user/%s/project/%s", owner, project)
    url = s.client.addToken(url)

    req, err := s.client.NewRequest("GET",
        url,
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
    url := s.client.addToken("/user/projects")
    req, err := s.client.NewRequest("GET", url, nil)
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
        }   `json:"data"`
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
    url := fmt.Sprintf("/user/%s/project/%s/git/blob/%s/%s", owner, project, branch, path)
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
            File File `json:"file"`
        }   `json:"data"`
    }{}

    _, err = s.client.Do(req, &out)
    return []byte(out.Data.File.Data), err
}
