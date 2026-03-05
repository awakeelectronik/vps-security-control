<template>
  <div class="bg-gray-800 rounded-lg p-6 border border-gray-700">
    <div class="flex justify-between items-start">
      <div>
        <p class="text-gray-400 text-sm">{{ title }}</p>
        <p class="text-3xl font-bold mt-2" :class="colorClass">{{ value }}</p>
      </div>
      <span class="text-4xl">{{ icon }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

interface Props {
  title: string
  value: string | number
  icon: string
  color: 'red' | 'orange' | 'blue' | 'green'
}

const props = defineProps<Props>()

const colorClass = computed(() => {
  const colors: Record<string, string> = {
    red: 'text-red-400',
    orange: 'text-orange-400',
    blue: 'text-blue-400',
    green: 'text-green-400',
  }
  return colors[props.color]
})
</script>
