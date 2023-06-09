import { MouseEvent, ReactNode } from 'react'
import { Group, ThemeIcon } from '@mantine/core'
import { useRouter } from 'next/router'
import useStyles from './NavbarLink.style'

export interface INavbarLinkProps {
  route: string
  label?: string
  icon?: ReactNode
  hideText?: boolean
  onClick?: (event: MouseEvent<HTMLAnchorElement>) => void
}
export default function NavbarLink(props: INavbarLinkProps) {
  const { route, label = '', icon, hideText = false, onClick } = props
  const router = useRouter()
  const { classes } = useStyles()

  const handleOnClick = (event: MouseEvent<HTMLAnchorElement>) => {
    event.preventDefault()
    if (onClick) {
      onClick(event)
      return
    }
    router.push(route)
  }

  return (
    <a href={route} className={classes.link} onClick={handleOnClick}>
      <Group position={'left'}>
        <ThemeIcon variant="light" size={'md'}>
          {icon}
        </ThemeIcon>
        <span>{label}</span>
      </Group>
    </a>
  )
  // return (
  //   <Tooltip
  //     label={label}
  //     position="right"
  //     zIndex={100}
  //     withinPortal={true}
  //     disabled={hideText}
  //   >
  //     <a href={route} className={classes.link} onClick={handleOnClick}>
  //       <Group position={'left'}>
  //         <ThemeIcon variant="light" size={'md'}>
  //           {icon}
  //         </ThemeIcon>
  //         <span>{label}</span>
  //       </Group>
  //     </a>
  //   </Tooltip>
  // )
}
