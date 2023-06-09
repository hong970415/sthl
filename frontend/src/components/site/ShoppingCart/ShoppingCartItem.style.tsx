import { createStyles, rem } from '@mantine/core'

interface SiteShoppingCartItemStyleProps {}
export default createStyles((theme, {}: SiteShoppingCartItemStyleProps) => ({
  root: {
    borderRadius: theme.radius.md,
    padding: theme.spacing.sm,
    ':hover': {
      backgroundColor: theme.colors.gray[3],
    },
  },
  img: {
    width: '100%',
    minHeight: '50px',
  },
}))
