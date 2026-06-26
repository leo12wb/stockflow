<template>
  <div style="padding:28px;">

    <div style="margin-bottom:24px;">
      <h1 style="margin:0 0 4px;font-size:22px;font-weight:700;color:#0f172a;">Relatórios</h1>
      <p style="margin:0;font-size:14px;color:#64748b;">Consulte dados do estoque por categoria</p>
    </div>

    <!-- Tabs -->
    <div style="display:flex;gap:8px;margin-bottom:20px;flex-wrap:wrap;">
      <button
        v-for="tab in tabs"
        :key="tab.key"
        @click="mudarAba(tab.key)"
        :style="`
          display:flex;align-items:center;gap:6px;padding:8px 16px;border-radius:8px;border:1px solid;
          font-size:13px;font-weight:500;cursor:pointer;transition:all .15s;
          background:${ativo===tab.key?'#6366f1':'#fff'};
          color:${ativo===tab.key?'#fff':'#475569'};
          border-color:${ativo===tab.key?'#6366f1':'#e2e8f0'};
        `"
      >
        <i :class="['pi', tab.icon]" style="font-size:13px;" />
        {{ tab.label }}
      </button>
    </div>

    <!-- Table card -->
    <div style="background:#fff;border-radius:12px;border:1px solid #e2e8f0;box-shadow:0 1px 3px rgba(0,0,0,.06);overflow:hidden;">
      <div style="padding:16px 20px;border-bottom:1px solid #f1f5f9;display:flex;align-items:center;justify-content:space-between;">
        <span style="font-weight:600;font-size:14px;color:#1e293b;">{{ tabAtual?.label }}</span>
        <span style="font-size:13px;color:#64748b;">{{ dados.length }} registro(s)</span>
      </div>

      <DataTable :value="dados" :loading="loading" striped-rows size="small" paginator :rows="15">

        <template v-if="isProdutoAba">
          <Column field="nome" header="Produto" sortable />
          <Column field="categoria" header="Categoria" style="width:140px;" />
          <Column header="Quantidade" style="width:110px;">
            <template #body="{ data }">
              <span
                :style="`font-weight:600;padding:3px 10px;border-radius:20px;font-size:12px;
                  background:${data.quantidade <= data.estoque_minimo ? '#fff7ed' : '#f0fdf4'};
                  color:${data.quantidade <= data.estoque_minimo ? '#ea580c' : '#16a34a'};`"
              >{{ data.quantidade }}</span>
            </template>
          </Column>
          <Column field="estoque_minimo" header="Mínimo" style="width:90px;" />
          <Column header="Preço Venda" style="width:130px;">
            <template #body="{ data }">{{ formatCurrency(data.preco_venda) }}</template>
          </Column>
          <Column field="fornecedor.nome" header="Fornecedor" />
        </template>

        <template v-else>
          <Column field="produto.nome" header="Produto" sortable />
          <Column header="Tipo" style="width:110px;">
            <template #body="{ data }">
              <span
                :style="`font-size:12px;font-weight:600;padding:3px 10px;border-radius:20px;border:1px solid;
                  background:${data.tipo==='ENTRADA'?'#f0fdf4':data.tipo==='SAIDA'?'#fff1f2':'#fffbeb'};
                  color:${data.tipo==='ENTRADA'?'#16a34a':data.tipo==='SAIDA'?'#dc2626':'#d97706'};
                  border-color:${data.tipo==='ENTRADA'?'#bbf7d0':data.tipo==='SAIDA'?'#fecaca':'#fde68a'};`"
              >{{ data.tipo }}</span>
            </template>
          </Column>
          <Column field="quantidade" header="Qtd" style="width:80px;" />
          <Column field="usuario.nome" header="Usuário" style="width:160px;" />
          <Column field="observacao" header="Observação" />
          <Column header="Data" style="width:150px;">
            <template #body="{ data }">{{ formatDate(data.created_at) }}</template>
          </Column>
        </template>

        <template #empty>
          <div style="text-align:center;color:#94a3b8;padding:40px;font-size:14px;">Nenhum dado encontrado</div>
        </template>
      </DataTable>
    </div>

  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { relatoriosApi } from '@/api/index'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'

const tabs = [
  { key: 'inventario',    label: 'Inventário',    icon: 'pi-list' },
  { key: 'estoque-baixo', label: 'Estoque Baixo', icon: 'pi-exclamation-triangle' },
  { key: 'entradas',      label: 'Entradas',      icon: 'pi-arrow-down' },
  { key: 'saidas',        label: 'Saídas',        icon: 'pi-arrow-up' },
  { key: 'mais-vendidos', label: 'Mais Vendidos', icon: 'pi-star' },
  { key: 'movimentacoes', label: 'Movimentações', icon: 'pi-arrows-h' },
]

const ativo = ref('inventario')
const dados = ref<any[]>([])
const loading = ref(false)

const tabAtual = computed(() => tabs.find(t => t.key === ativo.value))
const isProdutoAba = computed(() => ['inventario', 'estoque-baixo', 'mais-vendidos'].includes(ativo.value))

const formatCurrency = (v: number) => new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(v ?? 0)
const formatDate = (d: string) => new Date(d).toLocaleString('pt-BR')

const apiMap: Record<string, () => Promise<any>> = {
  'inventario':    relatoriosApi.inventario,
  'estoque-baixo': relatoriosApi.estoqueBaixo,
  'entradas':      relatoriosApi.entradas,
  'saidas':        relatoriosApi.saidas,
  'mais-vendidos': relatoriosApi.maisVendidos,
  'movimentacoes': relatoriosApi.movimentacoes,
}

async function mudarAba(key: string) {
  ativo.value = key; loading.value = true
  try {
    const res = await apiMap[key]()
    dados.value = res.data.data || []
  } finally { loading.value = false }
}

onMounted(() => mudarAba('inventario'))
</script>
