import { createStyles, rem } from '@mantine/core'

interface SiteShoppingCartStyleProps {}
export default createStyles((theme, {}: SiteShoppingCartStyleProps) => ({
  itemsContainer: {
    width: '100%',
    maxHeight: '40vh',
    overflowY: 'auto',
  },
}))
