import { ref, readonly } from 'vue'
import { getApiPath } from '../util/path-utils'

const deployments = ref([])

/**
 * Updates state from the API endpoint.
 */
async function getDeployments() {
  const response = await fetch(`${getApiPath()}/deployments`)

  if (!response.ok) {
    throw new Error(`Failed to get deployments: ${response.statusText}`)
  }

  deployments.value = await response.json()
}

/**
 * Composable that returns deployment state for use in components.
 */
export function useDeployments() {
  return {
    deployments: readonly(deployments),
    getDeployments
  }
}
