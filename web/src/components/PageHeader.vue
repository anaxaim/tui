<template>
  <div class="page_header">
    <div class="page_header_content">
      <slot name="icon" />
      <span class="page_header_text">{{ title }}</span>
    </div>
  </div>
</template>

<script setup>
defineProps({
  title: {
    type: String,
    required: true,
  },
});
</script>

<style lang="scss" scoped>
  .page_header {
    display: flex;
    overflow: hidden;
    border-radius: 0.375rem;
    border-width: 4px;
    box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
    margin-bottom: 2rem;
  }

  .page_header_content {
    display: flex;
    width: 100%;
    height: 5rem;
    align-items: center;
  }

  .page_header_text {
    font-size: 1.5rem;
    line-height: 2rem;
    margin: 0.75rem;
    font-weight: 500;
  }
</style>
