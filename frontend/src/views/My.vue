<template>
  <div class="my">
    <div>hello</div>
    <div>hello</div>
    <div>{{aaatest}}</div>

  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, toRefs, ref, watch } from "vue";
import { wraprequest } from "@/utils/request";
import { message } from "ant-design-vue";
import { useStore } from "vuex";
import { useRoute, useRouter } from "vue-router";
declare var Date: {
  parseRFC3339(item: any): any;
  toRFC3339UTCString(item: any): any;
  toRFC3339LocaleString(item: any): any;
};
export default defineComponent({
  components: {},

  data() {
    return {
      aaatest: 1
    };
  },
  props: {
    foo: {
      type: String,
      required: false,
    },
  },
  computed: {},

  methods: {
  },
  mounted() {
    console.log("mounted");
    wraprequest({
      umimethod: "GET",
      umipath: "/api/version",
    })
      .then((res: any) => {

      })
      .catch((err) => {
      })
      .finally(() => {
      });
  },
  setup() {

  },
});
</script>

<style lang="scss">
.my {}
</style>
