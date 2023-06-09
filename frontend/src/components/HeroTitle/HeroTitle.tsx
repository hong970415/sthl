import { Container, Text, Button, Group } from '@mantine/core'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faGithub } from '@fortawesome/free-brands-svg-icons'
import useTranslationData from '@/hooks/useTranslation'
import useStyles from './HeroTitle.style'

interface IHeroTitleProps {
  linkGetStarted: string
  linkGithub: string
}
export default function HeroTitle(props: IHeroTitleProps) {
  const { classes } = useStyles()
  const { t } = useTranslationData()
  const { linkGetStarted, linkGithub } = props

  return (
    <div className={classes.wrapper}>
      <Container size={700} className={classes.inner}>
        <h1 className={classes.title}>
          {t('hero_title')}{' '}
          <Text
            component="span"
            variant="gradient"
            gradient={{ from: 'blue', to: 'cyan' }}
            inherit
          >
            {t('hero_title_forfun')}
          </Text>
        </h1>

        <Text className={classes.description} color="dimmed">
          {t('hero_description')}
        </Text>

        <Group className={classes.controls}>
          <Button
            size="xl"
            className={classes.control}
            variant="gradient"
            gradient={{ from: 'blue', to: 'cyan' }}
            component="a"
            href={linkGetStarted}
          >
            {t('get_started')}
          </Button>

          <Button
            component="a"
            href={linkGithub}
            size="xl"
            variant="default"
            className={classes.control}
            leftIcon={<FontAwesomeIcon icon={faGithub} />}
          >
            {t('github')}
          </Button>
        </Group>
      </Container>
    </div>
  )
}
