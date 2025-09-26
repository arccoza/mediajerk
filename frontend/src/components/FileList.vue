<script lang="ts" setup>
import Column from "primevue/column"
import DataTable from "primevue/datatable"
import Tag from "primevue/tag"
import { type FileInfo, useFiles } from "../composables/useFiles"


type Emits = (e: "row-reorder", files: FileInfo[]) => void

const emit = defineEmits<Emits>()
const { files, selectedFiles, selectFiles } = useFiles()

const onRowReorder = (event: any) => {
  emit("row-reorder", event.value)
}
</script>

<template>
  <div class="file-list">
    <div class="file-list-header">
      <h3>Original Files</h3>
    </div>

    <DataTable
      :value="files"
      :selection="selectedFiles"
      :scrollable="true"
      scroll-height="100%"
      reorderableRows
      selectionMode="multiple"
      dataKey="path"
      @row-reorder="onRowReorder"
      @update:selection="selectFiles"
      class="file-table"
      :pt="{
        table: { style: 'min-width: 100%' },
        bodyRow: { style: 'height: 4rem' }
      }"
    >
      <Column selectionMode="multiple" headerStyle="width: 3rem"></Column>
      <Column :rowReorder="true" header="" style="width: 3rem" />

      <Column field="name" header="Filename" :sortable="true" style="min-width: 200px">
        <template #body="{ data }">
          <div class="filename-cell">
            <span class="filename">{{ data.name }}</span>
            <small class="file-path">{{ data.dir }}</small>
          </div>
        </template>
      </Column>

      <Column field="ext" header="Extension" :sortable="true" style="width: 120px">
        <template #body="{ data }">
          <div class="extension-cell">
            <Tag :value="data.ext" severity="info" class="extension-tag" />
          </div>
        </template>
      </Column>
    </DataTable>
  </div>
</template>

<style scoped>
.file-list {
  height: 100%;
  display: flex;
  flex-direction: column;
  gap: 1em;
  font-size: 0.9em;
}

.file-list-header {
  padding: 0.75rem;
  border-bottom: 1px solid var(--p-surface-border);
  --bg: light-dark(var(--p-surface-100), var(--p-surface-800));
  background-color: color-mix(in srgb, var(--bg), transparent 50%);
}

.file-list-header h3 {
  margin: 0;
  font-size: 0.75rem;
  font-weight: 500;
  text-transform: uppercase;
}

.file-table {
  flex: 1;
  border: none;
}

.filename-cell {
  display: flex;
  flex-direction: column;
  justify-content: center;
  height: 100%;
  padding: 0.5rem 0;
}

.filename {
  font-weight: 500;
  margin-bottom: 0.25rem;
  line-height: 1.2;
  text-wrap: nowrap;
}

.file-path {
  color: var(--p-text-muted-color);
  font-size: 0.875rem;
  line-height: 1.2;
}

.extension-cell {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 100%;
}

.extension-tag {
  font-family: monospace;
  font-size: 0.75rem;
  text-transform: uppercase;
}
</style>
