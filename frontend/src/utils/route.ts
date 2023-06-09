import { NextRouter } from 'next/router'
import { PublicRoutes } from '@/config/route'

export function checkPublicRoute(router: NextRouter) {
  const value = PublicRoutes.includes(router.pathname)
  return value
}
