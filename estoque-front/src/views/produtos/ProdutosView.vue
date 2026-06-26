<template>
  <div style="padding:28px;">

    <!-- Header -->
    <div style="display:flex;align-items:center;justify-content:space-between;margin-bottom:24px;">
      <div>
        <h1 style="margin:0 0 4px;font-size:22px;font-weight:700;color:#0f172a;">Produtos</h1>
        <p style="margin:0;font-size:14px;color:#64748b;">{{ produtos.length }} produto(s) cadastrado(s)</p>
      </div>
      <Button label="Novo Produto" icon="pi pi-plus" @click="openNew()" />
    </div>

    <!-- Table card -->
    <div style="background:#fff;border-radius:12px;border:1px solid #e2e8f0;box-shadow:0 1px 3px rgba(0,0,0,.06);overflow:hidden;">

      <!-- Search bar -->
      <div style="padding:16px 20px;border-bottom:1px solid #f1f5f9;display:flex;gap:10px;">
        <InputText v-model="search" placeholder="Buscar produto..." style="flex:1;" @keyup.enter="load()" />
        <Button icon="pi pi-search" @click="load()" />
        <Button icon="pi pi-refresh" severity="secondary" @click="load()" :loading="loading" />
      </div>

      <DataTable
        :value="produtos" :loading="loading" striped-rows size="small"
        lazy paginator :rows="perPage" :totalRecords="total"
        @page="onPage"
      >
        <Column field="codigo_barras" header="Código" style="width:130px;" />
        <Column field="nome" header="Nome" />
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
        <Column field="estoque_minimo" header="Mín." style="width:70px;" />
        <Column header="Preço Compra" style="width:120px;">
          <template #body="{ data }">{{ formatCurrency(data.preco_compra) }}</template>
        </Column>
        <Column header="Status" style="width:90px;">
          <template #body="{ data }">
            <Tag
              :value="data.quantidade <= data.estoque_minimo ? 'Baixo' : 'OK'"
              :severity="data.quantidade <= data.estoque_minimo ? 'warn' : 'success'"
            />
          </template>
        </Column>
        <Column header="Ações" style="width:100px;">
          <template #body="{ data }">
            <div style="display:flex;gap:4px;">
              <Button icon="pi pi-pencil" size="small" text @click="openEdit(data)" />
              <Button icon="pi pi-trash" size="small" text severity="danger" @click="confirmDelete(data)" />
            </div>
          </template>
        </Column>
        <template #empty>
          <div style="text-align:center;color:#94a3b8;padding:40px;font-size:14px;">Nenhum produto encontrado</div>
        </template>
      </DataTable>
    </div>

    <!-- Dialog -->
    <Dialog v-model:visible="dialogVisible" :header="editing ? 'Editar Produto' : 'Novo Produto'" modal style="width:580px;">
      <form @submit.prevent="save()" style="display:flex;flex-direction:column;gap:16px;padding-top:8px;">
        <div style="display:grid;grid-template-columns:1fr 1fr;gap:14px;">
          <div>
            <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Código de Barras</label>
            <InputText v-model="form.codigo_barras" placeholder="EAN13..." style="width:100%;" />
            <small style="color:#ef4444;font-size:12px;">{{ fieldError('codigo_barras') }}</small>
          </div>
          <div>
            <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Nome *</label>
            <InputText v-model="form.nome" placeholder="Nome do produto" style="width:100%;" />
            <small style="color:#ef4444;font-size:12px;">{{ fieldError('nome') }}</small>
          </div>
        </div>
        <div>
          <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Descrição</label>
          <Textarea v-model="form.descricao" rows="2" placeholder="Descrição..." style="width:100%;" />
        </div>
        <div style="display:grid;grid-template-columns:1fr 1fr;gap:14px;">
          <div>
            <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Categoria</label>
            <InputText v-model="form.categoria" placeholder="Categoria" style="width:100%;" />
            <small style="color:#ef4444;font-size:12px;">{{ fieldError('categoria') }}</small>
          </div>
          <div>
            <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Fornecedor</label>
            <Select v-model="form.fornecedor_id" :options="fornecedores" option-label="nome" option-value="id" placeholder="Selecione" filter style="width:100%;" />
          </div>
        </div>
        <div style="display:grid;grid-template-columns:1fr 1fr 1fr;gap:14px;">
          <div>
            <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Quantidade</label>
            <InputNumber v-model="form.quantidade" :min="0" style="width:100%;" />
          </div>
          <div>
            <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Estoque Mín.</label>
            <InputNumber v-model="form.estoque_minimo" :min="0" style="width:100%;" />
          </div>
          <div>
            <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Preço de Compra</label>
            <InputNumber v-model="form.preco_compra" mode="currency" currency="BRL" locale="pt-BR" style="width:100%;" />
          </div>
        </div>
        <Message v-if="saveError" severity="error" :closable="false">{{ saveError }}</Message>
      </form>
      <template #footer>
        <Button label="Cancelar" severity="secondary" @click="dialogVisible = false" />
        <Button label="Salvar" icon="pi pi-check" :loading="saving" @click="save()" />
      </template>
    </Dialog>

  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import InputNumber from 'primevue/inputnumber'
