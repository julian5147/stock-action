<template>
  <span
    class="inline-flex items-center rounded-md px-2.5 py-0.5 text-sm font-medium"
    :class="{
      'bg-emerald-100 text-emerald-800': recommendation === 'Strong Buy',
      'bg-green-100 text-green-800': recommendation === 'Buy',
      'bg-yellow-100 text-yellow-800': recommendation === 'Hold',
      'bg-orange-100 text-orange-800': recommendation === 'Sell',
      'bg-red-100 text-red-800': recommendation === 'Strong Sell',
    }"
  >
    {{ recommendation }}
  </span>
</template>

<script setup lang="ts">
defineProps<{
  recommendation: string;
}>();
</script>
