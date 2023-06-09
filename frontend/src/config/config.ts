import { AxiosRequestConfig } from 'axios'

export function getApiUrl() {
  return process.env.NEXT_PUBLIC_API_URL
}

export function getClientConfig(): AxiosRequestConfig {
  return {
    baseURL: getApiUrl(),
    withCredentials: true,
    headers: {
      Accept: 'application/json',
      'Access-Control-Allow-Origin': '*',
      'Access-Control-Allow-Headers': '*',
      'Access-Control-Allow-Methods': 'DELETE,GET,HEAD,OPTIONS,PATCH,POST,PUT',
      'Access-Control-Allow-Credentials': true,
      'Content-Type': 'application/json; charset=utf-8',
    },
    timeout: 1000 * 30,
  }
}

export enum LocalStorageKey {
  AccessToken = 'accessToken',
  RefreshToken = 'refreshToken',
}
