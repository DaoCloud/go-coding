package operations

// 本代码应该由go-swagger生成，
// 但可能由于go-swagger还没有开发完成，所以部分代码无法成功生成，
// 因此这部分代码为根据go-swagger工程中的示例脚本生成的代码结构为规范，手写完成的代码！

import (
	"fmt"
	"strconv"
	
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
TweetsListParams contains all the parameters to send to the API endpoint
for the get user by name operation typically these are written to a http.Request
*/
type TweetsListParams struct {
	Sort string
	LastId int32
	DefaultLikeCount int
	Filter bool
}

// WriteToRequest writes these params to a swagger request
// API: GET /api/social/tweet/public_tweets
func (o *TweetsListParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {
	// path param name
	if err := r.SetPathParam("social/tweet", "public_tweets"); err != nil {
		return err
	}
	if err := r.SetQueryParam("sort", o.Sort); err != nil {
		return err
	}
	if o.LastId>0 {
		if err := r.SetQueryParam("last_id",
				fmt.Sprintf("%d", o.LastId)); err != nil {
			return err
		}
	}
	if o.DefaultLikeCount>0 {
		if err := r.SetQueryParam("default_like_count",
			fmt.Sprintf("%d", o.DefaultLikeCount)); err != nil {
			return err
		}
	}
	if err := r.SetQueryParam("filter",
			strconv.FormatBool(o.Filter)); err != nil {
		return err
	}
	return nil
}
