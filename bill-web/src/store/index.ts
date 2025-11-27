import { reactive, readonly } from 'vue'

// 用户状态
const userState = reactive({
  isLoggedIn: false,
  userId: '',
  username: '',
  token: ''
})

// 设置用户信息
export const setUserInfo = (userInfo: {
  userId: string
  username: string
  token: string
}) => {
  userState.isLoggedIn = true
  userState.userId = userInfo.userId
  userState.username = userInfo.username
  userState.token = userInfo.token
  localStorage.setItem('token', userInfo.token)
}

// 清除用户信息
export const clearUserInfo = () => {
  userState.isLoggedIn = false
  userState.userId = ''
  userState.username = ''
  userState.token = ''
  localStorage.removeItem('token')
}

// 获取用户状态
export const getUserState = () => readonly(userState)

// 应用状态
const appState = reactive({
  loading: false,
  theme: 'light'
})

// 设置加载状态
export const setLoading = (loading: boolean) => {
  appState.loading = loading
}

// 设置主题
export const setTheme = (theme: 'light' | 'dark') => {
  appState.theme = theme
  document.documentElement.setAttribute('data-theme', theme)
}

// 获取应用状态
export const getAppState = () => readonly(appState)
