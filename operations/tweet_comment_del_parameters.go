package operations

// 本代码应该由go-swagger生成，
// 但可能由于go-swagger还没有开发完成，所以部分代码无法成功生成，
// 因此这部分代码为根据go-swagger工程中的示例脚本生成的代码结构为规范，手写完成的代码！

import (
	"fmt"
	
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
TweetCommentDelParams contains all the parameters to send to the API endpoint
for the get user by name operation typically these are written to a http.Request
*/
type TweetCommentDelParams struct {
	TweetId int32 //tweetId
	CommentId int32 //commentId
}

// WriteToRequest writes these params to a swagger request
// API: DELETE /api/social/tweet/{id}/comment/{comment_id}
func (o *TweetCommentDelParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {
	// path param name
	if err := r.SetPathParam("social/tweet",
			fmt.Sprintf("%d", o.TweetId)); err != nil {
		return err
	}
	if err := r.SetPathParam("comment",
			fmt.Sprintf("%d", o.CommentId)); err != nil {
		return err
	}
	if err := r.SetFormParam("comment",
			fmt.Sprintf("%d", o.CommentId)); err != nil {
		return err
	}
	return nil
}
