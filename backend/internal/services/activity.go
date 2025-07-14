package services

import (
	"encoding/json"
	"fmt"
	"myvault-backend/internal/models"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ActivityService struct {
	db        *gorm.DB
	redis     *redis.Client
	aiService *AIService
}

func NewActivityService(db *gorm.DB, redis *redis.Client, aiService *AIService) *ActivityService {
	return &ActivityService{
		db:        db,
		redis:     redis,
		aiService: aiService,
	}
}

func (s *ActivityService) GetUserActivities(userID uint, limit int, offset int) ([]models.Activity, error) {
	var activities []models.Activity
	
	query := s.db.Where("user_id = ?", userID).
		Preload("Commits").
		Preload("DataSources").
		Order("date DESC")

	if limit > 0 {
		query = query.Limit(limit)
	}
	
	if offset > 0 {
		query = query.Offset(offset)
	}

	if err := query.Find(&activities).Error; err != nil {
		return nil, err
	}

	return activities, nil
}

func (s *ActivityService) GetActivityByID(userID, activityID uint) (*models.Activity, error) {
	var activity models.Activity
	
	if err := s.db.Where("id = ? AND user_id = ?", activityID, userID).
		Preload("Commits").
		Preload("DataSources").
		First(&activity).Error; err != nil {
		return nil, err
	}

	return &activity, nil
}

func (s *ActivityService) CreateOrUpdateActivity(userID uint, date time.Time, commits []models.Commit, dataSources []models.DataSource) (*models.Activity, error) {
	// 检查是否已存在该日期的活动
	var activity models.Activity
	dateStart := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	dateEnd := dateStart.Add(24 * time.Hour)

	err := s.db.Where("user_id = ? AND date >= ? AND date < ?", userID, dateStart, dateEnd).
		Preload("Commits").
		Preload("DataSources").
		First(&activity).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// 如果不存在，创建新的活动记录
	if err == gorm.ErrRecordNotFound {
		activity = models.Activity{
			UserID:      userID,
			Date:        dateStart,
			HasActivity: len(commits) > 0,
			CommitCount: len(commits),
		}
		if err := s.db.Create(&activity).Error; err != nil {
			return nil, err
		}
	} else {
		// 更新现有活动
		activity.HasActivity = len(commits) > 0
		activity.CommitCount = len(commits)
		if err := s.db.Save(&activity).Error; err != nil {
			return nil, err
		}

		// 删除旧的提交记录和数据源
		s.db.Where("activity_id = ?", activity.ID).Delete(&models.Commit{})
		s.db.Where("activity_id = ?", activity.ID).Delete(&models.DataSource{})
	}

	// 添加新的提交记录
	for i := range commits {
		commits[i].ActivityID = activity.ID
	}
	if len(commits) > 0 {
		if err := s.db.Create(&commits).Error; err != nil {
			return nil, err
		}
	}

	// 添加新的数据源
	for i := range dataSources {
		dataSources[i].ActivityID = activity.ID
	}
	if len(dataSources) > 0 {
		if err := s.db.Create(&dataSources).Error; err != nil {
			return nil, err
		}
	}

	// 生成AI摘要
	if len(commits) > 0 {
		summary, err := s.generateAISummary(commits)
		if err == nil {
			activity.Summary = summary
			activity.AIGenerated = true
			s.db.Save(&activity)
		}
	} else {
		activity.Summary = "今日无编程活动"
		activity.AIGenerated = false
		s.db.Save(&activity)
	}

	// 重新加载活动数据
	s.db.Where("id = ?", activity.ID).
		Preload("Commits").
		Preload("DataSources").
		First(&activity)

	return &activity, nil
}

func (s *ActivityService) generateAISummary(commits []models.Commit) (string, error) {
	if len(commits) == 0 {
		return "今日无编程活动", nil
	}

	// 构建提示词
	var promptBuilder strings.Builder
	promptBuilder.WriteString("以下是今日的代码提交记录：\n\n")

	for _, commit := range commits {
		promptBuilder.WriteString(fmt.Sprintf("时间: %s\n", commit.Time.Format("15:04")))
		promptBuilder.WriteString(fmt.Sprintf("仓库: %s\n", commit.Repository))
		promptBuilder.WriteString(fmt.Sprintf("提交信息: %s\n", commit.Message))
		promptBuilder.WriteString(fmt.Sprintf("文件数: %d, 新增: %d行, 删除: %d行\n\n", 
			commit.Files, commit.Additions, commit.Deletions))
	}

	promptBuilder.WriteString("请基于以上信息生成一份简洁的每日编程活动摘要。")

	return s.aiService.GenerateSummary(promptBuilder.String())
}

func (s *ActivityService) SyncActivities(userID uint, force bool) error {
	// 这里应该实现从各种数据源同步活动的逻辑
	// 暂时返回nil，表示同步成功
	return nil
}

func (s *ActivityService) GetTodayActivity(userID uint) (*models.Activity, error) {
	now := time.Now()
	dateStart := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	dateEnd := dateStart.Add(24 * time.Hour)

	var activity models.Activity
	err := s.db.Where("user_id = ? AND date >= ? AND date < ?", userID, dateStart, dateEnd).
		Preload("Commits").
		Preload("DataSources").
		First(&activity).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &activity, nil
}