package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"myvault-backend/pkg/github"

	"github.com/gin-gonic/gin"
)

type GithubHandler struct {
	githubService GithubService
	userService   UserService
}

type GithubService interface {
	GetAccessToken(code string) (string, error)
	GetUser(accessToken string) (*github.User, error)
}

// 使用pkg/github中的User类型，移除重复定义
type GithubUser = github.User

func NewGithubHandler(githubService GithubService, userService UserService) *GithubHandler {
	return &GithubHandler{
		githubService: githubService,
		userService:   userService,
	}
}

func (h *GithubHandler) GithubLogin(c *gin.Context) {
	// 重定向到GitHub OAuth授权页面
	clientID := "your-client-id" // 应该从配置中获取
	redirectURL := fmt.Sprintf("https://github.com/login/oauth/authorize?client_id=%s&scope=user:email,repo", clientID)
	
	c.JSON(http.StatusOK, gin.H{
		"redirect_url": redirectURL,
	})
}

func (h *GithubHandler) GithubCallback(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code parameter"})
		return
	}

	// 获取访问令牌
	accessToken, err := h.githubService.GetAccessToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get access token"})
		return
	}

	// 获取用户信息
	githubUser, err := h.githubService.GetUser(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info"})
		return
	}

	// 创建或获取用户
	user, err := h.userService.GetOrCreateGithubUser(
		strconv.Itoa(githubUser.ID),
		githubUser.Login,
		githubUser.Email,
		githubUser.AvatarURL,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":         user,
		"access_token": accessToken,
	})
}