import { IUser } from '@/entities/user'
import { httpClient, IServerResponse } from './httpClient'

interface IPassport {
  accessToken: string
  refreshToken: string
}

type ILoginForm = Pick<IUser, 'email' | 'password'>
type IRefreshTokenForm = Pick<IPassport, 'refreshToken'>

async function postLogin(payload: ILoginForm) {
  return httpClient.post<IServerResponse<IPassport>>(
    '/api/v1/users/login',
    payload
  )
}
async function postRefreshToken(payload: IRefreshTokenForm) {
  return httpClient.postWithAuth<IServerResponse<IPassport>>(
    '/api/v1/users/refreshToken',
    payload
  )
}

export type { IPassport, ILoginForm, IRefreshTokenForm }
export { postLogin, postRefreshToken }
