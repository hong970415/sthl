import { ReactNode, useEffect } from 'react'
import { useRouter } from 'next/router'
import { Loader } from '@mantine/core'
import { LocalStorageKey } from '@/config/config'
import SidebarLayout from '@/layouts/SidebarLayout'
import useAuth from '@/hooks/useAuth'
import { useAppDispatch } from '@/redux/store'
import { setPassport } from '@/redux/auth/slice'
import { thunkGetMe } from '@/redux/auth/thunk'
import { checkPublicRoute } from '@/utils/route'

export interface IAuthProviderProps {
  children: ReactNode
}

function useAuthProvider() {
  const router = useRouter()
  const isPublicRoute = checkPublicRoute(router)
  const auth = useAuth()
  const dispatch = useAppDispatch()

  // validate when route change, by local storage at&rt
  useEffect(() => {
    const aToken = localStorage.getItem(LocalStorageKey.AccessToken)
    const rToken = localStorage.getItem(LocalStorageKey.RefreshToken)
    if (!isPublicRoute && aToken && rToken) {
      const payload = {
        accessToken: aToken,
        refreshToken: rToken,
      }
      dispatch(setPassport(payload))
      dispatch(thunkGetMe())
    }
  }, [router]) // eslint-disable-line react-hooks/exhaustive-deps
  return {
    isPublicRoute: isPublicRoute,
    isLoggedIn: auth.isLoggedIn,
  }
}
export default function AuthProvider(props: IAuthProviderProps) {
  const { children } = props
  const { isPublicRoute, isLoggedIn } = useAuthProvider()

  return (
    <>
      {isPublicRoute ? (
        children
      ) : isLoggedIn ? (
        <SidebarLayout>{children}</SidebarLayout>
      ) : (
        <div
          style={{
            width: '100%',
            height: '100%',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            fontSize: '50px',
          }}
        >
          <Loader size="xl" />
        </div>
      )}
    </>
  )
}
