import '@/styles/globals.css'
import '@/styles/nprogress.css'
import '@fortawesome/fontawesome-svg-core/styles.css'
import { config } from '@fortawesome/fontawesome-svg-core'
import type { AppProps } from 'next/app'
import Head from 'next/head'
import { appWithTranslation, SSRConfig } from 'next-i18next'
import { Provider } from 'react-redux'
import { ColorSchemeProvider, MantineProvider } from '@mantine/core'
import { Notifications } from '@mantine/notifications'
import { wrapper } from '@/redux/store'
import AuthProvider from '@/providers/AuthProvider'
import useColorScheme from '@/hooks/useColorScheme'
import useRouterTransition from '@/hooks/useRouterTransition'

// handle fontawesome icons oversize when page loaded
config.autoAddCss = false

function App({ Component, ...rest }: AppProps) {
  // redux
  const { store, props } = wrapper.useWrappedStore(rest)
  const { pageProps } = props
  const { colorScheme, toggleColorScheme } = useColorScheme()
  useRouterTransition()
  const theme = {
    /** Put your mantine theme override here */
    colorScheme: colorScheme,
    other: {
      isDarkTheme: colorScheme === 'dark',
    },
  }

  return (
    <>
      <Head>
        <title>sthl</title>
        <meta name="description" content="create your website" />
        <meta name="viewport" content="width=device-width, initial-scale=1" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Provider store={store}>
        <ColorSchemeProvider
          colorScheme={colorScheme}
          toggleColorScheme={toggleColorScheme}
        >
          <MantineProvider
            //
            withGlobalStyles
            withNormalizeCSS
            theme={theme}
          >
            <Notifications position="top-right" />
            <AuthProvider>
              <Component {...pageProps} />
            </AuthProvider>
          </MantineProvider>
        </ColorSchemeProvider>
      </Provider>
    </>
  )
}

export default appWithTranslation(App)
