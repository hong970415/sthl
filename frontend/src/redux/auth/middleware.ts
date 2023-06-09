import Router from 'next/router'
import { createListenerMiddleware } from '@reduxjs/toolkit'
import { showNotification } from '@mantine/notifications'
import { LocalStorageKey } from '@/config/config'
import { thunkGetMe, thunkLogin } from './thunk'
import { logout } from './slice'

export const authListenerMiddleware = createListenerMiddleware()

authListenerMiddleware.startListening({
  actionCreator: logout,
  effect: (action, listenerApi) => {
    showNotification({ color: 'green', message: 'ok' })
    localStorage.setItem(LocalStorageKey.AccessToken, '')
    localStorage.setItem(LocalStorageKey.RefreshToken, '')
    Router.push('/cms/login')
  },
})

// thunkLogin
authListenerMiddleware.startListening({
  actionCreator: thunkLogin.fulfilled,
  effect: (action, listenerApi) => {
    showNotification({ color: 'green', message: action.payload.msg })
    localStorage.setItem(
      LocalStorageKey.AccessToken,
      action.payload.data.accessToken
    )
    localStorage.setItem(
      LocalStorageKey.RefreshToken,
      action.payload.data.refreshToken
    )
    Router.push('/cms/dashboard')
  },
})
authListenerMiddleware.startListening({
  actionCreator: thunkLogin.rejected,
  effect: (action, listenerApi) => {
    showNotification({ color: 'red', message: action.payload as string })
    localStorage.setItem(LocalStorageKey.AccessToken, '')
    localStorage.setItem(LocalStorageKey.RefreshToken, '')
    Router.push('/cms/login')
  },
})

// thunkGetMe
authListenerMiddleware.startListening({
  actionCreator: thunkGetMe.fulfilled,
  effect: (action, listenerApi) => {
    // showNotification({ color: 'green', message: action.payload.msg })
  },
})
authListenerMiddleware.startListening({
  actionCreator: thunkGetMe.rejected,
  effect: (action, listenerApi) => {
    showNotification({ color: 'red', message: action.payload as string })
    Router.push('/cms/login')
  },
})
