import { Routes, Route } from 'react-router-dom'
import { useState, useEffect } from 'react'
import Home from './pages/Home'
import Login from './pages/Login'
import ActivityDetail from './pages/ActivityDetail'
import Layout from './components/Layout'
import { AuthProvider } from './contexts/AuthContext'

function App() {
  return (
    <AuthProvider>
      <Layout>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login />} />
          <Route path="/activity/:id" element={<ActivityDetail />} />
        </Routes>
      </Layout>
    </AuthProvider>
  )
}

export default App