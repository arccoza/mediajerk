import { createApp } from "vue"
import App from "./App.vue"
import "./style.css"
import { Tooltip } from "primevue"


import Aura from "@primeuix/themes/aura"
import PrimeVue from "primevue/config"
import "primeicons/primeicons.css"
import { Environment } from "../wailsjs/runtime/runtime"

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

// Add window focus / blur classes to body
window.addEventListener('focus', () => {
  document.body.classList.add('window-focused')
  document.body.classList.remove('window-blurred')
})

window.addEventListener('blur', () => {
  document.body.classList.remove('window-focused')
  document.body.classList.add('window-blurred')
})
