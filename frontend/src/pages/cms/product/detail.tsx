import { GetStaticProps } from 'next'
import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import PageContentContainer from '@/components/PageContentContainer/PageContentContainer'
import { useRouter } from 'next/router'
import useTranslationData from '@/hooks/useTranslation'
import { ActionIcon, Flex, Title } from '@mantine/core'
import ProductEditForm from '@/components/forms/ProductEditForm'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faArrowLeft } from '@fortawesome/free-solid-svg-icons'

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

export default function ProductDetailPage() {
  const router = useRouter()
  const { t } = useTranslationData()
  const productId = router.query.id
  const pathToBack = '/cms/product/list'

  const handleOnClickBackButton = (
    event: React.MouseEvent<HTMLButtonElement>
  ) => {
    router.push(pathToBack)
  }

  return (
    <PageContentContainer>
      <Flex mb="md" align={'center'} justify="space-between">
        <ActionIcon size="xl" radius="xl" onClick={handleOnClickBackButton}>
          <FontAwesomeIcon icon={faArrowLeft} />
        </ActionIcon>
        <Title color={'gray'} order={2}>
          {t('product.product_detail')}
        </Title>
      </Flex>
      <ProductEditForm productId={productId as string} />
    </PageContentContainer>
  )
}
