import { createStyles, rem } from '@mantine/core'

interface IAlbumImgStyleProps {
  hoverEffect: boolean
}
export default createStyles((theme, { hoverEffect }: IAlbumImgStyleProps) => ({
  imgRoot: {
    position: 'relative',
    overflow: 'hidden',
    borderRadius: theme.radius.md,
    boxShadow: 'rgba(0, 0, 0, 0.24) 0px 3px 8px',
    transition: 'all 0.1s ease-in-out',
    '&:hover': hoverEffect
      ? {
          cursor: 'pointer',
          transform: 'scale(1.05)',
        }
      : null,
  },
  img: {
    width: '100%',
    objectFit: 'cover',
    backgroundColor: '#ffffff',
  },
}))
