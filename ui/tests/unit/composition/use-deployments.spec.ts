import 'whatwg-fetch'
import { useDeployments } from '@/composition/use-deployments'
import { jsonArray as deploymentArray } from '@/example-data/deployments'

describe('useDeployments composable', () => {
  it('getDeployments calls out to the API on getDeployments', async () => {
    const { getDeployments } = useDeployments()

    const responseInit = {
      status: 200,
      statusText: 'OK',
      headers: new Headers({
        'Content-Type': 'application/json '
      })
    }
    const mockResponse = Promise.resolve(new Response('[]', responseInit))

    jest
      .spyOn(globalThis, 'fetch')
      .mockImplementation((_input, _init) => mockResponse)

    await getDeployments()

    expect(globalThis.fetch).toHaveBeenCalledTimes(1)
    expect(globalThis.fetch).toHaveBeenCalledWith('/api/v1/deployments')
  })

  it('getDeployments updates the deployments on success', async () => {
    const { deployments, getDeployments } = useDeployments()

    const responseInit = {
      status: 200,
      statusText: 'OK',
      headers: new Headers({
        'Content-Type': 'application/json '
      })
    }
    const mockResponse = Promise.resolve(
      new Response(deploymentArray, responseInit)
    )

    jest
      .spyOn(globalThis, 'fetch')
      .mockImplementation((_input, _init) => mockResponse)

    const origValue = deployments.value
    await getDeployments()

    expect(deployments.value).not.toEqual(origValue)
    expect(deployments.value.length).toEqual(1)
  })

  it('getDeployments throws an error if the API request is not successful', async () => {
    const { getDeployments } = useDeployments()

    const msg = 'Error listing deployments'
    const responseInit = {
      status: 500,
      statusText: 'Internal Server Error'
    }
    const mockResponse = Promise.resolve(new Response(msg, responseInit))

    jest
      .spyOn(globalThis, 'fetch')
      .mockImplementation((_input, _init) => mockResponse)

    try {
      await getDeployments()
      fail('call to getDeployments should have thrown an error')
    } catch (e) {
      expect(e.message).toMatch(responseInit.statusText)
    }
  })
})
