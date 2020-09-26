import { getApiPath, getContextPath } from '@/util/path-utils'

function createMetaTag(path: string) {
  const metaNode = document.createElement('meta')
  metaNode.setAttribute('name', 'context-path')
  metaNode.setAttribute('content', path)

  document.head.appendChild(metaNode)
}

describe('path utils', () => {
  afterEach(() => {
    document.head.innerHTML = ''
  })

  describe('getContextPath', () => {
    test('returns / by default', () => {
      expect(getContextPath()).toEqual('/')
    })

    test('returns meta tag value if present', () => {
      const expectedPath = '/test-path'
      createMetaTag(expectedPath)
      expect(getContextPath()).toEqual(expectedPath)
    })

    test('allows meta tag containing /', () => {
      const expectedPath = '/'
      createMetaTag(expectedPath)
      expect(getContextPath()).toEqual(expectedPath)
    })

    test('normalizes path by trimming trailing slash', () => {
      const expectedPath = '/test-path'
      createMetaTag(expectedPath + '/')
      expect(getContextPath()).toEqual(expectedPath)
    })
  })

  describe('getApiPath', () => {
    test('returns /api/v1 by default', () => {
      expect(getApiPath()).toEqual('/api/v1')
    })

    test('returns /api/v1 when running at the root context', () => {
      createMetaTag('/')
      expect(getApiPath()).toEqual('/api/v1')
    })

    test('respects the context path when set', () => {
      const contextPath = '/test-path'
      createMetaTag(contextPath)
      expect(getApiPath()).toEqual(contextPath + '/api/v1')
    })
  })
})
