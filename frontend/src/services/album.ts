import { IImgInfo } from '@/entities/imageinfo'
import { makeQueryString } from '@/utils/param'
import {
  httpClient,
  IPagingQuery,
  IServerResponse,
  IServerResponseWithPaging,
} from './httpClient'

async function postUploadAlbumImg(value: FormData) {
  const cfg = {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  }
  return httpClient.postWithAuth<IServerResponse<IImgInfo>>(
    '/api/v1/album',
    value,
    cfg
  )
}

async function getAlbumImgsByUserId(payload: IPagingQuery) {
  const params = makeQueryString(payload)
  return httpClient.getWithAuth<
    IServerResponseWithPaging<{ imgs: IImgInfo[] }>
  >(`/api/v1/album${params}`)
}

async function putUpdateImgDataById(id: number, value: FormData) {
  const cfg = {
    headers: {
      'Content-Type': 'multipart/form-data',
    },
  }
  return httpClient.putWithAuth<IServerResponse<IImgInfo>>(
    `/api/v1/album/${id}`,
    value,
    cfg
  )
}

export { postUploadAlbumImg, getAlbumImgsByUserId, putUpdateImgDataById }
