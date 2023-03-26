<template>
  <div class="page_content">
    <el-dialog v-model="showCreate" top="5vh" title="Create Module" width="60%">
      <el-form ref="createModuleRef" :model="newModule" label-position="top" label-width="auto">
        <div class="form_content">
          <el-form-item style="width: 45%;" label="Name" prop="name" required>
            <el-input v-model="newModule.name" />
            <span>The name of your module</span>
          </el-form-item>
          <el-form-item style="width: 55%;" label="Description" prop="description">
            <el-input v-model="newModule.description" type="textarea" />
            <span>The description of your module</span>
          </el-form-item>
        </div>

        <div class="form_content">
          <el-form-item style="width: 45%;" label="Main provider" prop="mainProvider" required>
            <el-input v-model="newModule.mainProvider" />
            <span>The main provider of this module</span>
          </el-form-item>
          <el-form-item style="width: 25%; margin-right: 5rem;" label="Provider version" prop="providerVersion">
            <el-input v-model="newModule.providerVersion" />
            <span>The version of the provider</span>
          </el-form-item>
          <el-form-item style="width: 20%;" label="Terraform version" prop="terraformVersion">
            <el-input v-model="newModule.terraformVersion" />
            <span>The version of Terraform</span>
          </el-form-item>
        </div>
        <div class="form_content">
          <el-form-item style="width: 60%;" label="Git Repository URL" prop="gitRepositoryUrl" required>
            <el-input v-model="newModule.gitRepositoryUrl" />
            <span>The URL of the module's git repository</span>
          </el-form-item>
          <el-form-item style="width: 40%;" label="Git repository directory" prop="directory">
            <el-input v-model="newModule.directory" />
            <span>The sub-directory of the module's code inside the repository</span>
          </el-form-item>
        </div>
      </el-form>
      <template #footer>
          <span class="dialog-footer">
            <el-button type="success" @click="createModule">Create</el-button>
            <el-button type="danger" @click="showCreate = false">Cancel</el-button>
          </span>
      </template>
    </el-dialog>

    <el-dialog v-model="showUpdate" top="5vh" title="Update Module" width="60%">
      <el-form ref="updateFormRef" :model="updatedModule" label-position="top" label-width="auto">
        <div class="form_content">
          <el-form-item style="width: 45%;" label="Name" prop="name" required>
            <el-input v-model="updatedModule.name" disabled/>
          </el-form-item>
          <el-form-item style="width: 55%;" label="Description" prop="description">
            <el-input v-model="updatedModule.description" type="textarea" />
            <span>The description of your module</span>
          </el-form-item>
        </div>
        <div class="form_content">
          <el-form-item style="width: 45%;" label="Main provider" prop="mainProvider">
            <el-input v-model="updatedModule.mainProvider" disabled/>
            <span>The main provider of this module</span>
          </el-form-item>
          <el-form-item style="width: 25%; margin-right: 5rem;" label="Provider version" prop="providerVersion">
            <el-input v-model="updatedModule.providerVersion" disabled/>
            <span>The version of the provider</span>
          </el-form-item>
          <el-form-item style="width: 20%;" label="Terraform version" prop="terraformVersion">
            <el-input v-model="updatedModule.terraformVersion" />
            <span>The version of Terraform</span>
          </el-form-item>
        </div>
        <div class="form_content">
          <el-form-item style="width: 60%;" label="Git Repository URL" prop="gitRepositoryUrl">
            <el-input v-model="newModule.gitRepositoryUrl" disabled/>
            <span>The URL of the module's git repository</span>
          </el-form-item>
          <el-form-item style="width: 40%;" label="Git repository directory" prop="directory">
            <el-input v-model="newModule.directory" disabled/>
            <span>The sub-directory of the module's code inside the repository</span>
          </el-form-item>
        </div>

      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button type="success" @click="updateModule()">Save</el-button>
          <el-button type="danger" @click="showUpdate = false">Cancel</el-button>
        </span>
      </template>
    </el-dialog>

    <div class="page_container">
      <PageHeader title="Modules">
        <template #icon>
          <CloudStorage style="margin-left: 1rem;" theme="outline" size="36" fill="#333" />
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
            <el-button type="success" :icon="CloudStorage" @Click="showCreate = true">Create</el-button>
          </div>
        </template>

        <el-table :data="filter" height="360">
          <el-table-column prop="name" label="Name" sortable min-width="40px"/>
          <el-table-column prop="gitRepositoryUrl" label="Repository" min-width="150px" />
          <el-table-column prop="mainProvider" label="Provider" />
          <el-table-column prop="createdAt" label="Created" sortable min-width="80px" />
          <el-table-column prop="updatedAt" label="Updated" sortable min-width="80px" />
          <el-table-column prop="Operations" label="Operations" min-width="100px">
            <template #default="scope">
              <el-button class="operation_icon" type="success" size="small" circle :icon="PlayOne" />
              <el-button class="operation_icon" type="warning" size="small" circle @click="editModule(scope.row)" :icon="Edit" />
              <el-button class="operation_icon" size="small" circle :icon="Log" />
              <el-popover :visible="showDelete === scope.$index" placement="top" :width="140">
                <template #reference>
                  <el-button size="small" type="danger" circle @click="showDelete = scope.$index" :icon="Delete" />
                </template>
                <p>Delete this module?</p>
                <span style="margin-left: 0.5rem;">
                  <el-button size="small" text @click="showDelete = -1">no</el-button>
                  <el-button size="small" type="danger" @click="deleteModule(scope.row)">yes</el-button>
                </span>
              </el-popover>
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
import PageHeader from '@/components/PageHeader.vue'
import {CloudStorage, Delete, Edit, Log, PlayOne, Search} from '@icon-park/vue-next'

