package es

type Filed struct {
	Uuid            string              `json:"uuid"`
	SiteDomain      string              `json:"site_domain"`
	SourceName      string              `json:"source_name"`
	Url             string              `json:"url"`
	Title           string              `json:"title"`
	Author          []string            `json:"author"`
	Content         string              `json:"content"`
	CommentCount    uint                `json:"comment_count"`
	ReadCount       uint                `json:"read_count"`
	LikeCount       uint                `json:"like_count"`
	ForwardCount    uint                `json:"forward_count"`
	NewsType        string              `json:"news_type"`
	Lang            string              `json:"lang"`
	Direction       string              `json:"direction"`
	BoardTheme      string              `json:"board_theme"`
	OriginTags      []string            `json:"origin_tags"`
	SiteBoardName   string              `json:"site_board_name"`
	RepostSource    string              `json:"repost_source"`
	IfRepost        uint                `json:"if_repost"`
	IfFrontPosition uint                `json:"if_front_position"`
	PublishTime     string              `json:"publish_time"`
	InsertTime      string              `json:"insert_time"`
	SiteId          string              `json:"site_id"`
	BoardId         string              `json:"board_id"`
	IndexCon        []map[string]string `json:"index_con"`
}
