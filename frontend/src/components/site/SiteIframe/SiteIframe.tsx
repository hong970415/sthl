import { ForwardedRef, forwardRef, RefObject, SyntheticEvent } from 'react'
import useStyles from './SiteIframe.style'

interface ISiteIframeProps {
  src: string
  onLoad?: (event: SyntheticEvent<HTMLIFrameElement, Event>) => void
}
// eslint-disable-next-line react/display-name
const SiteIframe = forwardRef(
  (props: ISiteIframeProps, ref: ForwardedRef<HTMLIFrameElement>) => {
    const { src, onLoad } = props
    const { classes } = useStyles()
    const handleOnLoad = (event: SyntheticEvent<HTMLIFrameElement, Event>) => {
      onLoad && onLoad(event)
    }
    return (
      <iframe
        ref={ref}
        className={classes.root}
        src={src}
        onLoad={handleOnLoad}
      />
    )
  }
)
export default SiteIframe
