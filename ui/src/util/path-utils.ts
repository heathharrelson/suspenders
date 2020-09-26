/**
 * Extracts the application's context path from the <meta> tag if present. The path will
 * be normalized by removing any trailing slash. If the context path is missing or unset,
 * the root context "/" is assumed.
 */
export function getContextPath(): string {
  const defaultContext = '/'

  const contextTag = document.querySelector('meta[name="context-path"]')
  if (contextTag === null) {
    return defaultContext
  }

  let contextPath = contextTag.getAttribute('content')
  if (contextPath && contextPath.endsWith('/')) {
    contextPath = contextPath.slice(0, -1)
  }

  return contextPath || defaultContext
}

/**
 * Computes the base path of API endpoints.
 */
export function getApiPath(): string {
  const apiPrefix = '/api/v1'
  const contextPath = getContextPath()
  return contextPath === '/' ? apiPrefix : contextPath + apiPrefix
}
