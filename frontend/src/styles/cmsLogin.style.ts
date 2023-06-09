import { createStyles } from '@mantine/core'

export default createStyles((theme) => ({
  container: {
    width: '100%',
    height: '100%',
  },
  cardContainer: {
    width: '500px',
    borderRadius: theme.radius.md,
    [`@media (max-width: ${theme.breakpoints.xs}px)`]: {
      width: '340px',
    },
  },
  actionText: {
    ':hover': {
      textDecoration: 'underline',
      cursor: 'pointer',
    },
  },
}))
