<script setup lang="ts">
import { onMounted } from "vue";
import { useStockStore } from "../stores/stocks";
import { RouterLink } from "vue-router";

const stockStore = useStockStore();

onMounted(async () => {
  await stockStore.fetchRecommendedStocks();
});

function getIndicatorColor(value: number): string {
  if (value > 0.7) return "text-green-600";
  if (value > 0.3) return "text-yellow-600";
  return "text-red-600";
}

function formatPercentage(value: number): string {
  return `${value.toFixed(2)}%`;
}
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

          <div v-else class="space-y-6">
            <div
              v-for="recommendation in stockStore.recommendedStocksList"
              :key="recommendation.stock.id"
              class="bg-white shadow rounded-lg overflow-hidden"
            >
              <!-- Encabezado -->
              <div class="px-4 py-5 sm:px-6 flex justify-between items-start">
                <div>
                  <div class="flex items-center">
                    <RouterLink
                      :to="{
                        name: 'stock-detail',
                        params: { symbol: recommendation.stock.ticker },
                      }"
                      class="text-lg font-medium text-indigo-600 hover:text-indigo-900"
                    >
                      {{ recommendation.stock.ticker }}
                    </RouterLink>
                    <span class="ml-2 text-sm text-gray-500">{{
                      recommendation.stock.company
                    }}</span>
                  </div>
                  <p class="mt-1 text-sm text-gray-500">
                    {{ recommendation.stock.brokerage }} -
                    {{ recommendation.stock.action }}
                  </p>
                </div>
                <span
                  class="inline-flex items-center rounded-md px-2.5 py-0.5 text-sm font-medium"
                  :class="{
                    'bg-emerald-100 text-emerald-800':
                      recommendation.recommendation === 'Strong Buy',
                    'bg-green-100 text-green-800':
                      recommendation.recommendation === 'Buy',
                    'bg-yellow-100 text-yellow-800':
                      recommendation.recommendation === 'Hold',
                    'bg-orange-100 text-orange-800':
                      recommendation.recommendation === 'Sell',
                    'bg-red-100 text-red-800':
                      recommendation.recommendation === 'Strong Sell',
                  }"
                >
                  {{ recommendation.recommendation }}
                </span>
              </div>

              <!-- Details -->
              <div class="border-t border-gray-200">
                <div
                  class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 px-4 py-5 sm:px-6"
                >
                  <!-- Price target -->
                  <div>
                    <dl class="text-sm font-medium text-gray-500">
                      Precio objetivo
                    </dl>
                    <dd class="mt-1 text-lg font-semibold text-gray-900">
                      ${{ recommendation.stock.target_to.toFixed(2) }}
                    </dd>
                    <dd
                      class="text-sm"
                      :class="
                        recommendation.stock.target_to >=
                        recommendation.stock.target_from
                          ? 'text-green-600'
                          : 'text-red-600'
                      "
                    >
                      {{
                        formatPercentage(
                          ((recommendation.stock.target_to -
                            recommendation.stock.target_from) /
                            recommendation.stock.target_from) *
                            100
                        )
                      }}
                    </dd>
                  </div>

                  <!-- Score -->
                  <div>
                    <dl class="text-sm font-medium text-gray-500">Score</dl>
                    <dd class="mt-1 text-lg font-semibold text-gray-900">
                      {{ (recommendation.score * 100).toFixed(0) }}%
                    </dd>
                  </div>

                  <!-- Indicators -->
                  <div class="sm:col-span-2">
                    <dl class="text-sm font-medium text-gray-500 mb-2">
                      Indicadores
                    </dl>
                    <dd class="space-y-2">
                      <div class="flex justify-between items-center">
                        <span class="text-sm text-gray-600"
                          >Confianza del broker</span
                        >
                        <span
                          :class="
                            getIndicatorColor(
                              recommendation.indicators.broker_confidence
                            )
                          "
                        >
                          {{
                            (
                              recommendation.indicators.broker_confidence * 100
                            ).toFixed(0)
                          }}%
                        </span>
                      </div>
                      <div class="flex justify-between items-center">
                        <span class="text-sm text-gray-600"
                          >Crecimiento del precio objetivo</span
                        >
                        <span
                          :class="
                            recommendation.indicators.price_target_growth > 0
                              ? 'text-green-600'
                              : 'text-red-600'
                          "
                        >
                          {{
                            formatPercentage(
                              recommendation.indicators.price_target_growth
                            )
                          }}
                        </span>
                      </div>
                      <div class="flex justify-between items-center">
                        <span class="text-sm text-gray-600"
                          >Impacto de calificación</span
                        >
                        <span
                          :class="
                            getIndicatorColor(
                              recommendation.indicators.rating_impact + 0.5
                            )
                          "
                        >
                          {{
                            formatPercentage(
                              recommendation.indicators.rating_impact * 100
                            )
                          }}
                        </span>
                      </div>
                    </dd>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
