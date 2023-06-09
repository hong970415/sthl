import { useEffect, useState } from 'react'
import { showNotification } from '@mantine/notifications'
import { isString } from 'lodash'
import { API, IServerResponse, StatusCode } from '@/services'
import { IOrder } from '@/entities/order'

/** useGetOrderById
 * @param orderId
 * @returns {object} order - api response data
 * @returns {boolean} isLoading - is calling api
 * @returns {function} refetch - Refetch api
 */
export default function useGetOrderById(orderId?: string) {
  const [state, setState] = useState<{
    response: IServerResponse<IOrder> | undefined
    isLoading: boolean
  }>({
    response: undefined,
    isLoading: true,
  })

  useEffect(() => {
    isString(orderId) && handleGetOrderById(orderId)
  }, [orderId]) // eslint-disable-line react-hooks/exhaustive-deps

  const handleGetOrderById = async (value: string) => {
    setState((prev) => ({ ...prev, isLoading: true }))
    const response = await API.getOrderById(value)
    if (
      response.success &&
      response.status === StatusCode.Ok &&
      response.data.data
    ) {
      setState((prev) => ({
        ...prev,
        response: response.data,
        isLoading: false,
      }))
    } else {
      // const message = response.errorMsg
      // showNotification({ color: 'red', message: message })
    }
  }

  const refetch = () => {
    isString(orderId) && handleGetOrderById(orderId)
  }

  const payload = {
    order: state.response?.data,
    isLoading: state.isLoading,
    refetch: refetch,
  }
  // console.log('useGetOrderById', payload)
  return payload
}
