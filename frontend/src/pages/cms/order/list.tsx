import { GetStaticProps } from 'next'
import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import PageContentContainer from '@/components/PageContentContainer/PageContentContainer'
import OrdersTable from '@/components/OrdersTable/OrdersTable'

/** getStaticProps (NextJS)
 * only on server-side
 * @param locale the type of locale
 * @returns
 */
export const getStaticProps: GetStaticProps = async (context) => {
  const { locale } = context!
  return {
    props: {
      ...(await serverSideTranslations(locale as string, ['data'])),
      // will be passed to the page component as props
    },
  }
}

export default function OrderListPage() {
  return (
    <PageContentContainer>
      <OrdersTable />
    </PageContentContainer>
  )
}
