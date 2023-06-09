import { GetStaticProps } from 'next'
import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import { Accordion, Title } from '@mantine/core'
import useTranslationData from '@/hooks/useTranslation'

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

export default function DashboardPage() {
  const { t } = useTranslationData()
  const titleOrder = 4
  const content = (
    <Accordion multiple defaultValue={['item-1', 'item-2']} variant="contained">
      <Accordion.Item value="item-1">
        <Accordion.Control>
          <Title order={titleOrder}>{t('sidebar.album')}</Title>
        </Accordion.Control>
        <Accordion.Panel>{t('dashboard.album_content')}</Accordion.Panel>
      </Accordion.Item>

      <Accordion.Item value="item-2">
        <Accordion.Control>
          <Title order={titleOrder}>{t('sidebar.sitedesign')}</Title>
        </Accordion.Control>
        <Accordion.Panel>{t('dashboard.sitedesign_content')}</Accordion.Panel>
      </Accordion.Item>

      <Accordion.Item value="item-3">
        <Accordion.Control>
          <Title order={titleOrder}>{t('sidebar.product')}</Title>
        </Accordion.Control>
        <Accordion.Panel>{t('dashboard.product_content')}</Accordion.Panel>
      </Accordion.Item>

      <Accordion.Item value="item-4">
        <Accordion.Control>
          <Title order={titleOrder}>{t('sidebar.order')}</Title>
        </Accordion.Control>
        <Accordion.Panel>{t('dashboard.order_content')}</Accordion.Panel>
      </Accordion.Item>

      <Accordion.Item value="item-5">
        <Accordion.Control>
          <Title order={titleOrder}>{t('sidebar.setting')}</Title>
        </Accordion.Control>
        <Accordion.Panel>{t('dashboard.setting_content')}</Accordion.Panel>
      </Accordion.Item>
    </Accordion>
  )
  return content
}
