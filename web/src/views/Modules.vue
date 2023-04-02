<template>
  <div class="page_content">
    <el-dialog
      v-model="showCreate"
      top="5vh"
      title="Create Module"
      width="60%"
    >
      <el-form
        ref="createModuleRef"
        :model="newModule"
        label-position="top"
        label-width="auto"
      >
        <div class="form_content">
          <el-form-item
            style="width: 45%;"
            label="Name"
            prop="name"
            required
          >
            <el-input v-model="newModule.name" />
            <span>The name of your module</span>
          </el-form-item>
          <el-form-item
            style="width: 55%;"
            label="Description"
            prop="description"
          >
            <el-input
              v-model="newModule.description"
              type="textarea"
            />
            <span>The description of your module</span>
          </el-form-item>
        </div>
        <div class="form_content">
          <el-form-item
            style="width: 45%;"
            label="Main provider"
            prop="mainProvider"
            required
          >
            <el-input v-model="newModule.mainProvider" />
            <span>The main provider of this module</span>
          </el-form-item>
          <el-form-item
            style="width: 25%; margin-right: 5rem;"
            label="Provider version"
            prop="providerVersion"
          >
            <el-input v-model="newModule.providerVersion" />
            <span>The version of the provider</span>
          </el-form-item>
          <el-form-item
            style="width: 20%;"
            label="Terraform version"
            prop="terraformVersion"
          >
            <el-input v-model="newModule.terraformVersion" />
            <span>The version of Terraform</span>
          </el-form-item>
        </div>
        <div class="form_content">
          <el-form-item
            style="width: 60%;"
            label="Git Repository URL"
            prop="gitRepositoryUrl"
            required
          >
            <el-input v-model="newModule.gitRepositoryUrl" />
            <span>The URL of the module's git repository</span>
          </el-form-item>
          <el-form-item
            style="width: 40%;"
            label="Git repository directory"
            prop="directory"
          >
            <el-input v-model="newModule.directory" />
            <span>The sub-directory of the module's code inside the repository</span>
          </el-form-item>
        </div>
        <el-divider />
        <div>
          <el-button
            type="primary"
            style="margin-bottom: 1rem;"
            :icon="Plus"
            size="default"
            @click="addVariable"
          >
            Add variable
          </el-button>
          <div
            v-for="(variable, index) in newModule.variables"
            :key="index"
            class="form_content"
          >
            <el-form-item>
              <el-button
                type="danger"
                :icon="Minus"
                @click="removeVariable(index)"
              />
            </el-form-item>
            <el-form-item
              style="width: 25%;"
              label="Name"
              :prop="`variables.${index}.name`"
              required
            >
              <el-input v-model="variable.name" />
            </el-form-item>
            <el-form-item
              style="width: 25%;"
              label="Default Value"
              :prop="`variables.${index}.defaultValue`"
            >
              <el-input v-model="variable.defaultValue" />
            </el-form-item>
            <el-form-item
              style="width: 30%;"
              label="Description"
              :prop="`variables.${index}.description`"
            >
              <el-input v-model="variable.description" />
            </el-form-item>
            <el-form-item style="width: 10%; margin-bottom: 0.5rem;">
              <el-checkbox
                v-model="variable.editable"
                label="Editable"
                style="margin-bottom: -6px;"
              />
              <el-checkbox
                v-model="variable.mandatory"
                label="Mandatory"
                style="margin-top: -6px;"
              />
            </el-form-item>
          </div>
        </div>
      </el-form>

      <template #footer>
        <span class="dialog-footer">
          <el-button
            type="success"
            @click="createModule"
          >Create</el-button>
          <el-button
            type="danger"
            @click="showCreate = false"
          >Cancel</el-button>
        </span>
      </template>
    </el-dialog>

    <el-dialog
      v-model="showUpdate"
      top="5vh"
      title="Update Module"
      width="60%"
    >
      <el-form
        ref="updateFormRef"
        :model="updatedModule"
        label-position="top"
        label-width="auto"
      >
        <div class="form_content">
          <el-form-item
            style="width: 45%;"
            label="Name"
            prop="name"
            required
          >
            <el-input v-model="updatedModule.name" />
            <span>The name of your module</span>
          </el-form-item>
          <el-form-item
            style="width: 55%;"
            label="Description"
            prop="description"
          >
            <el-input
              v-model="updatedModule.description"
              type="textarea"
            />
            <span>The description of your module</span>
          </el-form-item>
        </div>
        <div class="form_content">
          <el-form-item
            style="width: 45%;"
            label="Main provider"
            prop="mainProvider"
          >
            <el-input
              v-model="updatedModule.mainProvider"
              disabled
            />
            <span>The main provider of this module</span>
          </el-form-item>
          <el-form-item
            style="width: 25%; margin-right: 5rem;"
            label="Provider version"
            prop="providerVersion"
          >
            <el-input v-model="updatedModule.providerVersion" />
            <span>The version of the provider</span>
          </el-form-item>
          <el-form-item
            style="width: 20%;"
            label="Terraform version"
            prop="terraformVersion"
          >
            <el-input v-model="updatedModule.terraformVersion" />
            <span>The version of Terraform</span>
          </el-form-item>
        </div>
        <div class="form_content">
          <el-form-item
            style="width: 60%;"
            label="Git Repository URL"
            prop="gitRepositoryUrl"
          >
            <el-input
              v-model="newModule.gitRepositoryUrl"
              disabled
            />
            <span>The URL of the module's git repository</span>
          </el-form-item>
          <el-form-item
            style="width: 40%;"
            label="Git repository directory"
            prop="directory"
          >
            <el-input v-model="newModule.directory" />
            <span>The sub-directory of the module's code inside the repository</span>
          </el-form-item>
        </div>
        <el-divider />
        <div>
          <el-button
            type="primary"
            style="margin-bottom: 1rem;"
            :icon="Plus"
            size="default"
            @click="addVariable"
          >
            Add variable
          </el-button>
          <div
            v-for="(variable, index) in updatedModule.variables"
            :key="index"
            class="form_content"
          >
            <el-form-item>
              <el-button
                type="danger"
                :icon="Minus"
                @click="removeVariable(index)"
              />
            </el-form-item>
            <el-form-item
              style="width: 25%;"
              label="Name"
              :prop="`variables.${index}.name`"
              required
            >
              <el-input v-model="variable.name" />
            </el-form-item>
            <el-form-item
              style="width: 25%;"
              label="Default Value"
              :prop="`variables.${index}.defaultValue`"
            >
              <el-input v-model="variable.defaultValue" />
            </el-form-item>
            <el-form-item
              style="width: 30%;"
              label="Description"
              :prop="`variables.${index}.description`"
            >
              <el-input v-model="variable.description" />
            </el-form-item>
            <el-form-item style="width: 10%; margin-bottom: 0.5rem;">
              <el-checkbox
                v-model="variable.editable"
                label="Editable"
                style="margin-bottom: -6px;"
              />
              <el-checkbox
                v-model="variable.mandatory"
                label="Mandatory"
                style="margin-top: -6px;"
              />
            </el-form-item>
          </div>
        </div>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button
            type="success"
            @click="updateModule()"
          >Update</el-button>
          <el-button
            type="danger"
            @click="showUpdate = false"
          >Cancel</el-button>
        </span>
      </template>
    </el-dialog>

    <div class="page_container">
      <PageHeader title="Modules">
        <template #icon>
          <CloudStorage
            style="margin-left: 1rem;"
            theme="outline"
            size="36"
            fill="#333"
          />
        </template>
      </PageHeader>
      <el-card>
        <template #header>
          <div class="card_content">
            <el-input
              v-model="search"
              class="search"
              placeholder="Type to search"
            >
              <template #prefix>
                <el-icon>
                  <Search />
                </el-icon>
              </template>
            </el-input>
            <el-button
              type="success"
              :icon="CloudStorage"
              @click="showCreate = true"
            >
              Create
            </el-button>
          </div>
        </template>

        <el-table
          :data="filtration"
          height="360"
        >
          <el-table-column
            prop="name"
            label="Name"
            sortable
            min-width="50px"
          >
            <template #default="{row}">
              <el-popover
                placement="top-end"
                :width="200"
                :disabled="!row.description"
              >
                <template #reference>
                  <div>{{ row.name }}</div>
                </template>
                <p>{{ row.description }}</p>
              </el-popover>
            </template>
          </el-table-column>
          <el-table-column
            prop="registryDetails.projectId"
            label="Registry"
            min-width="150px"
          >
            <template #default="{row}">
              <span>
                <el-button link>
                  <a
                    :href="row.gitRepositoryUrl"
                    target="_blank"
                  >
                    <GithubOne
                      v-if="row.registryDetails.registryType === 'github'"
                      size="20"
                    />
                    <Gitlab
                      v-else-if="row.registryDetails.registryType === 'gitlab'"
                      theme="two-tone"
                      :fill="['#333' ,'#f6650a']"
                      :stroke-width="2"
                      size="20"
                    />
                  </a>
                </el-button>
                {{ row.registryDetails.projectId }}
              </span>
            </template>
          </el-table-column>
          <el-table-column
            prop="mainProvider"
            label="Provider"
          />
          <el-table-column
            prop="createdAtString"
            label="Created"
            sortable
            min-width="80px"
          />
          <el-table-column
            prop="updatedAtString"
            label="Updated"
            sortable
            min-width="80px"
          />
          <el-table-column
            prop="Operations"
            label="Operations"
            min-width="80px"
          >
            <template #default="scope">
              <el-button
                class="operation_icon"
                type="success"
                size="small"
                circle
                :icon="PlayOne"
                @click="importRegistry(scope.row.id)"
              />
              <el-button
                class="operation_icon"
                type="warning"
                size="small"
                circle
                :icon="Edit"
                @click="editModule(scope.row)"
              />
              <el-button
                class="operation_icon"
                size="small"
                circle
                :icon="Log"
              />
              <el-popover
                :visible="showDelete === scope.$index"
                placement="top-start"
              >
                <template #reference>
                  <el-button
                    size="small"
                    type="danger"
                    circle
                    :icon="Delete"
                    @click="showDelete = scope.$index"
                  />
                </template>
                <div style="margin-bottom: 0.5rem;">
                  Delete this module?
                </div>
                <span style="margin-left: 0.5rem;">
                  <el-button
                    size="small"
                    text
                    @click="showDelete = -1"
                  >no</el-button>
                  <el-button
                    size="small"
                    type="danger"
                    @click="deleteModule(scope.row)"
                  >yes</el-button>
                </span>
              </el-popover>
            </template>
          </el-table-column>
          <el-table-column
            prop="status"
            label="Status"
            min-width="90px"
          >
            <template #default="scope">
              <el-tag :type="getStatusType(scope.row.status)">
                {{ scope.row.status }}
              </el-tag>
            </template>
          </el-table-column>
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
  CloudStorage, Delete, Edit, Log, PlayOne, Search, GithubOne, Gitlab, Plus, Minus,
} from '@icon-park/vue-next';

