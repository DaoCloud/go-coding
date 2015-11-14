package coding

import (
	"fmt"
	"io/ioutil"
	"strings"
	
	"github.com/wangfeiping/go-coding/operations"
	"github.com/wangfeiping/go-coding/models"
	"github.com/wangfeiping/go-coding/httputil"
)

type TweetService struct {
	client *Client
}

// 冒泡广场列表
// API: GET   /api/tweet/public_tweets
func (s *TweetService) PublicTweets(params *operations.PublicTweetsParams) ([]*models.TweetDTO, error) {
	apiReq := new(httputil.ApiRequest)
	params.WriteToRequest(apiReq, nil)

	req, err := s.client.NewRequest("GET", apiReq.GetPath(), nil)
	if err != nil {
		return nil, err
	}
//	fmt.Println("req.RequestURI: "+req.RequestURI)
	
	out := &operations.PublicTweetsResponseBody{}
	_, err = s.client.Do(req, &out)
	return out.Data, err
}


// 冒泡列表
// API: GET /api/social/tweet/public_tweets
func (s *TweetService) List(params *operations.TweetsListParams) ([]*models.TweetDTO, error) {
	apiReq := new(httputil.ApiRequest)
	params.WriteToRequest(apiReq, nil)

	req, err := s.client.NewRequest("GET", apiReq.GetPath(), nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("req.RequestURI: "+req.RequestURI)

	out := &operations.TweetsListResponseBody{}
	_, err = s.client.Do(req, &out)
	return out.Data, err
}

// 查询last_id以后的最新冒泡列表
// API: GET /api/social/tweet/last
func (s *TweetService) Last(params *operations.TweetsLastParams) ([]*models.TweetDTO, error) {
	apiReq := new(httputil.ApiRequest)
	params.WriteToRequest(apiReq, nil)

	req, err := s.client.NewRequest("GET", apiReq.GetPath(), nil)
	if err != nil {
		return nil, err
	}
	
	out := &operations.TweetsLastResponseBody{}
	_, err = s.client.Do(req, &out)
	return out.Data, err
}

// 发送冒泡
// API: POST /api/social/tweet
func (s *TweetService) Create(params *operations.TweetCreateParams) (*models.TweetDTO, error) {
	apiReq := new(httputil.ApiRequest)
	params.WriteToRequest(apiReq, nil)

	if err := apiReq.SetFormParam("access_token", s.client.AccessToken); err != nil {
		return nil, err
	}
	
	v := apiReq.GetFormValues()
	formBody := ioutil.NopCloser(strings.NewReader(v.Encode()))
	req, err := s.client.NewRequest("POST", apiReq.GetPath(), formBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	out := &operations.TweetCreateResponseBody{}
	_, err = s.client.Do(req, &out)
	return out.Data, err
}

// 删除冒泡
// API: DELETE /api/social/tweet/{tweet_id}
func (s *TweetService) Delete(params *operations.TweetDeleteParams) (bool, error) {
	apiReq := new(httputil.ApiRequest)
	params.WriteToRequest(apiReq, nil)

	if err := apiReq.SetQueryParam("access_token", s.client.AccessToken); err != nil {
		return false, err
	}
	
	v := apiReq.GetFormValues()
	formBody := ioutil.NopCloser(strings.NewReader(v.Encode()))
	req, err := s.client.NewRequest("DELETE", apiReq.GetPath(), formBody)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	out := &operations.TweetDeleteResponseBody{}
	_, err = s.client.Do(req, &out)
	return out.Data, err
}

// 冒泡评论
// API: POST /api/social/tweet/{id}/comment
func (s *TweetService) Comment(params *operations.TweetCommentParams) (*models.TweetCommentDTO, error) {
	apiReq := new(httputil.ApiRequest)
	params.WriteToRequest(apiReq, nil)

	if err := apiReq.SetFormParam("access_token", s.client.AccessToken); err != nil {
		return nil, err
	}
	
	v := apiReq.GetFormValues()
	formBody := ioutil.NopCloser(strings.NewReader(v.Encode()))
	req, err := s.client.NewRequest("POST", apiReq.GetPath(), formBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	out := &operations.TweetCommentResponseBody{}
	_, err = s.client.Do(req, &out)
	return out.Data, err
}

// 删除冒泡评论
// API: DELETE /api/social/tweet/{id}/comment/{comment_id}
func (s *TweetService) DeleteComment(params *operations.TweetCommentDelParams) (bool, error) {
	apiReq := new(httputil.ApiRequest)
	params.WriteToRequest(apiReq, nil)

	if err := apiReq.SetQueryParam("access_token", s.client.AccessToken); err != nil {
		return false, err
	}
	
	v := apiReq.GetFormValues()
	formBody := ioutil.NopCloser(strings.NewReader(v.Encode()))
	req, err := s.client.NewRequest("DELETE", apiReq.GetPath(), formBody)
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	out := &operations.TweetCommentDelResponseBody{}
	_, err = s.client.Do(req, &out)
	return out.Data, err
}

// 冒泡点赞
// API: POST /api/social/tweet/{tweet_id}/like
func (s *TweetService) Like(params *operations.TweetLikeParams) (bool, error) {
	apiReq := new(httputil.ApiRequest)
	params.WriteToRequest(apiReq, nil)

	if err := apiReq.SetFormParam("access_token", s.client.AccessToken); err != nil {
		return false, err
	}
	
	v := apiReq.GetFormValues()
	formBody := ioutil.NopCloser(strings.NewReader(v.Encode()))
	req, err := s.client.NewRequest("POST", apiReq.GetPath(), formBody)
	if err != nil {
		return false, err
	}
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	out := &operations.TweetLikeResponseBody{}
	_, err = s.client.Do(req, &out)
	return out.Data, err
}

// 冒泡取消点赞
// API: POST /api/social/tweet/{tweet_id}/unlike
func (s *TweetService) Unlike(params *operations.TweetUnlikeParams) (bool, error) {
	apiReq := new(httputil.ApiRequest)
	params.WriteToRequest(apiReq, nil)

	if err := apiReq.SetFormParam("access_token", s.client.AccessToken); err != nil {
		return false, err
	}
	
	v := apiReq.GetFormValues()
	formBody := ioutil.NopCloser(strings.NewReader(v.Encode()))
	req, err := s.client.NewRequest("POST", apiReq.GetPath(), formBody)
	if err != nil {
		return false, err
	}
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	out := &operations.TweetUnlikeResponseBody{}
	_, err = s.client.Do(req, &out)
	return out.Data, err
}
