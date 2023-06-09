import { createStyles, rem } from '@mantine/core'
export enum TextPosition {
  TopLeft = 'topLeft',
  Top = 'top',
  TopRight = 'topRight',
  Left = 'left',
  Center = 'center',
  Right = 'right',
  BottomLeft = 'bottomLeft',
  Bottom = 'bottom',
  BottomRight = 'bottomRight',
}
export function getTextStyleByPosition(value: TextPosition) {
  switch (value) {
    case TextPosition.TopLeft:
      return { top: 0, left: 0 }
    case TextPosition.Top:
      return { top: 0, left: '50%', transform: 'translate(-50%, 0)' }
    case TextPosition.TopRight:
      return { top: 0, right: 0 }
    case TextPosition.Left:
      return { top: '50%', left: 0, transform: 'translate(0, -50%)' }
    case TextPosition.Center:
      return { top: '50%', left: '50%', transform: 'translate(-50%, -50%)' }
    case TextPosition.Right:
      return { top: '50%', right: 0, transform: 'translate(0%, -50%)' }
    case TextPosition.BottomLeft:
      return { bottom: 0, left: 0 }
    case TextPosition.Bottom:
      return { bottom: 0, left: '50%', transform: 'translate(-50%, 0)' }
    case TextPosition.BottomRight:
      return { bottom: 0, right: 0 }
    default:
      return { top: '50%', left: '50%', transform: 'translate(-50%, -50%)' }
  }
}
interface SiteHeroStyleProps {
  scrollMarginTop: string
  textPosition: TextPosition
  textColor: string
}
export default createStyles(
  (
    theme,
    { scrollMarginTop, textPosition, textColor }: SiteHeroStyleProps
  ) => ({
    root: {
      scrollMarginTop: scrollMarginTop,
      position: 'relative',
    },
    bgImg: {
      width: '100%',
      objectFit: 'cover',
      [theme.fn.smallerThan('sm')]: {
        height: 500,
      },
    },
    textContainer: {
      position: 'absolute',
      ...getTextStyleByPosition(textPosition),
    },
    text: {
      fontSize: '48px',
      color: textColor,
      fontWeight: 600,
    },
  })
)