import {
  computed, onMounted, ref, unref,
} from 'vue';
import { ElMessage } from 'element-plus';
import request from '@/axios';
import PageHeader from '@/components/PageHeader.vue';

/*
  modules list
*/
const modules = ref([]);
onMounted(
  () => {
    request
      .get('/api/v1/modules')
      .then((response) => {
        modules.value = response.data.data;
      });
  },
);

/*
  filtration and search
*/
const search = ref('');
const filtration = computed(() => modules.value.filter(
  (data) => !search.value || data.name.toLowerCase().includes(search.value.toLowerCase()),
));

/*
  create module
*/
const showCreate = ref(false);
const createModuleRef = ref();
const newModule = ref({
  name: '',
  description: '',
  provider: '',
  gitRepositoryUrl: '',
  directory: '',
  mainProvider: '',
  providerVersion: '',
  terraformVersion: '',
  registryDetails: {
    registryType: '',
    projectId: '',
  },
  variables: [],
});

const createModule = () => {
  const form = unref(createModuleRef);
  if (!form) {
    return;
  }

  form.validate((valid) => {
    if (valid) {
      request
        .post('/api/v1/modules', {
          name: newModule.value.name,
          description: newModule.value.description,
          gitRepositoryUrl: newModule.value.gitRepositoryUrl,
          directory: newModule.value.directory,
          mainProvider: newModule.value.mainProvider,
          providerVersion: newModule.value.providerVersion,
          terraformImage: {
            repository: 'hashicorp/terraform',
            tag: newModule.value.terraformVersion,
          },
          variables: newModule.value.variables,
        })
        .then((response) => {
          ElMessage.success('Create success');
          modules.value.push(response.data.data);
          showCreate.value = false;
          form.resetFields();
          newModule.value.variables = [];
        })
        .catch((err) => {
          console.error('Create module error:', err);
          ElMessage.error('Module creation error');
        });
    } else {
      ElMessage.error('Input invalid');
    }
  });
};
/*
  variables
*/
const addVariable = () => {
  newModule.value.variables.push({
    name: '',
    description: '',
    defaultValue: '',
    editable: false,
    mandatory: false,
  });
};

