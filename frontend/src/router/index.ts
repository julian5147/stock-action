import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      name: "home",
      component: () => import("@/views/StockDashboard.vue"),
    },
    {
      path: "/recommended",
      name: "recommended",
      component: () => import("@/views/RecommendedStocks.vue"),
    },
    {
      path: "/stock/:symbol",
      name: "stock-detail",
      component: () => import("@/views/StockDetail.vue"),
    },
  ],
});

export default router;
