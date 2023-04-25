import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import PlaygroundView from "../views/PlaygroundView.vue";

const routes = [
  {
    path: "/",
    name: "index",
    component: HomeView,
  },
  {
    path: "/playground",
    name: "playground",
    component: PlaygroundView,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});
router.beforeEach(() => {
  document.title = "CodeGPT";
});

export default router;
