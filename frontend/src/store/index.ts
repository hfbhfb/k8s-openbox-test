import { createStore } from "vuex";

export default createStore({
  state: {
    user: {
      name: "",
      account: -1,
    },
  },
  mutations: {
    update_user(state, item) {
      state.user.name = item.name;
      state.user.account = item.account;
    },
  },
  actions: {},
  modules: {},
});
