import { IOrder, IOrderItem } from '@/entities/order'
import { makeQueryString } from '@/utils/param'
import {
  httpClient,
  IPagingQuery,
  IServerResponse,
  IServerResponseWithPaging,
} from './httpClient'

type ICreateOrderItem = Pick<
  IOrderItem,
  'productId' | 'purchasedName' | 'purchasedPrice' | 'quantity'
>
type ICreateOrderForm = Pick<
  IOrder,
  'remark' | 'discount' | 'totalAmount' | 'paymentMethod' | 'shippingAddress'
> & {
  items: Array<ICreateOrderItem>
}

type IUpdateOrderForm = Pick<
  IOrder,
  | 'remark'
  | 'discount'
  | 'totalAmount'
  | 'paymentMethod'
  | 'shippingAddress'
  | 'status'
  | 'paymentStatus'
  | 'paymentMethod'
  | 'deliveryStatus'
>

async function postCreateOrder(userId: string, payload: ICreateOrderForm) {
  return httpClient.post<IServerResponse<IOrder>, ICreateOrderForm>(
    `/api/v1/orders/${userId}`,
    payload
  )
  // return httpClient.postWithAuth<IServerResponse<IOrder>, ICreateOrderForm>(
  //   `/api/v1/orders`,
  //   payload
  // )
}

async function getOrders(payload: IPagingQuery) {
  const params = makeQueryString(payload)
  return httpClient.getWithAuth<
    IServerResponseWithPaging<{ orders: IOrder[] }>
  >(`/api/v1/orders${params}`)
}

async function getOrderById(value: string) {
  return httpClient.getWithAuth<IServerResponse<IOrder>, string>(
    `/api/v1/orders/${value}`
  )
}

async function putUpdateOrderById(id: string, payload: IUpdateOrderForm) {
  const response = await httpClient.putWithAuth<
    IServerResponse<IOrder>,
    IUpdateOrderForm
  >(`/api/v1/orders/${id}`, payload)
  return response
}

async function deleteOrderById(value: string) {
  return httpClient.deleteWithAuth<IServerResponse, string>(
    `/api/v1/orders/${value}`
  )
}

export type { ICreateOrderItem, ICreateOrderForm, IUpdateOrderForm }
export {
  postCreateOrder,
  getOrders,
  getOrderById,
  putUpdateOrderById,
  deleteOrderById,
}
