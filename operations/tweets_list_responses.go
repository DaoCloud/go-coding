package operations

// 本代码应该由go-swagger生成，
// 但可能由于go-swagger还没有开发完成，所以部分代码无法成功生成，
// 因此这部分代码为根据go-swagger工程中的示例脚本生成的代码结构为规范，手写完成的代码！

import (
	"github.com/wangfeiping/go-coding/models"
)

type TweetsListResponseBody struct {
	Code int
	Data []*models.TweetDTO
}


