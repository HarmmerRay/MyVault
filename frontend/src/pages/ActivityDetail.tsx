import React, { useState, useEffect } from 'react'
import { useParams, useNavigate } from 'react-router-dom'
import { ArrowLeft, Github, Calendar, Clock, Code, GitCommit, FileText } from 'lucide-react'
import { motion } from 'framer-motion'

interface ActivityDetail {
  id: string
  date: string
  summary: string
  commitCount: number
  commits: Array<{
    id: string
    message: string
    time: string
    files: number
    additions: number
    deletions: number
  }>
  aiAnalysis: string
  totalTimeSpent: string
  repositories: Array<{
    name: string
    commits: number
    languages: string[]
  }>
}

const ActivityDetail: React.FC = () => {
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()
  const [activity, setActivity] = useState<ActivityDetail | null>(null)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    // 模拟获取详细数据
    const mockActivity: ActivityDetail = {
      id: id || '1',
      date: '2024-01-15',
      summary: '完成了React组件的优化，提交了3个bug修复，更新了文档',
      commitCount: 5,
      commits: [
        {
          id: '1',
          message: 'feat: 优化Timeline组件性能',
          time: '09:30',
          files: 3,
          additions: 45,
          deletions: 12
        },
        {
          id: '2',
          message: 'fix: 修复用户认证状态问题',
          time: '11:15',
          files: 2,
          additions: 23,
          deletions: 8
        },
        {
          id: '3',
          message: 'docs: 更新README文档',
          time: '14:20',
          files: 1,
          additions: 15,
          deletions: 3
        },
        {
          id: '4',
          message: 'style: 调整CSS样式',
          time: '16:45',
          files: 4,
          additions: 32,
          deletions: 18
        },
        {
          id: '5',
          message: 'test: 添加单元测试',
          time: '18:30',
          files: 2,
          additions: 67,
          deletions: 5
        }
      ],
      aiAnalysis: `今天的编程活动主要集中在MyVault项目的前端优化上。上午主要进行了Timeline组件的性能优化，通过代码重构减少了不必要的重渲染，提升了用户体验。

中午解决了一个关键的用户认证状态问题，这个bug可能会影响用户的登录体验。下午更新了项目文档，确保新功能的使用方法得到及时记录。

傍晚时分进行了样式调整，使界面更加美观和一致。最后添加了单元测试，提高了代码质量和可维护性。

总的来说，今天的工作涵盖了性能优化、bug修复、文档更新、样式调整和测试添加等多个方面，是一个比较全面的开发日。`,
      totalTimeSpent: '约6小时',
      repositories: [
        {
          name: 'MyVault',
          commits: 5,
          languages: ['TypeScript', 'CSS', 'HTML']
        }
      ]
    }

    setTimeout(() => {
      setActivity(mockActivity)
      setLoading(false)
    }, 500)
  }, [id])

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-[60vh]">
        <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-indigo-600"></div>
      </div>
    )
  }

  if (!activity) {
    return (
      <div className="text-center py-12">
        <p className="text-gray-500">活动详情不存在</p>
      </div>
    )
  }

  return (
    <div className="max-w-4xl mx-auto">
      <motion.div
        initial={{ opacity: 0, y: 20 }}
        animate={{ opacity: 1, y: 0 }}
        className="mb-6"
      >
        <button
          onClick={() => navigate('/')}
          className="flex items-center space-x-2 text-indigo-600 hover:text-indigo-700 transition-colors"
        >
          <ArrowLeft className="w-5 h-5" />
          <span>返回时间轴</span>
        </button>
      </motion.div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8">
        {/* 主要内容 */}
        <div className="lg:col-span-2 space-y-6">
          {/* 日期和摘要 */}
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.1 }}
            className="glass-effect rounded-xl p-6"
          >
            <div className="flex items-center space-x-3 mb-4">
              <Calendar className="w-6 h-6 text-indigo-600" />
              <h1 className="text-2xl font-bold text-gray-800">
                {new Date(activity.date).toLocaleDateString('zh-CN', {
                  year: 'numeric',
                  month: 'long',
                  day: 'numeric'
                })}
              </h1>
            </div>
            <p className="text-gray-700 text-lg">{activity.summary}</p>
          </motion.div>

          {/* AI分析 */}
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.2 }}
            className="glass-effect rounded-xl p-6"
          >
            <div className="flex items-center space-x-3 mb-4">
              <FileText className="w-6 h-6 text-purple-600" />
              <h2 className="text-xl font-semibold text-gray-800">AI分析</h2>
            </div>
            <div className="prose prose-gray max-w-none">
              {activity.aiAnalysis.split('\n\n').map((paragraph, index) => (
                <p key={index} className="text-gray-700 mb-3 leading-relaxed">
                  {paragraph}
                </p>
              ))}
            </div>
          </motion.div>

          {/* 提交记录 */}
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ delay: 0.3 }}
            className="glass-effect rounded-xl p-6"
          >
            <div className="flex items-center space-x-3 mb-4">
              <GitCommit className="w-6 h-6 text-green-600" />
              <h2 className="text-xl font-semibold text-gray-800">提交记录</h2>
            </div>
            <div className="space-y-3">
              {activity.commits.map((commit, index) => (
                <div key={commit.id} className="flex items-center justify-between p-3 bg-white/50 rounded-lg">
                  <div className="flex-1">
                    <p className="font-medium text-gray-800">{commit.message}</p>
                    <p className="text-sm text-gray-500">
                      {commit.time} • {commit.files} 个文件 • +{commit.additions} -{commit.deletions}
                    </p>
                  </div>
                  <div className="text-sm text-gray-500">
                    #{index + 1}
                  </div>
                </div>
              ))}
            </div>
          </motion.div>
        </div>

        {/* 侧边栏 */}
        <div className="space-y-6">
          {/* 统计信息 */}
          <motion.div
            initial={{ opacity: 0, x: 20 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: 0.4 }}
            className="glass-effect rounded-xl p-6"
          >
            <h3 className="text-lg font-semibold text-gray-800 mb-4">统计信息</h3>
            <div className="space-y-3">
              <div className="flex items-center justify-between">
                <span className="text-gray-600">总提交数</span>
                <span className="font-semibold text-indigo-600">{activity.commitCount}</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-gray-600">编程时长</span>
                <span className="font-semibold text-indigo-600">{activity.totalTimeSpent}</span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-gray-600">涉及文件</span>
                <span className="font-semibold text-indigo-600">
                  {activity.commits.reduce((sum, commit) => sum + commit.files, 0)}
                </span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-gray-600">代码增加</span>
                <span className="font-semibold text-green-600">
                  +{activity.commits.reduce((sum, commit) => sum + commit.additions, 0)}
                </span>
              </div>
              <div className="flex items-center justify-between">
                <span className="text-gray-600">代码删除</span>
                <span className="font-semibold text-red-600">
                  -{activity.commits.reduce((sum, commit) => sum + commit.deletions, 0)}
                </span>
              </div>
            </div>
          </motion.div>

          {/* 涉及仓库 */}
          <motion.div
            initial={{ opacity: 0, x: 20 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: 0.5 }}
            className="glass-effect rounded-xl p-6"
          >
            <h3 className="text-lg font-semibold text-gray-800 mb-4">涉及仓库</h3>
            <div className="space-y-3">
              {activity.repositories.map((repo, index) => (
                <div key={index} className="p-3 bg-white/50 rounded-lg">
                  <div className="flex items-center space-x-2 mb-2">
                    <Github className="w-4 h-4 text-gray-600" />
                    <span className="font-medium text-gray-800">{repo.name}</span>
                  </div>
                  <div className="text-sm text-gray-600 mb-2">
                    {repo.commits} 次提交
                  </div>
                  <div className="flex flex-wrap gap-1">
                    {repo.languages.map((lang, langIndex) => (
                      <span
                        key={langIndex}
                        className="px-2 py-1 text-xs bg-indigo-100 text-indigo-600 rounded-full"
                      >
                        {lang}
                      </span>
                    ))}
                  </div>
                </div>
              ))}
            </div>
          </motion.div>
        </div>
      </div>
    </div>
  )
}

export default ActivityDetail