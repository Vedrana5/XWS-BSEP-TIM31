import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";

import "jquery/src/jquery.js";

createApp(App).use(router).mount("#app");
