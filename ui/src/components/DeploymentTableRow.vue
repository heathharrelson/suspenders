<template>
  <tr>
    <td class="text-sm">
      {{ deployment.metadata.namespace }}/{{ deployment.metadata.name }}
    </td>
    <td class="text-sm">{{ deployment.spec.replicas }}</td>
    <td class="text-sm">{{ deployment.status.readyReplicas }}</td>
    <td class="text-sm">{{ deployment.status.updatedReplicas }}</td>
    <td class="text-sm">
      <time
        :datetime="deployment.metadata.creationTimestamp"
        :title="deployment.metadata.creationTimestamp"
      >
        {{ formatRelative(deployment.metadata.creationTimestamp) }}
      </time>
    </td>
    <td class="text-sm">
      {{ latestCondition.type }}:{{ latestCondition.status }}
    </td>
    <td class="text-sm">
      <time
        :datetime="latestCondition.lastTransitionTime"
        :title="latestCondition.lastTransitionTime"
      >
        {{ formatRelative(latestCondition.lastTransitionTime) }}
      </time>
    </td>
  </tr>
</template>

<script>
import { computed } from 'vue'
import formatDistanceToNow from 'date-fns/formatDistanceToNow'

export default {
  props: {
    deployment: Object
  },

  setup(props) {
    const formatRelative = (timestamp) => {
      const d = new Date(timestamp)
      return formatDistanceToNow(d, { addSuffix: true })
    }

    const creationTimestamp = computed(
      () => props.deployment.value.metadata.creationTimestamp
    )

    const latestCondition = computed(() => {
      const conditions = props.deployment.status.conditions
      return conditions.reduce((max, curr) =>
        max.lastTransitionTime < curr.lastTransitionTime ? curr : max
      )
    })

    return {
      creationTimestamp,
      latestCondition,
      formatRelative
    }
  }
}
</script>