import {computed, onMounted, ref, unref} from 'vue';
import request from '@/axios'
import {ElMessage} from "element-plus";

/*
  filtration and search
*/
  const search = ref('');
  const filter = computed(() =>
      modules.value.filter(
          (data) => !search.value || data.name.toLowerCase().includes(search.value.toLowerCase())
      )
  )

/*
  modules list
*/
  const modules = ref([]);
  onMounted(
      () => {
        request.get("/api/v1/modules").then((response) => {
          modules.value = response.data.data;
        })
      }
  )

/*
  create module
*/
  const showCreate = ref(false);
  const defaultTime = { timeout: "10000" }
  const createModuleRef = ref();
  const newModule = ref({
    name: '',
    description: '',
    provider: '',
    gitRepositoryUrl: '',
    directory: '',
    mainProvider: '',
    providerVersion: '',
    terraformVersion: ''
  });

  const createModule = () => {
    const form = unref(createModuleRef)
    if (!form) {
      return
    }

    form.validate((valid) => {
      if (valid) {
        const createdAt = new Date();
        request.post("/api/v1/modules", {
          name: newModule.value.name,
          description: newModule.value.description,
          gitRepositoryUrl: newModule.value.gitRepositoryUrl,
          directory: newModule.value.directory,
          mainProvider: newModule.value.mainProvider,
          providerVersion: newModule.value.providerVersion,
          terraformImage: {
            repository: "hashicorp/terraform",
            tag: newModule.value.terraformVersion,
          },
          CreatedAt: createdAt.toISOString(),
        }).then((response) => {
          ElMessage.success("Create success");
          modules.value.push(response.data.data);
          showCreate.value = false;
        })
      } else {
        ElMessage.error("Input invalid");
      }
    })
  }

/*
  delete module
*/
  const showDelete = ref(-1);

  const deleteModule = (row) => {
    request.delete("/api/v1/modules/" + row.id).then(() => {
      ElMessage.success("Delete success");
      const index = modules.value.findIndex(v => v.id === row.id);
      modules.value.splice(index, 1);
      showDelete.value = -1;
    })
  };
/*
 edit module
*/
  const updateFormRef = ref();
  const showUpdate = ref(false);
  const updatedModule = ref({
    name: '',
    description: '',
    provider: '',
    gitRepositoryUrl: '',
    directory: '',
    mainProvider: '',
    providerVersion: '',
    terraformVersion: '',
  });

  const editModule = (row) => {
    updatedModule.value = Object.assign({}, row);
    showUpdate.value = true;
  }

  const updateModule = () => {
    const form = unref(updateFormRef);
    if (!form) {
      return
    }

    form.validate((valid, err) => {
      if (valid) {
        updatedModule.value.updatedAt = new Date()
        request.put("/api/v1/modules/" + updatedModule.value.id, updatedModule.value).then(() => {
          ElMessage.success("Update success");
          const index = modules.value.findIndex(v => v.id === updatedModule.value.id);
          modules.value[index] = updatedModule.value;
          showUpdate.value = false;
        })
      } else {
        ElMessage.error("Input invalid", err);
      }
    });
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
  }

  .operation_icon {
    margin-left: 0.5rem;
  }
</style>
