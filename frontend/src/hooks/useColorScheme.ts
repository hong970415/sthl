import { ColorScheme } from '@mantine/core'
import { useLocalStorage } from '@mantine/hooks'

/** useColorScheme
 * @returns {object} colorScheme
 * @returns {function} toggleColorScheme - Toggle light or dark mode.
 */
export default function useColorScheme() {
  const [colorScheme, setColorScheme] = useLocalStorage<ColorScheme>({
    key: 'color-scheme',
    defaultValue: 'light',
    getInitialValueInEffect: true,
  })

  const toggleColorScheme = (value?: ColorScheme) =>
    setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'))

  return { colorScheme, toggleColorScheme }
}
