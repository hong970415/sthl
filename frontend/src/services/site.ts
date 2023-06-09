import { ISiteUi } from '@/entities/site'
import { httpClient, IServerResponse } from './httpClient'

type IUpsertSiteUiForm = Pick<
  ISiteUi,
  'sitename' | 'homepageImgUrl' | 'homepageText' | 'homepageTextColor'
>

async function getUserSiteUiDataById(userId: string) {
  return httpClient.get<IServerResponse<ISiteUi>>(`/api/v1/siteui/${userId}`)
}
async function putUpsertUserSiteUiDataById(payload: IUpsertSiteUiForm) {
  return httpClient.putWithAuth<IServerResponse<ISiteUi>, IUpsertSiteUiForm>(
    '/api/v1/siteui',
    payload
  )
}

export { getUserSiteUiDataById, putUpsertUserSiteUiDataById }
