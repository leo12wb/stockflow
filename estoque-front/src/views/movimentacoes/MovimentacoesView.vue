<template>
  <div style="padding:28px;">

    <div style="display:flex;align-items:center;justify-content:space-between;margin-bottom:24px;">
      <div>
        <h1 style="margin:0 0 4px;font-size:22px;font-weight:700;color:#0f172a;">Movimentações</h1>
        <p style="margin:0;font-size:14px;color:#64748b;">Histórico de entradas e saídas</p>
      </div>
      <Button label="Nova Movimentação" icon="pi pi-plus" @click="formVisible = true" />
    </div>

    <div style="background:#fff;border-radius:12px;border:1px solid #e2e8f0;box-shadow:0 1px 3px rgba(0,0,0,.06);overflow:hidden;">
      <DataTable
        :value="movimentacoes" :loading="loading" striped-rows size="small"
        lazy paginator :rows="perPage" :totalRecords="total"
        @page="onPage"
      >
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
        <template #empty>
          <div style="text-align:center;color:#94a3b8;padding:40px;font-size:14px;">Nenhuma movimentação encontrada</div>
        </template>
      </DataTable>
    </div>

    <Dialog v-model:visible="formVisible" header="Nova Movimentação" modal style="width:420px;">
      <form @submit.prevent="salvar" style="display:flex;flex-direction:column;gap:16px;padding-top:8px;">
        <div>
          <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Produto *</label>
          <Select v-model="formData.produto_id" :options="produtos" optionLabel="nome" optionValue="id" placeholder="Selecione..." :invalid="!!formErrors.produto_id" filter style="width:100%;" />
          <small v-if="formErrors.produto_id" style="color:#ef4444;font-size:12px;">{{ formErrors.produto_id }}</small>
        </div>
        <div>
          <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Tipo *</label>
          <Select v-model="formData.tipo" :options="tipos" placeholder="Selecione..." :invalid="!!formErrors.tipo" style="width:100%;" />
          <small v-if="formErrors.tipo" style="color:#ef4444;font-size:12px;">{{ formErrors.tipo }}</small>
        </div>
        <div>
          <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Quantidade *</label>
          <InputNumber v-model="formData.quantidade" :min="1" :invalid="!!formErrors.quantidade" style="width:100%;" />
          <small v-if="formErrors.quantidade" style="color:#ef4444;font-size:12px;">{{ formErrors.quantidade }}</small>
        </div>
        <div>
          <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Observação</label>
          <Textarea v-model="formData.observacao" rows="2" style="width:100%;" />
        </div>
        <div style="display:flex;justify-content:flex-end;gap:8px;padding-top:4px;">
          <Button label="Cancelar" severity="secondary" @click="fecharForm" type="button" />
          <Button label="Registrar" icon="pi pi-check" type="submit" :loading="salvando" />
        </div>
      </form>
    </Dialog>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useToast } from 'primevue/usetoast'
import { movimentacoesApi, produtosApi } from '@/api/index'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import Select from 'primevue/select'
import InputNumber from 'primevue/inputnumber'
import Textarea from 'primevue/textarea'

const toast = useToast()
const movimentacoes = ref<any[]>([])
const produtos = ref<any[]>([])
const loading = ref(false)
const page = ref(1)
const perPage = ref(10)
const total = ref(0)
const formVisible = ref(false)
const salvando = ref(false)
const formErrors = ref<any>({})
const tipos = ['ENTRADA', 'SAIDA', 'AJUSTE']
const formData = ref({ produto_id: null as number | null, tipo: '', quantidade: 1, observacao: '' })

const formatDate = (d: string) => new Date(d).toLocaleString('pt-BR')

async function carregar() {
  loading.value = true
  try {
    const [movRes, prodRes] = await Promise.all([
      movimentacoesApi.getAll({ page: page.value, per_page: perPage.value }),
      produtosApi.getAll({ per_page: 100 }),
    ])
    movimentacoes.value = movRes.data.data || []
    total.value = movRes.data.total ?? 0
    produtos.value = prodRes.data.data || []
  } finally { loading.value = false }
}

function onPage(event: any) {
  page.value = event.page + 1
  perPage.value = event.rows
  carregar()
}

function fecharForm() {
  formVisible.value = false; formErrors.value = {}
  formData.value = { produto_id: null, tipo: '', quantidade: 1, observacao: '' }
}

async function salvar() {
  formErrors.value = {}; salvando.value = true
  try {
    await movimentacoesApi.create(formData.value)
    toast.add({ severity: 'success', summary: 'Movimentação registrada!', life: 3000 })
    fecharForm(); carregar()
  } catch (err: any) {
    if (err.response?.data?.errors) {
      const e = err.response.data.errors
      Object.keys(e).forEach(k => formErrors.value[k] = e[k][0])
    } else {
      toast.add({ severity: 'error', summary: err.response?.data?.error || 'Erro ao registrar', life: 3000 })
    }
  } finally { salvando.value = false }
}

onMounted(carregar)
</script>
