<template>
  <div style="display:flex; height:100vh; overflow:hidden; background:#f1f5f9;">

    <!-- Sidebar -->
    <aside style="width:240px; min-width:240px; background:#0f172a; display:flex; flex-direction:column; height:100vh;">

      <!-- Logo -->
      <div style="padding:24px 20px 20px; border-bottom:1px solid rgba(255,255,255,.08);">
        <div style="display:flex; align-items:center; gap:10px;">
          <div style="width:36px;height:36px;background:linear-gradient(135deg,#6366f1,#8b5cf6);border-radius:10px;display:flex;align-items:center;justify-content:center;">
            <i class="pi pi-box" style="color:#fff;font-size:16px;" />
          </div>
          <div>
            <div style="font-weight:700;color:#f8fafc;font-size:15px;">Estoque</div>
            <div style="font-size:11px;color:#94a3b8;">Sistema de Controle</div>
          </div>
        </div>
      </div>

      <!-- Nav -->
      <nav style="flex:1;padding:12px 10px;overflow-y:auto;">
        <div style="font-size:10px;font-weight:600;color:#475569;letter-spacing:.08em;padding:8px 10px 4px;text-transform:uppercase;">Menu</div>
        <RouterLink
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          custom
          v-slot="{ isActive, navigate }"
        >
          <button
            @click="navigate"
            :style="{
              display:'flex', alignItems:'center', gap:'10px',
              width:'100%', padding:'9px 12px', borderRadius:'8px',
              border:'none', cursor:'pointer', marginBottom:'2px',
              background: isActive ? 'rgba(99,102,241,.2)' : 'transparent',
              color: isActive ? '#a5b4fc' : '#94a3b8',
              fontWeight: isActive ? '600' : '400',
              fontSize:'14px', textAlign:'left', transition:'all .15s',
            }"
            @mouseenter="e => { if(!isActive)(e.target as HTMLElement).style.background='rgba(255,255,255,.05)'; (e.target as HTMLElement).style.color='#e2e8f0' }"
            @mouseleave="e => { if(!isActive)(e.target as HTMLElement).style.background='transparent'; (e.target as HTMLElement).style.color='#94a3b8' }"
          >
            <i :class="['pi', item.icon]" style="font-size:15px;width:18px;text-align:center;" />
            {{ item.label }}
          </button>
        </RouterLink>
      </nav>

      <!-- User -->
      <div style="padding:12px 10px;border-top:1px solid rgba(255,255,255,.08);">
        <div style="display:flex;align-items:center;gap:10px;padding:8px 12px;border-radius:8px;margin-bottom:8px;">
          <div style="width:32px;height:32px;background:rgba(99,102,241,.3);border-radius:50%;display:flex;align-items:center;justify-content:center;">
            <i class="pi pi-user" style="color:#a5b4fc;font-size:13px;" />
          </div>
          <div style="flex:1;min-width:0;">
            <div style="font-size:13px;font-weight:600;color:#e2e8f0;white-space:nowrap;overflow:hidden;text-overflow:ellipsis;">{{ authStore.user?.nome }}</div>
            <div style="font-size:11px;color:#64748b;text-transform:capitalize;">{{ authStore.user?.perfil }}</div>
          </div>
        </div>
        <button
          @click="authStore.logout()"
          style="display:flex;align-items:center;gap:8px;width:100%;padding:8px 12px;border-radius:8px;border:none;cursor:pointer;background:rgba(239,68,68,.1);color:#f87171;font-size:13px;font-weight:500;transition:background .15s;"
          @mouseenter="e => (e.currentTarget as HTMLElement).style.background='rgba(239,68,68,.2)'"
          @mouseleave="e => (e.currentTarget as HTMLElement).style.background='rgba(239,68,68,.1)'"
        >
          <i class="pi pi-sign-out" style="font-size:13px;" />
          Sair
        </button>
      </div>
    </aside>

    <!-- Main -->
    <main style="flex:1;overflow-y:auto;min-width:0;">
      <RouterView />
    </main>

  </div>
</template>

<script setup lang="ts">
import { useAuthStore } from '@/stores/auth'

const authStore = useAuthStore()

const navItems = [
  { to: '/',              label: 'Dashboard',      icon: 'pi-chart-bar' },
  { to: '/produtos',      label: 'Produtos',        icon: 'pi-box' },
  { to: '/fornecedores',  label: 'Fornecedores',    icon: 'pi-truck' },
  { to: '/movimentacoes', label: 'Movimentações',   icon: 'pi-arrow-right-arrow-left' },
  { to: '/relatorios',    label: 'Relatórios',      icon: 'pi-file-pdf' },
]
</script>
