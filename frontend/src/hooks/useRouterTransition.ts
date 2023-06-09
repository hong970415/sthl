import { useEffect } from 'react'
import { useRouter } from 'next/router'
import NProgress from 'nprogress'

export default function useRouterTransition() {
  const router = useRouter()
  useEffect(() => {
    const handleRouteStart = () => NProgress.start()
    const handleRouteDone = () => NProgress.done()

    NProgress.configure({ showSpinner: false })

    router.events.on('routeChangeStart', handleRouteStart)
    router.events.on('routeChangeComplete', handleRouteDone)
    router.events.on('routeChangeError', handleRouteDone)

    return () => {
      // remove the event handler on unmount!
      router.events.off('routeChangeStart', handleRouteStart)
      router.events.off('routeChangeComplete', handleRouteDone)
      router.events.off('routeChangeError', handleRouteDone)
    }
  }, []) // eslint-disable-line react-hooks/exhaustive-deps
}
