package services

import (
	"myvault-backend/pkg/github"
)

type GithubService struct {
	client *github.Client
}

func NewGithubService(clientID, clientSecret string) *GithubService {
	return &GithubService{
		client: github.NewClient(clientID, clientSecret),
	}
}

func (s *GithubService) GetAccessToken(code string) (string, error) {
	return s.client.GetAccessToken(code)
}

func (s *GithubService) GetUser(accessToken string) (*github.User, error) {
	return s.client.GetUser(accessToken)
}

func (s *GithubService) GetUserCommits(accessToken, username string, since string) ([]github.Commit, error) {
	// 这里暂时返回空数组，实际实现需要解析since参数
	return []github.Commit{}, nil
}