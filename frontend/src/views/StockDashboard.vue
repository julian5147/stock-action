<script setup lang="ts">
import { onMounted } from "vue";
import { useStockStore } from "../stores/stocks";
import { RouterLink } from "vue-router";

const stockStore = useStockStore();

onMounted(async () => {
  await stockStore.fetchStocks();
});
</script>

<template>
  <div>
    <div class="sm:flex sm:items-center">
      <div class="sm:flex-auto">
        <h1 class="text-2xl font-semibold text-gray-900">Stocks</h1>
        <p class="mt-2 text-sm text-gray-700">
          Lista de stocks disponibles con sus precios actuales y cambios.
        </p>
      </div>
    </div>

    <div class="mt-8 flow-root">
      <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
        <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
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

          <table v-else class="min-w-full divide-y divide-gray-300">
            <thead>
              <tr>
                <th
                  scope="col"
                  class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0"
                >
                  Símbolo
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                >
                  Nombre
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                >
                  Precio
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                >
                  Cambio
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-right text-sm font-semibold text-gray-900"
                >
                  Volumen
                </th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr v-for="stock in stockStore.sortedStocks" :key="stock.ID">
                <td
                  class="whitespace-nowrap py-4 pl-4 pr-3 text-sm font-medium text-gray-900 sm:pl-0"
                >
                  <RouterLink
                    :to="{
                      name: 'stock-detail',
                      params: { symbol: stock.ticker },
                    }"
                    class="text-indigo-600 hover:text-indigo-900"
                  >
                    {{ stock.ticker }}
                  </RouterLink>
                </td>
                <td class="whitespace-nowrap px-3 py-4 text-sm text-gray-500">
                  {{ stock.company }}
                </td>
                <td
                  class="whitespace-nowrap px-3 py-4 text-sm text-right text-gray-500"
                >
                  ${{ stock.target_to.toFixed(2) }}
                </td>
                <td
                  class="whitespace-nowrap px-3 py-4 text-sm text-right"
                  :class="
                    stock.target_to >= stock.target_from
                      ? 'text-green-600'
                      : 'text-red-600'
                  "
                >
                  {{
                    (
                      ((stock.target_to - stock.target_from) /
                        stock.target_from) *
                      100
                    ).toFixed(2)
                  }}%
                </td>
                <td
                  class="whitespace-nowrap px-3 py-4 text-sm text-right text-gray-500"
                >
                  {{ stock.brokerage }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>
