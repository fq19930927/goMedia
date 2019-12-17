package defs

//requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//video
type VideoInfo struct {
	Id           string `json:"id"`
	AuthorId     int    `json:"author_id"`
	Name         string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
}

type Comments struct {
	Comments []*Comment `json:"comments"`
}

type Comment struct {
	Id      string `json:"id"`
	VideoId string `json:"video_id"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type SimpleSession struct {
	//login name
	Username string
	TTL      int64
}
