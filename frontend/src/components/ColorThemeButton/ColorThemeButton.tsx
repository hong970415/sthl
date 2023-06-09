import { useMemo } from 'react'
import { ActionIcon, useMantineTheme } from '@mantine/core'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faMoon, faSun } from '@fortawesome/free-solid-svg-icons'
import useColorScheme from '@/hooks/useColorScheme'

export default function ColorThemeButton() {
  const theme = useMantineTheme()
  const { isDarkTheme } = theme.other
  const { toggleColorScheme } = useColorScheme()

  // memorizeded elements
  const [icon, color] = useMemo(
    () =>
      isDarkTheme
        ? [<FontAwesomeIcon icon={faSun} />, 'yellow'] // eslint-disable-line react/jsx-key
        : [<FontAwesomeIcon icon={faMoon} />, 'blue'], // eslint-disable-line react/jsx-key
    [isDarkTheme]
  )

  return (
    <ActionIcon
      title="Toggle color scheme"
      variant="outline"
      color={color}
      onClick={() => toggleColorScheme()}
      size={'md'}
    >
      {icon}
    </ActionIcon>
  )
}
