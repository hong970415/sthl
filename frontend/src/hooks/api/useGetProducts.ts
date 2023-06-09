import { useEffect, useState } from 'react'
import { showNotification } from '@mantine/notifications'
import { useDebouncedValue } from '@mantine/hooks'
import { API, IServerResponseWithPaging, StatusCode } from '@/services'
import useFilter, { IUseFilter } from '../useFilter'
import useAuth from '../useAuth'
import { IProduct } from '@/entities/product'

interface IUseGetProductsProps extends IUseFilter {}
/** useGetProducts
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
 * @returns {object} products - product response data
 */
export default function useGetProducts(filterOptions: IUseGetProductsProps) {
  const { filter, setQuery, setPage, setLimit } = useFilter(filterOptions)
  const [state, setState] = useState<{
    response: IServerResponseWithPaging<{ products: IProduct[] }> | null
    isLoading: boolean
  }>({
    response: null,
    isLoading: false,
  })
  const { user } = useAuth()
  const [debouncedQuery] = useDebouncedValue(filter.query, 200)

  useEffect(() => {
    handleGetProducts()
  }, [debouncedQuery, filter.page, filter.limit]) // eslint-disable-line react-hooks/exhaustive-deps

  const handleGetProducts = async () => {
    const userId = user ? user.id : ''
    const payload = {
      page: filter.page,
      limit: filter.limit,
      query: filter.query,
    }
    const response = await API.getProducts(userId, payload)
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
    handleGetProducts()
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
    products: state.response?.data.products || [],
  }
  // console.log('useGetProducts', payload)
  return payload
}
