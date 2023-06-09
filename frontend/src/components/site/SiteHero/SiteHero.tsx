import { Box, Text } from '@mantine/core'
import useStyles, { TextPosition } from './SiteHero.style'
interface ISiteHeroProps {
  id: string
  scrollMarginTop: string
  text?: string
  textColor?: string
  textPosition?: TextPosition
  imgSrc: string
}

export default function SiteHero(props: ISiteHeroProps) {
  const {
    id,
    scrollMarginTop,
    text,
    textColor = '#ffffff',
    textPosition = TextPosition.Center,
    imgSrc,
  } = props

  const { classes } = useStyles({
    scrollMarginTop: scrollMarginTop,
    textPosition: textPosition,
    textColor: textColor,
  })

  return (
    <div id={id} className={classes.root}>
      <img className={classes.bgImg} src={imgSrc} />
      <Box className={classes.textContainer}>
        <Text className={classes.text}>{text && text}</Text>
      </Box>
    </div>
  )
}
