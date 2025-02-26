<script setup lang="ts">
import { onMounted } from "vue";
import { RouterLink, useRoute } from "vue-router";
import ErrorMessage from "../components/common/ErrorMessage.vue";
import LoadingSpinner from "../components/common/LoadingSpinner.vue";
import PriceChange from "../components/common/PriceChange.vue";
import { useStockStore } from "../stores/stocks";

const route = useRoute();
const stockStore = useStockStore();

onMounted(async () => {
  const symbol = route.params.symbol as string;
  await stockStore.fetchStockDetail(symbol);
});
</script>

<template>
  <div>
    <div class="mb-8">
      <RouterLink to="/" class="text-indigo-600 hover:text-indigo-900">
        ← Volver al listado
      </RouterLink>
    </div>

    <LoadingSpinner v-if="stockStore.loading" />

    <ErrorMessage v-else-if="stockStore.error" :message="stockStore.error" />

    <div
      v-else-if="stockStore.selectedStock"
      class="overflow-hidden bg-white shadow sm:rounded-lg"
    >
      <div class="px-4 py-6 sm:px-6">
        <h3 class="text-2xl font-semibold leading-6 text-gray-900">
          {{ stockStore.selectedStock.company }}
        </h3>
        <p class="mt-1 text-sm text-gray-500">
          {{ stockStore.selectedStock.ticker }}
        </p>
      </div>
      <div class="border-t border-gray-200 px-4 py-5 sm:p-0">
        <dl class="sm:divide-y sm:divide-gray-200">
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm text-left font-medium text-gray-500">
              Precio objetivo anterior
            </dt>
            <dd
              class="mt-1 text-sm text-left text-gray-900 sm:col-span-2 sm:mt-0"
            >
              ${{ stockStore.selectedStock.target_from.toFixed(2) }}
            </dd>
          </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm text-left font-medium text-gray-500">
              Nuevo precio objetivo
            </dt>
            <dd
              class="mt-1 text-sm text-left text-gray-900 sm:col-span-2 sm:mt-0"
            >
              ${{ stockStore.selectedStock.target_to.toFixed(2) }}
            </dd>
          </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm text-left font-medium text-gray-500">
              Cambio porcentual
            </dt>
            <dd class="mt-1 text-sm text-left sm:col-span-2 sm:mt-0">
              <PriceChange
                :from-value="stockStore.selectedStock.target_from"
                :to-value="stockStore.selectedStock.target_to"
              />
            </dd>
          </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm text-left font-medium text-gray-500">
              Casa de bolsa
            </dt>
            <dd
              class="mt-1 text-sm text-left text-gray-900 sm:col-span-2 sm:mt-0"
            >
              {{ stockStore.selectedStock.brokerage }}
            </dd>
          </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm text-left font-medium text-gray-500">Acción</dt>
            <dd
              class="mt-1 text-sm text-left text-gray-900 sm:col-span-2 sm:mt-0"
            >
              {{ stockStore.selectedStock.action }}
            </dd>
          </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm text-left font-medium text-gray-500">
              Calificación
            </dt>
            <dd
              class="mt-1 text-sm text-left text-gray-900 sm:col-span-2 sm:mt-0"
            >
              {{ stockStore.selectedStock.rating_from }} →
              {{ stockStore.selectedStock.rating_to }}
            </dd>
          </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm text-left font-medium text-gray-500">Fecha</dt>
            <dd
              class="mt-1 text-sm text-left text-gray-900 sm:col-span-2 sm:mt-0"
            >
              {{ new Date(stockStore.selectedStock.time).toLocaleString() }}
            </dd>
          </div>
        </dl>
      </div>
    </div>

    <div v-else class="text-center py-10">
      <p class="text-gray-500">No se encontró información del stock</p>
    </div>
  </div>
</template>
