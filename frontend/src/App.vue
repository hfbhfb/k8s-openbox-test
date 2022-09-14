<template>
  <div class="appout">
    <router-view />
    <!-- <div>{{ showinfo }}</div> -->
    <div class="menu">
      <div @click="changePage('home')" class="column" :class="fnMenuClass('home')">
        <div class="auto">首页</div>
      </div>
      <div @click="changePage('detail')" class="column" :class="fnMenuClass('detail')">
        <div class="auto">详情页</div>
      </div>
      <div @click="changePage('my')" class="column" :class="fnMenuClass('my')">
        <div class="auto">我的</div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, toRefs, ref, watch } from "vue";
// import { allmenu } from "@/store/";
import store from "@/store/";
import { mapState } from "vuex";
import { useRoute, useRouter } from "vue-router";
export default defineComponent({
  components: {},

  data() {
    return {};
  },
  props: {},

  methods: {
    fnMenuClass(id: string) {
      if (id == this.currentPath) {
        return "focus";
      }
    },
    changePage(page: string) {
      this.parchangePage(page);
    },
    fnHeight() {
      var pageWidth = window.innerWidth;
      var pageHeight = window.innerHeight;
      return {
        height: pageHeight,
      };
    },
  },
  // computed: mapState([
  //   "allmenu", // 映射 this.count 为 store.state.count
  // ]),
  setup() {
    let openKeys: Array<string> = [];
    const router = useRouter();
    const route = useRoute();
    let selectedKeys: Array<String> = [];
    const state = reactive({
      currentPath: "home",
      showinfo: "",
    });

    const handleClick = (e: Event) => {
      console.log("click", e);
    };
    const titleClick = (e: Event) => {
      console.log("titleClick", e);
    };

    let fnSetOpen = () => {
      if (route && route.path) {
        if (route.path == "/") {
          state.currentPath = "home";
        }

        if (route.path == "/my") {
          state.currentPath = "my";
        }

        if (route.path == "/detail") {
          state.currentPath = "detail";
        }
        state.showinfo = route.path;
        // if(route.path )
      }
      state.showinfo = "11w";
      console.log(route);
    };
    fnSetOpen();

    watch(
      () => route.path,
      (val, oval) => {
        fnSetOpen();
      }
    );

    let parchangePage = (val: string) => {
      if (val == "home") {
        router.push("/");
      } else {
        router.push(val);
      }
    };
    return {
      ...toRefs(state),
      handleClick,
      titleClick,
      parchangePage,
    };
  },
});
</script>

<style lang="scss">
body {
  font-size: 0.9rem;
  margin: 0px;
}
.ml-16 {
  margin-left: 0.3rem;
}
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
.appout {
  .menu {
    position: fixed;
    bottom: 0px;
    width: 100%;
    // height: 3rem;
    background: rgb(239, 237, 237);
    display: flex;
    // align-content: center;
    .column {
      float: left;
      font-size: 1.1rem;
      width: 33.33%;
      height: 3rem;
      display: flex;
      // border-right: 1px solid rgb(177, 177, 177);
      .auto {
        margin: auto;
        // height: 100%;
        width: 100%;
      }
    }
    .focus {
      background: chocolate;
    }
  }
}
</style>
