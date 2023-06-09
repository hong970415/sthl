import { ISensitiveUser, IUser } from '@/entities/user'
import { httpClient, IServerResponse } from './httpClient'

type ISignupForm = Pick<IUser, 'email' | 'password'>
type IUserUpdatePwForm = {
  currentPassword: string
  newPassword: string
}

async function getPing() {
  return httpClient.get<IServerResponse>('/api/v1/ping')
}

async function getMe() {
  return httpClient.getWithAuth<IServerResponse<ISensitiveUser>>(
    '/api/v1/users/me'
  )
}

async function postSignup(payload: ISignupForm) {
  return httpClient.post<IServerResponse<ISensitiveUser>, ISignupForm>(
    `/api/v1/users`,
    payload
  )
}

async function putUpdatePassword(payload: IUserUpdatePwForm) {
  return httpClient.putWithAuth<
    IServerResponse<ISensitiveUser>,
    IUserUpdatePwForm
  >('/api/v1/users/me/pw', payload)
}
async function checkUserExistById(userId: string) {
  return httpClient.get<IServerResponse>(`/api/v1/users/exist/${userId}`)
}

export type { ISignupForm, IUserUpdatePwForm }
export { getPing, getMe, postSignup, putUpdatePassword, checkUserExistById }