import Textarea from 'primevue/textarea'
import Select from 'primevue/select'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Dialog from 'primevue/dialog'
import Tag from 'primevue/tag'
import Message from 'primevue/message'
import { useConfirm } from 'primevue/useconfirm'
import { useToast } from 'primevue/usetoast'
import { produtosApi, fornecedoresApi } from '@/api'

const confirm = useConfirm()
const toast = useToast()
const loading = ref(false)
const saving = ref(false)
const produtos = ref<any[]>([])
const fornecedores = ref<any[]>([])
const search = ref('')
const page = ref(1)
const perPage = ref(10)
const total = ref(0)
const dialogVisible = ref(false)
const editing = ref<any>(null)
const saveError = ref('')
const validationErrors = ref<Record<string, string[]>>({})

const emptyForm = () => ({
  codigo_barras: '', nome: '', descricao: '', categoria: '',
  quantidade: 0, estoque_minimo: 5,
  preco_compra: 0, preco_venda: 0, fornecedor_id: null as number | null,
})

const form = ref(emptyForm())
const fieldError = (f: string) => validationErrors.value[f]?.[0] ?? ''
const formatCurrency = (v: number) => new Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(v ?? 0)

async function load() {
  loading.value = true
  try {
    const { data } = await produtosApi.getAll({ page: page.value, per_page: perPage.value, search: search.value || undefined })
    produtos.value = data.data ?? []
    total.value = data.total ?? 0
  } finally { loading.value = false }
}

function onPage(event: any) {
  page.value = event.page + 1
  perPage.value = event.rows
  load()
}

async function loadFornecedores() {
  const { data } = await fornecedoresApi.getAll()
  fornecedores.value = data.data ?? []
}

function openNew() {
  editing.value = null; form.value = emptyForm()
  saveError.value = ''; validationErrors.value = {}; dialogVisible.value = true
}

function openEdit(item: any) {
  editing.value = item
  form.value = {
    codigo_barras: item.codigo_barras ?? '', nome: item.nome,
    descricao: item.descricao ?? '', categoria: item.categoria ?? '',
    quantidade: item.quantidade, estoque_minimo: item.estoque_minimo,
    preco_compra: item.preco_compra, preco_venda: item.preco_venda,
    fornecedor_id: item.fornecedor_id ?? null,
  }
  saveError.value = ''; validationErrors.value = {}; dialogVisible.value = true
}

async function save() {
  saving.value = true; saveError.value = ''; validationErrors.value = {}
  try {
    if (editing.value) {
      await produtosApi.update(editing.value.id, form.value)
      toast.add({ severity: 'success', summary: 'Produto atualizado!', life: 3000 })
    } else {
      await produtosApi.create(form.value)
      toast.add({ severity: 'success', summary: 'Produto criado!', life: 3000 })
    }
    dialogVisible.value = false; load()
  } catch (err: any) {
    if (err.response?.data?.errors) validationErrors.value = err.response.data.errors
    else saveError.value = err.response?.data?.error || 'Erro ao salvar'
  } finally { saving.value = false }
}

function confirmDelete(item: any) {
  confirm.require({
    message: `Deseja remover o produto "${item.nome}"?`,
    header: 'Confirmar exclusão', icon: 'pi pi-trash',
    acceptProps: { label: 'Excluir', severity: 'danger' },
    rejectProps: { label: 'Cancelar', severity: 'secondary' },
    accept: async () => {
      await produtosApi.delete(item.id)
      toast.add({ severity: 'info', summary: 'Produto removido!', life: 3000 })
      load()
    },
  })
}

onMounted(() => { load(); loadFornecedores() })
</script>
