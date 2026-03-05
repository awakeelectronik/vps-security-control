<template>
  <div class="space-y-8">
    <div>
      <h2 class="text-3xl font-bold">Security Events</h2>
      <p class="text-gray-400">All recorded security events and incidents</p>
    </div>

    <div class="bg-gray-800 rounded-lg p-6 border border-gray-700">
      <!-- Filters -->
      <div class="mb-6 flex gap-4">
        <input
          v-model="filters.eventType"
          type="text"
          placeholder="Filter by event type..."
          class="flex-1 px-4 py-2 bg-gray-700 rounded border border-gray-600 focus:outline-none focus:border-red-500"
        />
        <button
          @click="fetchEvents"
          class="px-6 py-2 bg-red-600 hover:bg-red-700 rounded font-semibold transition"
        >
          Search
        </button>
      </div>

      <!-- Table -->
      <div v-if="loading" class="text-center py-8">Loading...</div>
      <div v-else-if="events.length === 0" class="text-center py-8 text-gray-400">
        No events found
      </div>
      <div v-else class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead class="border-b border-gray-700">
            <tr>
              <th class="text-left py-2">ID</th>
              <th class="text-left py-2">Type</th>
              <th class="text-left py-2">Source IP</th>
              <th class="text-left py-2">Username</th>
              <th class="text-left py-2">Service</th>
              <th class="text-left py-2">Severity</th>
              <th class="text-left py-2">Time</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="event in events" :key="event.id" class="border-b border-gray-700 hover:bg-gray-700 transition">
              <td class="py-2">{{ event.id }}</td>
              <td class="py-2">{{ event.event_type }}</td>
              <td class="py-2">{{ event.source_ip || 'N/A' }}</td>
              <td class="py-2">{{ event.username || 'N/A' }}</td>
              <td class="py-2">{{ event.service || 'N/A' }}</td>
              <td class="py-2">
                <span class="px-2 py-1 rounded text-xs" :class="getSeverityClass(event.severity)">
                  {{ event.severity }}
                </span>
              </td>
              <td class="py-2 text-gray-400">{{ formatTime(event.created_at) }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

interface SecurityEvent {
  id: number
  event_type: string
  source_ip?: string
  username?: string
  service?: string
  severity: string
  created_at: string
}

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

const events = ref<SecurityEvent[]>([])
const loading = ref(true)
const filters = ref({
  eventType: '',
})

const fetchEvents = async () => {
  try {
    loading.value = true
    const res = await axios.get(`${API_URL}/v1/security/events`, {
      params: {
        limit: 100,
      },
    })
    events.value = res.data.data || []
  } catch (error) {
    console.error('Error fetching events:', error)
  } finally {
    loading.value = false
  }
}

const formatTime = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleString()
}

const getSeverityClass = (severity: string) => {
  const classes: Record<string, string> = {
    critical: 'bg-red-900 text-red-200',
    high: 'bg-orange-900 text-orange-200',
    medium: 'bg-yellow-900 text-yellow-200',
    low: 'bg-blue-900 text-blue-200',
  }
  return classes[severity] || 'bg-gray-700 text-gray-200'
}

onMounted(() => {
  fetchEvents()
})
</script>
