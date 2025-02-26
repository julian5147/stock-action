<script setup lang="ts">
import { onMounted } from "vue";
import { RouterLink } from "vue-router";
import ErrorMessage from "../components/common/ErrorMessage.vue";
import LoadingSpinner from "../components/common/LoadingSpinner.vue";
import PriceChange from "../components/common/PriceChange.vue";
import { useStockStore } from "../stores/stocks";

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
          <LoadingSpinner v-if="stockStore.loading" />

          <ErrorMessage
            v-else-if="stockStore.error"
            :message="stockStore.error"
          />

          <table v-else class="min-w-full divide-y divide-gray-300">
            <thead>
              <tr>
                <th
                  scope="col"
                  class="py-3.5 pl-4 pr-3 text-left text-sm font-semibold text-gray-900 sm:pl-0"
                >
                  Ticker
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                >
                  Company
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                >
                  Price
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-center text-sm font-semibold text-gray-900"
                >
                  Change
                </th>
                <th
                  scope="col"
                  class="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                >
                  Broker
                </th>
              </tr>
            </thead>
            <tbody class="divide-y divide-gray-200">
              <tr v-for="stock in stockStore.sortedStocks" :key="stock.id">
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
                <td
                  class="whitespace-nowrap px-3 py-4 text-sm text-left text-gray-500"
                >
                  {{ stock.company }}
                </td>
                <td
                  class="whitespace-nowrap px-3 py-4 text-sm text-left text-gray-500"
                >
                  ${{ stock.target_to.toFixed(2) }}
                </td>
                <td class="whitespace-nowrap px-3 py-4 text-sm text-right">
                  <PriceChange
                    :from-value="stock.target_from"
                    :to-value="stock.target_to"
                  />
                </td>
                <td
                  class="whitespace-nowrap px-3 py-4 text-sm text-left text-gray-500"
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
