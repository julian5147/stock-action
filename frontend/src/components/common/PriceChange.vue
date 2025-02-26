<template>
  <span :class="fromValue <= toValue ? 'text-green-600' : 'text-red-600'">
    {{ calculatePercentageChange().toFixed(2) }}%
  </span>
</template>

<script setup lang="ts">
const props = defineProps<{
  fromValue: number;
  toValue: number;
}>();

const calculatePercentageChange = () => {
  return ((props.toValue - props.fromValue) / props.fromValue) * 100;
};
</script>
