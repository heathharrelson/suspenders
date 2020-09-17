import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Deployments from '../views/Deployments.vue'
import { getContextPath } from '../util/path-utils'

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Deployments',
    component: Deployments
  },
  {
    path: '/about',
    name: 'About',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ '../views/About.vue')
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
