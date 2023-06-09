import {
  httpClient,
  IPagingQuery,
  IServerResponse,
  IServerResponseWithPaging,
} from './httpClient'
import { makeQueryString } from '@/utils/param'
import { IProduct } from '@/entities/product'

type ICreateProductForm = Pick<
  IProduct,
  'name' | 'price' | 'quantity' | 'description' | 'imgUrl'
>
type IUpdateProductForm = ICreateProductForm

async function postCreateProduct(payload: ICreateProductForm) {
  return httpClient.postWithAuth<IServerResponse<IProduct>, ICreateProductForm>(
    `/api/v1/products`,
    payload
  )
}

async function getProducts(userId: string, payload: IPagingQuery) {
  const params = makeQueryString(payload)
  return httpClient.get<IServerResponseWithPaging<{ products: IProduct[] }>>(
    `/api/v1/products/${userId}${params}`
  )
}

async function getProductById(userId: string, productId: string) {
  return httpClient.get<IServerResponse<IProduct>, string>(
    `/api/v1/products/${userId}/${productId}`
  )
}

async function putUpdateProductById(
  userId: string,
  productId: string,
  payload: IUpdateProductForm
) {
  const response = await httpClient.putWithAuth<
    IServerResponse,
    IUpdateProductForm
  >(`/api/v1/products/${userId}/${productId}`, payload)
  return response
}

async function deleteProductById(userId: string, productId: string) {
  const response = await httpClient.deleteWithAuth<IServerResponse, string>(
    `/api/v1/products/${userId}/${productId}`
  )
  return response
}

export type { ICreateProductForm, IUpdateProductForm }
export {
  postCreateProduct,
  getProducts,
  getProductById,
  putUpdateProductById,
  deleteProductById,
}
