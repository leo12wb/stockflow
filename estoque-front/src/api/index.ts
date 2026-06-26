import api from './axios'

export const authApi = {
  login: (email: string, senha: string) =>
    api.post('/auth/login', { email, senha }),
  me: () => api.get('/auth/me'),
  register: (data: object) => api.post('/auth/register', data),
}

export const produtosApi = {
  getAll: (params?: object) => api.get('/produtos', { params }),
  getById: (id: number) => api.get(`/produtos/${id}`),
  getByCodigo: (codigo: string) => api.get(`/produtos/codigo/${codigo}`),
  create: (data: object) => api.post('/produtos', data),
  update: (id: number, data: object) => api.put(`/produtos/${id}`, data),
  delete: (id: number) => api.delete(`/produtos/${id}`),
  restore: (id: number) => api.patch(`/produtos/${id}/restaurar`),
  forceDelete: (id: number) => api.delete(`/produtos/${id}/force`),
}

export const fornecedoresApi = {
  getAll: (params?: object) => api.get('/fornecedores', { params }),
  getById: (id: number) => api.get(`/fornecedores/${id}`),
  create: (data: object) => api.post('/fornecedores', data),
  update: (id: number, data: object) => api.put(`/fornecedores/${id}`, data),
  delete: (id: number) => api.delete(`/fornecedores/${id}`),
  restore: (id: number) => api.patch(`/fornecedores/${id}/restaurar`),
}

export const movimentacoesApi = {
  getAll: (params?: object) => api.get('/movimentacoes', { params }),
  create: (data: object) => api.post('/movimentacoes', data),
}

export const relatoriosApi = {
  inventario: () => api.get('/relatorios/inventario'),
  estoqueBaixo: () => api.get('/relatorios/estoque-baixo'),
  entradas: () => api.get('/relatorios/entradas'),
  saidas: () => api.get('/relatorios/saidas'),
  maisVendidos: () => api.get('/relatorios/mais-vendidos'),
  movimentacoes: () => api.get('/relatorios/movimentacoes'),
}
