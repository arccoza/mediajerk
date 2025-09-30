import { createApp } from "vue"
import App from "./App.vue"
import "./style.css"
import { Tooltip } from "primevue"


import Aura from "@primeuix/themes/aura"
import PrimeVue from "primevue/config"
import "primeicons/primeicons.css"

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

window.addEventListener('focus', () => {
  document.body.classList.add('window-focused')
  document.body.classList.remove('window-blurred')
})

window.addEventListener('blur', () => {
  document.body.classList.remove('window-focused')
  document.body.classList.add('window-blurred')
})
