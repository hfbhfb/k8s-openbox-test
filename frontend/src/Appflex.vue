<template>
  <div class="appflex flex">
    <!-- 侧边栏 -->
    <div style="width: 200px; z-index: 2" class="ceBian baiSe yinYing">
      <!-- 头像栏 -->
      <div class="flex twocenter pd10" style="border-bottom: 1px solid #999">
        <img alt="logo" src="https://kangwenchang.com/logo.png" class="imgcss" />
        <div class="mgl8">何发宝</div>
      </div>
      <div class="flex1 flex column">
        <div class="daohan">导航</div>
        <div class="daohan">导航</div>
        <div class="daohan">导航</div>
        <div class="daohan">导航</div>
        <div class="daohan">导航</div>
        <div class="daohan">导航</div>
      </div>
    </div>
    <!-- 主区域 -->
    <div class="flex flex1 column">
      <!-- 头部栏 -->
      <div style="height: 60px; z-index: 1" class="baiSe yinYing"></div>
      <!-- 内容区 -->
      <div style="overflow: auto" class="neiRong flex1 flex">
        <!-- 左区 -->
        <div style="flex: 3" class="flex column mg8">
          <!-- 数据区 -->
          <div class="flex shuJuQu">
            <div class="shuJu flex1 baiSe yinYing mgr8">数据1</div>
            <div class="shuJu flex1 baiSe yinYing mgr8">数据2</div>
            <div class="shuJu flex1 baiSe yinYing mgr8">数据3</div>
            <div class="shuJu flex1 baiSe yinYing mgr8">数据4</div>
          </div>
          <!-- 列表区 -->
          <div class="flex column">
            <div style="height: 160px" class="baiSe yinYing mgt8">列表1</div>
            <div style="height: 160px" class="baiSe yinYing mgt8">列表2</div>
            <div style="height: 160px" class="baiSe yinYing mgt8">列表3</div>
            <div style="height: 160px" class="baiSe yinYing mgt8">列表4</div>
            <div style="height: 160px" class="baiSe yinYing mgt8">列表5</div>
            <div style="height: 160px" class="baiSe yinYing mgt8">列表6</div>
            <div style="height: 160px" class="baiSe yinYing mgt8">列表7</div>
            <div style="height: 160px" class="baiSe yinYing mgt8">列表8</div>
          </div>
        </div>
        <!-- 右区 -->
        <div style="flex: 1" class="flex column">
          <!-- 提示区 -->
          <div style="height: 150px" class="baiSe yinYing"></div>
          <!-- 消息区 -->
          <div style="height: 300px" class="baiSe yinYing"></div>
        </div>
      </div>
    </div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, toRefs, ref, watch } from "vue";

export default defineComponent({
  components: {},

  data() {
    return {};
  },
  props: {},

  methods: {},
  // computed: mapState([
  //   "allmenu", // 映射 this.count 为 store.state.count
  // ]),
  setup() {
    let openKeys: Array<string> = [];
    let parchangePage = (val: string) => {};
    return {};
  },
});
</script>

<style lang="scss">
body {
  margin: 0;
  height: 100vh;
  background-color: #eee;
}
// * {
//   border: 1px solid black;
// }
.twocenter {
  align-items: center;
  justify-content: center;
}
.imgcss {
  width: 40px;
  height: 40px;
}
.flex {
  display: flex;
}
.flex1 {
  flex: 1;
}
.baiSe {
  background-color: white;
}
.yinYing {
  box-shadow: rgb(0 0 0 / 20%) 0px 2px 1px -1px, rgb(0 0 0 / 14%) 0px 1px 1px 0px, rgb(0 0 0 / 12%) 0px 1px 3px 0px;
}
.column {
  flex-direction: column;
}
.mg8 {
  margin: 8px;
}
.mgr8 {
  margin-right: 8px;
}
.mgl8 {
  margin-left: 8px;
}
.mgt8 {
  margin-top: 8px;
}
.pd8 {
  padding: 8px;
}

.pd10 {
  padding: 10px;
}
.daohan {
  padding: 10px 20px;
  border-bottom: 1px solid #999;
  font-size: 16px;
}
.appflex {
  margin: 0;
  height: 100vh;
  background-color: #eee;
}
.shuJu {
  height: 100px;
}

@media (max-width: 600px) {
  .ceBian {
    display: none;
  }
  .neiRong {
    flex-direction: column;
  }
  .shuJuQu {
    flex-direction: column;
  }
  .shuJu {
    margin: 0 0 8px;
    flex: auto;
  }
}
</style>
