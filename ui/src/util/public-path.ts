/**
 * Updates the Webpack public path to match the application's context path
 * so that dynamic imports work correctly. Entrypoint modules should import
 * this module before importing any others.
 */
import { getContextPath } from './path-utils'

// eslint-disable-next-line @typescript-eslint/camelcase
__webpack_public_path__ = getContextPath()
