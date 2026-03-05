<template>
  <div class="space-y-8">
    <!-- Header -->
    <div>
      <h2 class="text-3xl font-bold">Security Dashboard</h2>
      <p class="text-gray-400">Real-time monitoring of VPS security events</p>
    </div>

    <!-- Stats Grid -->
    <div class="grid grid-cols-1 md:grid-cols-4 gap-4">
      <StatCard title="Failed Logins" :value="stats.failedLogins" icon="🔓" color="red" />
      <StatCard title="Detected Attacks" :value="stats.attacks" icon="🔴" color="orange" />
      <StatCard title="Active Agents" :value="stats.agents" icon="🖥️" color="blue" />
      <StatCard title="System Health" :value="stats.health" icon="💚" color="green" />
    </div>

    <!-- Recent Events -->
    <div class="bg-gray-800 rounded-lg p-6 border border-gray-700">
      <h3 class="text-xl font-bold mb-4">Recent Security Events</h3>
      <div v-if="loading" class="text-center py-8">Loading...</div>
      <div v-else-if="events.length === 0" class="text-center py-8 text-gray-400">
        No events recorded
      </div>
      <div v-else class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead class="border-b border-gray-700">
            <tr>
              <th class="text-left py-2">Type</th>
              <th class="text-left py-2">Source IP</th>
              <th class="text-left py-2">Country</th>
              <th class="text-left py-2">Severity</th>
              <th class="text-left py-2">Time</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="event in events" :key="event.id" class="border-b border-gray-700 hover:bg-gray-700 transition">
              <td class="py-2">
                <span class="px-2 py-1 rounded text-xs font-semibold" :class="getEventTypeClass(event.event_type)">
                  {{ event.event_type }}
                </span>
              </td>
              <td class="py-2">{{ event.source_ip || 'N/A' }}</td>
              <td class="py-2">{{ event.source_country || 'N/A' }}</td>
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
import StatCard from '../components/StatCard.vue'

interface Event {
  id: number
  event_type: string
  source_ip?: string
  source_country?: string
  severity: string
  created_at: string
}

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

const events = ref<Event[]>([])
const loading = ref(true)
const stats = ref({
  failedLogins: 0,
  attacks: 0,
  agents: 0,
  health: '100%',
})

const fetchData = async () => {
  try {
    loading.value = true
    
    const [eventsRes] = await Promise.all([
      axios.get(`${API_URL}/v1/security/events?limit=10`),
    ])
    
    events.value = eventsRes.data.data || []
    stats.value.failedLogins = events.value.filter((e: Event) => e.event_type === 'failed_login').length
    stats.value.attacks = events.value.filter((e: Event) => ['brute_force', 'port_scan', 'dos_attack'].includes(e.event_type)).length
  } catch (error) {
    console.error('Error fetching data:', error)
  } finally {
    loading.value = false
  }
}

const formatTime = (dateString: string) => {
  const date = new Date(dateString)
  return date.toLocaleString()
}

const getEventTypeClass = (type: string) => {
  const classes: Record<string, string> = {
    failed_login: 'bg-yellow-900 text-yellow-200',
    brute_force: 'bg-red-900 text-red-200',
    port_scan: 'bg-orange-900 text-orange-200',
    dos_attack: 'bg-red-900 text-red-200',
  }
  return classes[type] || 'bg-gray-700 text-gray-200'
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
  fetchData()
  // Refresh data every 30 seconds
  setInterval(fetchData, 30000)
})
</script>
