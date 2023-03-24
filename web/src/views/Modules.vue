<template>
  <div class="page_content">
    <el-dialog v-model="showCreate" top="5vh" title="Create Module" width="60%">
      <el-form ref="createModuleRef" :model="newModule" label-position="top" label-width="auto">
        <div class="form_content">
          <el-form-item style="width: 50%;" label="Name" prop="name" required>
            <el-input v-model="newModule.name" />
            <span>The name of your module</span>
          </el-form-item>
          <el-form-item style="width: 50%;" label="Description" prop="description">
            <el-input v-model="newModule.description" type="textarea" />
            <span>The description of your module</span>
          </el-form-item>
        </div>
        <el-form-item label="Git Repository URL" prop="repository" required>
          <el-input v-model="newModule.repository" />
          <span class="text-gray-400">The URL of the module's git repository</span>
        </el-form-item>
      </el-form>
      <template #footer>
          <span class="dialog-footer">
            <el-button type="success" @click="createModule">Confirm</el-button>
            <el-button type="danger" @click="showCreate = false">Cancel</el-button>
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
          <el-table-column prop="name" label="Name" sortable />
          <el-table-column prop="source" label="Source" />
          <el-table-column prop="status" label="Status"/>
          <el-table-column prop="published" label="Published" sortable min-width="120px" />
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
  import { CloudStorage, Search } from '@icon-park/vue-next'
  import { ref, unref, computed } from 'vue';
  import request from '@/axios'
  import { ElMessage } from "element-plus";

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

/*
  create module
*/
  const showCreate = ref(false);
  const defaultTime = { timeout: "10000" }
  const createModuleRef = ref();
  const newModule = ref({
    name: '',
    repository: '',
  });

  const createModule = () => {
    const form = unref(createModuleRef)
    if (!form) {
      return
    }

    form.validate((valid) => {
      if (valid) {
        request.post("/api/v1/modules", {
          name: newModule.value.name,
          repository: newModule.value.repository,
          description: newModule.description,
        }, defaultTime).then((response) => {
          ElMessage.success("Create success");
          modules.value.push(response.data.data);
          showCreate.value = false;
        })
      } else {
        ElMessage.error("Input invalid");
      }
    })
  }
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
</style>
