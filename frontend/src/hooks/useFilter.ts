import { IPagingQuery } from '@/services'
import { useState } from 'react'

export type IBaseFilter = {
  page?: number
  limit?: number
  query?: string
}
export type IUseFilter<T = {}> = IBaseFilter & T

/** useFilter
 * @returns {object} filter
 * @returns {string} filter.query
 * @returns {number} filter.page
 * @returns {number} filter.limit
 * @returns {function} setQuery - Set query value
 * @returns {function} setPage - Set page value
 * @returns {function} setLimit - Set limit value
 */
export default function useFilter(props: IUseFilter) {
  const { page = 1, limit = 20, query = '' } = props
  const [state, setState] = useState<IPagingQuery>({
    page: page,
    limit: limit,
    query: query,
  })

  const setPage = (value: number) =>
    setState((prev) => ({ ...prev, page: value }))
  const setLimit = (value: number) =>
    setState((prev) => ({ ...prev, limit: value }))
  const setQuery = (value: string) =>
    setState((prev) => ({ ...prev, query: value }))

  return {
    filter: state,
    setPage: setPage,
    setLimit: setLimit,
    setQuery: setQuery,
  }
}
