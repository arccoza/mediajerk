<script lang="ts" setup>
import AppHeader from "./components/AppHeader.vue"
import FileList from "./components/FileList.vue"
import FilePreview from "./components/FilePreview.vue"
import Splitter from "primevue/splitter"
import SplitterPanel from "primevue/splitterpanel"
</script>

<template>
  <div class="app-container">
    <AppHeader />

    <div class="app-main">
      <Splitter style="height: 100%; border: none">
        <SplitterPanel :size="50" :min-size="30">
          <FileList />
        </SplitterPanel>
        <SplitterPanel :size="50" :min-size="30">
          <FilePreview />
        </SplitterPanel>
      </Splitter>
    </div>
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html, body {
  height: 100%;
  font-family: var(--p-font-family);
  background-color: light-dark(var(--p-surface-0), var(--p-surface-900));
  color: var(--p-text-color);
}

#app {
  height: 100vh;
}

.app-container {
  height: 100%;
  display: flex;
  flex-direction: column;
  gap: 0em;
}

.app-main {
  flex: 1;
  overflow: hidden;
  /* margin: var(--p-toolbar-padding); */
}
</style>
