import { createRouter, createWebHistory } from 'vue-router'
import Dashboard from '../pages/Dashboard.vue'
import Events from '../pages/Events.vue'
import Agents from '../pages/Agents.vue'

const routes = [
  {
    path: '/',
    name: 'Dashboard',
    component: Dashboard,
  },
  {
    path: '/events',
    name: 'Events',
    component: Events,
  },
  {
    path: '/agents',
    name: 'Agents',
    component: Agents,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
