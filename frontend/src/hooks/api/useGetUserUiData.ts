import { useEffect, useState } from 'react'
import { isString } from 'lodash'
import { ISiteUi } from '@/entities/site'
import { API, IServerResponse, StatusCode } from '@/services'
import { getDefaultUiData } from '@/mocks'

export default function useGetUserUiData(userId: string) {
  const [state, setState] = useState<{
    responseData: ISiteUi | undefined
    isLoading: boolean
  }>({
    responseData: undefined,
    isLoading: true,
  })

  useEffect(() => {
    isString(userId) && handleGetUserSiteUiDataById(userId)
  }, [userId]) // eslint-disable-line react-hooks/exhaustive-deps

  const handleGetUserSiteUiDataById = async (uId: string) => {
    setState((prev) => ({ ...prev, isLoading: true }))
    const response = await API.getUserSiteUiDataById(uId)
    if (response.success && response.status === StatusCode.Ok) {
      setState((prev) => ({
        ...prev,
        responseData: response.data.data,
        isLoading: false,
      }))
      return
    } else if (response.status === StatusCode.NotFound) {
      setState((prev) => ({
        ...prev,
        responseData: getDefaultUiData(userId, userId),
        isLoading: false,
      }))
      return
    }
  }

  const refetch = () => {
    isString(userId) && handleGetUserSiteUiDataById(userId)
  }

  return {
    siteUiData: state.responseData,
    refetch: refetch,
  }
}