const removeVariable = (index) => {
  newModule.value.variables.splice(index, 1);
};

/*
  delete module
*/
const showDelete = ref(-1);

const deleteModule = (row) => {
  request
    .delete(`/api/v1/modules/${row.id}`)
    .then(() => {
      ElMessage.success('Delete success');
      const index = modules.value.findIndex((v) => v.id === row.id);
      modules.value.splice(index, 1);
      showDelete.value = -1;
    })
    .catch((err) => {
      console.error('Delete module error:', err);
      ElMessage.error('Delete module error');
    });
};
/*
 edit module
*/
const updateFormRef = ref();
const showUpdate = ref(false);
const updatedModule = ref({
  id: '',
  name: '',
  description: '',
  provider: '',
  gitRepositoryUrl: '',
  directory: '',
  mainProvider: '',
  providerVersion: '',
  terraformVersion: '',
  registryDetails: {
    registryType: '',
    projectId: '',
  },
  variables: [],
});

const editModule = (row) => {
  updatedModule.value = { ...row };
  showUpdate.value = true;
};

const updateModule = () => {
  const form = unref(updateFormRef);
  if (!form) {
    return;
  }

  form.validate((valid, err) => {
    if (valid) {
      request
        .put(`/api/v1/modules/${updatedModule.value.id}`, updatedModule.value)
        .then((response) => {
          ElMessage.success('Update success');
          const index = modules.value.findIndex((v) => v.id === updatedModule.value.id);
          updatedModule.value.status = response.data.data.status;
          updatedModule.value.updatedAtString = response.data.data.updatedAtString;
          modules.value[index] = updatedModule.value;
          showUpdate.value = false;
        })
        .catch((error) => {
          console.error('Update module error:', error);
          ElMessage.error('Update module error');
        });
    } else {
      ElMessage.error('Input invalid', err);
    }
  });
};

