<template>
  <div class="bg-white shadow rounded-lg overflow-hidden">
    <!-- Header -->
    <div class="px-4 py-5 sm:px-6 flex justify-between items-start">
      <div>
        <div class="flex items-center">
          <RouterLink
            :to="{
              name: 'stock-detail',
              params: { symbol: stock.ticker },
            }"
            class="text-lg font-medium text-indigo-600 hover:text-indigo-900"
          >
            {{ stock.ticker }}
          </RouterLink>
          <span class="ml-2 text-sm text-gray-500">{{ stock.company }}</span>
        </div>
        <p class="mt-1 text-sm text-gray-500">
          {{ stock.brokerage }} - {{ stock.action }}
        </p>
      </div>
      <RecommendationBadge :recommendation="recommendation" />
    </div>

    <!-- Details -->
    <div class="border-t border-gray-200">
      <div
        class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 px-4 py-5 sm:px-6"
      >
        <!-- Indicators -->
        <div class="sm:col-span-2">
          <StockIndicators :indicators="indicators" />
        </div>
        <!-- Price target -->
        <div>
          <dl class="text-md font-semibold text-gray-900">Precio objetivo</dl>
          <dd class="mt-1 text-md text-gray-600">
            ${{ stock.target_to.toFixed(2) }}
          </dd>
        </div>

        <!-- Score -->
        <div>
          <dl class="text-md font-semibold text-gray-900">Score</dl>
          <dd class="mt-1 text-md text-gray-600">
            {{ (score * 100).toFixed(0) }}%
          </dd>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { RouterLink } from "vue-router";
import RecommendationBadge from "./RecommendationBadge.vue";
import StockIndicators from "./StockIndicators.vue";

interface Stock {
  id: string;
  ticker: string;
  company: string;
  brokerage: string;
  action: string;
  target_to: number;
  target_from: number;
  rating_from: string;
  rating_to: string;
  time: string;
}

interface Indicators {
  broker_confidence: number;
  price_target_growth: number;
  rating_impact: number;
}

defineProps<{
  stock: Stock;
  recommendation: string;
  indicators: Indicators;
  score: number;
}>();
</script>
