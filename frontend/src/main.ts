import { createApp } from "vue"
import App from "./App.vue"
import "./style.css"
import { Tooltip } from "primevue"


import Aura from "@primeuix/themes/aura"
import PrimeVue from "primevue/config"
import "primeicons/primeicons.css"
import { Environment, WindowIsMaximised } from "../wailsjs/runtime/runtime"

const app = createApp(App)

app.use(PrimeVue, {
  theme: {
    preset: Aura,
    options: {
      darkModeSelector: "system",
      cssLayer: false,
    },
  },
})

app.directive('tooltip', Tooltip)

app.mount("#app")

// Add OS platform class to body
Environment().then((env) => {
  document.body.classList.add(`platform-${env.platform}`)
})

// Detect window resize (which happens on maximize/restore)
window.addEventListener('resize', () => {
  console.log('Window resized/maximized/restored')
  WindowIsMaximised().then((isMax) => {
    if (isMax) {
      document.body.classList.add('window-max')
    } else {
      document.body.classList.remove('window-max')
    }
  })
})

// Add window focus / blur classes to body
window.addEventListener('focus', () => {
  document.body.classList.add('window-focused')
  document.body.classList.remove('window-blurred')
})

window.addEventListener('blur', () => {
  document.body.classList.remove('window-focused')
  document.body.classList.add('window-blurred')
})
