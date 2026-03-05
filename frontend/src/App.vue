<template>
  <div id="app" class="min-h-screen bg-gray-900 text-gray-100">
    <nav class="bg-gray-800 border-b border-gray-700">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16">
          <div class="flex items-center">
            <h1 class="text-2xl font-bold text-red-500">🛡️ VPS Security Control</h1>
          </div>
          <div class="flex items-center space-x-4">
            <router-link to="/" class="hover:text-red-400">Dashboard</router-link>
            <router-link to="/events" class="hover:text-red-400">Events</router-link>
            <router-link to="/agents" class="hover:text-red-400">Agents</router-link>
          </div>
        </div>
      </div>
    </nav>

    <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <router-view></router-view>
    </main>
  </div>
</template>

<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<style scoped>
</style>
