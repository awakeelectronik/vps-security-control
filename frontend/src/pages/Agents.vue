<template>
  <div class="space-y-8">
    <div>
      <h2 class="text-3xl font-bold">Managed Agents</h2>
      <p class="text-gray-400">VPS servers reporting security events</p>
    </div>

    <div class="bg-gray-800 rounded-lg p-6 border border-gray-700">
      <!-- Register Agent Form -->
      <div class="mb-8 pb-8 border-b border-gray-700">
        <h3 class="text-xl font-bold mb-4">Register New Agent</h3>
        <form @submit.prevent="registerAgent" class="flex gap-4">
          <input
            v-model="newAgent.name"
            type="text"
            placeholder="Agent name (e.g., production-vps-1)"
            class="flex-1 px-4 py-2 bg-gray-700 rounded border border-gray-600 focus:outline-none focus:border-red-500"
            required
          />
          <input
            v-model="newAgent.ip_address"
            type="text"
            placeholder="IP address"
            class="flex-1 px-4 py-2 bg-gray-700 rounded border border-gray-600 focus:outline-none focus:border-red-500"
            required
          />
          <button
            type="submit"
            class="px-6 py-2 bg-green-600 hover:bg-green-700 rounded font-semibold transition"
          >
            Register
          </button>
        </form>
      </div>

      <!-- Agents List -->
      <div v-if="loading" class="text-center py-8">Loading...</div>
      <div v-else-if="agents.length === 0" class="text-center py-8 text-gray-400">
        No agents registered yet
      </div>
      <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-4">
        <div
          v-for="agent in agents"
          :key="agent.id"
          class="bg-gray-700 rounded-lg p-4 border" :class="agent.status === 'active' ? 'border-green-500' : 'border-red-500'"
        >
          <div class="flex justify-between items-start mb-2">
            <h4 class="text-lg font-bold">{{ agent.name }}</h4>
            <span class="px-2 py-1 rounded text-xs font-semibold" :class="agent.status === 'active' ? 'bg-green-900 text-green-200' : 'bg-red-900 text-red-200'">
              {{ agent.status }}
            </span>
          </div>
          <p class="text-gray-400 text-sm mb-2">IP: {{ agent.ip_address }}</p>
          <p class="text-gray-400 text-sm mb-4">UUID: {{ agent.uuid }}</p>
          <div class="flex gap-2">
            <button
              @click="deleteAgent(agent.id)"
              class="px-4 py-1 bg-red-600 hover:bg-red-700 rounded text-sm transition"
            >
              Delete
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

interface Agent {
  id: number
  uuid: string
  name: string
  ip_address: string
  status: string
  last_seen: string
  created_at: string
}

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

const agents = ref<Agent[]>([])
const loading = ref(true)
const newAgent = ref({
  name: '',
  ip_address: '',
})

const fetchAgents = async () => {
  try {
    loading.value = true
    const res = await axios.get(`${API_URL}/v1/agents`)
    agents.value = res.data.data || []
  } catch (error) {
    console.error('Error fetching agents:', error)
  } finally {
    loading.value = false
  }
}

const registerAgent = async () => {
  try {
    await axios.post(`${API_URL}/v1/agents/register`, {
      name: newAgent.value.name,
      ip_address: newAgent.value.ip_address,
    })
    newAgent.value = { name: '', ip_address: '' }
    fetchAgents()
  } catch (error) {
    console.error('Error registering agent:', error)
    alert('Failed to register agent')
  }
}

const deleteAgent = async (id: number) => {
  if (!confirm('Are you sure?')) return
  try {
    await axios.delete(`${API_URL}/v1/agents/${id}`)
    fetchAgents()
  } catch (error) {
    console.error('Error deleting agent:', error)
    alert('Failed to delete agent')
  }
}

onMounted(() => {
  fetchAgents()
})
</script>
