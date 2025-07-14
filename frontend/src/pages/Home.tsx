import React, { useState, useEffect } from 'react'
import { useAuth } from '../contexts/AuthContext'
import { Link } from 'react-router-dom'
import { motion } from 'framer-motion'
import { Github, Calendar, Clock, AlertCircle } from 'lucide-react'

interface DayActivity {
  id: string
  date: string
  summary: string
  commitCount: number
  aiGenerated: boolean
  hasActivity: boolean
}

const Home: React.FC = () => {
  const { isAuthenticated } = useAuth()
  const [activities, setActivities] = useState<DayActivity[]>([])
  const [currentTime, setCurrentTime] = useState(new Date())

  useEffect(() => {
    const timer = setInterval(() => {
      setCurrentTime(new Date())
    }, 1000)

    return () => clearInterval(timer)
  }, [])

  useEffect(() => {
    // 模拟数据
    const mockActivities: DayActivity[] = [
      {
        id: '1',
        date: '2024-01-15',
        summary: '完成了React组件的优化，提交了3个bug修复，更新了文档',
        commitCount: 5,
        aiGenerated: true,
        hasActivity: true
      },
      {
        id: '2',
        date: '2024-01-14',
        summary: '实现了用户认证系统，集成了GitHub OAuth，完成了API接口开发',
        commitCount: 8,
        aiGenerated: true,
        hasActivity: true
      },
      {
        id: '3',
        date: '2024-01-13',
        summary: '今日无编程活动',
        commitCount: 0,
        aiGenerated: false,
        hasActivity: false
      }
    ]
    setActivities(mockActivities)
  }, [])

  const isBeforeNoon = currentTime.getHours() < 12
  const todayActivity = activities.find(a => 
    new Date(a.date).toDateString() === currentTime.toDateString()
  )

  if (!isAuthenticated) {
    return (
      <div className="flex items-center justify-center min-h-[60vh]">
        <div className="text-center">
          <Calendar className="w-16 h-16 text-indigo-600 mx-auto mb-4" />
          <h2 className="text-2xl font-bold text-gray-800 mb-2">欢迎来到MyVault</h2>
          <p className="text-gray-600 mb-6">登录后查看您的每日活动时间轴</p>
          <Link
            to="/login"
            className="inline-flex items-center px-6 py-3 rounded-lg bg-indigo-600 text-white hover:bg-indigo-700 transition-colors"
          >
            <Github className="w-5 h-5 mr-2" />
            立即登录
          </Link>
        </div>
      </div>
    )
  }

  return (
    <div className="max-w-4xl mx-auto">
      <div className="text-center mb-8">
        <h1 className="text-4xl font-bold gradient-text mb-2">我的时间轴</h1>
        <p className="text-gray-600">记录每一天的编程足迹</p>
      </div>

      {/* 每日提醒 */}
      {isBeforeNoon && (!todayActivity || !todayActivity.hasActivity) && (
        <motion.div
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          className="mb-8 p-6 rounded-xl bg-gradient-to-r from-orange-100 to-red-100 border border-orange-200"
        >
          <div className="flex items-center space-x-3">
            <AlertCircle className="w-6 h-6 text-orange-600" />
            <div>
              <h3 className="text-lg font-semibold text-orange-800">醒醒吧少年，此时不搏何时搏！</h3>
              <p className="text-orange-700">今天还没有任何编程活动记录，赶紧开始写代码吧！</p>
            </div>
          </div>
        </motion.div>
      )}

      {/* 时间轴 */}
      <div className="space-y-6">
        {activities.map((activity, index) => (
          <motion.div
            key={activity.id}
            initial={{ opacity: 0, x: -20 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: index * 0.1 }}
            className="timeline-item"
          >
            <Link
              to={`/activity/${activity.id}`}
              className="block p-6 rounded-xl glass-effect hover:bg-white/20 transition-all duration-200 cursor-pointer group"
            >
              <div className="flex items-center justify-between mb-3">
                <div className="flex items-center space-x-3">
                  <Calendar className="w-5 h-5 text-indigo-600" />
                  <span className="font-semibold text-gray-800">
                    {new Date(activity.date).toLocaleDateString('zh-CN', {
                      year: 'numeric',
                      month: 'long',
                      day: 'numeric'
                    })}
                  </span>
                </div>
                <div className="flex items-center space-x-4">
                  {activity.commitCount > 0 && (
                    <div className="flex items-center space-x-1 text-sm text-gray-600">
                      <Github className="w-4 h-4" />
                      <span>{activity.commitCount} 次提交</span>
                    </div>
                  )}
                  {activity.aiGenerated && (
                    <span className="text-xs px-2 py-1 rounded-full bg-purple-100 text-purple-600">
                      AI生成
                    </span>
                  )}
                </div>
              </div>
              
              <p className={`text-gray-700 group-hover:text-gray-900 transition-colors ${
                !activity.hasActivity ? 'text-gray-500 italic' : ''
              }`}>
                {activity.summary}
              </p>
            </Link>
          </motion.div>
        ))}
      </div>
      
      {activities.length === 0 && (
        <div className="text-center py-12">
          <Clock className="w-16 h-16 text-gray-400 mx-auto mb-4" />
          <p className="text-gray-500">暂无活动记录</p>
        </div>
      )}
    </div>
  )
}

export default Home