import { useEffect, useState } from 'react'
import { showNotification } from '@mantine/notifications'
import { isString } from 'lodash'
import { API, StatusCode, IServerResponseWithPaging } from '@/services'
import { IOrder } from '@/entities/order'
import { IImgInfo } from '@/entities/imageinfo'
import useFilter, { IUseFilter } from '../useFilter'

interface IUseGetAlbumImgsByUserIdProps extends IUseFilter {}
/** useGetAlbumImgsByUserId
 * @returns {object} gallery - api response data
 * @returns {boolean} isLoading - is calling api
 * @returns {function} refetch - Refetch api
 */
export default function useGetAlbumImgsByUserId(
  filterOptions: IUseGetAlbumImgsByUserIdProps
) {
  const { filter, setQuery, setPage, setLimit } = useFilter(filterOptions)
  const [state, setState] = useState<{
    response: IServerResponseWithPaging<{ imgs: IImgInfo[] }> | undefined
    isLoading: boolean
    isFetched: boolean
    count: number //for reload img src
  }>({
    response: undefined,
    isLoading: true,
    isFetched: false,
    count: 1,
  })

  useEffect(() => {
    handleGetGalleryByUserId()
  }, [filter.page, filter.limit]) // eslint-disable-line react-hooks/exhaustive-deps

  const handleGetGalleryByUserId = async () => {
    setState((prev) => ({ ...prev, isLoading: true }))
    const response = await API.getAlbumImgsByUserId(filter)
    if (
      response.success &&
      response.status === StatusCode.Ok &&
      response.data.data
    ) {
      setState((prev) => ({
        ...prev,
        response: response.data,
        isLoading: false,
        isFetched: true,
        count: prev.count + 1,
      }))
    } else {
      // const message = response.errorMsg
      // showNotification({ color: 'red', message: message })
    }
  }

  const refetch = () => {
    handleGetGalleryByUserId()
  }

  const total = state.response?.data.total || 0
  const album = {
    ...state.response?.data,
    imgs:
      state.response?.data.imgs &&
      state.response?.data.imgs.map((el) => ({
        ...el,
        imgUrl: el.imgUrl + `?v${state.count}`,
      })),
  }

  const payload = {
    filter: filter,
    isLoading: state.isLoading,
    isFetched: state.isFetched,
    setQuery: setQuery,
    setPage: setPage,
    setLimit: setLimit,
    refetch: refetch,
    total: total,
    totalPage: Math.ceil((total > 0 ? total : 1) / filter.limit),
    album: album,
  }
  // console.log('useGetAlbumImgsByUserId', payload)
  return payload
}
