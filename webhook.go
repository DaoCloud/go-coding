package coding

import (
    "fmt"
    "strings"

    "github.com/google/go-querystring/query"
)

type Hook struct {
    DepotID   int    `json:"depot_id"`
    HookURL   string `json:"hook_url"`
    Type      int    `json:"type"`       // "type":1,
    Token     string `json:"token"`      // "token":"",
    Status    int    `json:"status"`     // "status":0,
    CreatedAt int64  `json:"created_at"` // "created_at":1431511538877,
    UpdatedAt int64  `json:"updated_at"` // "updated_at":1431511538877,
    SendType  int    `json:"send_type"`  // "send_type":3,
    ID        int    `json:"id"`         // "id":174
}

type HookService struct {
    client *Client
}

// webhook 列表 ( scope=project, project:depot )
// API: GET   /api/user/{user_name}/project/{project_name}/git/hooks
func (s *HookService) ListHooks(owner, project string) ([]*Hook, error) {
    url := fmt.Sprintf("/user/%s/project/%s/git/hooks", owner, project)
    url = s.client.addToken(url)
    req, err := s.client.NewRequest("GET",
        url,
        nil,
    )
    if err != nil {
        return nil, err
    }

    out := struct {
        Code  int     `json:"code"`
        Hooks []*Hook `json:"data"`
    }{}

    _, err = s.client.Do(req, &out)
    return out.Hooks, err
}

//获取 webhook ( scope=project, project:depot )
//API: GET   /api/user/{user_name}/project/{project_name}/git/hook/{hook_id}
func (s *HookService) GetHook(owner, project string, hookID int) (*Hook, error) {
    url := fmt.Sprintf("/user/%s/project/%s/git/hook/%d", owner, project, hookID)
    url = s.client.addToken(url)
    req, err := s.client.NewRequest("GET",
        url,
        nil,
    )
    if err != nil {
        return nil, err
    }

    out := struct {
        Code int   `json:"code"`
        Hook *Hook `json:"data"`
    }{}

    _, err = s.client.Do(req, &out)
    return out.Hook, err
}

type HookForm struct {
    HookURL      string `url:"hook_url"`      // hook_url	string	webhook 链接
    Token        string `url:"token"`         // token	string	自定义 webhook 秘钥
    TypePush     bool   `url:"type_push"`     // type_push	boolean	push代码 通知开关
    TypeMR_PR    bool   `url:"type_mr_pr"`    // type_mr_pr	boolean	MR/PR 通知开关
    TypeTopic    bool   `url:"type_topic"`    // type_topic	boolean	发布讨论 通知开关
    TypeMember   bool   `url:"type_member"`   // type_member	boolean	成员变动 通知开关
    TypeComment  bool   `url:"type_comment"`  // type_comment	boolean	发表评论 通知开关
    TypeDocument bool   `url:"type_document"` // type_document	boolean	文档 通知开关
    TypeWatch    bool   `url:"type_watch"`    // type_watch	boolean	项目被关注 通知开关
    TypeStar     bool   `url:"type_star"`     // type_star	boolean	项目被加星 通知开关
    TypeTask     bool   `url:"type_task"`     // type_task	boolean	项目任务 通知开关
}

// 增加 webhook ( scope=project, project:depot )
// API: POST   /api/user/{user_name}/project/{project_name}/git/hook/{hook_id}
func (s *HookService) CreateHook(owner, project string, hook *HookForm) error {
    v, err := query.Values(hook)
    if err != nil {
        return err
    }

    url := fmt.Sprintf("/user/%s/project/%s/git/hook", owner, project)
    url = s.client.addToken(url)

    req, err := s.client.NewRequest("POST",
        url,
        strings.NewReader(v.Encode()),
    )
    if err != nil {
        return err
    }
    req.Header.Add("content-type", "application/x-www-form-urlencoded")

    _, err = s.client.Do(req, nil)
    return err
}

// 编辑 webhook ( scope=project, project:depot )
// API: PUT   /api/user/{user_name}/project/{project_name}/git/hook/{hook_id}
func (s *HookService) EditHook(owner, project string, hookID int, hook *HookForm) error {
    v, err := query.Values(hook)
    if err != nil {
        return err
    }

    url := fmt.Sprintf("/user/%s/project/%s/git/hook/%d", owner, project, hookID)
    url = s.client.addToken(url)
    req, err := s.client.NewRequest("PUT",
        url,
        strings.NewReader(v.Encode()),
    )
    if err != nil {
        return err
    }
    req.Header.Add("content-type", "application/x-www-form-urlencoded")

    _, err = s.client.Do(req, nil)
    return err
}

// 删除 webhook ( scope=project, project:depot )
// API: DELETE   /api/user/{user_name}/project/{project_name}/git/hook
func (s *HookService) DeleteHook(owner, project string, hookID int) error {
    url := fmt.Sprintf("/user/%s/project/%s/git/hook/%d", owner, project, hookID)
    url = s.client.addToken(url)
    req, err := s.client.NewRequest("DELETE",
        url,
        nil,
    )
    if err != nil {
        return err
    }

    _, err = s.client.Do(req, nil)
    return err
}
