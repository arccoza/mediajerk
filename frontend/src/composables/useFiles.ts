import { computed, readonly, ref } from "vue"
import type { main } from "../../wailsjs/go/models"

export type FileInfo = main.FileInfo

// Shared reactive state
const files = ref<FileInfo[]>([])
const selectedFiles = ref<FileInfo[]>([])

export function useFiles() {
  // Computed properties for selection
  const hasSelection = computed(() => selectedFiles.value.length > 0)
  const selectedCount = computed(() => selectedFiles.value.length)

  // Add files to the collection
  const addFiles = (newFiles: FileInfo[]) => {
    // Filter out duplicates based on path
    const existingPaths = new Set(files.value.map((f) => f.path))
    const uniqueNewFiles = newFiles.filter((file) => !existingPaths.has(file.path))

    files.value.push(...uniqueNewFiles)
    console.log(files.value)
  }

  // Selection management functions
  const selectFiles = (newSelection: FileInfo[]) => {
    selectedFiles.value = newSelection
  }

  const clearSelection = () => {
    selectedFiles.value = []
  }

  // Remove file by path
  const removeFile = (path: string) => {
    files.value = files.value.filter((file) => file.path !== path)
    // Also remove from selection if it was selected
    selectedFiles.value = selectedFiles.value.filter((file) => file.path !== path)
  }

  // Remove currently selected files
  const removeSelectedFiles = () => {
    const selectedPaths = new Set(selectedFiles.value.map((f) => f.path))
    files.value = files.value.filter((file) => !selectedPaths.has(file.path))
    clearSelection()
  }

  // Remove file by index
  const removeFileByIndex = (index: number) => {
    files.value.splice(index, 1)
  }

  // Clear all files
  const clearFiles = () => {
    files.value = []
  }

  // Reorder files
  const reorderFiles = (reorderedFiles: FileInfo[]) => {
    files.value = [...reorderedFiles]
  }

  // Get file count
  const fileCount = computed(() => files.value.length)

  // Check if files list is empty
  const isEmpty = computed(() => files.value.length === 0)

  return {
    // Read-only access to files
    files: readonly(files),
    selectedFiles,

    // Computed properties
    fileCount,
    isEmpty,
    hasSelection,
    selectedCount,

    // Methods
    addFiles,
    removeFile,
    removeFileByIndex,
    removeSelectedFiles,
    clearFiles,
    reorderFiles,
    selectFiles,
    clearSelection,
  }
}
