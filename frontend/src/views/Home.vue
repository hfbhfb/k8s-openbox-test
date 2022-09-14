<template>
  <div class="home">
    <div class="urlcommit">
      <div class="divout">
        <div class="divide1">
          <div class="row1">
            <div>用户名: {{ user.name }}</div>
          </div>
          <div class="row2">
            <div>渲染币: {{ user.account }}</div>
            <div></div>
          </div>
        </div>

        <div class="wraphalf">
          <div class="half">将抖音的分享链接粘粘到框里并提交(它会自动识别抖音链接)</div>
        </div>
        <div class="areaclass">
          <textarea v-model="valuearea" name="a1" id="id1" cols="30" rows="6"></textarea>
          <a-button @click="fnCleanValue()" :loading="commitLoading">清空</a-button>
        </div>
        <div class="divide2">
          <div class="row1">
            <a-button @click="fnCommit()" :loading="commitLoading">提交</a-button>
          </div>
          <div class="row2">
            <div></div>
          </div>
        </div>
      </div>
    </div>
    <!-- <div>{{ pageinfo }}</div>
    <div>{{ pageinfo }}</div>
    <div>{{ pageinfo }}</div>
    <div>{{ pageinfo }}</div>
    <div>{{ pageinfo }}</div>
    <div>{{ pageinfo }}</div> -->
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, toRefs, ref, watch } from "vue";
import { wraprequest } from "@/utils/request";
import { message } from "ant-design-vue";
import { useStore } from "vuex";

export default defineComponent({
  components: {},

  data() {
    return {};
  },
  props: {
    foo: {
      type: String,
      required: false,
    },
  },

  methods: {
    getUserRepositories(): string {
      return "";
    }, // 1
    fnFlushAccount() {
      wraprequest({
        umimethod: "GET",
        umipath: "/my/userinfo",
      })
        .then((res: any) => {
          let { account } = res.data;
          this.store.commit("update_user", res.data);

          return null;
        })
        .catch((err) => {
          message.error(err);
        })
        .finally(() => {});
    },
    fnCleanValue() {
      this.valuearea = "";
    },
    fnCommit() {
      if (!this.valuearea) {
        message.error("内容为空");
        return;
      }
      // var reg = RegExp(/donyi/);
      // console.log(this.valuearea.match(reg));
      // console.log(this.valuearea);
      if (this.valuearea.indexOf("douyin") == -1) {
        message.error("链接出错,不是抖音链接");
        return;
      }

      // donyi

      if (this.commitLoading) {
        return;
      }
      this.commitLoading = true;
      wraprequest({
        umimethod: "POST",
        umipath: "/home/commit",
        url: this.valuearea,
      })
        .then((res: any) => {
          console.log(res);
          let { msg } = res;
          message.success("提交成功");
          this.fnFlushAccount();
          this.valuearea = "";
          return null;
        })
        .catch((err) => {
          message.error(err);
          console.log(err);
        })
        .finally(() => {
          this.commitLoading = false;
        });
    },
  },
  mounted() {
    if (this.user.account == -1) {
      this.fnFlushAccount();
    }
  },
  setup() {
    const store = useStore();
    // let selectedKeys: Array<String> = [];
    const state = reactive({
      pageinfo: "home",
      // valuearea: "https://douyin.com/lkdslfj",
      valuearea: "",
      commitLoading: false,
      user: store.state.user,
    });

    return {
      ...toRefs(state),
      store,
    };
  },
});
</script>

<style lang="scss">
.home {
  .urlcommit {
    position: absolute;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;

    display: flex;
    flex-wrap: wrap;
    align-items: center;
    justify-content: center;

    .divout {
      justify-content: center;
      display: flex;
      flex-direction: column;
      width: 70%;
    }

    .divide1 {
      position: fixed;
      top: 0;
      display: flex;
      flex-direction: row;
      height: 2rem;
      justify-content: center;
      .row1 {
        margin: auto;
      }
      .row2 {
        margin: auto;
        margin-left: 2rem;
      }
    }
    .areaclass {
      display: flex;
      flex-direction: row;
      justify-content: center;
      align-items: center;

      button {
        margin-left: 0.3rem;
      }
    }
    .wraphalf {
      display: flex;
      justify-content: center;
      .half {
        width: 80%;
      }
    }

    .divide2 {
      margin-top: 0.8rem;
      display: flex;
      flex-direction: row;
      justify-content: center;
      .row1 {
        margin: auto;
        font-size: 1rem;
        width: 35%;
        font-weight: 500;
        button {
          width: 5rem !important;
          // border-radius: 1px;
          height: 2rem;
        }
      }
      .row2 {
        width: 80%;
        height: 2rem;
        display: flex;
        justify-content: flex-start;
        input {
          width: 65%;
        }
      }
    }
  }
}
</style>
