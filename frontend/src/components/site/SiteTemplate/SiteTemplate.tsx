import Head from 'next/head'
import { rem } from '@mantine/core'
import { IProduct } from '@/entities/product'
import { ISiteUi } from '@/entities/site'
import SiteHeader from '../SiteHeader/SiteHeader'
import SiteHero from '../SiteHero/SiteHero'
import SiteProducts from '../SiteProducts/SiteProducts'

const HEADER_HEIGHT = rem(60)
const HOME_ANCHOR_ID = 'home'
const PRODUCTS_ANCHOR_ID = 'products'
export interface ISiteTemplateProps {
  userId: string
  products: IProduct[]
  siteUiData: ISiteUi
}
export default function SiteTemplate(props: ISiteTemplateProps) {
  const { userId, products, siteUiData } = props
  const links = [
    { link: HOME_ANCHOR_ID, label: 'Home' },
    { link: PRODUCTS_ANCHOR_ID, label: 'Products' },
  ]

  const customContent = {
    homepageImgUrl: siteUiData.homepageImgUrl,
    homepageText: siteUiData.homepageText,
    homepageTextColor: siteUiData.homepageTextColor,
  }

  return (
    <>
      <Head>
        <title>{siteUiData.sitename}</title>
      </Head>
      <SiteHeader height={HEADER_HEIGHT} links={links} />
      <SiteHero
        id={HOME_ANCHOR_ID}
        scrollMarginTop={HEADER_HEIGHT}
        imgSrc={customContent.homepageImgUrl}
        text={customContent.homepageText}
        textColor={customContent.homepageTextColor}
      />
      <SiteProducts id={PRODUCTS_ANCHOR_ID} products={products} />
    </>
  )
}
