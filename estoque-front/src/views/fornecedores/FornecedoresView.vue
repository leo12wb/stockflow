<template>
  <div style="padding:28px;">

    <div style="display:flex;align-items:center;justify-content:space-between;margin-bottom:24px;">
      <div>
        <h1 style="margin:0 0 4px;font-size:22px;font-weight:700;color:#0f172a;">Fornecedores</h1>
        <p style="margin:0;font-size:14px;color:#64748b;">{{ fornecedores.length }} fornecedor(es) cadastrado(s)</p>
      </div>
      <Button label="Novo Fornecedor" icon="pi pi-plus" @click="abrirForm()" />
    </div>

    <div style="background:#fff;border-radius:12px;border:1px solid #e2e8f0;box-shadow:0 1px 3px rgba(0,0,0,.06);overflow:hidden;">
      <DataTable
        :value="fornecedores" :loading="loading" striped-rows size="small"
        lazy paginator :rows="perPage" :totalRecords="total"
        @page="onPage"
      >
        <Column field="nome" header="Nome" sortable />
        <Column field="cnpj" header="CNPJ" style="width:160px;" />
        <Column field="email" header="E-mail" />
        <Column field="telefone" header="Telefone" style="width:140px;" />
        <Column field="endereco" header="Endereço" />
        <Column header="Ações" style="width:100px;">
          <template #body="{ data }">
            <div style="display:flex;gap:4px;">
              <Button icon="pi pi-pencil" size="small" text @click="abrirForm(data)" />
              <Button icon="pi pi-trash" size="small" text severity="danger" @click="confirmarDelete(data)" />
            </div>
          </template>
        </Column>
        <template #empty>
          <div style="text-align:center;color:#94a3b8;padding:40px;font-size:14px;">Nenhum fornecedor encontrado</div>
        </template>
      </DataTable>
    </div>

    <Dialog v-model:visible="formVisible" :header="formData.id ? 'Editar Fornecedor' : 'Novo Fornecedor'" modal style="width:480px;">
      <form @submit.prevent="salvar" style="display:flex;flex-direction:column;gap:16px;padding-top:8px;">
        <div>
          <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Nome *</label>
          <InputText v-model="formData.nome" :invalid="!!formErrors.nome" style="width:100%;" />
          <small v-if="formErrors.nome" style="color:#ef4444;font-size:12px;">{{ formErrors.nome }}</small>
        </div>
        <div>
          <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">CNPJ *</label>
          <InputText v-model="formData.cnpj" :invalid="!!formErrors.cnpj" style="width:100%;" />
          <small v-if="formErrors.cnpj" style="color:#ef4444;font-size:12px;">{{ formErrors.cnpj }}</small>
        </div>
        <div style="display:grid;grid-template-columns:1fr 1fr;gap:14px;">
          <div>
            <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">E-mail</label>
            <InputText v-model="formData.email" type="email" style="width:100%;" />
          </div>
          <div>
            <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Telefone</label>
            <InputText v-model="formData.telefone" style="width:100%;" />
          </div>
        </div>
        <div>
          <label style="display:block;font-size:13px;font-weight:600;color:#374151;margin-bottom:5px;">Endereço</label>
          <InputText v-model="formData.endereco" style="width:100%;" />
        </div>
        <div style="display:flex;justify-content:flex-end;gap:8px;padding-top:4px;">
          <Button label="Cancelar" severity="secondary" @click="formVisible = false" type="button" />
          <Button label="Salvar" icon="pi pi-check" type="submit" :loading="salvando" />
        </div>
      </form>
    </Dialog>

    <ConfirmDialog />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useConfirm } from 'primevue/useconfirm'
import { useToast } from 'primevue/usetoast'
import { fornecedoresApi } from '@/api/index'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import InputText from 'primevue/inputtext'
import ConfirmDialog from 'primevue/confirmdialog'

const confirm = useConfirm()
const toast = useToast()
const fornecedores = ref<any[]>([])
const loading = ref(false)
const page = ref(1)
const perPage = ref(10)
const total = ref(0)
const formVisible = ref(false)
const salvando = ref(false)
const formErrors = ref<any>({})
const formData = ref({ id: null as number | null, nome: '', cnpj: '', email: '', telefone: '', endereco: '' })

async function carregar() {
  loading.value = true
  try {
    const res = await fornecedoresApi.getAll({ page: page.value, per_page: perPage.value })
    fornecedores.value = res.data.data || []
    total.value = res.data.total ?? 0
  } finally { loading.value = false }
}

function onPage(event: any) {
  page.value = event.page + 1
  perPage.value = event.rows
  carregar()
}

function abrirForm(f?: any) {
  formErrors.value = {}
  formData.value = f ? { ...f } : { id: null, nome: '', cnpj: '', email: '', telefone: '', endereco: '' }
  formVisible.value = true
}

async function salvar() {
  formErrors.value = {}; salvando.value = true
  try {
    if (formData.value.id) {
      await fornecedoresApi.update(formData.value.id, formData.value)
      toast.add({ severity: 'success', summary: 'Fornecedor atualizado!', life: 3000 })
    } else {
      await fornecedoresApi.create(formData.value)
      toast.add({ severity: 'success', summary: 'Fornecedor criado!', life: 3000 })
    }
    formVisible.value = false; carregar()
  } catch (err: any) {
    if (err.response?.data?.errors) {
      const e = err.response.data.errors
      Object.keys(e).forEach(k => formErrors.value[k] = e[k][0])
    } else {
      toast.add({ severity: 'error', summary: err.response?.data?.error || 'Erro ao salvar', life: 3000 })
    }
  } finally { salvando.value = false }
}

function confirmarDelete(f: any) {
  confirm.require({
    message: `Deseja remover "${f.nome}"?`,
    header: 'Confirmar remoção', icon: 'pi pi-trash',
    acceptProps: { label: 'Remover', severity: 'danger' },
    rejectProps: { label: 'Cancelar', severity: 'secondary' },
    accept: async () => {
      await fornecedoresApi.delete(f.id)
      toast.add({ severity: 'success', summary: 'Fornecedor removido!', life: 3000 })
      carregar()
    }
  })
}

onMounted(carregar)
</script>
