import { GetStaticProps } from 'next'
import { useState } from 'react'
import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import { faPenToSquare } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { ActionIcon, Flex, Grid, Text } from '@mantine/core'
import useTranslationData from '@/hooks/useTranslation'
import ColorThemeButton from '@/components/ColorThemeButton/ColorThemeButton'
import UserUpdatePwForm from '@/components/forms/UserUpdatePwForm'
import PageContentContainer from '@/components/PageContentContainer/PageContentContainer'
import SelectLocale from '@/components/SelectLocale/SelectLocale'

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

export default function SettingPage() {
  const { t } = useTranslationData()
  const [state, setState] = useState({
    showUserUpdatePwForm: false,
  })
  const handleToggleShowUserUpdatePwForm = (
    event: React.MouseEvent<HTMLButtonElement>
  ) => {
    // toggle off will clear form data
    setState((prev) => ({
      ...prev,
      showUserUpdatePwForm: !prev.showUserUpdatePwForm,
    }))
  }
  const changePwEl = (
    <Flex direction={'column'} sx={{ width: '100%' }}>
      <ActionIcon
        variant="outline"
        onClick={handleToggleShowUserUpdatePwForm}
        mb={state.showUserUpdatePwForm ? 'lg' : undefined}
        color={state.showUserUpdatePwForm ? 'blue' : 'gray'}
      >
        <FontAwesomeIcon icon={faPenToSquare} />
      </ActionIcon>
      {state.showUserUpdatePwForm && <UserUpdatePwForm />}
    </Flex>
  )
  const sections = [
    {
      key: 'general.language',
      rightEl: <SelectLocale />,
    },
    {
      key: 'general.dark_mode',
      rightEl: <ColorThemeButton />,
    },
    {
      key: 'general.change_pw',
      rightEl: changePwEl,
    },
  ]
  return (
    <PageContentContainer>
      {sections &&
        sections.map((section, index) => {
          const leftSpan = 3
          const rightSpan = 9
          return (
            <Grid key={section.key} mt={index === 0 ? undefined : 'lg'}>
              <Grid.Col span={leftSpan}>
                <Text mr="lg">{t(section.key)}:</Text>
              </Grid.Col>
              <Grid.Col span={rightSpan}>{section.rightEl}</Grid.Col>
            </Grid>
          )
        })}
    </PageContentContainer>
  )
}
