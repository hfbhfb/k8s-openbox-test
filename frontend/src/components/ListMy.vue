<template>
  <div class="listmy">
    <div>1111</div>
    <div>1111</div>
    <div>1111</div>
    <div>1111</div>
    <div>1111</div>
  </div>
</template>
<script lang="ts">
import { defineComponent, reactive, toRefs, ref, watch } from "vue";
export default defineComponent({
  components: {},

  data() {
    return {};
  },
  props: {
    item: {
      type: Object,
      required: true,
    },
  },

  methods: {
    getUserRepositories(): string {
      return "";
    }, // 1
  },
  setup() {
    const state = reactive({
      pageinfo: "listmy",
    });

    return {
      ...toRefs(state),
    };
  },
});
</script>

<style lang="scss">
.listmy {
}
</style>
