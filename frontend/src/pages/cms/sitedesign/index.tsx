import { SyntheticEvent, useRef, useState } from 'react'
import { GetStaticProps } from 'next'
import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import { Grid, LoadingOverlay } from '@mantine/core'
import { SSRConfig } from 'next-i18next'
import UserSiteProvider from '@/providers/UserSiteProvider/UserSiteProvider'
import useAuth from '@/hooks/useAuth'
import useGetProducts from '@/hooks/api/useGetProducts'
import useGetUserUiData from '@/hooks/api/useGetUserUiData'
import useUserSite from '@/hooks/userUserSite'
import SiteIframe from '@/components/site/SiteIframe/SiteIframe'
import DesignPanel from '@/components/site/DesignPanel/DesignPanel'

/** getStaticProps (NextJS)
 * only on server-side
 * @param locale the type of locale
 * @returns
 */
export const getStaticProps: GetStaticProps = async (context) => {
  const { locale } = context!
  const props = {
    // will be passed to the page component as props
    ...(await serverSideTranslations(locale as string, ['data'])),
  }
  return {
    props: props,
  }
}
interface ISiteDesignPage {
  host: string
  refetchSiteUi: () => void
}
function SiteDesignPage(props: ISiteDesignPage) {
  const { host, refetchSiteUi } = props
  const { userSiteState } = useUserSite()
  const iframeRef = useRef<HTMLIFrameElement>(null)
  const [state, setState] = useState({
    isPreparing: true,
    reloadCount: 0,
  })

  const handleOnLoad = (event: SyntheticEvent<HTMLIFrameElement, Event>) => {
    setState((prev) => ({ ...prev, isPreparing: false }))
  }
  const handleReload = () => {
    setState((prev) => ({ ...prev, reloadCount: prev.reloadCount++ }))
  }

  // console.log('userSiteState', userSiteState)
  const sitePath = `${host}/site/${userSiteState.userId}`
  const siteEditPath = `${sitePath}?editMode=true`
  return (
    <Grid h={'100%'} sx={{ position: 'relative' }}>
      <LoadingOverlay
        visible={state.isPreparing}
        overlayBlur={1000}
        transitionDuration={400}
      />
      <Grid.Col span={3} pr="md">
        <DesignPanel
          iframeRef={iframeRef}
          sitePath={sitePath}
          refetchSiteUi={refetchSiteUi}
          reloadIframe={handleReload}
        />
      </Grid.Col>
      <Grid.Col
        span={9}
        sx={{
          border: '1px solid #d3d3d3',
          boxShadow: 'rgba(149, 157, 165, 0.2) 0px 8px 24px',
          borderRadius: '8px',
          overflow: 'hidden',
          padding: '0px',
        }}
      >
        <SiteIframe
          key={state.reloadCount}
          src={siteEditPath}
          onLoad={handleOnLoad}
          ref={iframeRef}
        />
      </Grid.Col>
    </Grid>
  )
}
interface ISiteDesignPageWrapperProps {
  _nextI18Next: SSRConfig
}
export default function SiteDesignPageWrapper(
  props: ISiteDesignPageWrapperProps
) {
  const { user } = useAuth()
  const { products } = useGetProducts({ limit: 1000 })
  const { siteUiData, refetch } = useGetUserUiData(user?.id || '')
  const userId = user?.id!
  const host = process.env.NEXT_PUBLIC_HOST
  console.log('userId', userId)
  if (!siteUiData || !user || !host) {
    return null
  }
  return (
    <UserSiteProvider
      userId={userId}
      products={products}
      userUiData={siteUiData}
      editMode={true}
    >
      <SiteDesignPage host={host} refetchSiteUi={refetch} />
    </UserSiteProvider>
  )
}
