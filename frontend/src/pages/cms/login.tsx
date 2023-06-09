import { useRouter } from 'next/router'
import { useReducer } from 'react'
import { GetStaticProps } from 'next'
import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import { Card, Center, Group, Text, Title } from '@mantine/core'
import useTranslationData from '@/hooks/useTranslation'
import LoginForm from '@/components/forms/LoginForm'
import SignupForm from '@/components/forms/SignupForm'
import SelectLocale from '@/components/SelectLocale/SelectLocale'
import ColorThemeButton from '@/components/ColorThemeButton/ColorThemeButton'
import useStyles from '@/styles/cmsLogin.style'

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

// login page useReducer
export interface ICmsLoginPageState {
  move: 'login' | 'signup'
}
const initialState: ICmsLoginPageState = {
  move: 'login',
}
export type ICmsLoginPageStateAction = { type: 'toggleMove' }
export const cmsLoginPageStateActionReducer = (
  state: ICmsLoginPageState,
  action: ICmsLoginPageStateAction
): ICmsLoginPageState => {
  switch (action.type) {
    case 'toggleMove':
      return {
        ...state,
        move: state.move === 'login' ? 'signup' : 'login',
      }
    default: {
      throw new Error(`Unhandled action type: ${JSON.stringify(action)}`)
    }
  }
}

export default function CmsLoginPage(props: any) {
  const router = useRouter()
  const { classes } = useStyles()
  const { t } = useTranslationData()
  const [state, stateDispatch] = useReducer(
    cmsLoginPageStateActionReducer,
    initialState
  )
  const toggleMove = () => stateDispatch({ type: 'toggleMove' })
  const handleSignupOnSuccess = () => toggleMove()

  // determine render login/signup compoment
  const isLoginMove = state.move === 'login'
  const renderByMove = isLoginMove ? (
    <>
      <LoginForm />
      <Center>
        <Group>
          <Text size="md">{t('general.no_account')}</Text>
          <Text
            size="md"
            c="blue"
            onClick={toggleMove}
            className={classes.actionText}
          >
            {t('general.signup')}
          </Text>
        </Group>
      </Center>
    </>
  ) : (
    <>
      <SignupForm onSuccess={handleSignupOnSuccess} />
      <Center>
        <Group>
          <Text size="md">{t('general.have_account')}</Text>
          <Text
            size="md"
            c="blue"
            onClick={toggleMove}
            className={classes.actionText}
          >
            {t('general.login')}
          </Text>
        </Group>
      </Center>
    </>
  )

  return (
    <Center className={classes.container}>
      <Card className={classes.cardContainer} withBorder>
        <Group mb="md" position="apart">
          <SelectLocale />
          <ColorThemeButton />
        </Group>

        <Title mb="md" order={2} align="center">
          {t('general.welcome')}
        </Title>

        {renderByMove}
      </Card>
    </Center>
  )
}
