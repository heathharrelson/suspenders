<template>
  <table class="w-full text-left table-collapse">
    <thead>
      <tr>
        <th class="text-sm font-semibold">Name</th>
        <th class="text-sm font-semibold">Desired</th>
        <th class="text-sm font-semibold">Ready</th>
        <th class="text-sm font-semibold">Up to Date</th>
        <th class="text-sm font-semibold">Created</th>
        <th class="text-sm font-semibold">Status</th>
        <th class="text-sm font-semibold">Last Transition</th>
      </tr>
    </thead>
    <tbody>
      <deployment-table-row
        v-for="deployment of deployments"
        :key="deploymentKey(deployment)"
        :deployment="deployment"
      />
    </tbody>
  </table>
</template>

<script lang="ts">
import DeploymentTableRow from './DeploymentTableRow.vue'

export default {
  props: {
    deployments: Array
  },

  components: {
    DeploymentTableRow
  },

  setup() {
    // TODO: proper type annotation
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    const deploymentKey = (deployment: any) => {
      return `${deployment.metadata.namespace}${deployment.metadata.name}`
    }

    return {
      deploymentKey
    }
  }
}
</script>
