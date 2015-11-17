package coding

import (
    "fmt"
    "strings"

    "github.com/google/go-querystring/query"
)

type TweetService struct {
    client *Client
}

type Tweet struct {
    ActivityID  int            `json:"activity_id,omitempty"`
    Address     string         `json:"address,omitempty"url:"address,omitempty"`
    CommentList []TweetComment `json:"comment_list,omitempty"`
    Comments    int            `json:"comments,omitempty"`
    Content     string         `json:"content,omitempty"url:"content,omitempty"`
    Coord       string         `json:"coord,omitempty"url:"coord,omitempty"`
    CreatedAt   int64          `json:"created_at,omitempty"`
    Device      string         `json:"device,omitempty"url:"device,omitempty"`
    ID          int            `json:"id,omitempty"`
    LikeUsers   []LikeUser     `json:"like_users,omitempty"`
    Liked       bool           `json:"liked,omitempty"`
    Likes       int            `json:"likes,omitempty"`
    Location    string         `json:"location,omitempty"url:"location,omitempty"`
    Owner       User           `json:"owner,omitempty"`
    OwnerID     int            `json:"owner_id,omitempty"`
    Path        string         `json:"path,omitempty"`
}

type ListOpt struct {
    Sort             string `url:"sort,omitempty"`
    LastId           int    `url:"last_id,omitempty"`
    DefaultLikeCount int    `url:"default_like_count,omitempty"`
    Filter           bool   `url:"filter,omitempty"`
}

func (l *ListOpt) Encode() string {
    v, err := query.Values(l)
    if err != nil {
        return ""
    }

    return v.Encode()
}

// 冒泡广场列表
// API: GET   /api/tweet/public_tweets
func (s *TweetService) PublicTweets(opt *ListOpt) ([]*Tweet, error) {

    url := "/api/tweet/public_tweets?" + opt.Encode()

    req, err := s.client.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    out := struct {
        Code   int      `json:"code"`
        Tweets []*Tweet `json:"data"`
    }{}

    _, err = s.client.Do(req, &out)
    return out.Tweets, err
}

// 冒泡列表
// API: GET /api/social/tweet/public_tweets
func (s *TweetService) List(opt *ListOpt) ([]*Tweet, error) {
    url := "/api/social/tweet/public_tweets?" + opt.Encode()

    req, err := s.client.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    out := struct {
        Code   int      `json:"code"`
        Tweets []*Tweet `json:"data"`
    }{}
    _, err = s.client.Do(req, &out)
    return out.Tweets, err
}

// 查询last_id以后的最新冒泡列表
// API: GET /api/social/tweet/last
func (s *TweetService) Last(opt *ListOpt) ([]*Tweet, error) {

    url := fmt.Sprintf("/api/social/tweet/last?" + opt.Encode())
    req, err := s.client.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    out := struct {
        Code   int      `json:"code"`
        Tweets []*Tweet `json:"data"`
    }{}

    _, err = s.client.Do(req, &out)
    return out.Tweets, err
}

// 发送冒泡
// API: POST /api/social/tweet
func (s *TweetService) Create(tweet *Tweet) (*Tweet, error) {
    v, err := query.Values(tweet)
    if err != nil {
        return nil, err
    }

    url := fmt.Sprintf("/api/social/tweet")
    url = s.client.addToken(url)
    req, err := s.client.NewRequest("POST", url, strings.NewReader(v.Encode()))
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    out := struct {
        Code  int    `json:"code"`
        Tweet *Tweet `json:"data"`
    }{}
    _, err = s.client.Do(req, &out)
    return out.Tweet, err
}

// 删除冒泡
// API: DELETE /api/social/tweet/{tweet_id}
func (s *TweetService) Delete(id int) (bool, error) {

    url := fmt.Sprintf("/api/social/tweet/%d", id)
    url = s.client.addToken(url)

    req, err := s.client.NewRequest("DELETE", url, nil)
    if err != nil {
        return false, err
    }

    out := struct {
        Code    int  `json:"code"`
        Deleted bool `json:"data"`
    }{}

    _, err = s.client.Do(req, &out)
    return out.Deleted, err
}

type TweetComment struct {
    Content   string `json:"content,omitempty"`
    CreatedAt int64  `json:"created_at,omitempty"` //Timestamp `json:"created_at,omitempty"`
    ID        int32  `json:"id,omitempty"`
    Owner     User   `json:"owner,omitempty"`
    OwnerID   int32  `json:"owner_id,omitempty"`
    TweetID   int32  `json:"tweet_id,omitempty"`
}

// 冒泡评论
// API: POST /api/social/tweet/{id}/comment
func (s *TweetService) Comment(id int, content string, afterFilter string) (*TweetComment, error) {

    url := fmt.Sprintf("/api/social/tweet/%d/comment", id)
    url = s.client.addToken(url)
    req, err := s.client.NewRequest("POST", url, nil)
    if err != nil {
        return nil, err
    }

    out := struct {
        Code         int           `json:"code"`
        TweetComment *TweetComment `json:"data"`
    }{}

    _, err = s.client.Do(req, &out)
    return out.TweetComment, err
}

// 删除冒泡评论
// API: DELETE /api/social/tweet/{id}/comment/{comment_id}
func (s *TweetService) DeleteComment(tweetId, commmentId int) (bool, error) {

    url := fmt.Sprintf("/api/social/tweet/%d/comment/%d", tweetId, commmentId)
    url = s.client.addToken(url)
    req, err := s.client.NewRequest("DELETE", url, nil)
    if err != nil {
        return false, err
    }

    out := struct {
        Code    int  `json:"code"`
        Deleted bool `json:"data"`
    }{}

    _, err = s.client.Do(req, &out)
    return out.Deleted, err
}

// 冒泡点赞
// API: POST /api/social/tweet/{tweet_id}/like
func (s *TweetService) Like(id int) (bool, error) {

    url := fmt.Sprintf("/api/social/tweet/%d/like", id)
    url = s.client.addToken(url)
    req, err := s.client.NewRequest("POST", url, nil)
    if err != nil {
        return false, err
    }

    out := struct {
        Code int  `json:"code"`
        Like bool `json:"data"`
    }{}
    _, err = s.client.Do(req, &out)
    return out.Like, err
}

// 冒泡取消点赞
// API: POST /api/social/tweet/{tweet_id}/unlike
func (s *TweetService) Unlike(id int) (bool, error) {

    url := fmt.Sprintf("/api/social/tweet/%d/unlike", id)
    url = s.client.addToken(url)
    req, err := s.client.NewRequest("POST", url, nil)
    if err != nil {
        return false, err
    }

    out := struct {
        Code   int  `json:"code"`
        Unlike bool `json:"data"`
    }{}
    _, err = s.client.Do(req, &out)
    return out.Unlike, err
}
