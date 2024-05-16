package convert

import (
	md "github.com/hxzhouh/html-to-markdown"
	"regexp"
	"strings"
	"time"
)

var converter *md.Converter

func init() {
	converter = md.NewConverter("", true, nil)

}

type Post struct {
	Title     string   `json:"title"`
	SubTitle  string   `json:"sub_title"`
	Content   string   `json:"content"`
	CreateAt  string   `json:"create_at"`
	Tags      []string `json:"tags"`
	SourceUrl string   `json:"source_url"`
	FileName  string   `json:"file_name"`
	Author    string   `json:"author"`
}

func Convert(name string, content []byte) (*Post, error) {
	temp, err := converter.ConvertString(string(content))
	if err != nil {
		return nil, err
	}
	post := &Post{Content: temp}
	post.CreateAt = nameToString(name)
	post.FileName = changeName(name)
	post.Author = "huizhou92"
	post.Title = strings.Split(temp, "\n")[0]
	post.SubTitle = strings.Split(temp, "\n")[1]
	post.SourceUrl = getSourceHttpLink(temp)
	return post, nil
}

func changeName(name string) string {
	mdFilePath := strings.Replace(name, "posts/", "posts-md/", 1)
	mdFilePath = strings.Replace(mdFilePath, ".html", ".md", 1)
	return mdFilePath
}
func nameToString(name string) string {
	d := strings.Split(name[len("posts/"):], "_")[0]
	createTime, err := time.Parse("2006-01-02", d)
	if err != nil || createTime.Unix() == 0 {
		return time.Now().Format("2006-01-02 15:04:05")
	}
	return createTime.Format("2006-01-02 15:04:05")
}

func getSourceHttpLink(text string) string {
	// 定义匹配http和https链接的正则表达式
	re := regexp.MustCompile(`https?://[^\s]+`)
	// 找到所有匹配的链接
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return ""
	}

	// 返回最后一个匹配的链接
	return matches[len(matches)-2]
}
