import { mount } from '@vue/test-utils'
import DeploymentTable from '@/components/DeploymentTable.vue'

import { deploymentArray } from '@/example-data/deployments'

describe('DeploymentTable.vue', () => {
  test('renders a row for each deployment', () => {
    const wrapper = mount(DeploymentTable, {
      props: {
        deployments: deploymentArray
      }
    })

    expect(wrapper.findAll('tbody tr').length).toEqual(deploymentArray.length)
  })
})
