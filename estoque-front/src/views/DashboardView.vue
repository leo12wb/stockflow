<template>
  <div style="padding:28px;">

    <!-- Header -->
    <div style="margin-bottom:28px;">
      <h1 style="margin:0 0 4px;font-size:22px;font-weight:700;color:#0f172a;">Dashboard</h1>
      <p style="margin:0;font-size:14px;color:#64748b;">Visão geral do sistema de estoque</p>
    </div>

    <!-- Cards -->
    <div style="display:grid;grid-template-columns:repeat(auto-fit,minmax(200px,1fr));gap:16px;margin-bottom:28px;">
      <div
        v-for="card in cards"
        :key="card.label"
        style="background:#fff;border-radius:12px;padding:20px;border:1px solid #e2e8f0;box-shadow:0 1px 3px rgba(0,0,0,.06);"
      >
        <div style="display:flex;align-items:center;justify-content:space-between;margin-bottom:12px;">
          <span style="font-size:13px;font-weight:500;color:#64748b;">{{ card.label }}</span>
          <div :style="`width:38px;height:38px;border-radius:10px;background:${card.bg};display:flex;align-items:center;justify-content:center;`">
            <i :class="['pi', card.icon]" :style="`color:${card.color};font-size:16px;`" />
          </div>
        </div>
        <div style="font-size:28px;font-weight:700;color:#0f172a;">
          <span v-if="loading">—</span>
          <span v-else>{{ card.value }}</span>
        </div>
      </div>
    </div>

    <!-- Panels -->
    <div style="display:grid;grid-template-columns:1fr 1fr;gap:20px;">

      <!-- Estoque Baixo -->
      <div style="background:#fff;border-radius:12px;border:1px solid #e2e8f0;box-shadow:0 1px 3px rgba(0,0,0,.06);overflow:hidden;">
        <div style="padding:16px 20px;border-bottom:1px solid #f1f5f9;display:flex;align-items:center;gap:8px;">
          <div style="width:28px;height:28px;background:#fff7ed;border-radius:8px;display:flex;align-items:center;justify-content:center;">
            <i class="pi pi-exclamation-triangle" style="color:#f97316;font-size:13px;" />
          </div>
          <span style="font-weight:600;font-size:14px;color:#1e293b;">Estoque Baixo</span>
          <span style="margin-left:auto;background:#fef3c7;color:#d97706;font-size:11px;font-weight:600;padding:2px 8px;border-radius:20px;">{{ estoqueBaixo.length }}</span>
        </div>
        <div style="padding:8px 0;max-height:280px;overflow-y:auto;">
          <div v-if="estoqueBaixo.length === 0" style="text-align:center;color:#94a3b8;padding:32px;font-size:14px;">
            Nenhum produto com estoque baixo
          </div>
          <div
            v-for="p in estoqueBaixo"
            :key="p.id"
            style="display:flex;align-items:center;justify-content:space-between;padding:10px 20px;border-bottom:1px solid #f8fafc;"
          >
            <span style="font-size:14px;color:#374151;font-weight:500;">{{ p.nome }}</span>
            <span style="background:#fff7ed;color:#ea580c;font-size:12px;font-weight:600;padding:3px 10px;border-radius:20px;border:1px solid #fed7aa;">
              {{ p.quantidade ?? p.quantidade_atual }} un
            </span>
          </div>
        </div>
      </div>

      <!-- Últimas Movimentações -->
      <div style="background:#fff;border-radius:12px;border:1px solid #e2e8f0;box-shadow:0 1px 3px rgba(0,0,0,.06);overflow:hidden;">
        <div style="padding:16px 20px;border-bottom:1px solid #f1f5f9;display:flex;align-items:center;gap:8px;">
          <div style="width:28px;height:28px;background:#eff6ff;border-radius:8px;display:flex;align-items:center;justify-content:center;">
            <i class="pi pi-history" style="color:#3b82f6;font-size:13px;" />
          </div>
          <span style="font-weight:600;font-size:14px;color:#1e293b;">Últimas Movimentações</span>
        </div>
        <div style="padding:8px 0;max-height:280px;overflow-y:auto;">
          <div v-if="movimentacoes.length === 0" style="text-align:center;color:#94a3b8;padding:32px;font-size:14px;">
            Nenhuma movimentação encontrada
          </div>
          <div
            v-for="m in movimentacoes.slice(0,5)"
            :key="m.id"
            style="display:flex;align-items:center;justify-content:space-between;padding:10px 20px;border-bottom:1px solid #f8fafc;"
          >
            <div>
              <p style="margin:0 0 2px;font-size:14px;font-weight:500;color:#374151;">{{ m.produto?.nome }}</p>
              <p style="margin:0;font-size:12px;color:#94a3b8;">{{ formatDate(m.created_at) }}</p>
            </div>
            <span
              :style="`font-size:13px;font-weight:700;padding:3px 10px;border-radius:20px;border:1px solid;
                background:${m.tipo==='ENTRADA'?'#f0fdf4':'#fff1f2'};
                color:${m.tipo==='ENTRADA'?'#16a34a':'#dc2626'};
                border-color:${m.tipo==='ENTRADA'?'#bbf7d0':'#fecaca'};`"
            >{{ m.tipo === 'ENTRADA' ? '+' : '-' }}{{ m.quantidade }}</span>
          </div>
        </div>
      </div>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { produtosApi, fornecedoresApi, movimentacoesApi, relatoriosApi } from '@/api'

const loading = ref(true)
const totalProdutos = ref(0)
const totalFornecedores = ref(0)
const movimentacoes = ref<any[]>([])
const estoqueBaixo = ref<any[]>([])

const cards = computed(() => [
  { label: 'Total de Produtos',  value: totalProdutos.value,      icon: 'pi-box',                    bg: '#eff6ff', color: '#3b82f6' },
  { label: 'Fornecedores',       value: totalFornecedores.value,  icon: 'pi-truck',                  bg: '#f0fdf4', color: '#22c55e' },
  { label: 'Movimentações',      value: movimentacoes.value.length, icon: 'pi-arrow-right-arrow-left', bg: '#f5f3ff', color: '#8b5cf6' },
  { label: 'Estoque Baixo',      value: estoqueBaixo.value.length, icon: 'pi-exclamation-triangle',   bg: '#fff7ed', color: '#f97316' },
])

const formatDate = (d: string) => new Date(d).toLocaleDateString('pt-BR')

onMounted(async () => {
  try {
    const [p, f, m, eb] = await Promise.all([
      produtosApi.getAll(),
      fornecedoresApi.getAll(),
      movimentacoesApi.getAll(),
      relatoriosApi.estoqueBaixo(),
    ])
    totalProdutos.value = p.data.data?.length ?? 0
    totalFornecedores.value = f.data.data?.length ?? 0
    movimentacoes.value = m.data.data ?? []
    estoqueBaixo.value = eb.data.data ?? []
  } finally {
    loading.value = false
  }
})
</script>
