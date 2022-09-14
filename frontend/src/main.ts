import { createApp } from "vue";
import App from "./App.vue";

// 仅用来测试布局
// import App from "./Appflex.vue";

import router from "./router";
import store from "./store";

require("@/utils/rfc3339date.js");

import Ant from "./utils/ant";
createApp(App).use(Ant).use(store).use(router).mount("#app");