/*
  import module
*/

const importRegistry = (id) => {
  request
    .post(`/api/v1/registry/import/${id}`)
    .then(() => {
      ElMessage.success('Launched module import');
      const index = modules.value.findIndex((v) => v.id === id);
      if (index !== -1) {
        modules.value[index].status = 'RUNNING';
        modules.value = [...modules.value];
      }
    })
    .catch((err) => {
      console.error('Import module error:', err);
      ElMessage.error('Import module error');
    });
};

/*
  status
*/
const getStatusType = (status) => {
  if (status === 'RUNNING') {
    return 'success';
  } if (status === 'CREATED' || status === 'UPDATED') {
    return 'warning';
  } if (status === 'ERROR') {
    return 'danger';
  }
  return 'info';
};
</script>

<style lang="scss">
  .page_content {
    width: 100%;
    display: flex;
    justify-content: center;
  }

  .page_container {
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
    padding: 0.8rem 4rem 0.8rem 1rem;
  }

  .card_content {
    display: flex;
    justify-content: space-between;
  }

  .search {
    margin-right: 1rem;
  }

  .form_content {
    display: flex;
    flex-direction: row;
    width: 100%;
    gap: 10px;
    align-items: flex-end;
  }

  .operation_icon {
    margin-left: 0.5rem;
  }
</style>
