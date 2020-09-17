/**
 * Extracts the application's context path from the <meta> tag if present. If missing
 * or unset, the root context "/" is assumed.
 */
export function getContextPath(): string {
  const defaultContext = '/'

  const contextTag = document.querySelector('meta[name="context-path"]')
  if (contextTag === null) {
    return defaultContext
  }

  return contextTag.getAttribute('content') || defaultContext
}

/**
 * Computes the base path of API endpoints.
 */
export function getApiPath(): string {
  const apiPrefix = '/api/v1'
  const contextPath = getContextPath()
  return contextPath === '/' ? apiPrefix : contextPath + apiPrefix
}
