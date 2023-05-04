```text
PUT news
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
        "type": "date"
      },
      "insert_time":{
        "type": "date"
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
```