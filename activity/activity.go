package activity

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GitHub用户活动的数据结构
type GitHubEvent struct {
	Type    string `json:"type"`
	Created string `json:"created_at"`
	Actor   struct {
		Login string `json:"login"`
	} `json:"actor"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
}

// GetUserActivity 获取指定用户的GitHub活动
func GetUserActivity(username string, token string) ([]GitHubEvent, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %v", err)
	}

	req.Header.Set("Authorization", "token "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}
	defer resp.Body.Close()

	var events []GitHubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
	}

	return events, nil
}
