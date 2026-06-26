import { defineStore } from 'pinia'
import { ref } from 'vue'
import { authApi } from '@/api'
import router from '@/router'

interface User {
  id: number
  nome: string
  email: string
  perfil: string
}

export const useAuthStore = defineStore('auth', () => {
  const user = ref<User | null>(JSON.parse(localStorage.getItem('user') || 'null'))
  const token = ref<string | null>(localStorage.getItem('token'))
  const loading = ref(false)

  async function login(email: string, senha: string) {
    loading.value = true
    try {
      const { data } = await authApi.login(email, senha)
      token.value = data.data.token
      user.value = data.data.usuario
      localStorage.setItem('token', data.data.token)
      localStorage.setItem('user', JSON.stringify(data.data.usuario))
      router.push('/')
    } finally {
      loading.value = false
    }
  }

  async function fetchUser() {
    try {
      const { data } = await authApi.me()
      user.value = data.data
      localStorage.setItem('user', JSON.stringify(data.data))
    } catch {
      logout()
    }
  }

  function logout() {
    user.value = null
    token.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('user')
    router.push('/login')
  }

  const isAuthenticated = () => !!token.value

  return { user, token, loading, login, logout, fetchUser, isAuthenticated }
})
