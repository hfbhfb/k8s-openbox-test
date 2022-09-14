import { Button, Tooltip, Tabs } from "ant-design-vue";

const { TabPane } = Tabs;
import { SyncOutlined } from "@ant-design/icons-vue";
import { App } from "vue"; // App是类型

const Ant = {
  install: function (Vue: App) {
    Vue.component("a-button", Button);
    Vue.component("a-tabs", Tabs);
    Vue.component("a-tab-pane", TabPane);
    Vue.component("a-tooltip", Tooltip);
    Vue.component("SyncOutlined", SyncOutlined);
  },
};

export default Ant;
