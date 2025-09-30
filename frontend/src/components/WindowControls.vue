<script lang="ts" setup>
import { ref, onMounted } from "vue"
import { WindowMinimise, WindowToggleMaximise, WindowIsMaximised, Quit, Environment } from "../../wailsjs/runtime/runtime"

const isMaximized = ref(false)
const platform = ref("")

onMounted(async () => {
  isMaximized.value = await WindowIsMaximised()
  const env = await Environment()
  platform.value = env.platform
})

const minimize = () => {
  WindowMinimise()
}

const toggleMaximize = () => {
  WindowToggleMaximise()
  isMaximized.value = !isMaximized.value
}

const close = () => {
  Quit()
}
</script>

<template>
  <div class="window-controls">
    <button class="control-button minimize" @click="minimize" aria-label="Minimize">
      <svg class="icon" viewBox="0 0 12 12">
        <line x1="3" y1="9.5" x2="9" y2="9.5" stroke="currentColor" stroke-width="1" stroke-linecap="round" />
      </svg>
    </button>
    <button class="control-button maximize" @click="toggleMaximize" :aria-label="isMaximized ? 'Restore' : 'Maximize'">
      <svg v-if="!isMaximized" class="icon" viewBox="0 0 12 12">
        <rect x="2.5" y="2.5" width="7" height="7" fill="none" stroke="currentColor" stroke-width="1" />
      </svg>
      <svg v-else class="icon" viewBox="0 0 12 12">
        <rect x="3.5" y="4.5" width="5" height="5" fill="none" stroke="currentColor" stroke-width="1" />
        <path d="M 4.5,4.5 V 2.5 H 9.5 V 7.5 H 7.5" fill="none" stroke="currentColor" stroke-width="1" />
      </svg>
    </button>
    <button class="control-button close" @click="close" aria-label="Close">
      <svg class="icon" viewBox="0 0 12 12">
        <path d="M 3,3 L 9,9 M 9,3 L 3,9" stroke="currentColor" stroke-width="1" stroke-linecap="round" />
      </svg>
    </button>
  </div>
</template>

<style scoped>
.window-controls {
  display: flex;
  align-items: center;
}

.icon {
  width: 12px;
  height: 12px;
}

.control-button {
  /* width: 2rem; */
  height: 2rem;
  aspect-ratio: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  border: none;
  background: transparent;
  color: var(--p-text-color);
  cursor: pointer;
  transition: background-color 0.2s, color 0.2s;
  border-radius: 0;
}

.control-button.control-button:hover {
  background-color: color-mix(in srgb, var(--p-text-color), transparent 89%);
}

/* .control-button.close:hover {
  background-color: #e81123;
  color: white;
} */

/* Windows style */
.platform-windows .window-controls {
  gap: 0;
}

.platform-windows .control-button {
  height: 2rem;
  border-radius: 0;
}

.platform-windows .control-button.close:hover {
  background-color: #e81123;
  color: white;
}

/* macOS style */
.platform-darwin .window-controls {
  gap: 0.5rem;
  /* order: -1; */
}

.platform-darwin .control-button {
  height: 0.875rem;
  border-radius: 50%;
  padding: 0;
  border: 0.5px solid color-mix(in srgb, currentColor, transparent 70%);
}

.platform-darwin .control-button .icon {
  display: none;
}

.platform-darwin .control-button.close {
  background-color: #ff5f57;
  border-color: #e0443e;
  color: transparent;
  order: 1;
}

.platform-darwin .control-button.minimize {
  background-color: #ffbd2e;
  border-color: #dea123;
  color: transparent;
  order: 2;
}

.platform-darwin .control-button.maximize {
  background-color: #28c840;
  border-color: #1aab29;
  color: transparent;
  order: 3;
}

.platform-darwin .control-button:hover .icon {
  display: block;
}

.platform-darwin .control-button.close:hover {
  background-color: #ff5f57;
}

.platform-darwin .control-button.minimize:hover {
  background-color: #ffbd2e;
}

.platform-darwin .control-button.maximize:hover {
  background-color: #28c840;
}

/* Linux style */
.platform-linux .window-controls {
  gap: 0.75rem;
}

.platform-linux .control-button {
  height: 1.4375rem;
  background-color: color-mix(in srgb, var(--p-text-color), transparent 93%);
  border-radius: 50%;
}
</style>
