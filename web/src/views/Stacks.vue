<template>
  <div class="page_content">
    <div class="modules">
      <PageHeader title="Stacks">
        <template #icon>
          <AllApplication style="margin-left: 1rem;" theme="outline" size="36" fill="#333" />
        </template>
      </PageHeader>
    </div>
  </div>
</template>

<script setup>
/*
  imports
*/
  import PageHeader from '@/components/PageHeader.vue'
  import { AllApplication } from '@icon-park/vue-next'
</script>

<style lang="scss">
  .page_content {
    width: 100%;
    display: flex;
    justify-content: center;
  }

  .modules {
    display: flex;
    flex-direction: column;
    width: 100%;
    height: 100%;
    padding: 0.8rem 4rem 0.8rem 1rem;
  }
</style>
