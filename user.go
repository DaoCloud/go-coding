package coding

type User struct {
    Avatar           string   `json:"avatar,omitempty"`
    Birthday         string   `json:"birthday,omitempty"`
    Company          string   `json:"company,omitempty"`
    CreatedAt        int64    `json:"created_at,omitempty"` //strfmt.DateTime `json:"created_at,omitempty"`
    Email            string   `json:"email,omitempty"`
    FansCount        int64    `json:"fans_count,omitempty"`
    Follow           bool     `json:"follow,omitempty"`
    Followed         bool     `json:"followed,omitempty"`
    FollowsCount     int64    `json:"follows_count,omitempty"`
    GlobalKey        string   `json:"global_key,omitempty"`
    Gravatar         string   `json:"gravatar,omitempty"`
    ID               int32    `json:"id,omitempty"`
    Introduction     string   `json:"introduction,omitempty"`
    IsMember         int32    `json:"is_member,omitempty"`
    IsPhoneValidated bool     `json:"is_phone_validated,omitempty"`
    Job              int32    `json:"job,omitempty"`
    JobStr           string   `json:"job_str,omitempty"`
    LastActivityAt   int64    `json:"last_activity_at,omitempty"` //strfmt.DateTime `json:"last_activity_at,omitempty"`
    LastLoginedAt    int64    `json:"last_logined_at,omitempty"`  //strfmt.DateTime `json:"last_logined_at,omitempty"`
    Lavatar          string   `json:"lavatar,omitempty"`
    Location         string   `json:"location,omitempty"`
    Name             string   `json:"name,omitempty"`
    NamePinyin       string   `json:"name_pinyin,omitempty"`
    Path             string   `json:"path,omitempty"`
    Phone            string   `json:"phone,omitempty"`
    Sex              int32    `json:"sex,omitempty"`
    Slogan           string   `json:"slogan,omitempty"`
    Status           int32    `json:"status,omitempty"`
    Tags             string   `json:"tags,omitempty"`
    TagsStr          string   `json:"tags_str,omitempty"`
    Test             []string `json:"test,omitempty"`
    TweetsCount      int64    `json:"tweets_count,omitempty"`
    UpdatedAt        int64    `json:"updated_at,omitempty"` //strfmt.DateTime `json:"updated_at,omitempty"`
}

type LikeUser struct {
    Path      string
    GlobalKey string
    Name      string
    Avatar    string
    Follow    int32
    Followed  int32
}
