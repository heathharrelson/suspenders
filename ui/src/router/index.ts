import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Deployments from '../views/Deployments.vue'
import { getContextPath } from '../util/path-utils'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Deployments',
    component: Deployments
  }
]

function buildRouter() {
  const contextPath = getContextPath()
  return createRouter({
    history: createWebHistory(contextPath),
    routes
  })
}

export default buildRouter
