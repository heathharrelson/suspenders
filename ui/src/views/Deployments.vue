<template>
  <div class="max-w-7xl py-6 sm:px-6 lg:px-8">
    <div class="py-6 sm:px-0">
      <deployment-table :deployments="deployments" />
    </div>
  </div>
</template>

<script lang="ts">
import DeploymentTable from '../components/DeploymentTable.vue'
import { useDeployments } from '../composition/use-deployments'

export default {
  components: {
    DeploymentTable
  },

  async setup() {
    const { deployments, getDeployments } = useDeployments()
    await getDeployments()
    return {
      deployments
    }
  }
}
</script>
