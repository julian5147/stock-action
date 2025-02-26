<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
import { useStockStore } from "./stores/stocks";

const stockStore = useStockStore();
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <header class="bg-white shadow">
      <nav class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
        <div class="flex h-16 justify-between">
          <div class="flex">
            <div class="flex flex-shrink-0 items-center">
              <RouterLink to="/" class="text-xl font-bold text-gray-900"
                >Stock Action</RouterLink
              >
            </div>
            <div class="hidden sm:ml-6 sm:flex sm:space-x-8">
              <RouterLink
                to="/"
                class="inline-flex items-center px-1 pt-1 text-sm font-medium text-gray-900"
                :class="{
                  'border-b-2 border-indigo-500': $route.name === 'home',
                }"
              >
                Stocks disponibles
              </RouterLink>
              <RouterLink
                to="/recommended"
                class="inline-flex items-center px-1 pt-1 text-sm font-medium text-gray-900"
                :class="{
                  'border-b-2 border-indigo-500': $route.name === 'recommended',
                }"
              >
                Recomendaciones de inversión
              </RouterLink>
            </div>
          </div>
          <div class="flex items-center">
            <button
              type="button"
              @click="stockStore.syncStocks"
              :disabled="stockStore.loading"
              class="rounded-full p-2 text-gray-600 hover:text-gray-900 focus:outline-none focus:ring-2 focus:ring-indigo-500 disabled:opacity-50 disabled:cursor-not-allowed"
              :title="
                stockStore.loading ? 'Sincronizando...' : 'Sincronizar Stocks'
              "
            >
              <svg
                :class="{ 'animate-spin': stockStore.loading }"
                class="h-6 w-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
                />
              </svg>
            </button>
          </div>
        </div>
      </nav>
    </header>

    <main>
      <div class="mx-auto max-w-7xl py-6 sm:px-6 lg:px-8">
        <RouterView />
      </div>
    </main>
  </div>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
