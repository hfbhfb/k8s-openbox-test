<template>
  <div class="detail">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="1" tab="今日渲染">
        <a-button class="freshbutton" @click="fnFreshTodayList()" type="primary">刷新列表</a-button>
        <div v-for="(item, index) in listToday" :key="index">
          <div class="divide1">
            <div class="row1">
              {{ item.url }}
              <div class="time">{{ TranslateDate(item.create_time) }}</div>
            </div>
            <div :class="getrow2(item.status)">
              {{ fnGetStatus(item.status) }}
            </div>
            <div class="row3">
              <div v-if="checkCopy(item.status)">
                <a-button size="small" @click="fnCopyUrl(item.callback_url)" type="primary">复制发布链接</a-button>
              </div>
            </div>
          </div>
        </div>
        <!-- <list-my :list="listToday"></list-my> -->
      </a-tab-pane>
      <a-tab-pane key="2" tab="历史渲染" force-render>
        <a-button class="freshbutton" @click="fnFreshHistoryList()" type="primary">刷新列表</a-button>
        <div v-for="(item, index) in listHistory" :key="index">
          <div class="divide1">
            <div class="row1">
              {{ item.url }}
              <div class="time">{{ TranslateDate(item.create_time) }}</div>
            </div>
            <div class="row2">
              {{ fnGetStatus(item.status) }}
            </div>
            <div class="row3">
              <div v-if="checkCopy(item.status)">
                <a-button size="small" @click="fnCopyUrl(item.callback_url)" type="primary">复制发布链接</a-button>
              </div>
            </div>
          </div>
        </div>
        <!-- <list-my :list="listHistory"></list-my> -->
      </a-tab-pane>
    </a-tabs>
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
import ListMy from "@/components/ListMy.vue";
declare var Date: {
  parseRFC3339(item: any): any;
  toRFC3339UTCString(item: any): any;
  toRFC3339LocaleString(item: any): any;
};
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
    checkCopy(status: number) {
      if (status == 3) {
        return true;
      }
    },
    fnFreshHistoryList() {
      wraprequest({
        umimethod: "GET",
        umipath: "/detail/history",
      })
        .then((res: any) => {
          let { consumerlist } = res.data;
          if (consumerlist) {
            this.listHistory = consumerlist;
          }
          return null;
        })
        .catch((err) => {
          message.error(err);
        })
        .finally(() => {});
    },
    fnFreshTodayList() {
      wraprequest({
        umimethod: "GET",
        umipath: "/detail/today",
      })
        .then((res: any) => {
          let { consumerlist } = res.data;
          if (consumerlist) {
            this.listToday = consumerlist;
          }
        })
        .catch((err) => {
          message.error(err);
        })
        .finally(() => {});
    },
    TranslateDate(item: string): any {
      let d1 = Date.parseRFC3339(item);
      return d1.toLocaleString("cn", { hour12: false });
      // return d1.toLocaleString("", { hour12: false });
    },
    fnCopyUrl(value: string) {
      var inputTest = document.createElement("input");
      inputTest.value = value;
      document.body.appendChild(inputTest);
      inputTest.select();
      document.execCommand("Copy");
      inputTest.className = "oInput";
      inputTest.style.display = "none";
      message.success("成功复制到链接");
    },
    getUserRepositories(): string {
      return "";
    }, // 1
    getrow2(status: number) {
      if (this.fnGetStatus(status) == "正在渲染") {
        return "rowgreen";
      }
      return "row2";
    },
    fnGetStatus(status: number): string {
      // 状态 1: 未处理  2:正在处理 3:已完成 4:失败
      if (status == 1) {
        return "等待渲染";
      }
      if (status == 2) {
        return "正在渲染";
      }
      if (status == 3) {
        return "已完成";
      }
      if (status == 4) {
        return "失败";
      }
      return "";
    },
  },
  mounted() {
    this.fnFreshTodayList();
    this.fnFreshHistoryList();
  },
  setup() {
    const state = reactive({
      pageinfo: "detail",
      activeKey: "1",
      listToday: [],
      listHistory: [],
    });

    return {
      ...toRefs(state),
    };
  },
});
</script>

<style lang="scss">
.detail {
  .freshbutton {
    position: absolute;
    right: 12px;
    top: 12px;
  }
  .divide1 {
    margin-top: 0.7rem;
    margin-bottom: 0.7rem;
    display: flex;
    flex-direction: row;
    border-bottom: 1px solid gray;
    .row1 {
      word-wrap: break-word;
      margin: auto;
      width: 50%;
      .time {
        color: #088900;
        text-align: left;
      }
    }
    .row2 {
      margin: auto;
      width: 20%;
    }

    .rowgreen {
      color: green;
      margin: auto;
      width: 20%;
    }
    .row3 {
      margin: auto;
      width: 30%;
    }
  }
}
</style>
