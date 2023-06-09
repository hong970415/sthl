import { useEffect, useState } from 'react'
import { showNotification } from '@mantine/notifications'
import { isString } from 'lodash'
import { API, IServerResponse, StatusCode } from '@/services'
import { IProduct } from '@/entities/product'

/** useGetProductById
 * @param userId
 * @param productId
 * @returns {object} product - api response data
 * @returns {boolean} isLoading - is calling api
 * @returns {function} refetch - Refetch api
 */
export default function useGetProductById(userId?: string, productId?: string) {
  const [state, setState] = useState<{
    response: IServerResponse<IProduct> | undefined
    isLoading: boolean
  }>({
    response: undefined,
    isLoading: true,
  })

  useEffect(() => {
    isString(userId) &&
      isString(productId) &&
      handleGetProductById(userId, productId)
  }, [userId, productId]) // eslint-disable-line react-hooks/exhaustive-deps

  const handleGetProductById = async (uId: string, pId: string) => {
    setState((prev) => ({ ...prev, isLoading: true }))
    const response = await API.getProductById(uId, pId)
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
    isString(userId) &&
      isString(productId) &&
      handleGetProductById(userId, productId)
  }

  const payload = {
    product: state.response?.data,
    isLoading: state.isLoading,
    refetch: refetch,
  }
  // console.log('useGetProductById', payload)
  return payload
}
