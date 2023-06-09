import { createStyles, rem } from '@mantine/core'

interface SiteProductsStyleProps {}
export default createStyles((theme, {}: SiteProductsStyleProps) => ({
  card: {
    overflow: 'hidden',
    borderRadius: theme.radius.lg,
    height: '100%',
    display: 'flex',
    flexDirection: 'column',
  },
  cardImg: {
    width: '100%',
    objectFit: 'cover',
  },
  cardPrice: {
    color: theme.colorScheme === 'dark' ? theme.white : theme.black,
  },
}))
