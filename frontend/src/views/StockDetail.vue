<script setup lang="ts">
import { onMounted } from "vue";
import { useRoute, RouterLink } from "vue-router";
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

    <div
      v-if="stockStore.loading"
      class="flex justify-center items-center py-10"
    >
      <div
        class="animate-spin rounded-full h-8 w-8 border-b-2 border-gray-900"
      ></div>
    </div>

    <div v-else-if="stockStore.error" class="rounded-lg bg-red-50 p-4">
      <div class="flex">
        <div class="ml-3">
          <h3 class="text-sm font-medium text-red-800">Error</h3>
          <div class="mt-2 text-sm text-red-700">
            <p>{{ stockStore.error }}</p>
          </div>
        </div>
      </div>
    </div>

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
            <dt class="text-sm font-medium text-gray-500">
              Precio objetivo anterior
            </dt>
            <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
              ${{ stockStore.selectedStock.target_from.toFixed(2) }}
            </dd>
          </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm font-medium text-gray-500">
              Nuevo precio objetivo
            </dt>
            <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
              ${{ stockStore.selectedStock.target_to.toFixed(2) }}
            </dd>
          </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm font-medium text-gray-500">Cambio porcentual</dt>
            <dd
              class="mt-1 text-sm sm:col-span-2 sm:mt-0"
              :class="
                stockStore.selectedStock.target_to >=
                stockStore.selectedStock.target_from
                  ? 'text-green-600'
                  : 'text-red-600'
              "
            >
              {{
                (
                  ((stockStore.selectedStock.target_to -
                    stockStore.selectedStock.target_from) /
                    stockStore.selectedStock.target_from) *
                  100
                ).toFixed(2)
              }}%
            </dd>
          </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm font-medium text-gray-500">Casa de bolsa</dt>
            <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
              {{ stockStore.selectedStock.brokerage }}
            </dd>
          </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm font-medium text-gray-500">Acción</dt>
            <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
              {{ stockStore.selectedStock.action }}
            </dd>
          </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm font-medium text-gray-500">Calificación</dt>
            <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
              {{ stockStore.selectedStock.rating_from }} →
              {{ stockStore.selectedStock.rating_to }}
            </dd>
          </div>
          <div class="py-4 sm:grid sm:grid-cols-3 sm:gap-4 sm:px-6 sm:py-5">
            <dt class="text-sm font-medium text-gray-500">Fecha</dt>
            <dd class="mt-1 text-sm text-gray-900 sm:col-span-2 sm:mt-0">
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
