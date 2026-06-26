import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/login',
      name: 'login',
      component: () => import('@/views/LoginView.vue'),
      meta: { guest: true },
    },
    {
      path: '/',
      component: () => import('@/components/AppLayout.vue'),
      meta: { requiresAuth: true },
      children: [
        {
          path: '',
          name: 'dashboard',
          component: () => import('@/views/DashboardView.vue'),
        },
        {
          path: 'produtos',
          name: 'produtos',
          component: () => import('@/views/produtos/ProdutosView.vue'),
        },
        {
          path: 'fornecedores',
          name: 'fornecedores',
          component: () => import('@/views/fornecedores/FornecedoresView.vue'),
        },
        {
          path: 'movimentacoes',
          name: 'movimentacoes',
          component: () => import('@/views/movimentacoes/MovimentacoesView.vue'),
        },
        {
          path: 'relatorios',
          name: 'relatorios',
          component: () => import('@/views/relatorios/RelatoriosView.vue'),
        },
      ],
    },
  ],
})

router.beforeEach((to, _from, next) => {
  const token = localStorage.getItem('token')
  if (to.meta.requiresAuth && !token) {
    next('/login')
  } else if (to.meta.guest && token) {
    next('/')
  } else {
    next()
  }
})

export default router
