package es

func getMapping() string {
	return `
{
  "mappings": {
    "properties": {
      "uuid":{
        "type": "keyword"
      },
      "site_domain":{
        "type": "keyword" 
      },
      "source_name":{
        "type": "keyword"
      },
      "url":{
        "type": "keyword"
      },
      "title":{
        "type": "text"
      },
      "author":{
        "type": "text"
      },
      "content":{
        "type": "keyword"
      },
      "comment_count":{
        "type": "integer"
      },
      "read_count":{
        "type": "integer"
      },
      "like_count":{
        "type": "integer"
      },
      "forward_count":{
        "type": "integer"
      },
      "news_type":{
        "type": "keyword"
      },
      "lang":{
        "type": "keyword"
      },
      "direction":{
        "type": "keyword"
      },
      "board_theme":{
        "type": "keyword"
      },
      "origin_tags":{
        "type": "text"
      },
      "site_board_name":{
        "type": "keyword"
      },
      "repost_source":{
        "type": "keyword"
      },
      "if_repost":{
        "type": "integer"
      },
      "if_front_position":{
        "type": "integer"
      },
      "publish_time":{
        "type": "date",
		"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
      },
      "insert_time":{
        "type": "date",
		"format": "yyyy-MM-dd HH:mm:ss||yyyy-MM-dd||epoch_millis"
      },
      "site_id":{
        "type": "keyword"
      },
      "board_id":{
        "type": "keyword"
      },
      "index_con":{
        "type": "nested",
        "properties": {
          "data":{
            "type":"text"
          }
        }
      }
      
    }
  },
  "settings": {
    "number_of_shards": 2,
    "number_of_replicas": 3
  }
}
`
}

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
