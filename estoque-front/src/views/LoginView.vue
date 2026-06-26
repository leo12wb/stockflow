<template>
  <div style="min-height:100vh;display:flex;align-items:center;justify-content:center;background:linear-gradient(135deg,#0f172a 0%,#1e1b4b 50%,#0f172a 100%);">

    <div style="width:100%;max-width:400px;padding:0 16px;">

      <!-- Card -->
      <div style="background:#fff;border-radius:16px;padding:40px;box-shadow:0 20px 60px rgba(0,0,0,.3);">

        <!-- Logo -->
        <div style="text-align:center;margin-bottom:32px;">
          <div style="width:56px;height:56px;background:linear-gradient(135deg,#6366f1,#8b5cf6);border-radius:14px;display:flex;align-items:center;justify-content:center;margin:0 auto 16px;">
            <i class="pi pi-box" style="color:#fff;font-size:24px;" />
          </div>
          <h1 style="margin:0 0 4px;font-size:22px;font-weight:700;color:#0f172a;">Sistema de Estoque</h1>
          <p style="margin:0;font-size:14px;color:#64748b;">Entre com suas credenciais</p>
        </div>

        <!-- Form -->
        <form @submit.prevent="handleLogin">
          <div style="margin-bottom:16px;">
            <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:6px;">E-mail</label>
            <InputText
              v-model="form.email"
              type="email"
              placeholder="seu@email.com"
              :invalid="!!errors.email"
              style="width:100%;"
            />
            <small v-if="errors.email" style="color:#ef4444;font-size:12px;">{{ errors.email }}</small>
          </div>

          <div style="margin-bottom:24px;">
            <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:6px;">Senha</label>
            <Password
              v-model="form.senha"
              placeholder="••••••"
              :feedback="false"
              toggle-mask
              :invalid="!!errors.senha"
              style="width:100%;"
              :input-style="{ width: '100%' }"
            />
            <small v-if="errors.senha" style="color:#ef4444;font-size:12px;">{{ errors.senha }}</small>
          </div>

          <Message v-if="errorMsg" severity="error" :closable="false" style="margin-bottom:16px;">{{ errorMsg }}</Message>

          <Button
            type="submit"
            label="Entrar"
            icon="pi pi-sign-in"
            :loading="authStore.loading"
            style="width:100%;"
          />
        </form>

        <!-- Footer hint -->
        <div style="margin-top:24px;padding-top:20px;border-top:1px solid #f1f5f9;text-align:center;">
          <p style="margin:0;font-size:12px;color:#94a3b8;">admin@estoque.com · senha: admin123</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import InputText from 'primevue/inputtext'
import Password from 'primevue/password'
import Button from 'primevue/button'
import Message from 'primevue/message'
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()
const form = ref({ email: '', senha: '' })
const errors = ref<Record<string, string>>({})
const errorMsg = ref('')

async function handleLogin() {
  errors.value = {}
  errorMsg.value = ''
  if (!form.value.email) errors.value.email = 'E-mail é obrigatório'
  if (!form.value.senha) errors.value.senha = 'Senha é obrigatória'
  if (Object.keys(errors.value).length) return
  try {
    await authStore.login(form.value.email, form.value.senha)
  } catch (err: any) {
    errorMsg.value = err.response?.data?.error || 'Credenciais inválidas'
  }
}
</script>
