import axios, {
  AxiosInstance,
  AxiosRequestConfig,
  AxiosResponse,
  AxiosError,
} from 'axios'
import { get } from 'lodash'
import { getClientConfig } from '@/config/config'
import { store } from '@/redux/store'

enum StatusCode {
  Ok = 200,
  Created = 201,
  BadRequest = 400,
  Unauthorized = 401,
  Forbidden = 403,
  NotFound = 404,
  TooManyRequests = 429,
  InternalServerError = 500,
}
interface IServerResponse<T = null> {
  code: number
  msg: string
  data: T
}
interface IPagingQuery {
  page: number
  limit: number
  query?: string
}
interface IPagingQueryResult {
  page: number
  limit: number
  total: number
}
interface IServerResponseWithPaging<T = null> {
  code: number
  msg: string
  data: T & IPagingQueryResult
}
interface CustomAxiosResponse<T = any, D = any> extends AxiosResponse<T, D> {
  success: boolean
  errorMsg?: string
}

class HttpClient {
  private api: AxiosInstance

  public constructor(config: AxiosRequestConfig, lang?: string) {
    this.api = axios.create(config)
    this.initRequestInterceptor()
    this.initReponseInterceptor()
  }

  // this middleware is been called right before the http request is made.
  private initRequestInterceptor(): void {
    // this.api.interceptors.request.use((param: AxiosRequestConfig) => param)
    this.api.interceptors.request.use((param: any) => param)
  }

  // this middleware is been called right before the response is get it by the method that triggers the request
  private initReponseInterceptor(): void {
    const noramlInterceptor = (param: AxiosResponse): CustomAxiosResponse => {
      const _param = { ...param, success: true }
      // console.log(
      //   'Response Interceptors:',
      //   _param.config.method?.toUpperCase(),
      //   _param.config.url,
      //   _param.status,
      //   _param
      // )
      return _param
    }

    const errorInterceptor = (error: AxiosError) => {
      // console.log(
      //   'Response Interceptors Error:',
      //   error.config?.method?.toUpperCase(),
      //   error.config?.url,
      //   error?.response?.status,
      //   error
      // )
      if (error.response) {
        // The request was made and the server responded with a status code
        // that falls out of the range of 2xx
        // if (error.response.status === StatusCode.Unauthorized) {
        //   const { locale } = router.router?.query!
        //   if (isString(locale) && validateLocale(locale)) {
        //     const path = `/${locale}/login`
        //     router.push(path)
        //   } else {
        //     const path = `/${LocaleKey.EN}/login`
        //     router.push(path)
        //   }
        // }
        return {
          ...error.response,
          success: false,
          errorMsg: get(error, 'response.data.msg', ''),
        }
      } else if (error.request) {
        // The request was made but no response was received
        // `error.request` is an instance of XMLHttpRequest in the browser and an instance of
        // http.ClientRequest in node.js
        return { success: false, errorMsg: error.message }
      } else {
        // Something happened in setting up the request that triggered an Error
        return { success: false, errorMsg: 'Internal Server Error' }
      }
    }

    this.api.interceptors.response.use(noramlInterceptor, errorInterceptor)
  }

  private getStorePassport() {
    const authHeaders = {
      Authorization: `bearer ${store.getState().auth.passport.accessToken}`,
    }
    return authHeaders
  }

  public getUri(config?: AxiosRequestConfig): string {
    return this.api.getUri(config)
  }
  // T=response data
  // D=resquest data
  public request<T = any, D = any>(
    config: AxiosRequestConfig<D>
  ): Promise<CustomAxiosResponse<T>> {
    return this.api.request<T, CustomAxiosResponse<T>, D>(config)
  }

  public head<T = any, D = any>(
    url: string,
    config?: AxiosRequestConfig<D>
  ): Promise<CustomAxiosResponse<T, D>> {
    return this.api.head<T, CustomAxiosResponse<T>, D>(url, config)
  }

  public get<T = any, D = any>(
    url: string,
    config?: AxiosRequestConfig<D>
  ): Promise<CustomAxiosResponse<T, D>> {
    return this.api.get<T, CustomAxiosResponse<T>, D>(url, config)
  }
  public getWithAuth<T = any, D = any>(
    url: string,
    config?: AxiosRequestConfig<D>
  ): Promise<CustomAxiosResponse<T, D>> {
    return this.api.get<T, CustomAxiosResponse<T>, D>(url, {
      ...config,
      headers: this.getStorePassport(),
    })
  }

  public post<T = any, D = any>(
    url: string,
    data?: D,
    config?: AxiosRequestConfig<D>
  ): Promise<CustomAxiosResponse<T, D>> {
    return this.api.post<T, CustomAxiosResponse<T>, D>(url, data, config)
  }
  public postWithAuth<T = any, D = any>(
    url: string,
    data?: D,
    config?: AxiosRequestConfig<D>
  ): Promise<CustomAxiosResponse<T, D>> {
    return this.api.post<T, CustomAxiosResponse<T>, D>(url, data, {
      ...config,
      headers: { ...config?.headers, ...this.getStorePassport() },
    })
  }

  public putWithAuth<T = any, D = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig<D>
  ): Promise<CustomAxiosResponse<T, D>> {
    return this.api.put<T, CustomAxiosResponse<T>, D>(url, data, {
      ...config,
      headers: { ...config?.headers, ...this.getStorePassport() },
    })
  }

  public patchWithAuth<T = any, D = any>(
    url: string,
    data?: any,
    config?: AxiosRequestConfig<D>
  ): Promise<CustomAxiosResponse<T, D>> {
    return this.api.patch<T, CustomAxiosResponse<T>, D>(url, data, {
      ...config,
      headers: { ...config?.headers, ...this.getStorePassport() },
    })
  }

  public deleteWithAuth<T = any, D = any>(
    url: string,
    config?: AxiosRequestConfig<D>
  ): Promise<CustomAxiosResponse<T, D>> {
    return this.api.delete<T, CustomAxiosResponse<T>, D>(url, {
      ...config,
      headers: { ...config?.headers, ...this.getStorePassport() },
    })
  }
}
const httpClient = new HttpClient(getClientConfig())

export type {
  IServerResponse,
  IPagingQuery,
  IPagingQueryResult,
  IServerResponseWithPaging,
  CustomAxiosResponse,
}
export { StatusCode, httpClient }
