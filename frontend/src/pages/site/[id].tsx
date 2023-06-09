import { useEffect } from 'react'
import { GetServerSideProps } from 'next'
import { useRouter } from 'next/router'
import { get } from 'lodash'
import { SSRConfig } from 'next-i18next'
import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import { ParsedUrlQuery } from 'querystring'
import UserSiteProvider from '@/providers/UserSiteProvider/UserSiteProvider'
import {
  createUpdateHomepageImgUrlAction,
  createUpdateHomepageTextAction,
  createUpdateHomepageTextColorAction,
  createUpdateSitenameAction,
  IUserSiteStateUiAction,
  UserSiteStateActionType,
} from '@/providers/UserSiteProvider/actions'
import { IProduct } from '@/entities/product'
import { ISiteUi } from '@/entities/site'
import { API, StatusCode } from '@/services'
import { getDefaultUiData } from '@/mocks'
import useUserSite from '@/hooks/userUserSite'
import SiteTemplate from '@/components/site/SiteTemplate/SiteTemplate'

/** getServerSideProps (NextJS)
 * only on server-side
 * @param locale the type of locale
 * @returns
 */
interface IParams extends ParsedUrlQuery {
  id: string
}
export const getServerSideProps: GetServerSideProps = async (context) => {
  const { locale } = context!
  const { id } = context.params as IParams

  // check user exist
  const resIsSiteExit = await API.checkUserExistById(id)
  if (resIsSiteExit.status !== StatusCode.Ok) {
    return {
      notFound: true,
    }
  }

  // get user ui data
  const resUiData = await API.getUserSiteUiDataById(id)
  let userCustomUIData: ISiteUi
  if (resUiData.status === StatusCode.Ok) {
    userCustomUIData = resUiData.data.data
  } else {
    // set default
    userCustomUIData = getDefaultUiData(id, id)
  }

  // get user's product
  const resProducts = await API.getProducts(id, { page: 1, limit: 1000 })
  if (!resProducts.success) {
    return {
      notFound: true,
    }
  }

  const productsData: IProduct[] = get(resProducts, 'data.data.products', [])

  const props: Omit<ISitePageProps, '_nextI18Next'> = {
    // will be passed to the page component as props
    ...(await serverSideTranslations(locale as string, ['data'])),
    userId: id,
    products: productsData,
    siteUiData: userCustomUIData,
  }
  return {
    props: props,
  }
}

export interface ISitePageProps {
  _nextI18Next: SSRConfig
  userId: string
  products: IProduct[]
  siteUiData: ISiteUi
}
function SitePage(props: ISitePageProps) {
  const { userSiteState, userSiteStateDispatch, resetUiData } = useUserSite()
  useEffect(() => {
    const handleReceiveActionFromIframeParent = (
      e: MessageEvent<IUserSiteStateUiAction>
    ) => {
      // console.log('ee data:', e.data)
      // if (e.origin !== 'http://localhost:3000') return
      const origin = process.env.NEXT_PUBLIC_HOST || 'http://localhost:3000'
      if (e.origin !== origin) return
      if (e.data.type === UserSiteStateActionType.UpdateSitename) {
        userSiteStateDispatch(createUpdateSitenameAction(e.data.payload))
      } else if (e.data.type === UserSiteStateActionType.UpdateHomepageImgUrl) {
        userSiteStateDispatch(createUpdateHomepageImgUrlAction(e.data.payload))
      } else if (e.data.type === UserSiteStateActionType.UpdateHomepageText) {
        userSiteStateDispatch(createUpdateHomepageTextAction(e.data.payload))
      } else if (e.data.type === UserSiteStateActionType.ResetUiData) {
        resetUiData()
      } else if (
        e.data.type === UserSiteStateActionType.UpdateHomepageTextColor
      ) {
        userSiteStateDispatch(
          createUpdateHomepageTextColorAction(e.data.payload)
        )
      }
    }

    if (userSiteState.editMode) {
      window.addEventListener('message', handleReceiveActionFromIframeParent)
    }
    return () =>
      window.removeEventListener('message', handleReceiveActionFromIframeParent)
  }, [userSiteState.editMode]) // eslint-disable-line react-hooks/exhaustive-deps

  return <SiteTemplate {...props} siteUiData={userSiteState.uiData} />
}
export default function SitePageWrapper(props: ISitePageProps) {
  const router = useRouter()
  const editMode = router.query.editMode === 'true'
  return (
    <UserSiteProvider
      userId={props.userId}
      products={props.products}
      userUiData={props.siteUiData}
      editMode={editMode}
    >
      <SitePage {...props} />
    </UserSiteProvider>
  )
}
