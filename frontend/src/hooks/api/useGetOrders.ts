import { useEffect, useState } from 'react'
import { useDebouncedValue } from '@mantine/hooks'
import { showNotification } from '@mantine/notifications'
import { IOrder } from '@/entities/order'
import { API, IServerResponseWithPaging, StatusCode } from '@/services'
import useFilter, { IUseFilter } from '../useFilter'

interface IUseGetOrdersProps extends IUseFilter {}
/** useGetOrders
 * @param filterOptions default page,limit,query
 *
 * @returns {object} filter - page,limit,query
 * @returns {boolean} isLoading - is calling api
 * @returns {function} setQuery - Set query value
 * @returns {function} setPage - Set page value
 * @returns {function} setLimit - Set limit value
 * @returns {function} refetch - Refetch api
 * @returns {number} total - Total items
 * @returns {number} totalPage - Total pages
 * @returns {object} orders - product response data
 */
export default function useGetOrders(filterOptions: IUseGetOrdersProps) {
  const { filter, setQuery, setPage, setLimit } = useFilter(filterOptions)
  const [state, setState] = useState<{
    response: IServerResponseWithPaging<{ orders: IOrder[] }> | null
    isLoading: boolean
  }>({
    response: null,
    isLoading: false,
  })
  const [debouncedQuery] = useDebouncedValue(filter.query, 200)

  useEffect(() => {
    handleGetOrders()
  }, [debouncedQuery, filter.page, filter.limit]) // eslint-disable-line react-hooks/exhaustive-deps

  const handleGetOrders = async () => {
    const payload = {
      page: filter.page,
      limit: filter.limit,
      query: filter.query,
    }
    const response = await API.getOrders(payload)
    setState((prev) => ({ ...prev, isLoading: true }))
    if (response.success && response.status === StatusCode.Ok) {
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
    handleGetOrders()
  }

  const total = state.response?.data.total || 0
  const payload = {
    filter: filter,
    isLoading: state.isLoading,
    setQuery: setQuery,
    setPage: setPage,
    setLimit: setLimit,
    refetch: refetch,
    total: total,
    totalPage: Math.ceil((total > 0 ? total : 1) / filter.limit),
    orders: state.response?.data.orders || [],
  }
  // console.log('useGetOrders', payload)
  return payload
}
