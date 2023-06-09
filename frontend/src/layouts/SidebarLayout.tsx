import { MouseEvent, ReactNode, useMemo } from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { ActionIcon, AppShell, Box, Group, Header, Navbar } from '@mantine/core'
import {
  faChevronLeft,
  faChevronRight,
} from '@fortawesome/free-solid-svg-icons'
import { getMainRoutes, getSettingRoutes } from '@/config/route'
import { useAppDispatch } from '@/redux/store'
import useTranslationData from '@/hooks/useTranslation'
import useSidebar from '@/hooks/useSidebar'
import ColorThemeButton from '@/components/ColorThemeButton/ColorThemeButton'
import NavbarLink from '@/components/NavbarLink/NavbarLink'
import useStyles from './SidebarLayout.style'
import { logout } from '@/redux/auth/slice'

export interface ISidebarLayoutProps {
  children: ReactNode
}
export default function SidebarLayout(props: ISidebarLayoutProps) {
  const { children } = props
  const { classes } = useStyles()
  const { t } = useTranslationData()
  const [navbarMainRoutes, navbarFooterRoutes] = useMemo(
    () => [getMainRoutes(), getSettingRoutes()],
    []
  )
  const sb = useSidebar()
  const dispatch = useAppDispatch()
  const handleLogout = (event: MouseEvent<HTMLAnchorElement>) => {
    dispatch(logout())
  }

  const headerEl = (
    <Header height={50} p="md">
      <Box
        sx={{
          display: 'flex',
          alignItems: 'center',
          height: '100%',
          width: '100%',
        }}
      >
        <ActionIcon size="lg" onClick={sb.toggleSidebar}>
          {true ? (
            <FontAwesomeIcon icon={faChevronLeft} />
          ) : (
            <FontAwesomeIcon icon={faChevronRight} />
          )}
        </ActionIcon>
        <Group position="right" sx={{ width: '100%' }}>
          <ColorThemeButton />
        </Group>
      </Box>
    </Header>
  )

  const navbarEl = (
    <Navbar
      height={'auto'}
      width={{ sm: sb.width }}
      hiddenBreakpoint="sm"
      style={{ visibility: sb.isCollapsed ? 'hidden' : 'visible' }}
      p="md"
      hidden={sb.isCollapsed}
    >
      <Navbar.Section grow>
        {navbarMainRoutes.map((item) => {
          return (
            <NavbarLink
              key={item.route}
              route={item.route}
              label={t<string>(item.labelKey)}
              icon={<FontAwesomeIcon icon={item.icon} />}
            />
          )
        })}
      </Navbar.Section>

      <Navbar.Section className={classes.footer}>
        {navbarFooterRoutes.map((item) => {
          if (item.group === 'logout') {
            return (
              <NavbarLink
                key={item.route}
                route={item.route}
                label={t<string>(item.labelKey)}
                icon={<FontAwesomeIcon icon={item.icon} />}
                onClick={handleLogout}
              />
            )
          } else {
            return (
              <NavbarLink
                key={item.route}
                route={item.route}
                label={t<string>(item.labelKey)}
                icon={<FontAwesomeIcon icon={item.icon} />}
              />
            )
          }
        })}
      </Navbar.Section>
    </Navbar>
  )

  return (
    <AppShell
      styles={(theme) => ({
        main: {
          backgroundColor:
            theme.colorScheme === 'dark'
              ? theme.colors.dark[8]
              : theme.colors.gray[0],
        },
      })}
      header={headerEl}
      navbar={navbarEl}
    >
      {children}
    </AppShell>
  )
}
