package main

type twitter struct {
  publicStream []message
}

// message publicStream filter response message
type message struct {
  Contributors interface{} `json:"contributors"`
  Coordinates  interface{} `json:"coordinates"`
  CreatedAt    string      `json:"created_at"`
  Entities     struct {
    Hashtags     []interface{} `json:"hashtags"`
    Symbols      []interface{} `json:"symbols"`
    Urls         []interface{} `json:"urls"`
    UserMentions []interface{} `json:"user_mentions"`
  } `json:"entities"`
  FavoriteCount        int         `json:"favorite_count"`
  Favorited            bool        `json:"favorited"`
  FilterLevel          string      `json:"filter_level"`
  Geo                  interface{} `json:"geo"`
  ID                   int         `json:"id"`
  IDStr                string      `json:"id_str"`
  InReplyToScreenName  interface{} `json:"in_reply_to_screen_name"`
  InReplyToStatusID    interface{} `json:"in_reply_to_status_id"`
  InReplyToStatusIDStr interface{} `json:"in_reply_to_status_id_str"`
  InReplyToUserID      interface{} `json:"in_reply_to_user_id"`
  InReplyToUserIDStr   interface{} `json:"in_reply_to_user_id_str"`
  IsQuoteStatus        bool        `json:"is_quote_status"`
  Lang                 string      `json:"lang"`
  Place                interface{} `json:"place"`
  RetweetCount         int         `json:"retweet_count"`
  Retweeted            bool        `json:"retweeted"`
  Source               string      `json:"source"`
  Text                 string      `json:"text"`
  TimestampMs          string      `json:"timestamp_ms"`
  Truncated            bool        `json:"truncated"`
  User                 struct {
    ContributorsEnabled            bool        `json:"contributors_enabled"`
    CreatedAt                      string      `json:"created_at"`
    DefaultProfile                 bool        `json:"default_profile"`
    DefaultProfileImage            bool        `json:"default_profile_image"`
    Description                    string      `json:"description"`
    FavouritesCount                int         `json:"favourites_count"`
    FollowRequestSent              interface{} `json:"follow_request_sent"`
    FollowersCount                 int         `json:"followers_count"`
    Following                      interface{} `json:"following"`
    FriendsCount                   int         `json:"friends_count"`
    GeoEnabled                     bool        `json:"geo_enabled"`
    ID                             int         `json:"id"`
    IDStr                          string      `json:"id_str"`
    IsTranslator                   bool        `json:"is_translator"`
    Lang                           string      `json:"lang"`
    ListedCount                    int         `json:"listed_count"`
    Location                       string      `json:"location"`
    Name                           string      `json:"name"`
    Notifications                  interface{} `json:"notifications"`
    ProfileBackgroundColor         string      `json:"profile_background_color"`
    ProfileBackgroundImageURL      string      `json:"profile_background_image_url"`
    ProfileBackgroundImageURLHTTPS string      `json:"profile_background_image_url_https"`
    ProfileBackgroundTile          bool        `json:"profile_background_tile"`
    ProfileBannerURL               string      `json:"profile_banner_url"`
    ProfileImageURL                string      `json:"profile_image_url"`
    ProfileImageURLHTTPS           string      `json:"profile_image_url_https"`
    ProfileLinkColor               string      `json:"profile_link_color"`
    ProfileSidebarBorderColor      string      `json:"profile_sidebar_border_color"`
    ProfileSidebarFillColor        string      `json:"profile_sidebar_fill_color"`
    ProfileTextColor               string      `json:"profile_text_color"`
    ProfileUseBackgroundImage      bool        `json:"profile_use_background_image"`
    Protected                      bool        `json:"protected"`
    ScreenName                     string      `json:"screen_name"`
    StatusesCount                  int         `json:"statuses_count"`
    TimeZone                       string      `json:"time_zone"`
    URL                            string      `json:"url"`
    UtcOffset                      int         `json:"utc_offset"`
    Verified                       bool        `json:"verified"`
  } `json:"user"`
}
