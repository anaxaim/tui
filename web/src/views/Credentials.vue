<template>
  <div class="page_content">
    <el-dialog v-model="showCreate" top="5vh" title="Create Credential" width="60%">
      <el-form ref="createCredentialRef" :model="credential" label-position="top" label-width="auto">
        <div class="form_content">
          <el-form-item style="width: 50%;" label="Name" prop="name" required>
            <el-input v-model="credential.name" />
            <span>The name of your credential</span>
          </el-form-item>
          <el-form-item style="width: 50%; " label="Description" prop="description">
            <el-input v-model="credential.description" />
            <span>The description of your credential</span>
          </el-form-item>
        </div>
        <div>
          <el-button type="primary" style="margin-bottom: 1rem;" :icon="Plus" size="default" @click="addSecret()">Add secret</el-button>
          <div v-for="(secret, index) in credential.secrets" :key="index" class="form_content">
            <el-form-item>
              <el-button type="danger" :icon="Minus" @click="removeSecret(index)" />
            </el-form-item>
            <el-form-item style="width: 45%;" label="Name" :prop="`secrets.${index}.name`" required>
              <el-input v-model="secret.name" />
            </el-form-item>
            <el-form-item style="width: 45%;" label="Value" :prop="`secrets.${index}.value`">
              <el-input v-model="secret.value" />
            </el-form-item>
          </div>
        </div>
      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button type="success" @click="createCredential()">Create</el-button>
          <el-button type="danger" @click="showCreate = false">Cancel</el-button>
        </span>
      </template>
    </el-dialog>

    <div class="page_container">
      <PageHeader title="Credentials">
        <template #icon>
          <Protect style="margin-left: 1rem;" theme="outline" size="36" fill="#333" />
        </template>
      </PageHeader>
      <el-card>
        <template #header>
          <div class="card_content">
            <el-input class="search" v-model="search" placeholder="Type to search">
              <template #prefix>
                <el-icon>
                  <Search />
                </el-icon>
              </template>
            </el-input>
            <el-button type="success" :icon="Protect" @click="showCreate = true">Create</el-button>
          </div>
        </template>

        <el-table :data="tableData" :border="true" :header-row-style="{color:'#00000F'}" :header-cell-style="{'background-color':'#f6f6f5'}">
          <el-table-column prop="Operations" label="Operations" min-width="23px">
            <template #default="scope">
              <el-popover :visible="showDelete === scope.$index" :width="155" placement="top-start">
                <template #reference>
                  <el-button size="small" type="danger" circle @click="showDelete = scope.$index" :icon="Delete" />
                </template>
                <div style="margin-bottom: 0.5rem;">Delete this credential?</div>
                <span style="margin-left: 0.5rem;">
                  <el-button size="small" text @click="showDelete = -1">no</el-button>
                  <el-button size="small" type="danger" @click="deleteCredential(scope.row)">yes</el-button>
                </span>
              </el-popover>
            </template>
          </el-table-column>
          <el-table-column prop="name" label="Name" sortable min-width="70px" />
          <el-table-column prop="description" label="Description" />
          <el-table-column prop="createdAtString" label="CreatedAt" sortable min-width="65px" />
          <el-table-column prop="createdBy" label="CreatedBy" sortable min-width="65px" />
        </el-table>
      </el-card>
    </div>
  </div>
</template>

<script setup>
/*
  imports
*/
import {
  Protect,
  Search,
  Minus,
  Plus,
  Delete,
} from '@icon-park/vue-next';
import {
  computed, ref, onMounted, unref,
} from 'vue';
import { ElMessage } from 'element-plus';
import request from '@/axios';
import PageHeader from '@/components/PageHeader.vue';
import { getUser } from '@/utils';

/*
  modules list
*/
const credentials = ref([]);
onMounted(
  () => {
    request
      .get('/api/v1/credentials')
      .then((response) => {
        credentials.value = response.data.data;
      });
  },
);

/*
  search
*/
const search = ref('');
const tableData = computed(() => credentials.value.filter(
  (data) => !search.value || data.name.toLowerCase().includes(search.value.toLowerCase()),
));

/*
  create credential
*/
const showCreate = ref(false);
const createCredentialRef = ref();
const credential = ref({
  name: '',
  description: '',
  createdBy: '',
  secrets: [],
});

const createCredential = () => {
  const form = unref(createCredentialRef);
  if (!form) {
    return;
  }

  form.validate((valid) => {
    if (valid) {
      const user = getUser();
      request
        .post('/api/v1/credentials', {
          name: credential.value.name,
          description: credential.value.description,
          secrets: credential.value.secrets,
          createdBy: user.name,
        })
        .then((response) => {
          ElMessage.success('Create success');
          credentials.value.push(response.data.data);
          showCreate.value = false;
          form.resetFields();
          credential.value.secrets = [];
        })
        .catch((err) => {
          console.error('Create credential error:', err);
          ElMessage.error('Credential creation error');
        });
    } else {
      ElMessage.error('Input invalid');
    }
  });
};

/*
  secrets
*/
const addSecret = () => {
  const secret = { name: '', value: '' };
  credential.value.secrets.push(secret);
};

const removeSecret = (index) => {
  credential.value.secrets.splice(index, 1);
};

/*
  delete credential
*/
const showDelete = ref(-1);

const deleteCredential = (row) => {
  request
    .delete(`/api/v1/credentials/${row.id}`)
    .then(() => {
      ElMessage.success('Delete success');
      const index = credentials.value.findIndex((v) => v.id === row.id);
      credentials.value.splice(index, 1);
      showDelete.value = -1;
    })
    .catch((err) => {
      console.error('Delete credential error:', err);
      ElMessage.error('Delete credential error');
    });
};
</script>

<style lang="scss">
  .page_content {
    width: 100%;
    display: flex;
    justify-content: center;
  }

  .credentials {
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
    padding: 0.8rem 4rem 0.8rem 1rem;
  }
</style>
