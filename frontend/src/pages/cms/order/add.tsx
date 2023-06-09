import { GetStaticProps } from 'next'
import { useRouter } from 'next/router'
import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import { ActionIcon, Flex, Title } from '@mantine/core'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faArrowLeft } from '@fortawesome/free-solid-svg-icons'
import PageContentContainer from '@/components/PageContentContainer/PageContentContainer'
import useTranslationData from '@/hooks/useTranslation'
import OrderCreateForm from '@/components/forms/OrderCreateForm'

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

export default function OrderAddPage() {
  const router = useRouter()
  const { t } = useTranslationData()
  const pathToBack = '/cms/order/list'

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
          {t('order.add_order')}
        </Title>
      </Flex>
      <OrderCreateForm />
    </PageContentContainer>
  )
}
