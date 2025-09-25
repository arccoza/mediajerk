<script lang="ts" setup>
import Splitter from "primevue/splitter"
import SplitterPanel from "primevue/splitterpanel"
import { ref } from "vue"
import AppHeader from "./components/AppHeader.vue"
import FileList from "./components/FileList.vue"
import FilePreview from "./components/FilePreview.vue"
import { useFiles } from "./composables/useFiles"
import type { MediaMetadata } from "./utils/templateProcessor"

interface PreviewItem {
  id: number
  originalFilename: string
  newFilename: string
  metadata: MediaMetadata
  status: "ready" | "warning" | "error"
  message?: string
}

// Use files composable for state management
const { reorderFiles } = useFiles()

const previewFiles = ref<PreviewItem[]>([
  {
    id: 1,
    originalFilename: "Game.of.Thrones.S01E01.Winter.is.Coming.mkv",
    newFilename: "Game of Thrones (2011) - S01E01.mkv",
    metadata: {
      title: "Game of Thrones",
      year: 2011,
      season: 1,
      episode: 1,
      episodeTitle: "Winter is Coming",
      type: "tv",
    },
    status: "ready",
  },
  {
    id: 2,
    originalFilename: "Breaking.Bad.S01E01.Pilot.mp4",
    newFilename: "Breaking Bad (2008) - S01E01.mp4",
    metadata: {
      title: "Breaking Bad",
      year: 2008,
      season: 1,
      episode: 1,
      episodeTitle: "Pilot",
      type: "tv",
    },
    status: "ready",
  },
  {
    id: 3,
    originalFilename: "The.Lord.of.the.Rings.The.Two.Towers.2002.mkv",
    newFilename: "The Lord of the Rings - The Two Towers (2002).mkv",
    metadata: {
      title: "The Two Towers",
      year: 2002,
      seriesName: "The Lord of the Rings",
      type: "movie",
    },
    status: "ready",
  },
  {
    id: 4,
    originalFilename: "The.Matrix.1999.mp4",
    newFilename: "The Matrix (1999).mp4",
    metadata: {
      title: "The Matrix",
      year: 1999,
      type: "movie",
    },
    status: "ready",
  },
])

// Synchronized reordering functions
const onFilesReorder = (reorderedFiles: any[]) => {
  reorderFiles(reorderedFiles)
  // TODO: Implement preview files reordering when needed
}

const onPreviewReorder = (reorderedPreviews: PreviewItem[]) => {
  previewFiles.value = reorderedPreviews
  // TODO: Implement files reordering to match preview order when needed
}
</script>

<template>
  <div class="app-container">
    <AppHeader />

    <div class="app-main">
      <Splitter style="height: 100%; border: none">
        <SplitterPanel :size="50" :min-size="30">
          <FileList
            @row-reorder="onFilesReorder"
          />
        </SplitterPanel>
        <SplitterPanel :size="50" :min-size="30">
          <FilePreview
            :previewFiles="previewFiles"
            @row-reorder="onPreviewReorder"
          />
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
