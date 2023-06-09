import { GetStaticProps } from 'next'
import { useRouter } from 'next/router'
import { serverSideTranslations } from 'next-i18next/serverSideTranslations'
import { Container, Flex } from '@mantine/core'
import HeroTitle from '@/components/HeroTitle/HeroTitle'
import ColorThemeButton from '@/components/ColorThemeButton/ColorThemeButton'
import SelectLocale from '@/components/SelectLocale/SelectLocale'
import useStyles from '@/styles/index.style'

/** getStaticProps (NextJS)
 * @param locale the type of locale
 * @returns
 */
export const getStaticProps: GetStaticProps = async (context) => {
  const { locale } = context!
  return {
    props: {
      ...(await serverSideTranslations(locale as any, ['data'])),
      // will be passed to the page component as props
    },
  }
}

export default function IndexPage() {
  const { classes } = useStyles()
  const router = useRouter()
  console.log('v1.1')
  const linkGetStarted = `/${router.locale}/cms/login`
  const linkGithub = 'https://github.com/hong970415/sthl'
  return (
    <Container className={classes.container}>
      <Flex justify="flex-end" pt={'lg'} gap="md">
        <SelectLocale />
        <ColorThemeButton />
      </Flex>
      <HeroTitle linkGetStarted={linkGetStarted} linkGithub={linkGithub} />
    </Container>
  )
}
