<template>
  <div>
    <dl class="text-md font-semibold text-gray-900 mb-2">Indicadores</dl>
    <dd class="space-y-2">
      <div class="flex justify-between items-center">
        <span class="text-sm text-gray-600">Confianza del broker</span>
        <span :class="getIndicatorColor(indicators.broker_confidence)">
          {{ (indicators.broker_confidence * 100).toFixed(0) }}%
        </span>
      </div>
      <div class="flex justify-between items-center">
        <span class="text-sm text-gray-600"
          >Crecimiento del precio objetivo</span
        >
        <span
          :class="
            indicators.price_target_growth > 0
              ? 'text-green-600'
              : 'text-red-600'
          "
        >
          {{ formatPercentage(indicators.price_target_growth) }}
        </span>
      </div>
      <div class="flex justify-between items-center">
        <span class="text-sm text-gray-600">Impacto de calificación</span>
        <span :class="getIndicatorColor(indicators.rating_impact + 0.5)">
          {{ formatPercentage(indicators.rating_impact * 100) }}
        </span>
      </div>
    </dd>
  </div>
</template>

<script setup lang="ts">
interface Indicators {
  broker_confidence: number;
  price_target_growth: number;
  rating_impact: number;
}

defineProps<{
  indicators: Indicators;
}>();

function getIndicatorColor(value: number): string {
  if (value > 0.7) return "text-green-600";
  if (value > 0.3) return "text-yellow-600";
  return "text-red-600";
}

function formatPercentage(value: number): string {
  return `${value.toFixed(2)}%`;
}
</script>
