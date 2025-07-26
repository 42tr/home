// src/router/index.js
import { createRouter, createWebHistory } from "vue-router";
import Resume from "../views/Resume.vue";
import Home from "../views/Home.vue";

const routes = [
  {
    path: "/resume",
    name: "Resume",
    component: Resume,
  },
  {
    path: "/",
    name: "Home",
    component: Home,
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL), // Vite 特有的
  routes,
});

export default router;
