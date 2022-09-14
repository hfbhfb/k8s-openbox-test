import { createRouter, createWebHashHistory, RouteRecordRaw } from "vue-router";
import Home from "../views/Home.vue";
import Detail from "../views/Detail.vue";
import My from "../views/My.vue";
const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/detail",
    name: "detail",
    component: Detail,
  },
  {
    path: "/my",
    name: "my",
    component: My,
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ "../views/About.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});
router.beforeEach((to, from, next) => {
  // ...
  // console.log(to);
  // console.log(to.path);
  // console.log(from);
  // console.log(next);
  if (to.path == "/my") {
    next();
    return;
  } else {
    let token = localStorage.getItem("access_token");
    if (!token) {
      next({
        path: "/my",
        query: {
          redirect: to.fullPath,
        },
      });
      setTimeout(function () {
        window.location.reload();
      }, 200);
      return;
    }
  }
  // store.state.allmenu
  next();
});
export default router;
