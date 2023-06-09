import { SyntheticEvent, useState } from 'react'
import { Box, LoadingOverlay } from '@mantine/core'
import { IImgInfo } from '@/entities/imageinfo'
import useStyles from './AlbumImg.style'

interface IAlbumImgProps {
  // data: IImgInfo
  imgSrc: string
  hoverEffect?: boolean
}
export default function AlbumImg(props: IAlbumImgProps) {
  const { imgSrc, hoverEffect = true } = props
  const { classes } = useStyles({ hoverEffect })
  const [state, setState] = useState({ isLoading: true })

  const handleOnLoad = (event: SyntheticEvent<HTMLImageElement, Event>) => {
    setState((prev) => ({ ...prev, isLoading: false }))
  }
  return (
    <Box component="div" className={classes.imgRoot}>
      <LoadingOverlay
        visible={state.isLoading}
        overlayBlur={1000}
        transitionDuration={400}
      />
      <img className={classes.img} src={imgSrc} onLoad={handleOnLoad} />
    </Box>
  )
}
