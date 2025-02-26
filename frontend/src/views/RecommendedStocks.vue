<script setup lang="ts">
import ErrorMessage from "@/components/common/ErrorMessage.vue";
import LoadingSpinner from "@/components/common/LoadingSpinner.vue";
import StockRecommendationCard from "@/components/StockRecommendationCard.vue";
import { onMounted } from "vue";
import { useStockStore } from "../stores/stocks";

const stockStore = useStockStore();

onMounted(async () => {
  await stockStore.fetchRecommendedStocks();
});
</script>

<template>
  <div>
    <div class="sm:flex sm:items-center">
      <div class="sm:flex-auto">
        <h1 class="text-2xl font-semibold text-gray-900">
          Acciones Recomendadas
        </h1>
        <p class="mt-2 text-sm text-gray-700">
          Lista de acciones recomendadas para invertir basadas en nuestro
          análisis.
        </p>
      </div>
    </div>

    <div class="mt-8 flow-root">
      <div class="-mx-4 -my-2 overflow-x-auto sm:-mx-6 lg:-mx-14">
        <div class="inline-block min-w-full py-2 align-middle sm:px-6 lg:px-8">
          <LoadingSpinner v-if="stockStore.loading" />

          <ErrorMessage
            v-else-if="stockStore.error"
            :message="stockStore.error"
          />

          <div v-else class="space-y-6">
            <StockRecommendationCard
              v-for="recommendation in stockStore.recommendedStocksList"
              :key="recommendation.stock.id"
              :stock="recommendation.stock"
              :recommendation="recommendation.recommendation"
              :indicators="recommendation.indicators"
              :score="recommendation.score"
            />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
